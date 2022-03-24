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

// V1UsageCommon v1 usage common
//
// swagger:model v1.UsageCommon
type V1UsageCommon struct {

	// the contract number attached to this entity
	// Required: true
	Contract *string `json:"contract"`

	// the debtor id attached to this entity
	// Required: true
	Debtorid *string `json:"debtorid"`

	// the project id of this entity
	// Required: true
	Projectid *string `json:"projectid"`

	// the project name of this entity
	// Required: true
	Projectname *string `json:"projectname"`

	// the tenant of this entity
	// Required: true
	Tenant *string `json:"tenant"`

	// the tenant name of this entity
	// Required: true
	Tenantname *string `json:"tenantname"`
}

// Validate validates this v1 usage common
func (m *V1UsageCommon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UsageCommon) validateContract(formats strfmt.Registry) error {

	if err := validate.Required("contract", "body", m.Contract); err != nil {
		return err
	}

	return nil
}

func (m *V1UsageCommon) validateDebtorid(formats strfmt.Registry) error {

	if err := validate.Required("debtorid", "body", m.Debtorid); err != nil {
		return err
	}

	return nil
}

func (m *V1UsageCommon) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1UsageCommon) validateProjectname(formats strfmt.Registry) error {

	if err := validate.Required("projectname", "body", m.Projectname); err != nil {
		return err
	}

	return nil
}

func (m *V1UsageCommon) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *V1UsageCommon) validateTenantname(formats strfmt.Registry) error {

	if err := validate.Required("tenantname", "body", m.Tenantname); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 usage common based on context it is used
func (m *V1UsageCommon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UsageCommon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UsageCommon) UnmarshalBinary(b []byte) error {
	var res V1UsageCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
