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

// DuoGetReader is a Reader for the DuoGet structure.
type DuoGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DuoGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDuoGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDuoGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDuoGetOK creates a DuoGetOK with default headers values
func NewDuoGetOK() *DuoGetOK {
	return &DuoGetOK{}
}

/*
DuoGetOK describes a response with status code 200, with default header values.

OK
*/
type DuoGetOK struct {
	Payload *models.Duo
}

// IsSuccess returns true when this duo get o k response has a 2xx status code
func (o *DuoGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this duo get o k response has a 3xx status code
func (o *DuoGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this duo get o k response has a 4xx status code
func (o *DuoGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this duo get o k response has a 5xx status code
func (o *DuoGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this duo get o k response a status code equal to that given
func (o *DuoGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the duo get o k response
func (o *DuoGetOK) Code() int {
	return 200
}

func (o *DuoGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/profiles/{owner.uuid}][%d] duoGetOK %s", 200, payload)
}

func (o *DuoGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/profiles/{owner.uuid}][%d] duoGetOK %s", 200, payload)
}

func (o *DuoGetOK) GetPayload() *models.Duo {
	return o.Payload
}

func (o *DuoGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Duo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDuoGetDefault creates a DuoGetDefault with default headers values
func NewDuoGetDefault(code int) *DuoGetDefault {
	return &DuoGetDefault{
		_statusCode: code,
	}
}

/*
DuoGetDefault describes a response with status code -1, with default header values.

Error
*/
type DuoGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this duo get default response has a 2xx status code
func (o *DuoGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this duo get default response has a 3xx status code
func (o *DuoGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this duo get default response has a 4xx status code
func (o *DuoGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this duo get default response has a 5xx status code
func (o *DuoGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this duo get default response a status code equal to that given
func (o *DuoGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the duo get default response
func (o *DuoGetDefault) Code() int {
	return o._statusCode
}

func (o *DuoGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/profiles/{owner.uuid}][%d] duo_get default %s", o._statusCode, payload)
}

func (o *DuoGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/profiles/{owner.uuid}][%d] duo_get default %s", o._statusCode, payload)
}

func (o *DuoGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DuoGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
