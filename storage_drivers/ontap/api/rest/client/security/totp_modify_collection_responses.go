// Code generated by go-swagger; DO NOT EDIT.

package security

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

// TotpModifyCollectionReader is a Reader for the TotpModifyCollection structure.
type TotpModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TotpModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTotpModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTotpModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTotpModifyCollectionOK creates a TotpModifyCollectionOK with default headers values
func NewTotpModifyCollectionOK() *TotpModifyCollectionOK {
	return &TotpModifyCollectionOK{}
}

/*
TotpModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type TotpModifyCollectionOK struct {
}

// IsSuccess returns true when this totp modify collection o k response has a 2xx status code
func (o *TotpModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this totp modify collection o k response has a 3xx status code
func (o *TotpModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this totp modify collection o k response has a 4xx status code
func (o *TotpModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this totp modify collection o k response has a 5xx status code
func (o *TotpModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this totp modify collection o k response a status code equal to that given
func (o *TotpModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the totp modify collection o k response
func (o *TotpModifyCollectionOK) Code() int {
	return 200
}

func (o *TotpModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /security/login/totps][%d] totpModifyCollectionOK", 200)
}

func (o *TotpModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /security/login/totps][%d] totpModifyCollectionOK", 200)
}

func (o *TotpModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTotpModifyCollectionDefault creates a TotpModifyCollectionDefault with default headers values
func NewTotpModifyCollectionDefault(code int) *TotpModifyCollectionDefault {
	return &TotpModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	TotpModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 144834564 | Only users with the admin role are allowed to modify the TOTP status. |
| 144834565 | Invalid option for the field -enabled |
| 144834566 | The user does not have a TOTP configuration available for modification. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type TotpModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this totp modify collection default response has a 2xx status code
func (o *TotpModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this totp modify collection default response has a 3xx status code
func (o *TotpModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this totp modify collection default response has a 4xx status code
func (o *TotpModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this totp modify collection default response has a 5xx status code
func (o *TotpModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this totp modify collection default response a status code equal to that given
func (o *TotpModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the totp modify collection default response
func (o *TotpModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *TotpModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/login/totps][%d] totp_modify_collection default %s", o._statusCode, payload)
}

func (o *TotpModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/login/totps][%d] totp_modify_collection default %s", o._statusCode, payload)
}

func (o *TotpModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TotpModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
TotpModifyCollectionBody totp modify collection body
swagger:model TotpModifyCollectionBody
*/
type TotpModifyCollectionBody struct {

	// links
	Links *models.TotpInlineLinks `json:"_links,omitempty"`

	// account
	Account *models.AccountReference `json:"account,omitempty"`

	// Optional comment for the TOTP profile.
	Comment *string `json:"comment,omitempty"`

	// Status of the TOTP profile.
	// Example: false
	Enabled *bool `json:"enabled,omitempty"`

	// owner
	Owner *models.TotpInlineOwner `json:"owner,omitempty"`

	// Scope of the entity. Set to "cluster" for cluster owned objects and to "svm" for SVM owned objects.
	// Read Only: true
	// Enum: ["cluster","svm"]
	Scope *string `json:"scope,omitempty"`

	// SHA fingerprint for the TOTP secret key.
	// Read Only: true
	ShaFingerprint *string `json:"sha_fingerprint,omitempty"`

	// totp response inline records
	TotpResponseInlineRecords []*models.Totp `json:"records,omitempty"`
}

// Validate validates this totp modify collection body
func (o *TotpModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotpResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotpModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *TotpModifyCollectionBody) validateAccount(formats strfmt.Registry) error {
	if swag.IsZero(o.Account) { // not required
		return nil
	}

	if o.Account != nil {
		if err := o.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "account")
			}
			return err
		}
	}

	return nil
}

func (o *TotpModifyCollectionBody) validateOwner(formats strfmt.Registry) error {
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

var totpModifyCollectionBodyTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster","svm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		totpModifyCollectionBodyTypeScopePropEnum = append(totpModifyCollectionBodyTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// TotpModifyCollectionBody
	// TotpModifyCollectionBody
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// TotpModifyCollectionBodyScopeCluster captures enum value "cluster"
	TotpModifyCollectionBodyScopeCluster string = "cluster"

	// BEGIN DEBUGGING
	// TotpModifyCollectionBody
	// TotpModifyCollectionBody
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// TotpModifyCollectionBodyScopeSvm captures enum value "svm"
	TotpModifyCollectionBodyScopeSvm string = "svm"
)

// prop value enum
func (o *TotpModifyCollectionBody) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, totpModifyCollectionBodyTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *TotpModifyCollectionBody) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(o.Scope) { // not required
		return nil
	}

	// value enum
	if err := o.validateScopeEnum("info"+"."+"scope", "body", *o.Scope); err != nil {
		return err
	}

	return nil
}

func (o *TotpModifyCollectionBody) validateTotpResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.TotpResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.TotpResponseInlineRecords); i++ {
		if swag.IsZero(o.TotpResponseInlineRecords[i]) { // not required
			continue
		}

		if o.TotpResponseInlineRecords[i] != nil {
			if err := o.TotpResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this totp modify collection body based on the context it is used
func (o *TotpModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateShaFingerprint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTotpResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotpModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *TotpModifyCollectionBody) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

	if o.Account != nil {
		if err := o.Account.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "account")
			}
			return err
		}
	}

	return nil
}

func (o *TotpModifyCollectionBody) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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

func (o *TotpModifyCollectionBody) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"scope", "body", o.Scope); err != nil {
		return err
	}

	return nil
}

func (o *TotpModifyCollectionBody) contextValidateShaFingerprint(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"sha_fingerprint", "body", o.ShaFingerprint); err != nil {
		return err
	}

	return nil
}

func (o *TotpModifyCollectionBody) contextValidateTotpResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.TotpResponseInlineRecords); i++ {

		if o.TotpResponseInlineRecords[i] != nil {
			if err := o.TotpResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TotpModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TotpModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res TotpModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TotpInlineLinks totp inline links
swagger:model totp_inline__links
*/
type TotpInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this totp inline links
func (o *TotpInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotpInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this totp inline links based on the context it is used
func (o *TotpInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotpInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *TotpInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TotpInlineLinks) UnmarshalBinary(b []byte) error {
	var res TotpInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TotpInlineOwner Owner name and UUID that uniquely identifies the TOTP profile.
swagger:model totp_inline_owner
*/
type TotpInlineOwner struct {

	// links
	Links *models.TotpInlineOwnerInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this totp inline owner
func (o *TotpInlineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotpInlineOwner) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this totp inline owner based on the context it is used
func (o *TotpInlineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotpInlineOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *TotpInlineOwner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TotpInlineOwner) UnmarshalBinary(b []byte) error {
	var res TotpInlineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TotpInlineOwnerInlineLinks totp inline owner inline links
swagger:model totp_inline_owner_inline__links
*/
type TotpInlineOwnerInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this totp inline owner inline links
func (o *TotpInlineOwnerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotpInlineOwnerInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this totp inline owner inline links based on the context it is used
func (o *TotpInlineOwnerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotpInlineOwnerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *TotpInlineOwnerInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TotpInlineOwnerInlineLinks) UnmarshalBinary(b []byte) error {
	var res TotpInlineOwnerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
