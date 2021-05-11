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

// V1PostgresBackupConfigCreateRequest v1 postgres backup config create request
//
// swagger:model v1.PostgresBackupConfigCreateRequest
type V1PostgresBackupConfigCreateRequest struct {

	// autocreate
	Autocreate bool `json:"autocreate,omitempty"`

	// created by
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// name
	Name string `json:"name,omitempty"`

	// partition
	Partition string `json:"partition,omitempty"`

	// project ID
	ProjectID string `json:"projectID,omitempty"`

	// retention
	Retention int32 `json:"retention,omitempty"`

	// s3 bucket name
	S3BucketName string `json:"s3BucketName,omitempty"`

	// s3 endpoint
	S3Endpoint string `json:"s3Endpoint,omitempty"`

	// s3 region
	S3Region string `json:"s3Region,omitempty"`

	// schedule
	Schedule string `json:"schedule,omitempty"`

	// secret
	Secret *V1PostgresBackupSecret `json:"secret,omitempty"`
}

// Validate validates this v1 postgres backup config create request
func (m *V1PostgresBackupConfigCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresBackupConfigCreateRequest) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("createdBy", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresBackupConfigCreateRequest) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if m.Secret != nil {
		if err := m.Secret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 postgres backup config create request based on the context it is used
func (m *V1PostgresBackupConfigCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresBackupConfigCreateRequest) contextValidateSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.Secret != nil {
		if err := m.Secret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresBackupConfigCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresBackupConfigCreateRequest) UnmarshalBinary(b []byte) error {
	var res V1PostgresBackupConfigCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
