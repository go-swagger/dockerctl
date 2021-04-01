// Code generated by go-swagger; DO NOT EDIT.

package service

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
	"github.com/go-openapi/swag"
)

// NewServiceInspectParams creates a new ServiceInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceInspectParams() *ServiceInspectParams {
	return &ServiceInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceInspectParamsWithTimeout creates a new ServiceInspectParams object
// with the ability to set a timeout on a request.
func NewServiceInspectParamsWithTimeout(timeout time.Duration) *ServiceInspectParams {
	return &ServiceInspectParams{
		timeout: timeout,
	}
}

// NewServiceInspectParamsWithContext creates a new ServiceInspectParams object
// with the ability to set a context for a request.
func NewServiceInspectParamsWithContext(ctx context.Context) *ServiceInspectParams {
	return &ServiceInspectParams{
		Context: ctx,
	}
}

// NewServiceInspectParamsWithHTTPClient creates a new ServiceInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceInspectParamsWithHTTPClient(client *http.Client) *ServiceInspectParams {
	return &ServiceInspectParams{
		HTTPClient: client,
	}
}

/* ServiceInspectParams contains all the parameters to send to the API endpoint
   for the service inspect operation.

   Typically these are written to a http.Request.
*/
type ServiceInspectParams struct {

	/* ID.

	   ID or name of service.
	*/
	ID string

	/* InsertDefaults.

	   Fill empty fields with default values.
	*/
	InsertDefaults *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceInspectParams) WithDefaults() *ServiceInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceInspectParams) SetDefaults() {
	var (
		insertDefaultsDefault = bool(false)
	)

	val := ServiceInspectParams{
		InsertDefaults: &insertDefaultsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the service inspect params
func (o *ServiceInspectParams) WithTimeout(timeout time.Duration) *ServiceInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service inspect params
func (o *ServiceInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service inspect params
func (o *ServiceInspectParams) WithContext(ctx context.Context) *ServiceInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service inspect params
func (o *ServiceInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service inspect params
func (o *ServiceInspectParams) WithHTTPClient(client *http.Client) *ServiceInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service inspect params
func (o *ServiceInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the service inspect params
func (o *ServiceInspectParams) WithID(id string) *ServiceInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the service inspect params
func (o *ServiceInspectParams) SetID(id string) {
	o.ID = id
}

// WithInsertDefaults adds the insertDefaults to the service inspect params
func (o *ServiceInspectParams) WithInsertDefaults(insertDefaults *bool) *ServiceInspectParams {
	o.SetInsertDefaults(insertDefaults)
	return o
}

// SetInsertDefaults adds the insertDefaults to the service inspect params
func (o *ServiceInspectParams) SetInsertDefaults(insertDefaults *bool) {
	o.InsertDefaults = insertDefaults
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.InsertDefaults != nil {

		// query param insertDefaults
		var qrInsertDefaults bool

		if o.InsertDefaults != nil {
			qrInsertDefaults = *o.InsertDefaults
		}
		qInsertDefaults := swag.FormatBool(qrInsertDefaults)
		if qInsertDefaults != "" {

			if err := r.SetQueryParam("insertDefaults", qInsertDefaults); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
