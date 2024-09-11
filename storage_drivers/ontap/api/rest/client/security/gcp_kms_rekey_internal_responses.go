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

// GcpKmsRekeyInternalReader is a Reader for the GcpKmsRekeyInternal structure.
type GcpKmsRekeyInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GcpKmsRekeyInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGcpKmsRekeyInternalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGcpKmsRekeyInternalAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGcpKmsRekeyInternalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGcpKmsRekeyInternalCreated creates a GcpKmsRekeyInternalCreated with default headers values
func NewGcpKmsRekeyInternalCreated() *GcpKmsRekeyInternalCreated {
	return &GcpKmsRekeyInternalCreated{}
}

/*
GcpKmsRekeyInternalCreated describes a response with status code 201, with default header values.

Created
*/
type GcpKmsRekeyInternalCreated struct {
	Payload *models.GcpKmsKeyJobLinkResponse
}

// IsSuccess returns true when this gcp kms rekey internal created response has a 2xx status code
func (o *GcpKmsRekeyInternalCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gcp kms rekey internal created response has a 3xx status code
func (o *GcpKmsRekeyInternalCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gcp kms rekey internal created response has a 4xx status code
func (o *GcpKmsRekeyInternalCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this gcp kms rekey internal created response has a 5xx status code
func (o *GcpKmsRekeyInternalCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this gcp kms rekey internal created response a status code equal to that given
func (o *GcpKmsRekeyInternalCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the gcp kms rekey internal created response
func (o *GcpKmsRekeyInternalCreated) Code() int {
	return 201
}

func (o *GcpKmsRekeyInternalCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{uuid}/rekey-internal][%d] gcpKmsRekeyInternalCreated %s", 201, payload)
}

func (o *GcpKmsRekeyInternalCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{uuid}/rekey-internal][%d] gcpKmsRekeyInternalCreated %s", 201, payload)
}

func (o *GcpKmsRekeyInternalCreated) GetPayload() *models.GcpKmsKeyJobLinkResponse {
	return o.Payload
}

func (o *GcpKmsRekeyInternalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GcpKmsKeyJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGcpKmsRekeyInternalAccepted creates a GcpKmsRekeyInternalAccepted with default headers values
func NewGcpKmsRekeyInternalAccepted() *GcpKmsRekeyInternalAccepted {
	return &GcpKmsRekeyInternalAccepted{}
}

/*
GcpKmsRekeyInternalAccepted describes a response with status code 202, with default header values.

Accepted
*/
type GcpKmsRekeyInternalAccepted struct {
	Payload *models.GcpKmsKeyJobLinkResponse
}

// IsSuccess returns true when this gcp kms rekey internal accepted response has a 2xx status code
func (o *GcpKmsRekeyInternalAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gcp kms rekey internal accepted response has a 3xx status code
func (o *GcpKmsRekeyInternalAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gcp kms rekey internal accepted response has a 4xx status code
func (o *GcpKmsRekeyInternalAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this gcp kms rekey internal accepted response has a 5xx status code
func (o *GcpKmsRekeyInternalAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this gcp kms rekey internal accepted response a status code equal to that given
func (o *GcpKmsRekeyInternalAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the gcp kms rekey internal accepted response
func (o *GcpKmsRekeyInternalAccepted) Code() int {
	return 202
}

func (o *GcpKmsRekeyInternalAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{uuid}/rekey-internal][%d] gcpKmsRekeyInternalAccepted %s", 202, payload)
}

func (o *GcpKmsRekeyInternalAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{uuid}/rekey-internal][%d] gcpKmsRekeyInternalAccepted %s", 202, payload)
}

func (o *GcpKmsRekeyInternalAccepted) GetPayload() *models.GcpKmsKeyJobLinkResponse {
	return o.Payload
}

func (o *GcpKmsRekeyInternalAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GcpKmsKeyJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGcpKmsRekeyInternalDefault creates a GcpKmsRekeyInternalDefault with default headers values
func NewGcpKmsRekeyInternalDefault(code int) *GcpKmsRekeyInternalDefault {
	return &GcpKmsRekeyInternalDefault{
		_statusCode: code,
	}
}

/*
	GcpKmsRekeyInternalDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537547 | One or more volume encryption keys for encrypted volumes of this data SVM are stored in the key manager configured for the admin SVM. Use the REST API POST method to migrate this data SVM's keys from the admin SVM's key manager to this data SVM's key manager before running the rekey operation. |
| 65537556 | ONTAP is not able to successfully encrypt or decrypt because the configured external key manager for this SVM is in a blocked state. Possible reasons for a blocked state include the top-level external key protection key not found, disabled or having insufficient privileges. |
| 65537559 | There are no existing internal keys for the SVM. A rekey operation is allowed for an SVM with one or more encryption keys. |
| 65537610 | Rekey cannot be performed on the SVM while the enabled keystore configuration is being switched. If a previous attempt to switch the keystore configuration failed, or was interrupted, the system will continue to prevent rekeying for the SVM. Use the REST API PATCH method "/api/security/key-stores/{uuid}" to re-run and complete the operation. |
| 65537721 | Google Cloud KMS is not configured for the given SVM. |
| 65539436 | Rekey cannot be performed on the SVM while the enabled keystore configuration is being initialized. Wait until the keystore is in the active state, and rerun the rekey operation. |
| 65539437 | Rekey cannot be performed on the SVM while the enabled keystore configuration is being disabled. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type GcpKmsRekeyInternalDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this gcp kms rekey internal default response has a 2xx status code
func (o *GcpKmsRekeyInternalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this gcp kms rekey internal default response has a 3xx status code
func (o *GcpKmsRekeyInternalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this gcp kms rekey internal default response has a 4xx status code
func (o *GcpKmsRekeyInternalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this gcp kms rekey internal default response has a 5xx status code
func (o *GcpKmsRekeyInternalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this gcp kms rekey internal default response a status code equal to that given
func (o *GcpKmsRekeyInternalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the gcp kms rekey internal default response
func (o *GcpKmsRekeyInternalDefault) Code() int {
	return o._statusCode
}

func (o *GcpKmsRekeyInternalDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{uuid}/rekey-internal][%d] gcp_kms_rekey_internal default %s", o._statusCode, payload)
}

func (o *GcpKmsRekeyInternalDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{uuid}/rekey-internal][%d] gcp_kms_rekey_internal default %s", o._statusCode, payload)
}

func (o *GcpKmsRekeyInternalDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GcpKmsRekeyInternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
