// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *KillOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *KillOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithSignal set field Signal to given value
func (o *KillOptions) WithSignal(value string) *KillOptions {
	o.Signal = &value
	return o
}

// GetSignal returns value of field Signal
func (o *KillOptions) GetSignal() string {
	if o.Signal == nil {
		var z string
		return z
	}
	return *o.Signal
}
