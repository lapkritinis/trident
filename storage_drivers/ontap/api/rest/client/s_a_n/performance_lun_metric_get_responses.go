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

// PerformanceLunMetricGetReader is a Reader for the PerformanceLunMetricGet structure.
type PerformanceLunMetricGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceLunMetricGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceLunMetricGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformanceLunMetricGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformanceLunMetricGetOK creates a PerformanceLunMetricGetOK with default headers values
func NewPerformanceLunMetricGetOK() *PerformanceLunMetricGetOK {
	return &PerformanceLunMetricGetOK{}
}

/*
PerformanceLunMetricGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceLunMetricGetOK struct {
	Payload *models.PerformanceLunMetric
}

// IsSuccess returns true when this performance lun metric get o k response has a 2xx status code
func (o *PerformanceLunMetricGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance lun metric get o k response has a 3xx status code
func (o *PerformanceLunMetricGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance lun metric get o k response has a 4xx status code
func (o *PerformanceLunMetricGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance lun metric get o k response has a 5xx status code
func (o *PerformanceLunMetricGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance lun metric get o k response a status code equal to that given
func (o *PerformanceLunMetricGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance lun metric get o k response
func (o *PerformanceLunMetricGetOK) Code() int {
	return 200
}

func (o *PerformanceLunMetricGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/metrics/{timestamp}][%d] performanceLunMetricGetOK %s", 200, payload)
}

func (o *PerformanceLunMetricGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/metrics/{timestamp}][%d] performanceLunMetricGetOK %s", 200, payload)
}

func (o *PerformanceLunMetricGetOK) GetPayload() *models.PerformanceLunMetric {
	return o.Payload
}

func (o *PerformanceLunMetricGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceLunMetric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformanceLunMetricGetDefault creates a PerformanceLunMetricGetDefault with default headers values
func NewPerformanceLunMetricGetDefault(code int) *PerformanceLunMetricGetDefault {
	return &PerformanceLunMetricGetDefault{
		_statusCode: code,
	}
}

/*
	PerformanceLunMetricGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8585947 | No metrics are available for the requested object. |
| 8586225 | An unexpected error occurred retrieving metrics for the requested object. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type PerformanceLunMetricGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this performance lun metric get default response has a 2xx status code
func (o *PerformanceLunMetricGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance lun metric get default response has a 3xx status code
func (o *PerformanceLunMetricGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance lun metric get default response has a 4xx status code
func (o *PerformanceLunMetricGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance lun metric get default response has a 5xx status code
func (o *PerformanceLunMetricGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance lun metric get default response a status code equal to that given
func (o *PerformanceLunMetricGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance lun metric get default response
func (o *PerformanceLunMetricGetDefault) Code() int {
	return o._statusCode
}

func (o *PerformanceLunMetricGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/metrics/{timestamp}][%d] performance_lun_metric_get default %s", o._statusCode, payload)
}

func (o *PerformanceLunMetricGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/metrics/{timestamp}][%d] performance_lun_metric_get default %s", o._statusCode, payload)
}

func (o *PerformanceLunMetricGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PerformanceLunMetricGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
