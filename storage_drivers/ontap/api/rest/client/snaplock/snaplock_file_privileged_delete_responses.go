// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// SnaplockFilePrivilegedDeleteReader is a Reader for the SnaplockFilePrivilegedDelete structure.
type SnaplockFilePrivilegedDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockFilePrivilegedDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockFilePrivilegedDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnaplockFilePrivilegedDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockFilePrivilegedDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockFilePrivilegedDeleteOK creates a SnaplockFilePrivilegedDeleteOK with default headers values
func NewSnaplockFilePrivilegedDeleteOK() *SnaplockFilePrivilegedDeleteOK {
	return &SnaplockFilePrivilegedDeleteOK{}
}

/*
SnaplockFilePrivilegedDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockFilePrivilegedDeleteOK struct {
	Payload *models.SnaplockFileRetentionJobLinkResponse
}

// IsSuccess returns true when this snaplock file privileged delete o k response has a 2xx status code
func (o *SnaplockFilePrivilegedDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock file privileged delete o k response has a 3xx status code
func (o *SnaplockFilePrivilegedDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock file privileged delete o k response has a 4xx status code
func (o *SnaplockFilePrivilegedDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock file privileged delete o k response has a 5xx status code
func (o *SnaplockFilePrivilegedDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock file privileged delete o k response a status code equal to that given
func (o *SnaplockFilePrivilegedDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock file privileged delete o k response
func (o *SnaplockFilePrivilegedDeleteOK) Code() int {
	return 200
}

func (o *SnaplockFilePrivilegedDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplockFilePrivilegedDeleteOK %s", 200, payload)
}

func (o *SnaplockFilePrivilegedDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplockFilePrivilegedDeleteOK %s", 200, payload)
}

func (o *SnaplockFilePrivilegedDeleteOK) GetPayload() *models.SnaplockFileRetentionJobLinkResponse {
	return o.Payload
}

func (o *SnaplockFilePrivilegedDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockFileRetentionJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockFilePrivilegedDeleteAccepted creates a SnaplockFilePrivilegedDeleteAccepted with default headers values
func NewSnaplockFilePrivilegedDeleteAccepted() *SnaplockFilePrivilegedDeleteAccepted {
	return &SnaplockFilePrivilegedDeleteAccepted{}
}

/*
SnaplockFilePrivilegedDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnaplockFilePrivilegedDeleteAccepted struct {
	Payload *models.SnaplockFileRetentionJobLinkResponse
}

// IsSuccess returns true when this snaplock file privileged delete accepted response has a 2xx status code
func (o *SnaplockFilePrivilegedDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock file privileged delete accepted response has a 3xx status code
func (o *SnaplockFilePrivilegedDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock file privileged delete accepted response has a 4xx status code
func (o *SnaplockFilePrivilegedDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock file privileged delete accepted response has a 5xx status code
func (o *SnaplockFilePrivilegedDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock file privileged delete accepted response a status code equal to that given
func (o *SnaplockFilePrivilegedDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snaplock file privileged delete accepted response
func (o *SnaplockFilePrivilegedDeleteAccepted) Code() int {
	return 202
}

func (o *SnaplockFilePrivilegedDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplockFilePrivilegedDeleteAccepted %s", 202, payload)
}

func (o *SnaplockFilePrivilegedDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplockFilePrivilegedDeleteAccepted %s", 202, payload)
}

func (o *SnaplockFilePrivilegedDeleteAccepted) GetPayload() *models.SnaplockFileRetentionJobLinkResponse {
	return o.Payload
}

func (o *SnaplockFilePrivilegedDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockFileRetentionJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockFilePrivilegedDeleteDefault creates a SnaplockFilePrivilegedDeleteDefault with default headers values
func NewSnaplockFilePrivilegedDeleteDefault(code int) *SnaplockFilePrivilegedDeleteDefault {
	return &SnaplockFilePrivilegedDeleteDefault{
		_statusCode: code,
	}
}

/*
	SnaplockFilePrivilegedDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 262179      | Unexpected argument \"<file_name>\"  |
| 6691623     | User is not authorized  |
| 13763162    | SnapLock audit log volume is not configured for the SVM |
| 13763280    | Only a user with the security login role \"vsadmin-snaplock\" is allowed to perform this operation |
| 14090347    | File path must be in the format \"\/\<dir\>\/\<file path\>\"  |
| 917804      | Path should be given in the format \"\/\vol\/\<volume name>\/\<file path>\".
*/
type SnaplockFilePrivilegedDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock file privileged delete default response has a 2xx status code
func (o *SnaplockFilePrivilegedDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock file privileged delete default response has a 3xx status code
func (o *SnaplockFilePrivilegedDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock file privileged delete default response has a 4xx status code
func (o *SnaplockFilePrivilegedDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock file privileged delete default response has a 5xx status code
func (o *SnaplockFilePrivilegedDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock file privileged delete default response a status code equal to that given
func (o *SnaplockFilePrivilegedDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock file privileged delete default response
func (o *SnaplockFilePrivilegedDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockFilePrivilegedDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplock_file_privileged_delete default %s", o._statusCode, payload)
}

func (o *SnaplockFilePrivilegedDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplock_file_privileged_delete default %s", o._statusCode, payload)
}

func (o *SnaplockFilePrivilegedDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockFilePrivilegedDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
