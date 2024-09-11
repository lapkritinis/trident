// Code generated by go-swagger; DO NOT EDIT.

package snapmirror

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

// NewSnapmirrorPolicyDeleteParams creates a new SnapmirrorPolicyDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapmirrorPolicyDeleteParams() *SnapmirrorPolicyDeleteParams {
	return &SnapmirrorPolicyDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapmirrorPolicyDeleteParamsWithTimeout creates a new SnapmirrorPolicyDeleteParams object
// with the ability to set a timeout on a request.
func NewSnapmirrorPolicyDeleteParamsWithTimeout(timeout time.Duration) *SnapmirrorPolicyDeleteParams {
	return &SnapmirrorPolicyDeleteParams{
		timeout: timeout,
	}
}

// NewSnapmirrorPolicyDeleteParamsWithContext creates a new SnapmirrorPolicyDeleteParams object
// with the ability to set a context for a request.
func NewSnapmirrorPolicyDeleteParamsWithContext(ctx context.Context) *SnapmirrorPolicyDeleteParams {
	return &SnapmirrorPolicyDeleteParams{
		Context: ctx,
	}
}

// NewSnapmirrorPolicyDeleteParamsWithHTTPClient creates a new SnapmirrorPolicyDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapmirrorPolicyDeleteParamsWithHTTPClient(client *http.Client) *SnapmirrorPolicyDeleteParams {
	return &SnapmirrorPolicyDeleteParams{
		HTTPClient: client,
	}
}

/*
SnapmirrorPolicyDeleteParams contains all the parameters to send to the API endpoint

	for the snapmirror policy delete operation.

	Typically these are written to a http.Request.
*/
type SnapmirrorPolicyDeleteParams struct {

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   SnapMirror policy UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapmirror policy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorPolicyDeleteParams) WithDefaults() *SnapmirrorPolicyDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapmirror policy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorPolicyDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := SnapmirrorPolicyDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) WithTimeout(timeout time.Duration) *SnapmirrorPolicyDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) WithContext(ctx context.Context) *SnapmirrorPolicyDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) WithHTTPClient(client *http.Client) *SnapmirrorPolicyDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnTimeout adds the returnTimeout to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) WithReturnTimeout(returnTimeout *int64) *SnapmirrorPolicyDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) WithUUID(uuid string) *SnapmirrorPolicyDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the snapmirror policy delete params
func (o *SnapmirrorPolicyDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapmirrorPolicyDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
