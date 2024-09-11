// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LocalHostModifyCollectionReader is a Reader for the LocalHostModifyCollection structure.
type LocalHostModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalHostModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalHostModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalHostModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalHostModifyCollectionOK creates a LocalHostModifyCollectionOK with default headers values
func NewLocalHostModifyCollectionOK() *LocalHostModifyCollectionOK {
	return &LocalHostModifyCollectionOK{}
}

/*
LocalHostModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type LocalHostModifyCollectionOK struct {
}

// IsSuccess returns true when this local host modify collection o k response has a 2xx status code
func (o *LocalHostModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local host modify collection o k response has a 3xx status code
func (o *LocalHostModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local host modify collection o k response has a 4xx status code
func (o *LocalHostModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local host modify collection o k response has a 5xx status code
func (o *LocalHostModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local host modify collection o k response a status code equal to that given
func (o *LocalHostModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the local host modify collection o k response
func (o *LocalHostModifyCollectionOK) Code() int {
	return 200
}

func (o *LocalHostModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/local-hosts][%d] localHostModifyCollectionOK", 200)
}

func (o *LocalHostModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /name-services/local-hosts][%d] localHostModifyCollectionOK", 200)
}

func (o *LocalHostModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLocalHostModifyCollectionDefault creates a LocalHostModifyCollectionDefault with default headers values
func NewLocalHostModifyCollectionDefault(code int) *LocalHostModifyCollectionDefault {
	return &LocalHostModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	LocalHostModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 23724055 | Internal error. Configuration for Vserver failed. Verify that the cluster is healthy, then try the command again. For further assistance, contact technical support. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LocalHostModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local host modify collection default response has a 2xx status code
func (o *LocalHostModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local host modify collection default response has a 3xx status code
func (o *LocalHostModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local host modify collection default response has a 4xx status code
func (o *LocalHostModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local host modify collection default response has a 5xx status code
func (o *LocalHostModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local host modify collection default response a status code equal to that given
func (o *LocalHostModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local host modify collection default response
func (o *LocalHostModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *LocalHostModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/local-hosts][%d] local_host_modify_collection default %s", o._statusCode, payload)
}

func (o *LocalHostModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/local-hosts][%d] local_host_modify_collection default %s", o._statusCode, payload)
}

func (o *LocalHostModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalHostModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
LocalHostModifyCollectionBody local host modify collection body
swagger:model LocalHostModifyCollectionBody
*/
type LocalHostModifyCollectionBody struct {

	// links
	Links *models.LocalHostInlineLinks `json:"_links,omitempty"`

	// IPv4/IPv6 address in dotted form.
	// Example: 123.123.123.123
	Address *string `json:"address,omitempty"`

	// Canonical hostname.
	// Example: host.sales.foo.com
	// Max Length: 255
	// Min Length: 1
	Hostname *string `json:"hostname,omitempty"`

	// The list of aliases.
	// Example: ["host1.sales.foo.com","host2.sakes.foo.com"]
	LocalHostInlineAliases []*string `json:"aliases,omitempty"`

	// local host response inline records
	LocalHostResponseInlineRecords []*models.LocalHost `json:"records,omitempty"`

	// owner
	Owner *models.LocalHostInlineOwner `json:"owner,omitempty"`

	// Scope of the entity. Set to "cluster" for cluster owned objects and to "svm" for SVM owned objects.
	// Read Only: true
	// Enum: ["cluster","svm"]
	Scope *string `json:"scope,omitempty"`
}

// Validate validates this local host modify collection body
func (o *LocalHostModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocalHostResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocalHostModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *LocalHostModifyCollectionBody) validateHostname(formats strfmt.Registry) error {
	if swag.IsZero(o.Hostname) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"hostname", "body", *o.Hostname, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"hostname", "body", *o.Hostname, 255); err != nil {
		return err
	}

	return nil
}

func (o *LocalHostModifyCollectionBody) validateLocalHostResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.LocalHostResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.LocalHostResponseInlineRecords); i++ {
		if swag.IsZero(o.LocalHostResponseInlineRecords[i]) { // not required
			continue
		}

		if o.LocalHostResponseInlineRecords[i] != nil {
			if err := o.LocalHostResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *LocalHostModifyCollectionBody) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(o.Owner) { // not required
		return nil
	}

	if o.Owner != nil {
		if err := o.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "owner")
			}
			return err
		}
	}

	return nil
}

var localHostModifyCollectionBodyTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster","svm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		localHostModifyCollectionBodyTypeScopePropEnum = append(localHostModifyCollectionBodyTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// LocalHostModifyCollectionBody
	// LocalHostModifyCollectionBody
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// LocalHostModifyCollectionBodyScopeCluster captures enum value "cluster"
	LocalHostModifyCollectionBodyScopeCluster string = "cluster"

	// BEGIN DEBUGGING
	// LocalHostModifyCollectionBody
	// LocalHostModifyCollectionBody
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// LocalHostModifyCollectionBodyScopeSvm captures enum value "svm"
	LocalHostModifyCollectionBodyScopeSvm string = "svm"
)

// prop value enum
func (o *LocalHostModifyCollectionBody) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, localHostModifyCollectionBodyTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *LocalHostModifyCollectionBody) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(o.Scope) { // not required
		return nil
	}

	// value enum
	if err := o.validateScopeEnum("info"+"."+"scope", "body", *o.Scope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this local host modify collection body based on the context it is used
func (o *LocalHostModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLocalHostResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocalHostModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *LocalHostModifyCollectionBody) contextValidateLocalHostResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.LocalHostResponseInlineRecords); i++ {

		if o.LocalHostResponseInlineRecords[i] != nil {
			if err := o.LocalHostResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *LocalHostModifyCollectionBody) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if o.Owner != nil {
		if err := o.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "owner")
			}
			return err
		}
	}

	return nil
}

func (o *LocalHostModifyCollectionBody) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"scope", "body", o.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LocalHostModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocalHostModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res LocalHostModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
LocalHostInlineLinks local host inline links
swagger:model local_host_inline__links
*/
type LocalHostInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this local host inline links
func (o *LocalHostInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocalHostInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local host inline links based on the context it is used
func (o *LocalHostInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocalHostInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LocalHostInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocalHostInlineLinks) UnmarshalBinary(b []byte) error {
	var res LocalHostInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
LocalHostInlineOwner SVM, applies only to SVM-scoped objects.
swagger:model local_host_inline_owner
*/
type LocalHostInlineOwner struct {

	// links
	Links *models.LocalHostInlineOwnerInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this local host inline owner
func (o *LocalHostInlineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocalHostInlineOwner) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local host inline owner based on the context it is used
func (o *LocalHostInlineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocalHostInlineOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LocalHostInlineOwner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocalHostInlineOwner) UnmarshalBinary(b []byte) error {
	var res LocalHostInlineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
LocalHostInlineOwnerInlineLinks local host inline owner inline links
swagger:model local_host_inline_owner_inline__links
*/
type LocalHostInlineOwnerInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this local host inline owner inline links
func (o *LocalHostInlineOwnerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocalHostInlineOwnerInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local host inline owner inline links based on the context it is used
func (o *LocalHostInlineOwnerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocalHostInlineOwnerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LocalHostInlineOwnerInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocalHostInlineOwnerInlineLinks) UnmarshalBinary(b []byte) error {
	var res LocalHostInlineOwnerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
