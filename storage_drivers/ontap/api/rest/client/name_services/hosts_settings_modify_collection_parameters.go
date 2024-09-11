// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewHostsSettingsModifyCollectionParams creates a new HostsSettingsModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsSettingsModifyCollectionParams() *HostsSettingsModifyCollectionParams {
	return &HostsSettingsModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsSettingsModifyCollectionParamsWithTimeout creates a new HostsSettingsModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewHostsSettingsModifyCollectionParamsWithTimeout(timeout time.Duration) *HostsSettingsModifyCollectionParams {
	return &HostsSettingsModifyCollectionParams{
		timeout: timeout,
	}
}

// NewHostsSettingsModifyCollectionParamsWithContext creates a new HostsSettingsModifyCollectionParams object
// with the ability to set a context for a request.
func NewHostsSettingsModifyCollectionParamsWithContext(ctx context.Context) *HostsSettingsModifyCollectionParams {
	return &HostsSettingsModifyCollectionParams{
		Context: ctx,
	}
}

// NewHostsSettingsModifyCollectionParamsWithHTTPClient creates a new HostsSettingsModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsSettingsModifyCollectionParamsWithHTTPClient(client *http.Client) *HostsSettingsModifyCollectionParams {
	return &HostsSettingsModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
HostsSettingsModifyCollectionParams contains all the parameters to send to the API endpoint

	for the hosts settings modify collection operation.

	Typically these are written to a http.Request.
*/
type HostsSettingsModifyCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* DNSTTLEnabled.

	   Filter by dns_ttl_enabled
	*/
	DNSTTLEnabled *bool

	/* Enabled.

	   Filter by enabled
	*/
	Enabled *bool

	/* Info.

	   Info specification
	*/
	Info HostsSettingsModifyCollectionBody

	/* NegativeCacheEnabled.

	   Filter by negative_cache_enabled
	*/
	NegativeCacheEnabled *bool

	/* NegativeTTL.

	   Filter by negative_ttl
	*/
	NegativeTTL *string

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

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* TTL.

	   Filter by ttl
	*/
	TTL *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hosts settings modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsSettingsModifyCollectionParams) WithDefaults() *HostsSettingsModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts settings modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsSettingsModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := HostsSettingsModifyCollectionParams{
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

// WithTimeout adds the timeout to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithTimeout(timeout time.Duration) *HostsSettingsModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithContext(ctx context.Context) *HostsSettingsModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithHTTPClient(client *http.Client) *HostsSettingsModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *HostsSettingsModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDNSTTLEnabled adds the dNSTTLEnabled to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithDNSTTLEnabled(dNSTTLEnabled *bool) *HostsSettingsModifyCollectionParams {
	o.SetDNSTTLEnabled(dNSTTLEnabled)
	return o
}

// SetDNSTTLEnabled adds the dnsTtlEnabled to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetDNSTTLEnabled(dNSTTLEnabled *bool) {
	o.DNSTTLEnabled = dNSTTLEnabled
}

// WithEnabled adds the enabled to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithEnabled(enabled *bool) *HostsSettingsModifyCollectionParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithInfo adds the info to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithInfo(info HostsSettingsModifyCollectionBody) *HostsSettingsModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetInfo(info HostsSettingsModifyCollectionBody) {
	o.Info = info
}

// WithNegativeCacheEnabled adds the negativeCacheEnabled to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithNegativeCacheEnabled(negativeCacheEnabled *bool) *HostsSettingsModifyCollectionParams {
	o.SetNegativeCacheEnabled(negativeCacheEnabled)
	return o
}

// SetNegativeCacheEnabled adds the negativeCacheEnabled to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetNegativeCacheEnabled(negativeCacheEnabled *bool) {
	o.NegativeCacheEnabled = negativeCacheEnabled
}

// WithNegativeTTL adds the negativeTTL to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithNegativeTTL(negativeTTL *string) *HostsSettingsModifyCollectionParams {
	o.SetNegativeTTL(negativeTTL)
	return o
}

// SetNegativeTTL adds the negativeTtl to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetNegativeTTL(negativeTTL *string) {
	o.NegativeTTL = negativeTTL
}

// WithReturnRecords adds the returnRecords to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithReturnRecords(returnRecords *bool) *HostsSettingsModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *HostsSettingsModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithSerialRecords(serialRecords *bool) *HostsSettingsModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithSvmName(svmName *string) *HostsSettingsModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithSvmUUID(svmUUID *string) *HostsSettingsModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithTTL adds the ttl to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithTTL(ttl *string) *HostsSettingsModifyCollectionParams {
	o.SetTTL(ttl)
	return o
}

// SetTTL adds the ttl to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetTTL(ttl *string) {
	o.TTL = ttl
}

// WithUUID adds the uuid to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) WithUUID(uuid *string) *HostsSettingsModifyCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the hosts settings modify collection params
func (o *HostsSettingsModifyCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *HostsSettingsModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DNSTTLEnabled != nil {

		// query param dns_ttl_enabled
		var qrDNSTTLEnabled bool

		if o.DNSTTLEnabled != nil {
			qrDNSTTLEnabled = *o.DNSTTLEnabled
		}
		qDNSTTLEnabled := swag.FormatBool(qrDNSTTLEnabled)
		if qDNSTTLEnabled != "" {

			if err := r.SetQueryParam("dns_ttl_enabled", qDNSTTLEnabled); err != nil {
				return err
			}
		}
	}

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.NegativeCacheEnabled != nil {

		// query param negative_cache_enabled
		var qrNegativeCacheEnabled bool

		if o.NegativeCacheEnabled != nil {
			qrNegativeCacheEnabled = *o.NegativeCacheEnabled
		}
		qNegativeCacheEnabled := swag.FormatBool(qrNegativeCacheEnabled)
		if qNegativeCacheEnabled != "" {

			if err := r.SetQueryParam("negative_cache_enabled", qNegativeCacheEnabled); err != nil {
				return err
			}
		}
	}

	if o.NegativeTTL != nil {

		// query param negative_ttl
		var qrNegativeTTL string

		if o.NegativeTTL != nil {
			qrNegativeTTL = *o.NegativeTTL
		}
		qNegativeTTL := qrNegativeTTL
		if qNegativeTTL != "" {

			if err := r.SetQueryParam("negative_ttl", qNegativeTTL); err != nil {
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

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.TTL != nil {

		// query param ttl
		var qrTTL string

		if o.TTL != nil {
			qrTTL = *o.TTL
		}
		qTTL := qrTTL
		if qTTL != "" {

			if err := r.SetQueryParam("ttl", qTTL); err != nil {
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
