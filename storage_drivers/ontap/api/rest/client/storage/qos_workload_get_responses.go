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

// QosWorkloadGetReader is a Reader for the QosWorkloadGet structure.
type QosWorkloadGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QosWorkloadGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQosWorkloadGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQosWorkloadGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQosWorkloadGetOK creates a QosWorkloadGetOK with default headers values
func NewQosWorkloadGetOK() *QosWorkloadGetOK {
	return &QosWorkloadGetOK{}
}

/*
QosWorkloadGetOK describes a response with status code 200, with default header values.

OK
*/
type QosWorkloadGetOK struct {
	Payload *models.QosWorkload
}

// IsSuccess returns true when this qos workload get o k response has a 2xx status code
func (o *QosWorkloadGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qos workload get o k response has a 3xx status code
func (o *QosWorkloadGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qos workload get o k response has a 4xx status code
func (o *QosWorkloadGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this qos workload get o k response has a 5xx status code
func (o *QosWorkloadGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this qos workload get o k response a status code equal to that given
func (o *QosWorkloadGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the qos workload get o k response
func (o *QosWorkloadGetOK) Code() int {
	return 200
}

func (o *QosWorkloadGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qos/workloads/{uuid}][%d] qosWorkloadGetOK %s", 200, payload)
}

func (o *QosWorkloadGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qos/workloads/{uuid}][%d] qosWorkloadGetOK %s", 200, payload)
}

func (o *QosWorkloadGetOK) GetPayload() *models.QosWorkload {
	return o.Payload
}

func (o *QosWorkloadGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QosWorkload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosWorkloadGetDefault creates a QosWorkloadGetDefault with default headers values
func NewQosWorkloadGetDefault(code int) *QosWorkloadGetDefault {
	return &QosWorkloadGetDefault{
		_statusCode: code,
	}
}

/*
QosWorkloadGetDefault describes a response with status code -1, with default header values.

Error
*/
type QosWorkloadGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this qos workload get default response has a 2xx status code
func (o *QosWorkloadGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qos workload get default response has a 3xx status code
func (o *QosWorkloadGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qos workload get default response has a 4xx status code
func (o *QosWorkloadGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qos workload get default response has a 5xx status code
func (o *QosWorkloadGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qos workload get default response a status code equal to that given
func (o *QosWorkloadGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the qos workload get default response
func (o *QosWorkloadGetDefault) Code() int {
	return o._statusCode
}

func (o *QosWorkloadGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qos/workloads/{uuid}][%d] qos_workload_get default %s", o._statusCode, payload)
}

func (o *QosWorkloadGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qos/workloads/{uuid}][%d] qos_workload_get default %s", o._statusCode, payload)
}

func (o *QosWorkloadGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QosWorkloadGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
