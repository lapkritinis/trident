// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewFileInfoFormDataModifyParams creates a new FileInfoFormDataModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileInfoFormDataModifyParams() *FileInfoFormDataModifyParams {
	return &FileInfoFormDataModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileInfoFormDataModifyParamsWithTimeout creates a new FileInfoFormDataModifyParams object
// with the ability to set a timeout on a request.
func NewFileInfoFormDataModifyParamsWithTimeout(timeout time.Duration) *FileInfoFormDataModifyParams {
	return &FileInfoFormDataModifyParams{
		timeout: timeout,
	}
}

// NewFileInfoFormDataModifyParamsWithContext creates a new FileInfoFormDataModifyParams object
// with the ability to set a context for a request.
func NewFileInfoFormDataModifyParamsWithContext(ctx context.Context) *FileInfoFormDataModifyParams {
	return &FileInfoFormDataModifyParams{
		Context: ctx,
	}
}

// NewFileInfoFormDataModifyParamsWithHTTPClient creates a new FileInfoFormDataModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileInfoFormDataModifyParamsWithHTTPClient(client *http.Client) *FileInfoFormDataModifyParams {
	return &FileInfoFormDataModifyParams{
		HTTPClient: client,
	}
}

/*
FileInfoFormDataModifyParams contains all the parameters to send to the API endpoint

	for the file info form data modify operation.

	Typically these are written to a http.Request.
*/
type FileInfoFormDataModifyParams struct {

	/* ByteOffset.

	   Indicates the number of bytes into the file to begin writing. Use "-1" to append (default). Note that the byte-offset field is only supported for writing to a new or existing file, which requires specifying the Content-Type as 'multipart/form-data'.
	*/
	ByteOffset *int64

	/* Data.

	   Data to write to the file.
	*/
	Data *string

	/* Path.

	   Relative path of a file in the volume. The path field requires using "%2E" to represent "." and "%2F" to represent "/" for the path provided.
	*/
	Path string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* StreamName.

	   Name of stream associated with the file to write data to.
	*/
	StreamName *string

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the file info form data modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileInfoFormDataModifyParams) WithDefaults() *FileInfoFormDataModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file info form data modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileInfoFormDataModifyParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := FileInfoFormDataModifyParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the file info form data modify params
func (o *FileInfoFormDataModifyParams) WithTimeout(timeout time.Duration) *FileInfoFormDataModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file info form data modify params
func (o *FileInfoFormDataModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file info form data modify params
func (o *FileInfoFormDataModifyParams) WithContext(ctx context.Context) *FileInfoFormDataModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file info form data modify params
func (o *FileInfoFormDataModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file info form data modify params
func (o *FileInfoFormDataModifyParams) WithHTTPClient(client *http.Client) *FileInfoFormDataModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file info form data modify params
func (o *FileInfoFormDataModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithByteOffset adds the byteOffset to the file info form data modify params
func (o *FileInfoFormDataModifyParams) WithByteOffset(byteOffset *int64) *FileInfoFormDataModifyParams {
	o.SetByteOffset(byteOffset)
	return o
}

// SetByteOffset adds the byteOffset to the file info form data modify params
func (o *FileInfoFormDataModifyParams) SetByteOffset(byteOffset *int64) {
	o.ByteOffset = byteOffset
}

// WithData adds the data to the file info form data modify params
func (o *FileInfoFormDataModifyParams) WithData(data *string) *FileInfoFormDataModifyParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the file info form data modify params
func (o *FileInfoFormDataModifyParams) SetData(data *string) {
	o.Data = data
}

// WithPath adds the path to the file info form data modify params
func (o *FileInfoFormDataModifyParams) WithPath(path string) *FileInfoFormDataModifyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the file info form data modify params
func (o *FileInfoFormDataModifyParams) SetPath(path string) {
	o.Path = path
}

// WithReturnRecords adds the returnRecords to the file info form data modify params
func (o *FileInfoFormDataModifyParams) WithReturnRecords(returnRecords *bool) *FileInfoFormDataModifyParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the file info form data modify params
func (o *FileInfoFormDataModifyParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithStreamName adds the streamName to the file info form data modify params
func (o *FileInfoFormDataModifyParams) WithStreamName(streamName *string) *FileInfoFormDataModifyParams {
	o.SetStreamName(streamName)
	return o
}

// SetStreamName adds the streamName to the file info form data modify params
func (o *FileInfoFormDataModifyParams) SetStreamName(streamName *string) {
	o.StreamName = streamName
}

// WithVolumeUUID adds the volumeUUID to the file info form data modify params
func (o *FileInfoFormDataModifyParams) WithVolumeUUID(volumeUUID string) *FileInfoFormDataModifyParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the file info form data modify params
func (o *FileInfoFormDataModifyParams) SetVolumeUUID(volumeUUID string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FileInfoFormDataModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ByteOffset != nil {

		// query param byte_offset
		var qrByteOffset int64

		if o.ByteOffset != nil {
			qrByteOffset = *o.ByteOffset
		}
		qByteOffset := swag.FormatInt64(qrByteOffset)
		if qByteOffset != "" {

			if err := r.SetQueryParam("byte_offset", qByteOffset); err != nil {
				return err
			}
		}
	}

	if o.Data != nil {

		// form param data
		var frData string
		if o.Data != nil {
			frData = *o.Data
		}
		fData := frData
		if fData != "" {
			if err := r.SetFormParam("data", fData); err != nil {
				return err
			}
		}
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.StreamName != nil {

		// query param stream_name
		var qrStreamName string

		if o.StreamName != nil {
			qrStreamName = *o.StreamName
		}
		qStreamName := qrStreamName
		if qStreamName != "" {

			if err := r.SetQueryParam("stream_name", qStreamName); err != nil {
				return err
			}
		}
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
