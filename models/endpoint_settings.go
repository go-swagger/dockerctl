// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointSettings Configuration for a network endpoint.
//
// swagger:model EndpointSettings
type EndpointSettings struct {

	// aliases
	// Example: ["server_x","server_y"]
	Aliases []string `json:"Aliases"`

	// DriverOpts is a mapping of driver options and values. These options
	// are passed directly to the driver and are driver specific.
	//
	// Example: {"com.example.some-label":"some-value","com.example.some-other-label":"some-other-value"}
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`

	// Unique ID for the service endpoint in a Sandbox.
	//
	// Example: b88f5b905aabf2893f3cbc4ee42d1ea7980bbc0a92e2c8922b1e1795298afb0b
	EndpointID string `json:"EndpointID,omitempty"`

	// Gateway address for this network.
	//
	// Example: 172.17.0.1
	Gateway string `json:"Gateway,omitempty"`

	// Global IPv6 address.
	//
	// Example: 2001:db8::5689
	GlobalIPV6Address string `json:"GlobalIPv6Address,omitempty"`

	// Mask length of the global IPv6 address.
	//
	// Example: 64
	GlobalIPV6PrefixLen int64 `json:"GlobalIPv6PrefixLen,omitempty"`

	// IP a m config
	IPAMConfig *EndpointIPAMConfig `json:"IPAMConfig,omitempty"`

	// IPv4 address.
	//
	// Example: 172.17.0.4
	IPAddress string `json:"IPAddress,omitempty"`

	// Mask length of the IPv4 address.
	//
	// Example: 16
	IPPrefixLen int64 `json:"IPPrefixLen,omitempty"`

	// IPv6 gateway address.
	//
	// Example: 2001:db8:2::100
	IPV6Gateway string `json:"IPv6Gateway,omitempty"`

	// links
	// Example: ["container_1","container_2"]
	Links []string `json:"Links"`

	// MAC address for the endpoint on this network.
	//
	// Example: 02:42:ac:11:00:04
	MacAddress string `json:"MacAddress,omitempty"`

	// Unique ID of the network.
	//
	// Example: 08754567f1f40222263eab4102e1c733ae697e8e354aa9cd6e18d7402835292a
	NetworkID string `json:"NetworkID,omitempty"`
}

// Validate validates this endpoint settings
func (m *EndpointSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAMConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointSettings) validateIPAMConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAMConfig) { // not required
		return nil
	}

	if m.IPAMConfig != nil {
		if err := m.IPAMConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IPAMConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this endpoint settings based on the context it is used
func (m *EndpointSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAMConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointSettings) contextValidateIPAMConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.IPAMConfig != nil {
		if err := m.IPAMConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IPAMConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointSettings) UnmarshalBinary(b []byte) error {
	var res EndpointSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
