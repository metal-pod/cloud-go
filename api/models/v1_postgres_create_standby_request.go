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

// V1PostgresCreateStandbyRequest v1 postgres create standby request
//
// swagger:model v1.PostgresCreateStandbyRequest
type V1PostgresCreateStandbyRequest struct {

	// backup
	Backup string `json:"backup,omitempty" yaml:"backup,omitempty"`

	// description
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// maintenance
	Maintenance []string `json:"maintenance" yaml:"maintenance"`

	// number of instances
	NumberOfInstances int32 `json:"numberOfInstances,omitempty" yaml:"numberOfInstances,omitempty"`

	// partition ID
	PartitionID string `json:"partitionID,omitempty" yaml:"partitionID,omitempty"`

	// primary Id
	// Required: true
	PrimaryID *string `json:"primaryId" yaml:"primaryId"`
}

// Validate validates this v1 postgres create standby request
func (m *V1PostgresCreateStandbyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrimaryID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresCreateStandbyRequest) validatePrimaryID(formats strfmt.Registry) error {

	if err := validate.Required("primaryId", "body", m.PrimaryID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 postgres create standby request based on context it is used
func (m *V1PostgresCreateStandbyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresCreateStandbyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresCreateStandbyRequest) UnmarshalBinary(b []byte) error {
	var res V1PostgresCreateStandbyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
