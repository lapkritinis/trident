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

// ApplicationTemplate Application templates
//
// swagger:model application_template
type ApplicationTemplate struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Description.
	// Read Only: true
	Description *string `json:"description,omitempty"`

	// Missing prerequisites.
	// Read Only: true
	MissingPrerequisites *string `json:"missing_prerequisites,omitempty"`

	// mongo db on san
	MongoDbOnSan *MongoDbOnSan `json:"mongo_db_on_san,omitempty"`

	// Template name.
	// Required: true
	// Read Only: true
	Name *string `json:"name"`

	// nas
	Nas *Nas `json:"nas,omitempty"`

	// nvme
	Nvme *ZappNvme `json:"nvme,omitempty"`

	// oracle on nfs
	OracleOnNfs *OracleOnNfs `json:"oracle_on_nfs,omitempty"`

	// oracle on san
	OracleOnSan *OracleOnSan `json:"oracle_on_san,omitempty"`

	// oracle rac on nfs
	OracleRacOnNfs *OracleRacOnNfs `json:"oracle_rac_on_nfs,omitempty"`

	// oracle rac on san
	OracleRacOnSan *OracleRacOnSan `json:"oracle_rac_on_san,omitempty"`

	// Access protocol.
	// Read Only: true
	// Enum: ["nas","nvme","s3","san"]
	Protocol *string `json:"protocol,omitempty"`

	// s3 bucket
	S3Bucket *ZappS3Bucket `json:"s3_bucket,omitempty"`

	// san
	San *San `json:"san,omitempty"`

	// sql on san
	SQLOnSan *SQLOnSan `json:"sql_on_san,omitempty"`

	// sql on smb
	SQLOnSmb *SQLOnSmb `json:"sql_on_smb,omitempty"`

	// vdi on nas
	VdiOnNas *VdiOnNas `json:"vdi_on_nas,omitempty"`

	// vdi on san
	VdiOnSan *VdiOnSan `json:"vdi_on_san,omitempty"`

	// vsi on nas
	VsiOnNas *VsiOnNas `json:"vsi_on_nas,omitempty"`

	// vsi on san
	VsiOnSan *VsiOnSan `json:"vsi_on_san,omitempty"`
}

// Validate validates this application template
func (m *ApplicationTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMongoDbOnSan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleOnNfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleOnSan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleRacOnNfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleRacOnSan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3Bucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSQLOnSan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSQLOnSmb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVdiOnNas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVdiOnSan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsiOnNas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsiOnSan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationTemplate) validateLinks(formats strfmt.Registry) error {
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

func (m *ApplicationTemplate) validateMongoDbOnSan(formats strfmt.Registry) error {
	if swag.IsZero(m.MongoDbOnSan) { // not required
		return nil
	}

	if m.MongoDbOnSan != nil {
		if err := m.MongoDbOnSan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongo_db_on_san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationTemplate) validateNas(formats strfmt.Registry) error {
	if swag.IsZero(m.Nas) { // not required
		return nil
	}

	if m.Nas != nil {
		if err := m.Nas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nas")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateNvme(formats strfmt.Registry) error {
	if swag.IsZero(m.Nvme) { // not required
		return nil
	}

	if m.Nvme != nil {
		if err := m.Nvme.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvme")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateOracleOnNfs(formats strfmt.Registry) error {
	if swag.IsZero(m.OracleOnNfs) { // not required
		return nil
	}

	if m.OracleOnNfs != nil {
		if err := m.OracleOnNfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_on_nfs")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateOracleOnSan(formats strfmt.Registry) error {
	if swag.IsZero(m.OracleOnSan) { // not required
		return nil
	}

	if m.OracleOnSan != nil {
		if err := m.OracleOnSan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_on_san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateOracleRacOnNfs(formats strfmt.Registry) error {
	if swag.IsZero(m.OracleRacOnNfs) { // not required
		return nil
	}

	if m.OracleRacOnNfs != nil {
		if err := m.OracleRacOnNfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_rac_on_nfs")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateOracleRacOnSan(formats strfmt.Registry) error {
	if swag.IsZero(m.OracleRacOnSan) { // not required
		return nil
	}

	if m.OracleRacOnSan != nil {
		if err := m.OracleRacOnSan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_rac_on_san")
			}
			return err
		}
	}

	return nil
}

var applicationTemplateTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nas","nvme","s3","san"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationTemplateTypeProtocolPropEnum = append(applicationTemplateTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// application_template
	// ApplicationTemplate
	// protocol
	// Protocol
	// nas
	// END DEBUGGING
	// ApplicationTemplateProtocolNas captures enum value "nas"
	ApplicationTemplateProtocolNas string = "nas"

	// BEGIN DEBUGGING
	// application_template
	// ApplicationTemplate
	// protocol
	// Protocol
	// nvme
	// END DEBUGGING
	// ApplicationTemplateProtocolNvme captures enum value "nvme"
	ApplicationTemplateProtocolNvme string = "nvme"

	// BEGIN DEBUGGING
	// application_template
	// ApplicationTemplate
	// protocol
	// Protocol
	// s3
	// END DEBUGGING
	// ApplicationTemplateProtocolS3 captures enum value "s3"
	ApplicationTemplateProtocolS3 string = "s3"

	// BEGIN DEBUGGING
	// application_template
	// ApplicationTemplate
	// protocol
	// Protocol
	// san
	// END DEBUGGING
	// ApplicationTemplateProtocolSan captures enum value "san"
	ApplicationTemplateProtocolSan string = "san"
)

// prop value enum
func (m *ApplicationTemplate) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationTemplateTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationTemplate) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationTemplate) validateS3Bucket(formats strfmt.Registry) error {
	if swag.IsZero(m.S3Bucket) { // not required
		return nil
	}

	if m.S3Bucket != nil {
		if err := m.S3Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3_bucket")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateSan(formats strfmt.Registry) error {
	if swag.IsZero(m.San) { // not required
		return nil
	}

	if m.San != nil {
		if err := m.San.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateSQLOnSan(formats strfmt.Registry) error {
	if swag.IsZero(m.SQLOnSan) { // not required
		return nil
	}

	if m.SQLOnSan != nil {
		if err := m.SQLOnSan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sql_on_san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateSQLOnSmb(formats strfmt.Registry) error {
	if swag.IsZero(m.SQLOnSmb) { // not required
		return nil
	}

	if m.SQLOnSmb != nil {
		if err := m.SQLOnSmb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sql_on_smb")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateVdiOnNas(formats strfmt.Registry) error {
	if swag.IsZero(m.VdiOnNas) { // not required
		return nil
	}

	if m.VdiOnNas != nil {
		if err := m.VdiOnNas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vdi_on_nas")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateVdiOnSan(formats strfmt.Registry) error {
	if swag.IsZero(m.VdiOnSan) { // not required
		return nil
	}

	if m.VdiOnSan != nil {
		if err := m.VdiOnSan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vdi_on_san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateVsiOnNas(formats strfmt.Registry) error {
	if swag.IsZero(m.VsiOnNas) { // not required
		return nil
	}

	if m.VsiOnNas != nil {
		if err := m.VsiOnNas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsi_on_nas")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) validateVsiOnSan(formats strfmt.Registry) error {
	if swag.IsZero(m.VsiOnSan) { // not required
		return nil
	}

	if m.VsiOnSan != nil {
		if err := m.VsiOnSan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsi_on_san")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application template based on the context it is used
func (m *ApplicationTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMissingPrerequisites(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMongoDbOnSan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOracleOnNfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOracleOnSan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOracleRacOnNfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOracleRacOnSan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3Bucket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSQLOnSan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSQLOnSmb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVdiOnNas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVdiOnSan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsiOnNas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsiOnSan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationTemplate) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ApplicationTemplate) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateMissingPrerequisites(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "missing_prerequisites", "body", m.MissingPrerequisites); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateMongoDbOnSan(ctx context.Context, formats strfmt.Registry) error {

	if m.MongoDbOnSan != nil {
		if err := m.MongoDbOnSan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongo_db_on_san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateNas(ctx context.Context, formats strfmt.Registry) error {

	if m.Nas != nil {
		if err := m.Nas.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nas")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateNvme(ctx context.Context, formats strfmt.Registry) error {

	if m.Nvme != nil {
		if err := m.Nvme.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvme")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateOracleOnNfs(ctx context.Context, formats strfmt.Registry) error {

	if m.OracleOnNfs != nil {
		if err := m.OracleOnNfs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_on_nfs")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateOracleOnSan(ctx context.Context, formats strfmt.Registry) error {

	if m.OracleOnSan != nil {
		if err := m.OracleOnSan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_on_san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateOracleRacOnNfs(ctx context.Context, formats strfmt.Registry) error {

	if m.OracleRacOnNfs != nil {
		if err := m.OracleRacOnNfs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_rac_on_nfs")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateOracleRacOnSan(ctx context.Context, formats strfmt.Registry) error {

	if m.OracleRacOnSan != nil {
		if err := m.OracleRacOnSan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_rac_on_san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateS3Bucket(ctx context.Context, formats strfmt.Registry) error {

	if m.S3Bucket != nil {
		if err := m.S3Bucket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3_bucket")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateSan(ctx context.Context, formats strfmt.Registry) error {

	if m.San != nil {
		if err := m.San.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateSQLOnSan(ctx context.Context, formats strfmt.Registry) error {

	if m.SQLOnSan != nil {
		if err := m.SQLOnSan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sql_on_san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateSQLOnSmb(ctx context.Context, formats strfmt.Registry) error {

	if m.SQLOnSmb != nil {
		if err := m.SQLOnSmb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sql_on_smb")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateVdiOnNas(ctx context.Context, formats strfmt.Registry) error {

	if m.VdiOnNas != nil {
		if err := m.VdiOnNas.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vdi_on_nas")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateVdiOnSan(ctx context.Context, formats strfmt.Registry) error {

	if m.VdiOnSan != nil {
		if err := m.VdiOnSan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vdi_on_san")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateVsiOnNas(ctx context.Context, formats strfmt.Registry) error {

	if m.VsiOnNas != nil {
		if err := m.VsiOnNas.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsi_on_nas")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationTemplate) contextValidateVsiOnSan(ctx context.Context, formats strfmt.Registry) error {

	if m.VsiOnSan != nil {
		if err := m.VsiOnSan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsi_on_san")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationTemplate) UnmarshalBinary(b []byte) error {
	var res ApplicationTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
