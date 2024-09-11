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

// CifsServiceGetReader is a Reader for the CifsServiceGet structure.
type CifsServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsServiceGetOK creates a CifsServiceGetOK with default headers values
func NewCifsServiceGetOK() *CifsServiceGetOK {
	return &CifsServiceGetOK{}
}

/*
CifsServiceGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsServiceGetOK struct {
	Payload *models.CifsService
}

// IsSuccess returns true when this cifs service get o k response has a 2xx status code
func (o *CifsServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs service get o k response has a 3xx status code
func (o *CifsServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs service get o k response has a 4xx status code
func (o *CifsServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs service get o k response has a 5xx status code
func (o *CifsServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs service get o k response a status code equal to that given
func (o *CifsServiceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs service get o k response
func (o *CifsServiceGetOK) Code() int {
	return 200
}

func (o *CifsServiceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services/{svm.uuid}][%d] cifsServiceGetOK %s", 200, payload)
}

func (o *CifsServiceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services/{svm.uuid}][%d] cifsServiceGetOK %s", 200, payload)
}

func (o *CifsServiceGetOK) GetPayload() *models.CifsService {
	return o.Payload
}

func (o *CifsServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsServiceGetDefault creates a CifsServiceGetDefault with default headers values
func NewCifsServiceGetDefault(code int) *CifsServiceGetDefault {
	return &CifsServiceGetDefault{
		_statusCode: code,
	}
}

/*
CifsServiceGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsServiceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs service get default response has a 2xx status code
func (o *CifsServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs service get default response has a 3xx status code
func (o *CifsServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs service get default response has a 4xx status code
func (o *CifsServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs service get default response has a 5xx status code
func (o *CifsServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs service get default response a status code equal to that given
func (o *CifsServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs service get default response
func (o *CifsServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsServiceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services/{svm.uuid}][%d] cifs_service_get default %s", o._statusCode, payload)
}

func (o *CifsServiceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services/{svm.uuid}][%d] cifs_service_get default %s", o._statusCode, payload)
}

func (o *CifsServiceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
