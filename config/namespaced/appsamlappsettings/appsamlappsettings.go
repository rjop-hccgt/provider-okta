package appsamlappsettings

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_saml_app_settings", func(r *config.Resource) {
		r.References["app_id"] = config.Reference{
			TerraformName: "okta_app_saml",
			Extractor:     "github.com/crossplane-contrib/provider-okta/apis/namespaced/app/v1alpha1.SamlID()",
		}
	})
}
