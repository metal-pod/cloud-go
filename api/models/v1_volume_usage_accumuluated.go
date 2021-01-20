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

// V1VolumeUsageAccumuluated v1 volume usage accumuluated
//
// swagger:model v1.VolumeUsageAccumuluated
type V1VolumeUsageAccumuluated struct {

	// the accumulated capacity seconds of the volumes in this response (byte*s)
	// Required: true
	Capacityseconds *string `json:"capacityseconds"`

	// the duration that this volume is running
	// Required: true
	Lifetime *int64 `json:"lifetime"`
}

// Validate validates this v1 volume usage accumuluated
func (m *V1VolumeUsageAccumuluated) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacityseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VolumeUsageAccumuluated) validateCapacityseconds(formats strfmt.Registry) error {

	if err := validate.Required("capacityseconds", "body", m.Capacityseconds); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeUsageAccumuluated) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 volume usage accumuluated based on context it is used
func (m *V1VolumeUsageAccumuluated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VolumeUsageAccumuluated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VolumeUsageAccumuluated) UnmarshalBinary(b []byte) error {
	var res V1VolumeUsageAccumuluated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
