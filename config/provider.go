/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	adminroletargetsCluster "github.com/crossplane-contrib/provider-okta/config/cluster/adminroletargets"
	appaccesspolicyassignmentCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appaccesspolicyassignment"
	appautologinCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appautologin"
	appbasicauthCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appbasicauth"
	appbookmarkCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appbookmark"
	appgroupassignmentCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appgroupassignment"
	appoauthCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appoauth"
	appoauthapiscopeCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appoauthapiscope"
	appoauthroleassignmentCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appoauthroleassignment"
	appsamlCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appsaml"
	appsamlappsettingsCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appsamlappsettings"
	appsecurepasswordstoreCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appsecurepasswordstore"
	appsharedcredentialsCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appsharedcredentials"
	appswaCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appswa"
	appthreefieldCluster "github.com/crossplane-contrib/provider-okta/config/cluster/appthreefield"
	authserverCluster "github.com/crossplane-contrib/provider-okta/config/cluster/authserver"
	authserverclaimCluster "github.com/crossplane-contrib/provider-okta/config/cluster/authserverclaim"
	authserverclaimdefaultCluster "github.com/crossplane-contrib/provider-okta/config/cluster/authserverclaimdefault"
	authserverdefaultCluster "github.com/crossplane-contrib/provider-okta/config/cluster/authserverdefault"
	authserverpolicyCluster "github.com/crossplane-contrib/provider-okta/config/cluster/authserverpolicy"
	authserverpolicyruleCluster "github.com/crossplane-contrib/provider-okta/config/cluster/authserverpolicyrule"
	authserverscopeCluster "github.com/crossplane-contrib/provider-okta/config/cluster/authserverscope"
	behaviorCluster "github.com/crossplane-contrib/provider-okta/config/cluster/behavior"
	groupCluster "github.com/crossplane-contrib/provider-okta/config/cluster/group"
	groupmembershipsCluster "github.com/crossplane-contrib/provider-okta/config/cluster/groupmemberships"
	grouproleCluster "github.com/crossplane-contrib/provider-okta/config/cluster/grouprole"
	groupruleCluster "github.com/crossplane-contrib/provider-okta/config/cluster/grouprule"
	idpoidcCluster "github.com/crossplane-contrib/provider-okta/config/cluster/idpoidc"
	idpsamlCluster "github.com/crossplane-contrib/provider-okta/config/cluster/idpsaml"
	userCluster "github.com/crossplane-contrib/provider-okta/config/cluster/user"
	useradminrolesCluster "github.com/crossplane-contrib/provider-okta/config/cluster/useradminroles"
	userbaseschemapropertyCluster "github.com/crossplane-contrib/provider-okta/config/cluster/userbaseschemaproperty"
	adminroletargetsNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/adminroletargets"
	appaccesspolicyassignmentNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appaccesspolicyassignment"
	appautologinNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appautologin"
	appbasicauthNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appbasicauth"
	appbookmarkNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appbookmark"
	appgroupassignmentNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appgroupassignment"
	appoauthNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appoauth"
	appoauthapiscopeNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appoauthapiscope"
	appoauthroleassignmentNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appoauthroleassignment"
	appsamlNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appsaml"
	appsamlappsettingsNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appsamlappsettings"
	appsecurepasswordstoreNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appsecurepasswordstore"
	appsharedcredentialsNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appsharedcredentials"
	appswaNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appswa"
	appthreefieldNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/appthreefield"
	authserverNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/authserver"
	authserverclaimNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/authserverclaim"
	authserverclaimdefaultNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/authserverclaimdefault"
	authserverdefaultNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/authserverdefault"
	authserverpolicyNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/authserverpolicy"
	authserverpolicyruleNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/authserverpolicyrule"
	authserverscopeNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/authserverscope"
	behaviorNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/behavior"
	groupNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/group"
	groupmembershipsNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/groupmemberships"
	grouproleNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/grouprole"
	groupruleNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/grouprule"
	idpoidcNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/idpoidc"
	idpsamlNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/idpsaml"
	userNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/user"
	useradminrolesNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/useradminroles"
	userbaseschemapropertyNamespaced "github.com/crossplane-contrib/provider-okta/config/namespaced/userbaseschemaproperty"
)

const (
	resourcePrefix = "okta"
	modulePath     = "github.com/crossplane-contrib/provider-okta"

	rootGroupCluster    = "okta.upbound.io"
	rootGroupNamespaced = "okta.m.upbound.io"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns the cluster-scoped provider configuration.
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroupCluster),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		adminroletargetsCluster.Configure,
		appaccesspolicyassignmentCluster.Configure,
		appautologinCluster.Configure,
		appbasicauthCluster.Configure,
		appbookmarkCluster.Configure,
		appgroupassignmentCluster.Configure,
		appoauthCluster.Configure,
		appoauthapiscopeCluster.Configure,
		appoauthroleassignmentCluster.Configure,
		appsamlCluster.Configure,
		appsamlappsettingsCluster.Configure,
		appsecurepasswordstoreCluster.Configure,
		appsharedcredentialsCluster.Configure,
		appswaCluster.Configure,
		appthreefieldCluster.Configure,
		authserverCluster.Configure,
		authserverclaimCluster.Configure,
		authserverclaimdefaultCluster.Configure,
		authserverdefaultCluster.Configure,
		authserverpolicyCluster.Configure,
		authserverpolicyruleCluster.Configure,
		authserverscopeCluster.Configure,
		behaviorCluster.Configure,
		groupCluster.Configure,
		groupmembershipsCluster.Configure,
		grouproleCluster.Configure,
		groupruleCluster.Configure,
		idpoidcCluster.Configure,
		idpsamlCluster.Configure,
		userCluster.Configure,
		useradminrolesCluster.Configure,
		userbaseschemapropertyCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration.
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroupNamespaced),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		adminroletargetsNamespaced.Configure,
		appaccesspolicyassignmentNamespaced.Configure,
		appautologinNamespaced.Configure,
		appbasicauthNamespaced.Configure,
		appbookmarkNamespaced.Configure,
		appgroupassignmentNamespaced.Configure,
		appoauthNamespaced.Configure,
		appoauthapiscopeNamespaced.Configure,
		appoauthroleassignmentNamespaced.Configure,
		appsamlNamespaced.Configure,
		appsamlappsettingsNamespaced.Configure,
		appsecurepasswordstoreNamespaced.Configure,
		appsharedcredentialsNamespaced.Configure,
		appswaNamespaced.Configure,
		appthreefieldNamespaced.Configure,
		authserverNamespaced.Configure,
		authserverclaimNamespaced.Configure,
		authserverclaimdefaultNamespaced.Configure,
		authserverdefaultNamespaced.Configure,
		authserverpolicyNamespaced.Configure,
		authserverpolicyruleNamespaced.Configure,
		authserverscopeNamespaced.Configure,
		behaviorNamespaced.Configure,
		groupNamespaced.Configure,
		groupmembershipsNamespaced.Configure,
		grouproleNamespaced.Configure,
		groupruleNamespaced.Configure,
		idpoidcNamespaced.Configure,
		idpsamlNamespaced.Configure,
		userNamespaced.Configure,
		useradminrolesNamespaced.Configure,
		userbaseschemapropertyNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
