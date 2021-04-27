// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1GatewayDeleteRequest v1 gateway delete request
//
// swagger:model v1.GatewayDeleteRequest
type V1GatewayDeleteRequest struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// project UID
	// Required: true
	ProjectUID *string `json:"projectUID"`
}

// Validate validates this v1 gateway delete request
func (m *V1GatewayDeleteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GatewayDeleteRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1GatewayDeleteRequest) validateProjectUID(formats strfmt.Registry) error {

	if err := validate.Required("projectUID", "body", m.ProjectUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 gateway delete request based on context it is used
func (m *V1GatewayDeleteRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GatewayDeleteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GatewayDeleteRequest) UnmarshalBinary(b []byte) error {
	var res V1GatewayDeleteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
