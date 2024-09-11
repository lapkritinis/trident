// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewLicenseManagerModifyCollectionParams creates a new LicenseManagerModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLicenseManagerModifyCollectionParams() *LicenseManagerModifyCollectionParams {
	return &LicenseManagerModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLicenseManagerModifyCollectionParamsWithTimeout creates a new LicenseManagerModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewLicenseManagerModifyCollectionParamsWithTimeout(timeout time.Duration) *LicenseManagerModifyCollectionParams {
	return &LicenseManagerModifyCollectionParams{
		timeout: timeout,
	}
}

// NewLicenseManagerModifyCollectionParamsWithContext creates a new LicenseManagerModifyCollectionParams object
// with the ability to set a context for a request.
func NewLicenseManagerModifyCollectionParamsWithContext(ctx context.Context) *LicenseManagerModifyCollectionParams {
	return &LicenseManagerModifyCollectionParams{
		Context: ctx,
	}
}

// NewLicenseManagerModifyCollectionParamsWithHTTPClient creates a new LicenseManagerModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewLicenseManagerModifyCollectionParamsWithHTTPClient(client *http.Client) *LicenseManagerModifyCollectionParams {
	return &LicenseManagerModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
LicenseManagerModifyCollectionParams contains all the parameters to send to the API endpoint

	for the license manager modify collection operation.

	Typically these are written to a http.Request.
*/
type LicenseManagerModifyCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Default.

	   Filter by default
	*/
	Default *bool

	/* Info.

	   Info specification
	*/
	Info LicenseManagerModifyCollectionBody

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

	/* URIHost.

	   Filter by uri.host
	*/
	URIHost *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the license manager modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseManagerModifyCollectionParams) WithDefaults() *LicenseManagerModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the license manager modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseManagerModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := LicenseManagerModifyCollectionParams{
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

// WithTimeout adds the timeout to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithTimeout(timeout time.Duration) *LicenseManagerModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithContext(ctx context.Context) *LicenseManagerModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithHTTPClient(client *http.Client) *LicenseManagerModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *LicenseManagerModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDefault adds the defaultVar to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithDefault(defaultVar *bool) *LicenseManagerModifyCollectionParams {
	o.SetDefault(defaultVar)
	return o
}

// SetDefault adds the default to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetDefault(defaultVar *bool) {
	o.Default = defaultVar
}

// WithInfo adds the info to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithInfo(info LicenseManagerModifyCollectionBody) *LicenseManagerModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetInfo(info LicenseManagerModifyCollectionBody) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithReturnRecords(returnRecords *bool) *LicenseManagerModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *LicenseManagerModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithSerialRecords(serialRecords *bool) *LicenseManagerModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithURIHost adds the uRIHost to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithURIHost(uRIHost *string) *LicenseManagerModifyCollectionParams {
	o.SetURIHost(uRIHost)
	return o
}

// SetURIHost adds the uriHost to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetURIHost(uRIHost *string) {
	o.URIHost = uRIHost
}

// WithUUID adds the uuid to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) WithUUID(uuid *string) *LicenseManagerModifyCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the license manager modify collection params
func (o *LicenseManagerModifyCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *LicenseManagerModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Default != nil {

		// query param default
		var qrDefault bool

		if o.Default != nil {
			qrDefault = *o.Default
		}
		qDefault := swag.FormatBool(qrDefault)
		if qDefault != "" {

			if err := r.SetQueryParam("default", qDefault); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
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

	if o.URIHost != nil {

		// query param uri.host
		var qrURIHost string

		if o.URIHost != nil {
			qrURIHost = *o.URIHost
		}
		qURIHost := qrURIHost
		if qURIHost != "" {

			if err := r.SetQueryParam("uri.host", qURIHost); err != nil {
				return err
			}
		}
	}

	if o.UUID != nil {

		// query param uuid
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
