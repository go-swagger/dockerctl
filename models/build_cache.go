// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BuildCache build cache
//
// swagger:model BuildCache
type BuildCache struct {

	// created at
	CreatedAt int64 `json:"CreatedAt,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// in use
	InUse bool `json:"InUse,omitempty"`

	// last used at
	LastUsedAt *int64 `json:"LastUsedAt,omitempty"`

	// parent
	Parent string `json:"Parent,omitempty"`

	// shared
	Shared bool `json:"Shared,omitempty"`

	// size
	Size int64 `json:"Size,omitempty"`

	// type
	Type string `json:"Type,omitempty"`

	// usage count
	UsageCount int64 `json:"UsageCount,omitempty"`
}

// Validate validates this build cache
func (m *BuildCache) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this build cache based on context it is used
func (m *BuildCache) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuildCache) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildCache) UnmarshalBinary(b []byte) error {
	var res BuildCache
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
