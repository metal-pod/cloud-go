// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/fi-ts/cloud-go/api/client/accounting"
	"github.com/fi-ts/cloud-go/api/client/audit"
	"github.com/fi-ts/cloud-go/api/client/cli"
	"github.com/fi-ts/cloud-go/api/client/cluster"
	"github.com/fi-ts/cloud-go/api/client/database"
	"github.com/fi-ts/cloud-go/api/client/health"
	"github.com/fi-ts/cloud-go/api/client/ip"
	"github.com/fi-ts/cloud-go/api/client/masterdata"
	"github.com/fi-ts/cloud-go/api/client/project"
	"github.com/fi-ts/cloud-go/api/client/s3"
	"github.com/fi-ts/cloud-go/api/client/tenant"
	"github.com/fi-ts/cloud-go/api/client/version"
	"github.com/fi-ts/cloud-go/api/client/volume"
)

// Default cloud API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new cloud API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CloudAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cloud API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *CloudAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cloud API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CloudAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(CloudAPI)
	cli.Transport = transport
	cli.Accounting = accounting.New(transport, formats)
	cli.Audit = audit.New(transport, formats)
	cli.Cli = cli.New(transport, formats)
	cli.Cluster = cluster.New(transport, formats)
	cli.Database = database.New(transport, formats)
	cli.Health = health.New(transport, formats)
	cli.IP = ip.New(transport, formats)
	cli.Masterdata = masterdata.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.S3 = s3.New(transport, formats)
	cli.Tenant = tenant.New(transport, formats)
	cli.Version = version.New(transport, formats)
	cli.Volume = volume.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// CloudAPI is a client for cloud API
type CloudAPI struct {
	Accounting accounting.ClientService

	Audit audit.ClientService

	Cli cli.ClientService

	Cluster cluster.ClientService

	Database database.ClientService

	Health health.ClientService

	IP ip.ClientService

	Masterdata masterdata.ClientService

	Project project.ClientService

	S3 s3.ClientService

	Tenant tenant.ClientService

	Version version.ClientService

	Volume volume.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CloudAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Accounting.SetTransport(transport)
	c.Audit.SetTransport(transport)
	c.Cli.SetTransport(transport)
	c.Cluster.SetTransport(transport)
	c.Database.SetTransport(transport)
	c.Health.SetTransport(transport)
	c.IP.SetTransport(transport)
	c.Masterdata.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.S3.SetTransport(transport)
	c.Tenant.SetTransport(transport)
	c.Version.SetTransport(transport)
	c.Volume.SetTransport(transport)
}
