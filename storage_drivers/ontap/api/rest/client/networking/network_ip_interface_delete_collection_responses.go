// Code generated by go-swagger; DO NOT EDIT.

package networking

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
	"github.com/go-openapi/validate"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetworkIPInterfaceDeleteCollectionReader is a Reader for the NetworkIPInterfaceDeleteCollection structure.
type NetworkIPInterfaceDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPInterfaceDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPInterfaceDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPInterfaceDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPInterfaceDeleteCollectionOK creates a NetworkIPInterfaceDeleteCollectionOK with default headers values
func NewNetworkIPInterfaceDeleteCollectionOK() *NetworkIPInterfaceDeleteCollectionOK {
	return &NetworkIPInterfaceDeleteCollectionOK{}
}

/*
NetworkIPInterfaceDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPInterfaceDeleteCollectionOK struct {
}

// IsSuccess returns true when this network Ip interface delete collection o k response has a 2xx status code
func (o *NetworkIPInterfaceDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip interface delete collection o k response has a 3xx status code
func (o *NetworkIPInterfaceDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip interface delete collection o k response has a 4xx status code
func (o *NetworkIPInterfaceDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip interface delete collection o k response has a 5xx status code
func (o *NetworkIPInterfaceDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip interface delete collection o k response a status code equal to that given
func (o *NetworkIPInterfaceDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network Ip interface delete collection o k response
func (o *NetworkIPInterfaceDeleteCollectionOK) Code() int {
	return 200
}

func (o *NetworkIPInterfaceDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ip/interfaces][%d] networkIpInterfaceDeleteCollectionOK", 200)
}

func (o *NetworkIPInterfaceDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /network/ip/interfaces][%d] networkIpInterfaceDeleteCollectionOK", 200)
}

func (o *NetworkIPInterfaceDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkIPInterfaceDeleteCollectionDefault creates a NetworkIPInterfaceDeleteCollectionDefault with default headers values
func NewNetworkIPInterfaceDeleteCollectionDefault(code int) *NetworkIPInterfaceDeleteCollectionDefault {
	return &NetworkIPInterfaceDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	NetworkIPInterfaceDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1966465 | LIF cannot be removed because it is required to maintain quorum on the node. |
| 5376953 | The interface is part of one or more portsets. Remove the interface from all portsets before deleting it. |
| 53281112 | Failed to delete the interface because it has an associated BGP peer group. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkIPInterfaceDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ip interface delete collection default response has a 2xx status code
func (o *NetworkIPInterfaceDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip interface delete collection default response has a 3xx status code
func (o *NetworkIPInterfaceDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip interface delete collection default response has a 4xx status code
func (o *NetworkIPInterfaceDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip interface delete collection default response has a 5xx status code
func (o *NetworkIPInterfaceDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip interface delete collection default response a status code equal to that given
func (o *NetworkIPInterfaceDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ip interface delete collection default response
func (o *NetworkIPInterfaceDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPInterfaceDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/interfaces][%d] network_ip_interface_delete_collection default %s", o._statusCode, payload)
}

func (o *NetworkIPInterfaceDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/interfaces][%d] network_ip_interface_delete_collection default %s", o._statusCode, payload)
}

func (o *NetworkIPInterfaceDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPInterfaceDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkIPInterfaceDeleteCollectionBody network IP interface delete collection body
swagger:model NetworkIPInterfaceDeleteCollectionBody
*/
type NetworkIPInterfaceDeleteCollectionBody struct {

	// ip interface response inline records
	IPInterfaceResponseInlineRecords []*models.IPInterface `json:"records,omitempty"`

	// recommend
	Recommend *models.IPInterfaceResponseInlineRecommend `json:"recommend,omitempty"`
}

// Validate validates this network IP interface delete collection body
func (o *NetworkIPInterfaceDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPInterfaceResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecommend(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkIPInterfaceDeleteCollectionBody) validateIPInterfaceResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.IPInterfaceResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.IPInterfaceResponseInlineRecords); i++ {
		if swag.IsZero(o.IPInterfaceResponseInlineRecords[i]) { // not required
			continue
		}

		if o.IPInterfaceResponseInlineRecords[i] != nil {
			if err := o.IPInterfaceResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NetworkIPInterfaceDeleteCollectionBody) validateRecommend(formats strfmt.Registry) error {
	if swag.IsZero(o.Recommend) { // not required
		return nil
	}

	if o.Recommend != nil {
		if err := o.Recommend.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "recommend")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network IP interface delete collection body based on the context it is used
func (o *NetworkIPInterfaceDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPInterfaceResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRecommend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkIPInterfaceDeleteCollectionBody) contextValidateIPInterfaceResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.IPInterfaceResponseInlineRecords); i++ {

		if o.IPInterfaceResponseInlineRecords[i] != nil {
			if err := o.IPInterfaceResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NetworkIPInterfaceDeleteCollectionBody) contextValidateRecommend(ctx context.Context, formats strfmt.Registry) error {

	if o.Recommend != nil {
		if err := o.Recommend.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "recommend")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetworkIPInterfaceDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkIPInterfaceDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res NetworkIPInterfaceDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
IPInterfaceResponseInlineRecommend Response properties specific to the LIF recommendation functionality.
//
swagger:model ip_interface_response_inline_recommend
*/
type IPInterfaceResponseInlineRecommend struct {

	// Messages describing the results of a LIF recommendation request.
	//
	// Read Only: true
	Messages []models.IPInterfaceRecommendMessage `json:"messages,omitempty"`
}

// Validate validates this ip interface response inline recommend
func (o *IPInterfaceResponseInlineRecommend) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ip interface response inline recommend based on the context it is used
func (o *IPInterfaceResponseInlineRecommend) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IPInterfaceResponseInlineRecommend) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"recommend"+"."+"messages", "body", []models.IPInterfaceRecommendMessage(o.Messages)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *IPInterfaceResponseInlineRecommend) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IPInterfaceResponseInlineRecommend) UnmarshalBinary(b []byte) error {
	var res IPInterfaceResponseInlineRecommend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
