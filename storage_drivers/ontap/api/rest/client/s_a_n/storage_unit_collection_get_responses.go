// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// StorageUnitCollectionGetReader is a Reader for the StorageUnitCollectionGet structure.
type StorageUnitCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageUnitCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageUnitCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageUnitCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageUnitCollectionGetOK creates a StorageUnitCollectionGetOK with default headers values
func NewStorageUnitCollectionGetOK() *StorageUnitCollectionGetOK {
	return &StorageUnitCollectionGetOK{}
}

/*
StorageUnitCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type StorageUnitCollectionGetOK struct {
	Payload *models.StorageUnitResponse
}

// IsSuccess returns true when this storage unit collection get o k response has a 2xx status code
func (o *StorageUnitCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage unit collection get o k response has a 3xx status code
func (o *StorageUnitCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage unit collection get o k response has a 4xx status code
func (o *StorageUnitCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage unit collection get o k response has a 5xx status code
func (o *StorageUnitCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this storage unit collection get o k response a status code equal to that given
func (o *StorageUnitCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the storage unit collection get o k response
func (o *StorageUnitCollectionGetOK) Code() int {
	return 200
}

func (o *StorageUnitCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/storage-units][%d] storageUnitCollectionGetOK %s", 200, payload)
}

func (o *StorageUnitCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/storage-units][%d] storageUnitCollectionGetOK %s", 200, payload)
}

func (o *StorageUnitCollectionGetOK) GetPayload() *models.StorageUnitResponse {
	return o.Payload
}

func (o *StorageUnitCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageUnitResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageUnitCollectionGetDefault creates a StorageUnitCollectionGetDefault with default headers values
func NewStorageUnitCollectionGetDefault(code int) *StorageUnitCollectionGetDefault {
	return &StorageUnitCollectionGetDefault{
		_statusCode: code,
	}
}

/*
StorageUnitCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type StorageUnitCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this storage unit collection get default response has a 2xx status code
func (o *StorageUnitCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this storage unit collection get default response has a 3xx status code
func (o *StorageUnitCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this storage unit collection get default response has a 4xx status code
func (o *StorageUnitCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this storage unit collection get default response has a 5xx status code
func (o *StorageUnitCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this storage unit collection get default response a status code equal to that given
func (o *StorageUnitCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the storage unit collection get default response
func (o *StorageUnitCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageUnitCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/storage-units][%d] storage_unit_collection_get default %s", o._statusCode, payload)
}

func (o *StorageUnitCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/storage-units][%d] storage_unit_collection_get default %s", o._statusCode, payload)
}

func (o *StorageUnitCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StorageUnitCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
