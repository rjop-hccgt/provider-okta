package adminroletargets

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_admin_role_targets", func(r *config.Resource) {
		r.References["user_id"] = config.Reference{
			TerraformName: "okta_user",
			Extractor:     "github.com/crossplane-contrib/provider-okta/apis/cluster/user/v1alpha1.UserID()",
		}
	})
}
