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

// AutoUpdateConfiguration auto update configuration
//
// swagger:model auto_update_configuration
type AutoUpdateConfiguration struct {

	// links
	Links *AutoUpdateConfigurationInlineLinks `json:"_links,omitempty"`

	// The action to be taken by the alert source as specified by the user.
	// Example: confirm
	// Enum: ["confirm","dismiss","automatic"]
	Action *string `json:"action,omitempty"`

	// Category for the configuration row.
	// Example: disk_fw
	// Read Only: true
	Category *string `json:"category,omitempty"`

	// description
	Description *AutoUpdateConfigurationInlineDescription `json:"description,omitempty"`

	// Unique identifier for the configuration row.
	// Example: 572361f3-e769-439d-9c04-2ba48a08ff47
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this auto update configuration
func (m *AutoUpdateConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateConfiguration) validateLinks(formats strfmt.Registry) error {
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

var autoUpdateConfigurationTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["confirm","dismiss","automatic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoUpdateConfigurationTypeActionPropEnum = append(autoUpdateConfigurationTypeActionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// auto_update_configuration
	// AutoUpdateConfiguration
	// action
	// Action
	// confirm
	// END DEBUGGING
	// AutoUpdateConfigurationActionConfirm captures enum value "confirm"
	AutoUpdateConfigurationActionConfirm string = "confirm"

	// BEGIN DEBUGGING
	// auto_update_configuration
	// AutoUpdateConfiguration
	// action
	// Action
	// dismiss
	// END DEBUGGING
	// AutoUpdateConfigurationActionDismiss captures enum value "dismiss"
	AutoUpdateConfigurationActionDismiss string = "dismiss"

	// BEGIN DEBUGGING
	// auto_update_configuration
	// AutoUpdateConfiguration
	// action
	// Action
	// automatic
	// END DEBUGGING
	// AutoUpdateConfigurationActionAutomatic captures enum value "automatic"
	AutoUpdateConfigurationActionAutomatic string = "automatic"
)

// prop value enum
func (m *AutoUpdateConfiguration) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, autoUpdateConfigurationTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AutoUpdateConfiguration) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *AutoUpdateConfiguration) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if m.Description != nil {
		if err := m.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auto update configuration based on the context it is used
func (m *AutoUpdateConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateConfiguration) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AutoUpdateConfiguration) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *AutoUpdateConfiguration) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if m.Description != nil {
		if err := m.Description.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

func (m *AutoUpdateConfiguration) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoUpdateConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoUpdateConfiguration) UnmarshalBinary(b []byte) error {
	var res AutoUpdateConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutoUpdateConfigurationInlineDescription Description of the configuration row.
// ONTAP Message Codes
// | Code       | Description |
// | ---------- | ----------- |
// | 131072401 | Storage Firmware |
// | 131072402 | SP/BMC Firmware |
// | 131072403 | System Files |
//
// swagger:model auto_update_configuration_inline_description
type AutoUpdateConfigurationInlineDescription struct {

	// Argument code
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Message argument
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this auto update configuration inline description
func (m *AutoUpdateConfigurationInlineDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this auto update configuration inline description based on the context it is used
func (m *AutoUpdateConfigurationInlineDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateConfigurationInlineDescription) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *AutoUpdateConfigurationInlineDescription) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoUpdateConfigurationInlineDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoUpdateConfigurationInlineDescription) UnmarshalBinary(b []byte) error {
	var res AutoUpdateConfigurationInlineDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutoUpdateConfigurationInlineLinks auto update configuration inline links
//
// swagger:model auto_update_configuration_inline__links
type AutoUpdateConfigurationInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this auto update configuration inline links
func (m *AutoUpdateConfigurationInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateConfigurationInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auto update configuration inline links based on the context it is used
func (m *AutoUpdateConfigurationInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateConfigurationInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoUpdateConfigurationInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoUpdateConfigurationInlineLinks) UnmarshalBinary(b []byte) error {
	var res AutoUpdateConfigurationInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
