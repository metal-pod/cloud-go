// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ConnectionInfo v1 connection info
//
// swagger:model v1.ConnectionInfo
type V1ConnectionInfo struct {

	// connected postgres ID
	ConnectedPostgresID string `json:"connectedPostgresID,omitempty"`

	// is replication primary
	IsReplicationPrimary bool `json:"isReplicationPrimary,omitempty"`

	// synchronous
	Synchronous bool `json:"synchronous,omitempty"`
}

// Validate validates this v1 connection info
func (m *V1ConnectionInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 connection info based on context it is used
func (m *V1ConnectionInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ConnectionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ConnectionInfo) UnmarshalBinary(b []byte) error {
	var res V1ConnectionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
