/*
Copyright 2025 The Crossplane Authors.
*/

package clients

import (
	"context"
	"testing"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	apisCluster "github.com/crossplane-contrib/provider-okta/apis/cluster"
	clustergroup "github.com/crossplane-contrib/provider-okta/apis/cluster/group/v1alpha1"
	clusterv1beta1 "github.com/crossplane-contrib/provider-okta/apis/cluster/v1beta1"
	apisNamespaced "github.com/crossplane-contrib/provider-okta/apis/namespaced"
	nsgroup "github.com/crossplane-contrib/provider-okta/apis/namespaced/group/v1alpha1"
	namespacedv1beta1 "github.com/crossplane-contrib/provider-okta/apis/namespaced/v1beta1"
)

const (
	systemNS = "crossplane-system"
	teamNS   = "team-a"
	credsKey = "credentials"
	pcName   = "default"
	credsDoc = `{"org_name":"example","base_url":"okta.com","api_token":"t0k3n","unknown":"ignored"}`
)

func testScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	s := runtime.NewScheme()
	for _, add := range []func(*runtime.Scheme) error{clientgoscheme.AddToScheme, apisCluster.AddToScheme, apisNamespaced.AddToScheme} {
		if err := add(s); err != nil {
			t.Fatalf("cannot build scheme: %v", err)
		}
	}
	return s
}

func secret(ns string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: "okta-creds", Namespace: ns},
		Data:       map[string][]byte{credsKey: []byte(credsDoc)},
	}
}

func secretSelector(ns string) xpv2.CommonCredentialSelectors {
	return xpv2.CommonCredentialSelectors{
		SecretRef: &xpv2.SecretKeySelector{
			SecretReference: xpv2.SecretReference{Name: "okta-creds", Namespace: ns},
			Key:             credsKey,
		},
	}
}

func spec(secretNS string) clusterv1beta1.ProviderConfigSpec {
	return clusterv1beta1.ProviderConfigSpec{Credentials: clusterv1beta1.ProviderCredentials{
		Source: xpv2.CredentialsSourceSecret, CommonCredentialSelectors: secretSelector(secretNS),
	}}
}

func nsSpec(secretNS string) namespacedv1beta1.ProviderConfigSpec {
	return namespacedv1beta1.ProviderConfigSpec{Credentials: namespacedv1beta1.ProviderCredentials{
		Source: xpv2.CredentialsSourceSecret, CommonCredentialSelectors: secretSelector(secretNS),
	}}
}

func legacyGroup(ref *xpv2.Reference) *clustergroup.Group {
	g := &clustergroup.Group{ObjectMeta: metav1.ObjectMeta{Name: "g", UID: types.UID("uid-legacy")}}
	g.SetGroupVersionKind(clustergroup.Group_GroupVersionKind)
	g.Spec.ProviderConfigReference = ref
	return g
}

func modernGroup(ref *xpv2.ProviderConfigReference) *nsgroup.Group {
	g := &nsgroup.Group{ObjectMeta: metav1.ObjectMeta{Name: "g", Namespace: teamNS, UID: types.UID("uid-modern")}}
	g.SetGroupVersionKind(nsgroup.Group_GroupVersionKind)
	g.Spec.ProviderConfigReference = ref
	return g
}

func TestTerraformSetupBuilder(t *testing.T) {
	wantCfg := map[string]any{"org_name": "example", "base_url": "okta.com", "api_token": "t0k3n"}

	cases := map[string]struct {
		reason  string
		objs    []client.Object
		mg      resource.Managed
		wantErr bool
		wantCfg map[string]any
	}{
		"LegacyClusterProviderConfig": {
			reason: "A cluster-scoped MR resolves the legacy cluster ProviderConfig.",
			objs: []client.Object{
				&clusterv1beta1.ProviderConfig{ObjectMeta: metav1.ObjectMeta{Name: pcName}, Spec: spec(systemNS)},
				secret(systemNS),
			},
			mg:      legacyGroup(&xpv2.Reference{Name: pcName}),
			wantCfg: wantCfg,
		},
		"NamespacedProviderConfigUsesMRNamespace": {
			reason: "A namespaced ProviderConfig may only read the secret from the MR's namespace, whatever namespace its secretRef names.",
			objs: []client.Object{
				&namespacedv1beta1.ProviderConfig{ObjectMeta: metav1.ObjectMeta{Name: pcName, Namespace: teamNS}, Spec: nsSpec("some-other-ns")},
				secret(teamNS),
			},
			mg:      modernGroup(&xpv2.ProviderConfigReference{Kind: "ProviderConfig", Name: pcName}),
			wantCfg: wantCfg,
		},
		"ClusterProviderConfigKeepsSecretNamespace": {
			reason: "A ClusterProviderConfig reads the secret from the namespace its secretRef names.",
			objs: []client.Object{
				&namespacedv1beta1.ClusterProviderConfig{ObjectMeta: metav1.ObjectMeta{Name: pcName}, Spec: nsSpec(systemNS)},
				secret(systemNS),
			},
			mg:      modernGroup(&xpv2.ProviderConfigReference{Kind: "ClusterProviderConfig", Name: pcName}),
			wantCfg: wantCfg,
		},
		"UnknownKind": {
			reason:  "An unknown providerConfigRef kind is an error.",
			mg:      modernGroup(&xpv2.ProviderConfigReference{Kind: "NotAProviderConfig", Name: pcName}),
			wantErr: true,
		},
		"NilLegacyRef": {
			reason:  "A cluster-scoped MR without a providerConfigRef is an error.",
			mg:      legacyGroup(nil),
			wantErr: true,
		},
		"NilModernRef": {
			reason:  "A namespaced MR without a providerConfigRef is an error.",
			mg:      modernGroup(nil),
			wantErr: true,
		},
		"MissingProviderConfig": {
			reason:  "A reference to a ProviderConfig that does not exist is an error.",
			mg:      legacyGroup(&xpv2.Reference{Name: "missing"}),
			wantErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := fake.NewClientBuilder().WithScheme(testScheme(t)).WithObjects(tc.objs...).Build()
			fn := TerraformSetupBuilder("1.5.7", "okta/okta", "4.6.3")

			got, err := fn(context.Background(), c, tc.mg)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("%s\nTerraformSetupBuilder(...): want error, got nil", tc.reason)
				}
				return
			}
			if err != nil {
				t.Fatalf("%s\nTerraformSetupBuilder(...): unexpected error: %v", tc.reason, err)
			}
			if diff := cmp.Diff(tc.wantCfg, map[string]any(got.Configuration)); diff != "" {
				t.Errorf("%s\nTerraformSetupBuilder(...): -want configuration, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestCredentialsToConfiguration(t *testing.T) {
	creds := map[string]string{"unknown": "x"}
	want := map[string]any{}
	for _, k := range credentialKeys {
		creds[k] = "v-" + k
		want[k] = "v-" + k
	}
	if diff := cmp.Diff(want, credentialsToConfiguration(creds)); diff != "" {
		t.Errorf("credentialsToConfiguration(...): -want, +got:\n%s", diff)
	}
	if diff := cmp.Diff(map[string]any{}, credentialsToConfiguration(nil)); diff != "" {
		t.Errorf("credentialsToConfiguration(nil): -want, +got:\n%s", diff)
	}
}
