// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// CifsSymlinkMappingDeleteCollectionReader is a Reader for the CifsSymlinkMappingDeleteCollection structure.
type CifsSymlinkMappingDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSymlinkMappingDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsSymlinkMappingDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSymlinkMappingDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSymlinkMappingDeleteCollectionOK creates a CifsSymlinkMappingDeleteCollectionOK with default headers values
func NewCifsSymlinkMappingDeleteCollectionOK() *CifsSymlinkMappingDeleteCollectionOK {
	return &CifsSymlinkMappingDeleteCollectionOK{}
}

/*
CifsSymlinkMappingDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type CifsSymlinkMappingDeleteCollectionOK struct {
}

// IsSuccess returns true when this cifs symlink mapping delete collection o k response has a 2xx status code
func (o *CifsSymlinkMappingDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs symlink mapping delete collection o k response has a 3xx status code
func (o *CifsSymlinkMappingDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs symlink mapping delete collection o k response has a 4xx status code
func (o *CifsSymlinkMappingDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs symlink mapping delete collection o k response has a 5xx status code
func (o *CifsSymlinkMappingDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs symlink mapping delete collection o k response a status code equal to that given
func (o *CifsSymlinkMappingDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs symlink mapping delete collection o k response
func (o *CifsSymlinkMappingDeleteCollectionOK) Code() int {
	return 200
}

func (o *CifsSymlinkMappingDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/unix-symlink-mapping][%d] cifsSymlinkMappingDeleteCollectionOK", 200)
}

func (o *CifsSymlinkMappingDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/unix-symlink-mapping][%d] cifsSymlinkMappingDeleteCollectionOK", 200)
}

func (o *CifsSymlinkMappingDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsSymlinkMappingDeleteCollectionDefault creates a CifsSymlinkMappingDeleteCollectionDefault with default headers values
func NewCifsSymlinkMappingDeleteCollectionDefault(code int) *CifsSymlinkMappingDeleteCollectionDefault {
	return &CifsSymlinkMappingDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
CifsSymlinkMappingDeleteCollectionDefault describes a response with status code -1, with default header values.

Error
*/
type CifsSymlinkMappingDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs symlink mapping delete collection default response has a 2xx status code
func (o *CifsSymlinkMappingDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs symlink mapping delete collection default response has a 3xx status code
func (o *CifsSymlinkMappingDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs symlink mapping delete collection default response has a 4xx status code
func (o *CifsSymlinkMappingDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs symlink mapping delete collection default response has a 5xx status code
func (o *CifsSymlinkMappingDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs symlink mapping delete collection default response a status code equal to that given
func (o *CifsSymlinkMappingDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs symlink mapping delete collection default response
func (o *CifsSymlinkMappingDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *CifsSymlinkMappingDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/unix-symlink-mapping][%d] cifs_symlink_mapping_delete_collection default %s", o._statusCode, payload)
}

func (o *CifsSymlinkMappingDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/unix-symlink-mapping][%d] cifs_symlink_mapping_delete_collection default %s", o._statusCode, payload)
}

func (o *CifsSymlinkMappingDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSymlinkMappingDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CifsSymlinkMappingDeleteCollectionBody cifs symlink mapping delete collection body
swagger:model CifsSymlinkMappingDeleteCollectionBody
*/
type CifsSymlinkMappingDeleteCollectionBody struct {

	// cifs symlink mapping response inline records
	CifsSymlinkMappingResponseInlineRecords []*models.CifsSymlinkMapping `json:"records,omitempty"`
}

// Validate validates this cifs symlink mapping delete collection body
func (o *CifsSymlinkMappingDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCifsSymlinkMappingResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CifsSymlinkMappingDeleteCollectionBody) validateCifsSymlinkMappingResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.CifsSymlinkMappingResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.CifsSymlinkMappingResponseInlineRecords); i++ {
		if swag.IsZero(o.CifsSymlinkMappingResponseInlineRecords[i]) { // not required
			continue
		}

		if o.CifsSymlinkMappingResponseInlineRecords[i] != nil {
			if err := o.CifsSymlinkMappingResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cifs symlink mapping delete collection body based on the context it is used
func (o *CifsSymlinkMappingDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCifsSymlinkMappingResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CifsSymlinkMappingDeleteCollectionBody) contextValidateCifsSymlinkMappingResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.CifsSymlinkMappingResponseInlineRecords); i++ {

		if o.CifsSymlinkMappingResponseInlineRecords[i] != nil {
			if err := o.CifsSymlinkMappingResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *CifsSymlinkMappingDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CifsSymlinkMappingDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res CifsSymlinkMappingDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
