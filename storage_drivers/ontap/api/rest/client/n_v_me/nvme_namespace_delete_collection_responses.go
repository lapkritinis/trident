// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NvmeNamespaceDeleteCollectionReader is a Reader for the NvmeNamespaceDeleteCollection structure.
type NvmeNamespaceDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeNamespaceDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeNamespaceDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewNvmeNamespaceDeleteCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeNamespaceDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeNamespaceDeleteCollectionOK creates a NvmeNamespaceDeleteCollectionOK with default headers values
func NewNvmeNamespaceDeleteCollectionOK() *NvmeNamespaceDeleteCollectionOK {
	return &NvmeNamespaceDeleteCollectionOK{}
}

/*
NvmeNamespaceDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type NvmeNamespaceDeleteCollectionOK struct {
}

// IsSuccess returns true when this nvme namespace delete collection o k response has a 2xx status code
func (o *NvmeNamespaceDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme namespace delete collection o k response has a 3xx status code
func (o *NvmeNamespaceDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme namespace delete collection o k response has a 4xx status code
func (o *NvmeNamespaceDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme namespace delete collection o k response has a 5xx status code
func (o *NvmeNamespaceDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme namespace delete collection o k response a status code equal to that given
func (o *NvmeNamespaceDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme namespace delete collection o k response
func (o *NvmeNamespaceDeleteCollectionOK) Code() int {
	return 200
}

func (o *NvmeNamespaceDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/namespaces][%d] nvmeNamespaceDeleteCollectionOK", 200)
}

func (o *NvmeNamespaceDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /storage/namespaces][%d] nvmeNamespaceDeleteCollectionOK", 200)
}

func (o *NvmeNamespaceDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNvmeNamespaceDeleteCollectionAccepted creates a NvmeNamespaceDeleteCollectionAccepted with default headers values
func NewNvmeNamespaceDeleteCollectionAccepted() *NvmeNamespaceDeleteCollectionAccepted {
	return &NvmeNamespaceDeleteCollectionAccepted{}
}

/*
NvmeNamespaceDeleteCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type NvmeNamespaceDeleteCollectionAccepted struct {
	Payload *models.NvmeNamespaceJobLinkResponse
}

// IsSuccess returns true when this nvme namespace delete collection accepted response has a 2xx status code
func (o *NvmeNamespaceDeleteCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme namespace delete collection accepted response has a 3xx status code
func (o *NvmeNamespaceDeleteCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme namespace delete collection accepted response has a 4xx status code
func (o *NvmeNamespaceDeleteCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme namespace delete collection accepted response has a 5xx status code
func (o *NvmeNamespaceDeleteCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme namespace delete collection accepted response a status code equal to that given
func (o *NvmeNamespaceDeleteCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the nvme namespace delete collection accepted response
func (o *NvmeNamespaceDeleteCollectionAccepted) Code() int {
	return 202
}

func (o *NvmeNamespaceDeleteCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/namespaces][%d] nvmeNamespaceDeleteCollectionAccepted %s", 202, payload)
}

func (o *NvmeNamespaceDeleteCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/namespaces][%d] nvmeNamespaceDeleteCollectionAccepted %s", 202, payload)
}

func (o *NvmeNamespaceDeleteCollectionAccepted) GetPayload() *models.NvmeNamespaceJobLinkResponse {
	return o.Payload
}

func (o *NvmeNamespaceDeleteCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeNamespaceJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeNamespaceDeleteCollectionDefault creates a NvmeNamespaceDeleteCollectionDefault with default headers values
func NewNvmeNamespaceDeleteCollectionDefault(code int) *NvmeNamespaceDeleteCollectionDefault {
	return &NvmeNamespaceDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	NvmeNamespaceDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 4 | The specified namespace was not found. |
| 72089796 | The namespace must be unmapped before deletion. |
| 72090016 | The namespace's aggregate is offline. The aggregate must be online to modify or remove the namespace. |
| 72090017 | The namespace's volume is offline. The volume must be online to modify or remove the namespace. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeNamespaceDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme namespace delete collection default response has a 2xx status code
func (o *NvmeNamespaceDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme namespace delete collection default response has a 3xx status code
func (o *NvmeNamespaceDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme namespace delete collection default response has a 4xx status code
func (o *NvmeNamespaceDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme namespace delete collection default response has a 5xx status code
func (o *NvmeNamespaceDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme namespace delete collection default response a status code equal to that given
func (o *NvmeNamespaceDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme namespace delete collection default response
func (o *NvmeNamespaceDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *NvmeNamespaceDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/namespaces][%d] nvme_namespace_delete_collection default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/namespaces][%d] nvme_namespace_delete_collection default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeNamespaceDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NvmeNamespaceDeleteCollectionBody nvme namespace delete collection body
swagger:model NvmeNamespaceDeleteCollectionBody
*/
type NvmeNamespaceDeleteCollectionBody struct {

	// nvme namespace response inline records
	NvmeNamespaceResponseInlineRecords []*models.NvmeNamespace `json:"records,omitempty"`
}

// Validate validates this nvme namespace delete collection body
func (o *NvmeNamespaceDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNvmeNamespaceResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NvmeNamespaceDeleteCollectionBody) validateNvmeNamespaceResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.NvmeNamespaceResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.NvmeNamespaceResponseInlineRecords); i++ {
		if swag.IsZero(o.NvmeNamespaceResponseInlineRecords[i]) { // not required
			continue
		}

		if o.NvmeNamespaceResponseInlineRecords[i] != nil {
			if err := o.NvmeNamespaceResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nvme namespace delete collection body based on the context it is used
func (o *NvmeNamespaceDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNvmeNamespaceResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NvmeNamespaceDeleteCollectionBody) contextValidateNvmeNamespaceResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.NvmeNamespaceResponseInlineRecords); i++ {

		if o.NvmeNamespaceResponseInlineRecords[i] != nil {
			if err := o.NvmeNamespaceResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *NvmeNamespaceDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NvmeNamespaceDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res NvmeNamespaceDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
