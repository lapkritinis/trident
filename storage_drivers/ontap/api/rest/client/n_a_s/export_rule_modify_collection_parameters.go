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

// NewExportRuleModifyCollectionParams creates a new ExportRuleModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportRuleModifyCollectionParams() *ExportRuleModifyCollectionParams {
	return &ExportRuleModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportRuleModifyCollectionParamsWithTimeout creates a new ExportRuleModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewExportRuleModifyCollectionParamsWithTimeout(timeout time.Duration) *ExportRuleModifyCollectionParams {
	return &ExportRuleModifyCollectionParams{
		timeout: timeout,
	}
}

// NewExportRuleModifyCollectionParamsWithContext creates a new ExportRuleModifyCollectionParams object
// with the ability to set a context for a request.
func NewExportRuleModifyCollectionParamsWithContext(ctx context.Context) *ExportRuleModifyCollectionParams {
	return &ExportRuleModifyCollectionParams{
		Context: ctx,
	}
}

// NewExportRuleModifyCollectionParamsWithHTTPClient creates a new ExportRuleModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportRuleModifyCollectionParamsWithHTTPClient(client *http.Client) *ExportRuleModifyCollectionParams {
	return &ExportRuleModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
ExportRuleModifyCollectionParams contains all the parameters to send to the API endpoint

	for the export rule modify collection operation.

	Typically these are written to a http.Request.
*/
type ExportRuleModifyCollectionParams struct {

	/* AllowDeviceCreation.

	   Filter by allow_device_creation
	*/
	AllowDeviceCreation *bool

	/* AllowSuid.

	   Filter by allow_suid
	*/
	AllowSuid *bool

	/* AnonymousUser.

	   Filter by anonymous_user
	*/
	AnonymousUser *string

	/* ChownMode.

	   Filter by chown_mode
	*/
	ChownMode *string

	/* ClientsMatch.

	   Filter by clients.match
	*/
	ClientsMatch *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Index.

	   Filter by index
	*/
	Index *int64

	/* Info.

	   Info specification
	*/
	Info ExportRuleModifyCollectionBody

	/* NewIndex.

	   New Export Rule Index
	*/
	NewIndex *int64

	/* NtfsUnixSecurity.

	   Filter by ntfs_unix_security
	*/
	NtfsUnixSecurity *string

	/* PolicyID.

	   Export Policy ID
	*/
	PolicyID int64

	/* PolicyName.

	   Filter by policy.name
	*/
	PolicyName *string

	/* Protocols.

	   Filter by protocols
	*/
	Protocols *string

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

	/* RoRule.

	   Filter by ro_rule
	*/
	RoRule *string

	/* RwRule.

	   Filter by rw_rule
	*/
	RwRule *string

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* Superuser.

	   Filter by superuser
	*/
	Superuser *string

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

// WithDefaults hydrates default values in the export rule modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleModifyCollectionParams) WithDefaults() *ExportRuleModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export rule modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := ExportRuleModifyCollectionParams{
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

// WithTimeout adds the timeout to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithTimeout(timeout time.Duration) *ExportRuleModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithContext(ctx context.Context) *ExportRuleModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithHTTPClient(client *http.Client) *ExportRuleModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowDeviceCreation adds the allowDeviceCreation to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithAllowDeviceCreation(allowDeviceCreation *bool) *ExportRuleModifyCollectionParams {
	o.SetAllowDeviceCreation(allowDeviceCreation)
	return o
}

// SetAllowDeviceCreation adds the allowDeviceCreation to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetAllowDeviceCreation(allowDeviceCreation *bool) {
	o.AllowDeviceCreation = allowDeviceCreation
}

// WithAllowSuid adds the allowSuid to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithAllowSuid(allowSuid *bool) *ExportRuleModifyCollectionParams {
	o.SetAllowSuid(allowSuid)
	return o
}

// SetAllowSuid adds the allowSuid to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetAllowSuid(allowSuid *bool) {
	o.AllowSuid = allowSuid
}

// WithAnonymousUser adds the anonymousUser to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithAnonymousUser(anonymousUser *string) *ExportRuleModifyCollectionParams {
	o.SetAnonymousUser(anonymousUser)
	return o
}

// SetAnonymousUser adds the anonymousUser to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetAnonymousUser(anonymousUser *string) {
	o.AnonymousUser = anonymousUser
}

// WithChownMode adds the chownMode to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithChownMode(chownMode *string) *ExportRuleModifyCollectionParams {
	o.SetChownMode(chownMode)
	return o
}

// SetChownMode adds the chownMode to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetChownMode(chownMode *string) {
	o.ChownMode = chownMode
}

// WithClientsMatch adds the clientsMatch to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithClientsMatch(clientsMatch *string) *ExportRuleModifyCollectionParams {
	o.SetClientsMatch(clientsMatch)
	return o
}

// SetClientsMatch adds the clientsMatch to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetClientsMatch(clientsMatch *string) {
	o.ClientsMatch = clientsMatch
}

// WithContinueOnFailure adds the continueOnFailure to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *ExportRuleModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithIndex adds the index to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithIndex(index *int64) *ExportRuleModifyCollectionParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetIndex(index *int64) {
	o.Index = index
}

// WithInfo adds the info to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithInfo(info ExportRuleModifyCollectionBody) *ExportRuleModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetInfo(info ExportRuleModifyCollectionBody) {
	o.Info = info
}

// WithNewIndex adds the newIndex to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithNewIndex(newIndex *int64) *ExportRuleModifyCollectionParams {
	o.SetNewIndex(newIndex)
	return o
}

// SetNewIndex adds the newIndex to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetNewIndex(newIndex *int64) {
	o.NewIndex = newIndex
}

// WithNtfsUnixSecurity adds the ntfsUnixSecurity to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithNtfsUnixSecurity(ntfsUnixSecurity *string) *ExportRuleModifyCollectionParams {
	o.SetNtfsUnixSecurity(ntfsUnixSecurity)
	return o
}

// SetNtfsUnixSecurity adds the ntfsUnixSecurity to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetNtfsUnixSecurity(ntfsUnixSecurity *string) {
	o.NtfsUnixSecurity = ntfsUnixSecurity
}

// WithPolicyID adds the policyID to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithPolicyID(policyID int64) *ExportRuleModifyCollectionParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WithPolicyName adds the policyName to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithPolicyName(policyName *string) *ExportRuleModifyCollectionParams {
	o.SetPolicyName(policyName)
	return o
}

// SetPolicyName adds the policyName to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetPolicyName(policyName *string) {
	o.PolicyName = policyName
}

// WithProtocols adds the protocols to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithProtocols(protocols *string) *ExportRuleModifyCollectionParams {
	o.SetProtocols(protocols)
	return o
}

// SetProtocols adds the protocols to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetProtocols(protocols *string) {
	o.Protocols = protocols
}

// WithReturnRecords adds the returnRecords to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithReturnRecords(returnRecords *bool) *ExportRuleModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *ExportRuleModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithRoRule adds the roRule to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithRoRule(roRule *string) *ExportRuleModifyCollectionParams {
	o.SetRoRule(roRule)
	return o
}

// SetRoRule adds the roRule to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetRoRule(roRule *string) {
	o.RoRule = roRule
}

// WithRwRule adds the rwRule to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithRwRule(rwRule *string) *ExportRuleModifyCollectionParams {
	o.SetRwRule(rwRule)
	return o
}

// SetRwRule adds the rwRule to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetRwRule(rwRule *string) {
	o.RwRule = rwRule
}

// WithSerialRecords adds the serialRecords to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithSerialRecords(serialRecords *bool) *ExportRuleModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSuperuser adds the superuser to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithSuperuser(superuser *string) *ExportRuleModifyCollectionParams {
	o.SetSuperuser(superuser)
	return o
}

// SetSuperuser adds the superuser to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetSuperuser(superuser *string) {
	o.Superuser = superuser
}

// WithSvmName adds the svmName to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithSvmName(svmName *string) *ExportRuleModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) WithSvmUUID(svmUUID *string) *ExportRuleModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the export rule modify collection params
func (o *ExportRuleModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ExportRuleModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowDeviceCreation != nil {

		// query param allow_device_creation
		var qrAllowDeviceCreation bool

		if o.AllowDeviceCreation != nil {
			qrAllowDeviceCreation = *o.AllowDeviceCreation
		}
		qAllowDeviceCreation := swag.FormatBool(qrAllowDeviceCreation)
		if qAllowDeviceCreation != "" {

			if err := r.SetQueryParam("allow_device_creation", qAllowDeviceCreation); err != nil {
				return err
			}
		}
	}

	if o.AllowSuid != nil {

		// query param allow_suid
		var qrAllowSuid bool

		if o.AllowSuid != nil {
			qrAllowSuid = *o.AllowSuid
		}
		qAllowSuid := swag.FormatBool(qrAllowSuid)
		if qAllowSuid != "" {

			if err := r.SetQueryParam("allow_suid", qAllowSuid); err != nil {
				return err
			}
		}
	}

	if o.AnonymousUser != nil {

		// query param anonymous_user
		var qrAnonymousUser string

		if o.AnonymousUser != nil {
			qrAnonymousUser = *o.AnonymousUser
		}
		qAnonymousUser := qrAnonymousUser
		if qAnonymousUser != "" {

			if err := r.SetQueryParam("anonymous_user", qAnonymousUser); err != nil {
				return err
			}
		}
	}

	if o.ChownMode != nil {

		// query param chown_mode
		var qrChownMode string

		if o.ChownMode != nil {
			qrChownMode = *o.ChownMode
		}
		qChownMode := qrChownMode
		if qChownMode != "" {

			if err := r.SetQueryParam("chown_mode", qChownMode); err != nil {
				return err
			}
		}
	}

	if o.ClientsMatch != nil {

		// query param clients.match
		var qrClientsMatch string

		if o.ClientsMatch != nil {
			qrClientsMatch = *o.ClientsMatch
		}
		qClientsMatch := qrClientsMatch
		if qClientsMatch != "" {

			if err := r.SetQueryParam("clients.match", qClientsMatch); err != nil {
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

	if o.Index != nil {

		// query param index
		var qrIndex int64

		if o.Index != nil {
			qrIndex = *o.Index
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.NewIndex != nil {

		// query param new_index
		var qrNewIndex int64

		if o.NewIndex != nil {
			qrNewIndex = *o.NewIndex
		}
		qNewIndex := swag.FormatInt64(qrNewIndex)
		if qNewIndex != "" {

			if err := r.SetQueryParam("new_index", qNewIndex); err != nil {
				return err
			}
		}
	}

	if o.NtfsUnixSecurity != nil {

		// query param ntfs_unix_security
		var qrNtfsUnixSecurity string

		if o.NtfsUnixSecurity != nil {
			qrNtfsUnixSecurity = *o.NtfsUnixSecurity
		}
		qNtfsUnixSecurity := qrNtfsUnixSecurity
		if qNtfsUnixSecurity != "" {

			if err := r.SetQueryParam("ntfs_unix_security", qNtfsUnixSecurity); err != nil {
				return err
			}
		}
	}

	// path param policy.id
	if err := r.SetPathParam("policy.id", swag.FormatInt64(o.PolicyID)); err != nil {
		return err
	}

	if o.PolicyName != nil {

		// query param policy.name
		var qrPolicyName string

		if o.PolicyName != nil {
			qrPolicyName = *o.PolicyName
		}
		qPolicyName := qrPolicyName
		if qPolicyName != "" {

			if err := r.SetQueryParam("policy.name", qPolicyName); err != nil {
				return err
			}
		}
	}

	if o.Protocols != nil {

		// query param protocols
		var qrProtocols string

		if o.Protocols != nil {
			qrProtocols = *o.Protocols
		}
		qProtocols := qrProtocols
		if qProtocols != "" {

			if err := r.SetQueryParam("protocols", qProtocols); err != nil {
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

	if o.RoRule != nil {

		// query param ro_rule
		var qrRoRule string

		if o.RoRule != nil {
			qrRoRule = *o.RoRule
		}
		qRoRule := qrRoRule
		if qRoRule != "" {

			if err := r.SetQueryParam("ro_rule", qRoRule); err != nil {
				return err
			}
		}
	}

	if o.RwRule != nil {

		// query param rw_rule
		var qrRwRule string

		if o.RwRule != nil {
			qrRwRule = *o.RwRule
		}
		qRwRule := qrRwRule
		if qRwRule != "" {

			if err := r.SetQueryParam("rw_rule", qRwRule); err != nil {
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

	if o.Superuser != nil {

		// query param superuser
		var qrSuperuser string

		if o.Superuser != nil {
			qrSuperuser = *o.Superuser
		}
		qSuperuser := qrSuperuser
		if qSuperuser != "" {

			if err := r.SetQueryParam("superuser", qSuperuser); err != nil {
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
