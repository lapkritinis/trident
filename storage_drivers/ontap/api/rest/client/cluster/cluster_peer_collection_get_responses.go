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

// ClusterPeerCollectionGetReader is a Reader for the ClusterPeerCollectionGet structure.
type ClusterPeerCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterPeerCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterPeerCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterPeerCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterPeerCollectionGetOK creates a ClusterPeerCollectionGetOK with default headers values
func NewClusterPeerCollectionGetOK() *ClusterPeerCollectionGetOK {
	return &ClusterPeerCollectionGetOK{}
}

/*
ClusterPeerCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterPeerCollectionGetOK struct {
	Payload *models.ClusterPeerResponse
}

// IsSuccess returns true when this cluster peer collection get o k response has a 2xx status code
func (o *ClusterPeerCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster peer collection get o k response has a 3xx status code
func (o *ClusterPeerCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster peer collection get o k response has a 4xx status code
func (o *ClusterPeerCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster peer collection get o k response has a 5xx status code
func (o *ClusterPeerCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster peer collection get o k response a status code equal to that given
func (o *ClusterPeerCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster peer collection get o k response
func (o *ClusterPeerCollectionGetOK) Code() int {
	return 200
}

func (o *ClusterPeerCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/peers][%d] clusterPeerCollectionGetOK %s", 200, payload)
}

func (o *ClusterPeerCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/peers][%d] clusterPeerCollectionGetOK %s", 200, payload)
}

func (o *ClusterPeerCollectionGetOK) GetPayload() *models.ClusterPeerResponse {
	return o.Payload
}

func (o *ClusterPeerCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterPeerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterPeerCollectionGetDefault creates a ClusterPeerCollectionGetDefault with default headers values
func NewClusterPeerCollectionGetDefault(code int) *ClusterPeerCollectionGetDefault {
	return &ClusterPeerCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ClusterPeerCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterPeerCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster peer collection get default response has a 2xx status code
func (o *ClusterPeerCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster peer collection get default response has a 3xx status code
func (o *ClusterPeerCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster peer collection get default response has a 4xx status code
func (o *ClusterPeerCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster peer collection get default response has a 5xx status code
func (o *ClusterPeerCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster peer collection get default response a status code equal to that given
func (o *ClusterPeerCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster peer collection get default response
func (o *ClusterPeerCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ClusterPeerCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/peers][%d] cluster_peer_collection_get default %s", o._statusCode, payload)
}

func (o *ClusterPeerCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/peers][%d] cluster_peer_collection_get default %s", o._statusCode, payload)
}

func (o *ClusterPeerCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterPeerCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
