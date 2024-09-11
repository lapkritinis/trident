// Code generated by go-swagger; DO NOT EDIT.

package support

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

// EmsFilterRuleDeleteReader is a Reader for the EmsFilterRuleDelete structure.
type EmsFilterRuleDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsFilterRuleDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsFilterRuleDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsFilterRuleDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsFilterRuleDeleteOK creates a EmsFilterRuleDeleteOK with default headers values
func NewEmsFilterRuleDeleteOK() *EmsFilterRuleDeleteOK {
	return &EmsFilterRuleDeleteOK{}
}

/*
EmsFilterRuleDeleteOK describes a response with status code 200, with default header values.

OK
*/
type EmsFilterRuleDeleteOK struct {
}

// IsSuccess returns true when this ems filter rule delete o k response has a 2xx status code
func (o *EmsFilterRuleDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems filter rule delete o k response has a 3xx status code
func (o *EmsFilterRuleDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems filter rule delete o k response has a 4xx status code
func (o *EmsFilterRuleDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems filter rule delete o k response has a 5xx status code
func (o *EmsFilterRuleDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems filter rule delete o k response a status code equal to that given
func (o *EmsFilterRuleDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ems filter rule delete o k response
func (o *EmsFilterRuleDeleteOK) Code() int {
	return 200
}

func (o *EmsFilterRuleDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /support/ems/filters/{name}/rules/{index}][%d] emsFilterRuleDeleteOK", 200)
}

func (o *EmsFilterRuleDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /support/ems/filters/{name}/rules/{index}][%d] emsFilterRuleDeleteOK", 200)
}

func (o *EmsFilterRuleDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmsFilterRuleDeleteDefault creates a EmsFilterRuleDeleteDefault with default headers values
func NewEmsFilterRuleDeleteDefault(code int) *EmsFilterRuleDeleteDefault {
	return &EmsFilterRuleDeleteDefault{
		_statusCode: code,
	}
}

/*
	EmsFilterRuleDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 983091 | A default rule cannot be removed |
| 983092 | The index of the rule provided is outside the allowed range for the filter provided |
| 983095 | The rule index provided is invalid for the filter provided |
| 983110 | There are no user defined rules in the filter provided |
| 983113 | Default filters cannot be modified or removed |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type EmsFilterRuleDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems filter rule delete default response has a 2xx status code
func (o *EmsFilterRuleDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems filter rule delete default response has a 3xx status code
func (o *EmsFilterRuleDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems filter rule delete default response has a 4xx status code
func (o *EmsFilterRuleDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems filter rule delete default response has a 5xx status code
func (o *EmsFilterRuleDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems filter rule delete default response a status code equal to that given
func (o *EmsFilterRuleDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems filter rule delete default response
func (o *EmsFilterRuleDeleteDefault) Code() int {
	return o._statusCode
}

func (o *EmsFilterRuleDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /support/ems/filters/{name}/rules/{index}][%d] ems_filter_rule_delete default %s", o._statusCode, payload)
}

func (o *EmsFilterRuleDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /support/ems/filters/{name}/rules/{index}][%d] ems_filter_rule_delete default %s", o._statusCode, payload)
}

func (o *EmsFilterRuleDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsFilterRuleDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
