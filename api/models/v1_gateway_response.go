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

// V1GatewayResponse v1 gateway response
//
// swagger:model v1.GatewayResponse
type V1GatewayResponse struct {

	// project UID
	// Required: true
	ProjectUID *string `json:"ProjectUID"`

	// name
	// Required: true
	Name *string `json:"name"`

	// peers
	// Required: true
	Peers []*V1PeerSpec `json:"peers"`

	// pipes
	// Required: true
	Pipes []*V1PipeSpec `json:"pipes"`

	// port
	Port int64 `json:"port,omitempty"`

	// private key
	// Required: true
	PrivateKey *string `json:"private_key"`

	// public key
	// Required: true
	PublicKey *string `json:"public_key"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this v1 gateway response
func (m *V1GatewayResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
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

func (m *V1GatewayResponse) validateProjectUID(formats strfmt.Registry) error {

	if err := validate.Required("ProjectUID", "body", m.ProjectUID); err != nil {
		return err
	}

	return nil
}

func (m *V1GatewayResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1GatewayResponse) validatePeers(formats strfmt.Registry) error {

	if err := validate.Required("peers", "body", m.Peers); err != nil {
		return err
	}

	for i := 0; i < len(m.Peers); i++ {
		if swag.IsZero(m.Peers[i]) { // not required
			continue
		}

		if m.Peers[i] != nil {
			if err := m.Peers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1GatewayResponse) validatePipes(formats strfmt.Registry) error {

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

func (m *V1GatewayResponse) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("private_key", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *V1GatewayResponse) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("public_key", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

func (m *V1GatewayResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 gateway response based on the context it is used
func (m *V1GatewayResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePipes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GatewayResponse) contextValidatePeers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Peers); i++ {

		if m.Peers[i] != nil {
			if err := m.Peers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1GatewayResponse) contextValidatePipes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1GatewayResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GatewayResponse) UnmarshalBinary(b []byte) error {
	var res V1GatewayResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
