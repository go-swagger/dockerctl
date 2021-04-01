// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/dockerctl/models"
)

// ContainerListReader is a Reader for the ContainerList structure.
type ContainerListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContainerListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerListOK creates a ContainerListOK with default headers values
func NewContainerListOK() *ContainerListOK {
	return &ContainerListOK{}
}

/* ContainerListOK describes a response with status code 200, with default header values.

no error
*/
type ContainerListOK struct {
	Payload models.ContainerSummary
}

func (o *ContainerListOK) Error() string {
	return fmt.Sprintf("[GET /containers/json][%d] containerListOK  %+v", 200, o.Payload)
}
func (o *ContainerListOK) GetPayload() models.ContainerSummary {
	return o.Payload
}

func (o *ContainerListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerListBadRequest creates a ContainerListBadRequest with default headers values
func NewContainerListBadRequest() *ContainerListBadRequest {
	return &ContainerListBadRequest{}
}

/* ContainerListBadRequest describes a response with status code 400, with default header values.

bad parameter
*/
type ContainerListBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ContainerListBadRequest) Error() string {
	return fmt.Sprintf("[GET /containers/json][%d] containerListBadRequest  %+v", 400, o.Payload)
}
func (o *ContainerListBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerListInternalServerError creates a ContainerListInternalServerError with default headers values
func NewContainerListInternalServerError() *ContainerListInternalServerError {
	return &ContainerListInternalServerError{}
}

/* ContainerListInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/json][%d] containerListInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerListInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
