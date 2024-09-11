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

// NewSvmPeerPermissionModifyCollectionParams creates a new SvmPeerPermissionModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmPeerPermissionModifyCollectionParams() *SvmPeerPermissionModifyCollectionParams {
	return &SvmPeerPermissionModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmPeerPermissionModifyCollectionParamsWithTimeout creates a new SvmPeerPermissionModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewSvmPeerPermissionModifyCollectionParamsWithTimeout(timeout time.Duration) *SvmPeerPermissionModifyCollectionParams {
	return &SvmPeerPermissionModifyCollectionParams{
		timeout: timeout,
	}
}

// NewSvmPeerPermissionModifyCollectionParamsWithContext creates a new SvmPeerPermissionModifyCollectionParams object
// with the ability to set a context for a request.
func NewSvmPeerPermissionModifyCollectionParamsWithContext(ctx context.Context) *SvmPeerPermissionModifyCollectionParams {
	return &SvmPeerPermissionModifyCollectionParams{
		Context: ctx,
	}
}

// NewSvmPeerPermissionModifyCollectionParamsWithHTTPClient creates a new SvmPeerPermissionModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmPeerPermissionModifyCollectionParamsWithHTTPClient(client *http.Client) *SvmPeerPermissionModifyCollectionParams {
	return &SvmPeerPermissionModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
SvmPeerPermissionModifyCollectionParams contains all the parameters to send to the API endpoint

	for the svm peer permission modify collection operation.

	Typically these are written to a http.Request.
*/
type SvmPeerPermissionModifyCollectionParams struct {

	/* Applications.

	   Filter by applications
	*/
	Applications *string

	/* ClusterPeerName.

	   Filter by cluster_peer.name
	*/
	ClusterPeerName *string

	/* ClusterPeerUUID.

	   Filter by cluster_peer.uuid
	*/
	ClusterPeerUUID *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info SvmPeerPermissionModifyCollectionBody

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

// WithDefaults hydrates default values in the svm peer permission modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerPermissionModifyCollectionParams) WithDefaults() *SvmPeerPermissionModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm peer permission modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerPermissionModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SvmPeerPermissionModifyCollectionParams{
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

// WithTimeout adds the timeout to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithTimeout(timeout time.Duration) *SvmPeerPermissionModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithContext(ctx context.Context) *SvmPeerPermissionModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithHTTPClient(client *http.Client) *SvmPeerPermissionModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplications adds the applications to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithApplications(applications *string) *SvmPeerPermissionModifyCollectionParams {
	o.SetApplications(applications)
	return o
}

// SetApplications adds the applications to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetApplications(applications *string) {
	o.Applications = applications
}

// WithClusterPeerName adds the clusterPeerName to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithClusterPeerName(clusterPeerName *string) *SvmPeerPermissionModifyCollectionParams {
	o.SetClusterPeerName(clusterPeerName)
	return o
}

// SetClusterPeerName adds the clusterPeerName to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetClusterPeerName(clusterPeerName *string) {
	o.ClusterPeerName = clusterPeerName
}

// WithClusterPeerUUID adds the clusterPeerUUID to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithClusterPeerUUID(clusterPeerUUID *string) *SvmPeerPermissionModifyCollectionParams {
	o.SetClusterPeerUUID(clusterPeerUUID)
	return o
}

// SetClusterPeerUUID adds the clusterPeerUuid to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetClusterPeerUUID(clusterPeerUUID *string) {
	o.ClusterPeerUUID = clusterPeerUUID
}

// WithContinueOnFailure adds the continueOnFailure to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SvmPeerPermissionModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithInfo(info SvmPeerPermissionModifyCollectionBody) *SvmPeerPermissionModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetInfo(info SvmPeerPermissionModifyCollectionBody) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithReturnRecords(returnRecords *bool) *SvmPeerPermissionModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *SvmPeerPermissionModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithSerialRecords(serialRecords *bool) *SvmPeerPermissionModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithSvmName(svmName *string) *SvmPeerPermissionModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) WithSvmUUID(svmUUID *string) *SvmPeerPermissionModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the svm peer permission modify collection params
func (o *SvmPeerPermissionModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SvmPeerPermissionModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ClusterPeerName != nil {

		// query param cluster_peer.name
		var qrClusterPeerName string

		if o.ClusterPeerName != nil {
			qrClusterPeerName = *o.ClusterPeerName
		}
		qClusterPeerName := qrClusterPeerName
		if qClusterPeerName != "" {

			if err := r.SetQueryParam("cluster_peer.name", qClusterPeerName); err != nil {
				return err
			}
		}
	}

	if o.ClusterPeerUUID != nil {

		// query param cluster_peer.uuid
		var qrClusterPeerUUID string

		if o.ClusterPeerUUID != nil {
			qrClusterPeerUUID = *o.ClusterPeerUUID
		}
		qClusterPeerUUID := qrClusterPeerUUID
		if qClusterPeerUUID != "" {

			if err := r.SetQueryParam("cluster_peer.uuid", qClusterPeerUUID); err != nil {
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
