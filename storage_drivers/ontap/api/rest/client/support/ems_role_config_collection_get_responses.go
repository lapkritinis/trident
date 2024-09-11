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

// EmsRoleConfigCollectionGetReader is a Reader for the EmsRoleConfigCollectionGet structure.
type EmsRoleConfigCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsRoleConfigCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsRoleConfigCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsRoleConfigCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsRoleConfigCollectionGetOK creates a EmsRoleConfigCollectionGetOK with default headers values
func NewEmsRoleConfigCollectionGetOK() *EmsRoleConfigCollectionGetOK {
	return &EmsRoleConfigCollectionGetOK{}
}

/*
EmsRoleConfigCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type EmsRoleConfigCollectionGetOK struct {
	Payload *models.EmsRoleConfigResponse
}

// IsSuccess returns true when this ems role config collection get o k response has a 2xx status code
func (o *EmsRoleConfigCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems role config collection get o k response has a 3xx status code
func (o *EmsRoleConfigCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems role config collection get o k response has a 4xx status code
func (o *EmsRoleConfigCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems role config collection get o k response has a 5xx status code
func (o *EmsRoleConfigCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems role config collection get o k response a status code equal to that given
func (o *EmsRoleConfigCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ems role config collection get o k response
func (o *EmsRoleConfigCollectionGetOK) Code() int {
	return 200
}

func (o *EmsRoleConfigCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/role-configs][%d] emsRoleConfigCollectionGetOK %s", 200, payload)
}

func (o *EmsRoleConfigCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/role-configs][%d] emsRoleConfigCollectionGetOK %s", 200, payload)
}

func (o *EmsRoleConfigCollectionGetOK) GetPayload() *models.EmsRoleConfigResponse {
	return o.Payload
}

func (o *EmsRoleConfigCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmsRoleConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsRoleConfigCollectionGetDefault creates a EmsRoleConfigCollectionGetDefault with default headers values
func NewEmsRoleConfigCollectionGetDefault(code int) *EmsRoleConfigCollectionGetDefault {
	return &EmsRoleConfigCollectionGetDefault{
		_statusCode: code,
	}
}

/*
EmsRoleConfigCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type EmsRoleConfigCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems role config collection get default response has a 2xx status code
func (o *EmsRoleConfigCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems role config collection get default response has a 3xx status code
func (o *EmsRoleConfigCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems role config collection get default response has a 4xx status code
func (o *EmsRoleConfigCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems role config collection get default response has a 5xx status code
func (o *EmsRoleConfigCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems role config collection get default response a status code equal to that given
func (o *EmsRoleConfigCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems role config collection get default response
func (o *EmsRoleConfigCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *EmsRoleConfigCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/role-configs][%d] ems_role_config_collection_get default %s", o._statusCode, payload)
}

func (o *EmsRoleConfigCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/role-configs][%d] ems_role_config_collection_get default %s", o._statusCode, payload)
}

func (o *EmsRoleConfigCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsRoleConfigCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
