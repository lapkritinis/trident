// Code generated by go-swagger; DO NOT EDIT.

package support

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

// SnmpUsersCollectionGetReader is a Reader for the SnmpUsersCollectionGet structure.
type SnmpUsersCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnmpUsersCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnmpUsersCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnmpUsersCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnmpUsersCollectionGetOK creates a SnmpUsersCollectionGetOK with default headers values
func NewSnmpUsersCollectionGetOK() *SnmpUsersCollectionGetOK {
	return &SnmpUsersCollectionGetOK{}
}

/*
SnmpUsersCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnmpUsersCollectionGetOK struct {
	Payload *models.SnmpUserResponse
}

// IsSuccess returns true when this snmp users collection get o k response has a 2xx status code
func (o *SnmpUsersCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snmp users collection get o k response has a 3xx status code
func (o *SnmpUsersCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snmp users collection get o k response has a 4xx status code
func (o *SnmpUsersCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snmp users collection get o k response has a 5xx status code
func (o *SnmpUsersCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snmp users collection get o k response a status code equal to that given
func (o *SnmpUsersCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snmp users collection get o k response
func (o *SnmpUsersCollectionGetOK) Code() int {
	return 200
}

func (o *SnmpUsersCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/users][%d] snmpUsersCollectionGetOK %s", 200, payload)
}

func (o *SnmpUsersCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/users][%d] snmpUsersCollectionGetOK %s", 200, payload)
}

func (o *SnmpUsersCollectionGetOK) GetPayload() *models.SnmpUserResponse {
	return o.Payload
}

func (o *SnmpUsersCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnmpUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnmpUsersCollectionGetDefault creates a SnmpUsersCollectionGetDefault with default headers values
func NewSnmpUsersCollectionGetDefault(code int) *SnmpUsersCollectionGetDefault {
	return &SnmpUsersCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SnmpUsersCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnmpUsersCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snmp users collection get default response has a 2xx status code
func (o *SnmpUsersCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snmp users collection get default response has a 3xx status code
func (o *SnmpUsersCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snmp users collection get default response has a 4xx status code
func (o *SnmpUsersCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snmp users collection get default response has a 5xx status code
func (o *SnmpUsersCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snmp users collection get default response a status code equal to that given
func (o *SnmpUsersCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snmp users collection get default response
func (o *SnmpUsersCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SnmpUsersCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/users][%d] snmp_users_collection_get default %s", o._statusCode, payload)
}

func (o *SnmpUsersCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/users][%d] snmp_users_collection_get default %s", o._statusCode, payload)
}

func (o *SnmpUsersCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnmpUsersCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
