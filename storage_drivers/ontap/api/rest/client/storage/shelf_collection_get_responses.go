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

// ShelfCollectionGetReader is a Reader for the ShelfCollectionGet structure.
type ShelfCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShelfCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShelfCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewShelfCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShelfCollectionGetOK creates a ShelfCollectionGetOK with default headers values
func NewShelfCollectionGetOK() *ShelfCollectionGetOK {
	return &ShelfCollectionGetOK{}
}

/*
ShelfCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ShelfCollectionGetOK struct {
	Payload *models.ShelfResponse
}

// IsSuccess returns true when this shelf collection get o k response has a 2xx status code
func (o *ShelfCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this shelf collection get o k response has a 3xx status code
func (o *ShelfCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this shelf collection get o k response has a 4xx status code
func (o *ShelfCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this shelf collection get o k response has a 5xx status code
func (o *ShelfCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this shelf collection get o k response a status code equal to that given
func (o *ShelfCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the shelf collection get o k response
func (o *ShelfCollectionGetOK) Code() int {
	return 200
}

func (o *ShelfCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/shelves][%d] shelfCollectionGetOK %s", 200, payload)
}

func (o *ShelfCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/shelves][%d] shelfCollectionGetOK %s", 200, payload)
}

func (o *ShelfCollectionGetOK) GetPayload() *models.ShelfResponse {
	return o.Payload
}

func (o *ShelfCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShelfResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShelfCollectionGetDefault creates a ShelfCollectionGetDefault with default headers values
func NewShelfCollectionGetDefault(code int) *ShelfCollectionGetDefault {
	return &ShelfCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ShelfCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ShelfCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this shelf collection get default response has a 2xx status code
func (o *ShelfCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this shelf collection get default response has a 3xx status code
func (o *ShelfCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this shelf collection get default response has a 4xx status code
func (o *ShelfCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this shelf collection get default response has a 5xx status code
func (o *ShelfCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this shelf collection get default response a status code equal to that given
func (o *ShelfCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the shelf collection get default response
func (o *ShelfCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ShelfCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/shelves][%d] shelf_collection_get default %s", o._statusCode, payload)
}

func (o *ShelfCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/shelves][%d] shelf_collection_get default %s", o._statusCode, payload)
}

func (o *ShelfCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ShelfCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
