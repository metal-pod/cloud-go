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

// V1IPAllocateRequest v1 IP allocate request
//
// swagger:model v1.IPAllocateRequest
type V1IPAllocateRequest struct {

	// specific IP
	// Required: true
	SpecificIP *string `json:"SpecificIP" yaml:"SpecificIP"`

	// description
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// machineid
	Machineid string `json:"machineid,omitempty" yaml:"machineid,omitempty"`

	// name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// networkid
	// Required: true
	Networkid *string `json:"networkid" yaml:"networkid"`

	// projectid
	// Required: true
	Projectid *string `json:"projectid" yaml:"projectid"`

	// tags
	// Required: true
	Tags []string `json:"tags" yaml:"tags"`

	// type
	// Required: true
	Type *string `json:"type" yaml:"type"`
}

// Validate validates this v1 IP allocate request
func (m *V1IPAllocateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpecificIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *V1IPAllocateRequest) validateSpecificIP(formats strfmt.Registry) error {

	if err := validate.Required("SpecificIP", "body", m.SpecificIP); err != nil {
		return err
	}

	return nil
}

func (m *V1IPAllocateRequest) validateNetworkid(formats strfmt.Registry) error {

	if err := validate.Required("networkid", "body", m.Networkid); err != nil {
		return err
	}

	return nil
}

func (m *V1IPAllocateRequest) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1IPAllocateRequest) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *V1IPAllocateRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 IP allocate request based on context it is used
func (m *V1IPAllocateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IPAllocateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPAllocateRequest) UnmarshalBinary(b []byte) error {
	var res V1IPAllocateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
