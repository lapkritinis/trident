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

// SnapmirrorSourceEndpoint Source endpoint of a SnapMirror relationship. For a GET request, the property "cluster" is populated when the endpoint is on a remote cluster. A POST request to establish a SnapMirror relationship between the source endpoint and destination endpoint and when the source SVM and the destination SVM are not peered, must specify the "cluster" property for the remote endpoint.
//
// swagger:model snapmirror_source_endpoint
type SnapmirrorSourceEndpoint struct {

	// cluster
	Cluster *SnapmirrorSourceEndpointInlineCluster `json:"cluster,omitempty"`

	// luns
	Luns *SnapmirrorSourceEndpointInlineLuns `json:"luns,omitempty"`

	// ONTAP FlexVol/FlexGroup - svm1:volume1
	// ONTAP SVM               - svm1:
	// ONTAP Consistency Group - svm1:/cg/cg_name
	// ONTAP S3                - svm1:/bucket/bucket1
	// NON-ONTAP               - objstore1:/objstore
	//
	// Example: svm1:volume1
	Path *string `json:"path,omitempty"`

	// This property specifies the list of FlexVol volumes or LUNs of a Consistency Group. Optional on the ASA r2 platform. Mandatory for all other platforms.
	SnapmirrorSourceEndpointInlineConsistencyGroupVolumes []*SnapmirrorSourceEndpointInlineConsistencyGroupVolumesInlineArrayItem `json:"consistency_group_volumes,omitempty"`

	// svm
	Svm *SnapmirrorSourceEndpointInlineSvm `json:"svm,omitempty"`
}

// Validate validates this snapmirror source endpoint
func (m *SnapmirrorSourceEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapmirrorSourceEndpointInlineConsistencyGroupVolumes(formats); err != nil {
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

func (m *SnapmirrorSourceEndpoint) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorSourceEndpoint) validateLuns(formats strfmt.Registry) error {
	if swag.IsZero(m.Luns) { // not required
		return nil
	}

	if m.Luns != nil {
		if err := m.Luns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("luns")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorSourceEndpoint) validateSnapmirrorSourceEndpointInlineConsistencyGroupVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapmirrorSourceEndpointInlineConsistencyGroupVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapmirrorSourceEndpointInlineConsistencyGroupVolumes); i++ {
		if swag.IsZero(m.SnapmirrorSourceEndpointInlineConsistencyGroupVolumes[i]) { // not required
			continue
		}

		if m.SnapmirrorSourceEndpointInlineConsistencyGroupVolumes[i] != nil {
			if err := m.SnapmirrorSourceEndpointInlineConsistencyGroupVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consistency_group_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapmirrorSourceEndpoint) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this snapmirror source endpoint based on the context it is used
func (m *SnapmirrorSourceEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapmirrorSourceEndpointInlineConsistencyGroupVolumes(ctx, formats); err != nil {
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

func (m *SnapmirrorSourceEndpoint) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorSourceEndpoint) contextValidateLuns(ctx context.Context, formats strfmt.Registry) error {

	if m.Luns != nil {
		if err := m.Luns.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("luns")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorSourceEndpoint) contextValidateSnapmirrorSourceEndpointInlineConsistencyGroupVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SnapmirrorSourceEndpointInlineConsistencyGroupVolumes); i++ {

		if m.SnapmirrorSourceEndpointInlineConsistencyGroupVolumes[i] != nil {
			if err := m.SnapmirrorSourceEndpointInlineConsistencyGroupVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consistency_group_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapmirrorSourceEndpoint) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnapmirrorSourceEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorSourceEndpoint) UnmarshalBinary(b []byte) error {
	var res SnapmirrorSourceEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorSourceEndpointInlineCluster snapmirror source endpoint inline cluster
//
// swagger:model snapmirror_source_endpoint_inline_cluster
type SnapmirrorSourceEndpointInlineCluster struct {

	// links
	Links *SnapmirrorSourceEndpointInlineClusterInlineLinks `json:"_links,omitempty"`

	// name
	// Example: cluster1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this snapmirror source endpoint inline cluster
func (m *SnapmirrorSourceEndpointInlineCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineCluster) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorSourceEndpointInlineCluster) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster"+"."+"uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snapmirror source endpoint inline cluster based on the context it is used
func (m *SnapmirrorSourceEndpointInlineCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineCluster) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineCluster) UnmarshalBinary(b []byte) error {
	var res SnapmirrorSourceEndpointInlineCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorSourceEndpointInlineClusterInlineLinks snapmirror source endpoint inline cluster inline links
//
// swagger:model snapmirror_source_endpoint_inline_cluster_inline__links
type SnapmirrorSourceEndpointInlineClusterInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapmirror source endpoint inline cluster inline links
func (m *SnapmirrorSourceEndpointInlineClusterInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineClusterInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapmirror source endpoint inline cluster inline links based on the context it is used
func (m *SnapmirrorSourceEndpointInlineClusterInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineClusterInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineClusterInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineClusterInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapmirrorSourceEndpointInlineClusterInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorSourceEndpointInlineConsistencyGroupVolumesInlineArrayItem snapmirror source endpoint inline consistency group volumes inline array item
//
// swagger:model snapmirror_source_endpoint_inline_consistency_group_volumes_inline_array_item
type SnapmirrorSourceEndpointInlineConsistencyGroupVolumesInlineArrayItem struct {

	// The name of the volume.
	// Example: volume1
	Name *string `json:"name,omitempty"`

	// Unique identifier of the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this snapmirror source endpoint inline consistency group volumes inline array item
func (m *SnapmirrorSourceEndpointInlineConsistencyGroupVolumesInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this snapmirror source endpoint inline consistency group volumes inline array item based on context it is used
func (m *SnapmirrorSourceEndpointInlineConsistencyGroupVolumesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineConsistencyGroupVolumesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineConsistencyGroupVolumesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res SnapmirrorSourceEndpointInlineConsistencyGroupVolumesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorSourceEndpointInlineLuns Optional property for a SnapMirror endpoint. Specifies the list of source LUNs and optionally list of destination LUNs during SnapMirror Consistency Group LUN restore operation.
//
// swagger:model snapmirror_source_endpoint_inline_luns
type SnapmirrorSourceEndpointInlineLuns struct {

	// links
	Links *SnapmirrorSourceEndpointInlineLunsInlineLinks `json:"_links,omitempty"`

	// The name of a LUN.
	// ### Platform Specifics
	// * **Unified ONTAP**:
	// A LUN is located within a volume. Optionally, it can be located within a qtree in a volume.<br/>
	// LUN names are paths of the form "/vol/\<volume>[/\<qtree>]/\<namespace>" where the qtree name is optional.
	// * **ASA r2**:
	// LUN names are simple names that share a namespace with LUNs within the same SVM. The name must begin with a letter or "\_" and contain only "\_" and alphanumeric characters. In specific cases, an optional snapshot-name can be used of the form "\<name>[@\<snapshot-name>]". The snapshot name must not begin or end with whitespace.
	//
	// Example: /vol/volume1/lun1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the LUN.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this snapmirror source endpoint inline luns
func (m *SnapmirrorSourceEndpointInlineLuns) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineLuns) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("luns" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapmirror source endpoint inline luns based on the context it is used
func (m *SnapmirrorSourceEndpointInlineLuns) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineLuns) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("luns" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineLuns) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineLuns) UnmarshalBinary(b []byte) error {
	var res SnapmirrorSourceEndpointInlineLuns
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorSourceEndpointInlineLunsInlineLinks snapmirror source endpoint inline luns inline links
//
// swagger:model snapmirror_source_endpoint_inline_luns_inline__links
type SnapmirrorSourceEndpointInlineLunsInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapmirror source endpoint inline luns inline links
func (m *SnapmirrorSourceEndpointInlineLunsInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineLunsInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("luns" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapmirror source endpoint inline luns inline links based on the context it is used
func (m *SnapmirrorSourceEndpointInlineLunsInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineLunsInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("luns" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineLunsInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineLunsInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapmirrorSourceEndpointInlineLunsInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorSourceEndpointInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model snapmirror_source_endpoint_inline_svm
type SnapmirrorSourceEndpointInlineSvm struct {

	// links
	Links *SnapmirrorSourceEndpointInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this snapmirror source endpoint inline svm
func (m *SnapmirrorSourceEndpointInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this snapmirror source endpoint inline svm based on the context it is used
func (m *SnapmirrorSourceEndpointInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnapmirrorSourceEndpointInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineSvm) UnmarshalBinary(b []byte) error {
	var res SnapmirrorSourceEndpointInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorSourceEndpointInlineSvmInlineLinks snapmirror source endpoint inline svm inline links
//
// swagger:model snapmirror_source_endpoint_inline_svm_inline__links
type SnapmirrorSourceEndpointInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapmirror source endpoint inline svm inline links
func (m *SnapmirrorSourceEndpointInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this snapmirror source endpoint inline svm inline links based on the context it is used
func (m *SnapmirrorSourceEndpointInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorSourceEndpointInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnapmirrorSourceEndpointInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorSourceEndpointInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapmirrorSourceEndpointInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
