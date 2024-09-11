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

// SoftwareErrorsReference software errors reference
//
// swagger:model software_errors_reference
type SoftwareErrorsReference struct {

	// Error code of message
	// Example: 177
	// Read Only: true
	Code *int64 `json:"code,omitempty"`

	// Error message
	// Example: Giveback of CFO aggregate is vetoed. Action: Use the \"storage failover show-giveback\" command to view detailed veto status information. Correct the vetoed update check. Use the \"storage failover giveback -ofnode \"node1\" command to complete the giveback.
	// Read Only: true
	Message *string `json:"message,omitempty"`

	// Severity of error
	// Example: warning
	// Read Only: true
	// Enum: ["informational","warning","error"]
	Severity *string `json:"severity,omitempty"`
}

// Validate validates this software errors reference
func (m *SoftwareErrorsReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var softwareErrorsReferenceTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["informational","warning","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwareErrorsReferenceTypeSeverityPropEnum = append(softwareErrorsReferenceTypeSeverityPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// software_errors_reference
	// SoftwareErrorsReference
	// severity
	// Severity
	// informational
	// END DEBUGGING
	// SoftwareErrorsReferenceSeverityInformational captures enum value "informational"
	SoftwareErrorsReferenceSeverityInformational string = "informational"

	// BEGIN DEBUGGING
	// software_errors_reference
	// SoftwareErrorsReference
	// severity
	// Severity
	// warning
	// END DEBUGGING
	// SoftwareErrorsReferenceSeverityWarning captures enum value "warning"
	SoftwareErrorsReferenceSeverityWarning string = "warning"

	// BEGIN DEBUGGING
	// software_errors_reference
	// SoftwareErrorsReference
	// severity
	// Severity
	// error
	// END DEBUGGING
	// SoftwareErrorsReferenceSeverityError captures enum value "error"
	SoftwareErrorsReferenceSeverityError string = "error"
)

// prop value enum
func (m *SoftwareErrorsReference) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, softwareErrorsReferenceTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SoftwareErrorsReference) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this software errors reference based on the context it is used
func (m *SoftwareErrorsReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareErrorsReference) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareErrorsReference) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareErrorsReference) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareErrorsReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareErrorsReference) UnmarshalBinary(b []byte) error {
	var res SoftwareErrorsReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
