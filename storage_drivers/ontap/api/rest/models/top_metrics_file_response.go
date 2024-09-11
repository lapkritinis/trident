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

// TopMetricsFileResponse top metrics file response
//
// swagger:model top_metrics_file_response
type TopMetricsFileResponse struct {

	// links
	Links *TopMetricsFileResponseInlineLinks `json:"_links,omitempty"`

	// incomplete response reason
	IncompleteResponseReason *TopMetricsFileResponseInlineIncompleteResponseReason `json:"incomplete_response_reason,omitempty"`

	// notice
	Notice *TopMetricsFileResponseInlineNotice `json:"notice,omitempty"`

	// Number of records.
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`

	// top metrics file response inline records
	TopMetricsFileResponseInlineRecords []*TopMetricsFile `json:"records,omitempty"`
}

// Validate validates this top metrics file response
func (m *TopMetricsFileResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncompleteResponseReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopMetricsFileResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsFileResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *TopMetricsFileResponse) validateIncompleteResponseReason(formats strfmt.Registry) error {
	if swag.IsZero(m.IncompleteResponseReason) { // not required
		return nil
	}

	if m.IncompleteResponseReason != nil {
		if err := m.IncompleteResponseReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incomplete_response_reason")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsFileResponse) validateNotice(formats strfmt.Registry) error {
	if swag.IsZero(m.Notice) { // not required
		return nil
	}

	if m.Notice != nil {
		if err := m.Notice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notice")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsFileResponse) validateTopMetricsFileResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.TopMetricsFileResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.TopMetricsFileResponseInlineRecords); i++ {
		if swag.IsZero(m.TopMetricsFileResponseInlineRecords[i]) { // not required
			continue
		}

		if m.TopMetricsFileResponseInlineRecords[i] != nil {
			if err := m.TopMetricsFileResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this top metrics file response based on the context it is used
func (m *TopMetricsFileResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncompleteResponseReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopMetricsFileResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsFileResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsFileResponse) contextValidateIncompleteResponseReason(ctx context.Context, formats strfmt.Registry) error {

	if m.IncompleteResponseReason != nil {
		if err := m.IncompleteResponseReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incomplete_response_reason")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsFileResponse) contextValidateNotice(ctx context.Context, formats strfmt.Registry) error {

	if m.Notice != nil {
		if err := m.Notice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notice")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsFileResponse) contextValidateTopMetricsFileResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TopMetricsFileResponseInlineRecords); i++ {

		if m.TopMetricsFileResponseInlineRecords[i] != nil {
			if err := m.TopMetricsFileResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsFileResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsFileResponse) UnmarshalBinary(b []byte) error {
	var res TopMetricsFileResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsFileResponseInlineIncompleteResponseReason Indicates that the metric report provides incomplete data.
//
// swagger:model top_metrics_file_response_inline_incomplete_response_reason
type TopMetricsFileResponseInlineIncompleteResponseReason struct {

	// Warning code indicating why partial data was reported.
	// Example: 111411207
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// A message describing the reason for partial data.
	// Example: Partial data has been returned for this metric report. Reason: The activity tracking report for this volume is not available because the system is busy collecting tracking data.
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this top metrics file response inline incomplete response reason
func (m *TopMetricsFileResponseInlineIncompleteResponseReason) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this top metrics file response inline incomplete response reason based on the context it is used
func (m *TopMetricsFileResponseInlineIncompleteResponseReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *TopMetricsFileResponseInlineIncompleteResponseReason) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "incomplete_response_reason"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsFileResponseInlineIncompleteResponseReason) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "incomplete_response_reason"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsFileResponseInlineIncompleteResponseReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsFileResponseInlineIncompleteResponseReason) UnmarshalBinary(b []byte) error {
	var res TopMetricsFileResponseInlineIncompleteResponseReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsFileResponseInlineLinks top metrics file response inline links
//
// swagger:model top_metrics_file_response_inline__links
type TopMetricsFileResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this top metrics file response inline links
func (m *TopMetricsFileResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsFileResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsFileResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics file response inline links based on the context it is used
func (m *TopMetricsFileResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsFileResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsFileResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsFileResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsFileResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsFileResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsFileResponseInlineNotice Optional field that indicates why no records are returned by the volume activity tracking REST API.
//
// swagger:model top_metrics_file_response_inline_notice
type TopMetricsFileResponseInlineNotice struct {

	// Warning code indicating why no records are returned.
	// Example: 111411207
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Details why no records are returned.
	// Example: No read/write traffic on volume.
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this top metrics file response inline notice
func (m *TopMetricsFileResponseInlineNotice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this top metrics file response inline notice based on the context it is used
func (m *TopMetricsFileResponseInlineNotice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *TopMetricsFileResponseInlineNotice) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "notice"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsFileResponseInlineNotice) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "notice"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsFileResponseInlineNotice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsFileResponseInlineNotice) UnmarshalBinary(b []byte) error {
	var res TopMetricsFileResponseInlineNotice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
