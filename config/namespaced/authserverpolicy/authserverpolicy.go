package authserverpolicy

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_auth_server_policy", func(r *config.Resource) {
		r.References["auth_server_id"] = config.Reference{
			TerraformName: "okta_auth_server",
			Extractor:     "github.com/crossplane-contrib/provider-okta/apis/namespaced/auth/v1alpha1.AuthServerID()",
		}
	})
}
