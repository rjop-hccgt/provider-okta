package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

// AuthServerID returns an extractor that returns the ID for a Server
func AuthServerID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		s, ok := any(mg).(*Server)
		if !ok {
			return ""
		}
		return s.GetID()
	}
}

// AuthServerPolicyID returns an extractor that returns the ID for a ServerPolicy
func AuthServerPolicyID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		s, ok := any(mg).(*ServerPolicy)
		if !ok {
			return ""
		}
		return s.GetID()
	}
}
