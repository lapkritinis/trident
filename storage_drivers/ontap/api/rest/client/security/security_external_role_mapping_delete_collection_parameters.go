// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityExternalRoleMappingDeleteCollectionParams creates a new SecurityExternalRoleMappingDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityExternalRoleMappingDeleteCollectionParams() *SecurityExternalRoleMappingDeleteCollectionParams {
	return &SecurityExternalRoleMappingDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityExternalRoleMappingDeleteCollectionParamsWithTimeout creates a new SecurityExternalRoleMappingDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewSecurityExternalRoleMappingDeleteCollectionParamsWithTimeout(timeout time.Duration) *SecurityExternalRoleMappingDeleteCollectionParams {
	return &SecurityExternalRoleMappingDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewSecurityExternalRoleMappingDeleteCollectionParamsWithContext creates a new SecurityExternalRoleMappingDeleteCollectionParams object
// with the ability to set a context for a request.
func NewSecurityExternalRoleMappingDeleteCollectionParamsWithContext(ctx context.Context) *SecurityExternalRoleMappingDeleteCollectionParams {
	return &SecurityExternalRoleMappingDeleteCollectionParams{
		Context: ctx,
	}
}

// NewSecurityExternalRoleMappingDeleteCollectionParamsWithHTTPClient creates a new SecurityExternalRoleMappingDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityExternalRoleMappingDeleteCollectionParamsWithHTTPClient(client *http.Client) *SecurityExternalRoleMappingDeleteCollectionParams {
	return &SecurityExternalRoleMappingDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
SecurityExternalRoleMappingDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the security external role mapping delete collection operation.

	Typically these are written to a http.Request.
*/
type SecurityExternalRoleMappingDeleteCollectionParams struct {

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* ExternalRole.

	   Filter by external_role
	*/
	ExternalRole *string

	/* Info.

	   Info specification
	*/
	Info SecurityExternalRoleMappingDeleteCollectionBody

	/* OntapRoleName.

	   Filter by ontap_role.name
	*/
	OntapRoleName *string

	/* Provider.

	   Filter by provider
	*/
	Provider *string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* Timestamp.

	   Filter by timestamp
	*/
	Timestamp *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security external role mapping delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithDefaults() *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security external role mapping delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SecurityExternalRoleMappingDeleteCollectionParams{
		ContinueOnFailure: &continueOnFailureDefault,
		ReturnRecords:     &returnRecordsDefault,
		ReturnTimeout:     &returnTimeoutDefault,
		SerialRecords:     &serialRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithTimeout(timeout time.Duration) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithContext(ctx context.Context) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithHTTPClient(client *http.Client) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithComment(comment *string) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContinueOnFailure adds the continueOnFailure to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithExternalRole adds the externalRole to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithExternalRole(externalRole *string) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetExternalRole(externalRole)
	return o
}

// SetExternalRole adds the externalRole to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetExternalRole(externalRole *string) {
	o.ExternalRole = externalRole
}

// WithInfo adds the info to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithInfo(info SecurityExternalRoleMappingDeleteCollectionBody) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetInfo(info SecurityExternalRoleMappingDeleteCollectionBody) {
	o.Info = info
}

// WithOntapRoleName adds the ontapRoleName to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithOntapRoleName(ontapRoleName *string) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetOntapRoleName(ontapRoleName)
	return o
}

// SetOntapRoleName adds the ontapRoleName to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetOntapRoleName(ontapRoleName *string) {
	o.OntapRoleName = ontapRoleName
}

// WithProvider adds the provider to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithProvider(provider *string) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithReturnRecords adds the returnRecords to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithTimestamp adds the timestamp to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WithTimestamp(timestamp *string) *SecurityExternalRoleMappingDeleteCollectionParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the security external role mapping delete collection params
func (o *SecurityExternalRoleMappingDeleteCollectionParams) SetTimestamp(timestamp *string) {
	o.Timestamp = timestamp
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityExternalRoleMappingDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.ContinueOnFailure != nil {

		// query param continue_on_failure
		var qrContinueOnFailure bool

		if o.ContinueOnFailure != nil {
			qrContinueOnFailure = *o.ContinueOnFailure
		}
		qContinueOnFailure := swag.FormatBool(qrContinueOnFailure)
		if qContinueOnFailure != "" {

			if err := r.SetQueryParam("continue_on_failure", qContinueOnFailure); err != nil {
				return err
			}
		}
	}

	if o.ExternalRole != nil {

		// query param external_role
		var qrExternalRole string

		if o.ExternalRole != nil {
			qrExternalRole = *o.ExternalRole
		}
		qExternalRole := qrExternalRole
		if qExternalRole != "" {

			if err := r.SetQueryParam("external_role", qExternalRole); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.OntapRoleName != nil {

		// query param ontap_role.name
		var qrOntapRoleName string

		if o.OntapRoleName != nil {
			qrOntapRoleName = *o.OntapRoleName
		}
		qOntapRoleName := qrOntapRoleName
		if qOntapRoleName != "" {

			if err := r.SetQueryParam("ontap_role.name", qOntapRoleName); err != nil {
				return err
			}
		}
	}

	if o.Provider != nil {

		// query param provider
		var qrProvider string

		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {

			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

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

	if o.SerialRecords != nil {

		// query param serial_records
		var qrSerialRecords bool

		if o.SerialRecords != nil {
			qrSerialRecords = *o.SerialRecords
		}
		qSerialRecords := swag.FormatBool(qrSerialRecords)
		if qSerialRecords != "" {

			if err := r.SetQueryParam("serial_records", qSerialRecords); err != nil {
				return err
			}
		}
	}

	if o.Timestamp != nil {

		// query param timestamp
		var qrTimestamp string

		if o.Timestamp != nil {
			qrTimestamp = *o.Timestamp
		}
		qTimestamp := qrTimestamp
		if qTimestamp != "" {

			if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
