// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// ContainerPauseReader is a Reader for the ContainerPause structure.
type ContainerPauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerPauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerPauseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerPauseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerPauseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerPauseNoContent creates a ContainerPauseNoContent with default headers values
func NewContainerPauseNoContent() *ContainerPauseNoContent {
	return &ContainerPauseNoContent{}
}

/* ContainerPauseNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerPauseNoContent struct {
}

// IsSuccess returns true when this container pause no content response has a 2xx status code
func (o *ContainerPauseNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container pause no content response has a 3xx status code
func (o *ContainerPauseNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container pause no content response has a 4xx status code
func (o *ContainerPauseNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this container pause no content response has a 5xx status code
func (o *ContainerPauseNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this container pause no content response a status code equal to that given
func (o *ContainerPauseNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ContainerPauseNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/pause][%d] containerPauseNoContent ", 204)
}

func (o *ContainerPauseNoContent) String() string {
	return fmt.Sprintf("[POST /containers/{id}/pause][%d] containerPauseNoContent ", 204)
}

func (o *ContainerPauseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerPauseNotFound creates a ContainerPauseNotFound with default headers values
func NewContainerPauseNotFound() *ContainerPauseNotFound {
	return &ContainerPauseNotFound{}
}

/* ContainerPauseNotFound describes a response with status code 404, with default header values.

no such container
*/
type ContainerPauseNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container pause not found response has a 2xx status code
func (o *ContainerPauseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container pause not found response has a 3xx status code
func (o *ContainerPauseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container pause not found response has a 4xx status code
func (o *ContainerPauseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container pause not found response has a 5xx status code
func (o *ContainerPauseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container pause not found response a status code equal to that given
func (o *ContainerPauseNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerPauseNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/pause][%d] containerPauseNotFound  %+v", 404, o.Payload)
}

func (o *ContainerPauseNotFound) String() string {
	return fmt.Sprintf("[POST /containers/{id}/pause][%d] containerPauseNotFound  %+v", 404, o.Payload)
}

func (o *ContainerPauseNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerPauseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerPauseInternalServerError creates a ContainerPauseInternalServerError with default headers values
func NewContainerPauseInternalServerError() *ContainerPauseInternalServerError {
	return &ContainerPauseInternalServerError{}
}

/* ContainerPauseInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerPauseInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container pause internal server error response has a 2xx status code
func (o *ContainerPauseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container pause internal server error response has a 3xx status code
func (o *ContainerPauseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container pause internal server error response has a 4xx status code
func (o *ContainerPauseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container pause internal server error response has a 5xx status code
func (o *ContainerPauseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container pause internal server error response a status code equal to that given
func (o *ContainerPauseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerPauseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/pause][%d] containerPauseInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerPauseInternalServerError) String() string {
	return fmt.Sprintf("[POST /containers/{id}/pause][%d] containerPauseInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerPauseInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerPauseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
