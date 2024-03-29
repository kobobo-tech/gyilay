package schematic

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OperationalConfigurationVersion is the current supported version.
const OperationalConfigurationVersion = "core.hydra.io/v1alpha1"

// OperationalConfiguration is the main resource type for a Hydra application.
type OperationalConfiguration struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`
	Spec            OperationalConfigurationSpec   `json:"spec,omitempty"`
	Status          OperationalConfigurationStatus `json:"status,omitempty"`
}

// OperationalConfigurationSpec is the main part of the a Hydra application configuration.
type OperationalConfigurationSpec struct {
	ParameterValues []ParameterValue         `json:"parameterValues,omitempty"`
	Scopes          []Scope                  `json:"scopes,omitempty"` // monkey trials
	Components      []ComponentConfiguration `json:"components,omitempty"`
}

// OperationalConfigurationStatus is the status of an ops config.
type OperationalConfigurationStatus struct {
	Name string `json:"name,omitempty"`
}

// OperationalConfigurationList describes a list of operational configurations.
type OperationalConfigurationList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`
	Items         []OperationalConfiguration `json:"items"`
}

// ParameterValue describes the value to be assigned to a given parameter.
type ParameterValue struct{}

// Scope defines the association to a named scope
type Scope struct{}

// ComponentConfiguration defines a configuration that attaches to a particular component.
type ComponentConfiguration struct{}
