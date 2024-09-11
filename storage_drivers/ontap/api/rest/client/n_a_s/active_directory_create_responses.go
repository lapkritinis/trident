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

// ActiveDirectoryCreateReader is a Reader for the ActiveDirectoryCreate structure.
type ActiveDirectoryCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActiveDirectoryCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewActiveDirectoryCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActiveDirectoryCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActiveDirectoryCreateCreated creates a ActiveDirectoryCreateCreated with default headers values
func NewActiveDirectoryCreateCreated() *ActiveDirectoryCreateCreated {
	return &ActiveDirectoryCreateCreated{}
}

/*
ActiveDirectoryCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ActiveDirectoryCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.ActiveDirectory
}

// IsSuccess returns true when this active directory create created response has a 2xx status code
func (o *ActiveDirectoryCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this active directory create created response has a 3xx status code
func (o *ActiveDirectoryCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this active directory create created response has a 4xx status code
func (o *ActiveDirectoryCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this active directory create created response has a 5xx status code
func (o *ActiveDirectoryCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this active directory create created response a status code equal to that given
func (o *ActiveDirectoryCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the active directory create created response
func (o *ActiveDirectoryCreateCreated) Code() int {
	return 201
}

func (o *ActiveDirectoryCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/active-directory][%d] activeDirectoryCreateCreated %s", 201, payload)
}

func (o *ActiveDirectoryCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/active-directory][%d] activeDirectoryCreateCreated %s", 201, payload)
}

func (o *ActiveDirectoryCreateCreated) GetPayload() *models.ActiveDirectory {
	return o.Payload
}

func (o *ActiveDirectoryCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.ActiveDirectory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActiveDirectoryCreateDefault creates a ActiveDirectoryCreateDefault with default headers values
func NewActiveDirectoryCreateDefault(code int) *ActiveDirectoryCreateDefault {
	return &ActiveDirectoryCreateDefault{
		_statusCode: code,
	}
}

/*
	ActiveDirectoryCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655562 | The NetBIOS name cannot be longer than 15 characters. |
| 655915 | A CIFS server for this SVM already exists. Having both a CIFS server and an Active Directory account for the same SVM is not supported. Use the \\\"vserver cifs delete\\\" command to delete the existing CIFS server and try the command again. |
| 656464 | Failed to create the Active Directory machine account. Reason: Invalid Credentials. |
| 656465 | Failed to create the Active Directory machine account. Reason: An account with this name already exists. |
| 656466 | Failed to create the Active Directory machine account. Reason: Unable to connect to any domain controllers. |
| 656467 | Failed to create the Active Directory machine account. Reason: Organizational-Unit not found. |
| 656490 | Unable to create the Active Directory account. The Active Directory account name is already used by another SVM. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ActiveDirectoryCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this active directory create default response has a 2xx status code
func (o *ActiveDirectoryCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this active directory create default response has a 3xx status code
func (o *ActiveDirectoryCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this active directory create default response has a 4xx status code
func (o *ActiveDirectoryCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this active directory create default response has a 5xx status code
func (o *ActiveDirectoryCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this active directory create default response a status code equal to that given
func (o *ActiveDirectoryCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the active directory create default response
func (o *ActiveDirectoryCreateDefault) Code() int {
	return o._statusCode
}

func (o *ActiveDirectoryCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/active-directory][%d] active_directory_create default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/active-directory][%d] active_directory_create default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActiveDirectoryCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
