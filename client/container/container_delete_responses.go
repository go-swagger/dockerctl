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

// ContainerDeleteReader is a Reader for the ContainerDelete structure.
type ContainerDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContainerDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewContainerDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerDeleteNoContent creates a ContainerDeleteNoContent with default headers values
func NewContainerDeleteNoContent() *ContainerDeleteNoContent {
	return &ContainerDeleteNoContent{}
}

/* ContainerDeleteNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerDeleteNoContent struct {
}

// IsSuccess returns true when this container delete no content response has a 2xx status code
func (o *ContainerDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container delete no content response has a 3xx status code
func (o *ContainerDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete no content response has a 4xx status code
func (o *ContainerDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this container delete no content response has a 5xx status code
func (o *ContainerDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this container delete no content response a status code equal to that given
func (o *ContainerDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ContainerDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteNoContent ", 204)
}

func (o *ContainerDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteNoContent ", 204)
}

func (o *ContainerDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerDeleteBadRequest creates a ContainerDeleteBadRequest with default headers values
func NewContainerDeleteBadRequest() *ContainerDeleteBadRequest {
	return &ContainerDeleteBadRequest{}
}

/* ContainerDeleteBadRequest describes a response with status code 400, with default header values.

bad parameter
*/
type ContainerDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container delete bad request response has a 2xx status code
func (o *ContainerDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container delete bad request response has a 3xx status code
func (o *ContainerDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete bad request response has a 4xx status code
func (o *ContainerDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this container delete bad request response has a 5xx status code
func (o *ContainerDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this container delete bad request response a status code equal to that given
func (o *ContainerDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ContainerDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerDeleteBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerDeleteNotFound creates a ContainerDeleteNotFound with default headers values
func NewContainerDeleteNotFound() *ContainerDeleteNotFound {
	return &ContainerDeleteNotFound{}
}

/* ContainerDeleteNotFound describes a response with status code 404, with default header values.

no such container
*/
type ContainerDeleteNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container delete not found response has a 2xx status code
func (o *ContainerDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container delete not found response has a 3xx status code
func (o *ContainerDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete not found response has a 4xx status code
func (o *ContainerDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container delete not found response has a 5xx status code
func (o *ContainerDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container delete not found response a status code equal to that given
func (o *ContainerDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ContainerDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ContainerDeleteNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerDeleteConflict creates a ContainerDeleteConflict with default headers values
func NewContainerDeleteConflict() *ContainerDeleteConflict {
	return &ContainerDeleteConflict{}
}

/* ContainerDeleteConflict describes a response with status code 409, with default header values.

conflict
*/
type ContainerDeleteConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container delete conflict response has a 2xx status code
func (o *ContainerDeleteConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container delete conflict response has a 3xx status code
func (o *ContainerDeleteConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete conflict response has a 4xx status code
func (o *ContainerDeleteConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this container delete conflict response has a 5xx status code
func (o *ContainerDeleteConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this container delete conflict response a status code equal to that given
func (o *ContainerDeleteConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ContainerDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteConflict  %+v", 409, o.Payload)
}

func (o *ContainerDeleteConflict) String() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteConflict  %+v", 409, o.Payload)
}

func (o *ContainerDeleteConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerDeleteInternalServerError creates a ContainerDeleteInternalServerError with default headers values
func NewContainerDeleteInternalServerError() *ContainerDeleteInternalServerError {
	return &ContainerDeleteInternalServerError{}
}

/* ContainerDeleteInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container delete internal server error response has a 2xx status code
func (o *ContainerDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container delete internal server error response has a 3xx status code
func (o *ContainerDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete internal server error response has a 4xx status code
func (o *ContainerDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container delete internal server error response has a 5xx status code
func (o *ContainerDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container delete internal server error response a status code equal to that given
func (o *ContainerDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerDeleteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
