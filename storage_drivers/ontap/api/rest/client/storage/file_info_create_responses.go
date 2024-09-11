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

// FileInfoCreateReader is a Reader for the FileInfoCreate structure.
type FileInfoCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileInfoCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFileInfoCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileInfoCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileInfoCreateCreated creates a FileInfoCreateCreated with default headers values
func NewFileInfoCreateCreated() *FileInfoCreateCreated {
	return &FileInfoCreateCreated{}
}

/*
FileInfoCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FileInfoCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this file info create created response has a 2xx status code
func (o *FileInfoCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file info create created response has a 3xx status code
func (o *FileInfoCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file info create created response has a 4xx status code
func (o *FileInfoCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this file info create created response has a 5xx status code
func (o *FileInfoCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this file info create created response a status code equal to that given
func (o *FileInfoCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the file info create created response
func (o *FileInfoCreateCreated) Code() int {
	return 201
}

func (o *FileInfoCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/files/{path}][%d] fileInfoCreateCreated", 201)
}

func (o *FileInfoCreateCreated) String() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/files/{path}][%d] fileInfoCreateCreated", 201)
}

func (o *FileInfoCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewFileInfoCreateDefault creates a FileInfoCreateDefault with default headers values
func NewFileInfoCreateDefault(code int) *FileInfoCreateDefault {
	return &FileInfoCreateDefault{
		_statusCode: code,
	}
}

/*
	FileInfoCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 917505 | The SVM does not exist. |
| 917525 | The volume in the symlink path does not exist in the SVM. |
| 917698 | Volume (vol name) in SVM (vserver) is not mounted in the namespace. |
| 917698 | The volume in the symlink path is not mounted in the namespace. |
| 918235 | A volume with UUID {volume.uuid} was not found. |
| 6488064 | This command is not supported. |
| 6488065 | The volume in the symlink path is invalid. |
| 6488066 | Mounting the unjunctioned volume in the symlink path failed. |
| 6488069 | Internal file error. |
| 6488084 | Failed to create {path} because the "unix_permissions" field was not specified. |
| 6488085 | Failed to create {path} because the "type" field was not specified. |
| 6488118 | File exists. |
| 8257536 | This operation is not supported for the system volume specified in the symlink path. |
| 8257541 | Failed to compute the SVM identification from this content. |
| 8257542 | This operation is not supported for the administrative SVM. |
| 9437549 | This operation is not allowed on SVMs with Infinite Volume. |
| 13172837 | This operation is not permitted because the SVM is locked for a migrate operation. |
| 111411203 | Failed to get the volume file system analytics report on the volume. |
| 111411204 | Internal error. Failed to retrieve the volume file system analytics report on the volume. |
| 111411207 | Volume file system analytics is not supported on volumes that contain LUNs. |
| 111411209 | Volume file system analytics is not supported on FlexCache volumes. |
| 111411210 | Volume file system analytics is not supported on audit staging volumes. |
| 111411211 | Volume file system analytics is not supported on object store server volumes. |
| 111411212 | Volume file system analytics is not supported on SnapMirror destination volumes. |
| 111411215 | Internal error. Volume file system analytics report timed out for volume volume.name in SVM svm.name. |
| 111411216 | Enabling or disabling volume file system analytics is not supported on individual FlexGroup constituents. |
| 111411217 | Volume file system analytics is not supported on SnapLock volumes. |
| 111411230 | Volume file system analytics is not supported on volumes that contain NVMe namespaces. |
| 111412203 | Volume file system analytics is not enabled on the volume. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FileInfoCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file info create default response has a 2xx status code
func (o *FileInfoCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file info create default response has a 3xx status code
func (o *FileInfoCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file info create default response has a 4xx status code
func (o *FileInfoCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file info create default response has a 5xx status code
func (o *FileInfoCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file info create default response a status code equal to that given
func (o *FileInfoCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file info create default response
func (o *FileInfoCreateDefault) Code() int {
	return o._statusCode
}

func (o *FileInfoCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/files/{path}][%d] file_info_create default %s", o._statusCode, payload)
}

func (o *FileInfoCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/files/{path}][%d] file_info_create default %s", o._statusCode, payload)
}

func (o *FileInfoCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileInfoCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
