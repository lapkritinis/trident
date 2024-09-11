// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewApplicationComponentSnapshotDeleteCollectionParams creates a new ApplicationComponentSnapshotDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationComponentSnapshotDeleteCollectionParams() *ApplicationComponentSnapshotDeleteCollectionParams {
	return &ApplicationComponentSnapshotDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationComponentSnapshotDeleteCollectionParamsWithTimeout creates a new ApplicationComponentSnapshotDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewApplicationComponentSnapshotDeleteCollectionParamsWithTimeout(timeout time.Duration) *ApplicationComponentSnapshotDeleteCollectionParams {
	return &ApplicationComponentSnapshotDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewApplicationComponentSnapshotDeleteCollectionParamsWithContext creates a new ApplicationComponentSnapshotDeleteCollectionParams object
// with the ability to set a context for a request.
func NewApplicationComponentSnapshotDeleteCollectionParamsWithContext(ctx context.Context) *ApplicationComponentSnapshotDeleteCollectionParams {
	return &ApplicationComponentSnapshotDeleteCollectionParams{
		Context: ctx,
	}
}

// NewApplicationComponentSnapshotDeleteCollectionParamsWithHTTPClient creates a new ApplicationComponentSnapshotDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationComponentSnapshotDeleteCollectionParamsWithHTTPClient(client *http.Client) *ApplicationComponentSnapshotDeleteCollectionParams {
	return &ApplicationComponentSnapshotDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
ApplicationComponentSnapshotDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the application component snapshot delete collection operation.

	Typically these are written to a http.Request.
*/
type ApplicationComponentSnapshotDeleteCollectionParams struct {

	/* ApplicationUUID.

	   Application UUID
	*/
	ApplicationUUID string

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ComponentName.

	   Filter by Application Component Name
	*/
	ComponentName *string

	/* ComponentUUID.

	   Application Component UUID
	*/
	ComponentUUID string

	/* ConsistencyType.

	   Filter by consistency_type
	*/
	ConsistencyType *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* CreateTime.

	   Filter by create_time
	*/
	CreateTime *string

	/* Info.

	   Info specification
	*/
	Info ApplicationComponentSnapshotDeleteCollectionBody

	/* IsPartial.

	   Filter by is_partial
	*/
	IsPartial *string

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

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application component snapshot delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithDefaults() *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application component snapshot delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := ApplicationComponentSnapshotDeleteCollectionParams{
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

// WithTimeout adds the timeout to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithTimeout(timeout time.Duration) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithContext(ctx context.Context) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithHTTPClient(client *http.Client) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationUUID adds the applicationUUID to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithApplicationUUID(applicationUUID string) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetApplicationUUID(applicationUUID)
	return o
}

// SetApplicationUUID adds the applicationUuid to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetApplicationUUID(applicationUUID string) {
	o.ApplicationUUID = applicationUUID
}

// WithComment adds the comment to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithComment(comment *string) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithComponentName adds the componentName to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithComponentName(componentName *string) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetComponentName(componentName *string) {
	o.ComponentName = componentName
}

// WithComponentUUID adds the componentUUID to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithComponentUUID(componentUUID string) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetComponentUUID(componentUUID)
	return o
}

// SetComponentUUID adds the componentUuid to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetComponentUUID(componentUUID string) {
	o.ComponentUUID = componentUUID
}

// WithConsistencyType adds the consistencyType to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithConsistencyType(consistencyType *string) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetConsistencyType(consistencyType)
	return o
}

// SetConsistencyType adds the consistencyType to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetConsistencyType(consistencyType *string) {
	o.ConsistencyType = consistencyType
}

// WithContinueOnFailure adds the continueOnFailure to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithCreateTime adds the createTime to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithCreateTime(createTime *string) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetCreateTime(createTime)
	return o
}

// SetCreateTime adds the createTime to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetCreateTime(createTime *string) {
	o.CreateTime = createTime
}

// WithInfo adds the info to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithInfo(info ApplicationComponentSnapshotDeleteCollectionBody) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetInfo(info ApplicationComponentSnapshotDeleteCollectionBody) {
	o.Info = info
}

// WithIsPartial adds the isPartial to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithIsPartial(isPartial *string) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetIsPartial(isPartial)
	return o
}

// SetIsPartial adds the isPartial to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetIsPartial(isPartial *string) {
	o.IsPartial = isPartial
}

// WithName adds the name to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithName(name *string) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReturnRecords adds the returnRecords to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithUUID adds the uuid to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WithUUID(uuid *string) *ApplicationComponentSnapshotDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the application component snapshot delete collection params
func (o *ApplicationComponentSnapshotDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationComponentSnapshotDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.uuid
	if err := r.SetPathParam("application.uuid", o.ApplicationUUID); err != nil {
		return err
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

	if o.ComponentName != nil {

		// query param component.name
		var qrComponentName string

		if o.ComponentName != nil {
			qrComponentName = *o.ComponentName
		}
		qComponentName := qrComponentName
		if qComponentName != "" {

			if err := r.SetQueryParam("component.name", qComponentName); err != nil {
				return err
			}
		}
	}

	// path param component.uuid
	if err := r.SetPathParam("component.uuid", o.ComponentUUID); err != nil {
		return err
	}

	if o.ConsistencyType != nil {

		// query param consistency_type
		var qrConsistencyType string

		if o.ConsistencyType != nil {
			qrConsistencyType = *o.ConsistencyType
		}
		qConsistencyType := qrConsistencyType
		if qConsistencyType != "" {

			if err := r.SetQueryParam("consistency_type", qConsistencyType); err != nil {
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

	if o.CreateTime != nil {

		// query param create_time
		var qrCreateTime string

		if o.CreateTime != nil {
			qrCreateTime = *o.CreateTime
		}
		qCreateTime := qrCreateTime
		if qCreateTime != "" {

			if err := r.SetQueryParam("create_time", qCreateTime); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.IsPartial != nil {

		// query param is_partial
		var qrIsPartial string

		if o.IsPartial != nil {
			qrIsPartial = *o.IsPartial
		}
		qIsPartial := qrIsPartial
		if qIsPartial != "" {

			if err := r.SetQueryParam("is_partial", qIsPartial); err != nil {
				return err
			}
		}
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
