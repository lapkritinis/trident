// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// StorageBridgeGetReader is a Reader for the StorageBridgeGet structure.
type StorageBridgeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageBridgeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageBridgeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageBridgeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageBridgeGetOK creates a StorageBridgeGetOK with default headers values
func NewStorageBridgeGetOK() *StorageBridgeGetOK {
	return &StorageBridgeGetOK{}
}

/*
StorageBridgeGetOK describes a response with status code 200, with default header values.

OK
*/
type StorageBridgeGetOK struct {
	Payload *models.StorageBridge
}

// IsSuccess returns true when this storage bridge get o k response has a 2xx status code
func (o *StorageBridgeGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage bridge get o k response has a 3xx status code
func (o *StorageBridgeGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage bridge get o k response has a 4xx status code
func (o *StorageBridgeGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage bridge get o k response has a 5xx status code
func (o *StorageBridgeGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this storage bridge get o k response a status code equal to that given
func (o *StorageBridgeGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the storage bridge get o k response
func (o *StorageBridgeGetOK) Code() int {
	return 200
}

func (o *StorageBridgeGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/bridges/{wwn}][%d] storageBridgeGetOK %s", 200, payload)
}

func (o *StorageBridgeGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/bridges/{wwn}][%d] storageBridgeGetOK %s", 200, payload)
}

func (o *StorageBridgeGetOK) GetPayload() *models.StorageBridge {
	return o.Payload
}

func (o *StorageBridgeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageBridge)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageBridgeGetDefault creates a StorageBridgeGetDefault with default headers values
func NewStorageBridgeGetDefault(code int) *StorageBridgeGetDefault {
	return &StorageBridgeGetDefault{
		_statusCode: code,
	}
}

/*
StorageBridgeGetDefault describes a response with status code -1, with default header values.

Error
*/
type StorageBridgeGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this storage bridge get default response has a 2xx status code
func (o *StorageBridgeGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this storage bridge get default response has a 3xx status code
func (o *StorageBridgeGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this storage bridge get default response has a 4xx status code
func (o *StorageBridgeGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this storage bridge get default response has a 5xx status code
func (o *StorageBridgeGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this storage bridge get default response a status code equal to that given
func (o *StorageBridgeGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the storage bridge get default response
func (o *StorageBridgeGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageBridgeGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/bridges/{wwn}][%d] storage_bridge_get default %s", o._statusCode, payload)
}

func (o *StorageBridgeGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/bridges/{wwn}][%d] storage_bridge_get default %s", o._statusCode, payload)
}

func (o *StorageBridgeGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StorageBridgeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
