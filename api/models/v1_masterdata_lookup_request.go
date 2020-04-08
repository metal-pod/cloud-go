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

// V1MasterdataLookupRequest v1 masterdata lookup request
// swagger:model v1.MasterdataLookupRequest
type V1MasterdataLookupRequest struct {

	// cluster ID
	// Required: true
	ClusterID *string `json:"ClusterID"`

	// cluster name project
	// Required: true
	ClusterNameProject *V1ClusterNameProject `json:"ClusterNameProject"`
}

// Validate validates this v1 masterdata lookup request
func (m *V1MasterdataLookupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNameProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MasterdataLookupRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("ClusterID", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *V1MasterdataLookupRequest) validateClusterNameProject(formats strfmt.Registry) error {

	if err := validate.Required("ClusterNameProject", "body", m.ClusterNameProject); err != nil {
		return err
	}

	if m.ClusterNameProject != nil {
		if err := m.ClusterNameProject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ClusterNameProject")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MasterdataLookupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MasterdataLookupRequest) UnmarshalBinary(b []byte) error {
	var res V1MasterdataLookupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
