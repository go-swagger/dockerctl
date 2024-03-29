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
)

// EngineDescription EngineDescription provides information about an engine.
//
// swagger:model EngineDescription
type EngineDescription struct {

	// engine version
	// Example: 17.06.0
	EngineVersion string `json:"EngineVersion,omitempty"`

	// labels
	// Example: {"foo":"bar"}
	Labels map[string]string `json:"Labels,omitempty"`

	// plugins
	// Example: [{"Name":"awslogs","Type":"Log"},{"Name":"fluentd","Type":"Log"},{"Name":"gcplogs","Type":"Log"},{"Name":"gelf","Type":"Log"},{"Name":"journald","Type":"Log"},{"Name":"json-file","Type":"Log"},{"Name":"logentries","Type":"Log"},{"Name":"splunk","Type":"Log"},{"Name":"syslog","Type":"Log"},{"Name":"bridge","Type":"Network"},{"Name":"host","Type":"Network"},{"Name":"ipvlan","Type":"Network"},{"Name":"macvlan","Type":"Network"},{"Name":"null","Type":"Network"},{"Name":"overlay","Type":"Network"},{"Name":"local","Type":"Volume"},{"Name":"localhost:5000/vieux/sshfs:latest","Type":"Volume"},{"Name":"vieux/sshfs:latest","Type":"Volume"}]
	Plugins []*EngineDescriptionPluginsItems0 `json:"Plugins"`
}

// Validate validates this engine description
func (m *EngineDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlugins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineDescription) validatePlugins(formats strfmt.Registry) error {
	if swag.IsZero(m.Plugins) { // not required
		return nil
	}

	for i := 0; i < len(m.Plugins); i++ {
		if swag.IsZero(m.Plugins[i]) { // not required
			continue
		}

		if m.Plugins[i] != nil {
			if err := m.Plugins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Plugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this engine description based on the context it is used
func (m *EngineDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlugins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineDescription) contextValidatePlugins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Plugins); i++ {

		if m.Plugins[i] != nil {
			if err := m.Plugins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Plugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineDescription) UnmarshalBinary(b []byte) error {
	var res EngineDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EngineDescriptionPluginsItems0 engine description plugins items0
//
// swagger:model EngineDescriptionPluginsItems0
type EngineDescriptionPluginsItems0 struct {

	// name
	Name string `json:"Name,omitempty"`

	// type
	Type string `json:"Type,omitempty"`
}

// Validate validates this engine description plugins items0
func (m *EngineDescriptionPluginsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine description plugins items0 based on context it is used
func (m *EngineDescriptionPluginsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineDescriptionPluginsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineDescriptionPluginsItems0) UnmarshalBinary(b []byte) error {
	var res EngineDescriptionPluginsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
