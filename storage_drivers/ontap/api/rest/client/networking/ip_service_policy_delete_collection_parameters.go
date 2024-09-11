// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewIPServicePolicyDeleteCollectionParams creates a new IPServicePolicyDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIPServicePolicyDeleteCollectionParams() *IPServicePolicyDeleteCollectionParams {
	return &IPServicePolicyDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIPServicePolicyDeleteCollectionParamsWithTimeout creates a new IPServicePolicyDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewIPServicePolicyDeleteCollectionParamsWithTimeout(timeout time.Duration) *IPServicePolicyDeleteCollectionParams {
	return &IPServicePolicyDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewIPServicePolicyDeleteCollectionParamsWithContext creates a new IPServicePolicyDeleteCollectionParams object
// with the ability to set a context for a request.
func NewIPServicePolicyDeleteCollectionParamsWithContext(ctx context.Context) *IPServicePolicyDeleteCollectionParams {
	return &IPServicePolicyDeleteCollectionParams{
		Context: ctx,
	}
}

// NewIPServicePolicyDeleteCollectionParamsWithHTTPClient creates a new IPServicePolicyDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewIPServicePolicyDeleteCollectionParamsWithHTTPClient(client *http.Client) *IPServicePolicyDeleteCollectionParams {
	return &IPServicePolicyDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
IPServicePolicyDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the ip service policy delete collection operation.

	Typically these are written to a http.Request.
*/
type IPServicePolicyDeleteCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info IPServicePolicyDeleteCollectionBody

	/* IpspaceName.

	   Filter by ipspace.name
	*/
	IpspaceName *string

	/* IpspaceUUID.

	   Filter by ipspace.uuid
	*/
	IpspaceUUID *string

	/* IsBuiltIn.

	   Filter by is_built_in
	*/
	IsBuiltIn *bool

	/* Name.

	   Filter by name
	*/
	Name *string

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

	/* Services.

	   Filter by services
	*/
	Services *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ip service policy delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPServicePolicyDeleteCollectionParams) WithDefaults() *IPServicePolicyDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ip service policy delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPServicePolicyDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := IPServicePolicyDeleteCollectionParams{
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

// WithTimeout adds the timeout to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithTimeout(timeout time.Duration) *IPServicePolicyDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithContext(ctx context.Context) *IPServicePolicyDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithHTTPClient(client *http.Client) *IPServicePolicyDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *IPServicePolicyDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithInfo(info IPServicePolicyDeleteCollectionBody) *IPServicePolicyDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetInfo(info IPServicePolicyDeleteCollectionBody) {
	o.Info = info
}

// WithIpspaceName adds the ipspaceName to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithIpspaceName(ipspaceName *string) *IPServicePolicyDeleteCollectionParams {
	o.SetIpspaceName(ipspaceName)
	return o
}

// SetIpspaceName adds the ipspaceName to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetIpspaceName(ipspaceName *string) {
	o.IpspaceName = ipspaceName
}

// WithIpspaceUUID adds the ipspaceUUID to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithIpspaceUUID(ipspaceUUID *string) *IPServicePolicyDeleteCollectionParams {
	o.SetIpspaceUUID(ipspaceUUID)
	return o
}

// SetIpspaceUUID adds the ipspaceUuid to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetIpspaceUUID(ipspaceUUID *string) {
	o.IpspaceUUID = ipspaceUUID
}

// WithIsBuiltIn adds the isBuiltIn to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithIsBuiltIn(isBuiltIn *bool) *IPServicePolicyDeleteCollectionParams {
	o.SetIsBuiltIn(isBuiltIn)
	return o
}

// SetIsBuiltIn adds the isBuiltIn to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetIsBuiltIn(isBuiltIn *bool) {
	o.IsBuiltIn = isBuiltIn
}

// WithName adds the name to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithName(name *string) *IPServicePolicyDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReturnRecords adds the returnRecords to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *IPServicePolicyDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *IPServicePolicyDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScope adds the scope to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithScope(scope *string) *IPServicePolicyDeleteCollectionParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSerialRecords adds the serialRecords to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *IPServicePolicyDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithServices adds the services to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithServices(services *string) *IPServicePolicyDeleteCollectionParams {
	o.SetServices(services)
	return o
}

// SetServices adds the services to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetServices(services *string) {
	o.Services = services
}

// WithSvmName adds the svmName to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithSvmName(svmName *string) *IPServicePolicyDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithSvmUUID(svmUUID *string) *IPServicePolicyDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) WithUUID(uuid *string) *IPServicePolicyDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the ip service policy delete collection params
func (o *IPServicePolicyDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IPServicePolicyDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.IpspaceName != nil {

		// query param ipspace.name
		var qrIpspaceName string

		if o.IpspaceName != nil {
			qrIpspaceName = *o.IpspaceName
		}
		qIpspaceName := qrIpspaceName
		if qIpspaceName != "" {

			if err := r.SetQueryParam("ipspace.name", qIpspaceName); err != nil {
				return err
			}
		}
	}

	if o.IpspaceUUID != nil {

		// query param ipspace.uuid
		var qrIpspaceUUID string

		if o.IpspaceUUID != nil {
			qrIpspaceUUID = *o.IpspaceUUID
		}
		qIpspaceUUID := qrIpspaceUUID
		if qIpspaceUUID != "" {

			if err := r.SetQueryParam("ipspace.uuid", qIpspaceUUID); err != nil {
				return err
			}
		}
	}

	if o.IsBuiltIn != nil {

		// query param is_built_in
		var qrIsBuiltIn bool

		if o.IsBuiltIn != nil {
			qrIsBuiltIn = *o.IsBuiltIn
		}
		qIsBuiltIn := swag.FormatBool(qrIsBuiltIn)
		if qIsBuiltIn != "" {

			if err := r.SetQueryParam("is_built_in", qIsBuiltIn); err != nil {
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

	if o.Services != nil {

		// query param services
		var qrServices string

		if o.Services != nil {
			qrServices = *o.Services
		}
		qServices := qrServices
		if qServices != "" {

			if err := r.SetQueryParam("services", qServices); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
