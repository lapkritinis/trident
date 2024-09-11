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

// NewQosPolicyModifyCollectionParams creates a new QosPolicyModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQosPolicyModifyCollectionParams() *QosPolicyModifyCollectionParams {
	return &QosPolicyModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQosPolicyModifyCollectionParamsWithTimeout creates a new QosPolicyModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewQosPolicyModifyCollectionParamsWithTimeout(timeout time.Duration) *QosPolicyModifyCollectionParams {
	return &QosPolicyModifyCollectionParams{
		timeout: timeout,
	}
}

// NewQosPolicyModifyCollectionParamsWithContext creates a new QosPolicyModifyCollectionParams object
// with the ability to set a context for a request.
func NewQosPolicyModifyCollectionParamsWithContext(ctx context.Context) *QosPolicyModifyCollectionParams {
	return &QosPolicyModifyCollectionParams{
		Context: ctx,
	}
}

// NewQosPolicyModifyCollectionParamsWithHTTPClient creates a new QosPolicyModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewQosPolicyModifyCollectionParamsWithHTTPClient(client *http.Client) *QosPolicyModifyCollectionParams {
	return &QosPolicyModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
QosPolicyModifyCollectionParams contains all the parameters to send to the API endpoint

	for the qos policy modify collection operation.

	Typically these are written to a http.Request.
*/
type QosPolicyModifyCollectionParams struct {

	/* AdaptiveAbsoluteMinIops.

	   Filter by adaptive.absolute_min_iops
	*/
	AdaptiveAbsoluteMinIops *int64

	/* AdaptiveBlockSize.

	   Filter by adaptive.block_size
	*/
	AdaptiveBlockSize *string

	/* AdaptiveExpectedIops.

	   Filter by adaptive.expected_iops
	*/
	AdaptiveExpectedIops *int64

	/* AdaptiveExpectedIopsAllocation.

	   Filter by adaptive.expected_iops_allocation
	*/
	AdaptiveExpectedIopsAllocation *string

	/* AdaptivePeakIops.

	   Filter by adaptive.peak_iops
	*/
	AdaptivePeakIops *int64

	/* AdaptivePeakIopsAllocation.

	   Filter by adaptive.peak_iops_allocation
	*/
	AdaptivePeakIopsAllocation *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* FixedCapacityShared.

	   Filter by fixed.capacity_shared
	*/
	FixedCapacityShared *bool

	/* FixedMaxThroughputIops.

	   Filter by fixed.max_throughput_iops
	*/
	FixedMaxThroughputIops *int64

	/* FixedMaxThroughputMbps.

	   Filter by fixed.max_throughput_mbps
	*/
	FixedMaxThroughputMbps *int64

	/* FixedMinThroughputIops.

	   Filter by fixed.min_throughput_iops
	*/
	FixedMinThroughputIops *int64

	/* FixedMinThroughputMbps.

	   Filter by fixed.min_throughput_mbps
	*/
	FixedMinThroughputMbps *int64

	/* Info.

	   Info specification
	*/
	Info QosPolicyModifyCollectionBody

	/* Name.

	   Filter by name
	*/
	Name *string

	/* ObjectCount.

	   Filter by object_count
	*/
	ObjectCount *int64

	/* Pgid.

	   Filter by pgid
	*/
	Pgid *int64

	/* PolicyClass.

	   Filter by policy_class
	*/
	PolicyClass *string

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

	/* Scope.

	   Filter by scope
	*/
	Scope *string

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

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the qos policy modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QosPolicyModifyCollectionParams) WithDefaults() *QosPolicyModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the qos policy modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QosPolicyModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := QosPolicyModifyCollectionParams{
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

// WithTimeout adds the timeout to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithTimeout(timeout time.Duration) *QosPolicyModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithContext(ctx context.Context) *QosPolicyModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithHTTPClient(client *http.Client) *QosPolicyModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdaptiveAbsoluteMinIops adds the adaptiveAbsoluteMinIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithAdaptiveAbsoluteMinIops(adaptiveAbsoluteMinIops *int64) *QosPolicyModifyCollectionParams {
	o.SetAdaptiveAbsoluteMinIops(adaptiveAbsoluteMinIops)
	return o
}

// SetAdaptiveAbsoluteMinIops adds the adaptiveAbsoluteMinIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetAdaptiveAbsoluteMinIops(adaptiveAbsoluteMinIops *int64) {
	o.AdaptiveAbsoluteMinIops = adaptiveAbsoluteMinIops
}

// WithAdaptiveBlockSize adds the adaptiveBlockSize to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithAdaptiveBlockSize(adaptiveBlockSize *string) *QosPolicyModifyCollectionParams {
	o.SetAdaptiveBlockSize(adaptiveBlockSize)
	return o
}

// SetAdaptiveBlockSize adds the adaptiveBlockSize to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetAdaptiveBlockSize(adaptiveBlockSize *string) {
	o.AdaptiveBlockSize = adaptiveBlockSize
}

// WithAdaptiveExpectedIops adds the adaptiveExpectedIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithAdaptiveExpectedIops(adaptiveExpectedIops *int64) *QosPolicyModifyCollectionParams {
	o.SetAdaptiveExpectedIops(adaptiveExpectedIops)
	return o
}

// SetAdaptiveExpectedIops adds the adaptiveExpectedIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetAdaptiveExpectedIops(adaptiveExpectedIops *int64) {
	o.AdaptiveExpectedIops = adaptiveExpectedIops
}

// WithAdaptiveExpectedIopsAllocation adds the adaptiveExpectedIopsAllocation to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithAdaptiveExpectedIopsAllocation(adaptiveExpectedIopsAllocation *string) *QosPolicyModifyCollectionParams {
	o.SetAdaptiveExpectedIopsAllocation(adaptiveExpectedIopsAllocation)
	return o
}

// SetAdaptiveExpectedIopsAllocation adds the adaptiveExpectedIopsAllocation to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetAdaptiveExpectedIopsAllocation(adaptiveExpectedIopsAllocation *string) {
	o.AdaptiveExpectedIopsAllocation = adaptiveExpectedIopsAllocation
}

// WithAdaptivePeakIops adds the adaptivePeakIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithAdaptivePeakIops(adaptivePeakIops *int64) *QosPolicyModifyCollectionParams {
	o.SetAdaptivePeakIops(adaptivePeakIops)
	return o
}

// SetAdaptivePeakIops adds the adaptivePeakIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetAdaptivePeakIops(adaptivePeakIops *int64) {
	o.AdaptivePeakIops = adaptivePeakIops
}

// WithAdaptivePeakIopsAllocation adds the adaptivePeakIopsAllocation to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithAdaptivePeakIopsAllocation(adaptivePeakIopsAllocation *string) *QosPolicyModifyCollectionParams {
	o.SetAdaptivePeakIopsAllocation(adaptivePeakIopsAllocation)
	return o
}

// SetAdaptivePeakIopsAllocation adds the adaptivePeakIopsAllocation to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetAdaptivePeakIopsAllocation(adaptivePeakIopsAllocation *string) {
	o.AdaptivePeakIopsAllocation = adaptivePeakIopsAllocation
}

// WithContinueOnFailure adds the continueOnFailure to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *QosPolicyModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithFixedCapacityShared adds the fixedCapacityShared to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithFixedCapacityShared(fixedCapacityShared *bool) *QosPolicyModifyCollectionParams {
	o.SetFixedCapacityShared(fixedCapacityShared)
	return o
}

// SetFixedCapacityShared adds the fixedCapacityShared to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetFixedCapacityShared(fixedCapacityShared *bool) {
	o.FixedCapacityShared = fixedCapacityShared
}

// WithFixedMaxThroughputIops adds the fixedMaxThroughputIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithFixedMaxThroughputIops(fixedMaxThroughputIops *int64) *QosPolicyModifyCollectionParams {
	o.SetFixedMaxThroughputIops(fixedMaxThroughputIops)
	return o
}

// SetFixedMaxThroughputIops adds the fixedMaxThroughputIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetFixedMaxThroughputIops(fixedMaxThroughputIops *int64) {
	o.FixedMaxThroughputIops = fixedMaxThroughputIops
}

// WithFixedMaxThroughputMbps adds the fixedMaxThroughputMbps to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithFixedMaxThroughputMbps(fixedMaxThroughputMbps *int64) *QosPolicyModifyCollectionParams {
	o.SetFixedMaxThroughputMbps(fixedMaxThroughputMbps)
	return o
}

// SetFixedMaxThroughputMbps adds the fixedMaxThroughputMbps to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetFixedMaxThroughputMbps(fixedMaxThroughputMbps *int64) {
	o.FixedMaxThroughputMbps = fixedMaxThroughputMbps
}

// WithFixedMinThroughputIops adds the fixedMinThroughputIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithFixedMinThroughputIops(fixedMinThroughputIops *int64) *QosPolicyModifyCollectionParams {
	o.SetFixedMinThroughputIops(fixedMinThroughputIops)
	return o
}

// SetFixedMinThroughputIops adds the fixedMinThroughputIops to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetFixedMinThroughputIops(fixedMinThroughputIops *int64) {
	o.FixedMinThroughputIops = fixedMinThroughputIops
}

// WithFixedMinThroughputMbps adds the fixedMinThroughputMbps to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithFixedMinThroughputMbps(fixedMinThroughputMbps *int64) *QosPolicyModifyCollectionParams {
	o.SetFixedMinThroughputMbps(fixedMinThroughputMbps)
	return o
}

// SetFixedMinThroughputMbps adds the fixedMinThroughputMbps to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetFixedMinThroughputMbps(fixedMinThroughputMbps *int64) {
	o.FixedMinThroughputMbps = fixedMinThroughputMbps
}

// WithInfo adds the info to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithInfo(info QosPolicyModifyCollectionBody) *QosPolicyModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetInfo(info QosPolicyModifyCollectionBody) {
	o.Info = info
}

// WithName adds the name to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithName(name *string) *QosPolicyModifyCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithObjectCount adds the objectCount to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithObjectCount(objectCount *int64) *QosPolicyModifyCollectionParams {
	o.SetObjectCount(objectCount)
	return o
}

// SetObjectCount adds the objectCount to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetObjectCount(objectCount *int64) {
	o.ObjectCount = objectCount
}

// WithPgid adds the pgid to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithPgid(pgid *int64) *QosPolicyModifyCollectionParams {
	o.SetPgid(pgid)
	return o
}

// SetPgid adds the pgid to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetPgid(pgid *int64) {
	o.Pgid = pgid
}

// WithPolicyClass adds the policyClass to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithPolicyClass(policyClass *string) *QosPolicyModifyCollectionParams {
	o.SetPolicyClass(policyClass)
	return o
}

// SetPolicyClass adds the policyClass to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetPolicyClass(policyClass *string) {
	o.PolicyClass = policyClass
}

// WithReturnRecords adds the returnRecords to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithReturnRecords(returnRecords *bool) *QosPolicyModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *QosPolicyModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScope adds the scope to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithScope(scope *string) *QosPolicyModifyCollectionParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSerialRecords adds the serialRecords to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithSerialRecords(serialRecords *bool) *QosPolicyModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithSvmName(svmName *string) *QosPolicyModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithSvmUUID(svmUUID *string) *QosPolicyModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) WithUUID(uuid *string) *QosPolicyModifyCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the qos policy modify collection params
func (o *QosPolicyModifyCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *QosPolicyModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdaptiveAbsoluteMinIops != nil {

		// query param adaptive.absolute_min_iops
		var qrAdaptiveAbsoluteMinIops int64

		if o.AdaptiveAbsoluteMinIops != nil {
			qrAdaptiveAbsoluteMinIops = *o.AdaptiveAbsoluteMinIops
		}
		qAdaptiveAbsoluteMinIops := swag.FormatInt64(qrAdaptiveAbsoluteMinIops)
		if qAdaptiveAbsoluteMinIops != "" {

			if err := r.SetQueryParam("adaptive.absolute_min_iops", qAdaptiveAbsoluteMinIops); err != nil {
				return err
			}
		}
	}

	if o.AdaptiveBlockSize != nil {

		// query param adaptive.block_size
		var qrAdaptiveBlockSize string

		if o.AdaptiveBlockSize != nil {
			qrAdaptiveBlockSize = *o.AdaptiveBlockSize
		}
		qAdaptiveBlockSize := qrAdaptiveBlockSize
		if qAdaptiveBlockSize != "" {

			if err := r.SetQueryParam("adaptive.block_size", qAdaptiveBlockSize); err != nil {
				return err
			}
		}
	}

	if o.AdaptiveExpectedIops != nil {

		// query param adaptive.expected_iops
		var qrAdaptiveExpectedIops int64

		if o.AdaptiveExpectedIops != nil {
			qrAdaptiveExpectedIops = *o.AdaptiveExpectedIops
		}
		qAdaptiveExpectedIops := swag.FormatInt64(qrAdaptiveExpectedIops)
		if qAdaptiveExpectedIops != "" {

			if err := r.SetQueryParam("adaptive.expected_iops", qAdaptiveExpectedIops); err != nil {
				return err
			}
		}
	}

	if o.AdaptiveExpectedIopsAllocation != nil {

		// query param adaptive.expected_iops_allocation
		var qrAdaptiveExpectedIopsAllocation string

		if o.AdaptiveExpectedIopsAllocation != nil {
			qrAdaptiveExpectedIopsAllocation = *o.AdaptiveExpectedIopsAllocation
		}
		qAdaptiveExpectedIopsAllocation := qrAdaptiveExpectedIopsAllocation
		if qAdaptiveExpectedIopsAllocation != "" {

			if err := r.SetQueryParam("adaptive.expected_iops_allocation", qAdaptiveExpectedIopsAllocation); err != nil {
				return err
			}
		}
	}

	if o.AdaptivePeakIops != nil {

		// query param adaptive.peak_iops
		var qrAdaptivePeakIops int64

		if o.AdaptivePeakIops != nil {
			qrAdaptivePeakIops = *o.AdaptivePeakIops
		}
		qAdaptivePeakIops := swag.FormatInt64(qrAdaptivePeakIops)
		if qAdaptivePeakIops != "" {

			if err := r.SetQueryParam("adaptive.peak_iops", qAdaptivePeakIops); err != nil {
				return err
			}
		}
	}

	if o.AdaptivePeakIopsAllocation != nil {

		// query param adaptive.peak_iops_allocation
		var qrAdaptivePeakIopsAllocation string

		if o.AdaptivePeakIopsAllocation != nil {
			qrAdaptivePeakIopsAllocation = *o.AdaptivePeakIopsAllocation
		}
		qAdaptivePeakIopsAllocation := qrAdaptivePeakIopsAllocation
		if qAdaptivePeakIopsAllocation != "" {

			if err := r.SetQueryParam("adaptive.peak_iops_allocation", qAdaptivePeakIopsAllocation); err != nil {
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

	if o.FixedCapacityShared != nil {

		// query param fixed.capacity_shared
		var qrFixedCapacityShared bool

		if o.FixedCapacityShared != nil {
			qrFixedCapacityShared = *o.FixedCapacityShared
		}
		qFixedCapacityShared := swag.FormatBool(qrFixedCapacityShared)
		if qFixedCapacityShared != "" {

			if err := r.SetQueryParam("fixed.capacity_shared", qFixedCapacityShared); err != nil {
				return err
			}
		}
	}

	if o.FixedMaxThroughputIops != nil {

		// query param fixed.max_throughput_iops
		var qrFixedMaxThroughputIops int64

		if o.FixedMaxThroughputIops != nil {
			qrFixedMaxThroughputIops = *o.FixedMaxThroughputIops
		}
		qFixedMaxThroughputIops := swag.FormatInt64(qrFixedMaxThroughputIops)
		if qFixedMaxThroughputIops != "" {

			if err := r.SetQueryParam("fixed.max_throughput_iops", qFixedMaxThroughputIops); err != nil {
				return err
			}
		}
	}

	if o.FixedMaxThroughputMbps != nil {

		// query param fixed.max_throughput_mbps
		var qrFixedMaxThroughputMbps int64

		if o.FixedMaxThroughputMbps != nil {
			qrFixedMaxThroughputMbps = *o.FixedMaxThroughputMbps
		}
		qFixedMaxThroughputMbps := swag.FormatInt64(qrFixedMaxThroughputMbps)
		if qFixedMaxThroughputMbps != "" {

			if err := r.SetQueryParam("fixed.max_throughput_mbps", qFixedMaxThroughputMbps); err != nil {
				return err
			}
		}
	}

	if o.FixedMinThroughputIops != nil {

		// query param fixed.min_throughput_iops
		var qrFixedMinThroughputIops int64

		if o.FixedMinThroughputIops != nil {
			qrFixedMinThroughputIops = *o.FixedMinThroughputIops
		}
		qFixedMinThroughputIops := swag.FormatInt64(qrFixedMinThroughputIops)
		if qFixedMinThroughputIops != "" {

			if err := r.SetQueryParam("fixed.min_throughput_iops", qFixedMinThroughputIops); err != nil {
				return err
			}
		}
	}

	if o.FixedMinThroughputMbps != nil {

		// query param fixed.min_throughput_mbps
		var qrFixedMinThroughputMbps int64

		if o.FixedMinThroughputMbps != nil {
			qrFixedMinThroughputMbps = *o.FixedMinThroughputMbps
		}
		qFixedMinThroughputMbps := swag.FormatInt64(qrFixedMinThroughputMbps)
		if qFixedMinThroughputMbps != "" {

			if err := r.SetQueryParam("fixed.min_throughput_mbps", qFixedMinThroughputMbps); err != nil {
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

	if o.ObjectCount != nil {

		// query param object_count
		var qrObjectCount int64

		if o.ObjectCount != nil {
			qrObjectCount = *o.ObjectCount
		}
		qObjectCount := swag.FormatInt64(qrObjectCount)
		if qObjectCount != "" {

			if err := r.SetQueryParam("object_count", qObjectCount); err != nil {
				return err
			}
		}
	}

	if o.Pgid != nil {

		// query param pgid
		var qrPgid int64

		if o.Pgid != nil {
			qrPgid = *o.Pgid
		}
		qPgid := swag.FormatInt64(qrPgid)
		if qPgid != "" {

			if err := r.SetQueryParam("pgid", qPgid); err != nil {
				return err
			}
		}
	}

	if o.PolicyClass != nil {

		// query param policy_class
		var qrPolicyClass string

		if o.PolicyClass != nil {
			qrPolicyClass = *o.PolicyClass
		}
		qPolicyClass := qrPolicyClass
		if qPolicyClass != "" {

			if err := r.SetQueryParam("policy_class", qPolicyClass); err != nil {
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

	if o.Scope != nil {

		// query param scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
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
