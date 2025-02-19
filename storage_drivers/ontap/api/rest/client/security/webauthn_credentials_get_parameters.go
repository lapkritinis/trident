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

// NewWebauthnCredentialsGetParams creates a new WebauthnCredentialsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWebauthnCredentialsGetParams() *WebauthnCredentialsGetParams {
	return &WebauthnCredentialsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWebauthnCredentialsGetParamsWithTimeout creates a new WebauthnCredentialsGetParams object
// with the ability to set a timeout on a request.
func NewWebauthnCredentialsGetParamsWithTimeout(timeout time.Duration) *WebauthnCredentialsGetParams {
	return &WebauthnCredentialsGetParams{
		timeout: timeout,
	}
}

// NewWebauthnCredentialsGetParamsWithContext creates a new WebauthnCredentialsGetParams object
// with the ability to set a context for a request.
func NewWebauthnCredentialsGetParamsWithContext(ctx context.Context) *WebauthnCredentialsGetParams {
	return &WebauthnCredentialsGetParams{
		Context: ctx,
	}
}

// NewWebauthnCredentialsGetParamsWithHTTPClient creates a new WebauthnCredentialsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewWebauthnCredentialsGetParamsWithHTTPClient(client *http.Client) *WebauthnCredentialsGetParams {
	return &WebauthnCredentialsGetParams{
		HTTPClient: client,
	}
}

/*
WebauthnCredentialsGetParams contains all the parameters to send to the API endpoint

	for the webauthn credentials get operation.

	Typically these are written to a http.Request.
*/
type WebauthnCredentialsGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Index.

	   Index.
	*/
	Index int64

	/* OwnerUUID.

	   Used to identify a cluster or an SVM.
	*/
	OwnerUUID string

	/* RelyingPartyID.

	   Relying Party ID.
	*/
	RelyingPartyID string

	/* Username.

	   Username.
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the webauthn credentials get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebauthnCredentialsGetParams) WithDefaults() *WebauthnCredentialsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the webauthn credentials get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebauthnCredentialsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) WithTimeout(timeout time.Duration) *WebauthnCredentialsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) WithContext(ctx context.Context) *WebauthnCredentialsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) WithHTTPClient(client *http.Client) *WebauthnCredentialsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) WithFields(fields []string) *WebauthnCredentialsGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIndex adds the index to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) WithIndex(index int64) *WebauthnCredentialsGetParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwnerUUID adds the ownerUUID to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) WithOwnerUUID(ownerUUID string) *WebauthnCredentialsGetParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WithRelyingPartyID adds the relyingPartyID to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) WithRelyingPartyID(relyingPartyID string) *WebauthnCredentialsGetParams {
	o.SetRelyingPartyID(relyingPartyID)
	return o
}

// SetRelyingPartyID adds the relyingPartyId to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) SetRelyingPartyID(relyingPartyID string) {
	o.RelyingPartyID = relyingPartyID
}

// WithUsername adds the username to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) WithUsername(username string) *WebauthnCredentialsGetParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the webauthn credentials get params
func (o *WebauthnCredentialsGetParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *WebauthnCredentialsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUID); err != nil {
		return err
	}

	// path param relying_party.id
	if err := r.SetPathParam("relying_party.id", o.RelyingPartyID); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamWebauthnCredentialsGet binds the parameter fields
func (o *WebauthnCredentialsGetParams) bindParamFields(formats strfmt.Registry) []string {
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
