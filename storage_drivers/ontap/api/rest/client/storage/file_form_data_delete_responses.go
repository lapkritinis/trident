// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// FileFormDataDeleteReader is a Reader for the FileFormDataDelete structure.
type FileFormDataDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileFormDataDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileFormDataDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewFileFormDataDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileFormDataDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileFormDataDeleteOK creates a FileFormDataDeleteOK with default headers values
func NewFileFormDataDeleteOK() *FileFormDataDeleteOK {
	return &FileFormDataDeleteOK{}
}

/*
FileFormDataDeleteOK describes a response with status code 200, with default header values.

OK
*/
type FileFormDataDeleteOK struct {
}

// IsSuccess returns true when this file form data delete o k response has a 2xx status code
func (o *FileFormDataDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file form data delete o k response has a 3xx status code
func (o *FileFormDataDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file form data delete o k response has a 4xx status code
func (o *FileFormDataDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this file form data delete o k response has a 5xx status code
func (o *FileFormDataDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this file form data delete o k response a status code equal to that given
func (o *FileFormDataDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the file form data delete o k response
func (o *FileFormDataDeleteOK) Code() int {
	return 200
}

func (o *FileFormDataDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE //storage/volumes/{volume.uuid}/files/{path}][%d] fileFormDataDeleteOK", 200)
}

func (o *FileFormDataDeleteOK) String() string {
	return fmt.Sprintf("[DELETE //storage/volumes/{volume.uuid}/files/{path}][%d] fileFormDataDeleteOK", 200)
}

func (o *FileFormDataDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFileFormDataDeleteAccepted creates a FileFormDataDeleteAccepted with default headers values
func NewFileFormDataDeleteAccepted() *FileFormDataDeleteAccepted {
	return &FileFormDataDeleteAccepted{}
}

/*
FileFormDataDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type FileFormDataDeleteAccepted struct {
	Payload *models.FileInfoJobLinkResponse
}

// IsSuccess returns true when this file form data delete accepted response has a 2xx status code
func (o *FileFormDataDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file form data delete accepted response has a 3xx status code
func (o *FileFormDataDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file form data delete accepted response has a 4xx status code
func (o *FileFormDataDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this file form data delete accepted response has a 5xx status code
func (o *FileFormDataDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this file form data delete accepted response a status code equal to that given
func (o *FileFormDataDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the file form data delete accepted response
func (o *FileFormDataDeleteAccepted) Code() int {
	return 202
}

func (o *FileFormDataDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE //storage/volumes/{volume.uuid}/files/{path}][%d] fileFormDataDeleteAccepted %s", 202, payload)
}

func (o *FileFormDataDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE //storage/volumes/{volume.uuid}/files/{path}][%d] fileFormDataDeleteAccepted %s", 202, payload)
}

func (o *FileFormDataDeleteAccepted) GetPayload() *models.FileInfoJobLinkResponse {
	return o.Payload
}

func (o *FileFormDataDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileInfoJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileFormDataDeleteDefault creates a FileFormDataDeleteDefault with default headers values
func NewFileFormDataDeleteDefault(code int) *FileFormDataDeleteDefault {
	return &FileFormDataDeleteDefault{
		_statusCode: code,
	}
}

/*
	FileFormDataDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 131102 | Read-only file system. |
| 131138 | Directory not empty. |
| 918235 | A volume with UUID {volume.uuid} was not found. |
| 6488081 | The {field} field is not supported for DELETE operations. |
| 6488110 | A volume delete is not supported on this endpoint. |
| 6684674 | No such file or directory. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FileFormDataDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file form data delete default response has a 2xx status code
func (o *FileFormDataDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file form data delete default response has a 3xx status code
func (o *FileFormDataDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file form data delete default response has a 4xx status code
func (o *FileFormDataDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file form data delete default response has a 5xx status code
func (o *FileFormDataDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file form data delete default response a status code equal to that given
func (o *FileFormDataDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file form data delete default response
func (o *FileFormDataDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FileFormDataDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE //storage/volumes/{volume.uuid}/files/{path}][%d] file_form_data_delete default %s", o._statusCode, payload)
}

func (o *FileFormDataDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE //storage/volumes/{volume.uuid}/files/{path}][%d] file_form_data_delete default %s", o._statusCode, payload)
}

func (o *FileFormDataDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileFormDataDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
