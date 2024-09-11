// Code generated by go-swagger; DO NOT EDIT.

package object_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// S3BucketDeleteCollectionReader is a Reader for the S3BucketDeleteCollection structure.
type S3BucketDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3BucketDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewS3BucketDeleteCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketDeleteCollectionOK creates a S3BucketDeleteCollectionOK with default headers values
func NewS3BucketDeleteCollectionOK() *S3BucketDeleteCollectionOK {
	return &S3BucketDeleteCollectionOK{}
}

/*
S3BucketDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type S3BucketDeleteCollectionOK struct {
	Payload *models.S3BucketJobLinkResponse
}

// IsSuccess returns true when this s3 bucket delete collection o k response has a 2xx status code
func (o *S3BucketDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket delete collection o k response has a 3xx status code
func (o *S3BucketDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket delete collection o k response has a 4xx status code
func (o *S3BucketDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket delete collection o k response has a 5xx status code
func (o *S3BucketDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket delete collection o k response a status code equal to that given
func (o *S3BucketDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 bucket delete collection o k response
func (o *S3BucketDeleteCollectionOK) Code() int {
	return 200
}

func (o *S3BucketDeleteCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/s3/buckets][%d] s3BucketDeleteCollectionOK %s", 200, payload)
}

func (o *S3BucketDeleteCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/s3/buckets][%d] s3BucketDeleteCollectionOK %s", 200, payload)
}

func (o *S3BucketDeleteCollectionOK) GetPayload() *models.S3BucketJobLinkResponse {
	return o.Payload
}

func (o *S3BucketDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3BucketJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketDeleteCollectionAccepted creates a S3BucketDeleteCollectionAccepted with default headers values
func NewS3BucketDeleteCollectionAccepted() *S3BucketDeleteCollectionAccepted {
	return &S3BucketDeleteCollectionAccepted{}
}

/*
S3BucketDeleteCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type S3BucketDeleteCollectionAccepted struct {
	Payload *models.S3BucketJobLinkResponse
}

// IsSuccess returns true when this s3 bucket delete collection accepted response has a 2xx status code
func (o *S3BucketDeleteCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket delete collection accepted response has a 3xx status code
func (o *S3BucketDeleteCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket delete collection accepted response has a 4xx status code
func (o *S3BucketDeleteCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket delete collection accepted response has a 5xx status code
func (o *S3BucketDeleteCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket delete collection accepted response a status code equal to that given
func (o *S3BucketDeleteCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the s3 bucket delete collection accepted response
func (o *S3BucketDeleteCollectionAccepted) Code() int {
	return 202
}

func (o *S3BucketDeleteCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/s3/buckets][%d] s3BucketDeleteCollectionAccepted %s", 202, payload)
}

func (o *S3BucketDeleteCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/s3/buckets][%d] s3BucketDeleteCollectionAccepted %s", 202, payload)
}

func (o *S3BucketDeleteCollectionAccepted) GetPayload() *models.S3BucketJobLinkResponse {
	return o.Payload
}

func (o *S3BucketDeleteCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3BucketJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketDeleteCollectionDefault creates a S3BucketDeleteCollectionDefault with default headers values
func NewS3BucketDeleteCollectionDefault(code int) *S3BucketDeleteCollectionDefault {
	return &S3BucketDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	S3BucketDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error code | Message |
| ---------- | ------- |
| 92405811   | "Failed to delete bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Wait a few minutes and try the operation again.";
| 92405858   | "Failed to \\\"delete\\\" the \\\"bucket\\\" because the operation is only supported on data SVMs.";
| 92405861   | "The specified SVM UUID or bucket UUID does not exist.";
| 92405779   | "Failed to remove bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Reason: {Reason for failure}. ";
| 92405813   | "Failed to delete the object store volume. Reason: {Reason for failure}.";
| 92405864   | "An error occurred when deleting an access policy. The reason for failure is detailed in the error message.";
*/
type S3BucketDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket delete collection default response has a 2xx status code
func (o *S3BucketDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket delete collection default response has a 3xx status code
func (o *S3BucketDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket delete collection default response has a 4xx status code
func (o *S3BucketDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket delete collection default response has a 5xx status code
func (o *S3BucketDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket delete collection default response a status code equal to that given
func (o *S3BucketDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket delete collection default response
func (o *S3BucketDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/s3/buckets][%d] s3_bucket_delete_collection default %s", o._statusCode, payload)
}

func (o *S3BucketDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/s3/buckets][%d] s3_bucket_delete_collection default %s", o._statusCode, payload)
}

func (o *S3BucketDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
S3BucketDeleteCollectionBody s3 bucket delete collection body
swagger:model S3BucketDeleteCollectionBody
*/
type S3BucketDeleteCollectionBody struct {

	// s3 bucket response inline records
	S3BucketResponseInlineRecords []*models.S3Bucket `json:"records,omitempty"`
}

// Validate validates this s3 bucket delete collection body
func (o *S3BucketDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateS3BucketResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *S3BucketDeleteCollectionBody) validateS3BucketResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.S3BucketResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.S3BucketResponseInlineRecords); i++ {
		if swag.IsZero(o.S3BucketResponseInlineRecords[i]) { // not required
			continue
		}

		if o.S3BucketResponseInlineRecords[i] != nil {
			if err := o.S3BucketResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this s3 bucket delete collection body based on the context it is used
func (o *S3BucketDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateS3BucketResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *S3BucketDeleteCollectionBody) contextValidateS3BucketResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.S3BucketResponseInlineRecords); i++ {

		if o.S3BucketResponseInlineRecords[i] != nil {
			if err := o.S3BucketResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *S3BucketDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3BucketDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res S3BucketDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
