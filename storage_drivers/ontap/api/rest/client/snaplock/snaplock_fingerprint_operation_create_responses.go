// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// SnaplockFingerprintOperationCreateReader is a Reader for the SnaplockFingerprintOperationCreate structure.
type SnaplockFingerprintOperationCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockFingerprintOperationCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnaplockFingerprintOperationCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockFingerprintOperationCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockFingerprintOperationCreateCreated creates a SnaplockFingerprintOperationCreateCreated with default headers values
func NewSnaplockFingerprintOperationCreateCreated() *SnaplockFingerprintOperationCreateCreated {
	return &SnaplockFingerprintOperationCreateCreated{}
}

/*
SnaplockFingerprintOperationCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnaplockFingerprintOperationCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this snaplock fingerprint operation create created response has a 2xx status code
func (o *SnaplockFingerprintOperationCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock fingerprint operation create created response has a 3xx status code
func (o *SnaplockFingerprintOperationCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock fingerprint operation create created response has a 4xx status code
func (o *SnaplockFingerprintOperationCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock fingerprint operation create created response has a 5xx status code
func (o *SnaplockFingerprintOperationCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock fingerprint operation create created response a status code equal to that given
func (o *SnaplockFingerprintOperationCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the snaplock fingerprint operation create created response
func (o *SnaplockFingerprintOperationCreateCreated) Code() int {
	return 201
}

func (o *SnaplockFingerprintOperationCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/file-fingerprints][%d] snaplockFingerprintOperationCreateCreated", 201)
}

func (o *SnaplockFingerprintOperationCreateCreated) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/file-fingerprints][%d] snaplockFingerprintOperationCreateCreated", 201)
}

func (o *SnaplockFingerprintOperationCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewSnaplockFingerprintOperationCreateDefault creates a SnaplockFingerprintOperationCreateDefault with default headers values
func NewSnaplockFingerprintOperationCreateDefault(code int) *SnaplockFingerprintOperationCreateDefault {
	return &SnaplockFingerprintOperationCreateDefault{
		_statusCode: code,
	}
}

/*
	SnaplockFingerprintOperationCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
|        2    | An invalid value was entered for one of the fields |
|  2621519    | Invalid SVM name format |
| 14090347    | Invalid path format  |
| 14090443    | Invalid volume name  |
| 14090444    | Invalid Vserver name |
| 14090447    | Invalid volume UUID  |
| 14090448    | Invalid key values. Provide valid Vserver name and volume name or Vserver UUID and volume UUID  |
*/
type SnaplockFingerprintOperationCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock fingerprint operation create default response has a 2xx status code
func (o *SnaplockFingerprintOperationCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock fingerprint operation create default response has a 3xx status code
func (o *SnaplockFingerprintOperationCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock fingerprint operation create default response has a 4xx status code
func (o *SnaplockFingerprintOperationCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock fingerprint operation create default response has a 5xx status code
func (o *SnaplockFingerprintOperationCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock fingerprint operation create default response a status code equal to that given
func (o *SnaplockFingerprintOperationCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock fingerprint operation create default response
func (o *SnaplockFingerprintOperationCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockFingerprintOperationCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/file-fingerprints][%d] snaplock_fingerprint_operation_create default %s", o._statusCode, payload)
}

func (o *SnaplockFingerprintOperationCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/file-fingerprints][%d] snaplock_fingerprint_operation_create default %s", o._statusCode, payload)
}

func (o *SnaplockFingerprintOperationCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockFingerprintOperationCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
