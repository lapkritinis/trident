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

// AccountModifyCollectionReader is a Reader for the AccountModifyCollection structure.
type AccountModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountModifyCollectionOK creates a AccountModifyCollectionOK with default headers values
func NewAccountModifyCollectionOK() *AccountModifyCollectionOK {
	return &AccountModifyCollectionOK{}
}

/*
AccountModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type AccountModifyCollectionOK struct {
}

// IsSuccess returns true when this account modify collection o k response has a 2xx status code
func (o *AccountModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account modify collection o k response has a 3xx status code
func (o *AccountModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account modify collection o k response has a 4xx status code
func (o *AccountModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account modify collection o k response has a 5xx status code
func (o *AccountModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account modify collection o k response a status code equal to that given
func (o *AccountModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account modify collection o k response
func (o *AccountModifyCollectionOK) Code() int {
	return 200
}

func (o *AccountModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /security/accounts][%d] accountModifyCollectionOK", 200)
}

func (o *AccountModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /security/accounts][%d] accountModifyCollectionOK", 200)
}

func (o *AccountModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountModifyCollectionDefault creates a AccountModifyCollectionDefault with default headers values
func NewAccountModifyCollectionDefault(code int) *AccountModifyCollectionDefault {
	return &AccountModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	AccountModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1261215 | The role was not found. |
| 1261218 | The user was not found. |
| 1263343 | Cannot lock user with password not set or non-password authentication method. |
| 5636096 | Cannot perform the operation for this user account since the password is not set. |
| 5636097 | The operation for user account failed since user password is not set. |
| 5636100 | Modification of a service-processor user's role to a non-admin role is not supported. |
| 5636125 | The operation not supported on AutoSupport user account which is reserved. |
| 5636129 | The role does not exist. |
| 5636136 | Specifying "is_ns_switch_group" as "true" is supported only for authentication method "nsswitch". |
| 5636154 | The second authentication method parameter is supported for SSH and Service Processor (SP) applications only. |
| 5636155 | The second-authentication-method parameter can be specified only if the authentication-method password or public key nsswitch. |
| 5636156 | Same value cannot be specified for the second-authentication-method and the authentication-method. |
| 5636159 | For a given user and application, if the second-authentication-method is specified, only one such login entry is supported. |
| 5636164 | If the value for either the authentication-method second-authentication-method is nsswitch or password, the other parameter must differ. |
| 5636165 | Second authentication method is not supported for NIS or LDAP group based accounts. |
| 5636197 | LDAP fastbind combination for application and authentication method is not supported. |
| 5636198 | LDAP fastbind authentication is supported only for nsswitch. |
| 5636210 | User creation failed because LDAP is not configured for the SVM or the LDAP connection is not secure. |
| 5636212 | TOTP is supported only when the primary authentication method is password or public key. |
| 5636214 | Configuring the user with TOTP as secondary authentication method requires an effective cluster version of 9.13.1 or later |
| 5636223 | Specifying "is_ns_switch_group" as "true" is supported only for SSH, ONTAPI and HTTP applications. |
| 5636224 | Configuring a Service Processor (SP) user with two-factor authentication requires an effective cluster version of 9.15.1 or later. |
| 5636225 | For a Service Processor (SP) user, the second factor of authentication must be one of publickey or none. |
| 5636226 | Internal error. Failed to check for ONTAP capability. |
| 7077896 | Cannot lock the account of the last console admin user. |
| 7077906 | A role with that name has not been defined for the Vserver. |
| 7077911 | The user is not configured to use the password authentication method. |
| 7077918 | The password cannot contain the username. |
| 7077919 | The minimum length for new password does not meet the policy. |
| 7077920 | The new password must have both letters and numbers. |
| 7077921 | The minimum number of special characters required do not meet the policy. |
| 7077924 | The new password must be different than last N passwords. |
| 7077925 | The new password must be different to the old password. |
| 7077929 | Cannot lock user with password not set or non-password authentication method. |
| 7077940 | The password exceeds maximum supported length. |
| 7077941 | Defined password composition exceeds the maximum password length of 128 characters. |
| 7078900 | An aAdmin password is not set. Set the password by including it in the request. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AccountModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this account modify collection default response has a 2xx status code
func (o *AccountModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this account modify collection default response has a 3xx status code
func (o *AccountModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this account modify collection default response has a 4xx status code
func (o *AccountModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this account modify collection default response has a 5xx status code
func (o *AccountModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this account modify collection default response a status code equal to that given
func (o *AccountModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the account modify collection default response
func (o *AccountModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *AccountModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/accounts][%d] account_modify_collection default %s", o._statusCode, payload)
}

func (o *AccountModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/accounts][%d] account_modify_collection default %s", o._statusCode, payload)
}

func (o *AccountModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AccountModifyCollectionBody account modify collection body
swagger:model AccountModifyCollectionBody
*/
type AccountModifyCollectionBody struct {

	// links
	Links *models.AccountInlineLinks `json:"_links,omitempty"`

	// account inline applications
	AccountInlineApplications []*models.AccountApplication `json:"applications,omitempty"`

	// account response inline records
	AccountResponseInlineRecords []*models.Account `json:"records,omitempty"`

	// Optional comment for the user account.
	Comment *string `json:"comment,omitempty"`

	// Locked status of the account.
	Locked *bool `json:"locked,omitempty"`

	// User or group account name
	// Example: joe.smith
	// Max Length: 64
	// Min Length: 3
	Name *string `json:"name,omitempty"`

	// owner
	Owner *models.AccountInlineOwner `json:"owner,omitempty"`

	// Password for the account. The password can contain a mix of lower and upper case alphabetic characters, digits, and special characters.
	// Max Length: 128
	// Min Length: 8
	// Format: password
	Password *strfmt.Password `json:"password,omitempty"`

	// Password hash algorithm used to generate a hash of the user's password for password matching.To modify "password_hash_algorithm", use REST API "/api/security/authentication/password".
	// Example: sha512
	// Read Only: true
	// Enum: ["sha512","sha256","md5"]
	PasswordHashAlgorithm *string `json:"password_hash_algorithm,omitempty"`

	// role
	Role *models.AccountInlineRole `json:"role,omitempty"`

	// Scope of the entity. Set to "cluster" for cluster owned objects and to "svm" for SVM owned objects.
	// Read Only: true
	// Enum: ["cluster","svm"]
	Scope *string `json:"scope,omitempty"`
}

// Validate validates this account modify collection body
func (o *AccountModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAccountInlineApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAccountResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePasswordHashAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRole(formats); err != nil {
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

func (o *AccountModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *AccountModifyCollectionBody) validateAccountInlineApplications(formats strfmt.Registry) error {
	if swag.IsZero(o.AccountInlineApplications) { // not required
		return nil
	}

	for i := 0; i < len(o.AccountInlineApplications); i++ {
		if swag.IsZero(o.AccountInlineApplications[i]) { // not required
			continue
		}

		if o.AccountInlineApplications[i] != nil {
			if err := o.AccountInlineApplications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AccountModifyCollectionBody) validateAccountResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.AccountResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.AccountResponseInlineRecords); i++ {
		if swag.IsZero(o.AccountResponseInlineRecords[i]) { // not required
			continue
		}

		if o.AccountResponseInlineRecords[i] != nil {
			if err := o.AccountResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AccountModifyCollectionBody) validateName(formats strfmt.Registry) error {
	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"name", "body", *o.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"name", "body", *o.Name, 64); err != nil {
		return err
	}

	return nil
}

func (o *AccountModifyCollectionBody) validateOwner(formats strfmt.Registry) error {
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

func (o *AccountModifyCollectionBody) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(o.Password) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"password", "body", o.Password.String(), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"password", "body", o.Password.String(), 128); err != nil {
		return err
	}

	if err := validate.FormatOf("info"+"."+"password", "body", "password", o.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

var accountModifyCollectionBodyTypePasswordHashAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha512","sha256","md5"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountModifyCollectionBodyTypePasswordHashAlgorithmPropEnum = append(accountModifyCollectionBodyTypePasswordHashAlgorithmPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// AccountModifyCollectionBody
	// AccountModifyCollectionBody
	// password_hash_algorithm
	// PasswordHashAlgorithm
	// sha512
	// END DEBUGGING
	// AccountModifyCollectionBodyPasswordHashAlgorithmSha512 captures enum value "sha512"
	AccountModifyCollectionBodyPasswordHashAlgorithmSha512 string = "sha512"

	// BEGIN DEBUGGING
	// AccountModifyCollectionBody
	// AccountModifyCollectionBody
	// password_hash_algorithm
	// PasswordHashAlgorithm
	// sha256
	// END DEBUGGING
	// AccountModifyCollectionBodyPasswordHashAlgorithmSha256 captures enum value "sha256"
	AccountModifyCollectionBodyPasswordHashAlgorithmSha256 string = "sha256"

	// BEGIN DEBUGGING
	// AccountModifyCollectionBody
	// AccountModifyCollectionBody
	// password_hash_algorithm
	// PasswordHashAlgorithm
	// md5
	// END DEBUGGING
	// AccountModifyCollectionBodyPasswordHashAlgorithmMd5 captures enum value "md5"
	AccountModifyCollectionBodyPasswordHashAlgorithmMd5 string = "md5"
)

// prop value enum
func (o *AccountModifyCollectionBody) validatePasswordHashAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountModifyCollectionBodyTypePasswordHashAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AccountModifyCollectionBody) validatePasswordHashAlgorithm(formats strfmt.Registry) error {
	if swag.IsZero(o.PasswordHashAlgorithm) { // not required
		return nil
	}

	// value enum
	if err := o.validatePasswordHashAlgorithmEnum("info"+"."+"password_hash_algorithm", "body", *o.PasswordHashAlgorithm); err != nil {
		return err
	}

	return nil
}

func (o *AccountModifyCollectionBody) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(o.Role) { // not required
		return nil
	}

	if o.Role != nil {
		if err := o.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "role")
			}
			return err
		}
	}

	return nil
}

var accountModifyCollectionBodyTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster","svm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountModifyCollectionBodyTypeScopePropEnum = append(accountModifyCollectionBodyTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// AccountModifyCollectionBody
	// AccountModifyCollectionBody
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// AccountModifyCollectionBodyScopeCluster captures enum value "cluster"
	AccountModifyCollectionBodyScopeCluster string = "cluster"

	// BEGIN DEBUGGING
	// AccountModifyCollectionBody
	// AccountModifyCollectionBody
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// AccountModifyCollectionBodyScopeSvm captures enum value "svm"
	AccountModifyCollectionBodyScopeSvm string = "svm"
)

// prop value enum
func (o *AccountModifyCollectionBody) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountModifyCollectionBodyTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AccountModifyCollectionBody) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(o.Scope) { // not required
		return nil
	}

	// value enum
	if err := o.validateScopeEnum("info"+"."+"scope", "body", *o.Scope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this account modify collection body based on the context it is used
func (o *AccountModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAccountInlineApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAccountResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePasswordHashAlgorithm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRole(ctx, formats); err != nil {
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

func (o *AccountModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *AccountModifyCollectionBody) contextValidateAccountInlineApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AccountInlineApplications); i++ {

		if o.AccountInlineApplications[i] != nil {
			if err := o.AccountInlineApplications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AccountModifyCollectionBody) contextValidateAccountResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AccountResponseInlineRecords); i++ {

		if o.AccountResponseInlineRecords[i] != nil {
			if err := o.AccountResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AccountModifyCollectionBody) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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

func (o *AccountModifyCollectionBody) contextValidatePasswordHashAlgorithm(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"password_hash_algorithm", "body", o.PasswordHashAlgorithm); err != nil {
		return err
	}

	return nil
}

func (o *AccountModifyCollectionBody) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if o.Role != nil {
		if err := o.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "role")
			}
			return err
		}
	}

	return nil
}

func (o *AccountModifyCollectionBody) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"scope", "body", o.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccountModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccountModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res AccountModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccountInlineLinks account inline links
swagger:model account_inline__links
*/
type AccountInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this account inline links
func (o *AccountInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this account inline links based on the context it is used
func (o *AccountInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *AccountInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccountInlineLinks) UnmarshalBinary(b []byte) error {
	var res AccountInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccountInlineOwner Owner name and UUID that uniquely identifies the user account.
swagger:model account_inline_owner
*/
type AccountInlineOwner struct {

	// links
	Links *models.AccountInlineOwnerInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this account inline owner
func (o *AccountInlineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineOwner) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this account inline owner based on the context it is used
func (o *AccountInlineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *AccountInlineOwner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccountInlineOwner) UnmarshalBinary(b []byte) error {
	var res AccountInlineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccountInlineOwnerInlineLinks account inline owner inline links
swagger:model account_inline_owner_inline__links
*/
type AccountInlineOwnerInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this account inline owner inline links
func (o *AccountInlineOwnerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineOwnerInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this account inline owner inline links based on the context it is used
func (o *AccountInlineOwnerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineOwnerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *AccountInlineOwnerInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccountInlineOwnerInlineLinks) UnmarshalBinary(b []byte) error {
	var res AccountInlineOwnerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccountInlineRole account inline role
swagger:model account_inline_role
*/
type AccountInlineRole struct {

	// links
	Links *models.AccountInlineRoleInlineLinks `json:"_links,omitempty"`

	// Role name
	// Example: admin
	Name *string `json:"name,omitempty"`
}

// Validate validates this account inline role
func (o *AccountInlineRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineRole) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "role" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account inline role based on the context it is used
func (o *AccountInlineRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineRole) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "role" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccountInlineRole) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccountInlineRole) UnmarshalBinary(b []byte) error {
	var res AccountInlineRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccountInlineRoleInlineLinks account inline role inline links
swagger:model account_inline_role_inline__links
*/
type AccountInlineRoleInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this account inline role inline links
func (o *AccountInlineRoleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineRoleInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "role" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account inline role inline links based on the context it is used
func (o *AccountInlineRoleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountInlineRoleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "role" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccountInlineRoleInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccountInlineRoleInlineLinks) UnmarshalBinary(b []byte) error {
	var res AccountInlineRoleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
