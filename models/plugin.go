// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Plugin A plugin for the Engine API
//
// swagger:model Plugin
type Plugin struct {

	// config
	// Required: true
	Config PluginConfig `json:"Config"`

	// True if the plugin is running. False if the plugin is not running, only installed.
	// Example: true
	// Required: true
	Enabled bool `json:"Enabled"`

	// Id
	// Example: 5724e2c8652da337ab2eedd19fc6fc0ec908e4bd907c7421bf6a8dfc70c4c078
	ID string `json:"Id,omitempty"`

	// name
	// Example: tiborvass/sample-volume-plugin
	// Required: true
	Name string `json:"Name"`

	// plugin remote reference used to push/pull the plugin
	// Example: localhost:5000/tiborvass/sample-volume-plugin:latest
	PluginReference string `json:"PluginReference,omitempty"`

	// settings
	// Required: true
	Settings PluginSettings `json:"Settings"`
}

// Validate validates this plugin
func (m *Plugin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plugin) validateConfig(formats strfmt.Registry) error {

	if err := m.Config.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config")
		}
		return err
	}

	return nil
}

func (m *Plugin) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("Enabled", "body", bool(m.Enabled)); err != nil {
		return err
	}

	return nil
}

func (m *Plugin) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Plugin) validateSettings(formats strfmt.Registry) error {

	if err := m.Settings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Settings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Settings")
		}
		return err
	}

	return nil
}

// ContextValidate validate this plugin based on the context it is used
func (m *Plugin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plugin) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Config.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config")
		}
		return err
	}

	return nil
}

func (m *Plugin) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Settings.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Settings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Settings")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plugin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plugin) UnmarshalBinary(b []byte) error {
	var res Plugin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginConfig The config of a plugin.
//
// swagger:model PluginConfig
type PluginConfig struct {

	// args
	// Required: true
	Args PluginConfigArgs `json:"Args"`

	// description
	// Example: A sample volume plugin for Docker
	// Required: true
	Description string `json:"Description"`

	// Docker Version used to create the plugin
	// Example: 17.06.0-ce
	DockerVersion string `json:"DockerVersion,omitempty"`

	// documentation
	// Example: https://docs.docker.com/engine/extend/plugins/
	// Required: true
	Documentation string `json:"Documentation"`

	// entrypoint
	// Example: ["/usr/bin/sample-volume-plugin","/data"]
	// Required: true
	Entrypoint []string `json:"Entrypoint"`

	// env
	// Example: [{"Description":"If set, prints debug messages","Name":"DEBUG","Settable":null,"Value":"0"}]
	// Required: true
	Env []PluginEnv `json:"Env"`

	// interface
	// Required: true
	Interface PluginConfigInterface `json:"Interface"`

	// ipc host
	// Example: false
	// Required: true
	IpcHost bool `json:"IpcHost"`

	// linux
	// Required: true
	Linux PluginConfigLinux `json:"Linux"`

	// mounts
	// Required: true
	Mounts []PluginMount `json:"Mounts"`

	// network
	// Required: true
	Network PluginConfigNetwork `json:"Network"`

	// pid host
	// Example: false
	// Required: true
	PidHost bool `json:"PidHost"`

	// propagated mount
	// Example: /mnt/volumes
	// Required: true
	PropagatedMount string `json:"PropagatedMount"`

	// user
	User PluginConfigUser `json:"User,omitempty"`

	// work dir
	// Example: /bin/
	// Required: true
	WorkDir string `json:"WorkDir"`

	// rootfs
	Rootfs *PluginConfigRootfs `json:"rootfs,omitempty"`
}

// Validate validates this plugin config
func (m *PluginConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocumentation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntrypoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpcHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinux(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePidHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropagatedMount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootfs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfig) validateArgs(formats strfmt.Registry) error {

	if err := m.Args.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "Args")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "Args")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("Config"+"."+"Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateDocumentation(formats strfmt.Registry) error {

	if err := validate.RequiredString("Config"+"."+"Documentation", "body", m.Documentation); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateEntrypoint(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"Entrypoint", "body", m.Entrypoint); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateEnv(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"Env", "body", m.Env); err != nil {
		return err
	}

	for i := 0; i < len(m.Env); i++ {

		if err := m.Env[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "Env" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "Env" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PluginConfig) validateInterface(formats strfmt.Registry) error {

	if err := m.Interface.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "Interface")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "Interface")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateIpcHost(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"IpcHost", "body", bool(m.IpcHost)); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateLinux(formats strfmt.Registry) error {

	if err := m.Linux.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "Linux")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "Linux")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateMounts(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"Mounts", "body", m.Mounts); err != nil {
		return err
	}

	for i := 0; i < len(m.Mounts); i++ {

		if err := m.Mounts[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "Mounts" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "Mounts" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PluginConfig) validateNetwork(formats strfmt.Registry) error {

	if err := m.Network.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "Network")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "Network")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validatePidHost(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"PidHost", "body", bool(m.PidHost)); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validatePropagatedMount(formats strfmt.Registry) error {

	if err := validate.RequiredString("Config"+"."+"PropagatedMount", "body", m.PropagatedMount); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if err := m.User.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "User")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "User")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateWorkDir(formats strfmt.Registry) error {

	if err := validate.RequiredString("Config"+"."+"WorkDir", "body", m.WorkDir); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateRootfs(formats strfmt.Registry) error {
	if swag.IsZero(m.Rootfs) { // not required
		return nil
	}

	if m.Rootfs != nil {
		if err := m.Rootfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "rootfs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "rootfs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plugin config based on the context it is used
func (m *PluginConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArgs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinux(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRootfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfig) contextValidateArgs(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Args.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "Args")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "Args")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Env); i++ {

		if err := m.Env[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "Env" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "Env" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PluginConfig) contextValidateInterface(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Interface.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "Interface")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "Interface")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) contextValidateLinux(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Linux.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "Linux")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "Linux")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) contextValidateMounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mounts); i++ {

		if err := m.Mounts[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "Mounts" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "Mounts" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PluginConfig) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Network.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "Network")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "Network")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if err := m.User.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config" + "." + "User")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Config" + "." + "User")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) contextValidateRootfs(ctx context.Context, formats strfmt.Registry) error {

	if m.Rootfs != nil {
		if err := m.Rootfs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "rootfs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "rootfs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfig) UnmarshalBinary(b []byte) error {
	var res PluginConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginConfigArgs plugin config args
//
// swagger:model PluginConfigArgs
type PluginConfigArgs struct {

	// description
	// Example: command line arguments
	// Required: true
	Description string `json:"Description"`

	// name
	// Example: args
	// Required: true
	Name string `json:"Name"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`

	// value
	// Required: true
	Value []string `json:"Value"`
}

// Validate validates this plugin config args
func (m *PluginConfigArgs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigArgs) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("Config"+"."+"Args"+"."+"Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigArgs) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("Config"+"."+"Args"+"."+"Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigArgs) validateSettable(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"Args"+"."+"Settable", "body", m.Settable); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigArgs) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"Args"+"."+"Value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugin config args based on context it is used
func (m *PluginConfigArgs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigArgs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigArgs) UnmarshalBinary(b []byte) error {
	var res PluginConfigArgs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginConfigInterface The interface between Docker and the plugin
//
// swagger:model PluginConfigInterface
type PluginConfigInterface struct {

	// Protocol to use for clients connecting to the plugin.
	// Example: some.protocol/v1.0
	// Enum: [ moby.plugins.http/v1]
	ProtocolScheme string `json:"ProtocolScheme,omitempty"`

	// socket
	// Example: plugins.sock
	// Required: true
	Socket string `json:"Socket"`

	// types
	// Example: ["docker.volumedriver/1.0"]
	// Required: true
	Types []PluginInterfaceType `json:"Types"`
}

// Validate validates this plugin config interface
func (m *PluginConfigInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocolScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSocket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pluginConfigInterfaceTypeProtocolSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","moby.plugins.http/v1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pluginConfigInterfaceTypeProtocolSchemePropEnum = append(pluginConfigInterfaceTypeProtocolSchemePropEnum, v)
	}
}

const (

	// PluginConfigInterfaceProtocolSchemeEmpty captures enum value ""
	PluginConfigInterfaceProtocolSchemeEmpty string = ""

	// PluginConfigInterfaceProtocolSchemeMobyDotPluginsDotHTTPV1 captures enum value "moby.plugins.http/v1"
	PluginConfigInterfaceProtocolSchemeMobyDotPluginsDotHTTPV1 string = "moby.plugins.http/v1"
)

// prop value enum
func (m *PluginConfigInterface) validateProtocolSchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pluginConfigInterfaceTypeProtocolSchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PluginConfigInterface) validateProtocolScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtocolScheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolSchemeEnum("Config"+"."+"Interface"+"."+"ProtocolScheme", "body", m.ProtocolScheme); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigInterface) validateSocket(formats strfmt.Registry) error {

	if err := validate.RequiredString("Config"+"."+"Interface"+"."+"Socket", "body", m.Socket); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigInterface) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"Interface"+"."+"Types", "body", m.Types); err != nil {
		return err
	}

	for i := 0; i < len(m.Types); i++ {

		if err := m.Types[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "Interface" + "." + "Types" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "Interface" + "." + "Types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this plugin config interface based on the context it is used
func (m *PluginConfigInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigInterface) contextValidateTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Types); i++ {

		if err := m.Types[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "Interface" + "." + "Types" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "Interface" + "." + "Types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigInterface) UnmarshalBinary(b []byte) error {
	var res PluginConfigInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginConfigLinux plugin config linux
//
// swagger:model PluginConfigLinux
type PluginConfigLinux struct {

	// allow all devices
	// Example: false
	// Required: true
	AllowAllDevices bool `json:"AllowAllDevices"`

	// capabilities
	// Example: ["CAP_SYS_ADMIN","CAP_SYSLOG"]
	// Required: true
	Capabilities []string `json:"Capabilities"`

	// devices
	// Required: true
	Devices []PluginDevice `json:"Devices"`
}

// Validate validates this plugin config linux
func (m *PluginConfigLinux) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowAllDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigLinux) validateAllowAllDevices(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"Linux"+"."+"AllowAllDevices", "body", bool(m.AllowAllDevices)); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigLinux) validateCapabilities(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"Linux"+"."+"Capabilities", "body", m.Capabilities); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigLinux) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("Config"+"."+"Linux"+"."+"Devices", "body", m.Devices); err != nil {
		return err
	}

	for i := 0; i < len(m.Devices); i++ {

		if err := m.Devices[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "Linux" + "." + "Devices" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "Linux" + "." + "Devices" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this plugin config linux based on the context it is used
func (m *PluginConfigLinux) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigLinux) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Devices); i++ {

		if err := m.Devices[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config" + "." + "Linux" + "." + "Devices" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Config" + "." + "Linux" + "." + "Devices" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigLinux) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigLinux) UnmarshalBinary(b []byte) error {
	var res PluginConfigLinux
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginConfigNetwork plugin config network
//
// swagger:model PluginConfigNetwork
type PluginConfigNetwork struct {

	// type
	// Example: host
	// Required: true
	Type string `json:"Type"`
}

// Validate validates this plugin config network
func (m *PluginConfigNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigNetwork) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("Config"+"."+"Network"+"."+"Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugin config network based on context it is used
func (m *PluginConfigNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigNetwork) UnmarshalBinary(b []byte) error {
	var res PluginConfigNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginConfigRootfs plugin config rootfs
//
// swagger:model PluginConfigRootfs
type PluginConfigRootfs struct {

	// diff ids
	// Example: ["sha256:675532206fbf3030b8458f88d6e26d4eb1577688a25efec97154c94e8b6b4887","sha256:e216a057b1cb1efc11f8a268f37ef62083e70b1b38323ba252e25ac88904a7e8"]
	DiffIds []string `json:"diff_ids"`

	// type
	// Example: layers
	Type string `json:"type,omitempty"`
}

// Validate validates this plugin config rootfs
func (m *PluginConfigRootfs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plugin config rootfs based on context it is used
func (m *PluginConfigRootfs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigRootfs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigRootfs) UnmarshalBinary(b []byte) error {
	var res PluginConfigRootfs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginConfigUser plugin config user
//
// swagger:model PluginConfigUser
type PluginConfigUser struct {

	// g ID
	// Example: 1000
	GID uint32 `json:"GID,omitempty"`

	// UID
	// Example: 1000
	UID uint32 `json:"UID,omitempty"`
}

// Validate validates this plugin config user
func (m *PluginConfigUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plugin config user based on context it is used
func (m *PluginConfigUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigUser) UnmarshalBinary(b []byte) error {
	var res PluginConfigUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginSettings Settings that can be modified by users.
//
// swagger:model PluginSettings
type PluginSettings struct {

	// args
	// Required: true
	Args []string `json:"Args"`

	// devices
	// Required: true
	Devices []PluginDevice `json:"Devices"`

	// env
	// Example: ["DEBUG=0"]
	// Required: true
	Env []string `json:"Env"`

	// mounts
	// Required: true
	Mounts []PluginMount `json:"Mounts"`
}

// Validate validates this plugin settings
func (m *PluginSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginSettings) validateArgs(formats strfmt.Registry) error {

	if err := validate.Required("Settings"+"."+"Args", "body", m.Args); err != nil {
		return err
	}

	return nil
}

func (m *PluginSettings) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("Settings"+"."+"Devices", "body", m.Devices); err != nil {
		return err
	}

	for i := 0; i < len(m.Devices); i++ {

		if err := m.Devices[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Settings" + "." + "Devices" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Settings" + "." + "Devices" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PluginSettings) validateEnv(formats strfmt.Registry) error {

	if err := validate.Required("Settings"+"."+"Env", "body", m.Env); err != nil {
		return err
	}

	return nil
}

func (m *PluginSettings) validateMounts(formats strfmt.Registry) error {

	if err := validate.Required("Settings"+"."+"Mounts", "body", m.Mounts); err != nil {
		return err
	}

	for i := 0; i < len(m.Mounts); i++ {

		if err := m.Mounts[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Settings" + "." + "Mounts" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Settings" + "." + "Mounts" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this plugin settings based on the context it is used
func (m *PluginSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginSettings) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Devices); i++ {

		if err := m.Devices[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Settings" + "." + "Devices" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Settings" + "." + "Devices" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PluginSettings) contextValidateMounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mounts); i++ {

		if err := m.Mounts[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Settings" + "." + "Mounts" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Settings" + "." + "Mounts" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginSettings) UnmarshalBinary(b []byte) error {
	var res PluginSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
