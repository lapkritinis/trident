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

// NewAwsKmsDeleteCollectionParams creates a new AwsKmsDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAwsKmsDeleteCollectionParams() *AwsKmsDeleteCollectionParams {
	return &AwsKmsDeleteCollectionParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewAwsKmsDeleteCollectionParamsWithTimeout creates a new AwsKmsDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewAwsKmsDeleteCollectionParamsWithTimeout(timeout time.Duration) *AwsKmsDeleteCollectionParams {
	return &AwsKmsDeleteCollectionParams{
		requestTimeout: timeout,
	}
}

// NewAwsKmsDeleteCollectionParamsWithContext creates a new AwsKmsDeleteCollectionParams object
// with the ability to set a context for a request.
func NewAwsKmsDeleteCollectionParamsWithContext(ctx context.Context) *AwsKmsDeleteCollectionParams {
	return &AwsKmsDeleteCollectionParams{
		Context: ctx,
	}
}

// NewAwsKmsDeleteCollectionParamsWithHTTPClient creates a new AwsKmsDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAwsKmsDeleteCollectionParamsWithHTTPClient(client *http.Client) *AwsKmsDeleteCollectionParams {
	return &AwsKmsDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
AwsKmsDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the aws kms delete collection operation.

	Typically these are written to a http.Request.
*/
type AwsKmsDeleteCollectionParams struct {

	/* AccessKeyID.

	   Filter by access_key_id
	*/
	AccessKeyID *string

	/* AmazonReachabilityCode.

	   Filter by amazon_reachability.code
	*/
	AmazonReachabilityCode *string

	/* AmazonReachabilityMessage.

	   Filter by amazon_reachability.message
	*/
	AmazonReachabilityMessage *string

	/* AmazonReachabilityReachable.

	   Filter by amazon_reachability.reachable
	*/
	AmazonReachabilityReachable *bool

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* DefaultDomain.

	   Filter by default_domain
	*/
	DefaultDomain *string

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

	/* EncryptionContext.

	   Filter by encryption_context
	*/
	EncryptionContext *string

	/* Host.

	   Filter by host
	*/
	Host *string

	/* Info.

	   Info specification
	*/
	Info AwsKmsDeleteCollectionBody

	/* KeyID.

	   Filter by key_id
	*/
	KeyID *string

	/* PollingPeriod.

	   Filter by polling_period
	*/
	PollingPeriod *int64

	/* Port.

	   Filter by port
	*/
	Port *int64

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

	/* Region.

	   Filter by region
	*/
	Region *string

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

	/* Service.

	   Filter by service
	*/
	Service *string

	/* SkipVerify.

	   Filter by skip_verify
	*/
	SkipVerify *bool

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

	/* Timeout.

	   Filter by timeout
	*/
	Timeout *int64

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	/* Verify.

	   Filter by verify
	*/
	Verify *bool

	/* VerifyHost.

	   Filter by verify_host
	*/
	VerifyHost *bool

	/* VerifyIP.

	   Filter by verify_ip
	*/
	VerifyIP *bool

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the aws kms delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AwsKmsDeleteCollectionParams) WithDefaults() *AwsKmsDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aws kms delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AwsKmsDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := AwsKmsDeleteCollectionParams{
		ContinueOnFailure: &continueOnFailureDefault,
		ReturnRecords:     &returnRecordsDefault,
		ReturnTimeout:     &returnTimeoutDefault,
		SerialRecords:     &serialRecordsDefault,
	}

	val.requestTimeout = o.requestTimeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithRequestTimeout adds the timeout to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithRequestTimeout(timeout time.Duration) *AwsKmsDeleteCollectionParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithContext(ctx context.Context) *AwsKmsDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithHTTPClient(client *http.Client) *AwsKmsDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKeyID adds the accessKeyID to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithAccessKeyID(accessKeyID *string) *AwsKmsDeleteCollectionParams {
	o.SetAccessKeyID(accessKeyID)
	return o
}

// SetAccessKeyID adds the accessKeyId to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetAccessKeyID(accessKeyID *string) {
	o.AccessKeyID = accessKeyID
}

// WithAmazonReachabilityCode adds the amazonReachabilityCode to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithAmazonReachabilityCode(amazonReachabilityCode *string) *AwsKmsDeleteCollectionParams {
	o.SetAmazonReachabilityCode(amazonReachabilityCode)
	return o
}

// SetAmazonReachabilityCode adds the amazonReachabilityCode to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetAmazonReachabilityCode(amazonReachabilityCode *string) {
	o.AmazonReachabilityCode = amazonReachabilityCode
}

// WithAmazonReachabilityMessage adds the amazonReachabilityMessage to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithAmazonReachabilityMessage(amazonReachabilityMessage *string) *AwsKmsDeleteCollectionParams {
	o.SetAmazonReachabilityMessage(amazonReachabilityMessage)
	return o
}

// SetAmazonReachabilityMessage adds the amazonReachabilityMessage to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetAmazonReachabilityMessage(amazonReachabilityMessage *string) {
	o.AmazonReachabilityMessage = amazonReachabilityMessage
}

// WithAmazonReachabilityReachable adds the amazonReachabilityReachable to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithAmazonReachabilityReachable(amazonReachabilityReachable *bool) *AwsKmsDeleteCollectionParams {
	o.SetAmazonReachabilityReachable(amazonReachabilityReachable)
	return o
}

// SetAmazonReachabilityReachable adds the amazonReachabilityReachable to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetAmazonReachabilityReachable(amazonReachabilityReachable *bool) {
	o.AmazonReachabilityReachable = amazonReachabilityReachable
}

// WithContinueOnFailure adds the continueOnFailure to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *AwsKmsDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDefaultDomain adds the defaultDomain to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithDefaultDomain(defaultDomain *string) *AwsKmsDeleteCollectionParams {
	o.SetDefaultDomain(defaultDomain)
	return o
}

// SetDefaultDomain adds the defaultDomain to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetDefaultDomain(defaultDomain *string) {
	o.DefaultDomain = defaultDomain
}

// WithEkmipReachabilityCode adds the ekmipReachabilityCode to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithEkmipReachabilityCode(ekmipReachabilityCode *string) *AwsKmsDeleteCollectionParams {
	o.SetEkmipReachabilityCode(ekmipReachabilityCode)
	return o
}

// SetEkmipReachabilityCode adds the ekmipReachabilityCode to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetEkmipReachabilityCode(ekmipReachabilityCode *string) {
	o.EkmipReachabilityCode = ekmipReachabilityCode
}

// WithEkmipReachabilityMessage adds the ekmipReachabilityMessage to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithEkmipReachabilityMessage(ekmipReachabilityMessage *string) *AwsKmsDeleteCollectionParams {
	o.SetEkmipReachabilityMessage(ekmipReachabilityMessage)
	return o
}

// SetEkmipReachabilityMessage adds the ekmipReachabilityMessage to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetEkmipReachabilityMessage(ekmipReachabilityMessage *string) {
	o.EkmipReachabilityMessage = ekmipReachabilityMessage
}

// WithEkmipReachabilityNodeName adds the ekmipReachabilityNodeName to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithEkmipReachabilityNodeName(ekmipReachabilityNodeName *string) *AwsKmsDeleteCollectionParams {
	o.SetEkmipReachabilityNodeName(ekmipReachabilityNodeName)
	return o
}

// SetEkmipReachabilityNodeName adds the ekmipReachabilityNodeName to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetEkmipReachabilityNodeName(ekmipReachabilityNodeName *string) {
	o.EkmipReachabilityNodeName = ekmipReachabilityNodeName
}

// WithEkmipReachabilityNodeUUID adds the ekmipReachabilityNodeUUID to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithEkmipReachabilityNodeUUID(ekmipReachabilityNodeUUID *string) *AwsKmsDeleteCollectionParams {
	o.SetEkmipReachabilityNodeUUID(ekmipReachabilityNodeUUID)
	return o
}

// SetEkmipReachabilityNodeUUID adds the ekmipReachabilityNodeUuid to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetEkmipReachabilityNodeUUID(ekmipReachabilityNodeUUID *string) {
	o.EkmipReachabilityNodeUUID = ekmipReachabilityNodeUUID
}

// WithEkmipReachabilityReachable adds the ekmipReachabilityReachable to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithEkmipReachabilityReachable(ekmipReachabilityReachable *bool) *AwsKmsDeleteCollectionParams {
	o.SetEkmipReachabilityReachable(ekmipReachabilityReachable)
	return o
}

// SetEkmipReachabilityReachable adds the ekmipReachabilityReachable to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetEkmipReachabilityReachable(ekmipReachabilityReachable *bool) {
	o.EkmipReachabilityReachable = ekmipReachabilityReachable
}

// WithEncryptionContext adds the encryptionContext to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithEncryptionContext(encryptionContext *string) *AwsKmsDeleteCollectionParams {
	o.SetEncryptionContext(encryptionContext)
	return o
}

// SetEncryptionContext adds the encryptionContext to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetEncryptionContext(encryptionContext *string) {
	o.EncryptionContext = encryptionContext
}

// WithHost adds the host to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithHost(host *string) *AwsKmsDeleteCollectionParams {
	o.SetHost(host)
	return o
}

// SetHost adds the host to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetHost(host *string) {
	o.Host = host
}

// WithInfo adds the info to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithInfo(info AwsKmsDeleteCollectionBody) *AwsKmsDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetInfo(info AwsKmsDeleteCollectionBody) {
	o.Info = info
}

// WithKeyID adds the keyID to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithKeyID(keyID *string) *AwsKmsDeleteCollectionParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetKeyID(keyID *string) {
	o.KeyID = keyID
}

// WithPollingPeriod adds the pollingPeriod to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithPollingPeriod(pollingPeriod *int64) *AwsKmsDeleteCollectionParams {
	o.SetPollingPeriod(pollingPeriod)
	return o
}

// SetPollingPeriod adds the pollingPeriod to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetPollingPeriod(pollingPeriod *int64) {
	o.PollingPeriod = pollingPeriod
}

// WithPort adds the port to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithPort(port *int64) *AwsKmsDeleteCollectionParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetPort(port *int64) {
	o.Port = port
}

// WithProxyHost adds the proxyHost to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithProxyHost(proxyHost *string) *AwsKmsDeleteCollectionParams {
	o.SetProxyHost(proxyHost)
	return o
}

// SetProxyHost adds the proxyHost to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetProxyHost(proxyHost *string) {
	o.ProxyHost = proxyHost
}

// WithProxyPort adds the proxyPort to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithProxyPort(proxyPort *int64) *AwsKmsDeleteCollectionParams {
	o.SetProxyPort(proxyPort)
	return o
}

// SetProxyPort adds the proxyPort to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetProxyPort(proxyPort *int64) {
	o.ProxyPort = proxyPort
}

// WithProxyType adds the proxyType to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithProxyType(proxyType *string) *AwsKmsDeleteCollectionParams {
	o.SetProxyType(proxyType)
	return o
}

// SetProxyType adds the proxyType to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetProxyType(proxyType *string) {
	o.ProxyType = proxyType
}

// WithProxyUsername adds the proxyUsername to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithProxyUsername(proxyUsername *string) *AwsKmsDeleteCollectionParams {
	o.SetProxyUsername(proxyUsername)
	return o
}

// SetProxyUsername adds the proxyUsername to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetProxyUsername(proxyUsername *string) {
	o.ProxyUsername = proxyUsername
}

// WithRegion adds the region to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithRegion(region *string) *AwsKmsDeleteCollectionParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetRegion(region *string) {
	o.Region = region
}

// WithReturnRecords adds the returnRecords to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *AwsKmsDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *AwsKmsDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScope adds the scope to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithScope(scope *string) *AwsKmsDeleteCollectionParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSerialRecords adds the serialRecords to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *AwsKmsDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithService adds the service to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithService(service *string) *AwsKmsDeleteCollectionParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetService(service *string) {
	o.Service = service
}

// WithSkipVerify adds the skipVerify to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithSkipVerify(skipVerify *bool) *AwsKmsDeleteCollectionParams {
	o.SetSkipVerify(skipVerify)
	return o
}

// SetSkipVerify adds the skipVerify to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetSkipVerify(skipVerify *bool) {
	o.SkipVerify = skipVerify
}

// WithStateClusterState adds the stateClusterState to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithStateClusterState(stateClusterState *bool) *AwsKmsDeleteCollectionParams {
	o.SetStateClusterState(stateClusterState)
	return o
}

// SetStateClusterState adds the stateClusterState to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetStateClusterState(stateClusterState *bool) {
	o.StateClusterState = stateClusterState
}

// WithStateCode adds the stateCode to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithStateCode(stateCode *string) *AwsKmsDeleteCollectionParams {
	o.SetStateCode(stateCode)
	return o
}

// SetStateCode adds the stateCode to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetStateCode(stateCode *string) {
	o.StateCode = stateCode
}

// WithStateMessage adds the stateMessage to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithStateMessage(stateMessage *string) *AwsKmsDeleteCollectionParams {
	o.SetStateMessage(stateMessage)
	return o
}

// SetStateMessage adds the stateMessage to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetStateMessage(stateMessage *string) {
	o.StateMessage = stateMessage
}

// WithSvmName adds the svmName to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithSvmName(svmName *string) *AwsKmsDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithSvmUUID(svmUUID *string) *AwsKmsDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithTimeout adds the timeout to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithTimeout(timeout *int64) *AwsKmsDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WithUUID adds the uuid to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithUUID(uuid *string) *AwsKmsDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WithVerify adds the verify to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithVerify(verify *bool) *AwsKmsDeleteCollectionParams {
	o.SetVerify(verify)
	return o
}

// SetVerify adds the verify to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetVerify(verify *bool) {
	o.Verify = verify
}

// WithVerifyHost adds the verifyHost to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithVerifyHost(verifyHost *bool) *AwsKmsDeleteCollectionParams {
	o.SetVerifyHost(verifyHost)
	return o
}

// SetVerifyHost adds the verifyHost to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetVerifyHost(verifyHost *bool) {
	o.VerifyHost = verifyHost
}

// WithVerifyIP adds the verifyIP to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) WithVerifyIP(verifyIP *bool) *AwsKmsDeleteCollectionParams {
	o.SetVerifyIP(verifyIP)
	return o
}

// SetVerifyIP adds the verifyIp to the aws kms delete collection params
func (o *AwsKmsDeleteCollectionParams) SetVerifyIP(verifyIP *bool) {
	o.VerifyIP = verifyIP
}

// WriteToRequest writes these params to a swagger request
func (o *AwsKmsDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	if o.AccessKeyID != nil {

		// query param access_key_id
		var qrAccessKeyID string

		if o.AccessKeyID != nil {
			qrAccessKeyID = *o.AccessKeyID
		}
		qAccessKeyID := qrAccessKeyID
		if qAccessKeyID != "" {

			if err := r.SetQueryParam("access_key_id", qAccessKeyID); err != nil {
				return err
			}
		}
	}

	if o.AmazonReachabilityCode != nil {

		// query param amazon_reachability.code
		var qrAmazonReachabilityCode string

		if o.AmazonReachabilityCode != nil {
			qrAmazonReachabilityCode = *o.AmazonReachabilityCode
		}
		qAmazonReachabilityCode := qrAmazonReachabilityCode
		if qAmazonReachabilityCode != "" {

			if err := r.SetQueryParam("amazon_reachability.code", qAmazonReachabilityCode); err != nil {
				return err
			}
		}
	}

	if o.AmazonReachabilityMessage != nil {

		// query param amazon_reachability.message
		var qrAmazonReachabilityMessage string

		if o.AmazonReachabilityMessage != nil {
			qrAmazonReachabilityMessage = *o.AmazonReachabilityMessage
		}
		qAmazonReachabilityMessage := qrAmazonReachabilityMessage
		if qAmazonReachabilityMessage != "" {

			if err := r.SetQueryParam("amazon_reachability.message", qAmazonReachabilityMessage); err != nil {
				return err
			}
		}
	}

	if o.AmazonReachabilityReachable != nil {

		// query param amazon_reachability.reachable
		var qrAmazonReachabilityReachable bool

		if o.AmazonReachabilityReachable != nil {
			qrAmazonReachabilityReachable = *o.AmazonReachabilityReachable
		}
		qAmazonReachabilityReachable := swag.FormatBool(qrAmazonReachabilityReachable)
		if qAmazonReachabilityReachable != "" {

			if err := r.SetQueryParam("amazon_reachability.reachable", qAmazonReachabilityReachable); err != nil {
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

	if o.DefaultDomain != nil {

		// query param default_domain
		var qrDefaultDomain string

		if o.DefaultDomain != nil {
			qrDefaultDomain = *o.DefaultDomain
		}
		qDefaultDomain := qrDefaultDomain
		if qDefaultDomain != "" {

			if err := r.SetQueryParam("default_domain", qDefaultDomain); err != nil {
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

	if o.EncryptionContext != nil {

		// query param encryption_context
		var qrEncryptionContext string

		if o.EncryptionContext != nil {
			qrEncryptionContext = *o.EncryptionContext
		}
		qEncryptionContext := qrEncryptionContext
		if qEncryptionContext != "" {

			if err := r.SetQueryParam("encryption_context", qEncryptionContext); err != nil {
				return err
			}
		}
	}

	if o.Host != nil {

		// query param host
		var qrHost string

		if o.Host != nil {
			qrHost = *o.Host
		}
		qHost := qrHost
		if qHost != "" {

			if err := r.SetQueryParam("host", qHost); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.KeyID != nil {

		// query param key_id
		var qrKeyID string

		if o.KeyID != nil {
			qrKeyID = *o.KeyID
		}
		qKeyID := qrKeyID
		if qKeyID != "" {

			if err := r.SetQueryParam("key_id", qKeyID); err != nil {
				return err
			}
		}
	}

	if o.PollingPeriod != nil {

		// query param polling_period
		var qrPollingPeriod int64

		if o.PollingPeriod != nil {
			qrPollingPeriod = *o.PollingPeriod
		}
		qPollingPeriod := swag.FormatInt64(qrPollingPeriod)
		if qPollingPeriod != "" {

			if err := r.SetQueryParam("polling_period", qPollingPeriod); err != nil {
				return err
			}
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

	if o.Region != nil {

		// query param region
		var qrRegion string

		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {

			if err := r.SetQueryParam("region", qRegion); err != nil {
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

	if o.Service != nil {

		// query param service
		var qrService string

		if o.Service != nil {
			qrService = *o.Service
		}
		qService := qrService
		if qService != "" {

			if err := r.SetQueryParam("service", qService); err != nil {
				return err
			}
		}
	}

	if o.SkipVerify != nil {

		// query param skip_verify
		var qrSkipVerify bool

		if o.SkipVerify != nil {
			qrSkipVerify = *o.SkipVerify
		}
		qSkipVerify := swag.FormatBool(qrSkipVerify)
		if qSkipVerify != "" {

			if err := r.SetQueryParam("skip_verify", qSkipVerify); err != nil {
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

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64

		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {

			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
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

	if o.Verify != nil {

		// query param verify
		var qrVerify bool

		if o.Verify != nil {
			qrVerify = *o.Verify
		}
		qVerify := swag.FormatBool(qrVerify)
		if qVerify != "" {

			if err := r.SetQueryParam("verify", qVerify); err != nil {
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
