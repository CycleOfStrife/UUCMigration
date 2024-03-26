// Code generated by go generate; DO NOT EDIT.
package kube

import (
	"net"
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *PlayOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *PlayOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAnnotations set field Annotations to given value
func (o *PlayOptions) WithAnnotations(value map[string]string) *PlayOptions {
	o.Annotations = value
	return o
}

// GetAnnotations returns value of field Annotations
func (o *PlayOptions) GetAnnotations() map[string]string {
	if o.Annotations == nil {
		var z map[string]string
		return z
	}
	return o.Annotations
}

// WithAuthfile set field Authfile to given value
func (o *PlayOptions) WithAuthfile(value string) *PlayOptions {
	o.Authfile = &value
	return o
}

// GetAuthfile returns value of field Authfile
func (o *PlayOptions) GetAuthfile() string {
	if o.Authfile == nil {
		var z string
		return z
	}
	return *o.Authfile
}

// WithCertDir set field CertDir to given value
func (o *PlayOptions) WithCertDir(value string) *PlayOptions {
	o.CertDir = &value
	return o
}

// GetCertDir returns value of field CertDir
func (o *PlayOptions) GetCertDir() string {
	if o.CertDir == nil {
		var z string
		return z
	}
	return *o.CertDir
}

// WithUsername set field Username to given value
func (o *PlayOptions) WithUsername(value string) *PlayOptions {
	o.Username = &value
	return o
}

// GetUsername returns value of field Username
func (o *PlayOptions) GetUsername() string {
	if o.Username == nil {
		var z string
		return z
	}
	return *o.Username
}

// WithPassword set field Password to given value
func (o *PlayOptions) WithPassword(value string) *PlayOptions {
	o.Password = &value
	return o
}

// GetPassword returns value of field Password
func (o *PlayOptions) GetPassword() string {
	if o.Password == nil {
		var z string
		return z
	}
	return *o.Password
}

// WithNetwork set field Network to given value
func (o *PlayOptions) WithNetwork(value []string) *PlayOptions {
	o.Network = &value
	return o
}

// GetNetwork returns value of field Network
func (o *PlayOptions) GetNetwork() []string {
	if o.Network == nil {
		var z []string
		return z
	}
	return *o.Network
}

// WithNoHosts set field NoHosts to given value
func (o *PlayOptions) WithNoHosts(value bool) *PlayOptions {
	o.NoHosts = &value
	return o
}

// GetNoHosts returns value of field NoHosts
func (o *PlayOptions) GetNoHosts() bool {
	if o.NoHosts == nil {
		var z bool
		return z
	}
	return *o.NoHosts
}

// WithQuiet set field Quiet to given value
func (o *PlayOptions) WithQuiet(value bool) *PlayOptions {
	o.Quiet = &value
	return o
}

// GetQuiet returns value of field Quiet
func (o *PlayOptions) GetQuiet() bool {
	if o.Quiet == nil {
		var z bool
		return z
	}
	return *o.Quiet
}

// WithSignaturePolicy set field SignaturePolicy to given value
func (o *PlayOptions) WithSignaturePolicy(value string) *PlayOptions {
	o.SignaturePolicy = &value
	return o
}

// GetSignaturePolicy returns value of field SignaturePolicy
func (o *PlayOptions) GetSignaturePolicy() string {
	if o.SignaturePolicy == nil {
		var z string
		return z
	}
	return *o.SignaturePolicy
}

// WithSkipTLSVerify set field SkipTLSVerify to given value
func (o *PlayOptions) WithSkipTLSVerify(value bool) *PlayOptions {
	o.SkipTLSVerify = &value
	return o
}

// GetSkipTLSVerify returns value of field SkipTLSVerify
func (o *PlayOptions) GetSkipTLSVerify() bool {
	if o.SkipTLSVerify == nil {
		var z bool
		return z
	}
	return *o.SkipTLSVerify
}

// WithSeccompProfileRoot set field SeccompProfileRoot to given value
func (o *PlayOptions) WithSeccompProfileRoot(value string) *PlayOptions {
	o.SeccompProfileRoot = &value
	return o
}

// GetSeccompProfileRoot returns value of field SeccompProfileRoot
func (o *PlayOptions) GetSeccompProfileRoot() string {
	if o.SeccompProfileRoot == nil {
		var z string
		return z
	}
	return *o.SeccompProfileRoot
}

// WithStaticIPs set field StaticIPs to given value
func (o *PlayOptions) WithStaticIPs(value []net.IP) *PlayOptions {
	o.StaticIPs = &value
	return o
}

// GetStaticIPs returns value of field StaticIPs
func (o *PlayOptions) GetStaticIPs() []net.IP {
	if o.StaticIPs == nil {
		var z []net.IP
		return z
	}
	return *o.StaticIPs
}

// WithStaticMACs set field StaticMACs to given value
func (o *PlayOptions) WithStaticMACs(value []net.HardwareAddr) *PlayOptions {
	o.StaticMACs = &value
	return o
}

// GetStaticMACs returns value of field StaticMACs
func (o *PlayOptions) GetStaticMACs() []net.HardwareAddr {
	if o.StaticMACs == nil {
		var z []net.HardwareAddr
		return z
	}
	return *o.StaticMACs
}

// WithConfigMaps set field ConfigMaps to given value
func (o *PlayOptions) WithConfigMaps(value []string) *PlayOptions {
	o.ConfigMaps = &value
	return o
}

// GetConfigMaps returns value of field ConfigMaps
func (o *PlayOptions) GetConfigMaps() []string {
	if o.ConfigMaps == nil {
		var z []string
		return z
	}
	return *o.ConfigMaps
}

// WithLogDriver set field LogDriver to given value
func (o *PlayOptions) WithLogDriver(value string) *PlayOptions {
	o.LogDriver = &value
	return o
}

// GetLogDriver returns value of field LogDriver
func (o *PlayOptions) GetLogDriver() string {
	if o.LogDriver == nil {
		var z string
		return z
	}
	return *o.LogDriver
}

// WithLogOptions set field LogOptions to given value
func (o *PlayOptions) WithLogOptions(value []string) *PlayOptions {
	o.LogOptions = &value
	return o
}

// GetLogOptions returns value of field LogOptions
func (o *PlayOptions) GetLogOptions() []string {
	if o.LogOptions == nil {
		var z []string
		return z
	}
	return *o.LogOptions
}

// WithReplace set field Replace to given value
func (o *PlayOptions) WithReplace(value bool) *PlayOptions {
	o.Replace = &value
	return o
}

// GetReplace returns value of field Replace
func (o *PlayOptions) GetReplace() bool {
	if o.Replace == nil {
		var z bool
		return z
	}
	return *o.Replace
}

// WithStart set field Start to given value
func (o *PlayOptions) WithStart(value bool) *PlayOptions {
	o.Start = &value
	return o
}

// GetStart returns value of field Start
func (o *PlayOptions) GetStart() bool {
	if o.Start == nil {
		var z bool
		return z
	}
	return *o.Start
}

// WithNoTrunc set field NoTrunc to given value
func (o *PlayOptions) WithNoTrunc(value bool) *PlayOptions {
	o.NoTrunc = &value
	return o
}

// GetNoTrunc returns value of field NoTrunc
func (o *PlayOptions) GetNoTrunc() bool {
	if o.NoTrunc == nil {
		var z bool
		return z
	}
	return *o.NoTrunc
}

// WithUserns set field Userns to given value
func (o *PlayOptions) WithUserns(value string) *PlayOptions {
	o.Userns = &value
	return o
}

// GetUserns returns value of field Userns
func (o *PlayOptions) GetUserns() string {
	if o.Userns == nil {
		var z string
		return z
	}
	return *o.Userns
}

// WithForce set field Force to given value
func (o *PlayOptions) WithForce(value bool) *PlayOptions {
	o.Force = &value
	return o
}

// GetForce returns value of field Force
func (o *PlayOptions) GetForce() bool {
	if o.Force == nil {
		var z bool
		return z
	}
	return *o.Force
}

// WithPublishPorts set field PublishPorts to given value
func (o *PlayOptions) WithPublishPorts(value []string) *PlayOptions {
	o.PublishPorts = value
	return o
}

// GetPublishPorts returns value of field PublishPorts
func (o *PlayOptions) GetPublishPorts() []string {
	if o.PublishPorts == nil {
		var z []string
		return z
	}
	return o.PublishPorts
}

// WithPublishAllPorts set field PublishAllPorts to given value
func (o *PlayOptions) WithPublishAllPorts(value bool) *PlayOptions {
	o.PublishAllPorts = &value
	return o
}

// GetPublishAllPorts returns value of field PublishAllPorts
func (o *PlayOptions) GetPublishAllPorts() bool {
	if o.PublishAllPorts == nil {
		var z bool
		return z
	}
	return *o.PublishAllPorts
}

// WithWait set field Wait to given value
func (o *PlayOptions) WithWait(value bool) *PlayOptions {
	o.Wait = &value
	return o
}

// GetWait returns value of field Wait
func (o *PlayOptions) GetWait() bool {
	if o.Wait == nil {
		var z bool
		return z
	}
	return *o.Wait
}

// WithServiceContainer set field ServiceContainer to given value
func (o *PlayOptions) WithServiceContainer(value bool) *PlayOptions {
	o.ServiceContainer = &value
	return o
}

// GetServiceContainer returns value of field ServiceContainer
func (o *PlayOptions) GetServiceContainer() bool {
	if o.ServiceContainer == nil {
		var z bool
		return z
	}
	return *o.ServiceContainer
}