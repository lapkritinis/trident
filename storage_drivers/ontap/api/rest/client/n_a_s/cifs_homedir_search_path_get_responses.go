// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// CifsHomedirSearchPathGetReader is a Reader for the CifsHomedirSearchPathGet structure.
type CifsHomedirSearchPathGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsHomedirSearchPathGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsHomedirSearchPathGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsHomedirSearchPathGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsHomedirSearchPathGetOK creates a CifsHomedirSearchPathGetOK with default headers values
func NewCifsHomedirSearchPathGetOK() *CifsHomedirSearchPathGetOK {
	return &CifsHomedirSearchPathGetOK{}
}

/*
CifsHomedirSearchPathGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsHomedirSearchPathGetOK struct {
	Payload *models.CifsSearchPath
}

// IsSuccess returns true when this cifs homedir search path get o k response has a 2xx status code
func (o *CifsHomedirSearchPathGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs homedir search path get o k response has a 3xx status code
func (o *CifsHomedirSearchPathGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs homedir search path get o k response has a 4xx status code
func (o *CifsHomedirSearchPathGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs homedir search path get o k response has a 5xx status code
func (o *CifsHomedirSearchPathGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs homedir search path get o k response a status code equal to that given
func (o *CifsHomedirSearchPathGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs homedir search path get o k response
func (o *CifsHomedirSearchPathGetOK) Code() int {
	return 200
}

func (o *CifsHomedirSearchPathGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifsHomedirSearchPathGetOK %s", 200, payload)
}

func (o *CifsHomedirSearchPathGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifsHomedirSearchPathGetOK %s", 200, payload)
}

func (o *CifsHomedirSearchPathGetOK) GetPayload() *models.CifsSearchPath {
	return o.Payload
}

func (o *CifsHomedirSearchPathGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsSearchPath)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsHomedirSearchPathGetDefault creates a CifsHomedirSearchPathGetDefault with default headers values
func NewCifsHomedirSearchPathGetDefault(code int) *CifsHomedirSearchPathGetDefault {
	return &CifsHomedirSearchPathGetDefault{
		_statusCode: code,
	}
}

/*
CifsHomedirSearchPathGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsHomedirSearchPathGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs homedir search path get default response has a 2xx status code
func (o *CifsHomedirSearchPathGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs homedir search path get default response has a 3xx status code
func (o *CifsHomedirSearchPathGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs homedir search path get default response has a 4xx status code
func (o *CifsHomedirSearchPathGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs homedir search path get default response has a 5xx status code
func (o *CifsHomedirSearchPathGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs homedir search path get default response a status code equal to that given
func (o *CifsHomedirSearchPathGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs homedir search path get default response
func (o *CifsHomedirSearchPathGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsHomedirSearchPathGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifs_homedir_search_path_get default %s", o._statusCode, payload)
}

func (o *CifsHomedirSearchPathGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifs_homedir_search_path_get default %s", o._statusCode, payload)
}

func (o *CifsHomedirSearchPathGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsHomedirSearchPathGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
