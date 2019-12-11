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

// V1Taint The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
// swagger:model v1.Taint
type V1Taint struct {

	// Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	// Required: true
	Effect *string `json:"effect"`

	// Required. The taint key to be applied to a node.
	// Required: true
	Key *string `json:"key"`

	// TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.
	TimeAdded string `json:"timeAdded,omitempty"`

	// Required. The taint value corresponding to the taint key.
	Value string `json:"value,omitempty"`
}

// Validate validates this v1 taint
func (m *V1Taint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Taint) validateEffect(formats strfmt.Registry) error {

	if err := validate.Required("effect", "body", m.Effect); err != nil {
		return err
	}

	return nil
}

func (m *V1Taint) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Taint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Taint) UnmarshalBinary(b []byte) error {
	var res V1Taint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
