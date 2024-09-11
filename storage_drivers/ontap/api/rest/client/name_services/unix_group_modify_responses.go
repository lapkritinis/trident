// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// UnixGroupModifyReader is a Reader for the UnixGroupModify structure.
type UnixGroupModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixGroupModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupModifyOK creates a UnixGroupModifyOK with default headers values
func NewUnixGroupModifyOK() *UnixGroupModifyOK {
	return &UnixGroupModifyOK{}
}

/*
UnixGroupModifyOK describes a response with status code 200, with default header values.

OK
*/
type UnixGroupModifyOK struct {
}

// IsSuccess returns true when this unix group modify o k response has a 2xx status code
func (o *UnixGroupModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix group modify o k response has a 3xx status code
func (o *UnixGroupModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix group modify o k response has a 4xx status code
func (o *UnixGroupModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix group modify o k response has a 5xx status code
func (o *UnixGroupModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unix group modify o k response a status code equal to that given
func (o *UnixGroupModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unix group modify o k response
func (o *UnixGroupModifyOK) Code() int {
	return 200
}

func (o *UnixGroupModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/unix-groups/{svm.uuid}/{name}][%d] unixGroupModifyOK", 200)
}

func (o *UnixGroupModifyOK) String() string {
	return fmt.Sprintf("[PATCH /name-services/unix-groups/{svm.uuid}/{name}][%d] unixGroupModifyOK", 200)
}

func (o *UnixGroupModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnixGroupModifyDefault creates a UnixGroupModifyDefault with default headers values
func NewUnixGroupModifyDefault(code int) *UnixGroupModifyDefault {
	return &UnixGroupModifyDefault{
		_statusCode: code,
	}
}

/*
	UnixGroupModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 23724141   | Duplicate group ID. Group ID must be unique.|
*/
type UnixGroupModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unix group modify default response has a 2xx status code
func (o *UnixGroupModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix group modify default response has a 3xx status code
func (o *UnixGroupModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix group modify default response has a 4xx status code
func (o *UnixGroupModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix group modify default response has a 5xx status code
func (o *UnixGroupModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix group modify default response a status code equal to that given
func (o *UnixGroupModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unix group modify default response
func (o *UnixGroupModifyDefault) Code() int {
	return o._statusCode
}

func (o *UnixGroupModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/unix-groups/{svm.uuid}/{name}][%d] unix_group_modify default %s", o._statusCode, payload)
}

func (o *UnixGroupModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/unix-groups/{svm.uuid}/{name}][%d] unix_group_modify default %s", o._statusCode, payload)
}

func (o *UnixGroupModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
