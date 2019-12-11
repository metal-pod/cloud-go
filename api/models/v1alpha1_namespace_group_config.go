// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V1alpha1NamespaceGroupConfig v1alpha1 namespace group config
// swagger:model v1alpha1.NamespaceGroupConfig
type V1alpha1NamespaceGroupConfig struct {

	// cluster groupname template
	ClusterGroupnameTemplate string `json:"clusterGroupnameTemplate,omitempty"`

	// excluded namespaces
	ExcludedNamespaces string `json:"excludedNamespaces,omitempty"`

	// expected groups list
	ExpectedGroupsList string `json:"expectedGroupsList,omitempty"`

	// namespace max length
	NamespaceMaxLength int32 `json:"namespaceMaxLength,omitempty"`

	// role binding name template
	RoleBindingNameTemplate string `json:"roleBindingNameTemplate,omitempty"`
}

// Validate validates this v1alpha1 namespace group config
func (m *V1alpha1NamespaceGroupConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1NamespaceGroupConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1NamespaceGroupConfig) UnmarshalBinary(b []byte) error {
	var res V1alpha1NamespaceGroupConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
