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

// QtreeCreateReader is a Reader for the QtreeCreate structure.
type QtreeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QtreeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewQtreeCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewQtreeCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQtreeCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQtreeCreateCreated creates a QtreeCreateCreated with default headers values
func NewQtreeCreateCreated() *QtreeCreateCreated {
	return &QtreeCreateCreated{}
}

/*
QtreeCreateCreated describes a response with status code 201, with default header values.

Created
*/
type QtreeCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.QtreeJobLinkResponse
}

// IsSuccess returns true when this qtree create created response has a 2xx status code
func (o *QtreeCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qtree create created response has a 3xx status code
func (o *QtreeCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qtree create created response has a 4xx status code
func (o *QtreeCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this qtree create created response has a 5xx status code
func (o *QtreeCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this qtree create created response a status code equal to that given
func (o *QtreeCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the qtree create created response
func (o *QtreeCreateCreated) Code() int {
	return 201
}

func (o *QtreeCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qtrees][%d] qtreeCreateCreated %s", 201, payload)
}

func (o *QtreeCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qtrees][%d] qtreeCreateCreated %s", 201, payload)
}

func (o *QtreeCreateCreated) GetPayload() *models.QtreeJobLinkResponse {
	return o.Payload
}

func (o *QtreeCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.QtreeJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQtreeCreateAccepted creates a QtreeCreateAccepted with default headers values
func NewQtreeCreateAccepted() *QtreeCreateAccepted {
	return &QtreeCreateAccepted{}
}

/*
QtreeCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QtreeCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.QtreeJobLinkResponse
}

// IsSuccess returns true when this qtree create accepted response has a 2xx status code
func (o *QtreeCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qtree create accepted response has a 3xx status code
func (o *QtreeCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qtree create accepted response has a 4xx status code
func (o *QtreeCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this qtree create accepted response has a 5xx status code
func (o *QtreeCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this qtree create accepted response a status code equal to that given
func (o *QtreeCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the qtree create accepted response
func (o *QtreeCreateAccepted) Code() int {
	return 202
}

func (o *QtreeCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qtrees][%d] qtreeCreateAccepted %s", 202, payload)
}

func (o *QtreeCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qtrees][%d] qtreeCreateAccepted %s", 202, payload)
}

func (o *QtreeCreateAccepted) GetPayload() *models.QtreeJobLinkResponse {
	return o.Payload
}

func (o *QtreeCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.QtreeJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQtreeCreateDefault creates a QtreeCreateDefault with default headers values
func NewQtreeCreateDefault(code int) *QtreeCreateDefault {
	return &QtreeCreateDefault{
		_statusCode: code,
	}
}

/*
	QtreeCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262245 | Invalid field value. |
| 262247 | Invalid field value. |
| 917525 | The specified volume does not exist in Vserver. |
| 917927 | The specified volume was not found. |
| 918232 | Either `volume.name` or `volume.uuid` must be provided. |
| 918236 | The specified `volume.uuid` and `volume.name` refer to different volumes. |
| 1703954 | Export Policy name specified is invalid. |
| 2621462 | The specified SVM does not exist. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be provided. |
| 5242881 | Cannot create qtree because the volume is read-only. |
| 5242886 | Failed to create qtree. |
| 5242894 | Qtree with empty name "" is not allowed, as that is reserved for the default qtree. |
| 5242900 | Qtree not supported on FlexCache volume |
| 5242948 | Qtree is not supported on FlexCache origin volume. |
| 5242951 | Export Policy supplied does not belong to the specified Export Policy ID. |
| 5242952 | Export Policy ID specified is invalid. |
| 5242953 | Qtree name must be provided. |
| 5242967 | UNIX user or group ID must be 32-bit unsigned integer. |
| 5242970 | FlexCache create is in progress for the volume. |
| 5242978 | The maximum number of qtrees for which extended performance monitoring can be enabled has been reached. Retry the POST request with ext_performance_monitoring.enabled set to false. |
| 5242979 | Qtree creation succeeded but failed to enable extended performance monitoring on the qtree. |
| 6622064 | Security-style NTFS is not supported on a SnapMirror active sync relationship volume. |
| 8454348 | QoS on qtrees is not supported because not all nodes in the cluster can support it. |
| 9437324 | The security style unified is not supported. |
| 23724050 | Failed to resolve user or group name. |
| 66846755 | Failed to determine whether volume is a FlexCache volume or not. |
| 66846839 | Failed to determine the effective cluster version of all the nodes hosting FlexCache volumes connected to FlexCache origin volume. |
| 92405926 | Qtree operation failed on the specified object store volume. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type QtreeCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this qtree create default response has a 2xx status code
func (o *QtreeCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qtree create default response has a 3xx status code
func (o *QtreeCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qtree create default response has a 4xx status code
func (o *QtreeCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qtree create default response has a 5xx status code
func (o *QtreeCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qtree create default response a status code equal to that given
func (o *QtreeCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the qtree create default response
func (o *QtreeCreateDefault) Code() int {
	return o._statusCode
}

func (o *QtreeCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qtrees][%d] qtree_create default %s", o._statusCode, payload)
}

func (o *QtreeCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qtrees][%d] qtree_create default %s", o._statusCode, payload)
}

func (o *QtreeCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QtreeCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
