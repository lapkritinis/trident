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

// CifsOpenFileDeleteReader is a Reader for the CifsOpenFileDelete structure.
type CifsOpenFileDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsOpenFileDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsOpenFileDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsOpenFileDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsOpenFileDeleteOK creates a CifsOpenFileDeleteOK with default headers values
func NewCifsOpenFileDeleteOK() *CifsOpenFileDeleteOK {
	return &CifsOpenFileDeleteOK{}
}

/*
CifsOpenFileDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsOpenFileDeleteOK struct {
}

// IsSuccess returns true when this cifs open file delete o k response has a 2xx status code
func (o *CifsOpenFileDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs open file delete o k response has a 3xx status code
func (o *CifsOpenFileDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs open file delete o k response has a 4xx status code
func (o *CifsOpenFileDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs open file delete o k response has a 5xx status code
func (o *CifsOpenFileDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs open file delete o k response a status code equal to that given
func (o *CifsOpenFileDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs open file delete o k response
func (o *CifsOpenFileDeleteOK) Code() int {
	return 200
}

func (o *CifsOpenFileDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/session/files/{node.uuid}/{svm.uuid}/{identifier}/{connection.identifier}/{session.identifier}][%d] cifsOpenFileDeleteOK", 200)
}

func (o *CifsOpenFileDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/session/files/{node.uuid}/{svm.uuid}/{identifier}/{connection.identifier}/{session.identifier}][%d] cifsOpenFileDeleteOK", 200)
}

func (o *CifsOpenFileDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsOpenFileDeleteDefault creates a CifsOpenFileDeleteDefault with default headers values
func NewCifsOpenFileDeleteDefault(code int) *CifsOpenFileDeleteDefault {
	return &CifsOpenFileDeleteDefault{
		_statusCode: code,
	}
}

/*
CifsOpenFileDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type CifsOpenFileDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs open file delete default response has a 2xx status code
func (o *CifsOpenFileDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs open file delete default response has a 3xx status code
func (o *CifsOpenFileDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs open file delete default response has a 4xx status code
func (o *CifsOpenFileDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs open file delete default response has a 5xx status code
func (o *CifsOpenFileDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs open file delete default response a status code equal to that given
func (o *CifsOpenFileDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs open file delete default response
func (o *CifsOpenFileDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsOpenFileDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/session/files/{node.uuid}/{svm.uuid}/{identifier}/{connection.identifier}/{session.identifier}][%d] cifs_open_file_delete default %s", o._statusCode, payload)
}

func (o *CifsOpenFileDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/session/files/{node.uuid}/{svm.uuid}/{identifier}/{connection.identifier}/{session.identifier}][%d] cifs_open_file_delete default %s", o._statusCode, payload)
}

func (o *CifsOpenFileDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsOpenFileDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
