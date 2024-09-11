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

// NewKerberosRealmDeleteCollectionParams creates a new KerberosRealmDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKerberosRealmDeleteCollectionParams() *KerberosRealmDeleteCollectionParams {
	return &KerberosRealmDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKerberosRealmDeleteCollectionParamsWithTimeout creates a new KerberosRealmDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewKerberosRealmDeleteCollectionParamsWithTimeout(timeout time.Duration) *KerberosRealmDeleteCollectionParams {
	return &KerberosRealmDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewKerberosRealmDeleteCollectionParamsWithContext creates a new KerberosRealmDeleteCollectionParams object
// with the ability to set a context for a request.
func NewKerberosRealmDeleteCollectionParamsWithContext(ctx context.Context) *KerberosRealmDeleteCollectionParams {
	return &KerberosRealmDeleteCollectionParams{
		Context: ctx,
	}
}

// NewKerberosRealmDeleteCollectionParamsWithHTTPClient creates a new KerberosRealmDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewKerberosRealmDeleteCollectionParamsWithHTTPClient(client *http.Client) *KerberosRealmDeleteCollectionParams {
	return &KerberosRealmDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
KerberosRealmDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the kerberos realm delete collection operation.

	Typically these are written to a http.Request.
*/
type KerberosRealmDeleteCollectionParams struct {

	/* AdServerAddress.

	   Filter by ad_server.address
	*/
	AdServerAddress *string

	/* AdServerName.

	   Filter by ad_server.name
	*/
	AdServerName *string

	/* AdminServerAddress.

	   Filter by admin_server.address
	*/
	AdminServerAddress *string

	/* AdminServerPort.

	   Filter by admin_server.port
	*/
	AdminServerPort *int64

	/* ClockSkew.

	   Filter by clock_skew
	*/
	ClockSkew *int64

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* EncryptionTypes.

	   Filter by encryption_types
	*/
	EncryptionTypes *string

	/* Info.

	   Info specification
	*/
	Info KerberosRealmDeleteCollectionBody

	/* KdcIP.

	   Filter by kdc.ip
	*/
	KdcIP *string

	/* KdcPort.

	   Filter by kdc.port
	*/
	KdcPort *int64

	/* KdcVendor.

	   Filter by kdc.vendor
	*/
	KdcVendor *string

	/* Name.

	   Filter by name
	*/
	Name *string

	/* PasswordServerAddress.

	   Filter by password_server.address
	*/
	PasswordServerAddress *string

	/* PasswordServerPort.

	   Filter by password_server.port
	*/
	PasswordServerPort *int64

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

// WithDefaults hydrates default values in the kerberos realm delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosRealmDeleteCollectionParams) WithDefaults() *KerberosRealmDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kerberos realm delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosRealmDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := KerberosRealmDeleteCollectionParams{
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

// WithTimeout adds the timeout to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithTimeout(timeout time.Duration) *KerberosRealmDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithContext(ctx context.Context) *KerberosRealmDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithHTTPClient(client *http.Client) *KerberosRealmDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdServerAddress adds the adServerAddress to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithAdServerAddress(adServerAddress *string) *KerberosRealmDeleteCollectionParams {
	o.SetAdServerAddress(adServerAddress)
	return o
}

// SetAdServerAddress adds the adServerAddress to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetAdServerAddress(adServerAddress *string) {
	o.AdServerAddress = adServerAddress
}

// WithAdServerName adds the adServerName to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithAdServerName(adServerName *string) *KerberosRealmDeleteCollectionParams {
	o.SetAdServerName(adServerName)
	return o
}

// SetAdServerName adds the adServerName to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetAdServerName(adServerName *string) {
	o.AdServerName = adServerName
}

// WithAdminServerAddress adds the adminServerAddress to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithAdminServerAddress(adminServerAddress *string) *KerberosRealmDeleteCollectionParams {
	o.SetAdminServerAddress(adminServerAddress)
	return o
}

// SetAdminServerAddress adds the adminServerAddress to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetAdminServerAddress(adminServerAddress *string) {
	o.AdminServerAddress = adminServerAddress
}

// WithAdminServerPort adds the adminServerPort to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithAdminServerPort(adminServerPort *int64) *KerberosRealmDeleteCollectionParams {
	o.SetAdminServerPort(adminServerPort)
	return o
}

// SetAdminServerPort adds the adminServerPort to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetAdminServerPort(adminServerPort *int64) {
	o.AdminServerPort = adminServerPort
}

// WithClockSkew adds the clockSkew to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithClockSkew(clockSkew *int64) *KerberosRealmDeleteCollectionParams {
	o.SetClockSkew(clockSkew)
	return o
}

// SetClockSkew adds the clockSkew to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetClockSkew(clockSkew *int64) {
	o.ClockSkew = clockSkew
}

// WithComment adds the comment to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithComment(comment *string) *KerberosRealmDeleteCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContinueOnFailure adds the continueOnFailure to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *KerberosRealmDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithEncryptionTypes adds the encryptionTypes to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithEncryptionTypes(encryptionTypes *string) *KerberosRealmDeleteCollectionParams {
	o.SetEncryptionTypes(encryptionTypes)
	return o
}

// SetEncryptionTypes adds the encryptionTypes to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetEncryptionTypes(encryptionTypes *string) {
	o.EncryptionTypes = encryptionTypes
}

// WithInfo adds the info to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithInfo(info KerberosRealmDeleteCollectionBody) *KerberosRealmDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetInfo(info KerberosRealmDeleteCollectionBody) {
	o.Info = info
}

// WithKdcIP adds the kdcIP to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithKdcIP(kdcIP *string) *KerberosRealmDeleteCollectionParams {
	o.SetKdcIP(kdcIP)
	return o
}

// SetKdcIP adds the kdcIp to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetKdcIP(kdcIP *string) {
	o.KdcIP = kdcIP
}

// WithKdcPort adds the kdcPort to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithKdcPort(kdcPort *int64) *KerberosRealmDeleteCollectionParams {
	o.SetKdcPort(kdcPort)
	return o
}

// SetKdcPort adds the kdcPort to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetKdcPort(kdcPort *int64) {
	o.KdcPort = kdcPort
}

// WithKdcVendor adds the kdcVendor to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithKdcVendor(kdcVendor *string) *KerberosRealmDeleteCollectionParams {
	o.SetKdcVendor(kdcVendor)
	return o
}

// SetKdcVendor adds the kdcVendor to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetKdcVendor(kdcVendor *string) {
	o.KdcVendor = kdcVendor
}

// WithName adds the name to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithName(name *string) *KerberosRealmDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithPasswordServerAddress adds the passwordServerAddress to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithPasswordServerAddress(passwordServerAddress *string) *KerberosRealmDeleteCollectionParams {
	o.SetPasswordServerAddress(passwordServerAddress)
	return o
}

// SetPasswordServerAddress adds the passwordServerAddress to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetPasswordServerAddress(passwordServerAddress *string) {
	o.PasswordServerAddress = passwordServerAddress
}

// WithPasswordServerPort adds the passwordServerPort to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithPasswordServerPort(passwordServerPort *int64) *KerberosRealmDeleteCollectionParams {
	o.SetPasswordServerPort(passwordServerPort)
	return o
}

// SetPasswordServerPort adds the passwordServerPort to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetPasswordServerPort(passwordServerPort *int64) {
	o.PasswordServerPort = passwordServerPort
}

// WithReturnRecords adds the returnRecords to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *KerberosRealmDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *KerberosRealmDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *KerberosRealmDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithSvmName(svmName *string) *KerberosRealmDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) WithSvmUUID(svmUUID *string) *KerberosRealmDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the kerberos realm delete collection params
func (o *KerberosRealmDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *KerberosRealmDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdServerAddress != nil {

		// query param ad_server.address
		var qrAdServerAddress string

		if o.AdServerAddress != nil {
			qrAdServerAddress = *o.AdServerAddress
		}
		qAdServerAddress := qrAdServerAddress
		if qAdServerAddress != "" {

			if err := r.SetQueryParam("ad_server.address", qAdServerAddress); err != nil {
				return err
			}
		}
	}

	if o.AdServerName != nil {

		// query param ad_server.name
		var qrAdServerName string

		if o.AdServerName != nil {
			qrAdServerName = *o.AdServerName
		}
		qAdServerName := qrAdServerName
		if qAdServerName != "" {

			if err := r.SetQueryParam("ad_server.name", qAdServerName); err != nil {
				return err
			}
		}
	}

	if o.AdminServerAddress != nil {

		// query param admin_server.address
		var qrAdminServerAddress string

		if o.AdminServerAddress != nil {
			qrAdminServerAddress = *o.AdminServerAddress
		}
		qAdminServerAddress := qrAdminServerAddress
		if qAdminServerAddress != "" {

			if err := r.SetQueryParam("admin_server.address", qAdminServerAddress); err != nil {
				return err
			}
		}
	}

	if o.AdminServerPort != nil {

		// query param admin_server.port
		var qrAdminServerPort int64

		if o.AdminServerPort != nil {
			qrAdminServerPort = *o.AdminServerPort
		}
		qAdminServerPort := swag.FormatInt64(qrAdminServerPort)
		if qAdminServerPort != "" {

			if err := r.SetQueryParam("admin_server.port", qAdminServerPort); err != nil {
				return err
			}
		}
	}

	if o.ClockSkew != nil {

		// query param clock_skew
		var qrClockSkew int64

		if o.ClockSkew != nil {
			qrClockSkew = *o.ClockSkew
		}
		qClockSkew := swag.FormatInt64(qrClockSkew)
		if qClockSkew != "" {

			if err := r.SetQueryParam("clock_skew", qClockSkew); err != nil {
				return err
			}
		}
	}

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

	if o.EncryptionTypes != nil {

		// query param encryption_types
		var qrEncryptionTypes string

		if o.EncryptionTypes != nil {
			qrEncryptionTypes = *o.EncryptionTypes
		}
		qEncryptionTypes := qrEncryptionTypes
		if qEncryptionTypes != "" {

			if err := r.SetQueryParam("encryption_types", qEncryptionTypes); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.KdcIP != nil {

		// query param kdc.ip
		var qrKdcIP string

		if o.KdcIP != nil {
			qrKdcIP = *o.KdcIP
		}
		qKdcIP := qrKdcIP
		if qKdcIP != "" {

			if err := r.SetQueryParam("kdc.ip", qKdcIP); err != nil {
				return err
			}
		}
	}

	if o.KdcPort != nil {

		// query param kdc.port
		var qrKdcPort int64

		if o.KdcPort != nil {
			qrKdcPort = *o.KdcPort
		}
		qKdcPort := swag.FormatInt64(qrKdcPort)
		if qKdcPort != "" {

			if err := r.SetQueryParam("kdc.port", qKdcPort); err != nil {
				return err
			}
		}
	}

	if o.KdcVendor != nil {

		// query param kdc.vendor
		var qrKdcVendor string

		if o.KdcVendor != nil {
			qrKdcVendor = *o.KdcVendor
		}
		qKdcVendor := qrKdcVendor
		if qKdcVendor != "" {

			if err := r.SetQueryParam("kdc.vendor", qKdcVendor); err != nil {
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

	if o.PasswordServerAddress != nil {

		// query param password_server.address
		var qrPasswordServerAddress string

		if o.PasswordServerAddress != nil {
			qrPasswordServerAddress = *o.PasswordServerAddress
		}
		qPasswordServerAddress := qrPasswordServerAddress
		if qPasswordServerAddress != "" {

			if err := r.SetQueryParam("password_server.address", qPasswordServerAddress); err != nil {
				return err
			}
		}
	}

	if o.PasswordServerPort != nil {

		// query param password_server.port
		var qrPasswordServerPort int64

		if o.PasswordServerPort != nil {
			qrPasswordServerPort = *o.PasswordServerPort
		}
		qPasswordServerPort := swag.FormatInt64(qrPasswordServerPort)
		if qPasswordServerPort != "" {

			if err := r.SetQueryParam("password_server.port", qPasswordServerPort); err != nil {
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
