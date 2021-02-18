// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Backup v1 backup
//
// swagger:model v1.Backup
type V1Backup struct {

	// project ID
	ProjectID string `json:"projectID,omitempty"`

	// retention
	Retention int32 `json:"retention,omitempty"`

	// s3 bucket URL
	S3BucketURL string `json:"s3BucketURL,omitempty"`

	// schedule
	Schedule string `json:"schedule,omitempty"`

	// secret
	Secret *V1BackupSecret `json:"secret,omitempty"`

	// tenant
	Tenant string `json:"tenant,omitempty"`
}

// Validate validates this v1 backup
func (m *V1Backup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Backup) validateSecret(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 backup based on the context it is used
func (m *V1Backup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Backup) contextValidateSecret(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1Backup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Backup) UnmarshalBinary(b []byte) error {
	var res V1Backup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
