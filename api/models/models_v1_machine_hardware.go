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

// ModelsV1MachineHardware models v1 machine hardware
//
// swagger:model models.V1MachineHardware
type ModelsV1MachineHardware struct {

	// cpu cores
	// Required: true
	CPUCores *int32 `json:"cpu_cores"`

	// cpus
	// Required: true
	Cpus []*ModelsV1MetalCPU `json:"cpus"`

	// disks
	// Required: true
	Disks []*ModelsV1MachineBlockDevice `json:"disks"`

	// gpus
	// Required: true
	Gpus []*ModelsV1MetalGPU `json:"gpus"`

	// memory
	// Required: true
	Memory *int64 `json:"memory"`

	// nics
	// Required: true
	Nics []*ModelsV1MachineNic `json:"nics"`
}

// Validate validates this models v1 machine hardware
func (m *ModelsV1MachineHardware) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1MachineHardware) validateCPUCores(formats strfmt.Registry) error {

	if err := validate.Required("cpu_cores", "body", m.CPUCores); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineHardware) validateCpus(formats strfmt.Registry) error {

	if err := validate.Required("cpus", "body", m.Cpus); err != nil {
		return err
	}

	for i := 0; i < len(m.Cpus); i++ {
		if swag.IsZero(m.Cpus[i]) { // not required
			continue
		}

		if m.Cpus[i] != nil {
			if err := m.Cpus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineHardware) validateDisks(formats strfmt.Registry) error {

	if err := validate.Required("disks", "body", m.Disks); err != nil {
		return err
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineHardware) validateGpus(formats strfmt.Registry) error {

	if err := validate.Required("gpus", "body", m.Gpus); err != nil {
		return err
	}

	for i := 0; i < len(m.Gpus); i++ {
		if swag.IsZero(m.Gpus[i]) { // not required
			continue
		}

		if m.Gpus[i] != nil {
			if err := m.Gpus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineHardware) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineHardware) validateNics(formats strfmt.Registry) error {

	if err := validate.Required("nics", "body", m.Nics); err != nil {
		return err
	}

	for i := 0; i < len(m.Nics); i++ {
		if swag.IsZero(m.Nics[i]) { // not required
			continue
		}

		if m.Nics[i] != nil {
			if err := m.Nics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this models v1 machine hardware based on the context it is used
func (m *ModelsV1MachineHardware) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCpus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1MachineHardware) contextValidateCpus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cpus); i++ {

		if m.Cpus[i] != nil {

			if swag.IsZero(m.Cpus[i]) { // not required
				return nil
			}

			if err := m.Cpus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineHardware) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {

			if swag.IsZero(m.Disks[i]) { // not required
				return nil
			}

			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineHardware) contextValidateGpus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Gpus); i++ {

		if m.Gpus[i] != nil {

			if swag.IsZero(m.Gpus[i]) { // not required
				return nil
			}

			if err := m.Gpus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsV1MachineHardware) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nics); i++ {

		if m.Nics[i] != nil {

			if swag.IsZero(m.Nics[i]) { // not required
				return nil
			}

			if err := m.Nics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1MachineHardware) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1MachineHardware) UnmarshalBinary(b []byte) error {
	var res ModelsV1MachineHardware
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
