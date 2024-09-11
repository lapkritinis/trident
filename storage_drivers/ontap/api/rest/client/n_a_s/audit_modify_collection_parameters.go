// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewAuditModifyCollectionParams creates a new AuditModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditModifyCollectionParams() *AuditModifyCollectionParams {
	return &AuditModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditModifyCollectionParamsWithTimeout creates a new AuditModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewAuditModifyCollectionParamsWithTimeout(timeout time.Duration) *AuditModifyCollectionParams {
	return &AuditModifyCollectionParams{
		timeout: timeout,
	}
}

// NewAuditModifyCollectionParamsWithContext creates a new AuditModifyCollectionParams object
// with the ability to set a context for a request.
func NewAuditModifyCollectionParamsWithContext(ctx context.Context) *AuditModifyCollectionParams {
	return &AuditModifyCollectionParams{
		Context: ctx,
	}
}

// NewAuditModifyCollectionParamsWithHTTPClient creates a new AuditModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditModifyCollectionParamsWithHTTPClient(client *http.Client) *AuditModifyCollectionParams {
	return &AuditModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
AuditModifyCollectionParams contains all the parameters to send to the API endpoint

	for the audit modify collection operation.

	Typically these are written to a http.Request.
*/
type AuditModifyCollectionParams struct {

	/* ChargeQos.

	   Filter by charge_qos
	*/
	ChargeQos *bool

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Enabled.

	   Filter by enabled
	*/
	Enabled *bool

	/* EventsAsyncDelete.

	   Filter by events.async_delete
	*/
	EventsAsyncDelete *bool

	/* EventsAuditPolicyChange.

	   Filter by events.audit_policy_change
	*/
	EventsAuditPolicyChange *bool

	/* EventsAuthorizationPolicy.

	   Filter by events.authorization_policy
	*/
	EventsAuthorizationPolicy *bool

	/* EventsCapStaging.

	   Filter by events.cap_staging
	*/
	EventsCapStaging *bool

	/* EventsCifsLogonLogoff.

	   Filter by events.cifs_logon_logoff
	*/
	EventsCifsLogonLogoff *bool

	/* EventsFileOperations.

	   Filter by events.file_operations
	*/
	EventsFileOperations *bool

	/* EventsFileShare.

	   Filter by events.file_share
	*/
	EventsFileShare *bool

	/* EventsSecurityGroup.

	   Filter by events.security_group
	*/
	EventsSecurityGroup *bool

	/* EventsUserAccount.

	   Filter by events.user_account
	*/
	EventsUserAccount *bool

	/* Guarantee.

	   Filter by guarantee
	*/
	Guarantee *bool

	/* Info.

	   Info specification
	*/
	Info AuditModifyCollectionBody

	/* LogFormat.

	   Filter by log.format
	*/
	LogFormat *string

	/* LogRetentionCount.

	   Filter by log.retention.count
	*/
	LogRetentionCount *int64

	/* LogRetentionDuration.

	   Filter by log.retention.duration
	*/
	LogRetentionDuration *string

	/* LogRotationScheduleDays.

	   Filter by log.rotation.schedule.days
	*/
	LogRotationScheduleDays *int64

	/* LogRotationScheduleHours.

	   Filter by log.rotation.schedule.hours
	*/
	LogRotationScheduleHours *int64

	/* LogRotationScheduleMinutes.

	   Filter by log.rotation.schedule.minutes
	*/
	LogRotationScheduleMinutes *int64

	/* LogRotationScheduleMonths.

	   Filter by log.rotation.schedule.months
	*/
	LogRotationScheduleMonths *int64

	/* LogRotationScheduleWeekdays.

	   Filter by log.rotation.schedule.weekdays
	*/
	LogRotationScheduleWeekdays *int64

	/* LogRotationSize.

	   Filter by log.rotation.size
	*/
	LogRotationSize *int64

	/* LogPath.

	   Filter by log_path
	*/
	LogPath *string

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

// WithDefaults hydrates default values in the audit modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditModifyCollectionParams) WithDefaults() *AuditModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := AuditModifyCollectionParams{
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

// WithTimeout adds the timeout to the audit modify collection params
func (o *AuditModifyCollectionParams) WithTimeout(timeout time.Duration) *AuditModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit modify collection params
func (o *AuditModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit modify collection params
func (o *AuditModifyCollectionParams) WithContext(ctx context.Context) *AuditModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit modify collection params
func (o *AuditModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit modify collection params
func (o *AuditModifyCollectionParams) WithHTTPClient(client *http.Client) *AuditModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit modify collection params
func (o *AuditModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChargeQos adds the chargeQos to the audit modify collection params
func (o *AuditModifyCollectionParams) WithChargeQos(chargeQos *bool) *AuditModifyCollectionParams {
	o.SetChargeQos(chargeQos)
	return o
}

// SetChargeQos adds the chargeQos to the audit modify collection params
func (o *AuditModifyCollectionParams) SetChargeQos(chargeQos *bool) {
	o.ChargeQos = chargeQos
}

// WithContinueOnFailure adds the continueOnFailure to the audit modify collection params
func (o *AuditModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *AuditModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the audit modify collection params
func (o *AuditModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithEnabled adds the enabled to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEnabled(enabled *bool) *AuditModifyCollectionParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithEventsAsyncDelete adds the eventsAsyncDelete to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEventsAsyncDelete(eventsAsyncDelete *bool) *AuditModifyCollectionParams {
	o.SetEventsAsyncDelete(eventsAsyncDelete)
	return o
}

// SetEventsAsyncDelete adds the eventsAsyncDelete to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEventsAsyncDelete(eventsAsyncDelete *bool) {
	o.EventsAsyncDelete = eventsAsyncDelete
}

// WithEventsAuditPolicyChange adds the eventsAuditPolicyChange to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEventsAuditPolicyChange(eventsAuditPolicyChange *bool) *AuditModifyCollectionParams {
	o.SetEventsAuditPolicyChange(eventsAuditPolicyChange)
	return o
}

// SetEventsAuditPolicyChange adds the eventsAuditPolicyChange to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEventsAuditPolicyChange(eventsAuditPolicyChange *bool) {
	o.EventsAuditPolicyChange = eventsAuditPolicyChange
}

// WithEventsAuthorizationPolicy adds the eventsAuthorizationPolicy to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEventsAuthorizationPolicy(eventsAuthorizationPolicy *bool) *AuditModifyCollectionParams {
	o.SetEventsAuthorizationPolicy(eventsAuthorizationPolicy)
	return o
}

// SetEventsAuthorizationPolicy adds the eventsAuthorizationPolicy to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEventsAuthorizationPolicy(eventsAuthorizationPolicy *bool) {
	o.EventsAuthorizationPolicy = eventsAuthorizationPolicy
}

// WithEventsCapStaging adds the eventsCapStaging to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEventsCapStaging(eventsCapStaging *bool) *AuditModifyCollectionParams {
	o.SetEventsCapStaging(eventsCapStaging)
	return o
}

// SetEventsCapStaging adds the eventsCapStaging to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEventsCapStaging(eventsCapStaging *bool) {
	o.EventsCapStaging = eventsCapStaging
}

// WithEventsCifsLogonLogoff adds the eventsCifsLogonLogoff to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEventsCifsLogonLogoff(eventsCifsLogonLogoff *bool) *AuditModifyCollectionParams {
	o.SetEventsCifsLogonLogoff(eventsCifsLogonLogoff)
	return o
}

// SetEventsCifsLogonLogoff adds the eventsCifsLogonLogoff to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEventsCifsLogonLogoff(eventsCifsLogonLogoff *bool) {
	o.EventsCifsLogonLogoff = eventsCifsLogonLogoff
}

// WithEventsFileOperations adds the eventsFileOperations to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEventsFileOperations(eventsFileOperations *bool) *AuditModifyCollectionParams {
	o.SetEventsFileOperations(eventsFileOperations)
	return o
}

// SetEventsFileOperations adds the eventsFileOperations to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEventsFileOperations(eventsFileOperations *bool) {
	o.EventsFileOperations = eventsFileOperations
}

// WithEventsFileShare adds the eventsFileShare to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEventsFileShare(eventsFileShare *bool) *AuditModifyCollectionParams {
	o.SetEventsFileShare(eventsFileShare)
	return o
}

// SetEventsFileShare adds the eventsFileShare to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEventsFileShare(eventsFileShare *bool) {
	o.EventsFileShare = eventsFileShare
}

// WithEventsSecurityGroup adds the eventsSecurityGroup to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEventsSecurityGroup(eventsSecurityGroup *bool) *AuditModifyCollectionParams {
	o.SetEventsSecurityGroup(eventsSecurityGroup)
	return o
}

// SetEventsSecurityGroup adds the eventsSecurityGroup to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEventsSecurityGroup(eventsSecurityGroup *bool) {
	o.EventsSecurityGroup = eventsSecurityGroup
}

// WithEventsUserAccount adds the eventsUserAccount to the audit modify collection params
func (o *AuditModifyCollectionParams) WithEventsUserAccount(eventsUserAccount *bool) *AuditModifyCollectionParams {
	o.SetEventsUserAccount(eventsUserAccount)
	return o
}

// SetEventsUserAccount adds the eventsUserAccount to the audit modify collection params
func (o *AuditModifyCollectionParams) SetEventsUserAccount(eventsUserAccount *bool) {
	o.EventsUserAccount = eventsUserAccount
}

// WithGuarantee adds the guarantee to the audit modify collection params
func (o *AuditModifyCollectionParams) WithGuarantee(guarantee *bool) *AuditModifyCollectionParams {
	o.SetGuarantee(guarantee)
	return o
}

// SetGuarantee adds the guarantee to the audit modify collection params
func (o *AuditModifyCollectionParams) SetGuarantee(guarantee *bool) {
	o.Guarantee = guarantee
}

// WithInfo adds the info to the audit modify collection params
func (o *AuditModifyCollectionParams) WithInfo(info AuditModifyCollectionBody) *AuditModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the audit modify collection params
func (o *AuditModifyCollectionParams) SetInfo(info AuditModifyCollectionBody) {
	o.Info = info
}

// WithLogFormat adds the logFormat to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogFormat(logFormat *string) *AuditModifyCollectionParams {
	o.SetLogFormat(logFormat)
	return o
}

// SetLogFormat adds the logFormat to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogFormat(logFormat *string) {
	o.LogFormat = logFormat
}

// WithLogRetentionCount adds the logRetentionCount to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogRetentionCount(logRetentionCount *int64) *AuditModifyCollectionParams {
	o.SetLogRetentionCount(logRetentionCount)
	return o
}

// SetLogRetentionCount adds the logRetentionCount to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogRetentionCount(logRetentionCount *int64) {
	o.LogRetentionCount = logRetentionCount
}

// WithLogRetentionDuration adds the logRetentionDuration to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogRetentionDuration(logRetentionDuration *string) *AuditModifyCollectionParams {
	o.SetLogRetentionDuration(logRetentionDuration)
	return o
}

// SetLogRetentionDuration adds the logRetentionDuration to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogRetentionDuration(logRetentionDuration *string) {
	o.LogRetentionDuration = logRetentionDuration
}

// WithLogRotationScheduleDays adds the logRotationScheduleDays to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogRotationScheduleDays(logRotationScheduleDays *int64) *AuditModifyCollectionParams {
	o.SetLogRotationScheduleDays(logRotationScheduleDays)
	return o
}

// SetLogRotationScheduleDays adds the logRotationScheduleDays to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogRotationScheduleDays(logRotationScheduleDays *int64) {
	o.LogRotationScheduleDays = logRotationScheduleDays
}

// WithLogRotationScheduleHours adds the logRotationScheduleHours to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogRotationScheduleHours(logRotationScheduleHours *int64) *AuditModifyCollectionParams {
	o.SetLogRotationScheduleHours(logRotationScheduleHours)
	return o
}

// SetLogRotationScheduleHours adds the logRotationScheduleHours to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogRotationScheduleHours(logRotationScheduleHours *int64) {
	o.LogRotationScheduleHours = logRotationScheduleHours
}

// WithLogRotationScheduleMinutes adds the logRotationScheduleMinutes to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogRotationScheduleMinutes(logRotationScheduleMinutes *int64) *AuditModifyCollectionParams {
	o.SetLogRotationScheduleMinutes(logRotationScheduleMinutes)
	return o
}

// SetLogRotationScheduleMinutes adds the logRotationScheduleMinutes to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogRotationScheduleMinutes(logRotationScheduleMinutes *int64) {
	o.LogRotationScheduleMinutes = logRotationScheduleMinutes
}

// WithLogRotationScheduleMonths adds the logRotationScheduleMonths to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogRotationScheduleMonths(logRotationScheduleMonths *int64) *AuditModifyCollectionParams {
	o.SetLogRotationScheduleMonths(logRotationScheduleMonths)
	return o
}

// SetLogRotationScheduleMonths adds the logRotationScheduleMonths to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogRotationScheduleMonths(logRotationScheduleMonths *int64) {
	o.LogRotationScheduleMonths = logRotationScheduleMonths
}

// WithLogRotationScheduleWeekdays adds the logRotationScheduleWeekdays to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogRotationScheduleWeekdays(logRotationScheduleWeekdays *int64) *AuditModifyCollectionParams {
	o.SetLogRotationScheduleWeekdays(logRotationScheduleWeekdays)
	return o
}

// SetLogRotationScheduleWeekdays adds the logRotationScheduleWeekdays to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogRotationScheduleWeekdays(logRotationScheduleWeekdays *int64) {
	o.LogRotationScheduleWeekdays = logRotationScheduleWeekdays
}

// WithLogRotationSize adds the logRotationSize to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogRotationSize(logRotationSize *int64) *AuditModifyCollectionParams {
	o.SetLogRotationSize(logRotationSize)
	return o
}

// SetLogRotationSize adds the logRotationSize to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogRotationSize(logRotationSize *int64) {
	o.LogRotationSize = logRotationSize
}

// WithLogPath adds the logPath to the audit modify collection params
func (o *AuditModifyCollectionParams) WithLogPath(logPath *string) *AuditModifyCollectionParams {
	o.SetLogPath(logPath)
	return o
}

// SetLogPath adds the logPath to the audit modify collection params
func (o *AuditModifyCollectionParams) SetLogPath(logPath *string) {
	o.LogPath = logPath
}

// WithReturnRecords adds the returnRecords to the audit modify collection params
func (o *AuditModifyCollectionParams) WithReturnRecords(returnRecords *bool) *AuditModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the audit modify collection params
func (o *AuditModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the audit modify collection params
func (o *AuditModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *AuditModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the audit modify collection params
func (o *AuditModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the audit modify collection params
func (o *AuditModifyCollectionParams) WithSerialRecords(serialRecords *bool) *AuditModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the audit modify collection params
func (o *AuditModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the audit modify collection params
func (o *AuditModifyCollectionParams) WithSvmName(svmName *string) *AuditModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the audit modify collection params
func (o *AuditModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the audit modify collection params
func (o *AuditModifyCollectionParams) WithSvmUUID(svmUUID *string) *AuditModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the audit modify collection params
func (o *AuditModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *AuditModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChargeQos != nil {

		// query param charge_qos
		var qrChargeQos bool

		if o.ChargeQos != nil {
			qrChargeQos = *o.ChargeQos
		}
		qChargeQos := swag.FormatBool(qrChargeQos)
		if qChargeQos != "" {

			if err := r.SetQueryParam("charge_qos", qChargeQos); err != nil {
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

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.EventsAsyncDelete != nil {

		// query param events.async_delete
		var qrEventsAsyncDelete bool

		if o.EventsAsyncDelete != nil {
			qrEventsAsyncDelete = *o.EventsAsyncDelete
		}
		qEventsAsyncDelete := swag.FormatBool(qrEventsAsyncDelete)
		if qEventsAsyncDelete != "" {

			if err := r.SetQueryParam("events.async_delete", qEventsAsyncDelete); err != nil {
				return err
			}
		}
	}

	if o.EventsAuditPolicyChange != nil {

		// query param events.audit_policy_change
		var qrEventsAuditPolicyChange bool

		if o.EventsAuditPolicyChange != nil {
			qrEventsAuditPolicyChange = *o.EventsAuditPolicyChange
		}
		qEventsAuditPolicyChange := swag.FormatBool(qrEventsAuditPolicyChange)
		if qEventsAuditPolicyChange != "" {

			if err := r.SetQueryParam("events.audit_policy_change", qEventsAuditPolicyChange); err != nil {
				return err
			}
		}
	}

	if o.EventsAuthorizationPolicy != nil {

		// query param events.authorization_policy
		var qrEventsAuthorizationPolicy bool

		if o.EventsAuthorizationPolicy != nil {
			qrEventsAuthorizationPolicy = *o.EventsAuthorizationPolicy
		}
		qEventsAuthorizationPolicy := swag.FormatBool(qrEventsAuthorizationPolicy)
		if qEventsAuthorizationPolicy != "" {

			if err := r.SetQueryParam("events.authorization_policy", qEventsAuthorizationPolicy); err != nil {
				return err
			}
		}
	}

	if o.EventsCapStaging != nil {

		// query param events.cap_staging
		var qrEventsCapStaging bool

		if o.EventsCapStaging != nil {
			qrEventsCapStaging = *o.EventsCapStaging
		}
		qEventsCapStaging := swag.FormatBool(qrEventsCapStaging)
		if qEventsCapStaging != "" {

			if err := r.SetQueryParam("events.cap_staging", qEventsCapStaging); err != nil {
				return err
			}
		}
	}

	if o.EventsCifsLogonLogoff != nil {

		// query param events.cifs_logon_logoff
		var qrEventsCifsLogonLogoff bool

		if o.EventsCifsLogonLogoff != nil {
			qrEventsCifsLogonLogoff = *o.EventsCifsLogonLogoff
		}
		qEventsCifsLogonLogoff := swag.FormatBool(qrEventsCifsLogonLogoff)
		if qEventsCifsLogonLogoff != "" {

			if err := r.SetQueryParam("events.cifs_logon_logoff", qEventsCifsLogonLogoff); err != nil {
				return err
			}
		}
	}

	if o.EventsFileOperations != nil {

		// query param events.file_operations
		var qrEventsFileOperations bool

		if o.EventsFileOperations != nil {
			qrEventsFileOperations = *o.EventsFileOperations
		}
		qEventsFileOperations := swag.FormatBool(qrEventsFileOperations)
		if qEventsFileOperations != "" {

			if err := r.SetQueryParam("events.file_operations", qEventsFileOperations); err != nil {
				return err
			}
		}
	}

	if o.EventsFileShare != nil {

		// query param events.file_share
		var qrEventsFileShare bool

		if o.EventsFileShare != nil {
			qrEventsFileShare = *o.EventsFileShare
		}
		qEventsFileShare := swag.FormatBool(qrEventsFileShare)
		if qEventsFileShare != "" {

			if err := r.SetQueryParam("events.file_share", qEventsFileShare); err != nil {
				return err
			}
		}
	}

	if o.EventsSecurityGroup != nil {

		// query param events.security_group
		var qrEventsSecurityGroup bool

		if o.EventsSecurityGroup != nil {
			qrEventsSecurityGroup = *o.EventsSecurityGroup
		}
		qEventsSecurityGroup := swag.FormatBool(qrEventsSecurityGroup)
		if qEventsSecurityGroup != "" {

			if err := r.SetQueryParam("events.security_group", qEventsSecurityGroup); err != nil {
				return err
			}
		}
	}

	if o.EventsUserAccount != nil {

		// query param events.user_account
		var qrEventsUserAccount bool

		if o.EventsUserAccount != nil {
			qrEventsUserAccount = *o.EventsUserAccount
		}
		qEventsUserAccount := swag.FormatBool(qrEventsUserAccount)
		if qEventsUserAccount != "" {

			if err := r.SetQueryParam("events.user_account", qEventsUserAccount); err != nil {
				return err
			}
		}
	}

	if o.Guarantee != nil {

		// query param guarantee
		var qrGuarantee bool

		if o.Guarantee != nil {
			qrGuarantee = *o.Guarantee
		}
		qGuarantee := swag.FormatBool(qrGuarantee)
		if qGuarantee != "" {

			if err := r.SetQueryParam("guarantee", qGuarantee); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.LogFormat != nil {

		// query param log.format
		var qrLogFormat string

		if o.LogFormat != nil {
			qrLogFormat = *o.LogFormat
		}
		qLogFormat := qrLogFormat
		if qLogFormat != "" {

			if err := r.SetQueryParam("log.format", qLogFormat); err != nil {
				return err
			}
		}
	}

	if o.LogRetentionCount != nil {

		// query param log.retention.count
		var qrLogRetentionCount int64

		if o.LogRetentionCount != nil {
			qrLogRetentionCount = *o.LogRetentionCount
		}
		qLogRetentionCount := swag.FormatInt64(qrLogRetentionCount)
		if qLogRetentionCount != "" {

			if err := r.SetQueryParam("log.retention.count", qLogRetentionCount); err != nil {
				return err
			}
		}
	}

	if o.LogRetentionDuration != nil {

		// query param log.retention.duration
		var qrLogRetentionDuration string

		if o.LogRetentionDuration != nil {
			qrLogRetentionDuration = *o.LogRetentionDuration
		}
		qLogRetentionDuration := qrLogRetentionDuration
		if qLogRetentionDuration != "" {

			if err := r.SetQueryParam("log.retention.duration", qLogRetentionDuration); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleDays != nil {

		// query param log.rotation.schedule.days
		var qrLogRotationScheduleDays int64

		if o.LogRotationScheduleDays != nil {
			qrLogRotationScheduleDays = *o.LogRotationScheduleDays
		}
		qLogRotationScheduleDays := swag.FormatInt64(qrLogRotationScheduleDays)
		if qLogRotationScheduleDays != "" {

			if err := r.SetQueryParam("log.rotation.schedule.days", qLogRotationScheduleDays); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleHours != nil {

		// query param log.rotation.schedule.hours
		var qrLogRotationScheduleHours int64

		if o.LogRotationScheduleHours != nil {
			qrLogRotationScheduleHours = *o.LogRotationScheduleHours
		}
		qLogRotationScheduleHours := swag.FormatInt64(qrLogRotationScheduleHours)
		if qLogRotationScheduleHours != "" {

			if err := r.SetQueryParam("log.rotation.schedule.hours", qLogRotationScheduleHours); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleMinutes != nil {

		// query param log.rotation.schedule.minutes
		var qrLogRotationScheduleMinutes int64

		if o.LogRotationScheduleMinutes != nil {
			qrLogRotationScheduleMinutes = *o.LogRotationScheduleMinutes
		}
		qLogRotationScheduleMinutes := swag.FormatInt64(qrLogRotationScheduleMinutes)
		if qLogRotationScheduleMinutes != "" {

			if err := r.SetQueryParam("log.rotation.schedule.minutes", qLogRotationScheduleMinutes); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleMonths != nil {

		// query param log.rotation.schedule.months
		var qrLogRotationScheduleMonths int64

		if o.LogRotationScheduleMonths != nil {
			qrLogRotationScheduleMonths = *o.LogRotationScheduleMonths
		}
		qLogRotationScheduleMonths := swag.FormatInt64(qrLogRotationScheduleMonths)
		if qLogRotationScheduleMonths != "" {

			if err := r.SetQueryParam("log.rotation.schedule.months", qLogRotationScheduleMonths); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleWeekdays != nil {

		// query param log.rotation.schedule.weekdays
		var qrLogRotationScheduleWeekdays int64

		if o.LogRotationScheduleWeekdays != nil {
			qrLogRotationScheduleWeekdays = *o.LogRotationScheduleWeekdays
		}
		qLogRotationScheduleWeekdays := swag.FormatInt64(qrLogRotationScheduleWeekdays)
		if qLogRotationScheduleWeekdays != "" {

			if err := r.SetQueryParam("log.rotation.schedule.weekdays", qLogRotationScheduleWeekdays); err != nil {
				return err
			}
		}
	}

	if o.LogRotationSize != nil {

		// query param log.rotation.size
		var qrLogRotationSize int64

		if o.LogRotationSize != nil {
			qrLogRotationSize = *o.LogRotationSize
		}
		qLogRotationSize := swag.FormatInt64(qrLogRotationSize)
		if qLogRotationSize != "" {

			if err := r.SetQueryParam("log.rotation.size", qLogRotationSize); err != nil {
				return err
			}
		}
	}

	if o.LogPath != nil {

		// query param log_path
		var qrLogPath string

		if o.LogPath != nil {
			qrLogPath = *o.LogPath
		}
		qLogPath := qrLogPath
		if qLogPath != "" {

			if err := r.SetQueryParam("log_path", qLogPath); err != nil {
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
