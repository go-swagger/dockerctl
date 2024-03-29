// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// PluginInspectReader is a Reader for the PluginInspect structure.
type PluginInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPluginInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPluginInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginInspectOK creates a PluginInspectOK with default headers values
func NewPluginInspectOK() *PluginInspectOK {
	return &PluginInspectOK{}
}

/* PluginInspectOK describes a response with status code 200, with default header values.

no error
*/
type PluginInspectOK struct {
	Payload *models.Plugin
}

// IsSuccess returns true when this plugin inspect o k response has a 2xx status code
func (o *PluginInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugin inspect o k response has a 3xx status code
func (o *PluginInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugin inspect o k response has a 4xx status code
func (o *PluginInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugin inspect o k response has a 5xx status code
func (o *PluginInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugin inspect o k response a status code equal to that given
func (o *PluginInspectOK) IsCode(code int) bool {
	return code == 200
}

func (o *PluginInspectOK) Error() string {
	return fmt.Sprintf("[GET /plugins/{name}/json][%d] pluginInspectOK  %+v", 200, o.Payload)
}

func (o *PluginInspectOK) String() string {
	return fmt.Sprintf("[GET /plugins/{name}/json][%d] pluginInspectOK  %+v", 200, o.Payload)
}

func (o *PluginInspectOK) GetPayload() *models.Plugin {
	return o.Payload
}

func (o *PluginInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Plugin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginInspectNotFound creates a PluginInspectNotFound with default headers values
func NewPluginInspectNotFound() *PluginInspectNotFound {
	return &PluginInspectNotFound{}
}

/* PluginInspectNotFound describes a response with status code 404, with default header values.

plugin is not installed
*/
type PluginInspectNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this plugin inspect not found response has a 2xx status code
func (o *PluginInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this plugin inspect not found response has a 3xx status code
func (o *PluginInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugin inspect not found response has a 4xx status code
func (o *PluginInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this plugin inspect not found response has a 5xx status code
func (o *PluginInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this plugin inspect not found response a status code equal to that given
func (o *PluginInspectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PluginInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /plugins/{name}/json][%d] pluginInspectNotFound  %+v", 404, o.Payload)
}

func (o *PluginInspectNotFound) String() string {
	return fmt.Sprintf("[GET /plugins/{name}/json][%d] pluginInspectNotFound  %+v", 404, o.Payload)
}

func (o *PluginInspectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PluginInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginInspectInternalServerError creates a PluginInspectInternalServerError with default headers values
func NewPluginInspectInternalServerError() *PluginInspectInternalServerError {
	return &PluginInspectInternalServerError{}
}

/* PluginInspectInternalServerError describes a response with status code 500, with default header values.

server error
*/
type PluginInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this plugin inspect internal server error response has a 2xx status code
func (o *PluginInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this plugin inspect internal server error response has a 3xx status code
func (o *PluginInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugin inspect internal server error response has a 4xx status code
func (o *PluginInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugin inspect internal server error response has a 5xx status code
func (o *PluginInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this plugin inspect internal server error response a status code equal to that given
func (o *PluginInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PluginInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /plugins/{name}/json][%d] pluginInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *PluginInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /plugins/{name}/json][%d] pluginInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *PluginInspectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PluginInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
