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

// V1PostgresCreateRequest v1 postgres create request
//
// swagger:model v1.PostgresCreateRequest
type V1PostgresCreateRequest struct {

	// ID
	// Required: true
	ID *string `json:"ID"`

	// access list
	AccessList *V1AccessList `json:"accessList,omitempty"`

	// backup
	Backup *V1Backup `json:"backup,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// maintenance
	Maintenance *V1MaintenanceWindow `json:"maintenance,omitempty"`

	// number of instances
	NumberOfInstances int32 `json:"numberOfInstances,omitempty"`

	// partition ID
	PartitionID string `json:"partitionID,omitempty"`

	// project ID
	ProjectID string `json:"projectID,omitempty"`

	// size
	Size *V1Size `json:"size,omitempty"`

	// tenant
	Tenant string `json:"tenant,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 postgres create request
func (m *V1PostgresCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresCreateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresCreateRequest) validateAccessList(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessList) { // not required
		return nil
	}

	if m.AccessList != nil {
		if err := m.AccessList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessList")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresCreateRequest) validateBackup(formats strfmt.Registry) error {
	if swag.IsZero(m.Backup) { // not required
		return nil
	}

	if m.Backup != nil {
		if err := m.Backup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresCreateRequest) validateMaintenance(formats strfmt.Registry) error {
	if swag.IsZero(m.Maintenance) { // not required
		return nil
	}

	if m.Maintenance != nil {
		if err := m.Maintenance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresCreateRequest) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 postgres create request based on the context it is used
func (m *V1PostgresCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaintenance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresCreateRequest) contextValidateAccessList(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessList != nil {
		if err := m.AccessList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessList")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresCreateRequest) contextValidateBackup(ctx context.Context, formats strfmt.Registry) error {

	if m.Backup != nil {
		if err := m.Backup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresCreateRequest) contextValidateMaintenance(ctx context.Context, formats strfmt.Registry) error {

	if m.Maintenance != nil {
		if err := m.Maintenance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresCreateRequest) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {
		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresCreateRequest) UnmarshalBinary(b []byte) error {
	var res V1PostgresCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
