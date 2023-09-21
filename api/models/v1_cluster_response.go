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

// V1ClusterResponse v1 cluster response
//
// swagger:model v1.ClusterResponse
type V1ClusterResponse struct {

	// additional networks
	// Required: true
	AdditionalNetworks []string `json:"AdditionalNetworks"`

	// cluster features
	// Required: true
	ClusterFeatures *V1ClusterFeatures `json:"ClusterFeatures"`

	// control plane feature gates
	// Required: true
	ControlPlaneFeatureGates []string `json:"ControlPlaneFeatureGates"`

	// creation timestamp
	// Required: true
	// Format: date-time
	CreationTimestamp *strfmt.DateTime `json:"CreationTimestamp"`

	// creator
	// Required: true
	Creator *string `json:"Creator"`

	// custom default storage class
	// Required: true
	CustomDefaultStorageClass *V1CustomDefaultStorageClass `json:"CustomDefaultStorageClass"`

	// DNS endpoint
	// Required: true
	DNSEndpoint *string `json:"DNSEndpoint"`

	// description
	// Required: true
	Description *string `json:"Description"`

	// egress rules
	// Required: true
	EgressRules []*V1EgressRule `json:"EgressRules"`

	// firewall controller version
	// Required: true
	FirewallControllerVersion *string `json:"FirewallControllerVersion"`

	// firewall image
	// Required: true
	FirewallImage *string `json:"FirewallImage"`

	// firewall size
	// Required: true
	FirewallSize *string `json:"FirewallSize"`

	// ID
	// Required: true
	ID *string `json:"ID"`

	// kubernetes
	// Required: true
	Kubernetes *V1Kubernetes `json:"Kubernetes"`

	// labels
	// Required: true
	Labels map[string]string `json:"Labels"`

	// maintenance
	// Required: true
	Maintenance *V1Maintenance `json:"Maintenance"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// networking
	// Required: true
	Networking *V1Networking `json:"Networking"`

	// partition ID
	// Required: true
	PartitionID *string `json:"PartitionID"`

	// project ID
	// Required: true
	ProjectID *string `json:"ProjectID"`

	// purpose
	// Required: true
	Purpose *string `json:"Purpose"`

	// status
	// Required: true
	Status *V1beta1ShootStatus `json:"Status"`

	// system components
	// Required: true
	SystemComponents *V1SystemComponents `json:"SystemComponents"`

	// tenant
	// Required: true
	Tenant *string `json:"Tenant"`

	// workers
	// Required: true
	Workers []*V1Worker `json:"Workers"`

	// cni
	// Required: true
	Cni *string `json:"cni"`

	// the firewalls which belong to this cluster
	Firewalls []*ModelsV1MachineResponse `json:"firewalls"`

	// the machines which belong to this cluster
	Machines []*ModelsV1MachineResponse `json:"machines"`
}

// Validate validates this v1 cluster response
func (m *V1ClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControlPlaneFeatureGates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomDefaultStorageClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEgressRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallControllerVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworking(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurpose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCni(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewalls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachines(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterResponse) validateAdditionalNetworks(formats strfmt.Registry) error {

	if err := validate.Required("AdditionalNetworks", "body", m.AdditionalNetworks); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateClusterFeatures(formats strfmt.Registry) error {

	if err := validate.Required("ClusterFeatures", "body", m.ClusterFeatures); err != nil {
		return err
	}

	if m.ClusterFeatures != nil {
		if err := m.ClusterFeatures.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ClusterFeatures")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ClusterFeatures")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) validateControlPlaneFeatureGates(formats strfmt.Registry) error {

	if err := validate.Required("ControlPlaneFeatureGates", "body", m.ControlPlaneFeatureGates); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateCreationTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("CreationTimestamp", "body", m.CreationTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("CreationTimestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateCreator(formats strfmt.Registry) error {

	if err := validate.Required("Creator", "body", m.Creator); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateCustomDefaultStorageClass(formats strfmt.Registry) error {

	if err := validate.Required("CustomDefaultStorageClass", "body", m.CustomDefaultStorageClass); err != nil {
		return err
	}

	if m.CustomDefaultStorageClass != nil {
		if err := m.CustomDefaultStorageClass.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CustomDefaultStorageClass")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CustomDefaultStorageClass")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) validateDNSEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("DNSEndpoint", "body", m.DNSEndpoint); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateEgressRules(formats strfmt.Registry) error {

	if err := validate.Required("EgressRules", "body", m.EgressRules); err != nil {
		return err
	}

	for i := 0; i < len(m.EgressRules); i++ {
		if swag.IsZero(m.EgressRules[i]) { // not required
			continue
		}

		if m.EgressRules[i] != nil {
			if err := m.EgressRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EgressRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("EgressRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterResponse) validateFirewallControllerVersion(formats strfmt.Registry) error {

	if err := validate.Required("FirewallControllerVersion", "body", m.FirewallControllerVersion); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateFirewallImage(formats strfmt.Registry) error {

	if err := validate.Required("FirewallImage", "body", m.FirewallImage); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateFirewallSize(formats strfmt.Registry) error {

	if err := validate.Required("FirewallSize", "body", m.FirewallSize); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateKubernetes(formats strfmt.Registry) error {

	if err := validate.Required("Kubernetes", "body", m.Kubernetes); err != nil {
		return err
	}

	if m.Kubernetes != nil {
		if err := m.Kubernetes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Kubernetes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("Labels", "body", m.Labels); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateMaintenance(formats strfmt.Registry) error {

	if err := validate.Required("Maintenance", "body", m.Maintenance); err != nil {
		return err
	}

	if m.Maintenance != nil {
		if err := m.Maintenance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Maintenance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Maintenance")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateNetworking(formats strfmt.Registry) error {

	if err := validate.Required("Networking", "body", m.Networking); err != nil {
		return err
	}

	if m.Networking != nil {
		if err := m.Networking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Networking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Networking")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) validatePartitionID(formats strfmt.Registry) error {

	if err := validate.Required("PartitionID", "body", m.PartitionID); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("ProjectID", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validatePurpose(formats strfmt.Registry) error {

	if err := validate.Required("Purpose", "body", m.Purpose); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) validateSystemComponents(formats strfmt.Registry) error {

	if err := validate.Required("SystemComponents", "body", m.SystemComponents); err != nil {
		return err
	}

	if m.SystemComponents != nil {
		if err := m.SystemComponents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SystemComponents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SystemComponents")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("Tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateWorkers(formats strfmt.Registry) error {

	if err := validate.Required("Workers", "body", m.Workers); err != nil {
		return err
	}

	for i := 0; i < len(m.Workers); i++ {
		if swag.IsZero(m.Workers[i]) { // not required
			continue
		}

		if m.Workers[i] != nil {
			if err := m.Workers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Workers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Workers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterResponse) validateCni(formats strfmt.Registry) error {

	if err := validate.Required("cni", "body", m.Cni); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterResponse) validateFirewalls(formats strfmt.Registry) error {
	if swag.IsZero(m.Firewalls) { // not required
		return nil
	}

	for i := 0; i < len(m.Firewalls); i++ {
		if swag.IsZero(m.Firewalls[i]) { // not required
			continue
		}

		if m.Firewalls[i] != nil {
			if err := m.Firewalls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("firewalls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("firewalls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterResponse) validateMachines(formats strfmt.Registry) error {
	if swag.IsZero(m.Machines) { // not required
		return nil
	}

	for i := 0; i < len(m.Machines); i++ {
		if swag.IsZero(m.Machines[i]) { // not required
			continue
		}

		if m.Machines[i] != nil {
			if err := m.Machines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 cluster response based on the context it is used
func (m *V1ClusterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomDefaultStorageClass(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEgressRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaintenance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirewalls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterResponse) contextValidateClusterFeatures(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterFeatures != nil {

		if err := m.ClusterFeatures.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ClusterFeatures")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ClusterFeatures")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) contextValidateCustomDefaultStorageClass(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomDefaultStorageClass != nil {

		if err := m.CustomDefaultStorageClass.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CustomDefaultStorageClass")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CustomDefaultStorageClass")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) contextValidateEgressRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EgressRules); i++ {

		if m.EgressRules[i] != nil {

			if swag.IsZero(m.EgressRules[i]) { // not required
				return nil
			}

			if err := m.EgressRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EgressRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("EgressRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterResponse) contextValidateKubernetes(ctx context.Context, formats strfmt.Registry) error {

	if m.Kubernetes != nil {

		if err := m.Kubernetes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Kubernetes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) contextValidateMaintenance(ctx context.Context, formats strfmt.Registry) error {

	if m.Maintenance != nil {

		if err := m.Maintenance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Maintenance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Maintenance")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) contextValidateNetworking(ctx context.Context, formats strfmt.Registry) error {

	if m.Networking != nil {

		if err := m.Networking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Networking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Networking")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) contextValidateSystemComponents(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemComponents != nil {

		if err := m.SystemComponents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SystemComponents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SystemComponents")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterResponse) contextValidateWorkers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workers); i++ {

		if m.Workers[i] != nil {

			if swag.IsZero(m.Workers[i]) { // not required
				return nil
			}

			if err := m.Workers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Workers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Workers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterResponse) contextValidateFirewalls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Firewalls); i++ {

		if m.Firewalls[i] != nil {

			if swag.IsZero(m.Firewalls[i]) { // not required
				return nil
			}

			if err := m.Firewalls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("firewalls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("firewalls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterResponse) contextValidateMachines(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Machines); i++ {

		if m.Machines[i] != nil {

			if swag.IsZero(m.Machines[i]) { // not required
				return nil
			}

			if err := m.Machines[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterResponse) UnmarshalBinary(b []byte) error {
	var res V1ClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
