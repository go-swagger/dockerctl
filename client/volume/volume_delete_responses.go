// Code generated by go-swagger; DO NOT EDIT.

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/dockerctl/models"
)

// VolumeDeleteReader is a Reader for the VolumeDelete structure.
type VolumeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVolumeDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewVolumeDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewVolumeDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeDeleteNoContent creates a VolumeDeleteNoContent with default headers values
func NewVolumeDeleteNoContent() *VolumeDeleteNoContent {
	return &VolumeDeleteNoContent{}
}

/* VolumeDeleteNoContent describes a response with status code 204, with default header values.

The volume was removed
*/
type VolumeDeleteNoContent struct {
}

func (o *VolumeDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{name}][%d] volumeDeleteNoContent ", 204)
}

func (o *VolumeDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVolumeDeleteNotFound creates a VolumeDeleteNotFound with default headers values
func NewVolumeDeleteNotFound() *VolumeDeleteNotFound {
	return &VolumeDeleteNotFound{}
}

/* VolumeDeleteNotFound describes a response with status code 404, with default header values.

No such volume or volume driver
*/
type VolumeDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *VolumeDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{name}][%d] volumeDeleteNotFound  %+v", 404, o.Payload)
}
func (o *VolumeDeleteNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeDeleteConflict creates a VolumeDeleteConflict with default headers values
func NewVolumeDeleteConflict() *VolumeDeleteConflict {
	return &VolumeDeleteConflict{}
}

/* VolumeDeleteConflict describes a response with status code 409, with default header values.

Volume is in use and cannot be removed
*/
type VolumeDeleteConflict struct {
	Payload *models.ErrorResponse
}

func (o *VolumeDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{name}][%d] volumeDeleteConflict  %+v", 409, o.Payload)
}
func (o *VolumeDeleteConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeDeleteInternalServerError creates a VolumeDeleteInternalServerError with default headers values
func NewVolumeDeleteInternalServerError() *VolumeDeleteInternalServerError {
	return &VolumeDeleteInternalServerError{}
}

/* VolumeDeleteInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type VolumeDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *VolumeDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{name}][%d] volumeDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *VolumeDeleteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}