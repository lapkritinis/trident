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

// ShelfGetReader is a Reader for the ShelfGet structure.
type ShelfGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShelfGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShelfGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewShelfGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShelfGetOK creates a ShelfGetOK with default headers values
func NewShelfGetOK() *ShelfGetOK {
	return &ShelfGetOK{}
}

/*
ShelfGetOK describes a response with status code 200, with default header values.

OK
*/
type ShelfGetOK struct {
	Payload *models.Shelf
}

// IsSuccess returns true when this shelf get o k response has a 2xx status code
func (o *ShelfGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this shelf get o k response has a 3xx status code
func (o *ShelfGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this shelf get o k response has a 4xx status code
func (o *ShelfGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this shelf get o k response has a 5xx status code
func (o *ShelfGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this shelf get o k response a status code equal to that given
func (o *ShelfGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the shelf get o k response
func (o *ShelfGetOK) Code() int {
	return 200
}

func (o *ShelfGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/shelves/{uid}][%d] shelfGetOK %s", 200, payload)
}

func (o *ShelfGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/shelves/{uid}][%d] shelfGetOK %s", 200, payload)
}

func (o *ShelfGetOK) GetPayload() *models.Shelf {
	return o.Payload
}

func (o *ShelfGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Shelf)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShelfGetDefault creates a ShelfGetDefault with default headers values
func NewShelfGetDefault(code int) *ShelfGetDefault {
	return &ShelfGetDefault{
		_statusCode: code,
	}
}

/*
ShelfGetDefault describes a response with status code -1, with default header values.

Error
*/
type ShelfGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this shelf get default response has a 2xx status code
func (o *ShelfGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this shelf get default response has a 3xx status code
func (o *ShelfGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this shelf get default response has a 4xx status code
func (o *ShelfGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this shelf get default response has a 5xx status code
func (o *ShelfGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this shelf get default response a status code equal to that given
func (o *ShelfGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the shelf get default response
func (o *ShelfGetDefault) Code() int {
	return o._statusCode
}

func (o *ShelfGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/shelves/{uid}][%d] shelf_get default %s", o._statusCode, payload)
}

func (o *ShelfGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/shelves/{uid}][%d] shelf_get default %s", o._statusCode, payload)
}

func (o *ShelfGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ShelfGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
