// Code generated by go-swagger; DO NOT EDIT.

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// VolumeInspectReader is a Reader for the VolumeInspect structure.
type VolumeInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewVolumeInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeInspectOK creates a VolumeInspectOK with default headers values
func NewVolumeInspectOK() *VolumeInspectOK {
	return &VolumeInspectOK{}
}

/* VolumeInspectOK describes a response with status code 200, with default header values.

No error
*/
type VolumeInspectOK struct {
	Payload *models.Volume
}

// IsSuccess returns true when this volume inspect o k response has a 2xx status code
func (o *VolumeInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume inspect o k response has a 3xx status code
func (o *VolumeInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume inspect o k response has a 4xx status code
func (o *VolumeInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume inspect o k response has a 5xx status code
func (o *VolumeInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume inspect o k response a status code equal to that given
func (o *VolumeInspectOK) IsCode(code int) bool {
	return code == 200
}

func (o *VolumeInspectOK) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectOK  %+v", 200, o.Payload)
}

func (o *VolumeInspectOK) String() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectOK  %+v", 200, o.Payload)
}

func (o *VolumeInspectOK) GetPayload() *models.Volume {
	return o.Payload
}

func (o *VolumeInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectNotFound creates a VolumeInspectNotFound with default headers values
func NewVolumeInspectNotFound() *VolumeInspectNotFound {
	return &VolumeInspectNotFound{}
}

/* VolumeInspectNotFound describes a response with status code 404, with default header values.

No such volume
*/
type VolumeInspectNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this volume inspect not found response has a 2xx status code
func (o *VolumeInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume inspect not found response has a 3xx status code
func (o *VolumeInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume inspect not found response has a 4xx status code
func (o *VolumeInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume inspect not found response has a 5xx status code
func (o *VolumeInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this volume inspect not found response a status code equal to that given
func (o *VolumeInspectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *VolumeInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectNotFound  %+v", 404, o.Payload)
}

func (o *VolumeInspectNotFound) String() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectNotFound  %+v", 404, o.Payload)
}

func (o *VolumeInspectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectInternalServerError creates a VolumeInspectInternalServerError with default headers values
func NewVolumeInspectInternalServerError() *VolumeInspectInternalServerError {
	return &VolumeInspectInternalServerError{}
}

/* VolumeInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type VolumeInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this volume inspect internal server error response has a 2xx status code
func (o *VolumeInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume inspect internal server error response has a 3xx status code
func (o *VolumeInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume inspect internal server error response has a 4xx status code
func (o *VolumeInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume inspect internal server error response has a 5xx status code
func (o *VolumeInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume inspect internal server error response a status code equal to that given
func (o *VolumeInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *VolumeInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInspectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
