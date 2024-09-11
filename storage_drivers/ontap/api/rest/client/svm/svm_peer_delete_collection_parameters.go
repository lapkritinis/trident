// Code generated by go-swagger; DO NOT EDIT.

package svm

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

// NewSvmPeerDeleteCollectionParams creates a new SvmPeerDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmPeerDeleteCollectionParams() *SvmPeerDeleteCollectionParams {
	return &SvmPeerDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmPeerDeleteCollectionParamsWithTimeout creates a new SvmPeerDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewSvmPeerDeleteCollectionParamsWithTimeout(timeout time.Duration) *SvmPeerDeleteCollectionParams {
	return &SvmPeerDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewSvmPeerDeleteCollectionParamsWithContext creates a new SvmPeerDeleteCollectionParams object
// with the ability to set a context for a request.
func NewSvmPeerDeleteCollectionParamsWithContext(ctx context.Context) *SvmPeerDeleteCollectionParams {
	return &SvmPeerDeleteCollectionParams{
		Context: ctx,
	}
}

// NewSvmPeerDeleteCollectionParamsWithHTTPClient creates a new SvmPeerDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmPeerDeleteCollectionParamsWithHTTPClient(client *http.Client) *SvmPeerDeleteCollectionParams {
	return &SvmPeerDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
SvmPeerDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the svm peer delete collection operation.

	Typically these are written to a http.Request.
*/
type SvmPeerDeleteCollectionParams struct {

	/* Applications.

	   Filter by applications
	*/
	Applications *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info SvmPeerDeleteCollectionBody

	/* Name.

	   Filter by name
	*/
	Name *string

	/* PeerClusterName.

	   Filter by peer.cluster.name
	*/
	PeerClusterName *string

	/* PeerClusterUUID.

	   Filter by peer.cluster.uuid
	*/
	PeerClusterUUID *string

	/* PeerSvmName.

	   Filter by peer.svm.name
	*/
	PeerSvmName *string

	/* PeerSvmUUID.

	   Filter by peer.svm.uuid
	*/
	PeerSvmUUID *string

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

	/* State.

	   Filter by state
	*/
	State *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm peer delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerDeleteCollectionParams) WithDefaults() *SvmPeerDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm peer delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SvmPeerDeleteCollectionParams{
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

// WithTimeout adds the timeout to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithTimeout(timeout time.Duration) *SvmPeerDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithContext(ctx context.Context) *SvmPeerDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithHTTPClient(client *http.Client) *SvmPeerDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplications adds the applications to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithApplications(applications *string) *SvmPeerDeleteCollectionParams {
	o.SetApplications(applications)
	return o
}

// SetApplications adds the applications to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetApplications(applications *string) {
	o.Applications = applications
}

// WithContinueOnFailure adds the continueOnFailure to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SvmPeerDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithInfo(info SvmPeerDeleteCollectionBody) *SvmPeerDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetInfo(info SvmPeerDeleteCollectionBody) {
	o.Info = info
}

// WithName adds the name to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithName(name *string) *SvmPeerDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithPeerClusterName adds the peerClusterName to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithPeerClusterName(peerClusterName *string) *SvmPeerDeleteCollectionParams {
	o.SetPeerClusterName(peerClusterName)
	return o
}

// SetPeerClusterName adds the peerClusterName to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetPeerClusterName(peerClusterName *string) {
	o.PeerClusterName = peerClusterName
}

// WithPeerClusterUUID adds the peerClusterUUID to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithPeerClusterUUID(peerClusterUUID *string) *SvmPeerDeleteCollectionParams {
	o.SetPeerClusterUUID(peerClusterUUID)
	return o
}

// SetPeerClusterUUID adds the peerClusterUuid to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetPeerClusterUUID(peerClusterUUID *string) {
	o.PeerClusterUUID = peerClusterUUID
}

// WithPeerSvmName adds the peerSvmName to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithPeerSvmName(peerSvmName *string) *SvmPeerDeleteCollectionParams {
	o.SetPeerSvmName(peerSvmName)
	return o
}

// SetPeerSvmName adds the peerSvmName to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetPeerSvmName(peerSvmName *string) {
	o.PeerSvmName = peerSvmName
}

// WithPeerSvmUUID adds the peerSvmUUID to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithPeerSvmUUID(peerSvmUUID *string) *SvmPeerDeleteCollectionParams {
	o.SetPeerSvmUUID(peerSvmUUID)
	return o
}

// SetPeerSvmUUID adds the peerSvmUuid to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetPeerSvmUUID(peerSvmUUID *string) {
	o.PeerSvmUUID = peerSvmUUID
}

// WithReturnRecords adds the returnRecords to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *SvmPeerDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *SvmPeerDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *SvmPeerDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithState adds the state to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithState(state *string) *SvmPeerDeleteCollectionParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetState(state *string) {
	o.State = state
}

// WithSvmName adds the svmName to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithSvmName(svmName *string) *SvmPeerDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithSvmUUID(svmUUID *string) *SvmPeerDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) WithUUID(uuid *string) *SvmPeerDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the svm peer delete collection params
func (o *SvmPeerDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SvmPeerDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Applications != nil {

		// query param applications
		var qrApplications string

		if o.Applications != nil {
			qrApplications = *o.Applications
		}
		qApplications := qrApplications
		if qApplications != "" {

			if err := r.SetQueryParam("applications", qApplications); err != nil {
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

	if o.PeerClusterName != nil {

		// query param peer.cluster.name
		var qrPeerClusterName string

		if o.PeerClusterName != nil {
			qrPeerClusterName = *o.PeerClusterName
		}
		qPeerClusterName := qrPeerClusterName
		if qPeerClusterName != "" {

			if err := r.SetQueryParam("peer.cluster.name", qPeerClusterName); err != nil {
				return err
			}
		}
	}

	if o.PeerClusterUUID != nil {

		// query param peer.cluster.uuid
		var qrPeerClusterUUID string

		if o.PeerClusterUUID != nil {
			qrPeerClusterUUID = *o.PeerClusterUUID
		}
		qPeerClusterUUID := qrPeerClusterUUID
		if qPeerClusterUUID != "" {

			if err := r.SetQueryParam("peer.cluster.uuid", qPeerClusterUUID); err != nil {
				return err
			}
		}
	}

	if o.PeerSvmName != nil {

		// query param peer.svm.name
		var qrPeerSvmName string

		if o.PeerSvmName != nil {
			qrPeerSvmName = *o.PeerSvmName
		}
		qPeerSvmName := qrPeerSvmName
		if qPeerSvmName != "" {

			if err := r.SetQueryParam("peer.svm.name", qPeerSvmName); err != nil {
				return err
			}
		}
	}

	if o.PeerSvmUUID != nil {

		// query param peer.svm.uuid
		var qrPeerSvmUUID string

		if o.PeerSvmUUID != nil {
			qrPeerSvmUUID = *o.PeerSvmUUID
		}
		qPeerSvmUUID := qrPeerSvmUUID
		if qPeerSvmUUID != "" {

			if err := r.SetQueryParam("peer.svm.uuid", qPeerSvmUUID); err != nil {
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

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
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
