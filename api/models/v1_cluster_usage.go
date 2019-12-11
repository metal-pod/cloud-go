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

// V1ClusterUsage v1 cluster usage
// swagger:model v1.ClusterUsage
type V1ClusterUsage struct {

	// the end time of this cluster
	// Required: true
	// Format: date-time
	Clusterend *strfmt.DateTime `json:"clusterend"`

	// the cluster id of this cluster
	// Required: true
	Clusterid *string `json:"clusterid"`

	// the cluster name of this cluster
	// Required: true
	Clustername *string `json:"clustername"`

	// the start time of this cluster
	// Required: true
	// Format: date-time
	Clusterstart *strfmt.DateTime `json:"clusterstart"`

	// the duration that this cluster is running
	// Required: true
	Lifetime *int64 `json:"lifetime"`

	// the partition of this cluster
	// Required: true
	Partition *string `json:"partition"`

	// the project id of this cluster
	// Required: true
	Projectid *string `json:"projectid"`

	// the project name of this cluster
	// Required: true
	Projectname *string `json:"projectname"`

	// the tenant of this cluster
	// Required: true
	Tenant *string `json:"tenant"`

	// warnings that occurred when calculating the usage of this cluster
	// Required: true
	Warnings []string `json:"warnings"`
}

// Validate validates this v1 cluster usage
func (m *V1ClusterUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterstart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterUsage) validateClusterend(formats strfmt.Registry) error {

	if err := validate.Required("clusterend", "body", m.Clusterend); err != nil {
		return err
	}

	if err := validate.FormatOf("clusterend", "body", "date-time", m.Clusterend.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateClusterid(formats strfmt.Registry) error {

	if err := validate.Required("clusterid", "body", m.Clusterid); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateClustername(formats strfmt.Registry) error {

	if err := validate.Required("clustername", "body", m.Clustername); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateClusterstart(formats strfmt.Registry) error {

	if err := validate.Required("clusterstart", "body", m.Clusterstart); err != nil {
		return err
	}

	if err := validate.FormatOf("clusterstart", "body", "date-time", m.Clusterstart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validatePartition(formats strfmt.Registry) error {

	if err := validate.Required("partition", "body", m.Partition); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateProjectname(formats strfmt.Registry) error {

	if err := validate.Required("projectname", "body", m.Projectname); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateWarnings(formats strfmt.Registry) error {

	if err := validate.Required("warnings", "body", m.Warnings); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterUsage) UnmarshalBinary(b []byte) error {
	var res V1ClusterUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
