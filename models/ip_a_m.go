// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IPAM IP a m
//
// swagger:model IPAM
type IPAM struct {

	// List of IPAM configuration options, specified as a map: `{"Subnet": <CIDR>, "IPRange": <CIDR>, "Gateway": <IP address>, "AuxAddress": <device_name:IP address>}`
	Config []map[string]string `json:"Config"`

	// Name of the IPAM driver to use.
	Driver *string `json:"Driver,omitempty"`

	// Driver-specific options, specified as a map.
	Options map[string]string `json:"Options,omitempty"`
}

// Validate validates this IP a m
func (m *IPAM) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP a m based on context it is used
func (m *IPAM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPAM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAM) UnmarshalBinary(b []byte) error {
	var res IPAM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
