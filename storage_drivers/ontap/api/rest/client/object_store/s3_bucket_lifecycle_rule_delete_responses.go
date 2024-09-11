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

// S3BucketLifecycleRuleDeleteReader is a Reader for the S3BucketLifecycleRuleDelete structure.
type S3BucketLifecycleRuleDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketLifecycleRuleDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3BucketLifecycleRuleDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketLifecycleRuleDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketLifecycleRuleDeleteOK creates a S3BucketLifecycleRuleDeleteOK with default headers values
func NewS3BucketLifecycleRuleDeleteOK() *S3BucketLifecycleRuleDeleteOK {
	return &S3BucketLifecycleRuleDeleteOK{}
}

/*
S3BucketLifecycleRuleDeleteOK describes a response with status code 200, with default header values.

OK
*/
type S3BucketLifecycleRuleDeleteOK struct {
}

// IsSuccess returns true when this s3 bucket lifecycle rule delete o k response has a 2xx status code
func (o *S3BucketLifecycleRuleDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket lifecycle rule delete o k response has a 3xx status code
func (o *S3BucketLifecycleRuleDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket lifecycle rule delete o k response has a 4xx status code
func (o *S3BucketLifecycleRuleDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket lifecycle rule delete o k response has a 5xx status code
func (o *S3BucketLifecycleRuleDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket lifecycle rule delete o k response a status code equal to that given
func (o *S3BucketLifecycleRuleDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 bucket lifecycle rule delete o k response
func (o *S3BucketLifecycleRuleDeleteOK) Code() int {
	return 200
}

func (o *S3BucketLifecycleRuleDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3BucketLifecycleRuleDeleteOK", 200)
}

func (o *S3BucketLifecycleRuleDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3BucketLifecycleRuleDeleteOK", 200)
}

func (o *S3BucketLifecycleRuleDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS3BucketLifecycleRuleDeleteDefault creates a S3BucketLifecycleRuleDeleteDefault with default headers values
func NewS3BucketLifecycleRuleDeleteDefault(code int) *S3BucketLifecycleRuleDeleteDefault {
	return &S3BucketLifecycleRuleDeleteDefault{
		_statusCode: code,
	}
}

/*
	S3BucketLifecycleRuleDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error code | Message |
| ---------- | ------- |
| 92405861   | "The specified SVM UUID or bucket UUID does not exist.";
| 92406139   | "Lifecycle Management rule for bucket in Vserver with action is a stale entry. Contact technical support for assistance.";
*/
type S3BucketLifecycleRuleDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket lifecycle rule delete default response has a 2xx status code
func (o *S3BucketLifecycleRuleDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket lifecycle rule delete default response has a 3xx status code
func (o *S3BucketLifecycleRuleDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket lifecycle rule delete default response has a 4xx status code
func (o *S3BucketLifecycleRuleDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket lifecycle rule delete default response has a 5xx status code
func (o *S3BucketLifecycleRuleDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket lifecycle rule delete default response a status code equal to that given
func (o *S3BucketLifecycleRuleDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket lifecycle rule delete default response
func (o *S3BucketLifecycleRuleDeleteDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketLifecycleRuleDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3_bucket_lifecycle_rule_delete default %s", o._statusCode, payload)
}

func (o *S3BucketLifecycleRuleDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3_bucket_lifecycle_rule_delete default %s", o._statusCode, payload)
}

func (o *S3BucketLifecycleRuleDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketLifecycleRuleDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
