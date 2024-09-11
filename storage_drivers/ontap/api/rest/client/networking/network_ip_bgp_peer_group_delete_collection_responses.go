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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetworkIPBgpPeerGroupDeleteCollectionReader is a Reader for the NetworkIPBgpPeerGroupDeleteCollection structure.
type NetworkIPBgpPeerGroupDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPBgpPeerGroupDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPBgpPeerGroupDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPBgpPeerGroupDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPBgpPeerGroupDeleteCollectionOK creates a NetworkIPBgpPeerGroupDeleteCollectionOK with default headers values
func NewNetworkIPBgpPeerGroupDeleteCollectionOK() *NetworkIPBgpPeerGroupDeleteCollectionOK {
	return &NetworkIPBgpPeerGroupDeleteCollectionOK{}
}

/*
NetworkIPBgpPeerGroupDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPBgpPeerGroupDeleteCollectionOK struct {
}

// IsSuccess returns true when this network Ip bgp peer group delete collection o k response has a 2xx status code
func (o *NetworkIPBgpPeerGroupDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip bgp peer group delete collection o k response has a 3xx status code
func (o *NetworkIPBgpPeerGroupDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip bgp peer group delete collection o k response has a 4xx status code
func (o *NetworkIPBgpPeerGroupDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip bgp peer group delete collection o k response has a 5xx status code
func (o *NetworkIPBgpPeerGroupDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip bgp peer group delete collection o k response a status code equal to that given
func (o *NetworkIPBgpPeerGroupDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network Ip bgp peer group delete collection o k response
func (o *NetworkIPBgpPeerGroupDeleteCollectionOK) Code() int {
	return 200
}

func (o *NetworkIPBgpPeerGroupDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ip/bgp/peer-groups][%d] networkIpBgpPeerGroupDeleteCollectionOK", 200)
}

func (o *NetworkIPBgpPeerGroupDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /network/ip/bgp/peer-groups][%d] networkIpBgpPeerGroupDeleteCollectionOK", 200)
}

func (o *NetworkIPBgpPeerGroupDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkIPBgpPeerGroupDeleteCollectionDefault creates a NetworkIPBgpPeerGroupDeleteCollectionDefault with default headers values
func NewNetworkIPBgpPeerGroupDeleteCollectionDefault(code int) *NetworkIPBgpPeerGroupDeleteCollectionDefault {
	return &NetworkIPBgpPeerGroupDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	NetworkIPBgpPeerGroupDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 53282019 | Internal error. Failed to remove BGP peer group on node. Wait a few minutes and try the command again. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkIPBgpPeerGroupDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ip bgp peer group delete collection default response has a 2xx status code
func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip bgp peer group delete collection default response has a 3xx status code
func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip bgp peer group delete collection default response has a 4xx status code
func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip bgp peer group delete collection default response has a 5xx status code
func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip bgp peer group delete collection default response a status code equal to that given
func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ip bgp peer group delete collection default response
func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/bgp/peer-groups][%d] network_ip_bgp_peer_group_delete_collection default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/bgp/peer-groups][%d] network_ip_bgp_peer_group_delete_collection default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkIPBgpPeerGroupDeleteCollectionBody network IP bgp peer group delete collection body
swagger:model NetworkIPBgpPeerGroupDeleteCollectionBody
*/
type NetworkIPBgpPeerGroupDeleteCollectionBody struct {

	// bgp peer group response inline records
	BgpPeerGroupResponseInlineRecords []*models.BgpPeerGroup `json:"records,omitempty"`
}

// Validate validates this network IP bgp peer group delete collection body
func (o *NetworkIPBgpPeerGroupDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBgpPeerGroupResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkIPBgpPeerGroupDeleteCollectionBody) validateBgpPeerGroupResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.BgpPeerGroupResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.BgpPeerGroupResponseInlineRecords); i++ {
		if swag.IsZero(o.BgpPeerGroupResponseInlineRecords[i]) { // not required
			continue
		}

		if o.BgpPeerGroupResponseInlineRecords[i] != nil {
			if err := o.BgpPeerGroupResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network IP bgp peer group delete collection body based on the context it is used
func (o *NetworkIPBgpPeerGroupDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBgpPeerGroupResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkIPBgpPeerGroupDeleteCollectionBody) contextValidateBgpPeerGroupResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.BgpPeerGroupResponseInlineRecords); i++ {

		if o.BgpPeerGroupResponseInlineRecords[i] != nil {
			if err := o.BgpPeerGroupResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *NetworkIPBgpPeerGroupDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkIPBgpPeerGroupDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res NetworkIPBgpPeerGroupDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
