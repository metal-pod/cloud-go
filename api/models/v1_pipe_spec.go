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

// V1PipeSpec v1 pipe spec
//
// swagger:model v1.PipeSpec
type V1PipeSpec struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// peer
	Peer string `json:"peer,omitempty"`

	// port
	// Required: true
	Port *int64 `json:"port"`

	// remote
	// Required: true
	Remote *string `json:"remote"`
}

// Validate validates this v1 pipe spec
func (m *V1PipeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PipeSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1PipeSpec) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *V1PipeSpec) validateRemote(formats strfmt.Registry) error {

	if err := validate.Required("remote", "body", m.Remote); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 pipe spec based on context it is used
func (m *V1PipeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PipeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PipeSpec) UnmarshalBinary(b []byte) error {
	var res V1PipeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
