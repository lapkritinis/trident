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

// ApplicationComponentSnapshot application component snapshot
//
// swagger:model application_component_snapshot
type ApplicationComponentSnapshot struct {

	// links
	Links *ApplicationComponentSnapshotInlineLinks `json:"_links,omitempty"`

	// application
	Application *ApplicationComponentSnapshotInlineApplication `json:"application,omitempty"`

	// Comment. Valid in POST
	// Max Length: 255
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// component
	Component *ApplicationComponentSnapshotInlineComponent `json:"component,omitempty"`

	// Consistency Type. This is for categorization only. A snapshot should not be set to application consistent unless the host application is quiesced for the snapshot. Valid in POST
	// Enum: ["crash","application"]
	ConsistencyType *string `json:"consistency_type,omitempty"`

	// Creation Time
	// Read Only: true
	CreateTime *string `json:"create_time,omitempty"`

	// A partial snapshot means that not all volumes in an application component were included in the snapshot.
	// Read Only: true
	IsPartial *bool `json:"is_partial,omitempty"`

	// Snapshot name. Valid in POST
	Name *string `json:"name,omitempty"`

	// svm
	Svm *ApplicationComponentSnapshotInlineSvm `json:"svm,omitempty"`

	// Snapshot UUID. Valid in URL
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this application component snapshot
func (m *ApplicationComponentSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationComponentSnapshot) validateLinks(formats strfmt.Registry) error {
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

func (m *ApplicationComponentSnapshot) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationComponentSnapshot) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 255); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationComponentSnapshot) validateComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

var applicationComponentSnapshotTypeConsistencyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["crash","application"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationComponentSnapshotTypeConsistencyTypePropEnum = append(applicationComponentSnapshotTypeConsistencyTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// application_component_snapshot
	// ApplicationComponentSnapshot
	// consistency_type
	// ConsistencyType
	// crash
	// END DEBUGGING
	// ApplicationComponentSnapshotConsistencyTypeCrash captures enum value "crash"
	ApplicationComponentSnapshotConsistencyTypeCrash string = "crash"

	// BEGIN DEBUGGING
	// application_component_snapshot
	// ApplicationComponentSnapshot
	// consistency_type
	// ConsistencyType
	// application
	// END DEBUGGING
	// ApplicationComponentSnapshotConsistencyTypeApplication captures enum value "application"
	ApplicationComponentSnapshotConsistencyTypeApplication string = "application"
)

// prop value enum
func (m *ApplicationComponentSnapshot) validateConsistencyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationComponentSnapshotTypeConsistencyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationComponentSnapshot) validateConsistencyType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateConsistencyTypeEnum("consistency_type", "body", *m.ConsistencyType); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationComponentSnapshot) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application component snapshot based on the context it is used
func (m *ApplicationComponentSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsPartial(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
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

func (m *ApplicationComponentSnapshot) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ApplicationComponentSnapshot) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {
		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationComponentSnapshot) contextValidateComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.Component != nil {
		if err := m.Component.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationComponentSnapshot) contextValidateCreateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "create_time", "body", m.CreateTime); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationComponentSnapshot) contextValidateIsPartial(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_partial", "body", m.IsPartial); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationComponentSnapshot) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationComponentSnapshot) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationComponentSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationComponentSnapshot) UnmarshalBinary(b []byte) error {
	var res ApplicationComponentSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationComponentSnapshotInlineApplication application component snapshot inline application
//
// swagger:model application_component_snapshot_inline_application
type ApplicationComponentSnapshotInlineApplication struct {

	// links
	Links *ApplicationComponentSnapshotInlineApplicationInlineLinks `json:"_links,omitempty"`

	// Application Name
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Application UUID. Valid in URL
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this application component snapshot inline application
func (m *ApplicationComponentSnapshotInlineApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationComponentSnapshotInlineApplication) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application component snapshot inline application based on the context it is used
func (m *ApplicationComponentSnapshotInlineApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
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

func (m *ApplicationComponentSnapshotInlineApplication) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationComponentSnapshotInlineApplication) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "application"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationComponentSnapshotInlineApplication) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "application"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineApplication) UnmarshalBinary(b []byte) error {
	var res ApplicationComponentSnapshotInlineApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationComponentSnapshotInlineApplicationInlineLinks application component snapshot inline application inline links
//
// swagger:model application_component_snapshot_inline_application_inline__links
type ApplicationComponentSnapshotInlineApplicationInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this application component snapshot inline application inline links
func (m *ApplicationComponentSnapshotInlineApplicationInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationComponentSnapshotInlineApplicationInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application component snapshot inline application inline links based on the context it is used
func (m *ApplicationComponentSnapshotInlineApplicationInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationComponentSnapshotInlineApplicationInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineApplicationInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineApplicationInlineLinks) UnmarshalBinary(b []byte) error {
	var res ApplicationComponentSnapshotInlineApplicationInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationComponentSnapshotInlineComponent application component snapshot inline component
//
// swagger:model application_component_snapshot_inline_component
type ApplicationComponentSnapshotInlineComponent struct {

	// links
	Links *ApplicationComponentSnapshotInlineComponentInlineLinks `json:"_links,omitempty"`

	// Component Name
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Component UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this application component snapshot inline component
func (m *ApplicationComponentSnapshotInlineComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationComponentSnapshotInlineComponent) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application component snapshot inline component based on the context it is used
func (m *ApplicationComponentSnapshotInlineComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
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

func (m *ApplicationComponentSnapshotInlineComponent) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationComponentSnapshotInlineComponent) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "component"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationComponentSnapshotInlineComponent) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "component"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineComponent) UnmarshalBinary(b []byte) error {
	var res ApplicationComponentSnapshotInlineComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationComponentSnapshotInlineComponentInlineLinks application component snapshot inline component inline links
//
// swagger:model application_component_snapshot_inline_component_inline__links
type ApplicationComponentSnapshotInlineComponentInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this application component snapshot inline component inline links
func (m *ApplicationComponentSnapshotInlineComponentInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationComponentSnapshotInlineComponentInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application component snapshot inline component inline links based on the context it is used
func (m *ApplicationComponentSnapshotInlineComponentInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationComponentSnapshotInlineComponentInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineComponentInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineComponentInlineLinks) UnmarshalBinary(b []byte) error {
	var res ApplicationComponentSnapshotInlineComponentInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationComponentSnapshotInlineLinks application component snapshot inline links
//
// swagger:model application_component_snapshot_inline__links
type ApplicationComponentSnapshotInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this application component snapshot inline links
func (m *ApplicationComponentSnapshotInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationComponentSnapshotInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this application component snapshot inline links based on the context it is used
func (m *ApplicationComponentSnapshotInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationComponentSnapshotInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ApplicationComponentSnapshotInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineLinks) UnmarshalBinary(b []byte) error {
	var res ApplicationComponentSnapshotInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationComponentSnapshotInlineSvm application component snapshot inline svm
//
// swagger:model application_component_snapshot_inline_svm
type ApplicationComponentSnapshotInlineSvm struct {

	// SVM Name
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// SVM UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this application component snapshot inline svm
func (m *ApplicationComponentSnapshotInlineSvm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this application component snapshot inline svm based on the context it is used
func (m *ApplicationComponentSnapshotInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
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

func (m *ApplicationComponentSnapshotInlineSvm) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "svm"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationComponentSnapshotInlineSvm) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "svm"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationComponentSnapshotInlineSvm) UnmarshalBinary(b []byte) error {
	var res ApplicationComponentSnapshotInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
