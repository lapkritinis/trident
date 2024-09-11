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

// NewDuogroupModifyCollectionParams creates a new DuogroupModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDuogroupModifyCollectionParams() *DuogroupModifyCollectionParams {
	return &DuogroupModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDuogroupModifyCollectionParamsWithTimeout creates a new DuogroupModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewDuogroupModifyCollectionParamsWithTimeout(timeout time.Duration) *DuogroupModifyCollectionParams {
	return &DuogroupModifyCollectionParams{
		timeout: timeout,
	}
}

// NewDuogroupModifyCollectionParamsWithContext creates a new DuogroupModifyCollectionParams object
// with the ability to set a context for a request.
func NewDuogroupModifyCollectionParamsWithContext(ctx context.Context) *DuogroupModifyCollectionParams {
	return &DuogroupModifyCollectionParams{
		Context: ctx,
	}
}

// NewDuogroupModifyCollectionParamsWithHTTPClient creates a new DuogroupModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDuogroupModifyCollectionParamsWithHTTPClient(client *http.Client) *DuogroupModifyCollectionParams {
	return &DuogroupModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
DuogroupModifyCollectionParams contains all the parameters to send to the API endpoint

	for the duogroup modify collection operation.

	Typically these are written to a http.Request.
*/
type DuogroupModifyCollectionParams struct {

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* ExcludedUsers.

	   Filter by excluded_users
	*/
	ExcludedUsers *string

	/* Info.

	   Info specification
	*/
	Info DuogroupModifyCollectionBody

	/* Name.

	   Filter by name
	*/
	Name *string

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerName *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUID *string

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

// WithDefaults hydrates default values in the duogroup modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuogroupModifyCollectionParams) WithDefaults() *DuogroupModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the duogroup modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuogroupModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := DuogroupModifyCollectionParams{
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

// WithTimeout adds the timeout to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithTimeout(timeout time.Duration) *DuogroupModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithContext(ctx context.Context) *DuogroupModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithHTTPClient(client *http.Client) *DuogroupModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithComment(comment *string) *DuogroupModifyCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContinueOnFailure adds the continueOnFailure to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *DuogroupModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithExcludedUsers adds the excludedUsers to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithExcludedUsers(excludedUsers *string) *DuogroupModifyCollectionParams {
	o.SetExcludedUsers(excludedUsers)
	return o
}

// SetExcludedUsers adds the excludedUsers to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetExcludedUsers(excludedUsers *string) {
	o.ExcludedUsers = excludedUsers
}

// WithInfo adds the info to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithInfo(info DuogroupModifyCollectionBody) *DuogroupModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetInfo(info DuogroupModifyCollectionBody) {
	o.Info = info
}

// WithName adds the name to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithName(name *string) *DuogroupModifyCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithOwnerName adds the ownerName to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithOwnerName(ownerName *string) *DuogroupModifyCollectionParams {
	o.SetOwnerName(ownerName)
	return o
}

// SetOwnerName adds the ownerName to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetOwnerName(ownerName *string) {
	o.OwnerName = ownerName
}

// WithOwnerUUID adds the ownerUUID to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithOwnerUUID(ownerUUID *string) *DuogroupModifyCollectionParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetOwnerUUID(ownerUUID *string) {
	o.OwnerUUID = ownerUUID
}

// WithReturnRecords adds the returnRecords to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithReturnRecords(returnRecords *bool) *DuogroupModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *DuogroupModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) WithSerialRecords(serialRecords *bool) *DuogroupModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the duogroup modify collection params
func (o *DuogroupModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WriteToRequest writes these params to a swagger request
func (o *DuogroupModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExcludedUsers != nil {

		// query param excluded_users
		var qrExcludedUsers string

		if o.ExcludedUsers != nil {
			qrExcludedUsers = *o.ExcludedUsers
		}
		qExcludedUsers := qrExcludedUsers
		if qExcludedUsers != "" {

			if err := r.SetQueryParam("excluded_users", qExcludedUsers); err != nil {
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
