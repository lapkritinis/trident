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

// AggregateGetReader is a Reader for the AggregateGet structure.
type AggregateGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAggregateGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateGetOK creates a AggregateGetOK with default headers values
func NewAggregateGetOK() *AggregateGetOK {
	return &AggregateGetOK{}
}

/*
AggregateGetOK describes a response with status code 200, with default header values.

OK
*/
type AggregateGetOK struct {
	Payload *models.Aggregate
}

// IsSuccess returns true when this aggregate get o k response has a 2xx status code
func (o *AggregateGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate get o k response has a 3xx status code
func (o *AggregateGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate get o k response has a 4xx status code
func (o *AggregateGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate get o k response has a 5xx status code
func (o *AggregateGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate get o k response a status code equal to that given
func (o *AggregateGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate get o k response
func (o *AggregateGetOK) Code() int {
	return 200
}

func (o *AggregateGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}][%d] aggregateGetOK %s", 200, payload)
}

func (o *AggregateGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}][%d] aggregateGetOK %s", 200, payload)
}

func (o *AggregateGetOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *AggregateGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateGetDefault creates a AggregateGetDefault with default headers values
func NewAggregateGetDefault(code int) *AggregateGetDefault {
	return &AggregateGetDefault{
		_statusCode: code,
	}
}

/*
	AggregateGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 787092 | The target field cannot be specified for this operation. |
| 7209049 | Cannot perform the operation because the aggregate is currently expanding. |
| 8586225 | Unexpected error encountered when retrieving metrics and statistics for this aggregate. |
| 19726382 | Another provisioning operation is in progress on this cluster. Wait a few minutes, and try the operation again. |
| 19726390 | Unable to provide a recommendation to expand the aggregate. |
| 19726391 | Too many unassigned disks visible to the node that owns this aggregate. |
| 19726392 | Layout of this aggregate is not a supported configuration. |
| 19726393 | Failed to expand the aggregate. Aggregate expansion is not supported on this system. |
| 19726394 | Automatic aggregate expansion is not supported on systems with multiple data aggregates. |
| 19726395 | Automatic aggregate expansion is not supported when MetroCluster is not configured |
| 19726396 | Automatic aggregate expansion is not supported when the DR group is not in a normal state |
| 19726397 | Aggregates must contain disks with identical disk-types and disk-sizes. |
| 19726402 | Internal error. Unable to determine the MetroCluster configuration state. |
| 19726538 | Cannot perform the operation because the aggregate is not in a healthy state. |
| 19726541 | Cannot perform the operation because the specified aggregate is a root aggregate. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AggregateGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this aggregate get default response has a 2xx status code
func (o *AggregateGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aggregate get default response has a 3xx status code
func (o *AggregateGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aggregate get default response has a 4xx status code
func (o *AggregateGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aggregate get default response has a 5xx status code
func (o *AggregateGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aggregate get default response a status code equal to that given
func (o *AggregateGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the aggregate get default response
func (o *AggregateGetDefault) Code() int {
	return o._statusCode
}

func (o *AggregateGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}][%d] aggregate_get default %s", o._statusCode, payload)
}

func (o *AggregateGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}][%d] aggregate_get default %s", o._statusCode, payload)
}

func (o *AggregateGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AggregateGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
