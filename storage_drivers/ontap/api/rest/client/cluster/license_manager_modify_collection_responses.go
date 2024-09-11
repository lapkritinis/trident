// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// LicenseManagerModifyCollectionReader is a Reader for the LicenseManagerModifyCollection structure.
type LicenseManagerModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LicenseManagerModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLicenseManagerModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewLicenseManagerModifyCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLicenseManagerModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLicenseManagerModifyCollectionOK creates a LicenseManagerModifyCollectionOK with default headers values
func NewLicenseManagerModifyCollectionOK() *LicenseManagerModifyCollectionOK {
	return &LicenseManagerModifyCollectionOK{}
}

/*
LicenseManagerModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type LicenseManagerModifyCollectionOK struct {
	Payload *models.LicenseManagerJobLinkResponse
}

// IsSuccess returns true when this license manager modify collection o k response has a 2xx status code
func (o *LicenseManagerModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this license manager modify collection o k response has a 3xx status code
func (o *LicenseManagerModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this license manager modify collection o k response has a 4xx status code
func (o *LicenseManagerModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this license manager modify collection o k response has a 5xx status code
func (o *LicenseManagerModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this license manager modify collection o k response a status code equal to that given
func (o *LicenseManagerModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the license manager modify collection o k response
func (o *LicenseManagerModifyCollectionOK) Code() int {
	return 200
}

func (o *LicenseManagerModifyCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/licensing/license-managers][%d] licenseManagerModifyCollectionOK %s", 200, payload)
}

func (o *LicenseManagerModifyCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/licensing/license-managers][%d] licenseManagerModifyCollectionOK %s", 200, payload)
}

func (o *LicenseManagerModifyCollectionOK) GetPayload() *models.LicenseManagerJobLinkResponse {
	return o.Payload
}

func (o *LicenseManagerModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseManagerJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLicenseManagerModifyCollectionAccepted creates a LicenseManagerModifyCollectionAccepted with default headers values
func NewLicenseManagerModifyCollectionAccepted() *LicenseManagerModifyCollectionAccepted {
	return &LicenseManagerModifyCollectionAccepted{}
}

/*
LicenseManagerModifyCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type LicenseManagerModifyCollectionAccepted struct {
	Payload *models.LicenseManagerJobLinkResponse
}

// IsSuccess returns true when this license manager modify collection accepted response has a 2xx status code
func (o *LicenseManagerModifyCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this license manager modify collection accepted response has a 3xx status code
func (o *LicenseManagerModifyCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this license manager modify collection accepted response has a 4xx status code
func (o *LicenseManagerModifyCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this license manager modify collection accepted response has a 5xx status code
func (o *LicenseManagerModifyCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this license manager modify collection accepted response a status code equal to that given
func (o *LicenseManagerModifyCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the license manager modify collection accepted response
func (o *LicenseManagerModifyCollectionAccepted) Code() int {
	return 202
}

func (o *LicenseManagerModifyCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/licensing/license-managers][%d] licenseManagerModifyCollectionAccepted %s", 202, payload)
}

func (o *LicenseManagerModifyCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/licensing/license-managers][%d] licenseManagerModifyCollectionAccepted %s", 202, payload)
}

func (o *LicenseManagerModifyCollectionAccepted) GetPayload() *models.LicenseManagerJobLinkResponse {
	return o.Payload
}

func (o *LicenseManagerModifyCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseManagerJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLicenseManagerModifyCollectionDefault creates a LicenseManagerModifyCollectionDefault with default headers values
func NewLicenseManagerModifyCollectionDefault(code int) *LicenseManagerModifyCollectionDefault {
	return &LicenseManagerModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	LicenseManagerModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1115532 | The requested update to the license manager information failed. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LicenseManagerModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponses
}

// IsSuccess returns true when this license manager modify collection default response has a 2xx status code
func (o *LicenseManagerModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this license manager modify collection default response has a 3xx status code
func (o *LicenseManagerModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this license manager modify collection default response has a 4xx status code
func (o *LicenseManagerModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this license manager modify collection default response has a 5xx status code
func (o *LicenseManagerModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this license manager modify collection default response a status code equal to that given
func (o *LicenseManagerModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the license manager modify collection default response
func (o *LicenseManagerModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *LicenseManagerModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/licensing/license-managers][%d] license_manager_modify_collection default %s", o._statusCode, payload)
}

func (o *LicenseManagerModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/licensing/license-managers][%d] license_manager_modify_collection default %s", o._statusCode, payload)
}

func (o *LicenseManagerModifyCollectionDefault) GetPayload() *models.ErrorResponses {
	return o.Payload
}

func (o *LicenseManagerModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
LicenseManagerModifyCollectionBody Information on license manager instances associated with the cluster.
swagger:model LicenseManagerModifyCollectionBody
*/
type LicenseManagerModifyCollectionBody struct {

	// links
	Links *models.SelfLink `json:"_links,omitempty"`

	// Flag that indicates whether it's the default license manager instance used by the cluster.'
	// When a capacity pool is created and if the license manager field is omitted, it is assumed that the license of the capacity pool is installed on the default license manager instance.
	//
	// Read Only: true
	Default *bool `json:"default,omitempty"`

	// license manager response inline records
	LicenseManagerResponseInlineRecords []*models.LicenseManagerResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`

	// uri
	URI *models.LicenseManagerResponseInlineRecordsInlineArrayItemInlineURI `json:"uri,omitempty"`

	// uuid
	// Example: 4ea7a442-86d1-11e0-ae1c-112233445566
	// Read Only: true
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this license manager modify collection body
func (o *LicenseManagerModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLicenseManagerResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LicenseManagerModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *LicenseManagerModifyCollectionBody) validateLicenseManagerResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.LicenseManagerResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.LicenseManagerResponseInlineRecords); i++ {
		if swag.IsZero(o.LicenseManagerResponseInlineRecords[i]) { // not required
			continue
		}

		if o.LicenseManagerResponseInlineRecords[i] != nil {
			if err := o.LicenseManagerResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *LicenseManagerModifyCollectionBody) validateURI(formats strfmt.Registry) error {
	if swag.IsZero(o.URI) { // not required
		return nil
	}

	if o.URI != nil {
		if err := o.URI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "uri")
			}
			return err
		}
	}

	return nil
}

func (o *LicenseManagerModifyCollectionBody) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(o.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("info"+"."+"uuid", "body", "uuid", o.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this license manager modify collection body based on the context it is used
func (o *LicenseManagerModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDefault(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLicenseManagerResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LicenseManagerModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *LicenseManagerModifyCollectionBody) contextValidateDefault(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"default", "body", o.Default); err != nil {
		return err
	}

	return nil
}

func (o *LicenseManagerModifyCollectionBody) contextValidateLicenseManagerResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.LicenseManagerResponseInlineRecords); i++ {

		if o.LicenseManagerResponseInlineRecords[i] != nil {
			if err := o.LicenseManagerResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *LicenseManagerModifyCollectionBody) contextValidateURI(ctx context.Context, formats strfmt.Registry) error {

	if o.URI != nil {
		if err := o.URI.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "uri")
			}
			return err
		}
	}

	return nil
}

func (o *LicenseManagerModifyCollectionBody) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"uuid", "body", o.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LicenseManagerModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LicenseManagerModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res LicenseManagerModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
LicenseManagerResponseInlineRecordsInlineArrayItem Information on a license manager instance associated with the cluster.
swagger:model license_manager_response_inline_records_inline_array_item
*/
type LicenseManagerResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *models.SelfLink `json:"_links,omitempty"`

	// Flag that indicates whether it's the default license manager instance used by the cluster.'
	// When a capacity pool is created and if the license manager field is omitted, it is assumed that the license of the capacity pool is installed on the default license manager instance.
	//
	// Read Only: true
	Default *bool `json:"default,omitempty"`

	// uri
	URI *models.LicenseManagerResponseInlineRecordsInlineArrayItemInlineURI `json:"uri,omitempty"`

	// uuid
	// Example: 4ea7a442-86d1-11e0-ae1c-112233445566
	// Read Only: true
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this license manager response inline records inline array item
func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) validateURI(formats strfmt.Registry) error {
	if swag.IsZero(o.URI) { // not required
		return nil
	}

	if o.URI != nil {
		if err := o.URI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uri")
			}
			return err
		}
	}

	return nil
}

func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(o.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("uuid", "body", "uuid", o.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this license manager response inline records inline array item based on the context it is used
func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDefault(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) contextValidateDefault(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "default", "body", o.Default); err != nil {
		return err
	}

	return nil
}

func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) contextValidateURI(ctx context.Context, formats strfmt.Registry) error {

	if o.URI != nil {
		if err := o.URI.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uri")
			}
			return err
		}
	}

	return nil
}

func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", o.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LicenseManagerResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res LicenseManagerResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
LicenseManagerResponseInlineRecordsInlineArrayItemInlineURI License manager URI.
swagger:model license_manager_response_inline_records_inline_array_item_inline_uri
*/
type LicenseManagerResponseInlineRecordsInlineArrayItemInlineURI struct {

	// License manager host name, IPv4 or IPv6 address.
	// Example: 10.1.1.1
	Host *string `json:"host,omitempty"`
}

// Validate validates this license manager response inline records inline array item inline uri
func (o *LicenseManagerResponseInlineRecordsInlineArrayItemInlineURI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this license manager response inline records inline array item inline uri based on context it is used
func (o *LicenseManagerResponseInlineRecordsInlineArrayItemInlineURI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LicenseManagerResponseInlineRecordsInlineArrayItemInlineURI) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LicenseManagerResponseInlineRecordsInlineArrayItemInlineURI) UnmarshalBinary(b []byte) error {
	var res LicenseManagerResponseInlineRecordsInlineArrayItemInlineURI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
