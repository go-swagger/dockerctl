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

// Task task
// Example: {"AssignedGenericResources":[{"DiscreteResourceSpec":{"Kind":"SSD","Value":3}},{"NamedResourceSpec":{"Kind":"GPU","Value":"UUID1"}},{"NamedResourceSpec":{"Kind":"GPU","Value":"UUID2"}}],"CreatedAt":"2016-06-07T21:07:31.171892745Z","DesiredState":"running","ID":"0kzzo1i0y4jz6027t0k7aezc7","NetworksAttachments":[{"Addresses":["10.255.0.10/16"],"Network":{"CreatedAt":"2016-06-07T20:31:11.912919752Z","DriverState":{"Name":"overlay","Options":{"com.docker.network.driver.overlay.vxlanid_list":"256"}},"ID":"4qvuz4ko70xaltuqbt8956gd1","IPAMOptions":{"Configs":[{"Gateway":"10.255.0.1","Subnet":"10.255.0.0/16"}],"Driver":{"Name":"default"}},"Spec":{"DriverConfiguration":{},"IPAMOptions":{"Configs":[{"Gateway":"10.255.0.1","Subnet":"10.255.0.0/16"}],"Driver":{}},"Labels":{"com.docker.swarm.internal":"true"},"Name":"ingress"},"UpdatedAt":"2016-06-07T21:07:29.955277358Z","Version":{"Index":18}}}],"NodeID":"60gvrl6tm78dmak4yl7srz94v","ServiceID":"9mnpnzenvg8p8tdbtq4wvbkcz","Slot":1,"Spec":{"ContainerSpec":{"Image":"redis"},"Placement":{},"Resources":{"Limits":{},"Reservations":{}},"RestartPolicy":{"Condition":"any","MaxAttempts":0}},"Status":{"ContainerStatus":{"ContainerID":"e5d62702a1b48d01c3e02ca1e0212a250801fa8d67caca0b6f35919ebc12f035","PID":677},"Message":"started","State":"running","Timestamp":"2016-06-07T21:07:31.290032978Z"},"UpdatedAt":"2016-06-07T21:07:31.376370513Z","Version":{"Index":71}}
//
// swagger:model Task
type Task struct {

	// assigned generic resources
	AssignedGenericResources GenericResources `json:"AssignedGenericResources,omitempty"`

	// created at
	CreatedAt string `json:"CreatedAt,omitempty"`

	// desired state
	DesiredState TaskState `json:"DesiredState,omitempty"`

	// The ID of the task.
	ID string `json:"ID,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// Name of the task.
	Name string `json:"Name,omitempty"`

	// The ID of the node that this task is on.
	NodeID string `json:"NodeID,omitempty"`

	// The ID of the service this task is part of.
	ServiceID string `json:"ServiceID,omitempty"`

	// slot
	Slot int64 `json:"Slot,omitempty"`

	// spec
	Spec *TaskSpec `json:"Spec,omitempty"`

	// status
	Status *TaskStatus `json:"Status,omitempty"`

	// updated at
	UpdatedAt string `json:"UpdatedAt,omitempty"`

	// version
	Version *ObjectVersion `json:"Version,omitempty"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedGenericResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateAssignedGenericResources(formats strfmt.Registry) error {
	if swag.IsZero(m.AssignedGenericResources) { // not required
		return nil
	}

	if err := m.AssignedGenericResources.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AssignedGenericResources")
		}
		return err
	}

	return nil
}

func (m *Task) validateDesiredState(formats strfmt.Registry) error {
	if swag.IsZero(m.DesiredState) { // not required
		return nil
	}

	if err := m.DesiredState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("DesiredState")
		}
		return err
	}

	return nil
}

func (m *Task) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

func (m *Task) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

func (m *Task) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this task based on the context it is used
func (m *Task) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignedGenericResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDesiredState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) contextValidateAssignedGenericResources(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AssignedGenericResources.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AssignedGenericResources")
		}
		return err
	}

	return nil
}

func (m *Task) contextValidateDesiredState(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DesiredState.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("DesiredState")
		}
		return err
	}

	return nil
}

func (m *Task) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.Spec != nil {
		if err := m.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

func (m *Task) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

func (m *Task) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {
		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Task) UnmarshalBinary(b []byte) error {
	var res Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TaskStatus task status
//
// swagger:model TaskStatus
type TaskStatus struct {

	// container status
	ContainerStatus *TaskStatusContainerStatus `json:"ContainerStatus,omitempty"`

	// err
	Err string `json:"Err,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// state
	State TaskState `json:"State,omitempty"`

	// timestamp
	Timestamp string `json:"Timestamp,omitempty"`
}

// Validate validates this task status
func (m *TaskStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskStatus) validateContainerStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerStatus) { // not required
		return nil
	}

	if m.ContainerStatus != nil {
		if err := m.ContainerStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status" + "." + "ContainerStatus")
			}
			return err
		}
	}

	return nil
}

func (m *TaskStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Status" + "." + "State")
		}
		return err
	}

	return nil
}

// ContextValidate validate this task status based on the context it is used
func (m *TaskStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskStatus) contextValidateContainerStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerStatus != nil {
		if err := m.ContainerStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status" + "." + "ContainerStatus")
			}
			return err
		}
	}

	return nil
}

func (m *TaskStatus) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Status" + "." + "State")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskStatus) UnmarshalBinary(b []byte) error {
	var res TaskStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TaskStatusContainerStatus task status container status
//
// swagger:model TaskStatusContainerStatus
type TaskStatusContainerStatus struct {

	// container ID
	ContainerID string `json:"ContainerID,omitempty"`

	// exit code
	ExitCode int64 `json:"ExitCode,omitempty"`

	// p ID
	PID int64 `json:"PID,omitempty"`
}

// Validate validates this task status container status
func (m *TaskStatusContainerStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this task status container status based on context it is used
func (m *TaskStatusContainerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskStatusContainerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskStatusContainerStatus) UnmarshalBinary(b []byte) error {
	var res TaskStatusContainerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}