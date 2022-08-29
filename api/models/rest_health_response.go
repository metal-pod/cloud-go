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

// RestHealthResponse rest health response
//
// swagger:model rest.HealthResponse
type RestHealthResponse struct {

	// message
	// Required: true
	Message *string `json:"message" yaml:"message"`

	// services
	// Required: true
	Services map[string]RestHealthResult `json:"services" yaml:"services"`

	// status
	// Required: true
	Status *string `json:"status" yaml:"status"`
}

// Validate validates this rest health response
func (m *RestHealthResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestHealthResponse) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *RestHealthResponse) validateServices(formats strfmt.Registry) error {

	if err := validate.Required("services", "body", m.Services); err != nil {
		return err
	}

	for k := range m.Services {

		if err := validate.Required("services"+"."+k, "body", m.Services[k]); err != nil {
			return err
		}
		if val, ok := m.Services[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestHealthResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rest health response based on the context it is used
func (m *RestHealthResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestHealthResponse) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("services", "body", m.Services); err != nil {
		return err
	}

	for k := range m.Services {

		if val, ok := m.Services[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestHealthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestHealthResponse) UnmarshalBinary(b []byte) error {
	var res RestHealthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
