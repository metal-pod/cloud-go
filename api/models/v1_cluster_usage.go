// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ClusterUsage v1 cluster usage
//
// swagger:model v1.ClusterUsage
type V1ClusterUsage struct {

	// the average amount of worker groups during the time window
	// Required: true
	Averageworkergroups *string `json:"averageworkergroups" yaml:"averageworkergroups"`

	// the end time of this cluster
	// Required: true
	// Format: date-time
	Clusterend *strfmt.DateTime `json:"clusterend" yaml:"clusterend"`

	// the cluster id of this cluster
	// Required: true
	Clusterid *string `json:"clusterid" yaml:"clusterid"`

	// the cluster name of this cluster
	// Required: true
	Clustername *string `json:"clustername" yaml:"clustername"`

	// the start time of this cluster
	// Required: true
	// Format: date-time
	Clusterstart *strfmt.DateTime `json:"clusterstart" yaml:"clusterstart"`

	// the contract number attached to this entity
	// Required: true
	Contract *string `json:"contract" yaml:"contract"`

	// the debtor id attached to this entity
	// Required: true
	Debtorid *string `json:"debtorid" yaml:"debtorid"`

	// the duration that this cluster is running
	// Required: true
	Lifetime *int64 `json:"lifetime" yaml:"lifetime"`

	// the partition of this cluster
	// Required: true
	Partition *string `json:"partition" yaml:"partition"`

	// the project id of this entity
	// Required: true
	Projectid *string `json:"projectid" yaml:"projectid"`

	// the project name of this entity
	// Required: true
	Projectname *string `json:"projectname" yaml:"projectname"`

	// the tenant of this entity
	// Required: true
	Tenant *string `json:"tenant" yaml:"tenant"`

	// the tenant name of this entity
	// Required: true
	Tenantname *string `json:"tenantname" yaml:"tenantname"`

	// the worker groups of this cluster
	// Required: true
	Workergroups []*V1ClusterWorker `json:"workergroups" yaml:"workergroups"`
}

// Validate validates this v1 cluster usage
func (m *V1ClusterUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAverageworkergroups(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorid(formats); err != nil {
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

	if err := m.validateTenantname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkergroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterUsage) validateAverageworkergroups(formats strfmt.Registry) error {

	if err := validate.Required("averageworkergroups", "body", m.Averageworkergroups); err != nil {
		return err
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

func (m *V1ClusterUsage) validateContract(formats strfmt.Registry) error {

	if err := validate.Required("contract", "body", m.Contract); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateDebtorid(formats strfmt.Registry) error {

	if err := validate.Required("debtorid", "body", m.Debtorid); err != nil {
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

func (m *V1ClusterUsage) validateTenantname(formats strfmt.Registry) error {

	if err := validate.Required("tenantname", "body", m.Tenantname); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterUsage) validateWorkergroups(formats strfmt.Registry) error {

	if err := validate.Required("workergroups", "body", m.Workergroups); err != nil {
		return err
	}

	for i := 0; i < len(m.Workergroups); i++ {
		if swag.IsZero(m.Workergroups[i]) { // not required
			continue
		}

		if m.Workergroups[i] != nil {
			if err := m.Workergroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workergroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workergroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 cluster usage based on the context it is used
func (m *V1ClusterUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkergroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterUsage) contextValidateWorkergroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workergroups); i++ {

		if m.Workergroups[i] != nil {
			if err := m.Workergroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workergroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workergroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

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
