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

// NewDuoModifyCollectionParams creates a new DuoModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDuoModifyCollectionParams() *DuoModifyCollectionParams {
	return &DuoModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDuoModifyCollectionParamsWithTimeout creates a new DuoModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewDuoModifyCollectionParamsWithTimeout(timeout time.Duration) *DuoModifyCollectionParams {
	return &DuoModifyCollectionParams{
		timeout: timeout,
	}
}

// NewDuoModifyCollectionParamsWithContext creates a new DuoModifyCollectionParams object
// with the ability to set a context for a request.
func NewDuoModifyCollectionParamsWithContext(ctx context.Context) *DuoModifyCollectionParams {
	return &DuoModifyCollectionParams{
		Context: ctx,
	}
}

// NewDuoModifyCollectionParamsWithHTTPClient creates a new DuoModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDuoModifyCollectionParamsWithHTTPClient(client *http.Client) *DuoModifyCollectionParams {
	return &DuoModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
DuoModifyCollectionParams contains all the parameters to send to the API endpoint

	for the duo modify collection operation.

	Typically these are written to a http.Request.
*/
type DuoModifyCollectionParams struct {

	/* APIHost.

	   Filter by api_host
	*/
	APIHost *string

	/* AutoPush.

	   Filter by auto_push
	*/
	AutoPush *bool

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* FailMode.

	   Filter by fail_mode
	*/
	FailMode *string

	/* Fingerprint.

	   Filter by fingerprint
	*/
	Fingerprint *string

	/* HTTPProxy.

	   Filter by http_proxy
	*/
	HTTPProxy *string

	/* Info.

	   Info specification
	*/
	Info DuoModifyCollectionBody

	/* IntegrationKey.

	   Filter by integration_key
	*/
	IntegrationKey *string

	/* IsEnabled.

	   Filter by is_enabled
	*/
	IsEnabled *bool

	/* MaxPrompts.

	   Filter by max_prompts
	*/
	MaxPrompts *int64

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerName *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUID *string

	/* PushInfo.

	   Filter by push_info
	*/
	PushInfo *bool

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

	/* Status.

	   Filter by status
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the duo modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuoModifyCollectionParams) WithDefaults() *DuoModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the duo modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuoModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := DuoModifyCollectionParams{
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

// WithTimeout adds the timeout to the duo modify collection params
func (o *DuoModifyCollectionParams) WithTimeout(timeout time.Duration) *DuoModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the duo modify collection params
func (o *DuoModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the duo modify collection params
func (o *DuoModifyCollectionParams) WithContext(ctx context.Context) *DuoModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the duo modify collection params
func (o *DuoModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the duo modify collection params
func (o *DuoModifyCollectionParams) WithHTTPClient(client *http.Client) *DuoModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the duo modify collection params
func (o *DuoModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIHost adds the aPIHost to the duo modify collection params
func (o *DuoModifyCollectionParams) WithAPIHost(aPIHost *string) *DuoModifyCollectionParams {
	o.SetAPIHost(aPIHost)
	return o
}

// SetAPIHost adds the apiHost to the duo modify collection params
func (o *DuoModifyCollectionParams) SetAPIHost(aPIHost *string) {
	o.APIHost = aPIHost
}

// WithAutoPush adds the autoPush to the duo modify collection params
func (o *DuoModifyCollectionParams) WithAutoPush(autoPush *bool) *DuoModifyCollectionParams {
	o.SetAutoPush(autoPush)
	return o
}

// SetAutoPush adds the autoPush to the duo modify collection params
func (o *DuoModifyCollectionParams) SetAutoPush(autoPush *bool) {
	o.AutoPush = autoPush
}

// WithComment adds the comment to the duo modify collection params
func (o *DuoModifyCollectionParams) WithComment(comment *string) *DuoModifyCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the duo modify collection params
func (o *DuoModifyCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContinueOnFailure adds the continueOnFailure to the duo modify collection params
func (o *DuoModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *DuoModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the duo modify collection params
func (o *DuoModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithFailMode adds the failMode to the duo modify collection params
func (o *DuoModifyCollectionParams) WithFailMode(failMode *string) *DuoModifyCollectionParams {
	o.SetFailMode(failMode)
	return o
}

// SetFailMode adds the failMode to the duo modify collection params
func (o *DuoModifyCollectionParams) SetFailMode(failMode *string) {
	o.FailMode = failMode
}

// WithFingerprint adds the fingerprint to the duo modify collection params
func (o *DuoModifyCollectionParams) WithFingerprint(fingerprint *string) *DuoModifyCollectionParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the duo modify collection params
func (o *DuoModifyCollectionParams) SetFingerprint(fingerprint *string) {
	o.Fingerprint = fingerprint
}

// WithHTTPProxy adds the hTTPProxy to the duo modify collection params
func (o *DuoModifyCollectionParams) WithHTTPProxy(hTTPProxy *string) *DuoModifyCollectionParams {
	o.SetHTTPProxy(hTTPProxy)
	return o
}

// SetHTTPProxy adds the httpProxy to the duo modify collection params
func (o *DuoModifyCollectionParams) SetHTTPProxy(hTTPProxy *string) {
	o.HTTPProxy = hTTPProxy
}

// WithInfo adds the info to the duo modify collection params
func (o *DuoModifyCollectionParams) WithInfo(info DuoModifyCollectionBody) *DuoModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the duo modify collection params
func (o *DuoModifyCollectionParams) SetInfo(info DuoModifyCollectionBody) {
	o.Info = info
}

// WithIntegrationKey adds the integrationKey to the duo modify collection params
func (o *DuoModifyCollectionParams) WithIntegrationKey(integrationKey *string) *DuoModifyCollectionParams {
	o.SetIntegrationKey(integrationKey)
	return o
}

// SetIntegrationKey adds the integrationKey to the duo modify collection params
func (o *DuoModifyCollectionParams) SetIntegrationKey(integrationKey *string) {
	o.IntegrationKey = integrationKey
}

// WithIsEnabled adds the isEnabled to the duo modify collection params
func (o *DuoModifyCollectionParams) WithIsEnabled(isEnabled *bool) *DuoModifyCollectionParams {
	o.SetIsEnabled(isEnabled)
	return o
}

// SetIsEnabled adds the isEnabled to the duo modify collection params
func (o *DuoModifyCollectionParams) SetIsEnabled(isEnabled *bool) {
	o.IsEnabled = isEnabled
}

// WithMaxPrompts adds the maxPrompts to the duo modify collection params
func (o *DuoModifyCollectionParams) WithMaxPrompts(maxPrompts *int64) *DuoModifyCollectionParams {
	o.SetMaxPrompts(maxPrompts)
	return o
}

// SetMaxPrompts adds the maxPrompts to the duo modify collection params
func (o *DuoModifyCollectionParams) SetMaxPrompts(maxPrompts *int64) {
	o.MaxPrompts = maxPrompts
}

// WithOwnerName adds the ownerName to the duo modify collection params
func (o *DuoModifyCollectionParams) WithOwnerName(ownerName *string) *DuoModifyCollectionParams {
	o.SetOwnerName(ownerName)
	return o
}

// SetOwnerName adds the ownerName to the duo modify collection params
func (o *DuoModifyCollectionParams) SetOwnerName(ownerName *string) {
	o.OwnerName = ownerName
}

// WithOwnerUUID adds the ownerUUID to the duo modify collection params
func (o *DuoModifyCollectionParams) WithOwnerUUID(ownerUUID *string) *DuoModifyCollectionParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the duo modify collection params
func (o *DuoModifyCollectionParams) SetOwnerUUID(ownerUUID *string) {
	o.OwnerUUID = ownerUUID
}

// WithPushInfo adds the pushInfo to the duo modify collection params
func (o *DuoModifyCollectionParams) WithPushInfo(pushInfo *bool) *DuoModifyCollectionParams {
	o.SetPushInfo(pushInfo)
	return o
}

// SetPushInfo adds the pushInfo to the duo modify collection params
func (o *DuoModifyCollectionParams) SetPushInfo(pushInfo *bool) {
	o.PushInfo = pushInfo
}

// WithReturnRecords adds the returnRecords to the duo modify collection params
func (o *DuoModifyCollectionParams) WithReturnRecords(returnRecords *bool) *DuoModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the duo modify collection params
func (o *DuoModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the duo modify collection params
func (o *DuoModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *DuoModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the duo modify collection params
func (o *DuoModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the duo modify collection params
func (o *DuoModifyCollectionParams) WithSerialRecords(serialRecords *bool) *DuoModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the duo modify collection params
func (o *DuoModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithStatus adds the status to the duo modify collection params
func (o *DuoModifyCollectionParams) WithStatus(status *string) *DuoModifyCollectionParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the duo modify collection params
func (o *DuoModifyCollectionParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DuoModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIHost != nil {

		// query param api_host
		var qrAPIHost string

		if o.APIHost != nil {
			qrAPIHost = *o.APIHost
		}
		qAPIHost := qrAPIHost
		if qAPIHost != "" {

			if err := r.SetQueryParam("api_host", qAPIHost); err != nil {
				return err
			}
		}
	}

	if o.AutoPush != nil {

		// query param auto_push
		var qrAutoPush bool

		if o.AutoPush != nil {
			qrAutoPush = *o.AutoPush
		}
		qAutoPush := swag.FormatBool(qrAutoPush)
		if qAutoPush != "" {

			if err := r.SetQueryParam("auto_push", qAutoPush); err != nil {
				return err
			}
		}
	}

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

	if o.FailMode != nil {

		// query param fail_mode
		var qrFailMode string

		if o.FailMode != nil {
			qrFailMode = *o.FailMode
		}
		qFailMode := qrFailMode
		if qFailMode != "" {

			if err := r.SetQueryParam("fail_mode", qFailMode); err != nil {
				return err
			}
		}
	}

	if o.Fingerprint != nil {

		// query param fingerprint
		var qrFingerprint string

		if o.Fingerprint != nil {
			qrFingerprint = *o.Fingerprint
		}
		qFingerprint := qrFingerprint
		if qFingerprint != "" {

			if err := r.SetQueryParam("fingerprint", qFingerprint); err != nil {
				return err
			}
		}
	}

	if o.HTTPProxy != nil {

		// query param http_proxy
		var qrHTTPProxy string

		if o.HTTPProxy != nil {
			qrHTTPProxy = *o.HTTPProxy
		}
		qHTTPProxy := qrHTTPProxy
		if qHTTPProxy != "" {

			if err := r.SetQueryParam("http_proxy", qHTTPProxy); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.IntegrationKey != nil {

		// query param integration_key
		var qrIntegrationKey string

		if o.IntegrationKey != nil {
			qrIntegrationKey = *o.IntegrationKey
		}
		qIntegrationKey := qrIntegrationKey
		if qIntegrationKey != "" {

			if err := r.SetQueryParam("integration_key", qIntegrationKey); err != nil {
				return err
			}
		}
	}

	if o.IsEnabled != nil {

		// query param is_enabled
		var qrIsEnabled bool

		if o.IsEnabled != nil {
			qrIsEnabled = *o.IsEnabled
		}
		qIsEnabled := swag.FormatBool(qrIsEnabled)
		if qIsEnabled != "" {

			if err := r.SetQueryParam("is_enabled", qIsEnabled); err != nil {
				return err
			}
		}
	}

	if o.MaxPrompts != nil {

		// query param max_prompts
		var qrMaxPrompts int64

		if o.MaxPrompts != nil {
			qrMaxPrompts = *o.MaxPrompts
		}
		qMaxPrompts := swag.FormatInt64(qrMaxPrompts)
		if qMaxPrompts != "" {

			if err := r.SetQueryParam("max_prompts", qMaxPrompts); err != nil {
				return err
			}
		}
	}

	if o.OwnerName != nil {

		// query param owner.name
		var qrOwnerName string

		if o.OwnerName != nil {
			qrOwnerName = *o.OwnerName
		}
		qOwnerName := qrOwnerName
		if qOwnerName != "" {

			if err := r.SetQueryParam("owner.name", qOwnerName); err != nil {
				return err
			}
		}
	}

	if o.OwnerUUID != nil {

		// query param owner.uuid
		var qrOwnerUUID string

		if o.OwnerUUID != nil {
			qrOwnerUUID = *o.OwnerUUID
		}
		qOwnerUUID := qrOwnerUUID
		if qOwnerUUID != "" {

			if err := r.SetQueryParam("owner.uuid", qOwnerUUID); err != nil {
				return err
			}
		}
	}

	if o.PushInfo != nil {

		// query param push_info
		var qrPushInfo bool

		if o.PushInfo != nil {
			qrPushInfo = *o.PushInfo
		}
		qPushInfo := swag.FormatBool(qrPushInfo)
		if qPushInfo != "" {

			if err := r.SetQueryParam("push_info", qPushInfo); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
