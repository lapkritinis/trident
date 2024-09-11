// Code generated by go-swagger; DO NOT EDIT.

package security

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

// GcpKmsCollectionGetReader is a Reader for the GcpKmsCollectionGet structure.
type GcpKmsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GcpKmsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGcpKmsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGcpKmsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGcpKmsCollectionGetOK creates a GcpKmsCollectionGetOK with default headers values
func NewGcpKmsCollectionGetOK() *GcpKmsCollectionGetOK {
	return &GcpKmsCollectionGetOK{}
}

/*
GcpKmsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type GcpKmsCollectionGetOK struct {
	Payload *models.GcpKmsResponse
}

// IsSuccess returns true when this gcp kms collection get o k response has a 2xx status code
func (o *GcpKmsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gcp kms collection get o k response has a 3xx status code
func (o *GcpKmsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gcp kms collection get o k response has a 4xx status code
func (o *GcpKmsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this gcp kms collection get o k response has a 5xx status code
func (o *GcpKmsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this gcp kms collection get o k response a status code equal to that given
func (o *GcpKmsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the gcp kms collection get o k response
func (o *GcpKmsCollectionGetOK) Code() int {
	return 200
}

func (o *GcpKmsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/gcp-kms][%d] gcpKmsCollectionGetOK %s", 200, payload)
}

func (o *GcpKmsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/gcp-kms][%d] gcpKmsCollectionGetOK %s", 200, payload)
}

func (o *GcpKmsCollectionGetOK) GetPayload() *models.GcpKmsResponse {
	return o.Payload
}

func (o *GcpKmsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GcpKmsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGcpKmsCollectionGetDefault creates a GcpKmsCollectionGetDefault with default headers values
func NewGcpKmsCollectionGetDefault(code int) *GcpKmsCollectionGetDefault {
	return &GcpKmsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	GcpKmsCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537551 | Top-level internal key protection key (KEK) unavailable on one or more nodes. |
| 65537552 | Embedded KMIP server status not available. |
| 65537730 | The Google Cloud Key Management Service is unreachable from one or more nodes. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type GcpKmsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this gcp kms collection get default response has a 2xx status code
func (o *GcpKmsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this gcp kms collection get default response has a 3xx status code
func (o *GcpKmsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this gcp kms collection get default response has a 4xx status code
func (o *GcpKmsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this gcp kms collection get default response has a 5xx status code
func (o *GcpKmsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this gcp kms collection get default response a status code equal to that given
func (o *GcpKmsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the gcp kms collection get default response
func (o *GcpKmsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *GcpKmsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/gcp-kms][%d] gcp_kms_collection_get default %s", o._statusCode, payload)
}

func (o *GcpKmsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/gcp-kms][%d] gcp_kms_collection_get default %s", o._statusCode, payload)
}

func (o *GcpKmsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GcpKmsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
