// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SnapmirrorError SnapMirror error
//
// swagger:model snapmirror_error
type SnapmirrorError struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// Arguments present in the error message encountered.
	SnapmirrorErrorInlineArguments []*string `json:"arguments,omitempty"`
}

// Validate validates this snapmirror error
func (m *SnapmirrorError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this snapmirror error based on context it is used
func (m *SnapmirrorError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorError) UnmarshalBinary(b []byte) error {
	var res SnapmirrorError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
