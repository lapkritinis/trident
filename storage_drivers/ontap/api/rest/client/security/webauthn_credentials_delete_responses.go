// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// WebauthnCredentialsDeleteReader is a Reader for the WebauthnCredentialsDelete structure.
type WebauthnCredentialsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebauthnCredentialsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebauthnCredentialsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebauthnCredentialsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebauthnCredentialsDeleteOK creates a WebauthnCredentialsDeleteOK with default headers values
func NewWebauthnCredentialsDeleteOK() *WebauthnCredentialsDeleteOK {
	return &WebauthnCredentialsDeleteOK{}
}

/*
WebauthnCredentialsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type WebauthnCredentialsDeleteOK struct {
}

// IsSuccess returns true when this webauthn credentials delete o k response has a 2xx status code
func (o *WebauthnCredentialsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webauthn credentials delete o k response has a 3xx status code
func (o *WebauthnCredentialsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webauthn credentials delete o k response has a 4xx status code
func (o *WebauthnCredentialsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webauthn credentials delete o k response has a 5xx status code
func (o *WebauthnCredentialsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webauthn credentials delete o k response a status code equal to that given
func (o *WebauthnCredentialsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webauthn credentials delete o k response
func (o *WebauthnCredentialsDeleteOK) Code() int {
	return 200
}

func (o *WebauthnCredentialsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/webauthn/credentials/{owner.uuid}/{username}/{index}/{relying_party.id}][%d] webauthnCredentialsDeleteOK", 200)
}

func (o *WebauthnCredentialsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/webauthn/credentials/{owner.uuid}/{username}/{index}/{relying_party.id}][%d] webauthnCredentialsDeleteOK", 200)
}

func (o *WebauthnCredentialsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWebauthnCredentialsDeleteDefault creates a WebauthnCredentialsDeleteDefault with default headers values
func NewWebauthnCredentialsDeleteDefault(code int) *WebauthnCredentialsDeleteDefault {
	return &WebauthnCredentialsDeleteDefault{
		_statusCode: code,
	}
}

/*
WebauthnCredentialsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type WebauthnCredentialsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this webauthn credentials delete default response has a 2xx status code
func (o *WebauthnCredentialsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this webauthn credentials delete default response has a 3xx status code
func (o *WebauthnCredentialsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this webauthn credentials delete default response has a 4xx status code
func (o *WebauthnCredentialsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this webauthn credentials delete default response has a 5xx status code
func (o *WebauthnCredentialsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this webauthn credentials delete default response a status code equal to that given
func (o *WebauthnCredentialsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the webauthn credentials delete default response
func (o *WebauthnCredentialsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *WebauthnCredentialsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/webauthn/credentials/{owner.uuid}/{username}/{index}/{relying_party.id}][%d] webauthn_credentials_delete default %s", o._statusCode, payload)
}

func (o *WebauthnCredentialsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/webauthn/credentials/{owner.uuid}/{username}/{index}/{relying_party.id}][%d] webauthn_credentials_delete default %s", o._statusCode, payload)
}

func (o *WebauthnCredentialsDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WebauthnCredentialsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
