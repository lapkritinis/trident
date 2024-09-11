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

// NewCifsShareACLCollectionGetParams creates a new CifsShareACLCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsShareACLCollectionGetParams() *CifsShareACLCollectionGetParams {
	return &CifsShareACLCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsShareACLCollectionGetParamsWithTimeout creates a new CifsShareACLCollectionGetParams object
// with the ability to set a timeout on a request.
func NewCifsShareACLCollectionGetParamsWithTimeout(timeout time.Duration) *CifsShareACLCollectionGetParams {
	return &CifsShareACLCollectionGetParams{
		timeout: timeout,
	}
}

// NewCifsShareACLCollectionGetParamsWithContext creates a new CifsShareACLCollectionGetParams object
// with the ability to set a context for a request.
func NewCifsShareACLCollectionGetParamsWithContext(ctx context.Context) *CifsShareACLCollectionGetParams {
	return &CifsShareACLCollectionGetParams{
		Context: ctx,
	}
}

// NewCifsShareACLCollectionGetParamsWithHTTPClient creates a new CifsShareACLCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsShareACLCollectionGetParamsWithHTTPClient(client *http.Client) *CifsShareACLCollectionGetParams {
	return &CifsShareACLCollectionGetParams{
		HTTPClient: client,
	}
}

/*
CifsShareACLCollectionGetParams contains all the parameters to send to the API endpoint

	for the cifs share acl collection get operation.

	Typically these are written to a http.Request.
*/
type CifsShareACLCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Permission.

	   Filter by permission
	*/
	Permission *string

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

	/* Share.

	   CIFS Share Name
	*/
	Share string

	/* Sid.

	   Filter by sid
	*/
	Sid *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	/* Type.

	   Filter by type
	*/
	Type *string

	/* UnixID.

	   Filter by unix_id
	*/
	UnixID *int64

	/* UserOrGroup.

	   Filter by user_or_group
	*/
	UserOrGroup *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs share acl collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareACLCollectionGetParams) WithDefaults() *CifsShareACLCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs share acl collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareACLCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := CifsShareACLCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithTimeout(timeout time.Duration) *CifsShareACLCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithContext(ctx context.Context) *CifsShareACLCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithHTTPClient(client *http.Client) *CifsShareACLCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithFields(fields []string) *CifsShareACLCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithMaxRecords(maxRecords *int64) *CifsShareACLCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithOrderBy(orderBy []string) *CifsShareACLCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPermission adds the permission to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithPermission(permission *string) *CifsShareACLCollectionGetParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetPermission(permission *string) {
	o.Permission = permission
}

// WithReturnRecords adds the returnRecords to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithReturnRecords(returnRecords *bool) *CifsShareACLCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *CifsShareACLCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithShare adds the share to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithShare(share string) *CifsShareACLCollectionGetParams {
	o.SetShare(share)
	return o
}

// SetShare adds the share to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetShare(share string) {
	o.Share = share
}

// WithSid adds the sid to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithSid(sid *string) *CifsShareACLCollectionGetParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetSid(sid *string) {
	o.Sid = sid
}

// WithSvmName adds the svmName to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithSvmName(svmName *string) *CifsShareACLCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithSvmUUID(svmUUID string) *CifsShareACLCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithType adds the typeVar to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithType(typeVar *string) *CifsShareACLCollectionGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUnixID adds the unixID to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithUnixID(unixID *int64) *CifsShareACLCollectionGetParams {
	o.SetUnixID(unixID)
	return o
}

// SetUnixID adds the unixId to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetUnixID(unixID *int64) {
	o.UnixID = unixID
}

// WithUserOrGroup adds the userOrGroup to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) WithUserOrGroup(userOrGroup *string) *CifsShareACLCollectionGetParams {
	o.SetUserOrGroup(userOrGroup)
	return o
}

// SetUserOrGroup adds the userOrGroup to the cifs share acl collection get params
func (o *CifsShareACLCollectionGetParams) SetUserOrGroup(userOrGroup *string) {
	o.UserOrGroup = userOrGroup
}

// WriteToRequest writes these params to a swagger request
func (o *CifsShareACLCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.Permission != nil {

		// query param permission
		var qrPermission string

		if o.Permission != nil {
			qrPermission = *o.Permission
		}
		qPermission := qrPermission
		if qPermission != "" {

			if err := r.SetQueryParam("permission", qPermission); err != nil {
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

	// path param share
	if err := r.SetPathParam("share", o.Share); err != nil {
		return err
	}

	if o.Sid != nil {

		// query param sid
		var qrSid string

		if o.Sid != nil {
			qrSid = *o.Sid
		}
		qSid := qrSid
		if qSid != "" {

			if err := r.SetQueryParam("sid", qSid); err != nil {
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
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

	if o.UnixID != nil {

		// query param unix_id
		var qrUnixID int64

		if o.UnixID != nil {
			qrUnixID = *o.UnixID
		}
		qUnixID := swag.FormatInt64(qrUnixID)
		if qUnixID != "" {

			if err := r.SetQueryParam("unix_id", qUnixID); err != nil {
				return err
			}
		}
	}

	if o.UserOrGroup != nil {

		// query param user_or_group
		var qrUserOrGroup string

		if o.UserOrGroup != nil {
			qrUserOrGroup = *o.UserOrGroup
		}
		qUserOrGroup := qrUserOrGroup
		if qUserOrGroup != "" {

			if err := r.SetQueryParam("user_or_group", qUserOrGroup); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsShareACLCollectionGet binds the parameter fields
func (o *CifsShareACLCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCifsShareACLCollectionGet binds the parameter order_by
func (o *CifsShareACLCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
