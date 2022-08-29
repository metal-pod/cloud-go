// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Restore v1 restore
//
// swagger:model v1.Restore
type V1Restore struct {

	// postgres ID
	PostgresID string `json:"postgresID,omitempty" yaml:"postgresID,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
}

// Validate validates this v1 restore
func (m *V1Restore) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 restore based on context it is used
func (m *V1Restore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Restore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Restore) UnmarshalBinary(b []byte) error {
	var res V1Restore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
