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

// IpsecCaCertificateDeleteReader is a Reader for the IpsecCaCertificateDelete structure.
type IpsecCaCertificateDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpsecCaCertificateDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpsecCaCertificateDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpsecCaCertificateDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpsecCaCertificateDeleteOK creates a IpsecCaCertificateDeleteOK with default headers values
func NewIpsecCaCertificateDeleteOK() *IpsecCaCertificateDeleteOK {
	return &IpsecCaCertificateDeleteOK{}
}

/*
IpsecCaCertificateDeleteOK describes a response with status code 200, with default header values.

OK
*/
type IpsecCaCertificateDeleteOK struct {
}

// IsSuccess returns true when this ipsec ca certificate delete o k response has a 2xx status code
func (o *IpsecCaCertificateDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipsec ca certificate delete o k response has a 3xx status code
func (o *IpsecCaCertificateDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipsec ca certificate delete o k response has a 4xx status code
func (o *IpsecCaCertificateDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipsec ca certificate delete o k response has a 5xx status code
func (o *IpsecCaCertificateDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipsec ca certificate delete o k response a status code equal to that given
func (o *IpsecCaCertificateDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipsec ca certificate delete o k response
func (o *IpsecCaCertificateDeleteOK) Code() int {
	return 200
}

func (o *IpsecCaCertificateDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsecCaCertificateDeleteOK", 200)
}

func (o *IpsecCaCertificateDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsecCaCertificateDeleteOK", 200)
}

func (o *IpsecCaCertificateDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpsecCaCertificateDeleteDefault creates a IpsecCaCertificateDeleteDefault with default headers values
func NewIpsecCaCertificateDeleteDefault(code int) *IpsecCaCertificateDeleteDefault {
	return &IpsecCaCertificateDeleteDefault{
		_statusCode: code,
	}
}

/*
	IpsecCaCertificateDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 66257298 | CA certificate is not installed for IPsec. |
| 66257303 | The CA certificate cannot be removed from IPsec because it is not installed. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IpsecCaCertificateDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ipsec ca certificate delete default response has a 2xx status code
func (o *IpsecCaCertificateDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipsec ca certificate delete default response has a 3xx status code
func (o *IpsecCaCertificateDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipsec ca certificate delete default response has a 4xx status code
func (o *IpsecCaCertificateDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipsec ca certificate delete default response has a 5xx status code
func (o *IpsecCaCertificateDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipsec ca certificate delete default response a status code equal to that given
func (o *IpsecCaCertificateDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipsec ca certificate delete default response
func (o *IpsecCaCertificateDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpsecCaCertificateDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsec_ca_certificate_delete default %s", o._statusCode, payload)
}

func (o *IpsecCaCertificateDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsec_ca_certificate_delete default %s", o._statusCode, payload)
}

func (o *IpsecCaCertificateDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IpsecCaCertificateDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
