// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// NewSnaplockLogModifyCollectionParams creates a new SnaplockLogModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLogModifyCollectionParams() *SnaplockLogModifyCollectionParams {
	return &SnaplockLogModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLogModifyCollectionParamsWithTimeout creates a new SnaplockLogModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewSnaplockLogModifyCollectionParamsWithTimeout(timeout time.Duration) *SnaplockLogModifyCollectionParams {
	return &SnaplockLogModifyCollectionParams{
		timeout: timeout,
	}
}

// NewSnaplockLogModifyCollectionParamsWithContext creates a new SnaplockLogModifyCollectionParams object
// with the ability to set a context for a request.
func NewSnaplockLogModifyCollectionParamsWithContext(ctx context.Context) *SnaplockLogModifyCollectionParams {
	return &SnaplockLogModifyCollectionParams{
		Context: ctx,
	}
}

// NewSnaplockLogModifyCollectionParamsWithHTTPClient creates a new SnaplockLogModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLogModifyCollectionParamsWithHTTPClient(client *http.Client) *SnaplockLogModifyCollectionParams {
	return &SnaplockLogModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
SnaplockLogModifyCollectionParams contains all the parameters to send to the API endpoint

	for the snaplock log modify collection operation.

	Typically these are written to a http.Request.
*/
type SnaplockLogModifyCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info SnaplockLogModifyCollectionBody

	/* LogFilesBaseName.

	   Filter by log_files.base_name
	*/
	LogFilesBaseName *string

	/* LogFilesExpiryTime.

	   Filter by log_files.expiry_time
	*/
	LogFilesExpiryTime *string

	/* LogFilesPath.

	   Filter by log_files.path
	*/
	LogFilesPath *string

	/* LogFilesSize.

	   Filter by log_files.size
	*/
	LogFilesSize *int64

	/* LogVolumeMaxLogSize.

	   Filter by log_volume.max_log_size
	*/
	LogVolumeMaxLogSize *int64

	/* LogVolumeRetentionPeriod.

	   Filter by log_volume.retention_period
	*/
	LogVolumeRetentionPeriod *string

	/* LogVolumeVolumeName.

	   Filter by log_volume.volume.name
	*/
	LogVolumeVolumeName *string

	/* LogVolumeVolumeUUID.

	   Filter by log_volume.volume.uuid
	*/
	LogVolumeVolumeUUID *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock log modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLogModifyCollectionParams) WithDefaults() *SnaplockLogModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock log modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLogModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SnaplockLogModifyCollectionParams{
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

// WithTimeout adds the timeout to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithTimeout(timeout time.Duration) *SnaplockLogModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithContext(ctx context.Context) *SnaplockLogModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithHTTPClient(client *http.Client) *SnaplockLogModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SnaplockLogModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithInfo(info SnaplockLogModifyCollectionBody) *SnaplockLogModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetInfo(info SnaplockLogModifyCollectionBody) {
	o.Info = info
}

// WithLogFilesBaseName adds the logFilesBaseName to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithLogFilesBaseName(logFilesBaseName *string) *SnaplockLogModifyCollectionParams {
	o.SetLogFilesBaseName(logFilesBaseName)
	return o
}

// SetLogFilesBaseName adds the logFilesBaseName to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetLogFilesBaseName(logFilesBaseName *string) {
	o.LogFilesBaseName = logFilesBaseName
}

// WithLogFilesExpiryTime adds the logFilesExpiryTime to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithLogFilesExpiryTime(logFilesExpiryTime *string) *SnaplockLogModifyCollectionParams {
	o.SetLogFilesExpiryTime(logFilesExpiryTime)
	return o
}

// SetLogFilesExpiryTime adds the logFilesExpiryTime to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetLogFilesExpiryTime(logFilesExpiryTime *string) {
	o.LogFilesExpiryTime = logFilesExpiryTime
}

// WithLogFilesPath adds the logFilesPath to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithLogFilesPath(logFilesPath *string) *SnaplockLogModifyCollectionParams {
	o.SetLogFilesPath(logFilesPath)
	return o
}

// SetLogFilesPath adds the logFilesPath to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetLogFilesPath(logFilesPath *string) {
	o.LogFilesPath = logFilesPath
}

// WithLogFilesSize adds the logFilesSize to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithLogFilesSize(logFilesSize *int64) *SnaplockLogModifyCollectionParams {
	o.SetLogFilesSize(logFilesSize)
	return o
}

// SetLogFilesSize adds the logFilesSize to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetLogFilesSize(logFilesSize *int64) {
	o.LogFilesSize = logFilesSize
}

// WithLogVolumeMaxLogSize adds the logVolumeMaxLogSize to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithLogVolumeMaxLogSize(logVolumeMaxLogSize *int64) *SnaplockLogModifyCollectionParams {
	o.SetLogVolumeMaxLogSize(logVolumeMaxLogSize)
	return o
}

// SetLogVolumeMaxLogSize adds the logVolumeMaxLogSize to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetLogVolumeMaxLogSize(logVolumeMaxLogSize *int64) {
	o.LogVolumeMaxLogSize = logVolumeMaxLogSize
}

// WithLogVolumeRetentionPeriod adds the logVolumeRetentionPeriod to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithLogVolumeRetentionPeriod(logVolumeRetentionPeriod *string) *SnaplockLogModifyCollectionParams {
	o.SetLogVolumeRetentionPeriod(logVolumeRetentionPeriod)
	return o
}

// SetLogVolumeRetentionPeriod adds the logVolumeRetentionPeriod to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetLogVolumeRetentionPeriod(logVolumeRetentionPeriod *string) {
	o.LogVolumeRetentionPeriod = logVolumeRetentionPeriod
}

// WithLogVolumeVolumeName adds the logVolumeVolumeName to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithLogVolumeVolumeName(logVolumeVolumeName *string) *SnaplockLogModifyCollectionParams {
	o.SetLogVolumeVolumeName(logVolumeVolumeName)
	return o
}

// SetLogVolumeVolumeName adds the logVolumeVolumeName to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetLogVolumeVolumeName(logVolumeVolumeName *string) {
	o.LogVolumeVolumeName = logVolumeVolumeName
}

// WithLogVolumeVolumeUUID adds the logVolumeVolumeUUID to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithLogVolumeVolumeUUID(logVolumeVolumeUUID *string) *SnaplockLogModifyCollectionParams {
	o.SetLogVolumeVolumeUUID(logVolumeVolumeUUID)
	return o
}

// SetLogVolumeVolumeUUID adds the logVolumeVolumeUuid to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetLogVolumeVolumeUUID(logVolumeVolumeUUID *string) {
	o.LogVolumeVolumeUUID = logVolumeVolumeUUID
}

// WithReturnRecords adds the returnRecords to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithReturnRecords(returnRecords *bool) *SnaplockLogModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *SnaplockLogModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithSerialRecords(serialRecords *bool) *SnaplockLogModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithSvmName(svmName *string) *SnaplockLogModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) WithSvmUUID(svmUUID *string) *SnaplockLogModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the snaplock log modify collection params
func (o *SnaplockLogModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLogModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LogFilesBaseName != nil {

		// query param log_files.base_name
		var qrLogFilesBaseName string

		if o.LogFilesBaseName != nil {
			qrLogFilesBaseName = *o.LogFilesBaseName
		}
		qLogFilesBaseName := qrLogFilesBaseName
		if qLogFilesBaseName != "" {

			if err := r.SetQueryParam("log_files.base_name", qLogFilesBaseName); err != nil {
				return err
			}
		}
	}

	if o.LogFilesExpiryTime != nil {

		// query param log_files.expiry_time
		var qrLogFilesExpiryTime string

		if o.LogFilesExpiryTime != nil {
			qrLogFilesExpiryTime = *o.LogFilesExpiryTime
		}
		qLogFilesExpiryTime := qrLogFilesExpiryTime
		if qLogFilesExpiryTime != "" {

			if err := r.SetQueryParam("log_files.expiry_time", qLogFilesExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.LogFilesPath != nil {

		// query param log_files.path
		var qrLogFilesPath string

		if o.LogFilesPath != nil {
			qrLogFilesPath = *o.LogFilesPath
		}
		qLogFilesPath := qrLogFilesPath
		if qLogFilesPath != "" {

			if err := r.SetQueryParam("log_files.path", qLogFilesPath); err != nil {
				return err
			}
		}
	}

	if o.LogFilesSize != nil {

		// query param log_files.size
		var qrLogFilesSize int64

		if o.LogFilesSize != nil {
			qrLogFilesSize = *o.LogFilesSize
		}
		qLogFilesSize := swag.FormatInt64(qrLogFilesSize)
		if qLogFilesSize != "" {

			if err := r.SetQueryParam("log_files.size", qLogFilesSize); err != nil {
				return err
			}
		}
	}

	if o.LogVolumeMaxLogSize != nil {

		// query param log_volume.max_log_size
		var qrLogVolumeMaxLogSize int64

		if o.LogVolumeMaxLogSize != nil {
			qrLogVolumeMaxLogSize = *o.LogVolumeMaxLogSize
		}
		qLogVolumeMaxLogSize := swag.FormatInt64(qrLogVolumeMaxLogSize)
		if qLogVolumeMaxLogSize != "" {

			if err := r.SetQueryParam("log_volume.max_log_size", qLogVolumeMaxLogSize); err != nil {
				return err
			}
		}
	}

	if o.LogVolumeRetentionPeriod != nil {

		// query param log_volume.retention_period
		var qrLogVolumeRetentionPeriod string

		if o.LogVolumeRetentionPeriod != nil {
			qrLogVolumeRetentionPeriod = *o.LogVolumeRetentionPeriod
		}
		qLogVolumeRetentionPeriod := qrLogVolumeRetentionPeriod
		if qLogVolumeRetentionPeriod != "" {

			if err := r.SetQueryParam("log_volume.retention_period", qLogVolumeRetentionPeriod); err != nil {
				return err
			}
		}
	}

	if o.LogVolumeVolumeName != nil {

		// query param log_volume.volume.name
		var qrLogVolumeVolumeName string

		if o.LogVolumeVolumeName != nil {
			qrLogVolumeVolumeName = *o.LogVolumeVolumeName
		}
		qLogVolumeVolumeName := qrLogVolumeVolumeName
		if qLogVolumeVolumeName != "" {

			if err := r.SetQueryParam("log_volume.volume.name", qLogVolumeVolumeName); err != nil {
				return err
			}
		}
	}

	if o.LogVolumeVolumeUUID != nil {

		// query param log_volume.volume.uuid
		var qrLogVolumeVolumeUUID string

		if o.LogVolumeVolumeUUID != nil {
			qrLogVolumeVolumeUUID = *o.LogVolumeVolumeUUID
		}
		qLogVolumeVolumeUUID := qrLogVolumeVolumeUUID
		if qLogVolumeVolumeUUID != "" {

			if err := r.SetQueryParam("log_volume.volume.uuid", qLogVolumeVolumeUUID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
