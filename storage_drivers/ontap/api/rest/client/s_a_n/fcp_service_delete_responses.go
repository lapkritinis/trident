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

// FcpServiceDeleteReader is a Reader for the FcpServiceDelete structure.
type FcpServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcpServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcpServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcpServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcpServiceDeleteOK creates a FcpServiceDeleteOK with default headers values
func NewFcpServiceDeleteOK() *FcpServiceDeleteOK {
	return &FcpServiceDeleteOK{}
}

/*
FcpServiceDeleteOK describes a response with status code 200, with default header values.

OK
*/
type FcpServiceDeleteOK struct {
}

// IsSuccess returns true when this fcp service delete o k response has a 2xx status code
func (o *FcpServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fcp service delete o k response has a 3xx status code
func (o *FcpServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fcp service delete o k response has a 4xx status code
func (o *FcpServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fcp service delete o k response has a 5xx status code
func (o *FcpServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fcp service delete o k response a status code equal to that given
func (o *FcpServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fcp service delete o k response
func (o *FcpServiceDeleteOK) Code() int {
	return 200
}

func (o *FcpServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/fcp/services/{svm.uuid}][%d] fcpServiceDeleteOK", 200)
}

func (o *FcpServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/fcp/services/{svm.uuid}][%d] fcpServiceDeleteOK", 200)
}

func (o *FcpServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFcpServiceDeleteDefault creates a FcpServiceDeleteDefault with default headers values
func NewFcpServiceDeleteDefault(code int) *FcpServiceDeleteDefault {
	return &FcpServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
	FcpServiceDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | An SVM with the specified UUID does not exist. |
| 5373960 | The Fibre Channel Protocol service cannot be removed while it is enabled. |
| 5374083 | There is no Fibre Channel Protocol service for the specified SVM. |
| 5376452 | Service POST and DELETE are not supported on ASA r2. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FcpServiceDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fcp service delete default response has a 2xx status code
func (o *FcpServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fcp service delete default response has a 3xx status code
func (o *FcpServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fcp service delete default response has a 4xx status code
func (o *FcpServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fcp service delete default response has a 5xx status code
func (o *FcpServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fcp service delete default response a status code equal to that given
func (o *FcpServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fcp service delete default response
func (o *FcpServiceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FcpServiceDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/fcp/services/{svm.uuid}][%d] fcp_service_delete default %s", o._statusCode, payload)
}

func (o *FcpServiceDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/fcp/services/{svm.uuid}][%d] fcp_service_delete default %s", o._statusCode, payload)
}

func (o *FcpServiceDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcpServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
