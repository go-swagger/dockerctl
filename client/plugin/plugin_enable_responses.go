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

// PluginEnableReader is a Reader for the PluginEnable structure.
type PluginEnableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginEnableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginEnableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPluginEnableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPluginEnableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginEnableOK creates a PluginEnableOK with default headers values
func NewPluginEnableOK() *PluginEnableOK {
	return &PluginEnableOK{}
}

/* PluginEnableOK describes a response with status code 200, with default header values.

no error
*/
type PluginEnableOK struct {
}

func (o *PluginEnableOK) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/enable][%d] pluginEnableOK ", 200)
}

func (o *PluginEnableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPluginEnableNotFound creates a PluginEnableNotFound with default headers values
func NewPluginEnableNotFound() *PluginEnableNotFound {
	return &PluginEnableNotFound{}
}

/* PluginEnableNotFound describes a response with status code 404, with default header values.

plugin is not installed
*/
type PluginEnableNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PluginEnableNotFound) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/enable][%d] pluginEnableNotFound  %+v", 404, o.Payload)
}
func (o *PluginEnableNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PluginEnableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginEnableInternalServerError creates a PluginEnableInternalServerError with default headers values
func NewPluginEnableInternalServerError() *PluginEnableInternalServerError {
	return &PluginEnableInternalServerError{}
}

/* PluginEnableInternalServerError describes a response with status code 500, with default header values.

server error
*/
type PluginEnableInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PluginEnableInternalServerError) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/enable][%d] pluginEnableInternalServerError  %+v", 500, o.Payload)
}
func (o *PluginEnableInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PluginEnableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
