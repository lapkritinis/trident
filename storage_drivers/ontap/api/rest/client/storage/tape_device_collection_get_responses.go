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

// TapeDeviceCollectionGetReader is a Reader for the TapeDeviceCollectionGet structure.
type TapeDeviceCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TapeDeviceCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTapeDeviceCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTapeDeviceCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTapeDeviceCollectionGetOK creates a TapeDeviceCollectionGetOK with default headers values
func NewTapeDeviceCollectionGetOK() *TapeDeviceCollectionGetOK {
	return &TapeDeviceCollectionGetOK{}
}

/*
TapeDeviceCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type TapeDeviceCollectionGetOK struct {
	Payload *models.TapeDeviceResponse
}

// IsSuccess returns true when this tape device collection get o k response has a 2xx status code
func (o *TapeDeviceCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tape device collection get o k response has a 3xx status code
func (o *TapeDeviceCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tape device collection get o k response has a 4xx status code
func (o *TapeDeviceCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tape device collection get o k response has a 5xx status code
func (o *TapeDeviceCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tape device collection get o k response a status code equal to that given
func (o *TapeDeviceCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tape device collection get o k response
func (o *TapeDeviceCollectionGetOK) Code() int {
	return 200
}

func (o *TapeDeviceCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/tape-devices][%d] tapeDeviceCollectionGetOK %s", 200, payload)
}

func (o *TapeDeviceCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/tape-devices][%d] tapeDeviceCollectionGetOK %s", 200, payload)
}

func (o *TapeDeviceCollectionGetOK) GetPayload() *models.TapeDeviceResponse {
	return o.Payload
}

func (o *TapeDeviceCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TapeDeviceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTapeDeviceCollectionGetDefault creates a TapeDeviceCollectionGetDefault with default headers values
func NewTapeDeviceCollectionGetDefault(code int) *TapeDeviceCollectionGetDefault {
	return &TapeDeviceCollectionGetDefault{
		_statusCode: code,
	}
}

/*
TapeDeviceCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type TapeDeviceCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tape device collection get default response has a 2xx status code
func (o *TapeDeviceCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tape device collection get default response has a 3xx status code
func (o *TapeDeviceCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tape device collection get default response has a 4xx status code
func (o *TapeDeviceCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tape device collection get default response has a 5xx status code
func (o *TapeDeviceCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tape device collection get default response a status code equal to that given
func (o *TapeDeviceCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tape device collection get default response
func (o *TapeDeviceCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *TapeDeviceCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/tape-devices][%d] tape_device_collection_get default %s", o._statusCode, payload)
}

func (o *TapeDeviceCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/tape-devices][%d] tape_device_collection_get default %s", o._statusCode, payload)
}

func (o *TapeDeviceCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TapeDeviceCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
