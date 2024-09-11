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

// PortMetrics Throughput performance for the Ethernet port.
//
// swagger:model port_metrics
type PortMetrics struct {

	// links
	Links *PortMetricsInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Enum: ["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]
	Duration *string `json:"duration,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Enum: ["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]
	Status *string `json:"status,omitempty"`

	// throughput
	Throughput *PortMetricsInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// Port UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this port metrics
func (m *PortMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetrics) validateLinks(formats strfmt.Registry) error {
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

var portMetricsTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portMetricsTypeDurationPropEnum = append(portMetricsTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PortMetricsDurationPT15S captures enum value "PT15S"
	PortMetricsDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PortMetricsDurationPT4M captures enum value "PT4M"
	PortMetricsDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PortMetricsDurationPT30M captures enum value "PT30M"
	PortMetricsDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PortMetricsDurationPT2H captures enum value "PT2H"
	PortMetricsDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PortMetricsDurationP1D captures enum value "P1D"
	PortMetricsDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PortMetricsDurationPT5M captures enum value "PT5M"
	PortMetricsDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PortMetrics) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portMetricsTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortMetrics) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

var portMetricsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portMetricsTypeStatusPropEnum = append(portMetricsTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// ok
	// END DEBUGGING
	// PortMetricsStatusOk captures enum value "ok"
	PortMetricsStatusOk string = "ok"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// error
	// END DEBUGGING
	// PortMetricsStatusError captures enum value "error"
	PortMetricsStatusError string = "error"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PortMetricsStatusPartialNoData captures enum value "partial_no_data"
	PortMetricsStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PortMetricsStatusPartialNoUUID captures enum value "partial_no_uuid"
	PortMetricsStatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PortMetricsStatusPartialNoResponse captures enum value "partial_no_response"
	PortMetricsStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PortMetricsStatusPartialOtherError captures enum value "partial_other_error"
	PortMetricsStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PortMetricsStatusNegativeDelta captures enum value "negative_delta"
	PortMetricsStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PortMetricsStatusBackfilledData captures enum value "backfilled_data"
	PortMetricsStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PortMetricsStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PortMetricsStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// port_metrics
	// PortMetrics
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PortMetricsStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PortMetricsStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *PortMetrics) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portMetricsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortMetrics) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PortMetrics) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PortMetrics) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this port metrics based on the context it is used
func (m *PortMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
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

func (m *PortMetrics) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortMetrics) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PortMetrics) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetrics) UnmarshalBinary(b []byte) error {
	var res PortMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortMetricsInlineLinks port metrics inline links
//
// swagger:model port_metrics_inline__links
type PortMetricsInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this port metrics inline links
func (m *PortMetricsInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetricsInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this port metrics inline links based on the context it is used
func (m *PortMetricsInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetricsInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PortMetricsInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsInlineLinks) UnmarshalBinary(b []byte) error {
	var res PortMetricsInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortMetricsInlineThroughput The rate of throughput bytes per second observed at the interface.
//
// swagger:model port_metrics_inline_throughput
type PortMetricsInlineThroughput struct {

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Performance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this port metrics inline throughput
func (m *PortMetricsInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port metrics inline throughput based on context it is used
func (m *PortMetricsInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortMetricsInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PortMetricsInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
