// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// CloudTargetModifyReader is a Reader for the CloudTargetModify structure.
type CloudTargetModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudTargetModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudTargetModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCloudTargetModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCloudTargetModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudTargetModifyOK creates a CloudTargetModifyOK with default headers values
func NewCloudTargetModifyOK() *CloudTargetModifyOK {
	return &CloudTargetModifyOK{}
}

/*
CloudTargetModifyOK describes a response with status code 200, with default header values.

OK
*/
type CloudTargetModifyOK struct {
	Payload *models.CloudTargetJobLinkResponse
}

// IsSuccess returns true when this cloud target modify o k response has a 2xx status code
func (o *CloudTargetModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud target modify o k response has a 3xx status code
func (o *CloudTargetModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud target modify o k response has a 4xx status code
func (o *CloudTargetModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud target modify o k response has a 5xx status code
func (o *CloudTargetModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud target modify o k response a status code equal to that given
func (o *CloudTargetModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cloud target modify o k response
func (o *CloudTargetModifyOK) Code() int {
	return 200
}

func (o *CloudTargetModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cloud/targets/{uuid}][%d] cloudTargetModifyOK %s", 200, payload)
}

func (o *CloudTargetModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cloud/targets/{uuid}][%d] cloudTargetModifyOK %s", 200, payload)
}

func (o *CloudTargetModifyOK) GetPayload() *models.CloudTargetJobLinkResponse {
	return o.Payload
}

func (o *CloudTargetModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudTargetJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudTargetModifyAccepted creates a CloudTargetModifyAccepted with default headers values
func NewCloudTargetModifyAccepted() *CloudTargetModifyAccepted {
	return &CloudTargetModifyAccepted{}
}

/*
CloudTargetModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CloudTargetModifyAccepted struct {
	Payload *models.CloudTargetJobLinkResponse
}

// IsSuccess returns true when this cloud target modify accepted response has a 2xx status code
func (o *CloudTargetModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud target modify accepted response has a 3xx status code
func (o *CloudTargetModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud target modify accepted response has a 4xx status code
func (o *CloudTargetModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud target modify accepted response has a 5xx status code
func (o *CloudTargetModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud target modify accepted response a status code equal to that given
func (o *CloudTargetModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cloud target modify accepted response
func (o *CloudTargetModifyAccepted) Code() int {
	return 202
}

func (o *CloudTargetModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cloud/targets/{uuid}][%d] cloudTargetModifyAccepted %s", 202, payload)
}

func (o *CloudTargetModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cloud/targets/{uuid}][%d] cloudTargetModifyAccepted %s", 202, payload)
}

func (o *CloudTargetModifyAccepted) GetPayload() *models.CloudTargetJobLinkResponse {
	return o.Payload
}

func (o *CloudTargetModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudTargetJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudTargetModifyDefault creates a CloudTargetModifyDefault with default headers values
func NewCloudTargetModifyDefault(code int) *CloudTargetModifyDefault {
	return &CloudTargetModifyDefault{
		_statusCode: code,
	}
}

/*
	CloudTargetModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 786436 | VLDB is not running. |
| 786908 | Capability check failed. |
| 787016 | An object store configuration with the same combination of server and container name already exists. |
| 787036 | Server name is invalid. A valid server name must be a fully qualified domain name. |
| 787038 | Object store provider type requires a FabricPool license. |
| 787039 | Failed to retrieve FabricPool capacity license information. |
| 787054 | An object store configuration with the same combination of server, azure account and container name already exists. |
| 787064 | Object store server provider type does not match object store provider type. Use the provider type that matches the server. |
| 787065 | Certificate validation must be enabled for the object store provider. |
| 787066 | Certificate validation cannot be specified when the \"-is-ssl-enabled\" parameter is false. |
| 787068 | Disabling certificate validation requires an effective cluster version of ONTAP 9.4.0 or later. |
| 787089 | The object store provider supports the virtual hosted-style, and the specified \"-server\" contains the container name.  The container specified in the \"-server\" parameter must be the same as the name of the container specified in the \"-container\" parameter. |
| 787133 | Could not retrieve temporary credentials from the CAP server. |
| 787134 | Could not retrieve temporary credentials from the CAP server due to an unexpected response. |
| 787136 | Specifying \"CAP\" as the \"-auth-type\" requires an effective cluster version of ONTAP 9.5.0 or later. |
| 787148 | The clock on the node is behind by more than the maximum of 5 minutes. |
| 787149 | The clock on the node is ahead by more than the maximum of 5 minutes. |
| 787184 | Using HTTP proxies with FabricPool requires an effective cluster version of ONTAP 9.7.0 or later. |
| 787185 | There is no HTTP proxy for IPspace. Refer to the \"vserver http-proxy\" man page for details. |
| 787188 | Object store configuration belongs to cluster and cannot be modified from the local cluster, unless the cluster is in switchover mode. |
| 787209 | Object store is not accessible from the partner cluster in a MetroCluster configuration. |
| 787216 | Cannot perform object store configuration operations on a cluster that is waiting for switchback. |
| 787223 | Specifying \"GCP_SA\" as the \"-auth-type\" requires an effective cluster version of ONTAP 9.7.0 or later. |
| 787227 | Specifying \"Azure_MSI\" as the \"-auth-type\" requires an effective cluster version of ONTAP 9.7 or later. |
| 787228 | SSL is required for this configuration. |
| 787229 | Cannot perform operation as URL style is not supported with object store provider type. |
| 787233 | The hash key for enabling this FabricPool feature is not present on the cluster. |
| 787234 | The hash key provided for the node to enable this FabricPool feature is not valid. |
| 787254 | The parameter is not supported on this system. |
| 787257 | The parameter \"-encryption-context\" is only applicable for AWS object store provider. |
| 787301 | ONTAP S3 Bucket is of type NAS and is not supported as a target for FabricPool. |
| 787302 | Cannot use HTTP port with \"-is-ssl-enabled\" set to true. |
| 787303 | Cannot use HTTPS port with \"-is-ssl-enabled\" set to false. |
| 787306 | Object store is not accessible from the partner cluster in a MetroCluster configuration. |
| 787350 | Modifying an object store configuration with a Managed Service Identity (MSI) token is only supported on Azure NetApp Files. |
| 787351 | Internal Error. Invalid authentication type. |
| 787352 | Modifying an object store configuration with a Managed Service Identity (MSI) token requires an effective cluster version of ONTAP 9.16.1 or later. |
| 787353 | Modifying an object store configuration with a Managed Service Identity (MSI) token is not supported for this owner. |
| 787354 | The specified properties are mutually exclusive. |
| 787355 | _azure_msi_ `authentication-type` is only supported with _Azure_Cloud_ `provider-type`. |
| 45940761 | Hostname cannot be resolved. Check the spelling of the hostname and check the system DNS availability using the \"vserver services name-service dns check\" command. |
| 45940778 | Bucket already exists. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type CloudTargetModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cloud target modify default response has a 2xx status code
func (o *CloudTargetModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cloud target modify default response has a 3xx status code
func (o *CloudTargetModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cloud target modify default response has a 4xx status code
func (o *CloudTargetModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cloud target modify default response has a 5xx status code
func (o *CloudTargetModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cloud target modify default response a status code equal to that given
func (o *CloudTargetModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cloud target modify default response
func (o *CloudTargetModifyDefault) Code() int {
	return o._statusCode
}

func (o *CloudTargetModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cloud/targets/{uuid}][%d] cloud_target_modify default %s", o._statusCode, payload)
}

func (o *CloudTargetModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cloud/targets/{uuid}][%d] cloud_target_modify default %s", o._statusCode, payload)
}

func (o *CloudTargetModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CloudTargetModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
