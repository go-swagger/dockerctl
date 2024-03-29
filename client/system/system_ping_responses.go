// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-swagger/dockerctl/models"
)

// SystemPingReader is a Reader for the SystemPing structure.
type SystemPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSystemPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemPingOK creates a SystemPingOK with default headers values
func NewSystemPingOK() *SystemPingOK {
	var (
		// initialize headers with default values
		cacheControlDefault = string("no-cache, no-store, must-revalidate")

		pragmaDefault = string("no-cache")
	)

	return &SystemPingOK{

		CacheControl: cacheControlDefault,
		Pragma:       pragmaDefault,
	}
}

/* SystemPingOK describes a response with status code 200, with default header values.

no error
*/
type SystemPingOK struct {

	/* Max API Version the server supports
	 */
	APIVersion string

	/* Default version of docker image builder
	 */
	BuildKitVersion string
	CacheControl    string

	/* If the server is running with experimental mode enabled
	 */
	DockerExperimental bool
	Pragma             string

	Payload string
}

// IsSuccess returns true when this system ping o k response has a 2xx status code
func (o *SystemPingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system ping o k response has a 3xx status code
func (o *SystemPingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system ping o k response has a 4xx status code
func (o *SystemPingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this system ping o k response has a 5xx status code
func (o *SystemPingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this system ping o k response a status code equal to that given
func (o *SystemPingOK) IsCode(code int) bool {
	return code == 200
}

func (o *SystemPingOK) Error() string {
	return fmt.Sprintf("[GET /_ping][%d] systemPingOK  %+v", 200, o.Payload)
}

func (o *SystemPingOK) String() string {
	return fmt.Sprintf("[GET /_ping][%d] systemPingOK  %+v", 200, o.Payload)
}

func (o *SystemPingOK) GetPayload() string {
	return o.Payload
}

func (o *SystemPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header API-Version
	hdrAPIVersion := response.GetHeader("API-Version")

	if hdrAPIVersion != "" {
		o.APIVersion = hdrAPIVersion
	}

	// hydrates response header BuildKit-Version
	hdrBuildKitVersion := response.GetHeader("BuildKit-Version")

	if hdrBuildKitVersion != "" {
		o.BuildKitVersion = hdrBuildKitVersion
	}

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Docker-Experimental
	hdrDockerExperimental := response.GetHeader("Docker-Experimental")

	if hdrDockerExperimental != "" {
		valdockerExperimental, err := swag.ConvertBool(hdrDockerExperimental)
		if err != nil {
			return errors.InvalidType("Docker-Experimental", "header", "bool", hdrDockerExperimental)
		}
		o.DockerExperimental = valdockerExperimental
	}

	// hydrates response header Pragma
	hdrPragma := response.GetHeader("Pragma")

	if hdrPragma != "" {
		o.Pragma = hdrPragma
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemPingInternalServerError creates a SystemPingInternalServerError with default headers values
func NewSystemPingInternalServerError() *SystemPingInternalServerError {
	var (
		// initialize headers with default values
		cacheControlDefault = string("no-cache, no-store, must-revalidate")

		pragmaDefault = string("no-cache")
	)

	return &SystemPingInternalServerError{

		CacheControl: cacheControlDefault,
		Pragma:       pragmaDefault,
	}
}

/* SystemPingInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SystemPingInternalServerError struct {
	CacheControl string
	Pragma       string

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this system ping internal server error response has a 2xx status code
func (o *SystemPingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this system ping internal server error response has a 3xx status code
func (o *SystemPingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system ping internal server error response has a 4xx status code
func (o *SystemPingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this system ping internal server error response has a 5xx status code
func (o *SystemPingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this system ping internal server error response a status code equal to that given
func (o *SystemPingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SystemPingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /_ping][%d] systemPingInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemPingInternalServerError) String() string {
	return fmt.Sprintf("[GET /_ping][%d] systemPingInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemPingInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SystemPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Pragma
	hdrPragma := response.GetHeader("Pragma")

	if hdrPragma != "" {
		o.Pragma = hdrPragma
	}

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
