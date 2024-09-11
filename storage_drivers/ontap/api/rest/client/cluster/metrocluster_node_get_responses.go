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

// MetroclusterNodeGetReader is a Reader for the MetroclusterNodeGet structure.
type MetroclusterNodeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterNodeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroclusterNodeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterNodeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterNodeGetOK creates a MetroclusterNodeGetOK with default headers values
func NewMetroclusterNodeGetOK() *MetroclusterNodeGetOK {
	return &MetroclusterNodeGetOK{}
}

/*
MetroclusterNodeGetOK describes a response with status code 200, with default header values.

OK
*/
type MetroclusterNodeGetOK struct {
	Payload *models.MetroclusterNode
}

// IsSuccess returns true when this metrocluster node get o k response has a 2xx status code
func (o *MetroclusterNodeGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster node get o k response has a 3xx status code
func (o *MetroclusterNodeGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster node get o k response has a 4xx status code
func (o *MetroclusterNodeGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster node get o k response has a 5xx status code
func (o *MetroclusterNodeGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster node get o k response a status code equal to that given
func (o *MetroclusterNodeGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the metrocluster node get o k response
func (o *MetroclusterNodeGetOK) Code() int {
	return 200
}

func (o *MetroclusterNodeGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/nodes/{node.uuid}][%d] metroclusterNodeGetOK %s", 200, payload)
}

func (o *MetroclusterNodeGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/nodes/{node.uuid}][%d] metroclusterNodeGetOK %s", 200, payload)
}

func (o *MetroclusterNodeGetOK) GetPayload() *models.MetroclusterNode {
	return o.Payload
}

func (o *MetroclusterNodeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetroclusterNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterNodeGetDefault creates a MetroclusterNodeGetDefault with default headers values
func NewMetroclusterNodeGetDefault(code int) *MetroclusterNodeGetDefault {
	return &MetroclusterNodeGetDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterNodeGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MetroclusterNodeGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this metrocluster node get default response has a 2xx status code
func (o *MetroclusterNodeGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster node get default response has a 3xx status code
func (o *MetroclusterNodeGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster node get default response has a 4xx status code
func (o *MetroclusterNodeGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster node get default response has a 5xx status code
func (o *MetroclusterNodeGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster node get default response a status code equal to that given
func (o *MetroclusterNodeGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the metrocluster node get default response
func (o *MetroclusterNodeGetDefault) Code() int {
	return o._statusCode
}

func (o *MetroclusterNodeGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/nodes/{node.uuid}][%d] metrocluster_node_get default %s", o._statusCode, payload)
}

func (o *MetroclusterNodeGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/nodes/{node.uuid}][%d] metrocluster_node_get default %s", o._statusCode, payload)
}

func (o *MetroclusterNodeGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterNodeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
