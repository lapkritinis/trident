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

// NewAccountPublickeyDeleteCollectionParams creates a new AccountPublickeyDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountPublickeyDeleteCollectionParams() *AccountPublickeyDeleteCollectionParams {
	return &AccountPublickeyDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountPublickeyDeleteCollectionParamsWithTimeout creates a new AccountPublickeyDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewAccountPublickeyDeleteCollectionParamsWithTimeout(timeout time.Duration) *AccountPublickeyDeleteCollectionParams {
	return &AccountPublickeyDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewAccountPublickeyDeleteCollectionParamsWithContext creates a new AccountPublickeyDeleteCollectionParams object
// with the ability to set a context for a request.
func NewAccountPublickeyDeleteCollectionParamsWithContext(ctx context.Context) *AccountPublickeyDeleteCollectionParams {
	return &AccountPublickeyDeleteCollectionParams{
		Context: ctx,
	}
}

// NewAccountPublickeyDeleteCollectionParamsWithHTTPClient creates a new AccountPublickeyDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountPublickeyDeleteCollectionParamsWithHTTPClient(client *http.Client) *AccountPublickeyDeleteCollectionParams {
	return &AccountPublickeyDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
AccountPublickeyDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the account publickey delete collection operation.

	Typically these are written to a http.Request.
*/
type AccountPublickeyDeleteCollectionParams struct {

	/* AccountName.

	   Filter by account.name
	*/
	AccountName *string

	/* Certificate.

	   Filter by certificate
	*/
	Certificate *string

	/* CertificateDetails.

	   Filter by certificate_details
	*/
	CertificateDetails *string

	/* CertificateExpired.

	   Filter by certificate_expired
	*/
	CertificateExpired *string

	/* CertificateRevoked.

	   Filter by certificate_revoked
	*/
	CertificateRevoked *string

	/* Comment.

	   Filter by comment
	*/
	Comment *string

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
	Info AccountPublickeyDeleteCollectionBody

	/* ObfuscatedFingerprint.

	   Filter by obfuscated_fingerprint
	*/
	ObfuscatedFingerprint *string

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerName *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUID *string

	/* PublicKey.

	   Filter by public_key
	*/
	PublicKey *string

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

	/* ShaFingerprint.

	   Filter by sha_fingerprint
	*/
	ShaFingerprint *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account publickey delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountPublickeyDeleteCollectionParams) WithDefaults() *AccountPublickeyDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account publickey delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountPublickeyDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := AccountPublickeyDeleteCollectionParams{
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

// WithTimeout adds the timeout to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithTimeout(timeout time.Duration) *AccountPublickeyDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithContext(ctx context.Context) *AccountPublickeyDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithHTTPClient(client *http.Client) *AccountPublickeyDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountName adds the accountName to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithAccountName(accountName *string) *AccountPublickeyDeleteCollectionParams {
	o.SetAccountName(accountName)
	return o
}

// SetAccountName adds the accountName to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetAccountName(accountName *string) {
	o.AccountName = accountName
}

// WithCertificate adds the certificate to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithCertificate(certificate *string) *AccountPublickeyDeleteCollectionParams {
	o.SetCertificate(certificate)
	return o
}

// SetCertificate adds the certificate to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetCertificate(certificate *string) {
	o.Certificate = certificate
}

// WithCertificateDetails adds the certificateDetails to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithCertificateDetails(certificateDetails *string) *AccountPublickeyDeleteCollectionParams {
	o.SetCertificateDetails(certificateDetails)
	return o
}

// SetCertificateDetails adds the certificateDetails to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetCertificateDetails(certificateDetails *string) {
	o.CertificateDetails = certificateDetails
}

// WithCertificateExpired adds the certificateExpired to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithCertificateExpired(certificateExpired *string) *AccountPublickeyDeleteCollectionParams {
	o.SetCertificateExpired(certificateExpired)
	return o
}

// SetCertificateExpired adds the certificateExpired to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetCertificateExpired(certificateExpired *string) {
	o.CertificateExpired = certificateExpired
}

// WithCertificateRevoked adds the certificateRevoked to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithCertificateRevoked(certificateRevoked *string) *AccountPublickeyDeleteCollectionParams {
	o.SetCertificateRevoked(certificateRevoked)
	return o
}

// SetCertificateRevoked adds the certificateRevoked to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetCertificateRevoked(certificateRevoked *string) {
	o.CertificateRevoked = certificateRevoked
}

// WithComment adds the comment to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithComment(comment *string) *AccountPublickeyDeleteCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContinueOnFailure adds the continueOnFailure to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *AccountPublickeyDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithIndex adds the index to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithIndex(index *int64) *AccountPublickeyDeleteCollectionParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetIndex(index *int64) {
	o.Index = index
}

// WithInfo adds the info to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithInfo(info AccountPublickeyDeleteCollectionBody) *AccountPublickeyDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetInfo(info AccountPublickeyDeleteCollectionBody) {
	o.Info = info
}

// WithObfuscatedFingerprint adds the obfuscatedFingerprint to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithObfuscatedFingerprint(obfuscatedFingerprint *string) *AccountPublickeyDeleteCollectionParams {
	o.SetObfuscatedFingerprint(obfuscatedFingerprint)
	return o
}

// SetObfuscatedFingerprint adds the obfuscatedFingerprint to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetObfuscatedFingerprint(obfuscatedFingerprint *string) {
	o.ObfuscatedFingerprint = obfuscatedFingerprint
}

// WithOwnerName adds the ownerName to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithOwnerName(ownerName *string) *AccountPublickeyDeleteCollectionParams {
	o.SetOwnerName(ownerName)
	return o
}

// SetOwnerName adds the ownerName to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetOwnerName(ownerName *string) {
	o.OwnerName = ownerName
}

// WithOwnerUUID adds the ownerUUID to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithOwnerUUID(ownerUUID *string) *AccountPublickeyDeleteCollectionParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetOwnerUUID(ownerUUID *string) {
	o.OwnerUUID = ownerUUID
}

// WithPublicKey adds the publicKey to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithPublicKey(publicKey *string) *AccountPublickeyDeleteCollectionParams {
	o.SetPublicKey(publicKey)
	return o
}

// SetPublicKey adds the publicKey to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetPublicKey(publicKey *string) {
	o.PublicKey = publicKey
}

// WithReturnRecords adds the returnRecords to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *AccountPublickeyDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *AccountPublickeyDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScope adds the scope to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithScope(scope *string) *AccountPublickeyDeleteCollectionParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSerialRecords adds the serialRecords to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *AccountPublickeyDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithShaFingerprint adds the shaFingerprint to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) WithShaFingerprint(shaFingerprint *string) *AccountPublickeyDeleteCollectionParams {
	o.SetShaFingerprint(shaFingerprint)
	return o
}

// SetShaFingerprint adds the shaFingerprint to the account publickey delete collection params
func (o *AccountPublickeyDeleteCollectionParams) SetShaFingerprint(shaFingerprint *string) {
	o.ShaFingerprint = shaFingerprint
}

// WriteToRequest writes these params to a swagger request
func (o *AccountPublickeyDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountName != nil {

		// query param account.name
		var qrAccountName string

		if o.AccountName != nil {
			qrAccountName = *o.AccountName
		}
		qAccountName := qrAccountName
		if qAccountName != "" {

			if err := r.SetQueryParam("account.name", qAccountName); err != nil {
				return err
			}
		}
	}

	if o.Certificate != nil {

		// query param certificate
		var qrCertificate string

		if o.Certificate != nil {
			qrCertificate = *o.Certificate
		}
		qCertificate := qrCertificate
		if qCertificate != "" {

			if err := r.SetQueryParam("certificate", qCertificate); err != nil {
				return err
			}
		}
	}

	if o.CertificateDetails != nil {

		// query param certificate_details
		var qrCertificateDetails string

		if o.CertificateDetails != nil {
			qrCertificateDetails = *o.CertificateDetails
		}
		qCertificateDetails := qrCertificateDetails
		if qCertificateDetails != "" {

			if err := r.SetQueryParam("certificate_details", qCertificateDetails); err != nil {
				return err
			}
		}
	}

	if o.CertificateExpired != nil {

		// query param certificate_expired
		var qrCertificateExpired string

		if o.CertificateExpired != nil {
			qrCertificateExpired = *o.CertificateExpired
		}
		qCertificateExpired := qrCertificateExpired
		if qCertificateExpired != "" {

			if err := r.SetQueryParam("certificate_expired", qCertificateExpired); err != nil {
				return err
			}
		}
	}

	if o.CertificateRevoked != nil {

		// query param certificate_revoked
		var qrCertificateRevoked string

		if o.CertificateRevoked != nil {
			qrCertificateRevoked = *o.CertificateRevoked
		}
		qCertificateRevoked := qrCertificateRevoked
		if qCertificateRevoked != "" {

			if err := r.SetQueryParam("certificate_revoked", qCertificateRevoked); err != nil {
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

	if o.ObfuscatedFingerprint != nil {

		// query param obfuscated_fingerprint
		var qrObfuscatedFingerprint string

		if o.ObfuscatedFingerprint != nil {
			qrObfuscatedFingerprint = *o.ObfuscatedFingerprint
		}
		qObfuscatedFingerprint := qrObfuscatedFingerprint
		if qObfuscatedFingerprint != "" {

			if err := r.SetQueryParam("obfuscated_fingerprint", qObfuscatedFingerprint); err != nil {
				return err
			}
		}
	}

	if o.OwnerName != nil {

		// query param owner.name
		var qrOwnerName string

		if o.OwnerName != nil {
			qrOwnerName = *o.OwnerName
		}
		qOwnerName := qrOwnerName
		if qOwnerName != "" {

			if err := r.SetQueryParam("owner.name", qOwnerName); err != nil {
				return err
			}
		}
	}

	if o.OwnerUUID != nil {

		// query param owner.uuid
		var qrOwnerUUID string

		if o.OwnerUUID != nil {
			qrOwnerUUID = *o.OwnerUUID
		}
		qOwnerUUID := qrOwnerUUID
		if qOwnerUUID != "" {

			if err := r.SetQueryParam("owner.uuid", qOwnerUUID); err != nil {
				return err
			}
		}
	}

	if o.PublicKey != nil {

		// query param public_key
		var qrPublicKey string

		if o.PublicKey != nil {
			qrPublicKey = *o.PublicKey
		}
		qPublicKey := qrPublicKey
		if qPublicKey != "" {

			if err := r.SetQueryParam("public_key", qPublicKey); err != nil {
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

	if o.ShaFingerprint != nil {

		// query param sha_fingerprint
		var qrShaFingerprint string

		if o.ShaFingerprint != nil {
			qrShaFingerprint = *o.ShaFingerprint
		}
		qShaFingerprint := qrShaFingerprint
		if qShaFingerprint != "" {

			if err := r.SetQueryParam("sha_fingerprint", qShaFingerprint); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
