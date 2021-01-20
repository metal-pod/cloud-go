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

// V1IPUsage v1 IP usage
//
// swagger:model v1.IPUsage
type V1IPUsage struct {

	// the end time of this ip
	// Required: true
	// Format: date-time
	End *strfmt.DateTime `json:"end"`

	// the address of this ip
	// Required: true
	IP *string `json:"ip"`

	// the duration that this ip is allocated
	// Required: true
	Lifetime *int64 `json:"lifetime"`

	// the project id of this ip
	// Required: true
	Projectid *string `json:"projectid"`

	// the project name of this ip
	// Required: true
	Projectname *string `json:"projectname"`

	// the start time of this ip
	// Required: true
	// Format: date-time
	Start *strfmt.DateTime `json:"start"`

	// the tenant of this ip
	// Required: true
	Tenant *string `json:"tenant"`

	// warnings that occurred when calculating the usage of this ip
	// Required: true
	Warnings []string `json:"warnings"`
}

// Validate validates this v1 IP usage
func (m *V1IPUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IPUsage) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IPUsage) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *V1IPUsage) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	return nil
}

func (m *V1IPUsage) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1IPUsage) validateProjectname(formats strfmt.Registry) error {

	if err := validate.Required("projectname", "body", m.Projectname); err != nil {
		return err
	}

	return nil
}

func (m *V1IPUsage) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IPUsage) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *V1IPUsage) validateWarnings(formats strfmt.Registry) error {

	if err := validate.Required("warnings", "body", m.Warnings); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 IP usage based on context it is used
func (m *V1IPUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IPUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPUsage) UnmarshalBinary(b []byte) error {
	var res V1IPUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
