// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// FpolicyPolicyCreateReader is a Reader for the FpolicyPolicyCreate structure.
type FpolicyPolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyPolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFpolicyPolicyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyPolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyPolicyCreateCreated creates a FpolicyPolicyCreateCreated with default headers values
func NewFpolicyPolicyCreateCreated() *FpolicyPolicyCreateCreated {
	return &FpolicyPolicyCreateCreated{}
}

/*
FpolicyPolicyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FpolicyPolicyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.FpolicyPolicyResponse
}

// IsSuccess returns true when this fpolicy policy create created response has a 2xx status code
func (o *FpolicyPolicyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy policy create created response has a 3xx status code
func (o *FpolicyPolicyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy policy create created response has a 4xx status code
func (o *FpolicyPolicyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy policy create created response has a 5xx status code
func (o *FpolicyPolicyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy policy create created response a status code equal to that given
func (o *FpolicyPolicyCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the fpolicy policy create created response
func (o *FpolicyPolicyCreateCreated) Code() int {
	return 201
}

func (o *FpolicyPolicyCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/policies][%d] fpolicyPolicyCreateCreated %s", 201, payload)
}

func (o *FpolicyPolicyCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/policies][%d] fpolicyPolicyCreateCreated %s", 201, payload)
}

func (o *FpolicyPolicyCreateCreated) GetPayload() *models.FpolicyPolicyResponse {
	return o.Payload
}

func (o *FpolicyPolicyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.FpolicyPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyPolicyCreateDefault creates a FpolicyPolicyCreateDefault with default headers values
func NewFpolicyPolicyCreateDefault(code int) *FpolicyPolicyCreateDefault {
	return &FpolicyPolicyCreateDefault{
		_statusCode: code,
	}
}

/*
	FpolicyPolicyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 9764875    | An FPolicy event does not exist |
| 9764888    | An FPolicy engine does not exist |
| 9764898    | An FPolicy policy cannot be created without defining its scope |
| 9765027    | FPolicy creation is successful but it cannot be enabled as the priority is already in use by another policy |
| 9765037    | FPolicy creation failed as passthrough-read cannot be enabled for policy without privileged user |
| 9765038    | Passthrough-read policies are not supported with asynchronous external engine |
| 9765056    | The specified Persistent Store does not exist |
| 9765059    | Persistent Store feature is not supported with native engine |
| 9765060    | Persistent Store feature is not supported with synchronous engine |
| 9765061    | Persistent Store feature is not supported with mandatory screening |
| 9765065    | A valid privileged user name must be in the form "domain-name\\user-name" |
| 9765066    | The privileged user contains characters that are not allowed |
*/
type FpolicyPolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy policy create default response has a 2xx status code
func (o *FpolicyPolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy policy create default response has a 3xx status code
func (o *FpolicyPolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy policy create default response has a 4xx status code
func (o *FpolicyPolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy policy create default response has a 5xx status code
func (o *FpolicyPolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy policy create default response a status code equal to that given
func (o *FpolicyPolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy policy create default response
func (o *FpolicyPolicyCreateDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyPolicyCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/policies][%d] fpolicy_policy_create default %s", o._statusCode, payload)
}

func (o *FpolicyPolicyCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/policies][%d] fpolicy_policy_create default %s", o._statusCode, payload)
}

func (o *FpolicyPolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyPolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
