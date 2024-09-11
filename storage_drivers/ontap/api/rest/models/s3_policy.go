// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// S3Policy An S3 policy is an object. It defines resource (bucket, folder or object) permissions. These policies get evaluated when an object store user user makes a request. Permissions in the policies determine whether the request is allowed or denied.
//
// swagger:model s3_policy
type S3Policy struct {

	// Can contain any additional information about the S3 policy.
	// Example: S3 policy.
	// Max Length: 256
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// Specifies the name of the policy. A policy name length can range from 1 to 128 characters and can only contain the following combination of characters 0-9, A-Z, a-z, "_", "+", "=", ",", ".","@", and "-". It cannot be specified in a PATCH method.
	// Example: Policy1
	// Max Length: 128
	// Min Length: 1
	// Pattern: ^[0-9A-Za-z_+=,.@-]{1,128}$
	Name *string `json:"name,omitempty"`

	// Specifies whether or not the s3 policy is read only. This parameter should not be specified in the POST method.
	// Read Only: true
	ReadOnly *bool `json:"read-only,omitempty"`

	// Specifies the policy statements.
	S3PolicyInlineStatements []*S3PolicyStatement `json:"statements,omitempty"`

	// svm
	Svm *S3PolicyInlineSvm `json:"svm,omitempty"`
}

// Validate validates this s3 policy
func (m *S3Policy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3PolicyInlineStatements(formats); err != nil {
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

func (m *S3Policy) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 256); err != nil {
		return err
	}

	return nil
}

func (m *S3Policy) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 128); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[0-9A-Za-z_+=,.@-]{1,128}$`); err != nil {
		return err
	}

	return nil
}

func (m *S3Policy) validateS3PolicyInlineStatements(formats strfmt.Registry) error {
	if swag.IsZero(m.S3PolicyInlineStatements) { // not required
		return nil
	}

	for i := 0; i < len(m.S3PolicyInlineStatements); i++ {
		if swag.IsZero(m.S3PolicyInlineStatements[i]) { // not required
			continue
		}

		if m.S3PolicyInlineStatements[i] != nil {
			if err := m.S3PolicyInlineStatements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *S3Policy) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 policy based on the context it is used
func (m *S3Policy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReadOnly(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3PolicyInlineStatements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3Policy) contextValidateReadOnly(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "read-only", "body", m.ReadOnly); err != nil {
		return err
	}

	return nil
}

func (m *S3Policy) contextValidateS3PolicyInlineStatements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.S3PolicyInlineStatements); i++ {

		if m.S3PolicyInlineStatements[i] != nil {
			if err := m.S3PolicyInlineStatements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *S3Policy) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *S3Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3Policy) UnmarshalBinary(b []byte) error {
	var res S3Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3PolicyInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model s3_policy_inline_svm
type S3PolicyInlineSvm struct {

	// links
	Links *S3PolicyInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this s3 policy inline svm
func (m *S3PolicyInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3PolicyInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 policy inline svm based on the context it is used
func (m *S3PolicyInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3PolicyInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3PolicyInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3PolicyInlineSvm) UnmarshalBinary(b []byte) error {
	var res S3PolicyInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3PolicyInlineSvmInlineLinks s3 policy inline svm inline links
//
// swagger:model s3_policy_inline_svm_inline__links
type S3PolicyInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 policy inline svm inline links
func (m *S3PolicyInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3PolicyInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 policy inline svm inline links based on the context it is used
func (m *S3PolicyInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3PolicyInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3PolicyInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3PolicyInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3PolicyInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
