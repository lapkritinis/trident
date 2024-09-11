// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ClusterPeerModifyReader is a Reader for the ClusterPeerModify structure.
type ClusterPeerModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterPeerModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterPeerModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterPeerModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterPeerModifyOK creates a ClusterPeerModifyOK with default headers values
func NewClusterPeerModifyOK() *ClusterPeerModifyOK {
	return &ClusterPeerModifyOK{}
}

/*
ClusterPeerModifyOK describes a response with status code 200, with default header values.

OK
*/
type ClusterPeerModifyOK struct {
	Payload *models.ClusterPeer
}

// IsSuccess returns true when this cluster peer modify o k response has a 2xx status code
func (o *ClusterPeerModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster peer modify o k response has a 3xx status code
func (o *ClusterPeerModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster peer modify o k response has a 4xx status code
func (o *ClusterPeerModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster peer modify o k response has a 5xx status code
func (o *ClusterPeerModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster peer modify o k response a status code equal to that given
func (o *ClusterPeerModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster peer modify o k response
func (o *ClusterPeerModifyOK) Code() int {
	return 200
}

func (o *ClusterPeerModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/peers/{uuid}][%d] clusterPeerModifyOK %s", 200, payload)
}

func (o *ClusterPeerModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/peers/{uuid}][%d] clusterPeerModifyOK %s", 200, payload)
}

func (o *ClusterPeerModifyOK) GetPayload() *models.ClusterPeer {
	return o.Payload
}

func (o *ClusterPeerModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterPeer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterPeerModifyDefault creates a ClusterPeerModifyDefault with default headers values
func NewClusterPeerModifyDefault(code int) *ClusterPeerModifyDefault {
	return &ClusterPeerModifyDefault{
		_statusCode: code,
	}
}

/*
	ClusterPeerModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | SVM name does not exist. |
| 4653061 | The specified remote cluster is invalid. |
| 4653218 | An introductory RPC to the peer address failed to connect. Verify that the peer address is correct, and then try the operation again. |
| 4653229 | Specified value for \"-offer-expiration\" is obsolete. |
| 4653236 | The specified passphrase is too short. |
| 4653257 | The vifmgr process is not running. |
| 4653261 | Error finding IPspace. |
| 4653671 | No operational intercluster LIFs of the IPv4 address family is available on this node for the specified IPspace. |
| 4655058 | Expiration time cannot be more than 7 days in the future. |
| 4655061 | SVM does not exist in the IPspace. |
| 4656070 | The encryption protocol is meaningful only with authenticated cluster peer relationships. |
| 4656072 | The name must conform to the same rules as a cluster name. |
| 4656073 | Changing the encryption state requires the refreshing of the authentication passphrase. |
| 4656075 | Cannot specify encryption: this operation requires an ECV of ONTAP 9.6.0 or later. |
| 4656076 | Cluster peer modify was attempted with mismatched IPv4 and IPv6 addresses. |
| 4656080 | Specify either a passphrase or set \"generate-passphrase\" to true. |
| 4656081 | The remote IP address list is empty. |
| 4656083 | Cannot auto-generate a passphrase when \"generate-passphrase\" is false. Modifying a passphrase using an auto-generated passphrase requires \"generate-passphrase\" be true. |
| 4656084 | Passphrase can only be modified with an authenticated cluster peer relationship. |
| 4656092 | Cluster peer modify was attempted with a host name that did not resolve to an IPv4 or IPv6 address. |
| 4656095 | The address family of the specified peer addresses is not valid in this IPspace. Use /api/network/interfaces/ to verify that required LIFs are present and operational on each cluster node. |
| 8847365 | Unknown Host |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ClusterPeerModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster peer modify default response has a 2xx status code
func (o *ClusterPeerModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster peer modify default response has a 3xx status code
func (o *ClusterPeerModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster peer modify default response has a 4xx status code
func (o *ClusterPeerModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster peer modify default response has a 5xx status code
func (o *ClusterPeerModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster peer modify default response a status code equal to that given
func (o *ClusterPeerModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster peer modify default response
func (o *ClusterPeerModifyDefault) Code() int {
	return o._statusCode
}

func (o *ClusterPeerModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/peers/{uuid}][%d] cluster_peer_modify default %s", o._statusCode, payload)
}

func (o *ClusterPeerModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/peers/{uuid}][%d] cluster_peer_modify default %s", o._statusCode, payload)
}

func (o *ClusterPeerModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterPeerModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
