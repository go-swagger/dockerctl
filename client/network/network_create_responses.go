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
	"github.com/go-openapi/validate"

	"github.com/go-swagger/dockerctl/models"
)

// NetworkCreateReader is a Reader for the NetworkCreate structure.
type NetworkCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNetworkCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNetworkCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNetworkCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkCreateCreated creates a NetworkCreateCreated with default headers values
func NewNetworkCreateCreated() *NetworkCreateCreated {
	return &NetworkCreateCreated{}
}

/* NetworkCreateCreated describes a response with status code 201, with default header values.

No error
*/
type NetworkCreateCreated struct {
	Payload *NetworkCreateCreatedBody
}

// IsSuccess returns true when this network create created response has a 2xx status code
func (o *NetworkCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network create created response has a 3xx status code
func (o *NetworkCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create created response has a 4xx status code
func (o *NetworkCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this network create created response has a 5xx status code
func (o *NetworkCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this network create created response a status code equal to that given
func (o *NetworkCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *NetworkCreateCreated) Error() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateCreated  %+v", 201, o.Payload)
}

func (o *NetworkCreateCreated) String() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateCreated  %+v", 201, o.Payload)
}

func (o *NetworkCreateCreated) GetPayload() *NetworkCreateCreatedBody {
	return o.Payload
}

func (o *NetworkCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkCreateCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkCreateForbidden creates a NetworkCreateForbidden with default headers values
func NewNetworkCreateForbidden() *NetworkCreateForbidden {
	return &NetworkCreateForbidden{}
}

/* NetworkCreateForbidden describes a response with status code 403, with default header values.

operation not supported for pre-defined networks
*/
type NetworkCreateForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network create forbidden response has a 2xx status code
func (o *NetworkCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network create forbidden response has a 3xx status code
func (o *NetworkCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create forbidden response has a 4xx status code
func (o *NetworkCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this network create forbidden response has a 5xx status code
func (o *NetworkCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this network create forbidden response a status code equal to that given
func (o *NetworkCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *NetworkCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateForbidden  %+v", 403, o.Payload)
}

func (o *NetworkCreateForbidden) String() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateForbidden  %+v", 403, o.Payload)
}

func (o *NetworkCreateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkCreateNotFound creates a NetworkCreateNotFound with default headers values
func NewNetworkCreateNotFound() *NetworkCreateNotFound {
	return &NetworkCreateNotFound{}
}

/* NetworkCreateNotFound describes a response with status code 404, with default header values.

plugin not found
*/
type NetworkCreateNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network create not found response has a 2xx status code
func (o *NetworkCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network create not found response has a 3xx status code
func (o *NetworkCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create not found response has a 4xx status code
func (o *NetworkCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this network create not found response has a 5xx status code
func (o *NetworkCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this network create not found response a status code equal to that given
func (o *NetworkCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NetworkCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateNotFound  %+v", 404, o.Payload)
}

func (o *NetworkCreateNotFound) String() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateNotFound  %+v", 404, o.Payload)
}

func (o *NetworkCreateNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkCreateInternalServerError creates a NetworkCreateInternalServerError with default headers values
func NewNetworkCreateInternalServerError() *NetworkCreateInternalServerError {
	return &NetworkCreateInternalServerError{}
}

/* NetworkCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type NetworkCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network create internal server error response has a 2xx status code
func (o *NetworkCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network create internal server error response has a 3xx status code
func (o *NetworkCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create internal server error response has a 4xx status code
func (o *NetworkCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this network create internal server error response has a 5xx status code
func (o *NetworkCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this network create internal server error response a status code equal to that given
func (o *NetworkCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NetworkCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkCreateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*NetworkCreateBody network create body
// Example: {"Attachable":false,"CheckDuplicate":false,"Driver":"bridge","EnableIPv6":true,"IPAM":{"Config":[{"Gateway":"172.20.10.11","IPRange":"172.20.10.0/24","Subnet":"172.20.0.0/16"},{"Gateway":"2001:db8:abcd::1011","Subnet":"2001:db8:abcd::/64"}],"Driver":"default","Options":{"foo":"bar"}},"Ingress":false,"Internal":true,"Labels":{"com.example.some-label":"some-value","com.example.some-other-label":"some-other-value"},"Name":"isolated_nw","Options":{"com.docker.network.bridge.default_bridge":"true","com.docker.network.bridge.enable_icc":"true","com.docker.network.bridge.enable_ip_masquerade":"true","com.docker.network.bridge.host_binding_ipv4":"0.0.0.0","com.docker.network.bridge.name":"docker0","com.docker.network.driver.mtu":"1500"}}
swagger:model NetworkCreateBody
*/
type NetworkCreateBody struct {

	// Globally scoped network is manually attachable by regular containers from workers in swarm mode.
	Attachable bool `json:"Attachable,omitempty"`

	// Check for networks with duplicate names. Since Network is primarily keyed based on a random ID and not on the name, and network name is strictly a user-friendly alias to the network which is uniquely identified using ID, there is no guaranteed way to check for duplicates. CheckDuplicate is there to provide a best effort checking of any networks which has the same name but it is not guaranteed to catch all name collisions.
	CheckDuplicate bool `json:"CheckDuplicate,omitempty"`

	// Name of the network driver plugin to use.
	Driver *string `json:"Driver,omitempty"`

	// Enable IPv6 on the network.
	EnableIPV6 bool `json:"EnableIPv6,omitempty"`

	// Optional custom IP scheme for the network.
	IPAM *models.IPAM `json:"IPAM,omitempty"`

	// Ingress network is the network which provides the routing-mesh in swarm mode.
	Ingress bool `json:"Ingress,omitempty"`

	// Restrict external access to the network.
	Internal bool `json:"Internal,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// The network's name.
	// Required: true
	Name *string `json:"Name"`

	// Network specific options to be used by the drivers.
	Options map[string]string `json:"Options,omitempty"`
}

// Validate validates this network create body
func (o *NetworkCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPAM(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkCreateBody) validateIPAM(formats strfmt.Registry) error {
	if swag.IsZero(o.IPAM) { // not required
		return nil
	}

	if o.IPAM != nil {
		if err := o.IPAM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkConfig" + "." + "IPAM")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkConfig" + "." + "IPAM")
			}
			return err
		}
	}

	return nil
}

func (o *NetworkCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("networkConfig"+"."+"Name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this network create body based on the context it is used
func (o *NetworkCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPAM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkCreateBody) contextValidateIPAM(ctx context.Context, formats strfmt.Registry) error {

	if o.IPAM != nil {
		if err := o.IPAM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkConfig" + "." + "IPAM")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkConfig" + "." + "IPAM")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetworkCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkCreateBody) UnmarshalBinary(b []byte) error {
	var res NetworkCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*NetworkCreateCreatedBody NetworkCreateResponse
// Example: {"Id":"22be93d5babb089c5aab8dbc369042fad48ff791584ca2da2100db837a1c7c30","Warning":""}
swagger:model NetworkCreateCreatedBody
*/
type NetworkCreateCreatedBody struct {

	// The ID of the created network.
	ID string `json:"Id,omitempty"`

	// warning
	Warning string `json:"Warning,omitempty"`
}

// Validate validates this network create created body
func (o *NetworkCreateCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network create created body based on context it is used
func (o *NetworkCreateCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkCreateCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkCreateCreatedBody) UnmarshalBinary(b []byte) error {
	var res NetworkCreateCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
