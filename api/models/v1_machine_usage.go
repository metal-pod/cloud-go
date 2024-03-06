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

// V1MachineUsage v1 machine usage
//
// swagger:model v1.MachineUsage
type V1MachineUsage struct {

	// the allocation id of this machine
	// Required: true
	Allocationid *string `json:"allocationid"`

	// the cluster id of this machine
	// Required: true
	Clusterid *string `json:"clusterid"`

	// the cluster name of this machine
	// Required: true
	Clustername *string `json:"clustername"`

	// the contract number attached to this entity
	// Required: true
	Contract *string `json:"contract"`

	// the debtor id attached to this entity
	// Required: true
	Debtorid *string `json:"debtorid"`

	// the image id of this machine
	// Required: true
	Imageid *string `json:"imageid"`

	// the duration that this machine is running
	// Required: true
	Lifetime *int64 `json:"lifetime"`

	// the end time of this machine
	// Required: true
	// Format: date-time
	Machineend *strfmt.DateTime `json:"machineend"`

	// the id of this machine
	// Required: true
	Machineid *string `json:"machineid"`

	// the name of this machine
	// Required: true
	Machinename *string `json:"machinename"`

	// the start time of this machine
	// Required: true
	// Format: date-time
	Machinestart *strfmt.DateTime `json:"machinestart"`

	// the partition of this machine
	// Required: true
	Partition *string `json:"partition"`

	// the project id of this entity
	// Required: true
	Projectid *string `json:"projectid"`

	// the project name of this entity
	// Required: true
	Projectname *string `json:"projectname"`

	// the size id of this machine
	// Required: true
	Sizeid *string `json:"sizeid"`

	// the tenant of this entity
	// Required: true
	Tenant *string `json:"tenant"`

	// the tenant name of this entity
	// Required: true
	Tenantname *string `json:"tenantname"`
}

// Validate validates this v1 machine usage
func (m *V1MachineUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocationid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinestart(formats); err != nil {
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

	if err := m.validateSizeid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineUsage) validateAllocationid(formats strfmt.Registry) error {

	if err := validate.Required("allocationid", "body", m.Allocationid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateClusterid(formats strfmt.Registry) error {

	if err := validate.Required("clusterid", "body", m.Clusterid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateClustername(formats strfmt.Registry) error {

	if err := validate.Required("clustername", "body", m.Clustername); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateContract(formats strfmt.Registry) error {

	if err := validate.Required("contract", "body", m.Contract); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateDebtorid(formats strfmt.Registry) error {

	if err := validate.Required("debtorid", "body", m.Debtorid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateImageid(formats strfmt.Registry) error {

	if err := validate.Required("imageid", "body", m.Imageid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateMachineend(formats strfmt.Registry) error {

	if err := validate.Required("machineend", "body", m.Machineend); err != nil {
		return err
	}

	if err := validate.FormatOf("machineend", "body", "date-time", m.Machineend.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateMachineid(formats strfmt.Registry) error {

	if err := validate.Required("machineid", "body", m.Machineid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateMachinename(formats strfmt.Registry) error {

	if err := validate.Required("machinename", "body", m.Machinename); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateMachinestart(formats strfmt.Registry) error {

	if err := validate.Required("machinestart", "body", m.Machinestart); err != nil {
		return err
	}

	if err := validate.FormatOf("machinestart", "body", "date-time", m.Machinestart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validatePartition(formats strfmt.Registry) error {

	if err := validate.Required("partition", "body", m.Partition); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateProjectname(formats strfmt.Registry) error {

	if err := validate.Required("projectname", "body", m.Projectname); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateSizeid(formats strfmt.Registry) error {

	if err := validate.Required("sizeid", "body", m.Sizeid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUsage) validateTenantname(formats strfmt.Registry) error {

	if err := validate.Required("tenantname", "body", m.Tenantname); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine usage based on context it is used
func (m *V1MachineUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineUsage) UnmarshalBinary(b []byte) error {
	var res V1MachineUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
