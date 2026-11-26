package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

// GroupID returns an extractor that returns the Group ID.
func GroupID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		u, ok := any(mg).(*Group)
		if !ok {
			return ""
		}
		return u.GetID()
	}
}
