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

// NewGcpKmsCollectionGetParams creates a new GcpKmsCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGcpKmsCollectionGetParams() *GcpKmsCollectionGetParams {
	return &GcpKmsCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGcpKmsCollectionGetParamsWithTimeout creates a new GcpKmsCollectionGetParams object
// with the ability to set a timeout on a request.
func NewGcpKmsCollectionGetParamsWithTimeout(timeout time.Duration) *GcpKmsCollectionGetParams {
	return &GcpKmsCollectionGetParams{
		timeout: timeout,
	}
}

// NewGcpKmsCollectionGetParamsWithContext creates a new GcpKmsCollectionGetParams object
// with the ability to set a context for a request.
func NewGcpKmsCollectionGetParamsWithContext(ctx context.Context) *GcpKmsCollectionGetParams {
	return &GcpKmsCollectionGetParams{
		Context: ctx,
	}
}

// NewGcpKmsCollectionGetParamsWithHTTPClient creates a new GcpKmsCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGcpKmsCollectionGetParamsWithHTTPClient(client *http.Client) *GcpKmsCollectionGetParams {
	return &GcpKmsCollectionGetParams{
		HTTPClient: client,
	}
}

/*
GcpKmsCollectionGetParams contains all the parameters to send to the API endpoint

	for the gcp kms collection get operation.

	Typically these are written to a http.Request.
*/
type GcpKmsCollectionGetParams struct {

	/* CallerAccount.

	   Filter by caller_account
	*/
	CallerAccount *string

	/* CloudkmsHost.

	   Filter by cloudkms_host
	*/
	CloudkmsHost *string

	/* EkmipReachabilityCode.

	   Filter by ekmip_reachability.code
	*/
	EkmipReachabilityCode *string

	/* EkmipReachabilityMessage.

	   Filter by ekmip_reachability.message
	*/
	EkmipReachabilityMessage *string

	/* EkmipReachabilityNodeName.

	   Filter by ekmip_reachability.node.name
	*/
	EkmipReachabilityNodeName *string

	/* EkmipReachabilityNodeUUID.

	   Filter by ekmip_reachability.node.uuid
	*/
	EkmipReachabilityNodeUUID *string

	/* EkmipReachabilityReachable.

	   Filter by ekmip_reachability.reachable
	*/
	EkmipReachabilityReachable *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* GoogleReachabilityCode.

	   Filter by google_reachability.code
	*/
	GoogleReachabilityCode *string

	/* GoogleReachabilityMessage.

	   Filter by google_reachability.message
	*/
	GoogleReachabilityMessage *string

	/* GoogleReachabilityReachable.

	   Filter by google_reachability.reachable
	*/
	GoogleReachabilityReachable *bool

	/* KeyName.

	   Filter by key_name
	*/
	KeyName *string

	/* KeyRingLocation.

	   Filter by key_ring_location
	*/
	KeyRingLocation *string

	/* KeyRingName.

	   Filter by key_ring_name
	*/
	KeyRingName *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OauthHost.

	   Filter by oauth_host
	*/
	OauthHost *string

	/* OauthURL.

	   Filter by oauth_url
	*/
	OauthURL *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Port.

	   Filter by port
	*/
	Port *int64

	/* PrivilegedAccount.

	   Filter by privileged_account
	*/
	PrivilegedAccount *string

	/* ProjectID.

	   Filter by project_id
	*/
	ProjectID *string

	/* ProxyHost.

	   Filter by proxy_host
	*/
	ProxyHost *string

	/* ProxyPort.

	   Filter by proxy_port
	*/
	ProxyPort *int64

	/* ProxyType.

	   Filter by proxy_type
	*/
	ProxyType *string

	/* ProxyUsername.

	   Filter by proxy_username
	*/
	ProxyUsername *string

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

	/* StateClusterState.

	   Filter by state.cluster_state
	*/
	StateClusterState *bool

	/* StateCode.

	   Filter by state.code
	*/
	StateCode *string

	/* StateMessage.

	   Filter by state.message
	*/
	StateMessage *string

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

	/* VerifyHost.

	   Filter by verify_host
	*/
	VerifyHost *bool

	/* VerifyIP.

	   Filter by verify_ip
	*/
	VerifyIP *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gcp kms collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsCollectionGetParams) WithDefaults() *GcpKmsCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gcp kms collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := GcpKmsCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithTimeout(timeout time.Duration) *GcpKmsCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithContext(ctx context.Context) *GcpKmsCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithHTTPClient(client *http.Client) *GcpKmsCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallerAccount adds the callerAccount to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithCallerAccount(callerAccount *string) *GcpKmsCollectionGetParams {
	o.SetCallerAccount(callerAccount)
	return o
}

// SetCallerAccount adds the callerAccount to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetCallerAccount(callerAccount *string) {
	o.CallerAccount = callerAccount
}

// WithCloudkmsHost adds the cloudkmsHost to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithCloudkmsHost(cloudkmsHost *string) *GcpKmsCollectionGetParams {
	o.SetCloudkmsHost(cloudkmsHost)
	return o
}

// SetCloudkmsHost adds the cloudkmsHost to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetCloudkmsHost(cloudkmsHost *string) {
	o.CloudkmsHost = cloudkmsHost
}

// WithEkmipReachabilityCode adds the ekmipReachabilityCode to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithEkmipReachabilityCode(ekmipReachabilityCode *string) *GcpKmsCollectionGetParams {
	o.SetEkmipReachabilityCode(ekmipReachabilityCode)
	return o
}

// SetEkmipReachabilityCode adds the ekmipReachabilityCode to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetEkmipReachabilityCode(ekmipReachabilityCode *string) {
	o.EkmipReachabilityCode = ekmipReachabilityCode
}

// WithEkmipReachabilityMessage adds the ekmipReachabilityMessage to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithEkmipReachabilityMessage(ekmipReachabilityMessage *string) *GcpKmsCollectionGetParams {
	o.SetEkmipReachabilityMessage(ekmipReachabilityMessage)
	return o
}

// SetEkmipReachabilityMessage adds the ekmipReachabilityMessage to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetEkmipReachabilityMessage(ekmipReachabilityMessage *string) {
	o.EkmipReachabilityMessage = ekmipReachabilityMessage
}

// WithEkmipReachabilityNodeName adds the ekmipReachabilityNodeName to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithEkmipReachabilityNodeName(ekmipReachabilityNodeName *string) *GcpKmsCollectionGetParams {
	o.SetEkmipReachabilityNodeName(ekmipReachabilityNodeName)
	return o
}

// SetEkmipReachabilityNodeName adds the ekmipReachabilityNodeName to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetEkmipReachabilityNodeName(ekmipReachabilityNodeName *string) {
	o.EkmipReachabilityNodeName = ekmipReachabilityNodeName
}

// WithEkmipReachabilityNodeUUID adds the ekmipReachabilityNodeUUID to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithEkmipReachabilityNodeUUID(ekmipReachabilityNodeUUID *string) *GcpKmsCollectionGetParams {
	o.SetEkmipReachabilityNodeUUID(ekmipReachabilityNodeUUID)
	return o
}

// SetEkmipReachabilityNodeUUID adds the ekmipReachabilityNodeUuid to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetEkmipReachabilityNodeUUID(ekmipReachabilityNodeUUID *string) {
	o.EkmipReachabilityNodeUUID = ekmipReachabilityNodeUUID
}

// WithEkmipReachabilityReachable adds the ekmipReachabilityReachable to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithEkmipReachabilityReachable(ekmipReachabilityReachable *bool) *GcpKmsCollectionGetParams {
	o.SetEkmipReachabilityReachable(ekmipReachabilityReachable)
	return o
}

// SetEkmipReachabilityReachable adds the ekmipReachabilityReachable to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetEkmipReachabilityReachable(ekmipReachabilityReachable *bool) {
	o.EkmipReachabilityReachable = ekmipReachabilityReachable
}

// WithFields adds the fields to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithFields(fields []string) *GcpKmsCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithGoogleReachabilityCode adds the googleReachabilityCode to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithGoogleReachabilityCode(googleReachabilityCode *string) *GcpKmsCollectionGetParams {
	o.SetGoogleReachabilityCode(googleReachabilityCode)
	return o
}

// SetGoogleReachabilityCode adds the googleReachabilityCode to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetGoogleReachabilityCode(googleReachabilityCode *string) {
	o.GoogleReachabilityCode = googleReachabilityCode
}

// WithGoogleReachabilityMessage adds the googleReachabilityMessage to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithGoogleReachabilityMessage(googleReachabilityMessage *string) *GcpKmsCollectionGetParams {
	o.SetGoogleReachabilityMessage(googleReachabilityMessage)
	return o
}

// SetGoogleReachabilityMessage adds the googleReachabilityMessage to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetGoogleReachabilityMessage(googleReachabilityMessage *string) {
	o.GoogleReachabilityMessage = googleReachabilityMessage
}

// WithGoogleReachabilityReachable adds the googleReachabilityReachable to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithGoogleReachabilityReachable(googleReachabilityReachable *bool) *GcpKmsCollectionGetParams {
	o.SetGoogleReachabilityReachable(googleReachabilityReachable)
	return o
}

// SetGoogleReachabilityReachable adds the googleReachabilityReachable to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetGoogleReachabilityReachable(googleReachabilityReachable *bool) {
	o.GoogleReachabilityReachable = googleReachabilityReachable
}

// WithKeyName adds the keyName to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithKeyName(keyName *string) *GcpKmsCollectionGetParams {
	o.SetKeyName(keyName)
	return o
}

// SetKeyName adds the keyName to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetKeyName(keyName *string) {
	o.KeyName = keyName
}

// WithKeyRingLocation adds the keyRingLocation to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithKeyRingLocation(keyRingLocation *string) *GcpKmsCollectionGetParams {
	o.SetKeyRingLocation(keyRingLocation)
	return o
}

// SetKeyRingLocation adds the keyRingLocation to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetKeyRingLocation(keyRingLocation *string) {
	o.KeyRingLocation = keyRingLocation
}

// WithKeyRingName adds the keyRingName to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithKeyRingName(keyRingName *string) *GcpKmsCollectionGetParams {
	o.SetKeyRingName(keyRingName)
	return o
}

// SetKeyRingName adds the keyRingName to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetKeyRingName(keyRingName *string) {
	o.KeyRingName = keyRingName
}

// WithMaxRecords adds the maxRecords to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithMaxRecords(maxRecords *int64) *GcpKmsCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOauthHost adds the oauthHost to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithOauthHost(oauthHost *string) *GcpKmsCollectionGetParams {
	o.SetOauthHost(oauthHost)
	return o
}

// SetOauthHost adds the oauthHost to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetOauthHost(oauthHost *string) {
	o.OauthHost = oauthHost
}

// WithOauthURL adds the oauthURL to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithOauthURL(oauthURL *string) *GcpKmsCollectionGetParams {
	o.SetOauthURL(oauthURL)
	return o
}

// SetOauthURL adds the oauthUrl to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetOauthURL(oauthURL *string) {
	o.OauthURL = oauthURL
}

// WithOrderBy adds the orderBy to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithOrderBy(orderBy []string) *GcpKmsCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPort adds the port to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithPort(port *int64) *GcpKmsCollectionGetParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetPort(port *int64) {
	o.Port = port
}

// WithPrivilegedAccount adds the privilegedAccount to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithPrivilegedAccount(privilegedAccount *string) *GcpKmsCollectionGetParams {
	o.SetPrivilegedAccount(privilegedAccount)
	return o
}

// SetPrivilegedAccount adds the privilegedAccount to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetPrivilegedAccount(privilegedAccount *string) {
	o.PrivilegedAccount = privilegedAccount
}

// WithProjectID adds the projectID to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithProjectID(projectID *string) *GcpKmsCollectionGetParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithProxyHost adds the proxyHost to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithProxyHost(proxyHost *string) *GcpKmsCollectionGetParams {
	o.SetProxyHost(proxyHost)
	return o
}

// SetProxyHost adds the proxyHost to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetProxyHost(proxyHost *string) {
	o.ProxyHost = proxyHost
}

// WithProxyPort adds the proxyPort to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithProxyPort(proxyPort *int64) *GcpKmsCollectionGetParams {
	o.SetProxyPort(proxyPort)
	return o
}

// SetProxyPort adds the proxyPort to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetProxyPort(proxyPort *int64) {
	o.ProxyPort = proxyPort
}

// WithProxyType adds the proxyType to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithProxyType(proxyType *string) *GcpKmsCollectionGetParams {
	o.SetProxyType(proxyType)
	return o
}

// SetProxyType adds the proxyType to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetProxyType(proxyType *string) {
	o.ProxyType = proxyType
}

// WithProxyUsername adds the proxyUsername to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithProxyUsername(proxyUsername *string) *GcpKmsCollectionGetParams {
	o.SetProxyUsername(proxyUsername)
	return o
}

// SetProxyUsername adds the proxyUsername to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetProxyUsername(proxyUsername *string) {
	o.ProxyUsername = proxyUsername
}

// WithReturnRecords adds the returnRecords to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithReturnRecords(returnRecords *bool) *GcpKmsCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *GcpKmsCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScope adds the scope to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithScope(scope *string) *GcpKmsCollectionGetParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithStateClusterState adds the stateClusterState to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithStateClusterState(stateClusterState *bool) *GcpKmsCollectionGetParams {
	o.SetStateClusterState(stateClusterState)
	return o
}

// SetStateClusterState adds the stateClusterState to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetStateClusterState(stateClusterState *bool) {
	o.StateClusterState = stateClusterState
}

// WithStateCode adds the stateCode to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithStateCode(stateCode *string) *GcpKmsCollectionGetParams {
	o.SetStateCode(stateCode)
	return o
}

// SetStateCode adds the stateCode to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetStateCode(stateCode *string) {
	o.StateCode = stateCode
}

// WithStateMessage adds the stateMessage to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithStateMessage(stateMessage *string) *GcpKmsCollectionGetParams {
	o.SetStateMessage(stateMessage)
	return o
}

// SetStateMessage adds the stateMessage to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetStateMessage(stateMessage *string) {
	o.StateMessage = stateMessage
}

// WithSvmName adds the svmName to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithSvmName(svmName *string) *GcpKmsCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithSvmUUID(svmUUID *string) *GcpKmsCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithUUID(uuid *string) *GcpKmsCollectionGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WithVerifyHost adds the verifyHost to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithVerifyHost(verifyHost *bool) *GcpKmsCollectionGetParams {
	o.SetVerifyHost(verifyHost)
	return o
}

// SetVerifyHost adds the verifyHost to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetVerifyHost(verifyHost *bool) {
	o.VerifyHost = verifyHost
}

// WithVerifyIP adds the verifyIP to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) WithVerifyIP(verifyIP *bool) *GcpKmsCollectionGetParams {
	o.SetVerifyIP(verifyIP)
	return o
}

// SetVerifyIP adds the verifyIp to the gcp kms collection get params
func (o *GcpKmsCollectionGetParams) SetVerifyIP(verifyIP *bool) {
	o.VerifyIP = verifyIP
}

// WriteToRequest writes these params to a swagger request
func (o *GcpKmsCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CallerAccount != nil {

		// query param caller_account
		var qrCallerAccount string

		if o.CallerAccount != nil {
			qrCallerAccount = *o.CallerAccount
		}
		qCallerAccount := qrCallerAccount
		if qCallerAccount != "" {

			if err := r.SetQueryParam("caller_account", qCallerAccount); err != nil {
				return err
			}
		}
	}

	if o.CloudkmsHost != nil {

		// query param cloudkms_host
		var qrCloudkmsHost string

		if o.CloudkmsHost != nil {
			qrCloudkmsHost = *o.CloudkmsHost
		}
		qCloudkmsHost := qrCloudkmsHost
		if qCloudkmsHost != "" {

			if err := r.SetQueryParam("cloudkms_host", qCloudkmsHost); err != nil {
				return err
			}
		}
	}

	if o.EkmipReachabilityCode != nil {

		// query param ekmip_reachability.code
		var qrEkmipReachabilityCode string

		if o.EkmipReachabilityCode != nil {
			qrEkmipReachabilityCode = *o.EkmipReachabilityCode
		}
		qEkmipReachabilityCode := qrEkmipReachabilityCode
		if qEkmipReachabilityCode != "" {

			if err := r.SetQueryParam("ekmip_reachability.code", qEkmipReachabilityCode); err != nil {
				return err
			}
		}
	}

	if o.EkmipReachabilityMessage != nil {

		// query param ekmip_reachability.message
		var qrEkmipReachabilityMessage string

		if o.EkmipReachabilityMessage != nil {
			qrEkmipReachabilityMessage = *o.EkmipReachabilityMessage
		}
		qEkmipReachabilityMessage := qrEkmipReachabilityMessage
		if qEkmipReachabilityMessage != "" {

			if err := r.SetQueryParam("ekmip_reachability.message", qEkmipReachabilityMessage); err != nil {
				return err
			}
		}
	}

	if o.EkmipReachabilityNodeName != nil {

		// query param ekmip_reachability.node.name
		var qrEkmipReachabilityNodeName string

		if o.EkmipReachabilityNodeName != nil {
			qrEkmipReachabilityNodeName = *o.EkmipReachabilityNodeName
		}
		qEkmipReachabilityNodeName := qrEkmipReachabilityNodeName
		if qEkmipReachabilityNodeName != "" {

			if err := r.SetQueryParam("ekmip_reachability.node.name", qEkmipReachabilityNodeName); err != nil {
				return err
			}
		}
	}

	if o.EkmipReachabilityNodeUUID != nil {

		// query param ekmip_reachability.node.uuid
		var qrEkmipReachabilityNodeUUID string

		if o.EkmipReachabilityNodeUUID != nil {
			qrEkmipReachabilityNodeUUID = *o.EkmipReachabilityNodeUUID
		}
		qEkmipReachabilityNodeUUID := qrEkmipReachabilityNodeUUID
		if qEkmipReachabilityNodeUUID != "" {

			if err := r.SetQueryParam("ekmip_reachability.node.uuid", qEkmipReachabilityNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.EkmipReachabilityReachable != nil {

		// query param ekmip_reachability.reachable
		var qrEkmipReachabilityReachable bool

		if o.EkmipReachabilityReachable != nil {
			qrEkmipReachabilityReachable = *o.EkmipReachabilityReachable
		}
		qEkmipReachabilityReachable := swag.FormatBool(qrEkmipReachabilityReachable)
		if qEkmipReachabilityReachable != "" {

			if err := r.SetQueryParam("ekmip_reachability.reachable", qEkmipReachabilityReachable); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.GoogleReachabilityCode != nil {

		// query param google_reachability.code
		var qrGoogleReachabilityCode string

		if o.GoogleReachabilityCode != nil {
			qrGoogleReachabilityCode = *o.GoogleReachabilityCode
		}
		qGoogleReachabilityCode := qrGoogleReachabilityCode
		if qGoogleReachabilityCode != "" {

			if err := r.SetQueryParam("google_reachability.code", qGoogleReachabilityCode); err != nil {
				return err
			}
		}
	}

	if o.GoogleReachabilityMessage != nil {

		// query param google_reachability.message
		var qrGoogleReachabilityMessage string

		if o.GoogleReachabilityMessage != nil {
			qrGoogleReachabilityMessage = *o.GoogleReachabilityMessage
		}
		qGoogleReachabilityMessage := qrGoogleReachabilityMessage
		if qGoogleReachabilityMessage != "" {

			if err := r.SetQueryParam("google_reachability.message", qGoogleReachabilityMessage); err != nil {
				return err
			}
		}
	}

	if o.GoogleReachabilityReachable != nil {

		// query param google_reachability.reachable
		var qrGoogleReachabilityReachable bool

		if o.GoogleReachabilityReachable != nil {
			qrGoogleReachabilityReachable = *o.GoogleReachabilityReachable
		}
		qGoogleReachabilityReachable := swag.FormatBool(qrGoogleReachabilityReachable)
		if qGoogleReachabilityReachable != "" {

			if err := r.SetQueryParam("google_reachability.reachable", qGoogleReachabilityReachable); err != nil {
				return err
			}
		}
	}

	if o.KeyName != nil {

		// query param key_name
		var qrKeyName string

		if o.KeyName != nil {
			qrKeyName = *o.KeyName
		}
		qKeyName := qrKeyName
		if qKeyName != "" {

			if err := r.SetQueryParam("key_name", qKeyName); err != nil {
				return err
			}
		}
	}

	if o.KeyRingLocation != nil {

		// query param key_ring_location
		var qrKeyRingLocation string

		if o.KeyRingLocation != nil {
			qrKeyRingLocation = *o.KeyRingLocation
		}
		qKeyRingLocation := qrKeyRingLocation
		if qKeyRingLocation != "" {

			if err := r.SetQueryParam("key_ring_location", qKeyRingLocation); err != nil {
				return err
			}
		}
	}

	if o.KeyRingName != nil {

		// query param key_ring_name
		var qrKeyRingName string

		if o.KeyRingName != nil {
			qrKeyRingName = *o.KeyRingName
		}
		qKeyRingName := qrKeyRingName
		if qKeyRingName != "" {

			if err := r.SetQueryParam("key_ring_name", qKeyRingName); err != nil {
				return err
			}
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OauthHost != nil {

		// query param oauth_host
		var qrOauthHost string

		if o.OauthHost != nil {
			qrOauthHost = *o.OauthHost
		}
		qOauthHost := qrOauthHost
		if qOauthHost != "" {

			if err := r.SetQueryParam("oauth_host", qOauthHost); err != nil {
				return err
			}
		}
	}

	if o.OauthURL != nil {

		// query param oauth_url
		var qrOauthURL string

		if o.OauthURL != nil {
			qrOauthURL = *o.OauthURL
		}
		qOauthURL := qrOauthURL
		if qOauthURL != "" {

			if err := r.SetQueryParam("oauth_url", qOauthURL); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.Port != nil {

		// query param port
		var qrPort int64

		if o.Port != nil {
			qrPort = *o.Port
		}
		qPort := swag.FormatInt64(qrPort)
		if qPort != "" {

			if err := r.SetQueryParam("port", qPort); err != nil {
				return err
			}
		}
	}

	if o.PrivilegedAccount != nil {

		// query param privileged_account
		var qrPrivilegedAccount string

		if o.PrivilegedAccount != nil {
			qrPrivilegedAccount = *o.PrivilegedAccount
		}
		qPrivilegedAccount := qrPrivilegedAccount
		if qPrivilegedAccount != "" {

			if err := r.SetQueryParam("privileged_account", qPrivilegedAccount); err != nil {
				return err
			}
		}
	}

	if o.ProjectID != nil {

		// query param project_id
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("project_id", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.ProxyHost != nil {

		// query param proxy_host
		var qrProxyHost string

		if o.ProxyHost != nil {
			qrProxyHost = *o.ProxyHost
		}
		qProxyHost := qrProxyHost
		if qProxyHost != "" {

			if err := r.SetQueryParam("proxy_host", qProxyHost); err != nil {
				return err
			}
		}
	}

	if o.ProxyPort != nil {

		// query param proxy_port
		var qrProxyPort int64

		if o.ProxyPort != nil {
			qrProxyPort = *o.ProxyPort
		}
		qProxyPort := swag.FormatInt64(qrProxyPort)
		if qProxyPort != "" {

			if err := r.SetQueryParam("proxy_port", qProxyPort); err != nil {
				return err
			}
		}
	}

	if o.ProxyType != nil {

		// query param proxy_type
		var qrProxyType string

		if o.ProxyType != nil {
			qrProxyType = *o.ProxyType
		}
		qProxyType := qrProxyType
		if qProxyType != "" {

			if err := r.SetQueryParam("proxy_type", qProxyType); err != nil {
				return err
			}
		}
	}

	if o.ProxyUsername != nil {

		// query param proxy_username
		var qrProxyUsername string

		if o.ProxyUsername != nil {
			qrProxyUsername = *o.ProxyUsername
		}
		qProxyUsername := qrProxyUsername
		if qProxyUsername != "" {

			if err := r.SetQueryParam("proxy_username", qProxyUsername); err != nil {
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

	if o.StateClusterState != nil {

		// query param state.cluster_state
		var qrStateClusterState bool

		if o.StateClusterState != nil {
			qrStateClusterState = *o.StateClusterState
		}
		qStateClusterState := swag.FormatBool(qrStateClusterState)
		if qStateClusterState != "" {

			if err := r.SetQueryParam("state.cluster_state", qStateClusterState); err != nil {
				return err
			}
		}
	}

	if o.StateCode != nil {

		// query param state.code
		var qrStateCode string

		if o.StateCode != nil {
			qrStateCode = *o.StateCode
		}
		qStateCode := qrStateCode
		if qStateCode != "" {

			if err := r.SetQueryParam("state.code", qStateCode); err != nil {
				return err
			}
		}
	}

	if o.StateMessage != nil {

		// query param state.message
		var qrStateMessage string

		if o.StateMessage != nil {
			qrStateMessage = *o.StateMessage
		}
		qStateMessage := qrStateMessage
		if qStateMessage != "" {

			if err := r.SetQueryParam("state.message", qStateMessage); err != nil {
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

	if o.VerifyHost != nil {

		// query param verify_host
		var qrVerifyHost bool

		if o.VerifyHost != nil {
			qrVerifyHost = *o.VerifyHost
		}
		qVerifyHost := swag.FormatBool(qrVerifyHost)
		if qVerifyHost != "" {

			if err := r.SetQueryParam("verify_host", qVerifyHost); err != nil {
				return err
			}
		}
	}

	if o.VerifyIP != nil {

		// query param verify_ip
		var qrVerifyIP bool

		if o.VerifyIP != nil {
			qrVerifyIP = *o.VerifyIP
		}
		qVerifyIP := swag.FormatBool(qrVerifyIP)
		if qVerifyIP != "" {

			if err := r.SetQueryParam("verify_ip", qVerifyIP); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGcpKmsCollectionGet binds the parameter fields
func (o *GcpKmsCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamGcpKmsCollectionGet binds the parameter order_by
func (o *GcpKmsCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
