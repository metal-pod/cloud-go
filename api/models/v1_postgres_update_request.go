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

// V1PostgresUpdateRequest v1 postgres update request
//
// swagger:model v1.PostgresUpdateRequest
type V1PostgresUpdateRequest struct {

	// access list
	AccessList *V1AccessList `json:"accessList,omitempty"`

	// audit logs
	AuditLogs bool `json:"auditLogs,omitempty"`

	// backup
	Backup string `json:"backup,omitempty"`

	// connection
	Connection *V1Connection `json:"connection,omitempty"`

	// dedicatedloadbalancerip
	// Required: true
	Dedicatedloadbalancerip *string `json:"dedicatedloadbalancerip"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// labels
	Labels V1PostgresUpdateRequestLabels `json:"labels,omitempty"`

	// maintenance
	Maintenance []string `json:"maintenance"`

	// number of instances
	NumberOfInstances int32 `json:"numberOfInstances,omitempty"`

	// partition ID
	PartitionID string `json:"partitionID,omitempty"`

	// postgres params
	PostgresParams map[string]string `json:"postgresParams,omitempty"`

	// project ID
	ProjectID string `json:"projectID,omitempty"`

	// size
	Size *V1PostgresSize `json:"size,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 postgres update request
func (m *V1PostgresUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDedicatedloadbalancerip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
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

func (m *V1PostgresUpdateRequest) validateAccessList(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessList) { // not required
		return nil
	}

	if m.AccessList != nil {
		if err := m.AccessList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessList")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresUpdateRequest) validateConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresUpdateRequest) validateDedicatedloadbalancerip(formats strfmt.Registry) error {

	if err := validate.Required("dedicatedloadbalancerip", "body", m.Dedicatedloadbalancerip); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresUpdateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresUpdateRequest) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if m.Labels != nil {
		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresUpdateRequest) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 postgres update request based on the context it is used
func (m *V1PostgresUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
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

func (m *V1PostgresUpdateRequest) contextValidateAccessList(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessList != nil {

		if swag.IsZero(m.AccessList) { // not required
			return nil
		}

		if err := m.AccessList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessList")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresUpdateRequest) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {

		if swag.IsZero(m.Connection) { // not required
			return nil
		}

		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresUpdateRequest) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labels")
		}
		return err
	}

	return nil
}

func (m *V1PostgresUpdateRequest) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {

		if swag.IsZero(m.Size) { // not required
			return nil
		}

		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresUpdateRequest) UnmarshalBinary(b []byte) error {
	var res V1PostgresUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
