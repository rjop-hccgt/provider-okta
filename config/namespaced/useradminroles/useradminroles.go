package useradminroles

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_user_admin_roles", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.References["user_id"] = config.Reference{
			TerraformName: "okta_user",
			Extractor:     "github.com/crossplane-contrib/provider-okta/apis/namespaced/user/v1alpha1.UserID()",
		}
	})
}
