// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// LunMapDeleteReader is a Reader for the LunMapDelete structure.
type LunMapDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunMapDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunMapDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunMapDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunMapDeleteOK creates a LunMapDeleteOK with default headers values
func NewLunMapDeleteOK() *LunMapDeleteOK {
	return &LunMapDeleteOK{}
}

/*
LunMapDeleteOK describes a response with status code 200, with default header values.

OK
*/
type LunMapDeleteOK struct {
}

// IsSuccess returns true when this lun map delete o k response has a 2xx status code
func (o *LunMapDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun map delete o k response has a 3xx status code
func (o *LunMapDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun map delete o k response has a 4xx status code
func (o *LunMapDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun map delete o k response has a 5xx status code
func (o *LunMapDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lun map delete o k response a status code equal to that given
func (o *LunMapDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lun map delete o k response
func (o *LunMapDeleteOK) Code() int {
	return 200
}

func (o *LunMapDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lunMapDeleteOK", 200)
}

func (o *LunMapDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lunMapDeleteOK", 200)
}

func (o *LunMapDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLunMapDeleteDefault creates a LunMapDeleteDefault with default headers values
func NewLunMapDeleteDefault(code int) *LunMapDeleteDefault {
	return &LunMapDeleteDefault{
		_statusCode: code,
	}
}

/*
	LunMapDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN does not exist or is not accessible to the caller. |
| 5374878 | The specified initiator group does not exist, is not accessible to the caller, or is not in the same SVM as the specified LUN. |
| 5374922 | The specified LUN map does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LunMapDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this lun map delete default response has a 2xx status code
func (o *LunMapDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun map delete default response has a 3xx status code
func (o *LunMapDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun map delete default response has a 4xx status code
func (o *LunMapDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun map delete default response has a 5xx status code
func (o *LunMapDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun map delete default response a status code equal to that given
func (o *LunMapDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lun map delete default response
func (o *LunMapDeleteDefault) Code() int {
	return o._statusCode
}

func (o *LunMapDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lun_map_delete default %s", o._statusCode, payload)
}

func (o *LunMapDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lun_map_delete default %s", o._statusCode, payload)
}

func (o *LunMapDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunMapDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
