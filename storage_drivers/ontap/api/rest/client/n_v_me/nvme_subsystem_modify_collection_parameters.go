// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NewNvmeSubsystemModifyCollectionParams creates a new NvmeSubsystemModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeSubsystemModifyCollectionParams() *NvmeSubsystemModifyCollectionParams {
	return &NvmeSubsystemModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeSubsystemModifyCollectionParamsWithTimeout creates a new NvmeSubsystemModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewNvmeSubsystemModifyCollectionParamsWithTimeout(timeout time.Duration) *NvmeSubsystemModifyCollectionParams {
	return &NvmeSubsystemModifyCollectionParams{
		timeout: timeout,
	}
}

// NewNvmeSubsystemModifyCollectionParamsWithContext creates a new NvmeSubsystemModifyCollectionParams object
// with the ability to set a context for a request.
func NewNvmeSubsystemModifyCollectionParamsWithContext(ctx context.Context) *NvmeSubsystemModifyCollectionParams {
	return &NvmeSubsystemModifyCollectionParams{
		Context: ctx,
	}
}

// NewNvmeSubsystemModifyCollectionParamsWithHTTPClient creates a new NvmeSubsystemModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeSubsystemModifyCollectionParamsWithHTTPClient(client *http.Client) *NvmeSubsystemModifyCollectionParams {
	return &NvmeSubsystemModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
NvmeSubsystemModifyCollectionParams contains all the parameters to send to the API endpoint

	for the nvme subsystem modify collection operation.

	Typically these are written to a http.Request.
*/
type NvmeSubsystemModifyCollectionParams struct {

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* DeleteOnUnmap.

	   Filter by delete_on_unmap
	*/
	DeleteOnUnmap *bool

	/* HostsDhHmacChapGroupSize.

	   Filter by hosts.dh_hmac_chap.group_size
	*/
	HostsDhHmacChapGroupSize *string

	/* HostsDhHmacChapHashFunction.

	   Filter by hosts.dh_hmac_chap.hash_function
	*/
	HostsDhHmacChapHashFunction *string

	/* HostsDhHmacChapMode.

	   Filter by hosts.dh_hmac_chap.mode
	*/
	HostsDhHmacChapMode *string

	/* HostsNqn.

	   Filter by hosts.nqn
	*/
	HostsNqn *string

	/* HostsPriority.

	   Filter by hosts.priority
	*/
	HostsPriority *string

	/* HostsTLSKeyType.

	   Filter by hosts.tls.key_type
	*/
	HostsTLSKeyType *string

	/* Info.

	   Info specification
	*/
	Info NvmeSubsystemModifyCollectionBody

	/* IoQueueDefaultCount.

	   Filter by io_queue.default.count
	*/
	IoQueueDefaultCount *int64

	/* IoQueueDefaultDepth.

	   Filter by io_queue.default.depth
	*/
	IoQueueDefaultDepth *int64

	/* Name.

	   Filter by name
	*/
	Name *string

	/* OsType.

	   Filter by os_type
	*/
	OsType *string

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

	/* SerialNumber.

	   Filter by serial_number
	*/
	SerialNumber *string

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* SubsystemMapsAnagrpid.

	   Filter by subsystem_maps.anagrpid
	*/
	SubsystemMapsAnagrpid *string

	/* SubsystemMapsNamespaceName.

	   Filter by subsystem_maps.namespace.name
	*/
	SubsystemMapsNamespaceName *string

	/* SubsystemMapsNamespaceUUID.

	   Filter by subsystem_maps.namespace.uuid
	*/
	SubsystemMapsNamespaceUUID *string

	/* SubsystemMapsNsid.

	   Filter by subsystem_maps.nsid
	*/
	SubsystemMapsNsid *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* TargetNqn.

	   Filter by target_nqn
	*/
	TargetNqn *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	/* VendorUuids.

	   Filter by vendor_uuids
	*/
	VendorUuids *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme subsystem modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemModifyCollectionParams) WithDefaults() *NvmeSubsystemModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme subsystem modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := NvmeSubsystemModifyCollectionParams{
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

// WithTimeout adds the timeout to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithTimeout(timeout time.Duration) *NvmeSubsystemModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithContext(ctx context.Context) *NvmeSubsystemModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithHTTPClient(client *http.Client) *NvmeSubsystemModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithComment(comment *string) *NvmeSubsystemModifyCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContinueOnFailure adds the continueOnFailure to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *NvmeSubsystemModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDeleteOnUnmap adds the deleteOnUnmap to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithDeleteOnUnmap(deleteOnUnmap *bool) *NvmeSubsystemModifyCollectionParams {
	o.SetDeleteOnUnmap(deleteOnUnmap)
	return o
}

// SetDeleteOnUnmap adds the deleteOnUnmap to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetDeleteOnUnmap(deleteOnUnmap *bool) {
	o.DeleteOnUnmap = deleteOnUnmap
}

// WithHostsDhHmacChapGroupSize adds the hostsDhHmacChapGroupSize to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithHostsDhHmacChapGroupSize(hostsDhHmacChapGroupSize *string) *NvmeSubsystemModifyCollectionParams {
	o.SetHostsDhHmacChapGroupSize(hostsDhHmacChapGroupSize)
	return o
}

// SetHostsDhHmacChapGroupSize adds the hostsDhHmacChapGroupSize to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetHostsDhHmacChapGroupSize(hostsDhHmacChapGroupSize *string) {
	o.HostsDhHmacChapGroupSize = hostsDhHmacChapGroupSize
}

// WithHostsDhHmacChapHashFunction adds the hostsDhHmacChapHashFunction to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithHostsDhHmacChapHashFunction(hostsDhHmacChapHashFunction *string) *NvmeSubsystemModifyCollectionParams {
	o.SetHostsDhHmacChapHashFunction(hostsDhHmacChapHashFunction)
	return o
}

// SetHostsDhHmacChapHashFunction adds the hostsDhHmacChapHashFunction to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetHostsDhHmacChapHashFunction(hostsDhHmacChapHashFunction *string) {
	o.HostsDhHmacChapHashFunction = hostsDhHmacChapHashFunction
}

// WithHostsDhHmacChapMode adds the hostsDhHmacChapMode to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithHostsDhHmacChapMode(hostsDhHmacChapMode *string) *NvmeSubsystemModifyCollectionParams {
	o.SetHostsDhHmacChapMode(hostsDhHmacChapMode)
	return o
}

// SetHostsDhHmacChapMode adds the hostsDhHmacChapMode to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetHostsDhHmacChapMode(hostsDhHmacChapMode *string) {
	o.HostsDhHmacChapMode = hostsDhHmacChapMode
}

// WithHostsNqn adds the hostsNqn to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithHostsNqn(hostsNqn *string) *NvmeSubsystemModifyCollectionParams {
	o.SetHostsNqn(hostsNqn)
	return o
}

// SetHostsNqn adds the hostsNqn to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetHostsNqn(hostsNqn *string) {
	o.HostsNqn = hostsNqn
}

// WithHostsPriority adds the hostsPriority to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithHostsPriority(hostsPriority *string) *NvmeSubsystemModifyCollectionParams {
	o.SetHostsPriority(hostsPriority)
	return o
}

// SetHostsPriority adds the hostsPriority to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetHostsPriority(hostsPriority *string) {
	o.HostsPriority = hostsPriority
}

// WithHostsTLSKeyType adds the hostsTLSKeyType to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithHostsTLSKeyType(hostsTLSKeyType *string) *NvmeSubsystemModifyCollectionParams {
	o.SetHostsTLSKeyType(hostsTLSKeyType)
	return o
}

// SetHostsTLSKeyType adds the hostsTlsKeyType to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetHostsTLSKeyType(hostsTLSKeyType *string) {
	o.HostsTLSKeyType = hostsTLSKeyType
}

// WithInfo adds the info to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithInfo(info NvmeSubsystemModifyCollectionBody) *NvmeSubsystemModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetInfo(info NvmeSubsystemModifyCollectionBody) {
	o.Info = info
}

// WithIoQueueDefaultCount adds the ioQueueDefaultCount to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithIoQueueDefaultCount(ioQueueDefaultCount *int64) *NvmeSubsystemModifyCollectionParams {
	o.SetIoQueueDefaultCount(ioQueueDefaultCount)
	return o
}

// SetIoQueueDefaultCount adds the ioQueueDefaultCount to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetIoQueueDefaultCount(ioQueueDefaultCount *int64) {
	o.IoQueueDefaultCount = ioQueueDefaultCount
}

// WithIoQueueDefaultDepth adds the ioQueueDefaultDepth to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithIoQueueDefaultDepth(ioQueueDefaultDepth *int64) *NvmeSubsystemModifyCollectionParams {
	o.SetIoQueueDefaultDepth(ioQueueDefaultDepth)
	return o
}

// SetIoQueueDefaultDepth adds the ioQueueDefaultDepth to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetIoQueueDefaultDepth(ioQueueDefaultDepth *int64) {
	o.IoQueueDefaultDepth = ioQueueDefaultDepth
}

// WithName adds the name to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithName(name *string) *NvmeSubsystemModifyCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithOsType adds the osType to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithOsType(osType *string) *NvmeSubsystemModifyCollectionParams {
	o.SetOsType(osType)
	return o
}

// SetOsType adds the osType to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetOsType(osType *string) {
	o.OsType = osType
}

// WithReturnRecords adds the returnRecords to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithReturnRecords(returnRecords *bool) *NvmeSubsystemModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *NvmeSubsystemModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialNumber adds the serialNumber to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithSerialNumber(serialNumber *string) *NvmeSubsystemModifyCollectionParams {
	o.SetSerialNumber(serialNumber)
	return o
}

// SetSerialNumber adds the serialNumber to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetSerialNumber(serialNumber *string) {
	o.SerialNumber = serialNumber
}

// WithSerialRecords adds the serialRecords to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithSerialRecords(serialRecords *bool) *NvmeSubsystemModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSubsystemMapsAnagrpid adds the subsystemMapsAnagrpid to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithSubsystemMapsAnagrpid(subsystemMapsAnagrpid *string) *NvmeSubsystemModifyCollectionParams {
	o.SetSubsystemMapsAnagrpid(subsystemMapsAnagrpid)
	return o
}

// SetSubsystemMapsAnagrpid adds the subsystemMapsAnagrpid to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetSubsystemMapsAnagrpid(subsystemMapsAnagrpid *string) {
	o.SubsystemMapsAnagrpid = subsystemMapsAnagrpid
}

// WithSubsystemMapsNamespaceName adds the subsystemMapsNamespaceName to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithSubsystemMapsNamespaceName(subsystemMapsNamespaceName *string) *NvmeSubsystemModifyCollectionParams {
	o.SetSubsystemMapsNamespaceName(subsystemMapsNamespaceName)
	return o
}

// SetSubsystemMapsNamespaceName adds the subsystemMapsNamespaceName to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetSubsystemMapsNamespaceName(subsystemMapsNamespaceName *string) {
	o.SubsystemMapsNamespaceName = subsystemMapsNamespaceName
}

// WithSubsystemMapsNamespaceUUID adds the subsystemMapsNamespaceUUID to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithSubsystemMapsNamespaceUUID(subsystemMapsNamespaceUUID *string) *NvmeSubsystemModifyCollectionParams {
	o.SetSubsystemMapsNamespaceUUID(subsystemMapsNamespaceUUID)
	return o
}

// SetSubsystemMapsNamespaceUUID adds the subsystemMapsNamespaceUuid to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetSubsystemMapsNamespaceUUID(subsystemMapsNamespaceUUID *string) {
	o.SubsystemMapsNamespaceUUID = subsystemMapsNamespaceUUID
}

// WithSubsystemMapsNsid adds the subsystemMapsNsid to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithSubsystemMapsNsid(subsystemMapsNsid *string) *NvmeSubsystemModifyCollectionParams {
	o.SetSubsystemMapsNsid(subsystemMapsNsid)
	return o
}

// SetSubsystemMapsNsid adds the subsystemMapsNsid to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetSubsystemMapsNsid(subsystemMapsNsid *string) {
	o.SubsystemMapsNsid = subsystemMapsNsid
}

// WithSvmName adds the svmName to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithSvmName(svmName *string) *NvmeSubsystemModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithSvmUUID(svmUUID *string) *NvmeSubsystemModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithTargetNqn adds the targetNqn to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithTargetNqn(targetNqn *string) *NvmeSubsystemModifyCollectionParams {
	o.SetTargetNqn(targetNqn)
	return o
}

// SetTargetNqn adds the targetNqn to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetTargetNqn(targetNqn *string) {
	o.TargetNqn = targetNqn
}

// WithUUID adds the uuid to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithUUID(uuid *string) *NvmeSubsystemModifyCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WithVendorUuids adds the vendorUuids to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) WithVendorUuids(vendorUuids *string) *NvmeSubsystemModifyCollectionParams {
	o.SetVendorUuids(vendorUuids)
	return o
}

// SetVendorUuids adds the vendorUuids to the nvme subsystem modify collection params
func (o *NvmeSubsystemModifyCollectionParams) SetVendorUuids(vendorUuids *string) {
	o.VendorUuids = vendorUuids
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeSubsystemModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DeleteOnUnmap != nil {

		// query param delete_on_unmap
		var qrDeleteOnUnmap bool

		if o.DeleteOnUnmap != nil {
			qrDeleteOnUnmap = *o.DeleteOnUnmap
		}
		qDeleteOnUnmap := swag.FormatBool(qrDeleteOnUnmap)
		if qDeleteOnUnmap != "" {

			if err := r.SetQueryParam("delete_on_unmap", qDeleteOnUnmap); err != nil {
				return err
			}
		}
	}

	if o.HostsDhHmacChapGroupSize != nil {

		// query param hosts.dh_hmac_chap.group_size
		var qrHostsDhHmacChapGroupSize string

		if o.HostsDhHmacChapGroupSize != nil {
			qrHostsDhHmacChapGroupSize = *o.HostsDhHmacChapGroupSize
		}
		qHostsDhHmacChapGroupSize := qrHostsDhHmacChapGroupSize
		if qHostsDhHmacChapGroupSize != "" {

			if err := r.SetQueryParam("hosts.dh_hmac_chap.group_size", qHostsDhHmacChapGroupSize); err != nil {
				return err
			}
		}
	}

	if o.HostsDhHmacChapHashFunction != nil {

		// query param hosts.dh_hmac_chap.hash_function
		var qrHostsDhHmacChapHashFunction string

		if o.HostsDhHmacChapHashFunction != nil {
			qrHostsDhHmacChapHashFunction = *o.HostsDhHmacChapHashFunction
		}
		qHostsDhHmacChapHashFunction := qrHostsDhHmacChapHashFunction
		if qHostsDhHmacChapHashFunction != "" {

			if err := r.SetQueryParam("hosts.dh_hmac_chap.hash_function", qHostsDhHmacChapHashFunction); err != nil {
				return err
			}
		}
	}

	if o.HostsDhHmacChapMode != nil {

		// query param hosts.dh_hmac_chap.mode
		var qrHostsDhHmacChapMode string

		if o.HostsDhHmacChapMode != nil {
			qrHostsDhHmacChapMode = *o.HostsDhHmacChapMode
		}
		qHostsDhHmacChapMode := qrHostsDhHmacChapMode
		if qHostsDhHmacChapMode != "" {

			if err := r.SetQueryParam("hosts.dh_hmac_chap.mode", qHostsDhHmacChapMode); err != nil {
				return err
			}
		}
	}

	if o.HostsNqn != nil {

		// query param hosts.nqn
		var qrHostsNqn string

		if o.HostsNqn != nil {
			qrHostsNqn = *o.HostsNqn
		}
		qHostsNqn := qrHostsNqn
		if qHostsNqn != "" {

			if err := r.SetQueryParam("hosts.nqn", qHostsNqn); err != nil {
				return err
			}
		}
	}

	if o.HostsPriority != nil {

		// query param hosts.priority
		var qrHostsPriority string

		if o.HostsPriority != nil {
			qrHostsPriority = *o.HostsPriority
		}
		qHostsPriority := qrHostsPriority
		if qHostsPriority != "" {

			if err := r.SetQueryParam("hosts.priority", qHostsPriority); err != nil {
				return err
			}
		}
	}

	if o.HostsTLSKeyType != nil {

		// query param hosts.tls.key_type
		var qrHostsTLSKeyType string

		if o.HostsTLSKeyType != nil {
			qrHostsTLSKeyType = *o.HostsTLSKeyType
		}
		qHostsTLSKeyType := qrHostsTLSKeyType
		if qHostsTLSKeyType != "" {

			if err := r.SetQueryParam("hosts.tls.key_type", qHostsTLSKeyType); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.IoQueueDefaultCount != nil {

		// query param io_queue.default.count
		var qrIoQueueDefaultCount int64

		if o.IoQueueDefaultCount != nil {
			qrIoQueueDefaultCount = *o.IoQueueDefaultCount
		}
		qIoQueueDefaultCount := swag.FormatInt64(qrIoQueueDefaultCount)
		if qIoQueueDefaultCount != "" {

			if err := r.SetQueryParam("io_queue.default.count", qIoQueueDefaultCount); err != nil {
				return err
			}
		}
	}

	if o.IoQueueDefaultDepth != nil {

		// query param io_queue.default.depth
		var qrIoQueueDefaultDepth int64

		if o.IoQueueDefaultDepth != nil {
			qrIoQueueDefaultDepth = *o.IoQueueDefaultDepth
		}
		qIoQueueDefaultDepth := swag.FormatInt64(qrIoQueueDefaultDepth)
		if qIoQueueDefaultDepth != "" {

			if err := r.SetQueryParam("io_queue.default.depth", qIoQueueDefaultDepth); err != nil {
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

	if o.OsType != nil {

		// query param os_type
		var qrOsType string

		if o.OsType != nil {
			qrOsType = *o.OsType
		}
		qOsType := qrOsType
		if qOsType != "" {

			if err := r.SetQueryParam("os_type", qOsType); err != nil {
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

	if o.SerialNumber != nil {

		// query param serial_number
		var qrSerialNumber string

		if o.SerialNumber != nil {
			qrSerialNumber = *o.SerialNumber
		}
		qSerialNumber := qrSerialNumber
		if qSerialNumber != "" {

			if err := r.SetQueryParam("serial_number", qSerialNumber); err != nil {
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

	if o.SubsystemMapsAnagrpid != nil {

		// query param subsystem_maps.anagrpid
		var qrSubsystemMapsAnagrpid string

		if o.SubsystemMapsAnagrpid != nil {
			qrSubsystemMapsAnagrpid = *o.SubsystemMapsAnagrpid
		}
		qSubsystemMapsAnagrpid := qrSubsystemMapsAnagrpid
		if qSubsystemMapsAnagrpid != "" {

			if err := r.SetQueryParam("subsystem_maps.anagrpid", qSubsystemMapsAnagrpid); err != nil {
				return err
			}
		}
	}

	if o.SubsystemMapsNamespaceName != nil {

		// query param subsystem_maps.namespace.name
		var qrSubsystemMapsNamespaceName string

		if o.SubsystemMapsNamespaceName != nil {
			qrSubsystemMapsNamespaceName = *o.SubsystemMapsNamespaceName
		}
		qSubsystemMapsNamespaceName := qrSubsystemMapsNamespaceName
		if qSubsystemMapsNamespaceName != "" {

			if err := r.SetQueryParam("subsystem_maps.namespace.name", qSubsystemMapsNamespaceName); err != nil {
				return err
			}
		}
	}

	if o.SubsystemMapsNamespaceUUID != nil {

		// query param subsystem_maps.namespace.uuid
		var qrSubsystemMapsNamespaceUUID string

		if o.SubsystemMapsNamespaceUUID != nil {
			qrSubsystemMapsNamespaceUUID = *o.SubsystemMapsNamespaceUUID
		}
		qSubsystemMapsNamespaceUUID := qrSubsystemMapsNamespaceUUID
		if qSubsystemMapsNamespaceUUID != "" {

			if err := r.SetQueryParam("subsystem_maps.namespace.uuid", qSubsystemMapsNamespaceUUID); err != nil {
				return err
			}
		}
	}

	if o.SubsystemMapsNsid != nil {

		// query param subsystem_maps.nsid
		var qrSubsystemMapsNsid string

		if o.SubsystemMapsNsid != nil {
			qrSubsystemMapsNsid = *o.SubsystemMapsNsid
		}
		qSubsystemMapsNsid := qrSubsystemMapsNsid
		if qSubsystemMapsNsid != "" {

			if err := r.SetQueryParam("subsystem_maps.nsid", qSubsystemMapsNsid); err != nil {
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

	if o.TargetNqn != nil {

		// query param target_nqn
		var qrTargetNqn string

		if o.TargetNqn != nil {
			qrTargetNqn = *o.TargetNqn
		}
		qTargetNqn := qrTargetNqn
		if qTargetNqn != "" {

			if err := r.SetQueryParam("target_nqn", qTargetNqn); err != nil {
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

	if o.VendorUuids != nil {

		// query param vendor_uuids
		var qrVendorUuids string

		if o.VendorUuids != nil {
			qrVendorUuids = *o.VendorUuids
		}
		qVendorUuids := qrVendorUuids
		if qVendorUuids != "" {

			if err := r.SetQueryParam("vendor_uuids", qVendorUuids); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
