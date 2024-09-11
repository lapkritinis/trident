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

// CifsShareDeleteReader is a Reader for the CifsShareDelete structure.
type CifsShareDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareDeleteOK creates a CifsShareDeleteOK with default headers values
func NewCifsShareDeleteOK() *CifsShareDeleteOK {
	return &CifsShareDeleteOK{}
}

/*
CifsShareDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareDeleteOK struct {
}

// IsSuccess returns true when this cifs share delete o k response has a 2xx status code
func (o *CifsShareDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs share delete o k response has a 3xx status code
func (o *CifsShareDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs share delete o k response has a 4xx status code
func (o *CifsShareDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs share delete o k response has a 5xx status code
func (o *CifsShareDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs share delete o k response a status code equal to that given
func (o *CifsShareDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs share delete o k response
func (o *CifsShareDeleteOK) Code() int {
	return 200
}

func (o *CifsShareDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifsShareDeleteOK", 200)
}

func (o *CifsShareDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifsShareDeleteOK", 200)
}

func (o *CifsShareDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsShareDeleteDefault creates a CifsShareDeleteDefault with default headers values
func NewCifsShareDeleteDefault(code int) *CifsShareDeleteDefault {
	return &CifsShareDeleteDefault{
		_statusCode: code,
	}
}

/*
	CifsShareDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655393     | Standard admin shares cannot be removed |
*/
type CifsShareDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs share delete default response has a 2xx status code
func (o *CifsShareDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs share delete default response has a 3xx status code
func (o *CifsShareDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs share delete default response has a 4xx status code
func (o *CifsShareDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs share delete default response has a 5xx status code
func (o *CifsShareDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs share delete default response a status code equal to that given
func (o *CifsShareDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs share delete default response
func (o *CifsShareDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifs_share_delete default %s", o._statusCode, payload)
}

func (o *CifsShareDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifs_share_delete default %s", o._statusCode, payload)
}

func (o *CifsShareDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
