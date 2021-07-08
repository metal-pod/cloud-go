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

// V1IPFindRequest v1 IP find request
//
// swagger:model v1.IPFindRequest
type V1IPFindRequest struct {

	// IP address
	// Required: true
	IPAddress *string `json:"IPAddress"`

	// machine ID
	// Required: true
	MachineID *string `json:"MachineID"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// network ID
	// Required: true
	NetworkID *string `json:"NetworkID"`

	// parent prefix cidr
	// Required: true
	ParentPrefixCidr *string `json:"ParentPrefixCidr"`

	// project ID
	// Required: true
	ProjectID *string `json:"ProjectID"`

	// tags
	// Required: true
	Tags []string `json:"Tags"`

	// type
	// Required: true
	Type *string `json:"Type"`
}

// Validate validates this v1 IP find request
func (m *V1IPFindRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentPrefixCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
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

func (m *V1IPFindRequest) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("IPAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *V1IPFindRequest) validateMachineID(formats strfmt.Registry) error {

	if err := validate.Required("MachineID", "body", m.MachineID); err != nil {
		return err
	}

	return nil
}

func (m *V1IPFindRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1IPFindRequest) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("NetworkID", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

func (m *V1IPFindRequest) validateParentPrefixCidr(formats strfmt.Registry) error {

	if err := validate.Required("ParentPrefixCidr", "body", m.ParentPrefixCidr); err != nil {
		return err
	}

	return nil
}

func (m *V1IPFindRequest) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("ProjectID", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *V1IPFindRequest) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("Tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *V1IPFindRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 IP find request based on context it is used
func (m *V1IPFindRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IPFindRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPFindRequest) UnmarshalBinary(b []byte) error {
	var res V1IPFindRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
