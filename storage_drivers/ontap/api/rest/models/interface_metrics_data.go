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

// InterfaceMetricsData Throughput performance for the interfaces.
//
// swagger:model interface_metrics_data
type InterfaceMetricsData struct {

	// links
	Links *InterfaceMetricsDataInlineLinks `json:"_links,omitempty"`

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
	Throughput *InterfaceMetricsDataInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this interface metrics data
func (m *InterfaceMetricsData) Validate(formats strfmt.Registry) error {
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

func (m *InterfaceMetricsData) validateLinks(formats strfmt.Registry) error {
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

var interfaceMetricsDataTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceMetricsDataTypeDurationPropEnum = append(interfaceMetricsDataTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// InterfaceMetricsDataDurationPT15S captures enum value "PT15S"
	InterfaceMetricsDataDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// InterfaceMetricsDataDurationPT4M captures enum value "PT4M"
	InterfaceMetricsDataDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// InterfaceMetricsDataDurationPT30M captures enum value "PT30M"
	InterfaceMetricsDataDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// InterfaceMetricsDataDurationPT2H captures enum value "PT2H"
	InterfaceMetricsDataDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// InterfaceMetricsDataDurationP1D captures enum value "P1D"
	InterfaceMetricsDataDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// InterfaceMetricsDataDurationPT5M captures enum value "PT5M"
	InterfaceMetricsDataDurationPT5M string = "PT5M"
)

// prop value enum
func (m *InterfaceMetricsData) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceMetricsDataTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceMetricsData) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

var interfaceMetricsDataTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceMetricsDataTypeStatusPropEnum = append(interfaceMetricsDataTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// ok
	// END DEBUGGING
	// InterfaceMetricsDataStatusOk captures enum value "ok"
	InterfaceMetricsDataStatusOk string = "ok"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// error
	// END DEBUGGING
	// InterfaceMetricsDataStatusError captures enum value "error"
	InterfaceMetricsDataStatusError string = "error"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// InterfaceMetricsDataStatusPartialNoData captures enum value "partial_no_data"
	InterfaceMetricsDataStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// InterfaceMetricsDataStatusPartialNoUUID captures enum value "partial_no_uuid"
	InterfaceMetricsDataStatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// InterfaceMetricsDataStatusPartialNoResponse captures enum value "partial_no_response"
	InterfaceMetricsDataStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// InterfaceMetricsDataStatusPartialOtherError captures enum value "partial_other_error"
	InterfaceMetricsDataStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// InterfaceMetricsDataStatusNegativeDelta captures enum value "negative_delta"
	InterfaceMetricsDataStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// InterfaceMetricsDataStatusBackfilledData captures enum value "backfilled_data"
	InterfaceMetricsDataStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// InterfaceMetricsDataStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	InterfaceMetricsDataStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// interface_metrics_data
	// InterfaceMetricsData
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// InterfaceMetricsDataStatusInconsistentOldData captures enum value "inconsistent_old_data"
	InterfaceMetricsDataStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *InterfaceMetricsData) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceMetricsDataTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceMetricsData) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceMetricsData) validateThroughput(formats strfmt.Registry) error {
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

func (m *InterfaceMetricsData) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this interface metrics data based on the context it is used
func (m *InterfaceMetricsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceMetricsData) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *InterfaceMetricsData) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *InterfaceMetricsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceMetricsData) UnmarshalBinary(b []byte) error {
	var res InterfaceMetricsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceMetricsDataInlineLinks interface metrics data inline links
//
// swagger:model interface_metrics_data_inline__links
type InterfaceMetricsDataInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this interface metrics data inline links
func (m *InterfaceMetricsDataInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceMetricsDataInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this interface metrics data inline links based on the context it is used
func (m *InterfaceMetricsDataInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceMetricsDataInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *InterfaceMetricsDataInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceMetricsDataInlineLinks) UnmarshalBinary(b []byte) error {
	var res InterfaceMetricsDataInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceMetricsDataInlineThroughput The rate of throughput bytes per second observed at the interface.
//
// swagger:model interface_metrics_data_inline_throughput
type InterfaceMetricsDataInlineThroughput struct {

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

// Validate validates this interface metrics data inline throughput
func (m *InterfaceMetricsDataInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this interface metrics data inline throughput based on context it is used
func (m *InterfaceMetricsDataInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceMetricsDataInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceMetricsDataInlineThroughput) UnmarshalBinary(b []byte) error {
	var res InterfaceMetricsDataInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
