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
)

// SvmPeerJobLinkResponse svm peer job link response
//
// swagger:model svm_peer_job_link_response
type SvmPeerJobLinkResponse struct {

	// job
	Job *JobLink `json:"job,omitempty"`

	// Number of records
	// Example: 1
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*SvmPeer `json:"records"`
}

// Validate validates this svm peer job link response
func (m *SvmPeerJobLinkResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerJobLinkResponse) validateJob(formats strfmt.Registry) error {
	if swag.IsZero(m.Job) { // not required
		return nil
	}

	if m.Job != nil {
		if err := m.Job.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *SvmPeerJobLinkResponse) validateRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this svm peer job link response based on the context it is used
func (m *SvmPeerJobLinkResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerJobLinkResponse) contextValidateJob(ctx context.Context, formats strfmt.Registry) error {

	if m.Job != nil {
		if err := m.Job.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *SvmPeerJobLinkResponse) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {
			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
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
func (m *SvmPeerJobLinkResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerJobLinkResponse) UnmarshalBinary(b []byte) error {
	var res SvmPeerJobLinkResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
