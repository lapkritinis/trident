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

// CloudStoreModifyReader is a Reader for the CloudStoreModify structure.
type CloudStoreModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudStoreModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudStoreModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCloudStoreModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCloudStoreModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudStoreModifyOK creates a CloudStoreModifyOK with default headers values
func NewCloudStoreModifyOK() *CloudStoreModifyOK {
	return &CloudStoreModifyOK{}
}

/*
CloudStoreModifyOK describes a response with status code 200, with default header values.

OK
*/
type CloudStoreModifyOK struct {
	Payload *models.CloudStoreJobLinkResponse
}

// IsSuccess returns true when this cloud store modify o k response has a 2xx status code
func (o *CloudStoreModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud store modify o k response has a 3xx status code
func (o *CloudStoreModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud store modify o k response has a 4xx status code
func (o *CloudStoreModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud store modify o k response has a 5xx status code
func (o *CloudStoreModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud store modify o k response a status code equal to that given
func (o *CloudStoreModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cloud store modify o k response
func (o *CloudStoreModifyOK) Code() int {
	return 200
}

func (o *CloudStoreModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/aggregates/{aggregate.uuid}/cloud-stores/{target.uuid}][%d] cloudStoreModifyOK %s", 200, payload)
}

func (o *CloudStoreModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/aggregates/{aggregate.uuid}/cloud-stores/{target.uuid}][%d] cloudStoreModifyOK %s", 200, payload)
}

func (o *CloudStoreModifyOK) GetPayload() *models.CloudStoreJobLinkResponse {
	return o.Payload
}

func (o *CloudStoreModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudStoreJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudStoreModifyAccepted creates a CloudStoreModifyAccepted with default headers values
func NewCloudStoreModifyAccepted() *CloudStoreModifyAccepted {
	return &CloudStoreModifyAccepted{}
}

/*
CloudStoreModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CloudStoreModifyAccepted struct {
	Payload *models.CloudStoreJobLinkResponse
}

// IsSuccess returns true when this cloud store modify accepted response has a 2xx status code
func (o *CloudStoreModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud store modify accepted response has a 3xx status code
func (o *CloudStoreModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud store modify accepted response has a 4xx status code
func (o *CloudStoreModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud store modify accepted response has a 5xx status code
func (o *CloudStoreModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud store modify accepted response a status code equal to that given
func (o *CloudStoreModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cloud store modify accepted response
func (o *CloudStoreModifyAccepted) Code() int {
	return 202
}

func (o *CloudStoreModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/aggregates/{aggregate.uuid}/cloud-stores/{target.uuid}][%d] cloudStoreModifyAccepted %s", 202, payload)
}

func (o *CloudStoreModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/aggregates/{aggregate.uuid}/cloud-stores/{target.uuid}][%d] cloudStoreModifyAccepted %s", 202, payload)
}

func (o *CloudStoreModifyAccepted) GetPayload() *models.CloudStoreJobLinkResponse {
	return o.Payload
}

func (o *CloudStoreModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudStoreJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudStoreModifyDefault creates a CloudStoreModifyDefault with default headers values
func NewCloudStoreModifyDefault(code int) *CloudStoreModifyDefault {
	return &CloudStoreModifyDefault{
		_statusCode: code,
	}
}

/*
	CloudStoreModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 787154 | The Object stores are currently not synchronized. Check the mirror resync status. |
| 787156 | Modifying the attributes of a mirror object store is not allowed. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type CloudStoreModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cloud store modify default response has a 2xx status code
func (o *CloudStoreModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cloud store modify default response has a 3xx status code
func (o *CloudStoreModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cloud store modify default response has a 4xx status code
func (o *CloudStoreModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cloud store modify default response has a 5xx status code
func (o *CloudStoreModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cloud store modify default response a status code equal to that given
func (o *CloudStoreModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cloud store modify default response
func (o *CloudStoreModifyDefault) Code() int {
	return o._statusCode
}

func (o *CloudStoreModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/aggregates/{aggregate.uuid}/cloud-stores/{target.uuid}][%d] cloud_store_modify default %s", o._statusCode, payload)
}

func (o *CloudStoreModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/aggregates/{aggregate.uuid}/cloud-stores/{target.uuid}][%d] cloud_store_modify default %s", o._statusCode, payload)
}

func (o *CloudStoreModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CloudStoreModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
