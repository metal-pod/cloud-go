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
)

// V1PostgresStatus v1 postgres status
//
// swagger:model v1.PostgresStatus
type V1PostgresStatus struct {

	// additionalsockets
	Additionalsockets []*V1PostgresSocket `json:"additionalsockets"`

	// child reference
	ChildReference string `json:"childReference,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// socket
	Socket *V1PostgresSocket `json:"socket,omitempty"`
}

// Validate validates this v1 postgres status
func (m *V1PostgresStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalsockets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSocket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresStatus) validateAdditionalsockets(formats strfmt.Registry) error {
	if swag.IsZero(m.Additionalsockets) { // not required
		return nil
	}

	for i := 0; i < len(m.Additionalsockets); i++ {
		if swag.IsZero(m.Additionalsockets[i]) { // not required
			continue
		}

		if m.Additionalsockets[i] != nil {
			if err := m.Additionalsockets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalsockets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalsockets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PostgresStatus) validateSocket(formats strfmt.Registry) error {
	if swag.IsZero(m.Socket) { // not required
		return nil
	}

	if m.Socket != nil {
		if err := m.Socket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("socket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("socket")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 postgres status based on the context it is used
func (m *V1PostgresStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalsockets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSocket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresStatus) contextValidateAdditionalsockets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Additionalsockets); i++ {

		if m.Additionalsockets[i] != nil {

			if swag.IsZero(m.Additionalsockets[i]) { // not required
				return nil
			}

			if err := m.Additionalsockets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalsockets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalsockets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PostgresStatus) contextValidateSocket(ctx context.Context, formats strfmt.Registry) error {

	if m.Socket != nil {

		if swag.IsZero(m.Socket) { // not required
			return nil
		}

		if err := m.Socket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("socket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("socket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresStatus) UnmarshalBinary(b []byte) error {
	var res V1PostgresStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
