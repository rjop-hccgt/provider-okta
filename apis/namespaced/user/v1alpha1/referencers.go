package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

// UserID returns an extractor that returns the User ID.
func UserID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		u, ok := any(mg).(*User)
		if !ok {
			return ""
		}
		return u.GetID()
	}
}
