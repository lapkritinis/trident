// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// SensorsCollectionGetReader is a Reader for the SensorsCollectionGet structure.
type SensorsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SensorsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSensorsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSensorsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSensorsCollectionGetOK creates a SensorsCollectionGetOK with default headers values
func NewSensorsCollectionGetOK() *SensorsCollectionGetOK {
	return &SensorsCollectionGetOK{}
}

/*
SensorsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SensorsCollectionGetOK struct {
	Payload *models.SensorsResponse
}

// IsSuccess returns true when this sensors collection get o k response has a 2xx status code
func (o *SensorsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sensors collection get o k response has a 3xx status code
func (o *SensorsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sensors collection get o k response has a 4xx status code
func (o *SensorsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sensors collection get o k response has a 5xx status code
func (o *SensorsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sensors collection get o k response a status code equal to that given
func (o *SensorsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sensors collection get o k response
func (o *SensorsCollectionGetOK) Code() int {
	return 200
}

func (o *SensorsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/sensors][%d] sensorsCollectionGetOK %s", 200, payload)
}

func (o *SensorsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/sensors][%d] sensorsCollectionGetOK %s", 200, payload)
}

func (o *SensorsCollectionGetOK) GetPayload() *models.SensorsResponse {
	return o.Payload
}

func (o *SensorsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SensorsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSensorsCollectionGetDefault creates a SensorsCollectionGetDefault with default headers values
func NewSensorsCollectionGetDefault(code int) *SensorsCollectionGetDefault {
	return &SensorsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SensorsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SensorsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this sensors collection get default response has a 2xx status code
func (o *SensorsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sensors collection get default response has a 3xx status code
func (o *SensorsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sensors collection get default response has a 4xx status code
func (o *SensorsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sensors collection get default response has a 5xx status code
func (o *SensorsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sensors collection get default response a status code equal to that given
func (o *SensorsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sensors collection get default response
func (o *SensorsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SensorsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/sensors][%d] sensors_collection_get default %s", o._statusCode, payload)
}

func (o *SensorsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/sensors][%d] sensors_collection_get default %s", o._statusCode, payload)
}

func (o *SensorsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SensorsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
