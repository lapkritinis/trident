// Code generated by go-swagger; DO NOT EDIT.

package ndmp

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

// NewNdmpNodeSessionDeleteCollectionParams creates a new NdmpNodeSessionDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNdmpNodeSessionDeleteCollectionParams() *NdmpNodeSessionDeleteCollectionParams {
	return &NdmpNodeSessionDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNdmpNodeSessionDeleteCollectionParamsWithTimeout creates a new NdmpNodeSessionDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewNdmpNodeSessionDeleteCollectionParamsWithTimeout(timeout time.Duration) *NdmpNodeSessionDeleteCollectionParams {
	return &NdmpNodeSessionDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewNdmpNodeSessionDeleteCollectionParamsWithContext creates a new NdmpNodeSessionDeleteCollectionParams object
// with the ability to set a context for a request.
func NewNdmpNodeSessionDeleteCollectionParamsWithContext(ctx context.Context) *NdmpNodeSessionDeleteCollectionParams {
	return &NdmpNodeSessionDeleteCollectionParams{
		Context: ctx,
	}
}

// NewNdmpNodeSessionDeleteCollectionParamsWithHTTPClient creates a new NdmpNodeSessionDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewNdmpNodeSessionDeleteCollectionParamsWithHTTPClient(client *http.Client) *NdmpNodeSessionDeleteCollectionParams {
	return &NdmpNodeSessionDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
NdmpNodeSessionDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the ndmp node session delete collection operation.

	Typically these are written to a http.Request.
*/
type NdmpNodeSessionDeleteCollectionParams struct {

	/* BackupEngine.

	   Filter by backup_engine
	*/
	BackupEngine *string

	/* ClientAddress.

	   Filter by client_address
	*/
	ClientAddress *string

	/* ClientPort.

	   Filter by client_port
	*/
	ClientPort *int64

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* DataBytesProcessed.

	   Filter by data.bytes_processed
	*/
	DataBytesProcessed *int64

	/* DataConnectionAddress.

	   Filter by data.connection.address
	*/
	DataConnectionAddress *string

	/* DataConnectionPort.

	   Filter by data.connection.port
	*/
	DataConnectionPort *int64

	/* DataConnectionType.

	   Filter by data.connection.type
	*/
	DataConnectionType *string

	/* DataOperation.

	   Filter by data.operation
	*/
	DataOperation *string

	/* DataReason.

	   Filter by data.reason
	*/
	DataReason *string

	/* DataState.

	   Filter by data.state
	*/
	DataState *string

	/* DataPath.

	   Filter by data_path
	*/
	DataPath *string

	/* ID.

	   Filter by id
	*/
	ID *string

	/* Info.

	   Info specification
	*/
	Info NdmpNodeSessionDeleteCollectionBody

	/* MoverBytesMoved.

	   Filter by mover.bytes_moved
	*/
	MoverBytesMoved *int64

	/* MoverConnectionAddress.

	   Filter by mover.connection.address
	*/
	MoverConnectionAddress *string

	/* MoverConnectionPort.

	   Filter by mover.connection.port
	*/
	MoverConnectionPort *int64

	/* MoverConnectionType.

	   Filter by mover.connection.type
	*/
	MoverConnectionType *string

	/* MoverMode.

	   Filter by mover.mode
	*/
	MoverMode *string

	/* MoverReason.

	   Filter by mover.reason
	*/
	MoverReason *string

	/* MoverState.

	   Filter by mover.state
	*/
	MoverState *string

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

	/* ScsiDeviceID.

	   Filter by scsi.device_id
	*/
	ScsiDeviceID *string

	/* ScsiHostAdapter.

	   Filter by scsi.host_adapter
	*/
	ScsiHostAdapter *int64

	/* ScsiLunID.

	   Filter by scsi.lun_id
	*/
	ScsiLunID *int64

	/* ScsiTargetID.

	   Filter by scsi.target_id
	*/
	ScsiTargetID *int64

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* SourceAddress.

	   Filter by source_address
	*/
	SourceAddress *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* TapeDevice.

	   Filter by tape_device
	*/
	TapeDevice *string

	/* TapeMode.

	   Filter by tape_mode
	*/
	TapeMode *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ndmp node session delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeSessionDeleteCollectionParams) WithDefaults() *NdmpNodeSessionDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ndmp node session delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeSessionDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := NdmpNodeSessionDeleteCollectionParams{
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

// WithTimeout adds the timeout to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithTimeout(timeout time.Duration) *NdmpNodeSessionDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithContext(ctx context.Context) *NdmpNodeSessionDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithHTTPClient(client *http.Client) *NdmpNodeSessionDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupEngine adds the backupEngine to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithBackupEngine(backupEngine *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetBackupEngine(backupEngine)
	return o
}

// SetBackupEngine adds the backupEngine to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetBackupEngine(backupEngine *string) {
	o.BackupEngine = backupEngine
}

// WithClientAddress adds the clientAddress to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithClientAddress(clientAddress *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetClientAddress(clientAddress)
	return o
}

// SetClientAddress adds the clientAddress to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetClientAddress(clientAddress *string) {
	o.ClientAddress = clientAddress
}

// WithClientPort adds the clientPort to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithClientPort(clientPort *int64) *NdmpNodeSessionDeleteCollectionParams {
	o.SetClientPort(clientPort)
	return o
}

// SetClientPort adds the clientPort to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetClientPort(clientPort *int64) {
	o.ClientPort = clientPort
}

// WithContinueOnFailure adds the continueOnFailure to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *NdmpNodeSessionDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDataBytesProcessed adds the dataBytesProcessed to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithDataBytesProcessed(dataBytesProcessed *int64) *NdmpNodeSessionDeleteCollectionParams {
	o.SetDataBytesProcessed(dataBytesProcessed)
	return o
}

// SetDataBytesProcessed adds the dataBytesProcessed to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetDataBytesProcessed(dataBytesProcessed *int64) {
	o.DataBytesProcessed = dataBytesProcessed
}

// WithDataConnectionAddress adds the dataConnectionAddress to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithDataConnectionAddress(dataConnectionAddress *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetDataConnectionAddress(dataConnectionAddress)
	return o
}

// SetDataConnectionAddress adds the dataConnectionAddress to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetDataConnectionAddress(dataConnectionAddress *string) {
	o.DataConnectionAddress = dataConnectionAddress
}

// WithDataConnectionPort adds the dataConnectionPort to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithDataConnectionPort(dataConnectionPort *int64) *NdmpNodeSessionDeleteCollectionParams {
	o.SetDataConnectionPort(dataConnectionPort)
	return o
}

// SetDataConnectionPort adds the dataConnectionPort to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetDataConnectionPort(dataConnectionPort *int64) {
	o.DataConnectionPort = dataConnectionPort
}

// WithDataConnectionType adds the dataConnectionType to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithDataConnectionType(dataConnectionType *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetDataConnectionType(dataConnectionType)
	return o
}

// SetDataConnectionType adds the dataConnectionType to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetDataConnectionType(dataConnectionType *string) {
	o.DataConnectionType = dataConnectionType
}

// WithDataOperation adds the dataOperation to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithDataOperation(dataOperation *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetDataOperation(dataOperation)
	return o
}

// SetDataOperation adds the dataOperation to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetDataOperation(dataOperation *string) {
	o.DataOperation = dataOperation
}

// WithDataReason adds the dataReason to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithDataReason(dataReason *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetDataReason(dataReason)
	return o
}

// SetDataReason adds the dataReason to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetDataReason(dataReason *string) {
	o.DataReason = dataReason
}

// WithDataState adds the dataState to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithDataState(dataState *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetDataState(dataState)
	return o
}

// SetDataState adds the dataState to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetDataState(dataState *string) {
	o.DataState = dataState
}

// WithDataPath adds the dataPath to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithDataPath(dataPath *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetDataPath(dataPath)
	return o
}

// SetDataPath adds the dataPath to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetDataPath(dataPath *string) {
	o.DataPath = dataPath
}

// WithID adds the id to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithID(id *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetID(id *string) {
	o.ID = id
}

// WithInfo adds the info to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithInfo(info NdmpNodeSessionDeleteCollectionBody) *NdmpNodeSessionDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetInfo(info NdmpNodeSessionDeleteCollectionBody) {
	o.Info = info
}

// WithMoverBytesMoved adds the moverBytesMoved to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithMoverBytesMoved(moverBytesMoved *int64) *NdmpNodeSessionDeleteCollectionParams {
	o.SetMoverBytesMoved(moverBytesMoved)
	return o
}

// SetMoverBytesMoved adds the moverBytesMoved to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetMoverBytesMoved(moverBytesMoved *int64) {
	o.MoverBytesMoved = moverBytesMoved
}

// WithMoverConnectionAddress adds the moverConnectionAddress to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithMoverConnectionAddress(moverConnectionAddress *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetMoverConnectionAddress(moverConnectionAddress)
	return o
}

// SetMoverConnectionAddress adds the moverConnectionAddress to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetMoverConnectionAddress(moverConnectionAddress *string) {
	o.MoverConnectionAddress = moverConnectionAddress
}

// WithMoverConnectionPort adds the moverConnectionPort to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithMoverConnectionPort(moverConnectionPort *int64) *NdmpNodeSessionDeleteCollectionParams {
	o.SetMoverConnectionPort(moverConnectionPort)
	return o
}

// SetMoverConnectionPort adds the moverConnectionPort to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetMoverConnectionPort(moverConnectionPort *int64) {
	o.MoverConnectionPort = moverConnectionPort
}

// WithMoverConnectionType adds the moverConnectionType to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithMoverConnectionType(moverConnectionType *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetMoverConnectionType(moverConnectionType)
	return o
}

// SetMoverConnectionType adds the moverConnectionType to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetMoverConnectionType(moverConnectionType *string) {
	o.MoverConnectionType = moverConnectionType
}

// WithMoverMode adds the moverMode to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithMoverMode(moverMode *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetMoverMode(moverMode)
	return o
}

// SetMoverMode adds the moverMode to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetMoverMode(moverMode *string) {
	o.MoverMode = moverMode
}

// WithMoverReason adds the moverReason to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithMoverReason(moverReason *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetMoverReason(moverReason)
	return o
}

// SetMoverReason adds the moverReason to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetMoverReason(moverReason *string) {
	o.MoverReason = moverReason
}

// WithMoverState adds the moverState to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithMoverState(moverState *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetMoverState(moverState)
	return o
}

// SetMoverState adds the moverState to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetMoverState(moverState *string) {
	o.MoverState = moverState
}

// WithNodeName adds the nodeName to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithNodeName(nodeName *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithNodeUUID adds the nodeUUID to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithNodeUUID(nodeUUID *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetNodeUUID(nodeUUID *string) {
	o.NodeUUID = nodeUUID
}

// WithReturnRecords adds the returnRecords to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *NdmpNodeSessionDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *NdmpNodeSessionDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScsiDeviceID adds the scsiDeviceID to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithScsiDeviceID(scsiDeviceID *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetScsiDeviceID(scsiDeviceID)
	return o
}

// SetScsiDeviceID adds the scsiDeviceId to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetScsiDeviceID(scsiDeviceID *string) {
	o.ScsiDeviceID = scsiDeviceID
}

// WithScsiHostAdapter adds the scsiHostAdapter to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithScsiHostAdapter(scsiHostAdapter *int64) *NdmpNodeSessionDeleteCollectionParams {
	o.SetScsiHostAdapter(scsiHostAdapter)
	return o
}

// SetScsiHostAdapter adds the scsiHostAdapter to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetScsiHostAdapter(scsiHostAdapter *int64) {
	o.ScsiHostAdapter = scsiHostAdapter
}

// WithScsiLunID adds the scsiLunID to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithScsiLunID(scsiLunID *int64) *NdmpNodeSessionDeleteCollectionParams {
	o.SetScsiLunID(scsiLunID)
	return o
}

// SetScsiLunID adds the scsiLunId to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetScsiLunID(scsiLunID *int64) {
	o.ScsiLunID = scsiLunID
}

// WithScsiTargetID adds the scsiTargetID to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithScsiTargetID(scsiTargetID *int64) *NdmpNodeSessionDeleteCollectionParams {
	o.SetScsiTargetID(scsiTargetID)
	return o
}

// SetScsiTargetID adds the scsiTargetId to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetScsiTargetID(scsiTargetID *int64) {
	o.ScsiTargetID = scsiTargetID
}

// WithSerialRecords adds the serialRecords to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *NdmpNodeSessionDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSourceAddress adds the sourceAddress to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithSourceAddress(sourceAddress *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetSourceAddress(sourceAddress)
	return o
}

// SetSourceAddress adds the sourceAddress to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetSourceAddress(sourceAddress *string) {
	o.SourceAddress = sourceAddress
}

// WithSvmName adds the svmName to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithSvmName(svmName *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithSvmUUID(svmUUID *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithTapeDevice adds the tapeDevice to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithTapeDevice(tapeDevice *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetTapeDevice(tapeDevice)
	return o
}

// SetTapeDevice adds the tapeDevice to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetTapeDevice(tapeDevice *string) {
	o.TapeDevice = tapeDevice
}

// WithTapeMode adds the tapeMode to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithTapeMode(tapeMode *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetTapeMode(tapeMode)
	return o
}

// SetTapeMode adds the tapeMode to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetTapeMode(tapeMode *string) {
	o.TapeMode = tapeMode
}

// WithUUID adds the uuid to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) WithUUID(uuid *string) *NdmpNodeSessionDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the ndmp node session delete collection params
func (o *NdmpNodeSessionDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NdmpNodeSessionDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BackupEngine != nil {

		// query param backup_engine
		var qrBackupEngine string

		if o.BackupEngine != nil {
			qrBackupEngine = *o.BackupEngine
		}
		qBackupEngine := qrBackupEngine
		if qBackupEngine != "" {

			if err := r.SetQueryParam("backup_engine", qBackupEngine); err != nil {
				return err
			}
		}
	}

	if o.ClientAddress != nil {

		// query param client_address
		var qrClientAddress string

		if o.ClientAddress != nil {
			qrClientAddress = *o.ClientAddress
		}
		qClientAddress := qrClientAddress
		if qClientAddress != "" {

			if err := r.SetQueryParam("client_address", qClientAddress); err != nil {
				return err
			}
		}
	}

	if o.ClientPort != nil {

		// query param client_port
		var qrClientPort int64

		if o.ClientPort != nil {
			qrClientPort = *o.ClientPort
		}
		qClientPort := swag.FormatInt64(qrClientPort)
		if qClientPort != "" {

			if err := r.SetQueryParam("client_port", qClientPort); err != nil {
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

	if o.DataBytesProcessed != nil {

		// query param data.bytes_processed
		var qrDataBytesProcessed int64

		if o.DataBytesProcessed != nil {
			qrDataBytesProcessed = *o.DataBytesProcessed
		}
		qDataBytesProcessed := swag.FormatInt64(qrDataBytesProcessed)
		if qDataBytesProcessed != "" {

			if err := r.SetQueryParam("data.bytes_processed", qDataBytesProcessed); err != nil {
				return err
			}
		}
	}

	if o.DataConnectionAddress != nil {

		// query param data.connection.address
		var qrDataConnectionAddress string

		if o.DataConnectionAddress != nil {
			qrDataConnectionAddress = *o.DataConnectionAddress
		}
		qDataConnectionAddress := qrDataConnectionAddress
		if qDataConnectionAddress != "" {

			if err := r.SetQueryParam("data.connection.address", qDataConnectionAddress); err != nil {
				return err
			}
		}
	}

	if o.DataConnectionPort != nil {

		// query param data.connection.port
		var qrDataConnectionPort int64

		if o.DataConnectionPort != nil {
			qrDataConnectionPort = *o.DataConnectionPort
		}
		qDataConnectionPort := swag.FormatInt64(qrDataConnectionPort)
		if qDataConnectionPort != "" {

			if err := r.SetQueryParam("data.connection.port", qDataConnectionPort); err != nil {
				return err
			}
		}
	}

	if o.DataConnectionType != nil {

		// query param data.connection.type
		var qrDataConnectionType string

		if o.DataConnectionType != nil {
			qrDataConnectionType = *o.DataConnectionType
		}
		qDataConnectionType := qrDataConnectionType
		if qDataConnectionType != "" {

			if err := r.SetQueryParam("data.connection.type", qDataConnectionType); err != nil {
				return err
			}
		}
	}

	if o.DataOperation != nil {

		// query param data.operation
		var qrDataOperation string

		if o.DataOperation != nil {
			qrDataOperation = *o.DataOperation
		}
		qDataOperation := qrDataOperation
		if qDataOperation != "" {

			if err := r.SetQueryParam("data.operation", qDataOperation); err != nil {
				return err
			}
		}
	}

	if o.DataReason != nil {

		// query param data.reason
		var qrDataReason string

		if o.DataReason != nil {
			qrDataReason = *o.DataReason
		}
		qDataReason := qrDataReason
		if qDataReason != "" {

			if err := r.SetQueryParam("data.reason", qDataReason); err != nil {
				return err
			}
		}
	}

	if o.DataState != nil {

		// query param data.state
		var qrDataState string

		if o.DataState != nil {
			qrDataState = *o.DataState
		}
		qDataState := qrDataState
		if qDataState != "" {

			if err := r.SetQueryParam("data.state", qDataState); err != nil {
				return err
			}
		}
	}

	if o.DataPath != nil {

		// query param data_path
		var qrDataPath string

		if o.DataPath != nil {
			qrDataPath = *o.DataPath
		}
		qDataPath := qrDataPath
		if qDataPath != "" {

			if err := r.SetQueryParam("data_path", qDataPath); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.MoverBytesMoved != nil {

		// query param mover.bytes_moved
		var qrMoverBytesMoved int64

		if o.MoverBytesMoved != nil {
			qrMoverBytesMoved = *o.MoverBytesMoved
		}
		qMoverBytesMoved := swag.FormatInt64(qrMoverBytesMoved)
		if qMoverBytesMoved != "" {

			if err := r.SetQueryParam("mover.bytes_moved", qMoverBytesMoved); err != nil {
				return err
			}
		}
	}

	if o.MoverConnectionAddress != nil {

		// query param mover.connection.address
		var qrMoverConnectionAddress string

		if o.MoverConnectionAddress != nil {
			qrMoverConnectionAddress = *o.MoverConnectionAddress
		}
		qMoverConnectionAddress := qrMoverConnectionAddress
		if qMoverConnectionAddress != "" {

			if err := r.SetQueryParam("mover.connection.address", qMoverConnectionAddress); err != nil {
				return err
			}
		}
	}

	if o.MoverConnectionPort != nil {

		// query param mover.connection.port
		var qrMoverConnectionPort int64

		if o.MoverConnectionPort != nil {
			qrMoverConnectionPort = *o.MoverConnectionPort
		}
		qMoverConnectionPort := swag.FormatInt64(qrMoverConnectionPort)
		if qMoverConnectionPort != "" {

			if err := r.SetQueryParam("mover.connection.port", qMoverConnectionPort); err != nil {
				return err
			}
		}
	}

	if o.MoverConnectionType != nil {

		// query param mover.connection.type
		var qrMoverConnectionType string

		if o.MoverConnectionType != nil {
			qrMoverConnectionType = *o.MoverConnectionType
		}
		qMoverConnectionType := qrMoverConnectionType
		if qMoverConnectionType != "" {

			if err := r.SetQueryParam("mover.connection.type", qMoverConnectionType); err != nil {
				return err
			}
		}
	}

	if o.MoverMode != nil {

		// query param mover.mode
		var qrMoverMode string

		if o.MoverMode != nil {
			qrMoverMode = *o.MoverMode
		}
		qMoverMode := qrMoverMode
		if qMoverMode != "" {

			if err := r.SetQueryParam("mover.mode", qMoverMode); err != nil {
				return err
			}
		}
	}

	if o.MoverReason != nil {

		// query param mover.reason
		var qrMoverReason string

		if o.MoverReason != nil {
			qrMoverReason = *o.MoverReason
		}
		qMoverReason := qrMoverReason
		if qMoverReason != "" {

			if err := r.SetQueryParam("mover.reason", qMoverReason); err != nil {
				return err
			}
		}
	}

	if o.MoverState != nil {

		// query param mover.state
		var qrMoverState string

		if o.MoverState != nil {
			qrMoverState = *o.MoverState
		}
		qMoverState := qrMoverState
		if qMoverState != "" {

			if err := r.SetQueryParam("mover.state", qMoverState); err != nil {
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

	if o.ScsiDeviceID != nil {

		// query param scsi.device_id
		var qrScsiDeviceID string

		if o.ScsiDeviceID != nil {
			qrScsiDeviceID = *o.ScsiDeviceID
		}
		qScsiDeviceID := qrScsiDeviceID
		if qScsiDeviceID != "" {

			if err := r.SetQueryParam("scsi.device_id", qScsiDeviceID); err != nil {
				return err
			}
		}
	}

	if o.ScsiHostAdapter != nil {

		// query param scsi.host_adapter
		var qrScsiHostAdapter int64

		if o.ScsiHostAdapter != nil {
			qrScsiHostAdapter = *o.ScsiHostAdapter
		}
		qScsiHostAdapter := swag.FormatInt64(qrScsiHostAdapter)
		if qScsiHostAdapter != "" {

			if err := r.SetQueryParam("scsi.host_adapter", qScsiHostAdapter); err != nil {
				return err
			}
		}
	}

	if o.ScsiLunID != nil {

		// query param scsi.lun_id
		var qrScsiLunID int64

		if o.ScsiLunID != nil {
			qrScsiLunID = *o.ScsiLunID
		}
		qScsiLunID := swag.FormatInt64(qrScsiLunID)
		if qScsiLunID != "" {

			if err := r.SetQueryParam("scsi.lun_id", qScsiLunID); err != nil {
				return err
			}
		}
	}

	if o.ScsiTargetID != nil {

		// query param scsi.target_id
		var qrScsiTargetID int64

		if o.ScsiTargetID != nil {
			qrScsiTargetID = *o.ScsiTargetID
		}
		qScsiTargetID := swag.FormatInt64(qrScsiTargetID)
		if qScsiTargetID != "" {

			if err := r.SetQueryParam("scsi.target_id", qScsiTargetID); err != nil {
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

	if o.SourceAddress != nil {

		// query param source_address
		var qrSourceAddress string

		if o.SourceAddress != nil {
			qrSourceAddress = *o.SourceAddress
		}
		qSourceAddress := qrSourceAddress
		if qSourceAddress != "" {

			if err := r.SetQueryParam("source_address", qSourceAddress); err != nil {
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

	if o.TapeDevice != nil {

		// query param tape_device
		var qrTapeDevice string

		if o.TapeDevice != nil {
			qrTapeDevice = *o.TapeDevice
		}
		qTapeDevice := qrTapeDevice
		if qTapeDevice != "" {

			if err := r.SetQueryParam("tape_device", qTapeDevice); err != nil {
				return err
			}
		}
	}

	if o.TapeMode != nil {

		// query param tape_mode
		var qrTapeMode string

		if o.TapeMode != nil {
			qrTapeMode = *o.TapeMode
		}
		qTapeMode := qrTapeMode
		if qTapeMode != "" {

			if err := r.SetQueryParam("tape_mode", qTapeMode); err != nil {
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
