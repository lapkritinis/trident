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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AccountPublickeyDeleteCollectionReader is a Reader for the AccountPublickeyDeleteCollection structure.
type AccountPublickeyDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountPublickeyDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountPublickeyDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountPublickeyDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountPublickeyDeleteCollectionOK creates a AccountPublickeyDeleteCollectionOK with default headers values
func NewAccountPublickeyDeleteCollectionOK() *AccountPublickeyDeleteCollectionOK {
	return &AccountPublickeyDeleteCollectionOK{}
}

/*
AccountPublickeyDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type AccountPublickeyDeleteCollectionOK struct {
}

// IsSuccess returns true when this account publickey delete collection o k response has a 2xx status code
func (o *AccountPublickeyDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account publickey delete collection o k response has a 3xx status code
func (o *AccountPublickeyDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account publickey delete collection o k response has a 4xx status code
func (o *AccountPublickeyDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account publickey delete collection o k response has a 5xx status code
func (o *AccountPublickeyDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account publickey delete collection o k response a status code equal to that given
func (o *AccountPublickeyDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account publickey delete collection o k response
func (o *AccountPublickeyDeleteCollectionOK) Code() int {
	return 200
}

func (o *AccountPublickeyDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /security/authentication/publickeys][%d] accountPublickeyDeleteCollectionOK", 200)
}

func (o *AccountPublickeyDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /security/authentication/publickeys][%d] accountPublickeyDeleteCollectionOK", 200)
}

func (o *AccountPublickeyDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountPublickeyDeleteCollectionDefault creates a AccountPublickeyDeleteCollectionDefault with default headers values
func NewAccountPublickeyDeleteCollectionDefault(code int) *AccountPublickeyDeleteCollectionDefault {
	return &AccountPublickeyDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
AccountPublickeyDeleteCollectionDefault describes a response with status code -1, with default header values.

Error
*/
type AccountPublickeyDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this account publickey delete collection default response has a 2xx status code
func (o *AccountPublickeyDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this account publickey delete collection default response has a 3xx status code
func (o *AccountPublickeyDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this account publickey delete collection default response has a 4xx status code
func (o *AccountPublickeyDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this account publickey delete collection default response has a 5xx status code
func (o *AccountPublickeyDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this account publickey delete collection default response a status code equal to that given
func (o *AccountPublickeyDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the account publickey delete collection default response
func (o *AccountPublickeyDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *AccountPublickeyDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/authentication/publickeys][%d] account_publickey_delete_collection default %s", o._statusCode, payload)
}

func (o *AccountPublickeyDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/authentication/publickeys][%d] account_publickey_delete_collection default %s", o._statusCode, payload)
}

func (o *AccountPublickeyDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountPublickeyDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AccountPublickeyDeleteCollectionBody account publickey delete collection body
swagger:model AccountPublickeyDeleteCollectionBody
*/
type AccountPublickeyDeleteCollectionBody struct {

	// publickey response inline records
	PublickeyResponseInlineRecords []*models.Publickey `json:"records,omitempty"`
}

// Validate validates this account publickey delete collection body
func (o *AccountPublickeyDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePublickeyResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountPublickeyDeleteCollectionBody) validatePublickeyResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.PublickeyResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.PublickeyResponseInlineRecords); i++ {
		if swag.IsZero(o.PublickeyResponseInlineRecords[i]) { // not required
			continue
		}

		if o.PublickeyResponseInlineRecords[i] != nil {
			if err := o.PublickeyResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this account publickey delete collection body based on the context it is used
func (o *AccountPublickeyDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePublickeyResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountPublickeyDeleteCollectionBody) contextValidatePublickeyResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PublickeyResponseInlineRecords); i++ {

		if o.PublickeyResponseInlineRecords[i] != nil {
			if err := o.PublickeyResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *AccountPublickeyDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccountPublickeyDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res AccountPublickeyDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
