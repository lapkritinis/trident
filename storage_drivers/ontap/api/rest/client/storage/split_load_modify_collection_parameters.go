// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewSplitLoadModifyCollectionParams creates a new SplitLoadModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSplitLoadModifyCollectionParams() *SplitLoadModifyCollectionParams {
	return &SplitLoadModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSplitLoadModifyCollectionParamsWithTimeout creates a new SplitLoadModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewSplitLoadModifyCollectionParamsWithTimeout(timeout time.Duration) *SplitLoadModifyCollectionParams {
	return &SplitLoadModifyCollectionParams{
		timeout: timeout,
	}
}

// NewSplitLoadModifyCollectionParamsWithContext creates a new SplitLoadModifyCollectionParams object
// with the ability to set a context for a request.
func NewSplitLoadModifyCollectionParamsWithContext(ctx context.Context) *SplitLoadModifyCollectionParams {
	return &SplitLoadModifyCollectionParams{
		Context: ctx,
	}
}

// NewSplitLoadModifyCollectionParamsWithHTTPClient creates a new SplitLoadModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSplitLoadModifyCollectionParamsWithHTTPClient(client *http.Client) *SplitLoadModifyCollectionParams {
	return &SplitLoadModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
SplitLoadModifyCollectionParams contains all the parameters to send to the API endpoint

	for the split load modify collection operation.

	Typically these are written to a http.Request.
*/
type SplitLoadModifyCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info SplitLoadModifyCollectionBody

	/* LoadAllowable.

	   Filter by load.allowable
	*/
	LoadAllowable *int64

	/* LoadCurrent.

	   Filter by load.current
	*/
	LoadCurrent *int64

	/* LoadMaximum.

	   Filter by load.maximum
	*/
	LoadMaximum *int64

	/* LoadTokenReserved.

	   Filter by load.token_reserved
	*/
	LoadTokenReserved *int64

	/* NodeName.

	   Filter by node.name
	*/
	NodeName *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUID *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the split load modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SplitLoadModifyCollectionParams) WithDefaults() *SplitLoadModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the split load modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SplitLoadModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SplitLoadModifyCollectionParams{
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

// WithTimeout adds the timeout to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithTimeout(timeout time.Duration) *SplitLoadModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithContext(ctx context.Context) *SplitLoadModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithHTTPClient(client *http.Client) *SplitLoadModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SplitLoadModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithInfo(info SplitLoadModifyCollectionBody) *SplitLoadModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetInfo(info SplitLoadModifyCollectionBody) {
	o.Info = info
}

// WithLoadAllowable adds the loadAllowable to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithLoadAllowable(loadAllowable *int64) *SplitLoadModifyCollectionParams {
	o.SetLoadAllowable(loadAllowable)
	return o
}

// SetLoadAllowable adds the loadAllowable to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetLoadAllowable(loadAllowable *int64) {
	o.LoadAllowable = loadAllowable
}

// WithLoadCurrent adds the loadCurrent to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithLoadCurrent(loadCurrent *int64) *SplitLoadModifyCollectionParams {
	o.SetLoadCurrent(loadCurrent)
	return o
}

// SetLoadCurrent adds the loadCurrent to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetLoadCurrent(loadCurrent *int64) {
	o.LoadCurrent = loadCurrent
}

// WithLoadMaximum adds the loadMaximum to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithLoadMaximum(loadMaximum *int64) *SplitLoadModifyCollectionParams {
	o.SetLoadMaximum(loadMaximum)
	return o
}

// SetLoadMaximum adds the loadMaximum to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetLoadMaximum(loadMaximum *int64) {
	o.LoadMaximum = loadMaximum
}

// WithLoadTokenReserved adds the loadTokenReserved to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithLoadTokenReserved(loadTokenReserved *int64) *SplitLoadModifyCollectionParams {
	o.SetLoadTokenReserved(loadTokenReserved)
	return o
}

// SetLoadTokenReserved adds the loadTokenReserved to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetLoadTokenReserved(loadTokenReserved *int64) {
	o.LoadTokenReserved = loadTokenReserved
}

// WithNodeName adds the nodeName to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithNodeName(nodeName *string) *SplitLoadModifyCollectionParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithNodeUUID adds the nodeUUID to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithNodeUUID(nodeUUID *string) *SplitLoadModifyCollectionParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetNodeUUID(nodeUUID *string) {
	o.NodeUUID = nodeUUID
}

// WithReturnRecords adds the returnRecords to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithReturnRecords(returnRecords *bool) *SplitLoadModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *SplitLoadModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) WithSerialRecords(serialRecords *bool) *SplitLoadModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the split load modify collection params
func (o *SplitLoadModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WriteToRequest writes these params to a swagger request
func (o *SplitLoadModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.LoadAllowable != nil {

		// query param load.allowable
		var qrLoadAllowable int64

		if o.LoadAllowable != nil {
			qrLoadAllowable = *o.LoadAllowable
		}
		qLoadAllowable := swag.FormatInt64(qrLoadAllowable)
		if qLoadAllowable != "" {

			if err := r.SetQueryParam("load.allowable", qLoadAllowable); err != nil {
				return err
			}
		}
	}

	if o.LoadCurrent != nil {

		// query param load.current
		var qrLoadCurrent int64

		if o.LoadCurrent != nil {
			qrLoadCurrent = *o.LoadCurrent
		}
		qLoadCurrent := swag.FormatInt64(qrLoadCurrent)
		if qLoadCurrent != "" {

			if err := r.SetQueryParam("load.current", qLoadCurrent); err != nil {
				return err
			}
		}
	}

	if o.LoadMaximum != nil {

		// query param load.maximum
		var qrLoadMaximum int64

		if o.LoadMaximum != nil {
			qrLoadMaximum = *o.LoadMaximum
		}
		qLoadMaximum := swag.FormatInt64(qrLoadMaximum)
		if qLoadMaximum != "" {

			if err := r.SetQueryParam("load.maximum", qLoadMaximum); err != nil {
				return err
			}
		}
	}

	if o.LoadTokenReserved != nil {

		// query param load.token_reserved
		var qrLoadTokenReserved int64

		if o.LoadTokenReserved != nil {
			qrLoadTokenReserved = *o.LoadTokenReserved
		}
		qLoadTokenReserved := swag.FormatInt64(qrLoadTokenReserved)
		if qLoadTokenReserved != "" {

			if err := r.SetQueryParam("load.token_reserved", qLoadTokenReserved); err != nil {
				return err
			}
		}
	}

	if o.NodeName != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeName != nil {
			qrNodeName = *o.NodeName
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUID != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUID != nil {
			qrNodeUUID = *o.NodeUUID
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
