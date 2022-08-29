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

// V1NetworkUsageAccumuluated v1 network usage accumuluated
//
// swagger:model v1.NetworkUsageAccumuluated
type V1NetworkUsageAccumuluated struct {

	// the accumuluated ingoing traffic (byte)
	// Required: true
	In *string `json:"in" yaml:"in"`

	// the duration for that the network usage is accounted
	// Required: true
	Lifetime *int64 `json:"lifetime" yaml:"lifetime"`

	// the accumulated outgoing traffic (byte)
	// Required: true
	Out *string `json:"out" yaml:"out"`

	// the accumulated total traffic (byte)
	// Required: true
	Total *string `json:"total" yaml:"total"`
}

// Validate validates this v1 network usage accumuluated
func (m *V1NetworkUsageAccumuluated) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOut(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkUsageAccumuluated) validateIn(formats strfmt.Registry) error {

	if err := validate.Required("in", "body", m.In); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsageAccumuluated) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsageAccumuluated) validateOut(formats strfmt.Registry) error {

	if err := validate.Required("out", "body", m.Out); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsageAccumuluated) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 network usage accumuluated based on context it is used
func (m *V1NetworkUsageAccumuluated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkUsageAccumuluated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkUsageAccumuluated) UnmarshalBinary(b []byte) error {
	var res V1NetworkUsageAccumuluated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
