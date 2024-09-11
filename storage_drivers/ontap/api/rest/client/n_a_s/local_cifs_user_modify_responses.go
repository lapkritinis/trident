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

// LocalCifsUserModifyReader is a Reader for the LocalCifsUserModify structure.
type LocalCifsUserModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsUserModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsUserModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsUserModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsUserModifyOK creates a LocalCifsUserModifyOK with default headers values
func NewLocalCifsUserModifyOK() *LocalCifsUserModifyOK {
	return &LocalCifsUserModifyOK{}
}

/*
LocalCifsUserModifyOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsUserModifyOK struct {
}

// IsSuccess returns true when this local cifs user modify o k response has a 2xx status code
func (o *LocalCifsUserModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs user modify o k response has a 3xx status code
func (o *LocalCifsUserModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs user modify o k response has a 4xx status code
func (o *LocalCifsUserModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs user modify o k response has a 5xx status code
func (o *LocalCifsUserModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs user modify o k response a status code equal to that given
func (o *LocalCifsUserModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the local cifs user modify o k response
func (o *LocalCifsUserModifyOK) Code() int {
	return 200
}

func (o *LocalCifsUserModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] localCifsUserModifyOK", 200)
}

func (o *LocalCifsUserModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] localCifsUserModifyOK", 200)
}

func (o *LocalCifsUserModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLocalCifsUserModifyDefault creates a LocalCifsUserModifyDefault with default headers values
func NewLocalCifsUserModifyDefault(code int) *LocalCifsUserModifyDefault {
	return &LocalCifsUserModifyDefault{
		_statusCode: code,
	}
}

/*
	LocalCifsUserModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655661     | The user name should not exceed 20 characters. Also full_name and description should not exceed 256 characters. |
| 655668     | The specified user name contains illegal characters. |
| 655675     | The local domain name specified in the user name does not exist. |
| 655682     | The user name cannot be blank. |
| 655732     | Failed to rename a user. The error code returned details the failure along with the reason for the failure. Take corrective actions as per the specified reason. |
| 655733     | The specified password does not meet the password complexity requirements. |
| 655737     | To rename an existing user, the local domain specified in name must match the local domain of the user to be renamed. |
*/
type LocalCifsUserModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local cifs user modify default response has a 2xx status code
func (o *LocalCifsUserModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs user modify default response has a 3xx status code
func (o *LocalCifsUserModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs user modify default response has a 4xx status code
func (o *LocalCifsUserModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs user modify default response has a 5xx status code
func (o *LocalCifsUserModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs user modify default response a status code equal to that given
func (o *LocalCifsUserModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local cifs user modify default response
func (o *LocalCifsUserModifyDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsUserModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] local_cifs_user_modify default %s", o._statusCode, payload)
}

func (o *LocalCifsUserModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] local_cifs_user_modify default %s", o._statusCode, payload)
}

func (o *LocalCifsUserModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsUserModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
