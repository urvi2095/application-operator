package apis

import (
	"github.com/urvi2095/application-operator/pkg/apis/application/v2alpha2"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v2alpha2.SchemeBuilder.AddToScheme)
}
