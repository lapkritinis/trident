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

// NewSecurityCertificateDeleteCollectionParams creates a new SecurityCertificateDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityCertificateDeleteCollectionParams() *SecurityCertificateDeleteCollectionParams {
	return &SecurityCertificateDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityCertificateDeleteCollectionParamsWithTimeout creates a new SecurityCertificateDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewSecurityCertificateDeleteCollectionParamsWithTimeout(timeout time.Duration) *SecurityCertificateDeleteCollectionParams {
	return &SecurityCertificateDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewSecurityCertificateDeleteCollectionParamsWithContext creates a new SecurityCertificateDeleteCollectionParams object
// with the ability to set a context for a request.
func NewSecurityCertificateDeleteCollectionParamsWithContext(ctx context.Context) *SecurityCertificateDeleteCollectionParams {
	return &SecurityCertificateDeleteCollectionParams{
		Context: ctx,
	}
}

// NewSecurityCertificateDeleteCollectionParamsWithHTTPClient creates a new SecurityCertificateDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityCertificateDeleteCollectionParamsWithHTTPClient(client *http.Client) *SecurityCertificateDeleteCollectionParams {
	return &SecurityCertificateDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
SecurityCertificateDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the security certificate delete collection operation.

	Typically these are written to a http.Request.
*/
type SecurityCertificateDeleteCollectionParams struct {

	/* AuthorityKeyIdentifier.

	   Filter by authority_key_identifier
	*/
	AuthorityKeyIdentifier *string

	/* Ca.

	   Filter by ca
	*/
	Ca *string

	/* CommonName.

	   Filter by common_name
	*/
	CommonName *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* ExpiryTime.

	   Filter by expiry_time
	*/
	ExpiryTime *string

	/* HashFunction.

	   Filter by hash_function
	*/
	HashFunction *string

	/* Info.

	   Info specification
	*/
	Info SecurityCertificateDeleteCollectionBody

	/* KeySize.

	   Filter by key_size
	*/
	KeySize *int64

	/* Name.

	   Filter by name
	*/
	Name *string

	/* PublicCertificate.

	   Filter by public_certificate
	*/
	PublicCertificate *string

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

	/* SerialNumber.

	   Filter by serial_number
	*/
	SerialNumber *string

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* SubjectAlternativesDNS.

	   Filter by subject_alternatives.dns
	*/
	SubjectAlternativesDNS *string

	/* SubjectAlternativesEmail.

	   Filter by subject_alternatives.email
	*/
	SubjectAlternativesEmail *string

	/* SubjectAlternativesIP.

	   Filter by subject_alternatives.ip
	*/
	SubjectAlternativesIP *string

	/* SubjectAlternativesURI.

	   Filter by subject_alternatives.uri
	*/
	SubjectAlternativesURI *string

	/* SubjectKeyIdentifier.

	   Filter by subject_key_identifier
	*/
	SubjectKeyIdentifier *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* Type.

	   Filter by type
	*/
	Type *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security certificate delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityCertificateDeleteCollectionParams) WithDefaults() *SecurityCertificateDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security certificate delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityCertificateDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SecurityCertificateDeleteCollectionParams{
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

// WithTimeout adds the timeout to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithTimeout(timeout time.Duration) *SecurityCertificateDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithContext(ctx context.Context) *SecurityCertificateDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithHTTPClient(client *http.Client) *SecurityCertificateDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorityKeyIdentifier adds the authorityKeyIdentifier to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithAuthorityKeyIdentifier(authorityKeyIdentifier *string) *SecurityCertificateDeleteCollectionParams {
	o.SetAuthorityKeyIdentifier(authorityKeyIdentifier)
	return o
}

// SetAuthorityKeyIdentifier adds the authorityKeyIdentifier to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetAuthorityKeyIdentifier(authorityKeyIdentifier *string) {
	o.AuthorityKeyIdentifier = authorityKeyIdentifier
}

// WithCa adds the ca to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithCa(ca *string) *SecurityCertificateDeleteCollectionParams {
	o.SetCa(ca)
	return o
}

// SetCa adds the ca to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetCa(ca *string) {
	o.Ca = ca
}

// WithCommonName adds the commonName to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithCommonName(commonName *string) *SecurityCertificateDeleteCollectionParams {
	o.SetCommonName(commonName)
	return o
}

// SetCommonName adds the commonName to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetCommonName(commonName *string) {
	o.CommonName = commonName
}

// WithContinueOnFailure adds the continueOnFailure to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SecurityCertificateDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithExpiryTime adds the expiryTime to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithExpiryTime(expiryTime *string) *SecurityCertificateDeleteCollectionParams {
	o.SetExpiryTime(expiryTime)
	return o
}

// SetExpiryTime adds the expiryTime to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetExpiryTime(expiryTime *string) {
	o.ExpiryTime = expiryTime
}

// WithHashFunction adds the hashFunction to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithHashFunction(hashFunction *string) *SecurityCertificateDeleteCollectionParams {
	o.SetHashFunction(hashFunction)
	return o
}

// SetHashFunction adds the hashFunction to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetHashFunction(hashFunction *string) {
	o.HashFunction = hashFunction
}

// WithInfo adds the info to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithInfo(info SecurityCertificateDeleteCollectionBody) *SecurityCertificateDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetInfo(info SecurityCertificateDeleteCollectionBody) {
	o.Info = info
}

// WithKeySize adds the keySize to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithKeySize(keySize *int64) *SecurityCertificateDeleteCollectionParams {
	o.SetKeySize(keySize)
	return o
}

// SetKeySize adds the keySize to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetKeySize(keySize *int64) {
	o.KeySize = keySize
}

// WithName adds the name to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithName(name *string) *SecurityCertificateDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithPublicCertificate adds the publicCertificate to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithPublicCertificate(publicCertificate *string) *SecurityCertificateDeleteCollectionParams {
	o.SetPublicCertificate(publicCertificate)
	return o
}

// SetPublicCertificate adds the publicCertificate to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetPublicCertificate(publicCertificate *string) {
	o.PublicCertificate = publicCertificate
}

// WithReturnRecords adds the returnRecords to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *SecurityCertificateDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *SecurityCertificateDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScope adds the scope to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithScope(scope *string) *SecurityCertificateDeleteCollectionParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSerialNumber adds the serialNumber to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithSerialNumber(serialNumber *string) *SecurityCertificateDeleteCollectionParams {
	o.SetSerialNumber(serialNumber)
	return o
}

// SetSerialNumber adds the serialNumber to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetSerialNumber(serialNumber *string) {
	o.SerialNumber = serialNumber
}

// WithSerialRecords adds the serialRecords to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *SecurityCertificateDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSubjectAlternativesDNS adds the subjectAlternativesDNS to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithSubjectAlternativesDNS(subjectAlternativesDNS *string) *SecurityCertificateDeleteCollectionParams {
	o.SetSubjectAlternativesDNS(subjectAlternativesDNS)
	return o
}

// SetSubjectAlternativesDNS adds the subjectAlternativesDns to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetSubjectAlternativesDNS(subjectAlternativesDNS *string) {
	o.SubjectAlternativesDNS = subjectAlternativesDNS
}

// WithSubjectAlternativesEmail adds the subjectAlternativesEmail to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithSubjectAlternativesEmail(subjectAlternativesEmail *string) *SecurityCertificateDeleteCollectionParams {
	o.SetSubjectAlternativesEmail(subjectAlternativesEmail)
	return o
}

// SetSubjectAlternativesEmail adds the subjectAlternativesEmail to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetSubjectAlternativesEmail(subjectAlternativesEmail *string) {
	o.SubjectAlternativesEmail = subjectAlternativesEmail
}

// WithSubjectAlternativesIP adds the subjectAlternativesIP to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithSubjectAlternativesIP(subjectAlternativesIP *string) *SecurityCertificateDeleteCollectionParams {
	o.SetSubjectAlternativesIP(subjectAlternativesIP)
	return o
}

// SetSubjectAlternativesIP adds the subjectAlternativesIp to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetSubjectAlternativesIP(subjectAlternativesIP *string) {
	o.SubjectAlternativesIP = subjectAlternativesIP
}

// WithSubjectAlternativesURI adds the subjectAlternativesURI to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithSubjectAlternativesURI(subjectAlternativesURI *string) *SecurityCertificateDeleteCollectionParams {
	o.SetSubjectAlternativesURI(subjectAlternativesURI)
	return o
}

// SetSubjectAlternativesURI adds the subjectAlternativesUri to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetSubjectAlternativesURI(subjectAlternativesURI *string) {
	o.SubjectAlternativesURI = subjectAlternativesURI
}

// WithSubjectKeyIdentifier adds the subjectKeyIdentifier to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithSubjectKeyIdentifier(subjectKeyIdentifier *string) *SecurityCertificateDeleteCollectionParams {
	o.SetSubjectKeyIdentifier(subjectKeyIdentifier)
	return o
}

// SetSubjectKeyIdentifier adds the subjectKeyIdentifier to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetSubjectKeyIdentifier(subjectKeyIdentifier *string) {
	o.SubjectKeyIdentifier = subjectKeyIdentifier
}

// WithSvmName adds the svmName to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithSvmName(svmName *string) *SecurityCertificateDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithSvmUUID(svmUUID *string) *SecurityCertificateDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithType adds the typeVar to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithType(typeVar *string) *SecurityCertificateDeleteCollectionParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUUID adds the uuid to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) WithUUID(uuid *string) *SecurityCertificateDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the security certificate delete collection params
func (o *SecurityCertificateDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityCertificateDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthorityKeyIdentifier != nil {

		// query param authority_key_identifier
		var qrAuthorityKeyIdentifier string

		if o.AuthorityKeyIdentifier != nil {
			qrAuthorityKeyIdentifier = *o.AuthorityKeyIdentifier
		}
		qAuthorityKeyIdentifier := qrAuthorityKeyIdentifier
		if qAuthorityKeyIdentifier != "" {

			if err := r.SetQueryParam("authority_key_identifier", qAuthorityKeyIdentifier); err != nil {
				return err
			}
		}
	}

	if o.Ca != nil {

		// query param ca
		var qrCa string

		if o.Ca != nil {
			qrCa = *o.Ca
		}
		qCa := qrCa
		if qCa != "" {

			if err := r.SetQueryParam("ca", qCa); err != nil {
				return err
			}
		}
	}

	if o.CommonName != nil {

		// query param common_name
		var qrCommonName string

		if o.CommonName != nil {
			qrCommonName = *o.CommonName
		}
		qCommonName := qrCommonName
		if qCommonName != "" {

			if err := r.SetQueryParam("common_name", qCommonName); err != nil {
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

	if o.HashFunction != nil {

		// query param hash_function
		var qrHashFunction string

		if o.HashFunction != nil {
			qrHashFunction = *o.HashFunction
		}
		qHashFunction := qrHashFunction
		if qHashFunction != "" {

			if err := r.SetQueryParam("hash_function", qHashFunction); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.KeySize != nil {

		// query param key_size
		var qrKeySize int64

		if o.KeySize != nil {
			qrKeySize = *o.KeySize
		}
		qKeySize := swag.FormatInt64(qrKeySize)
		if qKeySize != "" {

			if err := r.SetQueryParam("key_size", qKeySize); err != nil {
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

	if o.PublicCertificate != nil {

		// query param public_certificate
		var qrPublicCertificate string

		if o.PublicCertificate != nil {
			qrPublicCertificate = *o.PublicCertificate
		}
		qPublicCertificate := qrPublicCertificate
		if qPublicCertificate != "" {

			if err := r.SetQueryParam("public_certificate", qPublicCertificate); err != nil {
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

	if o.SubjectAlternativesDNS != nil {

		// query param subject_alternatives.dns
		var qrSubjectAlternativesDNS string

		if o.SubjectAlternativesDNS != nil {
			qrSubjectAlternativesDNS = *o.SubjectAlternativesDNS
		}
		qSubjectAlternativesDNS := qrSubjectAlternativesDNS
		if qSubjectAlternativesDNS != "" {

			if err := r.SetQueryParam("subject_alternatives.dns", qSubjectAlternativesDNS); err != nil {
				return err
			}
		}
	}

	if o.SubjectAlternativesEmail != nil {

		// query param subject_alternatives.email
		var qrSubjectAlternativesEmail string

		if o.SubjectAlternativesEmail != nil {
			qrSubjectAlternativesEmail = *o.SubjectAlternativesEmail
		}
		qSubjectAlternativesEmail := qrSubjectAlternativesEmail
		if qSubjectAlternativesEmail != "" {

			if err := r.SetQueryParam("subject_alternatives.email", qSubjectAlternativesEmail); err != nil {
				return err
			}
		}
	}

	if o.SubjectAlternativesIP != nil {

		// query param subject_alternatives.ip
		var qrSubjectAlternativesIP string

		if o.SubjectAlternativesIP != nil {
			qrSubjectAlternativesIP = *o.SubjectAlternativesIP
		}
		qSubjectAlternativesIP := qrSubjectAlternativesIP
		if qSubjectAlternativesIP != "" {

			if err := r.SetQueryParam("subject_alternatives.ip", qSubjectAlternativesIP); err != nil {
				return err
			}
		}
	}

	if o.SubjectAlternativesURI != nil {

		// query param subject_alternatives.uri
		var qrSubjectAlternativesURI string

		if o.SubjectAlternativesURI != nil {
			qrSubjectAlternativesURI = *o.SubjectAlternativesURI
		}
		qSubjectAlternativesURI := qrSubjectAlternativesURI
		if qSubjectAlternativesURI != "" {

			if err := r.SetQueryParam("subject_alternatives.uri", qSubjectAlternativesURI); err != nil {
				return err
			}
		}
	}

	if o.SubjectKeyIdentifier != nil {

		// query param subject_key_identifier
		var qrSubjectKeyIdentifier string

		if o.SubjectKeyIdentifier != nil {
			qrSubjectKeyIdentifier = *o.SubjectKeyIdentifier
		}
		qSubjectKeyIdentifier := qrSubjectKeyIdentifier
		if qSubjectKeyIdentifier != "" {

			if err := r.SetQueryParam("subject_key_identifier", qSubjectKeyIdentifier); err != nil {
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
