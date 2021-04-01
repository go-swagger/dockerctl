// Code generated by go-swagger; DO NOT EDIT.

package network

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

	"github.com/go-openapi/dockerctl/models"
)

// NetworkConnectReader is a Reader for the NetworkConnect structure.
type NetworkConnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkConnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkConnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNetworkConnectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNetworkConnectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkConnectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkConnectOK creates a NetworkConnectOK with default headers values
func NewNetworkConnectOK() *NetworkConnectOK {
	return &NetworkConnectOK{}
}

/* NetworkConnectOK describes a response with status code 200, with default header values.

No error
*/
type NetworkConnectOK struct {
}

func (o *NetworkConnectOK) Error() string {
	return fmt.Sprintf("[POST /networks/{id}/connect][%d] networkConnectOK ", 200)
}

func (o *NetworkConnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkConnectForbidden creates a NetworkConnectForbidden with default headers values
func NewNetworkConnectForbidden() *NetworkConnectForbidden {
	return &NetworkConnectForbidden{}
}

/* NetworkConnectForbidden describes a response with status code 403, with default header values.

Operation not supported for swarm scoped networks
*/
type NetworkConnectForbidden struct {
	Payload *models.ErrorResponse
}

func (o *NetworkConnectForbidden) Error() string {
	return fmt.Sprintf("[POST /networks/{id}/connect][%d] networkConnectForbidden  %+v", 403, o.Payload)
}
func (o *NetworkConnectForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkConnectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkConnectNotFound creates a NetworkConnectNotFound with default headers values
func NewNetworkConnectNotFound() *NetworkConnectNotFound {
	return &NetworkConnectNotFound{}
}

/* NetworkConnectNotFound describes a response with status code 404, with default header values.

Network or container not found
*/
type NetworkConnectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *NetworkConnectNotFound) Error() string {
	return fmt.Sprintf("[POST /networks/{id}/connect][%d] networkConnectNotFound  %+v", 404, o.Payload)
}
func (o *NetworkConnectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkConnectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkConnectInternalServerError creates a NetworkConnectInternalServerError with default headers values
func NewNetworkConnectInternalServerError() *NetworkConnectInternalServerError {
	return &NetworkConnectInternalServerError{}
}

/* NetworkConnectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type NetworkConnectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *NetworkConnectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /networks/{id}/connect][%d] networkConnectInternalServerError  %+v", 500, o.Payload)
}
func (o *NetworkConnectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkConnectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*NetworkConnectBody network connect body
// Example: {"Container":"3613f73ba0e4","EndpointConfig":{"IPAMConfig":{"IPv4Address":"172.24.56.89","IPv6Address":"2001:db8::5689"}}}
swagger:model NetworkConnectBody
*/
type NetworkConnectBody struct {

	// The ID or name of the container to connect to the network.
	Container string `json:"Container,omitempty"`

	// endpoint config
	EndpointConfig *models.EndpointSettings `json:"EndpointConfig,omitempty"`
}

// Validate validates this network connect body
func (o *NetworkConnectBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEndpointConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkConnectBody) validateEndpointConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.EndpointConfig) { // not required
		return nil
	}

	if o.EndpointConfig != nil {
		if err := o.EndpointConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container" + "." + "EndpointConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network connect body based on the context it is used
func (o *NetworkConnectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEndpointConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkConnectBody) contextValidateEndpointConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.EndpointConfig != nil {
		if err := o.EndpointConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container" + "." + "EndpointConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetworkConnectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkConnectBody) UnmarshalBinary(b []byte) error {
	var res NetworkConnectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
