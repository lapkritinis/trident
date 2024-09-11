// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// FabricGetReader is a Reader for the FabricGet structure.
type FabricGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FabricGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFabricGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFabricGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFabricGetOK creates a FabricGetOK with default headers values
func NewFabricGetOK() *FabricGetOK {
	return &FabricGetOK{}
}

/*
FabricGetOK describes a response with status code 200, with default header values.

OK
*/
type FabricGetOK struct {
	Payload *models.Fabric
}

// IsSuccess returns true when this fabric get o k response has a 2xx status code
func (o *FabricGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fabric get o k response has a 3xx status code
func (o *FabricGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fabric get o k response has a 4xx status code
func (o *FabricGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fabric get o k response has a 5xx status code
func (o *FabricGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fabric get o k response a status code equal to that given
func (o *FabricGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fabric get o k response
func (o *FabricGetOK) Code() int {
	return 200
}

func (o *FabricGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/fabrics/{name}][%d] fabricGetOK %s", 200, payload)
}

func (o *FabricGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/fabrics/{name}][%d] fabricGetOK %s", 200, payload)
}

func (o *FabricGetOK) GetPayload() *models.Fabric {
	return o.Payload
}

func (o *FabricGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fabric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFabricGetDefault creates a FabricGetDefault with default headers values
func NewFabricGetDefault(code int) *FabricGetDefault {
	return &FabricGetDefault{
		_statusCode: code,
	}
}

/*
	FabricGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5373982 | An invalid WWN was specified. The length is incorrect. |
| 5373983 | An invalid WWN was specified. The format is incorrect. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FabricGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fabric get default response has a 2xx status code
func (o *FabricGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fabric get default response has a 3xx status code
func (o *FabricGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fabric get default response has a 4xx status code
func (o *FabricGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fabric get default response has a 5xx status code
func (o *FabricGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fabric get default response a status code equal to that given
func (o *FabricGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fabric get default response
func (o *FabricGetDefault) Code() int {
	return o._statusCode
}

func (o *FabricGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/fabrics/{name}][%d] fabric_get default %s", o._statusCode, payload)
}

func (o *FabricGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/fabrics/{name}][%d] fabric_get default %s", o._statusCode, payload)
}

func (o *FabricGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FabricGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
