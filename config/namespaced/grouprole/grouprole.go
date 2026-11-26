package grouprole

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_group_role", func(r *config.Resource) {
		r.ShortGroup = "group"
		r.References["group_id"] = config.Reference{
			TerraformName: "okta_group",
			Extractor:     "github.com/crossplane-contrib/provider-okta/apis/namespaced/group/v1alpha1.GroupID()",
		}
	})
}
