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

// MultiAdminVerifyRequestCollectionGetReader is a Reader for the MultiAdminVerifyRequestCollectionGet structure.
type MultiAdminVerifyRequestCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiAdminVerifyRequestCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiAdminVerifyRequestCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMultiAdminVerifyRequestCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMultiAdminVerifyRequestCollectionGetOK creates a MultiAdminVerifyRequestCollectionGetOK with default headers values
func NewMultiAdminVerifyRequestCollectionGetOK() *MultiAdminVerifyRequestCollectionGetOK {
	return &MultiAdminVerifyRequestCollectionGetOK{}
}

/*
MultiAdminVerifyRequestCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type MultiAdminVerifyRequestCollectionGetOK struct {
	Payload *models.MultiAdminVerifyRequestResponse
}

// IsSuccess returns true when this multi admin verify request collection get o k response has a 2xx status code
func (o *MultiAdminVerifyRequestCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this multi admin verify request collection get o k response has a 3xx status code
func (o *MultiAdminVerifyRequestCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this multi admin verify request collection get o k response has a 4xx status code
func (o *MultiAdminVerifyRequestCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this multi admin verify request collection get o k response has a 5xx status code
func (o *MultiAdminVerifyRequestCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this multi admin verify request collection get o k response a status code equal to that given
func (o *MultiAdminVerifyRequestCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the multi admin verify request collection get o k response
func (o *MultiAdminVerifyRequestCollectionGetOK) Code() int {
	return 200
}

func (o *MultiAdminVerifyRequestCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/requests][%d] multiAdminVerifyRequestCollectionGetOK %s", 200, payload)
}

func (o *MultiAdminVerifyRequestCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/requests][%d] multiAdminVerifyRequestCollectionGetOK %s", 200, payload)
}

func (o *MultiAdminVerifyRequestCollectionGetOK) GetPayload() *models.MultiAdminVerifyRequestResponse {
	return o.Payload
}

func (o *MultiAdminVerifyRequestCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MultiAdminVerifyRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiAdminVerifyRequestCollectionGetDefault creates a MultiAdminVerifyRequestCollectionGetDefault with default headers values
func NewMultiAdminVerifyRequestCollectionGetDefault(code int) *MultiAdminVerifyRequestCollectionGetDefault {
	return &MultiAdminVerifyRequestCollectionGetDefault{
		_statusCode: code,
	}
}

/*
MultiAdminVerifyRequestCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type MultiAdminVerifyRequestCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this multi admin verify request collection get default response has a 2xx status code
func (o *MultiAdminVerifyRequestCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this multi admin verify request collection get default response has a 3xx status code
func (o *MultiAdminVerifyRequestCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this multi admin verify request collection get default response has a 4xx status code
func (o *MultiAdminVerifyRequestCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this multi admin verify request collection get default response has a 5xx status code
func (o *MultiAdminVerifyRequestCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this multi admin verify request collection get default response a status code equal to that given
func (o *MultiAdminVerifyRequestCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the multi admin verify request collection get default response
func (o *MultiAdminVerifyRequestCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *MultiAdminVerifyRequestCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/requests][%d] multi_admin_verify_request_collection_get default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRequestCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/requests][%d] multi_admin_verify_request_collection_get default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRequestCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MultiAdminVerifyRequestCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
