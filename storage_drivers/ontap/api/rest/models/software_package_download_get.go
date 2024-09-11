// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SoftwarePackageDownloadGet software package download get
//
// swagger:model software_package_download_get
type SoftwarePackageDownloadGet struct {

	// Code returned corresponds to a download message
	// Example: 10551382
	Code *int64 `json:"code,omitempty"`

	// Download progress details
	// Example: Package download in progress
	Message *string `json:"message,omitempty"`

	// Download status of the package
	// Example: success
	// Enum: ["not_started","running","success","failure"]
	State *string `json:"state,omitempty"`
}

// Validate validates this software package download get
func (m *SoftwarePackageDownloadGet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var softwarePackageDownloadGetTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_started","running","success","failure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwarePackageDownloadGetTypeStatePropEnum = append(softwarePackageDownloadGetTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// software_package_download_get
	// SoftwarePackageDownloadGet
	// state
	// State
	// not_started
	// END DEBUGGING
	// SoftwarePackageDownloadGetStateNotStarted captures enum value "not_started"
	SoftwarePackageDownloadGetStateNotStarted string = "not_started"

	// BEGIN DEBUGGING
	// software_package_download_get
	// SoftwarePackageDownloadGet
	// state
	// State
	// running
	// END DEBUGGING
	// SoftwarePackageDownloadGetStateRunning captures enum value "running"
	SoftwarePackageDownloadGetStateRunning string = "running"

	// BEGIN DEBUGGING
	// software_package_download_get
	// SoftwarePackageDownloadGet
	// state
	// State
	// success
	// END DEBUGGING
	// SoftwarePackageDownloadGetStateSuccess captures enum value "success"
	SoftwarePackageDownloadGetStateSuccess string = "success"

	// BEGIN DEBUGGING
	// software_package_download_get
	// SoftwarePackageDownloadGet
	// state
	// State
	// failure
	// END DEBUGGING
	// SoftwarePackageDownloadGetStateFailure captures enum value "failure"
	SoftwarePackageDownloadGetStateFailure string = "failure"
)

// prop value enum
func (m *SoftwarePackageDownloadGet) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, softwarePackageDownloadGetTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SoftwarePackageDownloadGet) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this software package download get based on context it is used
func (m *SoftwarePackageDownloadGet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwarePackageDownloadGet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwarePackageDownloadGet) UnmarshalBinary(b []byte) error {
	var res SoftwarePackageDownloadGet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
