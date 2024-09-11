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

// NewSnapmirrorRelationshipTransferGetParams creates a new SnapmirrorRelationshipTransferGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapmirrorRelationshipTransferGetParams() *SnapmirrorRelationshipTransferGetParams {
	return &SnapmirrorRelationshipTransferGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapmirrorRelationshipTransferGetParamsWithTimeout creates a new SnapmirrorRelationshipTransferGetParams object
// with the ability to set a timeout on a request.
func NewSnapmirrorRelationshipTransferGetParamsWithTimeout(timeout time.Duration) *SnapmirrorRelationshipTransferGetParams {
	return &SnapmirrorRelationshipTransferGetParams{
		timeout: timeout,
	}
}

// NewSnapmirrorRelationshipTransferGetParamsWithContext creates a new SnapmirrorRelationshipTransferGetParams object
// with the ability to set a context for a request.
func NewSnapmirrorRelationshipTransferGetParamsWithContext(ctx context.Context) *SnapmirrorRelationshipTransferGetParams {
	return &SnapmirrorRelationshipTransferGetParams{
		Context: ctx,
	}
}

// NewSnapmirrorRelationshipTransferGetParamsWithHTTPClient creates a new SnapmirrorRelationshipTransferGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapmirrorRelationshipTransferGetParamsWithHTTPClient(client *http.Client) *SnapmirrorRelationshipTransferGetParams {
	return &SnapmirrorRelationshipTransferGetParams{
		HTTPClient: client,
	}
}

/*
SnapmirrorRelationshipTransferGetParams contains all the parameters to send to the API endpoint

	for the snapmirror relationship transfer get operation.

	Typically these are written to a http.Request.
*/
type SnapmirrorRelationshipTransferGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* RelationshipUUID.

	   SnapMirror relationship UUID
	*/
	RelationshipUUID string

	/* UUID.

	   SnapMirror transfer UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapmirror relationship transfer get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorRelationshipTransferGetParams) WithDefaults() *SnapmirrorRelationshipTransferGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapmirror relationship transfer get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorRelationshipTransferGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) WithTimeout(timeout time.Duration) *SnapmirrorRelationshipTransferGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) WithContext(ctx context.Context) *SnapmirrorRelationshipTransferGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) WithHTTPClient(client *http.Client) *SnapmirrorRelationshipTransferGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) WithFields(fields []string) *SnapmirrorRelationshipTransferGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithRelationshipUUID adds the relationshipUUID to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) WithRelationshipUUID(relationshipUUID string) *SnapmirrorRelationshipTransferGetParams {
	o.SetRelationshipUUID(relationshipUUID)
	return o
}

// SetRelationshipUUID adds the relationshipUuid to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) SetRelationshipUUID(relationshipUUID string) {
	o.RelationshipUUID = relationshipUUID
}

// WithUUID adds the uuid to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) WithUUID(uuid string) *SnapmirrorRelationshipTransferGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the snapmirror relationship transfer get params
func (o *SnapmirrorRelationshipTransferGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapmirrorRelationshipTransferGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param relationship.uuid
	if err := r.SetPathParam("relationship.uuid", o.RelationshipUUID); err != nil {
		return err
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

// bindParamSnapmirrorRelationshipTransferGet binds the parameter fields
func (o *SnapmirrorRelationshipTransferGetParams) bindParamFields(formats strfmt.Registry) []string {
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
