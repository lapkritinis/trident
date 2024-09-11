// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnalyticsInfo File system analytics information summarizing all descendants of a specific directory.
//
// swagger:model analytics_info
type AnalyticsInfo struct {

	// by accessed time
	ByAccessedTime *AnalyticsInfoInlineByAccessedTime `json:"by_accessed_time,omitempty"`

	// by modified time
	ByModifiedTime *AnalyticsInfoInlineByModifiedTime `json:"by_modified_time,omitempty"`

	// Number of bytes used on-disk
	// Example: 15436648448
	BytesUsed *int64 `json:"bytes_used,omitempty"`

	// Number of descendants
	// Example: 21134
	FileCount *int64 `json:"file_count,omitempty"`

	// Returns true if data collection is incomplete for this directory tree.
	IncompleteData *bool `json:"incomplete_data,omitempty"`

	// Number of sub directories
	// Example: 35
	SubdirCount *int64 `json:"subdir_count,omitempty"`
}

// Validate validates this analytics info
func (m *AnalyticsInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateByAccessedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateByModifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfo) validateByAccessedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ByAccessedTime) { // not required
		return nil
	}

	if m.ByAccessedTime != nil {
		if err := m.ByAccessedTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_accessed_time")
			}
			return err
		}
	}

	return nil
}

func (m *AnalyticsInfo) validateByModifiedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ByModifiedTime) { // not required
		return nil
	}

	if m.ByModifiedTime != nil {
		if err := m.ByModifiedTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_modified_time")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this analytics info based on the context it is used
func (m *AnalyticsInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateByAccessedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateByModifiedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfo) contextValidateByAccessedTime(ctx context.Context, formats strfmt.Registry) error {

	if m.ByAccessedTime != nil {
		if err := m.ByAccessedTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_accessed_time")
			}
			return err
		}
	}

	return nil
}

func (m *AnalyticsInfo) contextValidateByModifiedTime(ctx context.Context, formats strfmt.Registry) error {

	if m.ByModifiedTime != nil {
		if err := m.ByModifiedTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_modified_time")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsInfo) UnmarshalBinary(b []byte) error {
	var res AnalyticsInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AnalyticsInfoInlineByAccessedTime File system analytics information, broken down by date of last access.
//
// swagger:model analytics_info_inline_by_accessed_time
type AnalyticsInfoInlineByAccessedTime struct {

	// bytes used
	BytesUsed *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed `json:"bytes_used,omitempty"`
}

// Validate validates this analytics info inline by accessed time
func (m *AnalyticsInfoInlineByAccessedTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytesUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfoInlineByAccessedTime) validateBytesUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.BytesUsed) { // not required
		return nil
	}

	if m.BytesUsed != nil {
		if err := m.BytesUsed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_accessed_time" + "." + "bytes_used")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this analytics info inline by accessed time based on the context it is used
func (m *AnalyticsInfoInlineByAccessedTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBytesUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfoInlineByAccessedTime) contextValidateBytesUsed(ctx context.Context, formats strfmt.Registry) error {

	if m.BytesUsed != nil {
		if err := m.BytesUsed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_accessed_time" + "." + "bytes_used")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsInfoInlineByAccessedTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsInfoInlineByAccessedTime) UnmarshalBinary(b []byte) error {
	var res AnalyticsInfoInlineByAccessedTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AnalyticsInfoInlineByAccessedTimeInlineBytesUsed Number of bytes used on-disk, broken down by date of last access.
//
// swagger:model analytics_info_inline_by_accessed_time_inline_bytes_used
type AnalyticsInfoInlineByAccessedTimeInlineBytesUsed struct {

	// labels
	Labels AnalyticsHistogramByTimeLabelsArrayInline `json:"labels,omitempty"`

	// The newest time label with a non-zero histogram value.
	NewestLabel *AnalyticsHistogramTimeLabel `json:"newest_label,omitempty"`

	// The oldest time label with a non-zero histogram value.
	OldestLabel *AnalyticsHistogramTimeLabel `json:"oldest_label,omitempty"`

	// Percentages for this histogram
	// Example: [0.1,11.24,0.18,15.75,0.75,83.5,0]
	Percentages []*float64 `json:"percentages,omitempty"`

	// Values for this histogram
	// Example: [15925248,1735569408,27672576,2430595072,116105216,12889948160,0]
	Values []*int64 `json:"values,omitempty"`
}

// Validate validates this analytics info inline by accessed time inline bytes used
func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewestLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOldestLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("by_accessed_time" + "." + "bytes_used" + "." + "labels")
		}
		return err
	}

	return nil
}

func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) validateNewestLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.NewestLabel) { // not required
		return nil
	}

	if m.NewestLabel != nil {
		if err := m.NewestLabel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_accessed_time" + "." + "bytes_used" + "." + "newest_label")
			}
			return err
		}
	}

	return nil
}

func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) validateOldestLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.OldestLabel) { // not required
		return nil
	}

	if m.OldestLabel != nil {
		if err := m.OldestLabel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_accessed_time" + "." + "bytes_used" + "." + "oldest_label")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this analytics info inline by accessed time inline bytes used based on the context it is used
func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewestLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOldestLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("by_accessed_time" + "." + "bytes_used" + "." + "labels")
		}
		return err
	}

	return nil
}

func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) contextValidateNewestLabel(ctx context.Context, formats strfmt.Registry) error {

	if m.NewestLabel != nil {
		if err := m.NewestLabel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_accessed_time" + "." + "bytes_used" + "." + "newest_label")
			}
			return err
		}
	}

	return nil
}

func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) contextValidateOldestLabel(ctx context.Context, formats strfmt.Registry) error {

	if m.OldestLabel != nil {
		if err := m.OldestLabel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_accessed_time" + "." + "bytes_used" + "." + "oldest_label")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsInfoInlineByAccessedTimeInlineBytesUsed) UnmarshalBinary(b []byte) error {
	var res AnalyticsInfoInlineByAccessedTimeInlineBytesUsed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AnalyticsInfoInlineByModifiedTime File system analytics information, broken down by date of last modification.
//
// swagger:model analytics_info_inline_by_modified_time
type AnalyticsInfoInlineByModifiedTime struct {

	// bytes used
	BytesUsed *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed `json:"bytes_used,omitempty"`
}

// Validate validates this analytics info inline by modified time
func (m *AnalyticsInfoInlineByModifiedTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytesUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfoInlineByModifiedTime) validateBytesUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.BytesUsed) { // not required
		return nil
	}

	if m.BytesUsed != nil {
		if err := m.BytesUsed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_modified_time" + "." + "bytes_used")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this analytics info inline by modified time based on the context it is used
func (m *AnalyticsInfoInlineByModifiedTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBytesUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfoInlineByModifiedTime) contextValidateBytesUsed(ctx context.Context, formats strfmt.Registry) error {

	if m.BytesUsed != nil {
		if err := m.BytesUsed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_modified_time" + "." + "bytes_used")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsInfoInlineByModifiedTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsInfoInlineByModifiedTime) UnmarshalBinary(b []byte) error {
	var res AnalyticsInfoInlineByModifiedTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AnalyticsInfoInlineByModifiedTimeInlineBytesUsed Number of bytes used on-disk, broken down by date of last modification.
//
// swagger:model analytics_info_inline_by_modified_time_inline_bytes_used
type AnalyticsInfoInlineByModifiedTimeInlineBytesUsed struct {

	// labels
	Labels AnalyticsHistogramByTimeLabelsArrayInline `json:"labels,omitempty"`

	// The newest time label with a non-zero histogram value.
	NewestLabel *AnalyticsHistogramTimeLabel `json:"newest_label,omitempty"`

	// The oldest time label with a non-zero histogram value.
	OldestLabel *AnalyticsHistogramTimeLabel `json:"oldest_label,omitempty"`

	// Percentages for this histogram
	// Example: [0.1,11.24,0.18,15.75,0.75,83.5,0]
	Percentages []*float64 `json:"percentages,omitempty"`

	// Values for this histogram
	// Example: [15925248,1735569408,27672576,2430595072,116105216,12889948160,0]
	Values []*int64 `json:"values,omitempty"`
}

// Validate validates this analytics info inline by modified time inline bytes used
func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewestLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOldestLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("by_modified_time" + "." + "bytes_used" + "." + "labels")
		}
		return err
	}

	return nil
}

func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) validateNewestLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.NewestLabel) { // not required
		return nil
	}

	if m.NewestLabel != nil {
		if err := m.NewestLabel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_modified_time" + "." + "bytes_used" + "." + "newest_label")
			}
			return err
		}
	}

	return nil
}

func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) validateOldestLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.OldestLabel) { // not required
		return nil
	}

	if m.OldestLabel != nil {
		if err := m.OldestLabel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_modified_time" + "." + "bytes_used" + "." + "oldest_label")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this analytics info inline by modified time inline bytes used based on the context it is used
func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewestLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOldestLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("by_modified_time" + "." + "bytes_used" + "." + "labels")
		}
		return err
	}

	return nil
}

func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) contextValidateNewestLabel(ctx context.Context, formats strfmt.Registry) error {

	if m.NewestLabel != nil {
		if err := m.NewestLabel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_modified_time" + "." + "bytes_used" + "." + "newest_label")
			}
			return err
		}
	}

	return nil
}

func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) contextValidateOldestLabel(ctx context.Context, formats strfmt.Registry) error {

	if m.OldestLabel != nil {
		if err := m.OldestLabel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("by_modified_time" + "." + "bytes_used" + "." + "oldest_label")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsInfoInlineByModifiedTimeInlineBytesUsed) UnmarshalBinary(b []byte) error {
	var res AnalyticsInfoInlineByModifiedTimeInlineBytesUsed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
