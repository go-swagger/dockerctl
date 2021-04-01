// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/dockerctl/models"
)

// NewSystemAuthParams creates a new SystemAuthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemAuthParams() *SystemAuthParams {
	return &SystemAuthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemAuthParamsWithTimeout creates a new SystemAuthParams object
// with the ability to set a timeout on a request.
func NewSystemAuthParamsWithTimeout(timeout time.Duration) *SystemAuthParams {
	return &SystemAuthParams{
		timeout: timeout,
	}
}

// NewSystemAuthParamsWithContext creates a new SystemAuthParams object
// with the ability to set a context for a request.
func NewSystemAuthParamsWithContext(ctx context.Context) *SystemAuthParams {
	return &SystemAuthParams{
		Context: ctx,
	}
}

// NewSystemAuthParamsWithHTTPClient creates a new SystemAuthParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemAuthParamsWithHTTPClient(client *http.Client) *SystemAuthParams {
	return &SystemAuthParams{
		HTTPClient: client,
	}
}

/* SystemAuthParams contains all the parameters to send to the API endpoint
   for the system auth operation.

   Typically these are written to a http.Request.
*/
type SystemAuthParams struct {

	/* AuthConfig.

	   Authentication to check
	*/
	AuthConfig *models.AuthConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemAuthParams) WithDefaults() *SystemAuthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemAuthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system auth params
func (o *SystemAuthParams) WithTimeout(timeout time.Duration) *SystemAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system auth params
func (o *SystemAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system auth params
func (o *SystemAuthParams) WithContext(ctx context.Context) *SystemAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system auth params
func (o *SystemAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system auth params
func (o *SystemAuthParams) WithHTTPClient(client *http.Client) *SystemAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system auth params
func (o *SystemAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthConfig adds the authConfig to the system auth params
func (o *SystemAuthParams) WithAuthConfig(authConfig *models.AuthConfig) *SystemAuthParams {
	o.SetAuthConfig(authConfig)
	return o
}

// SetAuthConfig adds the authConfig to the system auth params
func (o *SystemAuthParams) SetAuthConfig(authConfig *models.AuthConfig) {
	o.AuthConfig = authConfig
}

// WriteToRequest writes these params to a swagger request
func (o *SystemAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AuthConfig != nil {
		if err := r.SetBodyParam(o.AuthConfig); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
