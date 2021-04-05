// Code generated by go-swagger; DO NOT EDIT.

package exec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-swagger/dockerctl/models"
)

// ExecInspectReader is a Reader for the ExecInspect structure.
type ExecInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewExecInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExecInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecInspectOK creates a ExecInspectOK with default headers values
func NewExecInspectOK() *ExecInspectOK {
	return &ExecInspectOK{}
}

/* ExecInspectOK describes a response with status code 200, with default header values.

No error
*/
type ExecInspectOK struct {
	Payload *ExecInspectOKBody
}

func (o *ExecInspectOK) Error() string {
	return fmt.Sprintf("[GET /exec/{id}/json][%d] execInspectOK  %+v", 200, o.Payload)
}
func (o *ExecInspectOK) GetPayload() *ExecInspectOKBody {
	return o.Payload
}

func (o *ExecInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExecInspectOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecInspectNotFound creates a ExecInspectNotFound with default headers values
func NewExecInspectNotFound() *ExecInspectNotFound {
	return &ExecInspectNotFound{}
}

/* ExecInspectNotFound describes a response with status code 404, with default header values.

No such exec instance
*/
type ExecInspectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ExecInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /exec/{id}/json][%d] execInspectNotFound  %+v", 404, o.Payload)
}
func (o *ExecInspectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExecInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecInspectInternalServerError creates a ExecInspectInternalServerError with default headers values
func NewExecInspectInternalServerError() *ExecInspectInternalServerError {
	return &ExecInspectInternalServerError{}
}

/* ExecInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ExecInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ExecInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /exec/{id}/json][%d] execInspectInternalServerError  %+v", 500, o.Payload)
}
func (o *ExecInspectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExecInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ExecInspectOKBody ExecInspectResponse
swagger:model ExecInspectOKBody
*/
type ExecInspectOKBody struct {

	// can remove
	CanRemove bool `json:"CanRemove,omitempty"`

	// container ID
	ContainerID string `json:"ContainerID,omitempty"`

	// detach keys
	DetachKeys string `json:"DetachKeys,omitempty"`

	// exit code
	ExitCode int64 `json:"ExitCode,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// open stderr
	OpenStderr bool `json:"OpenStderr,omitempty"`

	// open stdin
	OpenStdin bool `json:"OpenStdin,omitempty"`

	// open stdout
	OpenStdout bool `json:"OpenStdout,omitempty"`

	// The system process ID for the exec process.
	Pid int64 `json:"Pid,omitempty"`

	// process config
	ProcessConfig *models.ProcessConfig `json:"ProcessConfig,omitempty"`

	// running
	Running bool `json:"Running,omitempty"`
}

// Validate validates this exec inspect o k body
func (o *ExecInspectOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProcessConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExecInspectOKBody) validateProcessConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.ProcessConfig) { // not required
		return nil
	}

	if o.ProcessConfig != nil {
		if err := o.ProcessConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execInspectOK" + "." + "ProcessConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this exec inspect o k body based on the context it is used
func (o *ExecInspectOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateProcessConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExecInspectOKBody) contextValidateProcessConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.ProcessConfig != nil {
		if err := o.ProcessConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execInspectOK" + "." + "ProcessConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExecInspectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExecInspectOKBody) UnmarshalBinary(b []byte) error {
	var res ExecInspectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
