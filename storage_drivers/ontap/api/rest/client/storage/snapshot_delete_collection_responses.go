// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnapshotDeleteCollectionReader is a Reader for the SnapshotDeleteCollection structure.
type SnapshotDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnapshotDeleteCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotDeleteCollectionOK creates a SnapshotDeleteCollectionOK with default headers values
func NewSnapshotDeleteCollectionOK() *SnapshotDeleteCollectionOK {
	return &SnapshotDeleteCollectionOK{}
}

/*
SnapshotDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotDeleteCollectionOK struct {
	Payload *models.SnapshotJobLinkResponse
}

// IsSuccess returns true when this snapshot delete collection o k response has a 2xx status code
func (o *SnapshotDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot delete collection o k response has a 3xx status code
func (o *SnapshotDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot delete collection o k response has a 4xx status code
func (o *SnapshotDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot delete collection o k response has a 5xx status code
func (o *SnapshotDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot delete collection o k response a status code equal to that given
func (o *SnapshotDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapshot delete collection o k response
func (o *SnapshotDeleteCollectionOK) Code() int {
	return 200
}

func (o *SnapshotDeleteCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/snapshots][%d] snapshotDeleteCollectionOK %s", 200, payload)
}

func (o *SnapshotDeleteCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/snapshots][%d] snapshotDeleteCollectionOK %s", 200, payload)
}

func (o *SnapshotDeleteCollectionOK) GetPayload() *models.SnapshotJobLinkResponse {
	return o.Payload
}

func (o *SnapshotDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotDeleteCollectionAccepted creates a SnapshotDeleteCollectionAccepted with default headers values
func NewSnapshotDeleteCollectionAccepted() *SnapshotDeleteCollectionAccepted {
	return &SnapshotDeleteCollectionAccepted{}
}

/*
SnapshotDeleteCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapshotDeleteCollectionAccepted struct {
	Payload *models.SnapshotJobLinkResponse
}

// IsSuccess returns true when this snapshot delete collection accepted response has a 2xx status code
func (o *SnapshotDeleteCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot delete collection accepted response has a 3xx status code
func (o *SnapshotDeleteCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot delete collection accepted response has a 4xx status code
func (o *SnapshotDeleteCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot delete collection accepted response has a 5xx status code
func (o *SnapshotDeleteCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot delete collection accepted response a status code equal to that given
func (o *SnapshotDeleteCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snapshot delete collection accepted response
func (o *SnapshotDeleteCollectionAccepted) Code() int {
	return 202
}

func (o *SnapshotDeleteCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/snapshots][%d] snapshotDeleteCollectionAccepted %s", 202, payload)
}

func (o *SnapshotDeleteCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/snapshots][%d] snapshotDeleteCollectionAccepted %s", 202, payload)
}

func (o *SnapshotDeleteCollectionAccepted) GetPayload() *models.SnapshotJobLinkResponse {
	return o.Payload
}

func (o *SnapshotDeleteCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotDeleteCollectionDefault creates a SnapshotDeleteCollectionDefault with default headers values
func NewSnapshotDeleteCollectionDefault(code int) *SnapshotDeleteCollectionDefault {
	return &SnapshotDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	SnapshotDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Code

| Error Code | Description |
| ---------- | ----------- |
| 2          | An invalid value was entered for one of the fields. |
| 524481     | Snapshot was not deleted because the associated volume is locked or fenced. |
| 1638521    | Snapshots can only be deleted on read/write (RW) volumes. |
| 1638538    | Cannot determine the status of the snapshot delete operation for the specified volume. |
| 1638543    | Failed to delete snapshot because it has an owner. |
| 1638555    | The specified snapshot has not expired or is locked. |
| 1638600    | The snapshot does not exist. |
*/
type SnapshotDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapshot delete collection default response has a 2xx status code
func (o *SnapshotDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot delete collection default response has a 3xx status code
func (o *SnapshotDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot delete collection default response has a 4xx status code
func (o *SnapshotDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot delete collection default response has a 5xx status code
func (o *SnapshotDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot delete collection default response a status code equal to that given
func (o *SnapshotDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapshot delete collection default response
func (o *SnapshotDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/snapshots][%d] snapshot_delete_collection default %s", o._statusCode, payload)
}

func (o *SnapshotDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/snapshots][%d] snapshot_delete_collection default %s", o._statusCode, payload)
}

func (o *SnapshotDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SnapshotDeleteCollectionBody snapshot delete collection body
swagger:model SnapshotDeleteCollectionBody
*/
type SnapshotDeleteCollectionBody struct {

	// delta
	Delta *models.SnapshotDelta `json:"delta,omitempty"`

	// Space reclaimed when the snapshot is deleted, in bytes.
	ReclaimableSpace *int64 `json:"reclaimable_space,omitempty"`

	// snapshot response inline records
	SnapshotResponseInlineRecords []*models.Snapshot `json:"records,omitempty"`
}

// Validate validates this snapshot delete collection body
func (o *SnapshotDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDelta(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSnapshotResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnapshotDeleteCollectionBody) validateDelta(formats strfmt.Registry) error {
	if swag.IsZero(o.Delta) { // not required
		return nil
	}

	if o.Delta != nil {
		if err := o.Delta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "delta")
			}
			return err
		}
	}

	return nil
}

func (o *SnapshotDeleteCollectionBody) validateSnapshotResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.SnapshotResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.SnapshotResponseInlineRecords); i++ {
		if swag.IsZero(o.SnapshotResponseInlineRecords[i]) { // not required
			continue
		}

		if o.SnapshotResponseInlineRecords[i] != nil {
			if err := o.SnapshotResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this snapshot delete collection body based on the context it is used
func (o *SnapshotDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDelta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSnapshotResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnapshotDeleteCollectionBody) contextValidateDelta(ctx context.Context, formats strfmt.Registry) error {

	if o.Delta != nil {
		if err := o.Delta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "delta")
			}
			return err
		}
	}

	return nil
}

func (o *SnapshotDeleteCollectionBody) contextValidateSnapshotResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SnapshotResponseInlineRecords); i++ {

		if o.SnapshotResponseInlineRecords[i] != nil {
			if err := o.SnapshotResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SnapshotDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SnapshotDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res SnapshotDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
