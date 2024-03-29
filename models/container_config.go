// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContainerConfig Configuration for a container that is portable between hosts
//
// swagger:model ContainerConfig
type ContainerConfig struct {

	// Command is already escaped (Windows only)
	ArgsEscaped bool `json:"ArgsEscaped,omitempty"`

	// Whether to attach to `stderr`.
	AttachStderr *bool `json:"AttachStderr,omitempty"`

	// Whether to attach to `stdin`.
	AttachStdin *bool `json:"AttachStdin,omitempty"`

	// Whether to attach to `stdout`.
	AttachStdout *bool `json:"AttachStdout,omitempty"`

	// Command to run specified as a string or an array of strings.
	Cmd []string `json:"Cmd"`

	// The domain name to use for the container.
	Domainname string `json:"Domainname,omitempty"`

	// The entry point for the container as a string or an array of strings.
	//
	// If the array consists of exactly one empty string (`[""]`) then the entry point is reset to system default (i.e., the entry point used by docker when there is no `ENTRYPOINT` instruction in the `Dockerfile`).
	//
	Entrypoint []string `json:"Entrypoint"`

	// A list of environment variables to set inside the container in the form `["VAR=value", ...]`. A variable without `=` is removed from the environment, rather than to have an empty value.
	//
	Env []string `json:"Env"`

	// An object mapping ports to an empty object in the form:
	//
	// `{"<port>/<tcp|udp|sctp>": {}}`
	//
	ExposedPorts map[string]interface{} `json:"ExposedPorts,omitempty"`

	// healthcheck
	Healthcheck *HealthConfig `json:"Healthcheck,omitempty"`

	// The hostname to use for the container, as a valid RFC 1123 hostname.
	Hostname string `json:"Hostname,omitempty"`

	// The name of the image to use when creating the container
	Image string `json:"Image,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// MAC address of the container.
	MacAddress string `json:"MacAddress,omitempty"`

	// Disable networking for the container.
	NetworkDisabled bool `json:"NetworkDisabled,omitempty"`

	// `ONBUILD` metadata that were defined in the image's `Dockerfile`.
	OnBuild []string `json:"OnBuild"`

	// Open `stdin`
	OpenStdin *bool `json:"OpenStdin,omitempty"`

	// Shell for when `RUN`, `CMD`, and `ENTRYPOINT` uses a shell.
	Shell []string `json:"Shell"`

	// Close `stdin` after one attached client disconnects
	StdinOnce *bool `json:"StdinOnce,omitempty"`

	// Signal to stop a container as a string or unsigned integer.
	StopSignal *string `json:"StopSignal,omitempty"`

	// Timeout to stop a container in seconds.
	StopTimeout *int64 `json:"StopTimeout,omitempty"`

	// Attach standard streams to a TTY, including `stdin` if it is not closed.
	Tty *bool `json:"Tty,omitempty"`

	// The user that commands are run as inside the container.
	User string `json:"User,omitempty"`

	// An object mapping mount point paths inside the container to empty objects.
	Volumes map[string]interface{} `json:"Volumes,omitempty"`

	// The working directory for commands to run in.
	WorkingDir string `json:"WorkingDir,omitempty"`
}

// Validate validates this container config
func (m *ContainerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExposedPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthcheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// additional properties value enum
var containerConfigExposedPortsValueEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`[{}]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerConfigExposedPortsValueEnum = append(containerConfigExposedPortsValueEnum, v)
	}
}

func (m *ContainerConfig) validateExposedPortsValueEnum(path, location string, value interface{}) error {
	if err := validate.EnumCase(path, location, value, containerConfigExposedPortsValueEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContainerConfig) validateExposedPorts(formats strfmt.Registry) error {
	if swag.IsZero(m.ExposedPorts) { // not required
		return nil
	}

	for k := range m.ExposedPorts {

		if err := m.validateExposedPortsValueEnum("ExposedPorts"+"."+k, "body", m.ExposedPorts[k]); err != nil {
			return err
		}
	}

	return nil
}

func (m *ContainerConfig) validateHealthcheck(formats strfmt.Registry) error {
	if swag.IsZero(m.Healthcheck) { // not required
		return nil
	}

	if m.Healthcheck != nil {
		if err := m.Healthcheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Healthcheck")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Healthcheck")
			}
			return err
		}
	}

	return nil
}

// additional properties value enum
var containerConfigVolumesValueEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`[{}]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerConfigVolumesValueEnum = append(containerConfigVolumesValueEnum, v)
	}
}

func (m *ContainerConfig) validateVolumesValueEnum(path, location string, value interface{}) error {
	if err := validate.EnumCase(path, location, value, containerConfigVolumesValueEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContainerConfig) validateVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for k := range m.Volumes {

		if err := m.validateVolumesValueEnum("Volumes"+"."+k, "body", m.Volumes[k]); err != nil {
			return err
		}
	}

	return nil
}

// ContextValidate validate this container config based on the context it is used
func (m *ContainerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHealthcheck(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerConfig) contextValidateHealthcheck(ctx context.Context, formats strfmt.Registry) error {

	if m.Healthcheck != nil {
		if err := m.Healthcheck.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Healthcheck")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Healthcheck")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerConfig) UnmarshalBinary(b []byte) error {
	var res ContainerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
