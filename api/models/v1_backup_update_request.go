// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1BackupUpdateRequest v1 backup update request
//
// swagger:model v1.BackupUpdateRequest
type V1BackupUpdateRequest struct {

	// id
	ID string `json:"id,omitempty"`

	// retention
	Retention int32 `json:"retention,omitempty"`

	// schedule
	Schedule string `json:"schedule,omitempty"`
}

// Validate validates this v1 backup update request
func (m *V1BackupUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 backup update request based on context it is used
func (m *V1BackupUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1BackupUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BackupUpdateRequest) UnmarshalBinary(b []byte) error {
	var res V1BackupUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
