package userbaseschemaproperty

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_user_base_schema_property", func(r *config.Resource) {
	})
}
