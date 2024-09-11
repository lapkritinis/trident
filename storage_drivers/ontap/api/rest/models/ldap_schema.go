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

// LdapSchema ldap schema
//
// swagger:model ldap_schema
type LdapSchema struct {

	// links
	Links *LdapSchemaInlineLinks `json:"_links,omitempty"`

	// Comment to associate with the schema.
	// Example: Schema based on Active Directory Services for UNIX (read-only).
	Comment *string `json:"comment,omitempty"`

	// A global schema that can be used by all the SVMs.
	// Example: true
	// Read Only: true
	GlobalSchema *bool `json:"global_schema,omitempty"`

	// The name of the schema being created, modified or deleted.
	// Example: AD-SFU-v1
	// Max Length: 32
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// name mapping
	NameMapping *LdapSchemaNameMapping `json:"name_mapping,omitempty"`

	// owner
	Owner *LdapSchemaInlineOwner `json:"owner,omitempty"`

	// rfc2307
	Rfc2307 *Rfc2307 `json:"rfc2307,omitempty"`

	// rfc2307bis
	Rfc2307bis *Rfc2307bis `json:"rfc2307bis,omitempty"`

	// Scope of the entity. Set to "cluster" for cluster owned objects and to "svm" for SVM owned objects.
	// Read Only: true
	// Enum: ["cluster","svm"]
	Scope *string `json:"scope,omitempty"`

	// template
	Template *LdapSchemaInlineTemplate `json:"template,omitempty"`
}

// Validate validates this ldap schema
func (m *LdapSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRfc2307(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRfc2307bis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchema) validateLinks(formats strfmt.Registry) error {
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

func (m *LdapSchema) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 32); err != nil {
		return err
	}

	return nil
}

func (m *LdapSchema) validateNameMapping(formats strfmt.Registry) error {
	if swag.IsZero(m.NameMapping) { // not required
		return nil
	}

	if m.NameMapping != nil {
		if err := m.NameMapping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name_mapping")
			}
			return err
		}
	}

	return nil
}

func (m *LdapSchema) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *LdapSchema) validateRfc2307(formats strfmt.Registry) error {
	if swag.IsZero(m.Rfc2307) { // not required
		return nil
	}

	if m.Rfc2307 != nil {
		if err := m.Rfc2307.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rfc2307")
			}
			return err
		}
	}

	return nil
}

func (m *LdapSchema) validateRfc2307bis(formats strfmt.Registry) error {
	if swag.IsZero(m.Rfc2307bis) { // not required
		return nil
	}

	if m.Rfc2307bis != nil {
		if err := m.Rfc2307bis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rfc2307bis")
			}
			return err
		}
	}

	return nil
}

var ldapSchemaTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster","svm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapSchemaTypeScopePropEnum = append(ldapSchemaTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ldap_schema
	// LdapSchema
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// LdapSchemaScopeCluster captures enum value "cluster"
	LdapSchemaScopeCluster string = "cluster"

	// BEGIN DEBUGGING
	// ldap_schema
	// LdapSchema
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// LdapSchemaScopeSvm captures enum value "svm"
	LdapSchemaScopeSvm string = "svm"
)

// prop value enum
func (m *LdapSchema) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapSchemaTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapSchema) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *LdapSchema) validateTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap schema based on the context it is used
func (m *LdapSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNameMapping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRfc2307(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRfc2307bis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchema) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LdapSchema) contextValidateGlobalSchema(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "global_schema", "body", m.GlobalSchema); err != nil {
		return err
	}

	return nil
}

func (m *LdapSchema) contextValidateNameMapping(ctx context.Context, formats strfmt.Registry) error {

	if m.NameMapping != nil {
		if err := m.NameMapping.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name_mapping")
			}
			return err
		}
	}

	return nil
}

func (m *LdapSchema) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *LdapSchema) contextValidateRfc2307(ctx context.Context, formats strfmt.Registry) error {

	if m.Rfc2307 != nil {
		if err := m.Rfc2307.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rfc2307")
			}
			return err
		}
	}

	return nil
}

func (m *LdapSchema) contextValidateRfc2307bis(ctx context.Context, formats strfmt.Registry) error {

	if m.Rfc2307bis != nil {
		if err := m.Rfc2307bis.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rfc2307bis")
			}
			return err
		}
	}

	return nil
}

func (m *LdapSchema) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *LdapSchema) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapSchema) UnmarshalBinary(b []byte) error {
	var res LdapSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapSchemaInlineLinks ldap schema inline links
//
// swagger:model ldap_schema_inline__links
type LdapSchemaInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ldap schema inline links
func (m *LdapSchemaInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ldap schema inline links based on the context it is used
func (m *LdapSchemaInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LdapSchemaInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapSchemaInlineLinks) UnmarshalBinary(b []byte) error {
	var res LdapSchemaInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapSchemaInlineOwner SVM, applies only to SVM-scoped objects.
//
// swagger:model ldap_schema_inline_owner
type LdapSchemaInlineOwner struct {

	// links
	Links *LdapSchemaInlineOwnerInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this ldap schema inline owner
func (m *LdapSchemaInlineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineOwner) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap schema inline owner based on the context it is used
func (m *LdapSchemaInlineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapSchemaInlineOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapSchemaInlineOwner) UnmarshalBinary(b []byte) error {
	var res LdapSchemaInlineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapSchemaInlineOwnerInlineLinks ldap schema inline owner inline links
//
// swagger:model ldap_schema_inline_owner_inline__links
type LdapSchemaInlineOwnerInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ldap schema inline owner inline links
func (m *LdapSchemaInlineOwnerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineOwnerInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap schema inline owner inline links based on the context it is used
func (m *LdapSchemaInlineOwnerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineOwnerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapSchemaInlineOwnerInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapSchemaInlineOwnerInlineLinks) UnmarshalBinary(b []byte) error {
	var res LdapSchemaInlineOwnerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapSchemaInlineTemplate The existing schema template you want to copy.
//
// swagger:model ldap_schema_inline_template
type LdapSchemaInlineTemplate struct {

	// links
	Links *LdapSchemaInlineTemplateInlineLinks `json:"_links,omitempty"`

	// The name of the schema.
	// Example: AD-SFU-v1
	// Max Length: 32
	// Min Length: 1
	Name *string `json:"name,omitempty"`
}

// Validate validates this ldap schema inline template
func (m *LdapSchemaInlineTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineTemplate) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *LdapSchemaInlineTemplate) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("template"+"."+"name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("template"+"."+"name", "body", *m.Name, 32); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ldap schema inline template based on the context it is used
func (m *LdapSchemaInlineTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineTemplate) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapSchemaInlineTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapSchemaInlineTemplate) UnmarshalBinary(b []byte) error {
	var res LdapSchemaInlineTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapSchemaInlineTemplateInlineLinks ldap schema inline template inline links
//
// swagger:model ldap_schema_inline_template_inline__links
type LdapSchemaInlineTemplateInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ldap schema inline template inline links
func (m *LdapSchemaInlineTemplateInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineTemplateInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap schema inline template inline links based on the context it is used
func (m *LdapSchemaInlineTemplateInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapSchemaInlineTemplateInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapSchemaInlineTemplateInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapSchemaInlineTemplateInlineLinks) UnmarshalBinary(b []byte) error {
	var res LdapSchemaInlineTemplateInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
