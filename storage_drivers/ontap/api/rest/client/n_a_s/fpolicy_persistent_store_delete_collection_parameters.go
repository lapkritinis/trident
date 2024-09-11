// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewFpolicyPersistentStoreDeleteCollectionParams creates a new FpolicyPersistentStoreDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyPersistentStoreDeleteCollectionParams() *FpolicyPersistentStoreDeleteCollectionParams {
	return &FpolicyPersistentStoreDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyPersistentStoreDeleteCollectionParamsWithTimeout creates a new FpolicyPersistentStoreDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewFpolicyPersistentStoreDeleteCollectionParamsWithTimeout(timeout time.Duration) *FpolicyPersistentStoreDeleteCollectionParams {
	return &FpolicyPersistentStoreDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewFpolicyPersistentStoreDeleteCollectionParamsWithContext creates a new FpolicyPersistentStoreDeleteCollectionParams object
// with the ability to set a context for a request.
func NewFpolicyPersistentStoreDeleteCollectionParamsWithContext(ctx context.Context) *FpolicyPersistentStoreDeleteCollectionParams {
	return &FpolicyPersistentStoreDeleteCollectionParams{
		Context: ctx,
	}
}

// NewFpolicyPersistentStoreDeleteCollectionParamsWithHTTPClient creates a new FpolicyPersistentStoreDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyPersistentStoreDeleteCollectionParamsWithHTTPClient(client *http.Client) *FpolicyPersistentStoreDeleteCollectionParams {
	return &FpolicyPersistentStoreDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
FpolicyPersistentStoreDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the fpolicy persistent store delete collection operation.

	Typically these are written to a http.Request.
*/
type FpolicyPersistentStoreDeleteCollectionParams struct {

	/* AutosizeMode.

	   Filter by autosize_mode
	*/
	AutosizeMode *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* DeleteVolume.

	   Delete the associated volume of the Persistent Store.
	*/
	DeleteVolume *bool

	/* Info.

	   Info specification
	*/
	Info FpolicyPersistentStoreDeleteCollectionBody

	/* Name.

	   Filter by name
	*/
	Name *string

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

	/* Size.

	   Filter by size
	*/
	Size *int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	/* Volume.

	   Filter by volume
	*/
	Volume *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fpolicy persistent store delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithDefaults() *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy persistent store delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		deleteVolumeDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := FpolicyPersistentStoreDeleteCollectionParams{
		ContinueOnFailure: &continueOnFailureDefault,
		DeleteVolume:      &deleteVolumeDefault,
		ReturnRecords:     &returnRecordsDefault,
		ReturnTimeout:     &returnTimeoutDefault,
		SerialRecords:     &serialRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithTimeout(timeout time.Duration) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithContext(ctx context.Context) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithHTTPClient(client *http.Client) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAutosizeMode adds the autosizeMode to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithAutosizeMode(autosizeMode *string) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetAutosizeMode(autosizeMode)
	return o
}

// SetAutosizeMode adds the autosizeMode to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetAutosizeMode(autosizeMode *string) {
	o.AutosizeMode = autosizeMode
}

// WithContinueOnFailure adds the continueOnFailure to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDeleteVolume adds the deleteVolume to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithDeleteVolume(deleteVolume *bool) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetDeleteVolume(deleteVolume)
	return o
}

// SetDeleteVolume adds the deleteVolume to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetDeleteVolume(deleteVolume *bool) {
	o.DeleteVolume = deleteVolume
}

// WithInfo adds the info to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithInfo(info FpolicyPersistentStoreDeleteCollectionBody) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetInfo(info FpolicyPersistentStoreDeleteCollectionBody) {
	o.Info = info
}

// WithName adds the name to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithName(name *string) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReturnRecords adds the returnRecords to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSize adds the size to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithSize(size *int64) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetSize(size *int64) {
	o.Size = size
}

// WithSvmUUID adds the svmUUID to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithSvmUUID(svmUUID string) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithVolume adds the volume to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) WithVolume(volume *string) *FpolicyPersistentStoreDeleteCollectionParams {
	o.SetVolume(volume)
	return o
}

// SetVolume adds the volume to the fpolicy persistent store delete collection params
func (o *FpolicyPersistentStoreDeleteCollectionParams) SetVolume(volume *string) {
	o.Volume = volume
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyPersistentStoreDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AutosizeMode != nil {

		// query param autosize_mode
		var qrAutosizeMode string

		if o.AutosizeMode != nil {
			qrAutosizeMode = *o.AutosizeMode
		}
		qAutosizeMode := qrAutosizeMode
		if qAutosizeMode != "" {

			if err := r.SetQueryParam("autosize_mode", qAutosizeMode); err != nil {
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

	if o.DeleteVolume != nil {

		// query param delete_volume
		var qrDeleteVolume bool

		if o.DeleteVolume != nil {
			qrDeleteVolume = *o.DeleteVolume
		}
		qDeleteVolume := swag.FormatBool(qrDeleteVolume)
		if qDeleteVolume != "" {

			if err := r.SetQueryParam("delete_volume", qDeleteVolume); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.Size != nil {

		// query param size
		var qrSize int64

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if o.Volume != nil {

		// query param volume
		var qrVolume string

		if o.Volume != nil {
			qrVolume = *o.Volume
		}
		qVolume := qrVolume
		if qVolume != "" {

			if err := r.SetQueryParam("volume", qVolume); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
