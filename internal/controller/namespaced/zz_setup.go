/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	roletargets "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/admin/roletargets"
	accesspolicyassignment "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/accesspolicyassignment"
	autologin "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/autologin"
	basicauth "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/basicauth"
	bookmark "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/bookmark"
	groupassignment "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/groupassignment"
	oauth "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/oauth"
	oauthapiscope "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/oauthapiscope"
	oauthroleassignment "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/oauthroleassignment"
	saml "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/saml"
	samlappsettings "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/samlappsettings"
	securepasswordstore "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/securepasswordstore"
	sharedcredentials "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/sharedcredentials"
	swa "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/swa"
	threefield "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/app/threefield"
	server "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/auth/server"
	serverclaim "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/auth/serverclaim"
	serverclaimdefault "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/auth/serverclaimdefault"
	serverdefault "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/auth/serverdefault"
	serverpolicy "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/auth/serverpolicy"
	serverpolicyrule "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/auth/serverpolicyrule"
	serverscope "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/auth/serverscope"
	group "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/group/group"
	role "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/group/role"
	rule "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/group/rule"
	memberships "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/groupmembership/memberships"
	oidc "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/idp/oidc"
	samlidp "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/idp/saml"
	behavior "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/okta/behavior"
	providerconfig "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/providerconfig"
	adminroles "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/user/adminroles"
	baseschemaproperty "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/user/baseschemaproperty"
	user "github.com/crossplane-contrib/provider-okta/internal/controller/namespaced/user/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		roletargets.Setup,
		accesspolicyassignment.Setup,
		autologin.Setup,
		basicauth.Setup,
		bookmark.Setup,
		groupassignment.Setup,
		oauth.Setup,
		oauthapiscope.Setup,
		oauthroleassignment.Setup,
		saml.Setup,
		samlappsettings.Setup,
		securepasswordstore.Setup,
		sharedcredentials.Setup,
		swa.Setup,
		threefield.Setup,
		server.Setup,
		serverclaim.Setup,
		serverclaimdefault.Setup,
		serverdefault.Setup,
		serverpolicy.Setup,
		serverpolicyrule.Setup,
		serverscope.Setup,
		group.Setup,
		role.Setup,
		rule.Setup,
		memberships.Setup,
		oidc.Setup,
		samlidp.Setup,
		behavior.Setup,
		providerconfig.Setup,
		adminroles.Setup,
		baseschemaproperty.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		roletargets.SetupGated,
		accesspolicyassignment.SetupGated,
		autologin.SetupGated,
		basicauth.SetupGated,
		bookmark.SetupGated,
		groupassignment.SetupGated,
		oauth.SetupGated,
		oauthapiscope.SetupGated,
		oauthroleassignment.SetupGated,
		saml.SetupGated,
		samlappsettings.SetupGated,
		securepasswordstore.SetupGated,
		sharedcredentials.SetupGated,
		swa.SetupGated,
		threefield.SetupGated,
		server.SetupGated,
		serverclaim.SetupGated,
		serverclaimdefault.SetupGated,
		serverdefault.SetupGated,
		serverpolicy.SetupGated,
		serverpolicyrule.SetupGated,
		serverscope.SetupGated,
		group.SetupGated,
		role.SetupGated,
		rule.SetupGated,
		memberships.SetupGated,
		oidc.SetupGated,
		samlidp.SetupGated,
		behavior.SetupGated,
		providerconfig.SetupGated,
		adminroles.SetupGated,
		baseschemaproperty.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		roletargets.SetupWebhookWithManager,
		accesspolicyassignment.SetupWebhookWithManager,
		autologin.SetupWebhookWithManager,
		basicauth.SetupWebhookWithManager,
		bookmark.SetupWebhookWithManager,
		groupassignment.SetupWebhookWithManager,
		oauth.SetupWebhookWithManager,
		oauthapiscope.SetupWebhookWithManager,
		oauthroleassignment.SetupWebhookWithManager,
		saml.SetupWebhookWithManager,
		samlappsettings.SetupWebhookWithManager,
		securepasswordstore.SetupWebhookWithManager,
		sharedcredentials.SetupWebhookWithManager,
		swa.SetupWebhookWithManager,
		threefield.SetupWebhookWithManager,
		server.SetupWebhookWithManager,
		serverclaim.SetupWebhookWithManager,
		serverclaimdefault.SetupWebhookWithManager,
		serverdefault.SetupWebhookWithManager,
		serverpolicy.SetupWebhookWithManager,
		serverpolicyrule.SetupWebhookWithManager,
		serverscope.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		rule.SetupWebhookWithManager,
		memberships.SetupWebhookWithManager,
		oidc.SetupWebhookWithManager,
		samlidp.SetupWebhookWithManager,
		behavior.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		adminroles.SetupWebhookWithManager,
		baseschemaproperty.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
