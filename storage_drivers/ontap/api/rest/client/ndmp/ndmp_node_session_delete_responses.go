// Code generated by go-swagger; DO NOT EDIT.

package ndmp

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

// NdmpNodeSessionDeleteReader is a Reader for the NdmpNodeSessionDelete structure.
type NdmpNodeSessionDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NdmpNodeSessionDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNdmpNodeSessionDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNdmpNodeSessionDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNdmpNodeSessionDeleteOK creates a NdmpNodeSessionDeleteOK with default headers values
func NewNdmpNodeSessionDeleteOK() *NdmpNodeSessionDeleteOK {
	return &NdmpNodeSessionDeleteOK{}
}

/*
NdmpNodeSessionDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NdmpNodeSessionDeleteOK struct {
}

// IsSuccess returns true when this ndmp node session delete o k response has a 2xx status code
func (o *NdmpNodeSessionDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ndmp node session delete o k response has a 3xx status code
func (o *NdmpNodeSessionDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ndmp node session delete o k response has a 4xx status code
func (o *NdmpNodeSessionDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ndmp node session delete o k response has a 5xx status code
func (o *NdmpNodeSessionDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ndmp node session delete o k response a status code equal to that given
func (o *NdmpNodeSessionDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ndmp node session delete o k response
func (o *NdmpNodeSessionDeleteOK) Code() int {
	return 200
}

func (o *NdmpNodeSessionDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/ndmp/sessions/{owner.uuid}/{session.id}][%d] ndmpNodeSessionDeleteOK", 200)
}

func (o *NdmpNodeSessionDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/ndmp/sessions/{owner.uuid}/{session.id}][%d] ndmpNodeSessionDeleteOK", 200)
}

func (o *NdmpNodeSessionDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNdmpNodeSessionDeleteDefault creates a NdmpNodeSessionDeleteDefault with default headers values
func NewNdmpNodeSessionDeleteDefault(code int) *NdmpNodeSessionDeleteDefault {
	return &NdmpNodeSessionDeleteDefault{
		_statusCode: code,
	}
}

/*
	NdmpNodeSessionDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 68812802    | The UUID is not valid.|
| 68812803    | Failed to get the SVM name from the specified SVM UUID.|
| 68812804    | Failed to get the node name from the specified node UUID.|
| 68812805    | Failed to obtain the NDMP mode of operation.|
| 68812806    | UUID and Session ID are required.|
*/
type NdmpNodeSessionDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ndmp node session delete default response has a 2xx status code
func (o *NdmpNodeSessionDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ndmp node session delete default response has a 3xx status code
func (o *NdmpNodeSessionDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ndmp node session delete default response has a 4xx status code
func (o *NdmpNodeSessionDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ndmp node session delete default response has a 5xx status code
func (o *NdmpNodeSessionDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ndmp node session delete default response a status code equal to that given
func (o *NdmpNodeSessionDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ndmp node session delete default response
func (o *NdmpNodeSessionDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NdmpNodeSessionDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/ndmp/sessions/{owner.uuid}/{session.id}][%d] ndmp_node_session_delete default %s", o._statusCode, payload)
}

func (o *NdmpNodeSessionDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/ndmp/sessions/{owner.uuid}/{session.id}][%d] ndmp_node_session_delete default %s", o._statusCode, payload)
}

func (o *NdmpNodeSessionDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NdmpNodeSessionDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
