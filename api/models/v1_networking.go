// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Networking v1 networking
// swagger:model v1.Networking
type V1Networking struct {

	// nodes
	// Required: true
	Nodes *string `json:"Nodes"`

	// pods
	// Required: true
	Pods *string `json:"Pods"`

	// services
	// Required: true
	Services *string `json:"Services"`

	// type
	// Required: true
	Type *string `json:"Type"`
}

// Validate validates this v1 networking
func (m *V1Networking) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Networking) validateNodes(formats strfmt.Registry) error {

	if err := validate.Required("Nodes", "body", m.Nodes); err != nil {
		return err
	}

	return nil
}

func (m *V1Networking) validatePods(formats strfmt.Registry) error {

	if err := validate.Required("Pods", "body", m.Pods); err != nil {
		return err
	}

	return nil
}

func (m *V1Networking) validateServices(formats strfmt.Registry) error {

	if err := validate.Required("Services", "body", m.Services); err != nil {
		return err
	}

	return nil
}

func (m *V1Networking) validateType(formats strfmt.Registry) error {

	if err := validate.Required("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Networking) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Networking) UnmarshalBinary(b []byte) error {
	var res V1Networking
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
