// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NvmeNamespaceCreateReader is a Reader for the NvmeNamespaceCreate structure.
type NvmeNamespaceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeNamespaceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNvmeNamespaceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewNvmeNamespaceCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeNamespaceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeNamespaceCreateCreated creates a NvmeNamespaceCreateCreated with default headers values
func NewNvmeNamespaceCreateCreated() *NvmeNamespaceCreateCreated {
	return &NvmeNamespaceCreateCreated{}
}

/*
NvmeNamespaceCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NvmeNamespaceCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.NvmeNamespaceResponse
}

// IsSuccess returns true when this nvme namespace create created response has a 2xx status code
func (o *NvmeNamespaceCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme namespace create created response has a 3xx status code
func (o *NvmeNamespaceCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme namespace create created response has a 4xx status code
func (o *NvmeNamespaceCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme namespace create created response has a 5xx status code
func (o *NvmeNamespaceCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme namespace create created response a status code equal to that given
func (o *NvmeNamespaceCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the nvme namespace create created response
func (o *NvmeNamespaceCreateCreated) Code() int {
	return 201
}

func (o *NvmeNamespaceCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/namespaces][%d] nvmeNamespaceCreateCreated %s", 201, payload)
}

func (o *NvmeNamespaceCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/namespaces][%d] nvmeNamespaceCreateCreated %s", 201, payload)
}

func (o *NvmeNamespaceCreateCreated) GetPayload() *models.NvmeNamespaceResponse {
	return o.Payload
}

func (o *NvmeNamespaceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.NvmeNamespaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeNamespaceCreateAccepted creates a NvmeNamespaceCreateAccepted with default headers values
func NewNvmeNamespaceCreateAccepted() *NvmeNamespaceCreateAccepted {
	return &NvmeNamespaceCreateAccepted{}
}

/*
NvmeNamespaceCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type NvmeNamespaceCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.NvmeNamespaceJobLinkResponse
}

// IsSuccess returns true when this nvme namespace create accepted response has a 2xx status code
func (o *NvmeNamespaceCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme namespace create accepted response has a 3xx status code
func (o *NvmeNamespaceCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme namespace create accepted response has a 4xx status code
func (o *NvmeNamespaceCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme namespace create accepted response has a 5xx status code
func (o *NvmeNamespaceCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme namespace create accepted response a status code equal to that given
func (o *NvmeNamespaceCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the nvme namespace create accepted response
func (o *NvmeNamespaceCreateAccepted) Code() int {
	return 202
}

func (o *NvmeNamespaceCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/namespaces][%d] nvmeNamespaceCreateAccepted %s", 202, payload)
}

func (o *NvmeNamespaceCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/namespaces][%d] nvmeNamespaceCreateAccepted %s", 202, payload)
}

func (o *NvmeNamespaceCreateAccepted) GetPayload() *models.NvmeNamespaceJobLinkResponse {
	return o.Payload
}

func (o *NvmeNamespaceCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.NvmeNamespaceJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeNamespaceCreateDefault creates a NvmeNamespaceCreateDefault with default headers values
func NewNvmeNamespaceCreateDefault(code int) *NvmeNamespaceCreateDefault {
	return &NvmeNamespaceCreateDefault{
		_statusCode: code,
	}
}

/*
	NvmeNamespaceCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 917927 | The specified volume was not found. |
| 918236 | The specified `location.volume.uuid` and `location.volume.name` do not refer to the same volume. |
| 1254197 | The LUN specified for conversion to a namespace is mapped. |
| 1260121 | Cloning a namespace to a volume different than the source volume is not supported. |
| 1260136 | The specified destination for a clone operation already exists as a LUN, namespace, or file. |
| 2621462 | The supplied SVM does not exist. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5242927 | The specified qtree was not found. |
| 5242950 | The specified `location.qtree.id` and `location.qtree.name` do not refer to the same qtree. |
| 5374127 | The specified namespace name is invalid. |
| 5374140 | LUN has a non-zero prefix and/or suffix size. |
| 5374141 | LUN is part of a SnapMirror active sync relationship. |
| 5374156 | A protocol endpoint LUN cannot be converted to an NVMe namespace. |
| 5374157 | LUN in an SVM with MetroCluster configured cannot be converted to an NVMe namespace. |
| 5374158 | LUN contains an operating system type that is not supported for NVMe namespace. |
| 5374352 | An invalid name was provided for the NVMe namespace. |
| 5374858 | The volume specified by `name` is not the same as that specified by `location.volume`. |
| 5374860 | The qtree specified by `name` is not the same as that specified by `location.qtree`. |
| 5374861 | The NVME namespace base name specified by `name` is not the same as that specified by `location.name`. |
| 5374862 | No NVMe namespace path base name was provided for the namespace. |
| 5374876 | The LUN specified for conversion to a namespace was not found. |
| 5376461 | The specified namespace name is invalid. |
| 5376462 | The specified namespace name is too long. |
| 5376463 | The snapshot portion of the specified namespace name is too long. |
| 13565952 | The NVMe namespace clone request failed. |
| 72089636 | Creating NVMe namespaces with `os_type` AIX is not supported until the effective cluster version is 9.13.1 or later. |
| 72089720 | NVMe namespaces cannot be created in snapshots. |
| 72089721 | The volume specified is in a load sharing mirror relationship. Namespaces are not supported in load sharing mirrors. |
| 72089722 | A negative size was provided for the NVMe namespace. |
| 72089723 | The specified size is too small for the NVMe namespace. |
| 72089724 | The specified size is too large for the NVMe namespace. |
| 72089725 | A LUN or NVMe namespace already exists at the specified path. |
| 72089727 | NVMe namespaces cannot be created on an SVM root volume. |
| 72089728 | NVMe namespaces cannot be created on a FlexGroup volume. |
| 72089732 | An NVMe namespace name can only contain characters A-Z, a-z, 0-9, "-", ".", "_", "{" and "}". |
| 72090005 | The specified `clone.source.uuid` and `clone.source.name` do not refer to the same NVMe namespace. |
| 72090006 | The specified `clone.source` was not found. |
| 72090007 | The specified `clone.source` was not found. |
| 72090009 | An error occurred after successfully creating the NVMe namespace. Some properties were not set. |
| 72090012 | The property cannot be specified when creating an NVMe namespace clone. The `target` property of the error object identifies the property. |
| 72090013 | The property is required except when creating an NVMe namespace clone. The `target` property of the error object identifies the property. |
| 72090014 | No volume was specified for the NVMe namespace. |
| 72090015 | An error occurred after successfully creating the NVMe namespace preventing the retrieval of its properties. |
| 72090033 | The `clone.source.uuid` property is not supported when specifying a source NVMe namespace from a snapshot. |
| 72090039 | The property cannot be specified at the same time when creating an NVMe namespace as a clone. The `target` property of the error object identifies the other property given with clone. |
| 72090040 | The property cannot be specified when converting a LUN into an NVMe namespace. The `target` property of the error object identifies the property. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeNamespaceCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme namespace create default response has a 2xx status code
func (o *NvmeNamespaceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme namespace create default response has a 3xx status code
func (o *NvmeNamespaceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme namespace create default response has a 4xx status code
func (o *NvmeNamespaceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme namespace create default response has a 5xx status code
func (o *NvmeNamespaceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme namespace create default response a status code equal to that given
func (o *NvmeNamespaceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme namespace create default response
func (o *NvmeNamespaceCreateDefault) Code() int {
	return o._statusCode
}

func (o *NvmeNamespaceCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/namespaces][%d] nvme_namespace_create default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/namespaces][%d] nvme_namespace_create default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeNamespaceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
