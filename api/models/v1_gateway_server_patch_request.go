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

// V1GatewayServerPatchRequest v1 gateway server patch request
//
// swagger:model v1.GatewayServerPatchRequest
type V1GatewayServerPatchRequest struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// pipes
	// Required: true
	Pipes []*V1PipeSpec `json:"pipes"`

	// project UID
	// Required: true
	ProjectUID *string `json:"projectUID"`

	// uid
	// Required: true
	UID *string `json:"uid"`
}

// Validate validates this v1 gateway server patch request
func (m *V1GatewayServerPatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GatewayServerPatchRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1GatewayServerPatchRequest) validatePipes(formats strfmt.Registry) error {

	if err := validate.Required("pipes", "body", m.Pipes); err != nil {
		return err
	}

	for i := 0; i < len(m.Pipes); i++ {
		if swag.IsZero(m.Pipes[i]) { // not required
			continue
		}

		if m.Pipes[i] != nil {
			if err := m.Pipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1GatewayServerPatchRequest) validateProjectUID(formats strfmt.Registry) error {

	if err := validate.Required("projectUID", "body", m.ProjectUID); err != nil {
		return err
	}

	return nil
}

func (m *V1GatewayServerPatchRequest) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 gateway server patch request based on the context it is used
func (m *V1GatewayServerPatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePipes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GatewayServerPatchRequest) contextValidatePipes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Pipes); i++ {

		if m.Pipes[i] != nil {
			if err := m.Pipes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GatewayServerPatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GatewayServerPatchRequest) UnmarshalBinary(b []byte) error {
	var res V1GatewayServerPatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
