// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ClusterUsageResponse v1 cluster usage response
//
// swagger:model v1.ClusterUsageResponse
type V1ClusterUsageResponse struct {

	// just the usage data of the individual clusters summed up
	// Required: true
	Accumulatedusage *V1ClusterUsageAccumuluated `json:"accumulatedusage" yaml:"accumulatedusage"`

	// the start time in the accounting window to look at
	// Required: true
	// Format: date-time
	From *strfmt.DateTime `json:"from" yaml:"from"`

	// the end time in the accounting window to look at (defaults to current system time)
	// Format: date-time
	To strfmt.DateTime `json:"to,omitempty" yaml:"to,omitempty"`

	// the usage data of the individual clusters
	// Required: true
	Usage []*V1ClusterUsage `json:"usage" yaml:"usage"`
}

// Validate validates this v1 cluster usage response
func (m *V1ClusterUsageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccumulatedusage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterUsageResponse) validateAccumulatedusage(formats strfmt.Registry) error {

	if err := validate.Required("accumulatedusage", "body", m.Accumulatedusage); err != nil {
		return err
	}

	if m.Accumulatedusage != nil {
		if err := m.Accumulatedusage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accumulatedusage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accumulatedusage")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUsageResponse) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsageResponse) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsageResponse) validateUsage(formats strfmt.Registry) error {

	if err := validate.Required("usage", "body", m.Usage); err != nil {
		return err
	}

	for i := 0; i < len(m.Usage); i++ {
		if swag.IsZero(m.Usage[i]) { // not required
			continue
		}

		if m.Usage[i] != nil {
			if err := m.Usage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 cluster usage response based on the context it is used
func (m *V1ClusterUsageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccumulatedusage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterUsageResponse) contextValidateAccumulatedusage(ctx context.Context, formats strfmt.Registry) error {

	if m.Accumulatedusage != nil {
		if err := m.Accumulatedusage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accumulatedusage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accumulatedusage")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterUsageResponse) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Usage); i++ {

		if m.Usage[i] != nil {
			if err := m.Usage[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterUsageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterUsageResponse) UnmarshalBinary(b []byte) error {
	var res V1ClusterUsageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
