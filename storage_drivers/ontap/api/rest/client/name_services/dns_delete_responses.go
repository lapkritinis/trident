// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// DNSDeleteReader is a Reader for the DNSDelete structure.
type DNSDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDNSDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDNSDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDNSDeleteOK creates a DNSDeleteOK with default headers values
func NewDNSDeleteOK() *DNSDeleteOK {
	return &DNSDeleteOK{}
}

/*
DNSDeleteOK describes a response with status code 200, with default header values.

OK
*/
type DNSDeleteOK struct {
}

// IsSuccess returns true when this dns delete o k response has a 2xx status code
func (o *DNSDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns delete o k response has a 3xx status code
func (o *DNSDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns delete o k response has a 4xx status code
func (o *DNSDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns delete o k response has a 5xx status code
func (o *DNSDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dns delete o k response a status code equal to that given
func (o *DNSDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dns delete o k response
func (o *DNSDeleteOK) Code() int {
	return 200
}

func (o *DNSDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /name-services/dns/{uuid}][%d] dnsDeleteOK", 200)
}

func (o *DNSDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /name-services/dns/{uuid}][%d] dnsDeleteOK", 200)
}

func (o *DNSDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDNSDeleteDefault creates a DNSDeleteDefault with default headers values
func NewDNSDeleteDefault(code int) *DNSDeleteDefault {
	return &DNSDeleteDefault{
		_statusCode: code,
	}
}

/*
DNSDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type DNSDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this dns delete default response has a 2xx status code
func (o *DNSDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dns delete default response has a 3xx status code
func (o *DNSDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dns delete default response has a 4xx status code
func (o *DNSDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dns delete default response has a 5xx status code
func (o *DNSDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dns delete default response a status code equal to that given
func (o *DNSDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dns delete default response
func (o *DNSDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DNSDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/dns/{uuid}][%d] dns_delete default %s", o._statusCode, payload)
}

func (o *DNSDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/dns/{uuid}][%d] dns_delete default %s", o._statusCode, payload)
}

func (o *DNSDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DNSDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
