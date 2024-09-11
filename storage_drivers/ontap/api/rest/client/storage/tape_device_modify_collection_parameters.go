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

// NewTapeDeviceModifyCollectionParams creates a new TapeDeviceModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTapeDeviceModifyCollectionParams() *TapeDeviceModifyCollectionParams {
	return &TapeDeviceModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTapeDeviceModifyCollectionParamsWithTimeout creates a new TapeDeviceModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewTapeDeviceModifyCollectionParamsWithTimeout(timeout time.Duration) *TapeDeviceModifyCollectionParams {
	return &TapeDeviceModifyCollectionParams{
		timeout: timeout,
	}
}

// NewTapeDeviceModifyCollectionParamsWithContext creates a new TapeDeviceModifyCollectionParams object
// with the ability to set a context for a request.
func NewTapeDeviceModifyCollectionParamsWithContext(ctx context.Context) *TapeDeviceModifyCollectionParams {
	return &TapeDeviceModifyCollectionParams{
		Context: ctx,
	}
}

// NewTapeDeviceModifyCollectionParamsWithHTTPClient creates a new TapeDeviceModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewTapeDeviceModifyCollectionParamsWithHTTPClient(client *http.Client) *TapeDeviceModifyCollectionParams {
	return &TapeDeviceModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
TapeDeviceModifyCollectionParams contains all the parameters to send to the API endpoint

	for the tape device modify collection operation.

	Typically these are written to a http.Request.
*/
type TapeDeviceModifyCollectionParams struct {

	/* AliasMapping.

	   Filter by alias.mapping
	*/
	AliasMapping *string

	/* AliasName.

	   Filter by alias.name
	*/
	AliasName *string

	/* AliasesMapping.

	   Filter by aliases.mapping
	*/
	AliasesMapping *string

	/* AliasesName.

	   Filter by aliases.name
	*/
	AliasesName *string

	/* BlockNumber.

	   Filter by block_number
	*/
	BlockNumber *int64

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Density.

	   Filter by density
	*/
	Density *string

	/* Description.

	   Filter by description
	*/
	Description *string

	/* DeviceID.

	   Filter by device_id
	*/
	DeviceID *string

	/* DeviceNamesNoRewindDevice.

	   Filter by device_names.no_rewind_device
	*/
	DeviceNamesNoRewindDevice *string

	/* DeviceNamesRewindDevice.

	   Filter by device_names.rewind_device
	*/
	DeviceNamesRewindDevice *string

	/* DeviceNamesUnloadReloadDevice.

	   Filter by device_names.unload_reload_device
	*/
	DeviceNamesUnloadReloadDevice *string

	/* DeviceState.

	   Filter by device_state
	*/
	DeviceState *string

	/* FileNumber.

	   Filter by file_number
	*/
	FileNumber *int64

	/* Formats.

	   Filter by formats
	*/
	Formats *string

	/* Info.

	   Info specification
	*/
	Info TapeDeviceModifyCollectionBody

	/* Interface.

	   Filter by interface
	*/
	Interface *string

	/* NodeName.

	   Filter by node.name
	*/
	NodeName *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUID *string

	/* Online.

	   Filter by online
	*/
	Online *bool

	/* ReservationType.

	   Filter by reservation_type
	*/
	ReservationType *string

	/* ResidualCount.

	   Filter by residual_count
	*/
	ResidualCount *int64

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

	/* StoragePortName.

	   Filter by storage_port.name
	*/
	StoragePortName *string

	/* Type.

	   Filter by type
	*/
	Type *string

	/* Wwnn.

	   Filter by wwnn
	*/
	Wwnn *string

	/* Wwpn.

	   Filter by wwpn
	*/
	Wwpn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tape device modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TapeDeviceModifyCollectionParams) WithDefaults() *TapeDeviceModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tape device modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TapeDeviceModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := TapeDeviceModifyCollectionParams{
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

// WithTimeout adds the timeout to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithTimeout(timeout time.Duration) *TapeDeviceModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithContext(ctx context.Context) *TapeDeviceModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithHTTPClient(client *http.Client) *TapeDeviceModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAliasMapping adds the aliasMapping to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithAliasMapping(aliasMapping *string) *TapeDeviceModifyCollectionParams {
	o.SetAliasMapping(aliasMapping)
	return o
}

// SetAliasMapping adds the aliasMapping to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetAliasMapping(aliasMapping *string) {
	o.AliasMapping = aliasMapping
}

// WithAliasName adds the aliasName to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithAliasName(aliasName *string) *TapeDeviceModifyCollectionParams {
	o.SetAliasName(aliasName)
	return o
}

// SetAliasName adds the aliasName to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetAliasName(aliasName *string) {
	o.AliasName = aliasName
}

// WithAliasesMapping adds the aliasesMapping to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithAliasesMapping(aliasesMapping *string) *TapeDeviceModifyCollectionParams {
	o.SetAliasesMapping(aliasesMapping)
	return o
}

// SetAliasesMapping adds the aliasesMapping to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetAliasesMapping(aliasesMapping *string) {
	o.AliasesMapping = aliasesMapping
}

// WithAliasesName adds the aliasesName to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithAliasesName(aliasesName *string) *TapeDeviceModifyCollectionParams {
	o.SetAliasesName(aliasesName)
	return o
}

// SetAliasesName adds the aliasesName to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetAliasesName(aliasesName *string) {
	o.AliasesName = aliasesName
}

// WithBlockNumber adds the blockNumber to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithBlockNumber(blockNumber *int64) *TapeDeviceModifyCollectionParams {
	o.SetBlockNumber(blockNumber)
	return o
}

// SetBlockNumber adds the blockNumber to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetBlockNumber(blockNumber *int64) {
	o.BlockNumber = blockNumber
}

// WithContinueOnFailure adds the continueOnFailure to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *TapeDeviceModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDensity adds the density to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithDensity(density *string) *TapeDeviceModifyCollectionParams {
	o.SetDensity(density)
	return o
}

// SetDensity adds the density to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetDensity(density *string) {
	o.Density = density
}

// WithDescription adds the description to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithDescription(description *string) *TapeDeviceModifyCollectionParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetDescription(description *string) {
	o.Description = description
}

// WithDeviceID adds the deviceID to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithDeviceID(deviceID *string) *TapeDeviceModifyCollectionParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceNamesNoRewindDevice adds the deviceNamesNoRewindDevice to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithDeviceNamesNoRewindDevice(deviceNamesNoRewindDevice *string) *TapeDeviceModifyCollectionParams {
	o.SetDeviceNamesNoRewindDevice(deviceNamesNoRewindDevice)
	return o
}

// SetDeviceNamesNoRewindDevice adds the deviceNamesNoRewindDevice to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetDeviceNamesNoRewindDevice(deviceNamesNoRewindDevice *string) {
	o.DeviceNamesNoRewindDevice = deviceNamesNoRewindDevice
}

// WithDeviceNamesRewindDevice adds the deviceNamesRewindDevice to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithDeviceNamesRewindDevice(deviceNamesRewindDevice *string) *TapeDeviceModifyCollectionParams {
	o.SetDeviceNamesRewindDevice(deviceNamesRewindDevice)
	return o
}

// SetDeviceNamesRewindDevice adds the deviceNamesRewindDevice to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetDeviceNamesRewindDevice(deviceNamesRewindDevice *string) {
	o.DeviceNamesRewindDevice = deviceNamesRewindDevice
}

// WithDeviceNamesUnloadReloadDevice adds the deviceNamesUnloadReloadDevice to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithDeviceNamesUnloadReloadDevice(deviceNamesUnloadReloadDevice *string) *TapeDeviceModifyCollectionParams {
	o.SetDeviceNamesUnloadReloadDevice(deviceNamesUnloadReloadDevice)
	return o
}

// SetDeviceNamesUnloadReloadDevice adds the deviceNamesUnloadReloadDevice to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetDeviceNamesUnloadReloadDevice(deviceNamesUnloadReloadDevice *string) {
	o.DeviceNamesUnloadReloadDevice = deviceNamesUnloadReloadDevice
}

// WithDeviceState adds the deviceState to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithDeviceState(deviceState *string) *TapeDeviceModifyCollectionParams {
	o.SetDeviceState(deviceState)
	return o
}

// SetDeviceState adds the deviceState to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetDeviceState(deviceState *string) {
	o.DeviceState = deviceState
}

// WithFileNumber adds the fileNumber to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithFileNumber(fileNumber *int64) *TapeDeviceModifyCollectionParams {
	o.SetFileNumber(fileNumber)
	return o
}

// SetFileNumber adds the fileNumber to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetFileNumber(fileNumber *int64) {
	o.FileNumber = fileNumber
}

// WithFormats adds the formats to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithFormats(formats *string) *TapeDeviceModifyCollectionParams {
	o.SetFormats(formats)
	return o
}

// SetFormats adds the formats to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetFormats(formats *string) {
	o.Formats = formats
}

// WithInfo adds the info to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithInfo(info TapeDeviceModifyCollectionBody) *TapeDeviceModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetInfo(info TapeDeviceModifyCollectionBody) {
	o.Info = info
}

// WithInterface adds the interfaceVar to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithInterface(interfaceVar *string) *TapeDeviceModifyCollectionParams {
	o.SetInterface(interfaceVar)
	return o
}

// SetInterface adds the interface to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetInterface(interfaceVar *string) {
	o.Interface = interfaceVar
}

// WithNodeName adds the nodeName to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithNodeName(nodeName *string) *TapeDeviceModifyCollectionParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithNodeUUID adds the nodeUUID to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithNodeUUID(nodeUUID *string) *TapeDeviceModifyCollectionParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetNodeUUID(nodeUUID *string) {
	o.NodeUUID = nodeUUID
}

// WithOnline adds the online to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithOnline(online *bool) *TapeDeviceModifyCollectionParams {
	o.SetOnline(online)
	return o
}

// SetOnline adds the online to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetOnline(online *bool) {
	o.Online = online
}

// WithReservationType adds the reservationType to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithReservationType(reservationType *string) *TapeDeviceModifyCollectionParams {
	o.SetReservationType(reservationType)
	return o
}

// SetReservationType adds the reservationType to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetReservationType(reservationType *string) {
	o.ReservationType = reservationType
}

// WithResidualCount adds the residualCount to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithResidualCount(residualCount *int64) *TapeDeviceModifyCollectionParams {
	o.SetResidualCount(residualCount)
	return o
}

// SetResidualCount adds the residualCount to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetResidualCount(residualCount *int64) {
	o.ResidualCount = residualCount
}

// WithReturnRecords adds the returnRecords to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithReturnRecords(returnRecords *bool) *TapeDeviceModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *TapeDeviceModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialNumber adds the serialNumber to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithSerialNumber(serialNumber *string) *TapeDeviceModifyCollectionParams {
	o.SetSerialNumber(serialNumber)
	return o
}

// SetSerialNumber adds the serialNumber to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetSerialNumber(serialNumber *string) {
	o.SerialNumber = serialNumber
}

// WithSerialRecords adds the serialRecords to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithSerialRecords(serialRecords *bool) *TapeDeviceModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithStoragePortName adds the storagePortName to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithStoragePortName(storagePortName *string) *TapeDeviceModifyCollectionParams {
	o.SetStoragePortName(storagePortName)
	return o
}

// SetStoragePortName adds the storagePortName to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetStoragePortName(storagePortName *string) {
	o.StoragePortName = storagePortName
}

// WithType adds the typeVar to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithType(typeVar *string) *TapeDeviceModifyCollectionParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithWwnn adds the wwnn to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithWwnn(wwnn *string) *TapeDeviceModifyCollectionParams {
	o.SetWwnn(wwnn)
	return o
}

// SetWwnn adds the wwnn to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetWwnn(wwnn *string) {
	o.Wwnn = wwnn
}

// WithWwpn adds the wwpn to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) WithWwpn(wwpn *string) *TapeDeviceModifyCollectionParams {
	o.SetWwpn(wwpn)
	return o
}

// SetWwpn adds the wwpn to the tape device modify collection params
func (o *TapeDeviceModifyCollectionParams) SetWwpn(wwpn *string) {
	o.Wwpn = wwpn
}

// WriteToRequest writes these params to a swagger request
func (o *TapeDeviceModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AliasMapping != nil {

		// query param alias.mapping
		var qrAliasMapping string

		if o.AliasMapping != nil {
			qrAliasMapping = *o.AliasMapping
		}
		qAliasMapping := qrAliasMapping
		if qAliasMapping != "" {

			if err := r.SetQueryParam("alias.mapping", qAliasMapping); err != nil {
				return err
			}
		}
	}

	if o.AliasName != nil {

		// query param alias.name
		var qrAliasName string

		if o.AliasName != nil {
			qrAliasName = *o.AliasName
		}
		qAliasName := qrAliasName
		if qAliasName != "" {

			if err := r.SetQueryParam("alias.name", qAliasName); err != nil {
				return err
			}
		}
	}

	if o.AliasesMapping != nil {

		// query param aliases.mapping
		var qrAliasesMapping string

		if o.AliasesMapping != nil {
			qrAliasesMapping = *o.AliasesMapping
		}
		qAliasesMapping := qrAliasesMapping
		if qAliasesMapping != "" {

			if err := r.SetQueryParam("aliases.mapping", qAliasesMapping); err != nil {
				return err
			}
		}
	}

	if o.AliasesName != nil {

		// query param aliases.name
		var qrAliasesName string

		if o.AliasesName != nil {
			qrAliasesName = *o.AliasesName
		}
		qAliasesName := qrAliasesName
		if qAliasesName != "" {

			if err := r.SetQueryParam("aliases.name", qAliasesName); err != nil {
				return err
			}
		}
	}

	if o.BlockNumber != nil {

		// query param block_number
		var qrBlockNumber int64

		if o.BlockNumber != nil {
			qrBlockNumber = *o.BlockNumber
		}
		qBlockNumber := swag.FormatInt64(qrBlockNumber)
		if qBlockNumber != "" {

			if err := r.SetQueryParam("block_number", qBlockNumber); err != nil {
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

	if o.Density != nil {

		// query param density
		var qrDensity string

		if o.Density != nil {
			qrDensity = *o.Density
		}
		qDensity := qrDensity
		if qDensity != "" {

			if err := r.SetQueryParam("density", qDensity); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string

		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {

			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}
	}

	if o.DeviceNamesNoRewindDevice != nil {

		// query param device_names.no_rewind_device
		var qrDeviceNamesNoRewindDevice string

		if o.DeviceNamesNoRewindDevice != nil {
			qrDeviceNamesNoRewindDevice = *o.DeviceNamesNoRewindDevice
		}
		qDeviceNamesNoRewindDevice := qrDeviceNamesNoRewindDevice
		if qDeviceNamesNoRewindDevice != "" {

			if err := r.SetQueryParam("device_names.no_rewind_device", qDeviceNamesNoRewindDevice); err != nil {
				return err
			}
		}
	}

	if o.DeviceNamesRewindDevice != nil {

		// query param device_names.rewind_device
		var qrDeviceNamesRewindDevice string

		if o.DeviceNamesRewindDevice != nil {
			qrDeviceNamesRewindDevice = *o.DeviceNamesRewindDevice
		}
		qDeviceNamesRewindDevice := qrDeviceNamesRewindDevice
		if qDeviceNamesRewindDevice != "" {

			if err := r.SetQueryParam("device_names.rewind_device", qDeviceNamesRewindDevice); err != nil {
				return err
			}
		}
	}

	if o.DeviceNamesUnloadReloadDevice != nil {

		// query param device_names.unload_reload_device
		var qrDeviceNamesUnloadReloadDevice string

		if o.DeviceNamesUnloadReloadDevice != nil {
			qrDeviceNamesUnloadReloadDevice = *o.DeviceNamesUnloadReloadDevice
		}
		qDeviceNamesUnloadReloadDevice := qrDeviceNamesUnloadReloadDevice
		if qDeviceNamesUnloadReloadDevice != "" {

			if err := r.SetQueryParam("device_names.unload_reload_device", qDeviceNamesUnloadReloadDevice); err != nil {
				return err
			}
		}
	}

	if o.DeviceState != nil {

		// query param device_state
		var qrDeviceState string

		if o.DeviceState != nil {
			qrDeviceState = *o.DeviceState
		}
		qDeviceState := qrDeviceState
		if qDeviceState != "" {

			if err := r.SetQueryParam("device_state", qDeviceState); err != nil {
				return err
			}
		}
	}

	if o.FileNumber != nil {

		// query param file_number
		var qrFileNumber int64

		if o.FileNumber != nil {
			qrFileNumber = *o.FileNumber
		}
		qFileNumber := swag.FormatInt64(qrFileNumber)
		if qFileNumber != "" {

			if err := r.SetQueryParam("file_number", qFileNumber); err != nil {
				return err
			}
		}
	}

	if o.Formats != nil {

		// query param formats
		var qrFormats string

		if o.Formats != nil {
			qrFormats = *o.Formats
		}
		qFormats := qrFormats
		if qFormats != "" {

			if err := r.SetQueryParam("formats", qFormats); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.Interface != nil {

		// query param interface
		var qrInterface string

		if o.Interface != nil {
			qrInterface = *o.Interface
		}
		qInterface := qrInterface
		if qInterface != "" {

			if err := r.SetQueryParam("interface", qInterface); err != nil {
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

	if o.Online != nil {

		// query param online
		var qrOnline bool

		if o.Online != nil {
			qrOnline = *o.Online
		}
		qOnline := swag.FormatBool(qrOnline)
		if qOnline != "" {

			if err := r.SetQueryParam("online", qOnline); err != nil {
				return err
			}
		}
	}

	if o.ReservationType != nil {

		// query param reservation_type
		var qrReservationType string

		if o.ReservationType != nil {
			qrReservationType = *o.ReservationType
		}
		qReservationType := qrReservationType
		if qReservationType != "" {

			if err := r.SetQueryParam("reservation_type", qReservationType); err != nil {
				return err
			}
		}
	}

	if o.ResidualCount != nil {

		// query param residual_count
		var qrResidualCount int64

		if o.ResidualCount != nil {
			qrResidualCount = *o.ResidualCount
		}
		qResidualCount := swag.FormatInt64(qrResidualCount)
		if qResidualCount != "" {

			if err := r.SetQueryParam("residual_count", qResidualCount); err != nil {
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

	if o.StoragePortName != nil {

		// query param storage_port.name
		var qrStoragePortName string

		if o.StoragePortName != nil {
			qrStoragePortName = *o.StoragePortName
		}
		qStoragePortName := qrStoragePortName
		if qStoragePortName != "" {

			if err := r.SetQueryParam("storage_port.name", qStoragePortName); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Wwnn != nil {

		// query param wwnn
		var qrWwnn string

		if o.Wwnn != nil {
			qrWwnn = *o.Wwnn
		}
		qWwnn := qrWwnn
		if qWwnn != "" {

			if err := r.SetQueryParam("wwnn", qWwnn); err != nil {
				return err
			}
		}
	}

	if o.Wwpn != nil {

		// query param wwpn
		var qrWwpn string

		if o.Wwpn != nil {
			qrWwpn = *o.Wwpn
		}
		qWwpn := qrWwpn
		if qWwpn != "" {

			if err := r.SetQueryParam("wwpn", qWwpn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
