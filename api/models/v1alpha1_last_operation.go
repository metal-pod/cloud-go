// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1alpha1LastOperation v1alpha1 last operation
// swagger:model v1alpha1.LastOperation
type V1alpha1LastOperation struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// last update time
	// Required: true
	LastUpdateTime *string `json:"lastUpdateTime"`

	// progress
	// Required: true
	Progress *int32 `json:"progress"`

	// state
	// Required: true
	State *string `json:"state"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this v1alpha1 last operation
func (m *V1alpha1LastOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *V1alpha1LastOperation) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1LastOperation) validateLastUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdateTime", "body", m.LastUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1LastOperation) validateProgress(formats strfmt.Registry) error {

	if err := validate.Required("progress", "body", m.Progress); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1LastOperation) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1LastOperation) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1LastOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1LastOperation) UnmarshalBinary(b []byte) error {
	var res V1alpha1LastOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
