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

// SecurityKeyManagerKeyServersDeleteReader is a Reader for the SecurityKeyManagerKeyServersDelete structure.
type SecurityKeyManagerKeyServersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerKeyServersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerKeyServersDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerKeyServersDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerKeyServersDeleteOK creates a SecurityKeyManagerKeyServersDeleteOK with default headers values
func NewSecurityKeyManagerKeyServersDeleteOK() *SecurityKeyManagerKeyServersDeleteOK {
	return &SecurityKeyManagerKeyServersDeleteOK{}
}

/*
SecurityKeyManagerKeyServersDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerKeyServersDeleteOK struct {
}

// IsSuccess returns true when this security key manager key servers delete o k response has a 2xx status code
func (o *SecurityKeyManagerKeyServersDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager key servers delete o k response has a 3xx status code
func (o *SecurityKeyManagerKeyServersDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager key servers delete o k response has a 4xx status code
func (o *SecurityKeyManagerKeyServersDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager key servers delete o k response has a 5xx status code
func (o *SecurityKeyManagerKeyServersDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager key servers delete o k response a status code equal to that given
func (o *SecurityKeyManagerKeyServersDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security key manager key servers delete o k response
func (o *SecurityKeyManagerKeyServersDeleteOK) Code() int {
	return 200
}

func (o *SecurityKeyManagerKeyServersDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}/key-servers/{server}][%d] securityKeyManagerKeyServersDeleteOK", 200)
}

func (o *SecurityKeyManagerKeyServersDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}/key-servers/{server}][%d] securityKeyManagerKeyServersDeleteOK", 200)
}

func (o *SecurityKeyManagerKeyServersDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityKeyManagerKeyServersDeleteDefault creates a SecurityKeyManagerKeyServersDeleteDefault with default headers values
func NewSecurityKeyManagerKeyServersDeleteDefault(code int) *SecurityKeyManagerKeyServersDeleteDefault {
	return &SecurityKeyManagerKeyServersDeleteDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeyManagerKeyServersDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536600 | Cannot remove a key server while a node is out of quorum. |
| 65536700 | The key server contains keys that are currently in use and not available from any other configured key server in the SVM. |
| 65536843 | The key management server is not configured for the SVM. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeyManagerKeyServersDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager key servers delete default response has a 2xx status code
func (o *SecurityKeyManagerKeyServersDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager key servers delete default response has a 3xx status code
func (o *SecurityKeyManagerKeyServersDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager key servers delete default response has a 4xx status code
func (o *SecurityKeyManagerKeyServersDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager key servers delete default response has a 5xx status code
func (o *SecurityKeyManagerKeyServersDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager key servers delete default response a status code equal to that given
func (o *SecurityKeyManagerKeyServersDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager key servers delete default response
func (o *SecurityKeyManagerKeyServersDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerKeyServersDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}/key-servers/{server}][%d] security_key_manager_key_servers_delete default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}/key-servers/{server}][%d] security_key_manager_key_servers_delete default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
