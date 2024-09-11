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

// SvmSSHServerCollectionGetReader is a Reader for the SvmSSHServerCollectionGet structure.
type SvmSSHServerCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmSSHServerCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmSSHServerCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmSSHServerCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmSSHServerCollectionGetOK creates a SvmSSHServerCollectionGetOK with default headers values
func NewSvmSSHServerCollectionGetOK() *SvmSSHServerCollectionGetOK {
	return &SvmSSHServerCollectionGetOK{}
}

/*
SvmSSHServerCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SvmSSHServerCollectionGetOK struct {
	Payload *models.SvmSSHServerResponse
}

// IsSuccess returns true when this svm Ssh server collection get o k response has a 2xx status code
func (o *SvmSSHServerCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm Ssh server collection get o k response has a 3xx status code
func (o *SvmSSHServerCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm Ssh server collection get o k response has a 4xx status code
func (o *SvmSSHServerCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm Ssh server collection get o k response has a 5xx status code
func (o *SvmSSHServerCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm Ssh server collection get o k response a status code equal to that given
func (o *SvmSSHServerCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm Ssh server collection get o k response
func (o *SvmSSHServerCollectionGetOK) Code() int {
	return 200
}

func (o *SvmSSHServerCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/ssh/svms][%d] svmSshServerCollectionGetOK %s", 200, payload)
}

func (o *SvmSSHServerCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/ssh/svms][%d] svmSshServerCollectionGetOK %s", 200, payload)
}

func (o *SvmSSHServerCollectionGetOK) GetPayload() *models.SvmSSHServerResponse {
	return o.Payload
}

func (o *SvmSSHServerCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SvmSSHServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmSSHServerCollectionGetDefault creates a SvmSSHServerCollectionGetDefault with default headers values
func NewSvmSSHServerCollectionGetDefault(code int) *SvmSSHServerCollectionGetDefault {
	return &SvmSSHServerCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SvmSSHServerCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SvmSSHServerCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm ssh server collection get default response has a 2xx status code
func (o *SvmSSHServerCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm ssh server collection get default response has a 3xx status code
func (o *SvmSSHServerCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm ssh server collection get default response has a 4xx status code
func (o *SvmSSHServerCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm ssh server collection get default response has a 5xx status code
func (o *SvmSSHServerCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm ssh server collection get default response a status code equal to that given
func (o *SvmSSHServerCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm ssh server collection get default response
func (o *SvmSSHServerCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SvmSSHServerCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/ssh/svms][%d] svm_ssh_server_collection_get default %s", o._statusCode, payload)
}

func (o *SvmSSHServerCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/ssh/svms][%d] svm_ssh_server_collection_get default %s", o._statusCode, payload)
}

func (o *SvmSSHServerCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmSSHServerCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
