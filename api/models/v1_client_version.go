// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClientVersion v1 client version
//
// swagger:model v1.ClientVersion
type V1ClientVersion struct {

	// min version
	MinVersion string `json:"min_version,omitempty"`
}

// Validate validates this v1 client version
func (m *V1ClientVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 client version based on context it is used
func (m *V1ClientVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClientVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClientVersion) UnmarshalBinary(b []byte) error {
	var res V1ClientVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
