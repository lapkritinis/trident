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

// StoragePoolGetReader is a Reader for the StoragePoolGet structure.
type StoragePoolGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StoragePoolGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStoragePoolGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStoragePoolGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStoragePoolGetOK creates a StoragePoolGetOK with default headers values
func NewStoragePoolGetOK() *StoragePoolGetOK {
	return &StoragePoolGetOK{}
}

/*
StoragePoolGetOK describes a response with status code 200, with default header values.

OK
*/
type StoragePoolGetOK struct {
	Payload *models.StoragePool
}

// IsSuccess returns true when this storage pool get o k response has a 2xx status code
func (o *StoragePoolGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage pool get o k response has a 3xx status code
func (o *StoragePoolGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage pool get o k response has a 4xx status code
func (o *StoragePoolGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage pool get o k response has a 5xx status code
func (o *StoragePoolGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this storage pool get o k response a status code equal to that given
func (o *StoragePoolGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the storage pool get o k response
func (o *StoragePoolGetOK) Code() int {
	return 200
}

func (o *StoragePoolGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/pools/{uuid}][%d] storagePoolGetOK %s", 200, payload)
}

func (o *StoragePoolGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/pools/{uuid}][%d] storagePoolGetOK %s", 200, payload)
}

func (o *StoragePoolGetOK) GetPayload() *models.StoragePool {
	return o.Payload
}

func (o *StoragePoolGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoragePoolGetDefault creates a StoragePoolGetDefault with default headers values
func NewStoragePoolGetDefault(code int) *StoragePoolGetDefault {
	return &StoragePoolGetDefault{
		_statusCode: code,
	}
}

/*
	StoragePoolGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 11206662 | There is no storage pool matching the specified UUID or name. |
| 11215856 | The specified storage pool is not healthy. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type StoragePoolGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this storage pool get default response has a 2xx status code
func (o *StoragePoolGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this storage pool get default response has a 3xx status code
func (o *StoragePoolGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this storage pool get default response has a 4xx status code
func (o *StoragePoolGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this storage pool get default response has a 5xx status code
func (o *StoragePoolGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this storage pool get default response a status code equal to that given
func (o *StoragePoolGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the storage pool get default response
func (o *StoragePoolGetDefault) Code() int {
	return o._statusCode
}

func (o *StoragePoolGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/pools/{uuid}][%d] storage_pool_get default %s", o._statusCode, payload)
}

func (o *StoragePoolGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/pools/{uuid}][%d] storage_pool_get default %s", o._statusCode, payload)
}

func (o *StoragePoolGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StoragePoolGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
