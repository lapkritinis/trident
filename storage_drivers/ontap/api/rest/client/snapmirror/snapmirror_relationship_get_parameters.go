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

// NewSnapmirrorRelationshipGetParams creates a new SnapmirrorRelationshipGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapmirrorRelationshipGetParams() *SnapmirrorRelationshipGetParams {
	return &SnapmirrorRelationshipGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapmirrorRelationshipGetParamsWithTimeout creates a new SnapmirrorRelationshipGetParams object
// with the ability to set a timeout on a request.
func NewSnapmirrorRelationshipGetParamsWithTimeout(timeout time.Duration) *SnapmirrorRelationshipGetParams {
	return &SnapmirrorRelationshipGetParams{
		timeout: timeout,
	}
}

// NewSnapmirrorRelationshipGetParamsWithContext creates a new SnapmirrorRelationshipGetParams object
// with the ability to set a context for a request.
func NewSnapmirrorRelationshipGetParamsWithContext(ctx context.Context) *SnapmirrorRelationshipGetParams {
	return &SnapmirrorRelationshipGetParams{
		Context: ctx,
	}
}

// NewSnapmirrorRelationshipGetParamsWithHTTPClient creates a new SnapmirrorRelationshipGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapmirrorRelationshipGetParamsWithHTTPClient(client *http.Client) *SnapmirrorRelationshipGetParams {
	return &SnapmirrorRelationshipGetParams{
		HTTPClient: client,
	}
}

/*
SnapmirrorRelationshipGetParams contains all the parameters to send to the API endpoint

	for the snapmirror relationship get operation.

	Typically these are written to a http.Request.
*/
type SnapmirrorRelationshipGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* ListDestinationsOnly.

	   Set to true to show relationships from the source only.
	*/
	ListDestinationsOnly *bool

	/* UUID.

	   SnapMirror relationship UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapmirror relationship get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorRelationshipGetParams) WithDefaults() *SnapmirrorRelationshipGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapmirror relationship get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorRelationshipGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) WithTimeout(timeout time.Duration) *SnapmirrorRelationshipGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) WithContext(ctx context.Context) *SnapmirrorRelationshipGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) WithHTTPClient(client *http.Client) *SnapmirrorRelationshipGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) WithFields(fields []string) *SnapmirrorRelationshipGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithListDestinationsOnly adds the listDestinationsOnly to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) WithListDestinationsOnly(listDestinationsOnly *bool) *SnapmirrorRelationshipGetParams {
	o.SetListDestinationsOnly(listDestinationsOnly)
	return o
}

// SetListDestinationsOnly adds the listDestinationsOnly to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) SetListDestinationsOnly(listDestinationsOnly *bool) {
	o.ListDestinationsOnly = listDestinationsOnly
}

// WithUUID adds the uuid to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) WithUUID(uuid string) *SnapmirrorRelationshipGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the snapmirror relationship get params
func (o *SnapmirrorRelationshipGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapmirrorRelationshipGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.ListDestinationsOnly != nil {

		// query param list_destinations_only
		var qrListDestinationsOnly bool

		if o.ListDestinationsOnly != nil {
			qrListDestinationsOnly = *o.ListDestinationsOnly
		}
		qListDestinationsOnly := swag.FormatBool(qrListDestinationsOnly)
		if qListDestinationsOnly != "" {

			if err := r.SetQueryParam("list_destinations_only", qListDestinationsOnly); err != nil {
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

// bindParamSnapmirrorRelationshipGet binds the parameter fields
func (o *SnapmirrorRelationshipGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
