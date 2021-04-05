// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-swagger/dockerctl/models"
)

// PluginPullReader is a Reader for the PluginPull structure.
type PluginPullReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginPullReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginPullNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPluginPullInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginPullNoContent creates a PluginPullNoContent with default headers values
func NewPluginPullNoContent() *PluginPullNoContent {
	return &PluginPullNoContent{}
}

/* PluginPullNoContent describes a response with status code 204, with default header values.

no error
*/
type PluginPullNoContent struct {
}

func (o *PluginPullNoContent) Error() string {
	return fmt.Sprintf("[POST /plugins/pull][%d] pluginPullNoContent ", 204)
}

func (o *PluginPullNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPluginPullInternalServerError creates a PluginPullInternalServerError with default headers values
func NewPluginPullInternalServerError() *PluginPullInternalServerError {
	return &PluginPullInternalServerError{}
}

/* PluginPullInternalServerError describes a response with status code 500, with default header values.

server error
*/
type PluginPullInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PluginPullInternalServerError) Error() string {
	return fmt.Sprintf("[POST /plugins/pull][%d] pluginPullInternalServerError  %+v", 500, o.Payload)
}
func (o *PluginPullInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PluginPullInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PluginPullParamsBodyItems0 Describes a permission accepted by the user upon installing the plugin.
swagger:model PluginPullParamsBodyItems0
*/
type PluginPullParamsBodyItems0 struct {

	// description
	Description string `json:"Description,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// value
	Value []string `json:"Value"`
}

// Validate validates this plugin pull params body items0
func (o *PluginPullParamsBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plugin pull params body items0 based on context it is used
func (o *PluginPullParamsBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PluginPullParamsBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PluginPullParamsBodyItems0) UnmarshalBinary(b []byte) error {
	var res PluginPullParamsBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
