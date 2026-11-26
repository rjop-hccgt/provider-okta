package appgroupassignment

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_group_assignment", func(r *config.Resource) {
		r.References["group_id"] = config.Reference{
			TerraformName: "okta_group",
			Extractor:     "github.com/crossplane-contrib/provider-okta/apis/namespaced/group/v1alpha1.GroupID()",
		}
	})
}
