// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/dockerctl/models"
)

// PluginSetReader is a Reader for the PluginSet structure.
type PluginSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginSetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPluginSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPluginSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginSetNoContent creates a PluginSetNoContent with default headers values
func NewPluginSetNoContent() *PluginSetNoContent {
	return &PluginSetNoContent{}
}

/* PluginSetNoContent describes a response with status code 204, with default header values.

No error
*/
type PluginSetNoContent struct {
}

func (o *PluginSetNoContent) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/set][%d] pluginSetNoContent ", 204)
}

func (o *PluginSetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPluginSetNotFound creates a PluginSetNotFound with default headers values
func NewPluginSetNotFound() *PluginSetNotFound {
	return &PluginSetNotFound{}
}

/* PluginSetNotFound describes a response with status code 404, with default header values.

Plugin not installed
*/
type PluginSetNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PluginSetNotFound) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/set][%d] pluginSetNotFound  %+v", 404, o.Payload)
}
func (o *PluginSetNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PluginSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginSetInternalServerError creates a PluginSetInternalServerError with default headers values
func NewPluginSetInternalServerError() *PluginSetInternalServerError {
	return &PluginSetInternalServerError{}
}

/* PluginSetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type PluginSetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PluginSetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/set][%d] pluginSetInternalServerError  %+v", 500, o.Payload)
}
func (o *PluginSetInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PluginSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}