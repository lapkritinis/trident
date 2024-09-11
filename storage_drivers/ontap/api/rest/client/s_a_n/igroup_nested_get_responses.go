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

// IgroupNestedGetReader is a Reader for the IgroupNestedGet structure.
type IgroupNestedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupNestedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIgroupNestedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupNestedGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupNestedGetOK creates a IgroupNestedGetOK with default headers values
func NewIgroupNestedGetOK() *IgroupNestedGetOK {
	return &IgroupNestedGetOK{}
}

/*
IgroupNestedGetOK describes a response with status code 200, with default header values.

OK
*/
type IgroupNestedGetOK struct {
	Payload *models.IgroupNested
}

// IsSuccess returns true when this igroup nested get o k response has a 2xx status code
func (o *IgroupNestedGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this igroup nested get o k response has a 3xx status code
func (o *IgroupNestedGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this igroup nested get o k response has a 4xx status code
func (o *IgroupNestedGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this igroup nested get o k response has a 5xx status code
func (o *IgroupNestedGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this igroup nested get o k response a status code equal to that given
func (o *IgroupNestedGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the igroup nested get o k response
func (o *IgroupNestedGetOK) Code() int {
	return 200
}

func (o *IgroupNestedGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/igroups/{igroup.uuid}/igroups/{uuid}][%d] igroupNestedGetOK %s", 200, payload)
}

func (o *IgroupNestedGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/igroups/{igroup.uuid}/igroups/{uuid}][%d] igroupNestedGetOK %s", 200, payload)
}

func (o *IgroupNestedGetOK) GetPayload() *models.IgroupNested {
	return o.Payload
}

func (o *IgroupNestedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IgroupNested)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIgroupNestedGetDefault creates a IgroupNestedGetDefault with default headers values
func NewIgroupNestedGetDefault(code int) *IgroupNestedGetDefault {
	return &IgroupNestedGetDefault{
		_statusCode: code,
	}
}

/*
	IgroupNestedGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 4 | The nested initiator group is not a member of the initiator group. |
| 5374852 | The parent initiator group specified in the URI does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IgroupNestedGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this igroup nested get default response has a 2xx status code
func (o *IgroupNestedGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this igroup nested get default response has a 3xx status code
func (o *IgroupNestedGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this igroup nested get default response has a 4xx status code
func (o *IgroupNestedGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this igroup nested get default response has a 5xx status code
func (o *IgroupNestedGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this igroup nested get default response a status code equal to that given
func (o *IgroupNestedGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the igroup nested get default response
func (o *IgroupNestedGetDefault) Code() int {
	return o._statusCode
}

func (o *IgroupNestedGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/igroups/{igroup.uuid}/igroups/{uuid}][%d] igroup_nested_get default %s", o._statusCode, payload)
}

func (o *IgroupNestedGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/igroups/{igroup.uuid}/igroups/{uuid}][%d] igroup_nested_get default %s", o._statusCode, payload)
}

func (o *IgroupNestedGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupNestedGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
