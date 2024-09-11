// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

// S3GroupGetReader is a Reader for the S3GroupGet structure.
type S3GroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3GroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3GroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3GroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3GroupGetOK creates a S3GroupGetOK with default headers values
func NewS3GroupGetOK() *S3GroupGetOK {
	return &S3GroupGetOK{}
}

/*
S3GroupGetOK describes a response with status code 200, with default header values.

OK
*/
type S3GroupGetOK struct {
	Payload *models.S3Group
}

// IsSuccess returns true when this s3 group get o k response has a 2xx status code
func (o *S3GroupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 group get o k response has a 3xx status code
func (o *S3GroupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 group get o k response has a 4xx status code
func (o *S3GroupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 group get o k response has a 5xx status code
func (o *S3GroupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 group get o k response a status code equal to that given
func (o *S3GroupGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 group get o k response
func (o *S3GroupGetOK) Code() int {
	return 200
}

func (o *S3GroupGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups/{id}][%d] s3GroupGetOK %s", 200, payload)
}

func (o *S3GroupGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups/{id}][%d] s3GroupGetOK %s", 200, payload)
}

func (o *S3GroupGetOK) GetPayload() *models.S3Group {
	return o.Payload
}

func (o *S3GroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3GroupGetDefault creates a S3GroupGetDefault with default headers values
func NewS3GroupGetDefault(code int) *S3GroupGetDefault {
	return &S3GroupGetDefault{
		_statusCode: code,
	}
}

/*
S3GroupGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3GroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 group get default response has a 2xx status code
func (o *S3GroupGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 group get default response has a 3xx status code
func (o *S3GroupGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 group get default response has a 4xx status code
func (o *S3GroupGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 group get default response has a 5xx status code
func (o *S3GroupGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 group get default response a status code equal to that given
func (o *S3GroupGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 group get default response
func (o *S3GroupGetDefault) Code() int {
	return o._statusCode
}

func (o *S3GroupGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups/{id}][%d] s3_group_get default %s", o._statusCode, payload)
}

func (o *S3GroupGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups/{id}][%d] s3_group_get default %s", o._statusCode, payload)
}

func (o *S3GroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3GroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
