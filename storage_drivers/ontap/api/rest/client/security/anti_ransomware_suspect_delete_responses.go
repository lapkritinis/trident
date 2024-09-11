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

// AntiRansomwareSuspectDeleteReader is a Reader for the AntiRansomwareSuspectDelete structure.
type AntiRansomwareSuspectDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AntiRansomwareSuspectDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAntiRansomwareSuspectDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAntiRansomwareSuspectDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /security/anti-ransomware/suspects/{volume.uuid}] anti_ransomware_suspect_delete", response, response.Code())
	}
}

// NewAntiRansomwareSuspectDeleteOK creates a AntiRansomwareSuspectDeleteOK with default headers values
func NewAntiRansomwareSuspectDeleteOK() *AntiRansomwareSuspectDeleteOK {
	return &AntiRansomwareSuspectDeleteOK{}
}

/*
AntiRansomwareSuspectDeleteOK describes a response with status code 200, with default header values.

OK
*/
type AntiRansomwareSuspectDeleteOK struct {
	Payload *models.AntiRansomwareSuspectJobLinkResponse
}

// IsSuccess returns true when this anti ransomware suspect delete o k response has a 2xx status code
func (o *AntiRansomwareSuspectDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this anti ransomware suspect delete o k response has a 3xx status code
func (o *AntiRansomwareSuspectDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this anti ransomware suspect delete o k response has a 4xx status code
func (o *AntiRansomwareSuspectDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this anti ransomware suspect delete o k response has a 5xx status code
func (o *AntiRansomwareSuspectDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this anti ransomware suspect delete o k response a status code equal to that given
func (o *AntiRansomwareSuspectDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the anti ransomware suspect delete o k response
func (o *AntiRansomwareSuspectDeleteOK) Code() int {
	return 200
}

func (o *AntiRansomwareSuspectDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects/{volume.uuid}][%d] antiRansomwareSuspectDeleteOK %s", 200, payload)
}

func (o *AntiRansomwareSuspectDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects/{volume.uuid}][%d] antiRansomwareSuspectDeleteOK %s", 200, payload)
}

func (o *AntiRansomwareSuspectDeleteOK) GetPayload() *models.AntiRansomwareSuspectJobLinkResponse {
	return o.Payload
}

func (o *AntiRansomwareSuspectDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AntiRansomwareSuspectJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAntiRansomwareSuspectDeleteAccepted creates a AntiRansomwareSuspectDeleteAccepted with default headers values
func NewAntiRansomwareSuspectDeleteAccepted() *AntiRansomwareSuspectDeleteAccepted {
	return &AntiRansomwareSuspectDeleteAccepted{}
}

/*
AntiRansomwareSuspectDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AntiRansomwareSuspectDeleteAccepted struct {
	Payload *models.AntiRansomwareSuspectJobLinkResponse
}

// IsSuccess returns true when this anti ransomware suspect delete accepted response has a 2xx status code
func (o *AntiRansomwareSuspectDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this anti ransomware suspect delete accepted response has a 3xx status code
func (o *AntiRansomwareSuspectDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this anti ransomware suspect delete accepted response has a 4xx status code
func (o *AntiRansomwareSuspectDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this anti ransomware suspect delete accepted response has a 5xx status code
func (o *AntiRansomwareSuspectDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this anti ransomware suspect delete accepted response a status code equal to that given
func (o *AntiRansomwareSuspectDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the anti ransomware suspect delete accepted response
func (o *AntiRansomwareSuspectDeleteAccepted) Code() int {
	return 202
}

func (o *AntiRansomwareSuspectDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects/{volume.uuid}][%d] antiRansomwareSuspectDeleteAccepted %s", 202, payload)
}

func (o *AntiRansomwareSuspectDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects/{volume.uuid}][%d] antiRansomwareSuspectDeleteAccepted %s", 202, payload)
}

func (o *AntiRansomwareSuspectDeleteAccepted) GetPayload() *models.AntiRansomwareSuspectJobLinkResponse {
	return o.Payload
}

func (o *AntiRansomwareSuspectDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AntiRansomwareSuspectJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
