/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/crossplane-contrib/provider-okta/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/crossplane-contrib/provider-okta/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal okta credentials as JSON"
	errResolveProviderCfg   = "cannot resolve provider config"
	errUnknownPCKind        = "unknown GVK for ProviderConfig"
	errNotManaged           = "resource is not a managed resource"
	errUnknownPCType        = "unknown provider config type"
	errNotObject            = "provider config is not a client.Object"
)

// credentialKeys are the keys read from the credentials secret and passed
// verbatim to the Okta Terraform provider configuration.
var credentialKeys = []string{
	"org_name",
	"base_url",
	"http_proxy",
	"access_token",
	"api_token",
	"client_id",
	"scopes",
	"private_key",
	"private_key_id",
	"backoff",
	"min_wait_seconds",
	"max_wait_seconds",
	"max_retries",
	"request_timeout",
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, errResolveProviderCfg)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}
		ps.Configuration = credentialsToConfiguration(creds)
		return ps, nil
	}
}

// credentialsToConfiguration copies the supported Okta keys from the
// credentials into the Terraform provider configuration. Unknown keys are
// ignored.
func credentialsToConfiguration(creds map[string]string) map[string]any {
	cfg := map[string]any{}
	for _, k := range credentialKeys {
		if v, ok := creds[k]; ok {
			cfg[k] = v
		}
	}
	return cfg
}

func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}
	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged: //nolint:staticcheck // still handling cluster-scoped behavior
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New(errNotManaged)
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) { //nolint:staticcheck // still handling cluster-scoped behavior
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, errUnknownPCKind)
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		return nil, errors.New(errNotObject)
	}

	key := types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}
	if _, clusterScoped := pcObj.(*namespacedv1beta1.ClusterProviderConfig); clusterScoped {
		key.Namespace = ""
	}
	if err := crClient.Get(ctx, key, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		// A namespaced ProviderConfig may only read secrets from the
		// managed resource's own namespace.
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New(errUnknownPCType)
	}
	t := resource.NewProviderConfigUsageTracker(crClient, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
