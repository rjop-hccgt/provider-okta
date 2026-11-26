package appoauthroleassignment

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_oauth_role_assignment", func(r *config.Resource) {
		r.References["client_id"] = config.Reference{
			TerraformName: "okta_app_oauth",
			Extractor:     "github.com/crossplane-contrib/provider-okta/apis/cluster/app/v1alpha1.OauthClientID()",
		}
	})
}
