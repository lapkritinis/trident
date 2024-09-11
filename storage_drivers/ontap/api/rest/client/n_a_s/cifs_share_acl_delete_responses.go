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

// CifsShareACLDeleteReader is a Reader for the CifsShareACLDelete structure.
type CifsShareACLDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareACLDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareACLDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareACLDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareACLDeleteOK creates a CifsShareACLDeleteOK with default headers values
func NewCifsShareACLDeleteOK() *CifsShareACLDeleteOK {
	return &CifsShareACLDeleteOK{}
}

/*
CifsShareACLDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareACLDeleteOK struct {
}

// IsSuccess returns true when this cifs share Acl delete o k response has a 2xx status code
func (o *CifsShareACLDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs share Acl delete o k response has a 3xx status code
func (o *CifsShareACLDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs share Acl delete o k response has a 4xx status code
func (o *CifsShareACLDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs share Acl delete o k response has a 5xx status code
func (o *CifsShareACLDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs share Acl delete o k response a status code equal to that given
func (o *CifsShareACLDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs share Acl delete o k response
func (o *CifsShareACLDeleteOK) Code() int {
	return 200
}

func (o *CifsShareACLDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type}][%d] cifsShareAclDeleteOK", 200)
}

func (o *CifsShareACLDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type}][%d] cifsShareAclDeleteOK", 200)
}

func (o *CifsShareACLDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsShareACLDeleteDefault creates a CifsShareACLDeleteDefault with default headers values
func NewCifsShareACLDeleteDefault(code int) *CifsShareACLDeleteDefault {
	return &CifsShareACLDeleteDefault{
		_statusCode: code,
	}
}

/*
CifsShareACLDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type CifsShareACLDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs share acl delete default response has a 2xx status code
func (o *CifsShareACLDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs share acl delete default response has a 3xx status code
func (o *CifsShareACLDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs share acl delete default response has a 4xx status code
func (o *CifsShareACLDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs share acl delete default response has a 5xx status code
func (o *CifsShareACLDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs share acl delete default response a status code equal to that given
func (o *CifsShareACLDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs share acl delete default response
func (o *CifsShareACLDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareACLDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type}][%d] cifs_share_acl_delete default %s", o._statusCode, payload)
}

func (o *CifsShareACLDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type}][%d] cifs_share_acl_delete default %s", o._statusCode, payload)
}

func (o *CifsShareACLDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareACLDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
