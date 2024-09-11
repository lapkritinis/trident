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

// SplitLoadCollectionGetReader is a Reader for the SplitLoadCollectionGet structure.
type SplitLoadCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SplitLoadCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSplitLoadCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSplitLoadCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSplitLoadCollectionGetOK creates a SplitLoadCollectionGetOK with default headers values
func NewSplitLoadCollectionGetOK() *SplitLoadCollectionGetOK {
	return &SplitLoadCollectionGetOK{}
}

/*
SplitLoadCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SplitLoadCollectionGetOK struct {
	Payload *models.SplitLoadResponse
}

// IsSuccess returns true when this split load collection get o k response has a 2xx status code
func (o *SplitLoadCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this split load collection get o k response has a 3xx status code
func (o *SplitLoadCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this split load collection get o k response has a 4xx status code
func (o *SplitLoadCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this split load collection get o k response has a 5xx status code
func (o *SplitLoadCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this split load collection get o k response a status code equal to that given
func (o *SplitLoadCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the split load collection get o k response
func (o *SplitLoadCollectionGetOK) Code() int {
	return 200
}

func (o *SplitLoadCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-loads][%d] splitLoadCollectionGetOK %s", 200, payload)
}

func (o *SplitLoadCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-loads][%d] splitLoadCollectionGetOK %s", 200, payload)
}

func (o *SplitLoadCollectionGetOK) GetPayload() *models.SplitLoadResponse {
	return o.Payload
}

func (o *SplitLoadCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SplitLoadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSplitLoadCollectionGetDefault creates a SplitLoadCollectionGetDefault with default headers values
func NewSplitLoadCollectionGetDefault(code int) *SplitLoadCollectionGetDefault {
	return &SplitLoadCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SplitLoadCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SplitLoadCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this split load collection get default response has a 2xx status code
func (o *SplitLoadCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this split load collection get default response has a 3xx status code
func (o *SplitLoadCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this split load collection get default response has a 4xx status code
func (o *SplitLoadCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this split load collection get default response has a 5xx status code
func (o *SplitLoadCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this split load collection get default response a status code equal to that given
func (o *SplitLoadCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the split load collection get default response
func (o *SplitLoadCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SplitLoadCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-loads][%d] split_load_collection_get default %s", o._statusCode, payload)
}

func (o *SplitLoadCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-loads][%d] split_load_collection_get default %s", o._statusCode, payload)
}

func (o *SplitLoadCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SplitLoadCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
