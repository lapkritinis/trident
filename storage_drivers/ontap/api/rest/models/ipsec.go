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

// Ipsec Manages IPsec configuration via REST APIs.
//
// swagger:model ipsec
type Ipsec struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Indicates whether or not IPsec is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Indicates whether or not IPsec hardware offload is enabled.
	OffloadEnabled *bool `json:"offload_enabled,omitempty"`

	// Replay window size in packets, where 0 indicates that the relay window is disabled.
	// Enum: [0,64,128,256,512,1024]
	ReplayWindow *int64 `json:"replay_window,omitempty"`
}

// Validate validates this ipsec
func (m *Ipsec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplayWindow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ipsec) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

var ipsecTypeReplayWindowPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,64,128,256,512,1024]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipsecTypeReplayWindowPropEnum = append(ipsecTypeReplayWindowPropEnum, v)
	}
}

// prop value enum
func (m *Ipsec) validateReplayWindowEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, ipsecTypeReplayWindowPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Ipsec) validateReplayWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplayWindow) { // not required
		return nil
	}

	// value enum
	if err := m.validateReplayWindowEnum("replay_window", "body", *m.ReplayWindow); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ipsec based on the context it is used
func (m *Ipsec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ipsec) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Ipsec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ipsec) UnmarshalBinary(b []byte) error {
	var res Ipsec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
