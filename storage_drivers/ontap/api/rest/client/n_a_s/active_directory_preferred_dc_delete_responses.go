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

// ActiveDirectoryPreferredDcDeleteReader is a Reader for the ActiveDirectoryPreferredDcDelete structure.
type ActiveDirectoryPreferredDcDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActiveDirectoryPreferredDcDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActiveDirectoryPreferredDcDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActiveDirectoryPreferredDcDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActiveDirectoryPreferredDcDeleteOK creates a ActiveDirectoryPreferredDcDeleteOK with default headers values
func NewActiveDirectoryPreferredDcDeleteOK() *ActiveDirectoryPreferredDcDeleteOK {
	return &ActiveDirectoryPreferredDcDeleteOK{}
}

/*
ActiveDirectoryPreferredDcDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ActiveDirectoryPreferredDcDeleteOK struct {
}

// IsSuccess returns true when this active directory preferred dc delete o k response has a 2xx status code
func (o *ActiveDirectoryPreferredDcDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this active directory preferred dc delete o k response has a 3xx status code
func (o *ActiveDirectoryPreferredDcDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this active directory preferred dc delete o k response has a 4xx status code
func (o *ActiveDirectoryPreferredDcDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this active directory preferred dc delete o k response has a 5xx status code
func (o *ActiveDirectoryPreferredDcDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this active directory preferred dc delete o k response a status code equal to that given
func (o *ActiveDirectoryPreferredDcDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the active directory preferred dc delete o k response
func (o *ActiveDirectoryPreferredDcDeleteOK) Code() int {
	return 200
}

func (o *ActiveDirectoryPreferredDcDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/active-directory/{svm.uuid}/preferred-domain-controllers/{fqdn}/{server_ip}][%d] activeDirectoryPreferredDcDeleteOK", 200)
}

func (o *ActiveDirectoryPreferredDcDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/active-directory/{svm.uuid}/preferred-domain-controllers/{fqdn}/{server_ip}][%d] activeDirectoryPreferredDcDeleteOK", 200)
}

func (o *ActiveDirectoryPreferredDcDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActiveDirectoryPreferredDcDeleteDefault creates a ActiveDirectoryPreferredDcDeleteDefault with default headers values
func NewActiveDirectoryPreferredDcDeleteDefault(code int) *ActiveDirectoryPreferredDcDeleteDefault {
	return &ActiveDirectoryPreferredDcDeleteDefault{
		_statusCode: code,
	}
}

/*
	ActiveDirectoryPreferredDcDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655507 | Failed to remove preferred-dc. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ActiveDirectoryPreferredDcDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this active directory preferred dc delete default response has a 2xx status code
func (o *ActiveDirectoryPreferredDcDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this active directory preferred dc delete default response has a 3xx status code
func (o *ActiveDirectoryPreferredDcDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this active directory preferred dc delete default response has a 4xx status code
func (o *ActiveDirectoryPreferredDcDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this active directory preferred dc delete default response has a 5xx status code
func (o *ActiveDirectoryPreferredDcDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this active directory preferred dc delete default response a status code equal to that given
func (o *ActiveDirectoryPreferredDcDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the active directory preferred dc delete default response
func (o *ActiveDirectoryPreferredDcDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ActiveDirectoryPreferredDcDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/active-directory/{svm.uuid}/preferred-domain-controllers/{fqdn}/{server_ip}][%d] active_directory_preferred_dc_delete default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryPreferredDcDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/active-directory/{svm.uuid}/preferred-domain-controllers/{fqdn}/{server_ip}][%d] active_directory_preferred_dc_delete default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryPreferredDcDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActiveDirectoryPreferredDcDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
