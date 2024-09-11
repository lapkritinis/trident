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

// NewSnapshotDeleteCollectionParams creates a new SnapshotDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotDeleteCollectionParams() *SnapshotDeleteCollectionParams {
	return &SnapshotDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotDeleteCollectionParamsWithTimeout creates a new SnapshotDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewSnapshotDeleteCollectionParamsWithTimeout(timeout time.Duration) *SnapshotDeleteCollectionParams {
	return &SnapshotDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewSnapshotDeleteCollectionParamsWithContext creates a new SnapshotDeleteCollectionParams object
// with the ability to set a context for a request.
func NewSnapshotDeleteCollectionParamsWithContext(ctx context.Context) *SnapshotDeleteCollectionParams {
	return &SnapshotDeleteCollectionParams{
		Context: ctx,
	}
}

// NewSnapshotDeleteCollectionParamsWithHTTPClient creates a new SnapshotDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotDeleteCollectionParamsWithHTTPClient(client *http.Client) *SnapshotDeleteCollectionParams {
	return &SnapshotDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
SnapshotDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the snapshot delete collection operation.

	Typically these are written to a http.Request.
*/
type SnapshotDeleteCollectionParams struct {

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* CreateTime.

	   Filter by create_time
	*/
	CreateTime *string

	/* DeltaSizeConsumed.

	   Filter by delta.size_consumed
	*/
	DeltaSizeConsumed *int64

	/* DeltaTimeElapsed.

	   Filter by delta.time_elapsed
	*/
	DeltaTimeElapsed *string

	/* ExpiryTime.

	   Filter by expiry_time
	*/
	ExpiryTime *string

	/* Info.

	   Info specification
	*/
	Info SnapshotDeleteCollectionBody

	/* LogicalSize.

	   Filter by logical_size
	*/
	LogicalSize *int64

	/* Name.

	   Filter by name
	*/
	Name *string

	/* Owners.

	   Filter by owners
	*/
	Owners *string

	/* ProvenanceVolumeUUID.

	   Filter by provenance_volume.uuid
	*/
	ProvenanceVolumeUUID *string

	/* ReclaimableSpace.

	   Filter by reclaimable_space
	*/
	ReclaimableSpace *int64

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

	/* SnaplockExpired.

	   Filter by snaplock.expired
	*/
	SnaplockExpired *bool

	/* SnaplockTimeUntilExpiry.

	   Filter by snaplock.time_until_expiry
	*/
	SnaplockTimeUntilExpiry *string

	/* SnaplockExpiryTime.

	   Filter by snaplock_expiry_time
	*/
	SnaplockExpiryTime *string

	/* SnapmirrorLabel.

	   Filter by snapmirror_label
	*/
	SnapmirrorLabel *string

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

	/* VersionUUID.

	   Filter by version_uuid
	*/
	VersionUUID *string

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeName *string

	/* VolumeUUID.

	   Volume
	*/
	VolumeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotDeleteCollectionParams) WithDefaults() *SnapshotDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SnapshotDeleteCollectionParams{
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

// WithTimeout adds the timeout to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithTimeout(timeout time.Duration) *SnapshotDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithContext(ctx context.Context) *SnapshotDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithHTTPClient(client *http.Client) *SnapshotDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithComment(comment *string) *SnapshotDeleteCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContinueOnFailure adds the continueOnFailure to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SnapshotDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithCreateTime adds the createTime to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithCreateTime(createTime *string) *SnapshotDeleteCollectionParams {
	o.SetCreateTime(createTime)
	return o
}

// SetCreateTime adds the createTime to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetCreateTime(createTime *string) {
	o.CreateTime = createTime
}

// WithDeltaSizeConsumed adds the deltaSizeConsumed to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithDeltaSizeConsumed(deltaSizeConsumed *int64) *SnapshotDeleteCollectionParams {
	o.SetDeltaSizeConsumed(deltaSizeConsumed)
	return o
}

// SetDeltaSizeConsumed adds the deltaSizeConsumed to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetDeltaSizeConsumed(deltaSizeConsumed *int64) {
	o.DeltaSizeConsumed = deltaSizeConsumed
}

// WithDeltaTimeElapsed adds the deltaTimeElapsed to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithDeltaTimeElapsed(deltaTimeElapsed *string) *SnapshotDeleteCollectionParams {
	o.SetDeltaTimeElapsed(deltaTimeElapsed)
	return o
}

// SetDeltaTimeElapsed adds the deltaTimeElapsed to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetDeltaTimeElapsed(deltaTimeElapsed *string) {
	o.DeltaTimeElapsed = deltaTimeElapsed
}

// WithExpiryTime adds the expiryTime to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithExpiryTime(expiryTime *string) *SnapshotDeleteCollectionParams {
	o.SetExpiryTime(expiryTime)
	return o
}

// SetExpiryTime adds the expiryTime to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetExpiryTime(expiryTime *string) {
	o.ExpiryTime = expiryTime
}

// WithInfo adds the info to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithInfo(info SnapshotDeleteCollectionBody) *SnapshotDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetInfo(info SnapshotDeleteCollectionBody) {
	o.Info = info
}

// WithLogicalSize adds the logicalSize to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithLogicalSize(logicalSize *int64) *SnapshotDeleteCollectionParams {
	o.SetLogicalSize(logicalSize)
	return o
}

// SetLogicalSize adds the logicalSize to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetLogicalSize(logicalSize *int64) {
	o.LogicalSize = logicalSize
}

// WithName adds the name to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithName(name *string) *SnapshotDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithOwners adds the owners to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithOwners(owners *string) *SnapshotDeleteCollectionParams {
	o.SetOwners(owners)
	return o
}

// SetOwners adds the owners to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetOwners(owners *string) {
	o.Owners = owners
}

// WithProvenanceVolumeUUID adds the provenanceVolumeUUID to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithProvenanceVolumeUUID(provenanceVolumeUUID *string) *SnapshotDeleteCollectionParams {
	o.SetProvenanceVolumeUUID(provenanceVolumeUUID)
	return o
}

// SetProvenanceVolumeUUID adds the provenanceVolumeUuid to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetProvenanceVolumeUUID(provenanceVolumeUUID *string) {
	o.ProvenanceVolumeUUID = provenanceVolumeUUID
}

// WithReclaimableSpace adds the reclaimableSpace to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithReclaimableSpace(reclaimableSpace *int64) *SnapshotDeleteCollectionParams {
	o.SetReclaimableSpace(reclaimableSpace)
	return o
}

// SetReclaimableSpace adds the reclaimableSpace to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetReclaimableSpace(reclaimableSpace *int64) {
	o.ReclaimableSpace = reclaimableSpace
}

// WithReturnRecords adds the returnRecords to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *SnapshotDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *SnapshotDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *SnapshotDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSize adds the size to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithSize(size *int64) *SnapshotDeleteCollectionParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetSize(size *int64) {
	o.Size = size
}

// WithSnaplockExpired adds the snaplockExpired to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithSnaplockExpired(snaplockExpired *bool) *SnapshotDeleteCollectionParams {
	o.SetSnaplockExpired(snaplockExpired)
	return o
}

// SetSnaplockExpired adds the snaplockExpired to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetSnaplockExpired(snaplockExpired *bool) {
	o.SnaplockExpired = snaplockExpired
}

// WithSnaplockTimeUntilExpiry adds the snaplockTimeUntilExpiry to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithSnaplockTimeUntilExpiry(snaplockTimeUntilExpiry *string) *SnapshotDeleteCollectionParams {
	o.SetSnaplockTimeUntilExpiry(snaplockTimeUntilExpiry)
	return o
}

// SetSnaplockTimeUntilExpiry adds the snaplockTimeUntilExpiry to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetSnaplockTimeUntilExpiry(snaplockTimeUntilExpiry *string) {
	o.SnaplockTimeUntilExpiry = snaplockTimeUntilExpiry
}

// WithSnaplockExpiryTime adds the snaplockExpiryTime to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithSnaplockExpiryTime(snaplockExpiryTime *string) *SnapshotDeleteCollectionParams {
	o.SetSnaplockExpiryTime(snaplockExpiryTime)
	return o
}

// SetSnaplockExpiryTime adds the snaplockExpiryTime to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetSnaplockExpiryTime(snaplockExpiryTime *string) {
	o.SnaplockExpiryTime = snaplockExpiryTime
}

// WithSnapmirrorLabel adds the snapmirrorLabel to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithSnapmirrorLabel(snapmirrorLabel *string) *SnapshotDeleteCollectionParams {
	o.SetSnapmirrorLabel(snapmirrorLabel)
	return o
}

// SetSnapmirrorLabel adds the snapmirrorLabel to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetSnapmirrorLabel(snapmirrorLabel *string) {
	o.SnapmirrorLabel = snapmirrorLabel
}

// WithState adds the state to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithState(state *string) *SnapshotDeleteCollectionParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetState(state *string) {
	o.State = state
}

// WithSvmName adds the svmName to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithSvmName(svmName *string) *SnapshotDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithSvmUUID(svmUUID *string) *SnapshotDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithUUID(uuid *string) *SnapshotDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WithVersionUUID adds the versionUUID to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithVersionUUID(versionUUID *string) *SnapshotDeleteCollectionParams {
	o.SetVersionUUID(versionUUID)
	return o
}

// SetVersionUUID adds the versionUuid to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetVersionUUID(versionUUID *string) {
	o.VersionUUID = versionUUID
}

// WithVolumeName adds the volumeName to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithVolumeName(volumeName *string) *SnapshotDeleteCollectionParams {
	o.SetVolumeName(volumeName)
	return o
}

// SetVolumeName adds the volumeName to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetVolumeName(volumeName *string) {
	o.VolumeName = volumeName
}

// WithVolumeUUID adds the volumeUUID to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) WithVolumeUUID(volumeUUID string) *SnapshotDeleteCollectionParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the snapshot delete collection params
func (o *SnapshotDeleteCollectionParams) SetVolumeUUID(volumeUUID string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DeltaSizeConsumed != nil {

		// query param delta.size_consumed
		var qrDeltaSizeConsumed int64

		if o.DeltaSizeConsumed != nil {
			qrDeltaSizeConsumed = *o.DeltaSizeConsumed
		}
		qDeltaSizeConsumed := swag.FormatInt64(qrDeltaSizeConsumed)
		if qDeltaSizeConsumed != "" {

			if err := r.SetQueryParam("delta.size_consumed", qDeltaSizeConsumed); err != nil {
				return err
			}
		}
	}

	if o.DeltaTimeElapsed != nil {

		// query param delta.time_elapsed
		var qrDeltaTimeElapsed string

		if o.DeltaTimeElapsed != nil {
			qrDeltaTimeElapsed = *o.DeltaTimeElapsed
		}
		qDeltaTimeElapsed := qrDeltaTimeElapsed
		if qDeltaTimeElapsed != "" {

			if err := r.SetQueryParam("delta.time_elapsed", qDeltaTimeElapsed); err != nil {
				return err
			}
		}
	}

	if o.ExpiryTime != nil {

		// query param expiry_time
		var qrExpiryTime string

		if o.ExpiryTime != nil {
			qrExpiryTime = *o.ExpiryTime
		}
		qExpiryTime := qrExpiryTime
		if qExpiryTime != "" {

			if err := r.SetQueryParam("expiry_time", qExpiryTime); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.LogicalSize != nil {

		// query param logical_size
		var qrLogicalSize int64

		if o.LogicalSize != nil {
			qrLogicalSize = *o.LogicalSize
		}
		qLogicalSize := swag.FormatInt64(qrLogicalSize)
		if qLogicalSize != "" {

			if err := r.SetQueryParam("logical_size", qLogicalSize); err != nil {
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

	if o.Owners != nil {

		// query param owners
		var qrOwners string

		if o.Owners != nil {
			qrOwners = *o.Owners
		}
		qOwners := qrOwners
		if qOwners != "" {

			if err := r.SetQueryParam("owners", qOwners); err != nil {
				return err
			}
		}
	}

	if o.ProvenanceVolumeUUID != nil {

		// query param provenance_volume.uuid
		var qrProvenanceVolumeUUID string

		if o.ProvenanceVolumeUUID != nil {
			qrProvenanceVolumeUUID = *o.ProvenanceVolumeUUID
		}
		qProvenanceVolumeUUID := qrProvenanceVolumeUUID
		if qProvenanceVolumeUUID != "" {

			if err := r.SetQueryParam("provenance_volume.uuid", qProvenanceVolumeUUID); err != nil {
				return err
			}
		}
	}

	if o.ReclaimableSpace != nil {

		// query param reclaimable_space
		var qrReclaimableSpace int64

		if o.ReclaimableSpace != nil {
			qrReclaimableSpace = *o.ReclaimableSpace
		}
		qReclaimableSpace := swag.FormatInt64(qrReclaimableSpace)
		if qReclaimableSpace != "" {

			if err := r.SetQueryParam("reclaimable_space", qReclaimableSpace); err != nil {
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

	if o.SnaplockExpired != nil {

		// query param snaplock.expired
		var qrSnaplockExpired bool

		if o.SnaplockExpired != nil {
			qrSnaplockExpired = *o.SnaplockExpired
		}
		qSnaplockExpired := swag.FormatBool(qrSnaplockExpired)
		if qSnaplockExpired != "" {

			if err := r.SetQueryParam("snaplock.expired", qSnaplockExpired); err != nil {
				return err
			}
		}
	}

	if o.SnaplockTimeUntilExpiry != nil {

		// query param snaplock.time_until_expiry
		var qrSnaplockTimeUntilExpiry string

		if o.SnaplockTimeUntilExpiry != nil {
			qrSnaplockTimeUntilExpiry = *o.SnaplockTimeUntilExpiry
		}
		qSnaplockTimeUntilExpiry := qrSnaplockTimeUntilExpiry
		if qSnaplockTimeUntilExpiry != "" {

			if err := r.SetQueryParam("snaplock.time_until_expiry", qSnaplockTimeUntilExpiry); err != nil {
				return err
			}
		}
	}

	if o.SnaplockExpiryTime != nil {

		// query param snaplock_expiry_time
		var qrSnaplockExpiryTime string

		if o.SnaplockExpiryTime != nil {
			qrSnaplockExpiryTime = *o.SnaplockExpiryTime
		}
		qSnaplockExpiryTime := qrSnaplockExpiryTime
		if qSnaplockExpiryTime != "" {

			if err := r.SetQueryParam("snaplock_expiry_time", qSnaplockExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.SnapmirrorLabel != nil {

		// query param snapmirror_label
		var qrSnapmirrorLabel string

		if o.SnapmirrorLabel != nil {
			qrSnapmirrorLabel = *o.SnapmirrorLabel
		}
		qSnapmirrorLabel := qrSnapmirrorLabel
		if qSnapmirrorLabel != "" {

			if err := r.SetQueryParam("snapmirror_label", qSnapmirrorLabel); err != nil {
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

	if o.VersionUUID != nil {

		// query param version_uuid
		var qrVersionUUID string

		if o.VersionUUID != nil {
			qrVersionUUID = *o.VersionUUID
		}
		qVersionUUID := qrVersionUUID
		if qVersionUUID != "" {

			if err := r.SetQueryParam("version_uuid", qVersionUUID); err != nil {
				return err
			}
		}
	}

	if o.VolumeName != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeName != nil {
			qrVolumeName = *o.VolumeName
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
