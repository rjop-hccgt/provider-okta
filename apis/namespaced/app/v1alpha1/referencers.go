package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

// OauthClientID  returns an extractor that returns the ClientID for an Oauth APP
func OauthClientID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		o, ok := any(mg).(*Oauth)
		if !ok {
			return ""
		}
		if o.Status.AtProvider.ClientID == nil {
			return ""
		}
		return *o.Status.AtProvider.ClientID
	}
}

// SamlID returns an extractor that returns the ID for a SAML app
func SamlID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		s, ok := any(mg).(*SAML)
		if !ok {
			return ""
		}
		return s.GetID()
	}
}
