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

// ResourceObject An object describing the resources which can be advertised by a node and requested by a task
//
// swagger:model ResourceObject
type ResourceObject struct {

	// generic resources
	GenericResources GenericResources `json:"GenericResources,omitempty"`

	// memory bytes
	// Example: 8272408576
	MemoryBytes int64 `json:"MemoryBytes,omitempty"`

	// nano c p us
	// Example: 4000000000
	NanoCPUs int64 `json:"NanoCPUs,omitempty"`
}

// Validate validates this resource object
func (m *ResourceObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGenericResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceObject) validateGenericResources(formats strfmt.Registry) error {
	if swag.IsZero(m.GenericResources) { // not required
		return nil
	}

	if err := m.GenericResources.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("GenericResources")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("GenericResources")
		}
		return err
	}

	return nil
}

// ContextValidate validate this resource object based on the context it is used
func (m *ResourceObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGenericResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceObject) contextValidateGenericResources(ctx context.Context, formats strfmt.Registry) error {

	if err := m.GenericResources.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("GenericResources")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("GenericResources")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject) UnmarshalBinary(b []byte) error {
	var res ResourceObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
