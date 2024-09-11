// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSoftwareModifyParams creates a new SoftwareModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSoftwareModifyParams() *SoftwareModifyParams {
	return &SoftwareModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwareModifyParamsWithTimeout creates a new SoftwareModifyParams object
// with the ability to set a timeout on a request.
func NewSoftwareModifyParamsWithTimeout(timeout time.Duration) *SoftwareModifyParams {
	return &SoftwareModifyParams{
		timeout: timeout,
	}
}

// NewSoftwareModifyParamsWithContext creates a new SoftwareModifyParams object
// with the ability to set a context for a request.
func NewSoftwareModifyParamsWithContext(ctx context.Context) *SoftwareModifyParams {
	return &SoftwareModifyParams{
		Context: ctx,
	}
}

// NewSoftwareModifyParamsWithHTTPClient creates a new SoftwareModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSoftwareModifyParamsWithHTTPClient(client *http.Client) *SoftwareModifyParams {
	return &SoftwareModifyParams{
		HTTPClient: client,
	}
}

/*
SoftwareModifyParams contains all the parameters to send to the API endpoint

	for the software modify operation.

	Typically these are written to a http.Request.
*/
type SoftwareModifyParams struct {

	/* Action.

	     Requests an upgrade to pause, resume, or cancel.
	Note that not all upgrades support these actions. An upgrade can only be resumed if it is in the paused state. When a request to cancel an upgrade is successful, the upgrade state changes to either `success` or `failure`.

	*/
	Action *string

	/* EstimateOnly.

	     Generates an estimate of the time required for the overall update operation for the specified package.
	No update is performed when this option is used. The default is false.

	*/
	EstimateOnly *bool

	/* ForceRolling.

	   Forces a rolling upgrade on the cluster. This option is not applicable for a single-node cluster and for a two-node MetroCluster. The default is false.

	*/
	ForceRolling *bool

	// Info.
	Info *models.SoftwareReference

	/* NodesToUpdate.

	   A comma separated list of node names to be updated. The nodes must be a part of a HA Pair. The default is all nodes. If the nodes_to_update parameter is empty then upgrade will error out and will not proceed.

	*/
	NodesToUpdate *string

	/* PauseAfter.

	   The pause after specified tasks option. When ANDU is paused user interaction is required to resume the update. The default is none.

	*/
	PauseAfter *string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* ShowValidationDetails.

	   If the value is set to true, then all validation details will be shown in the output.

	*/
	ShowValidationDetails *bool

	/* SkipNodesAtTargetVersion.

	     Defaults to true in non-MetroCluster configurations. When set to true, nodes already at the target version will not be upgraded.
	Setting this field to false will force nodes at the target version to undergo upgrade including migrating LIFs, performing takeover/giveback, and rebooting.
	It's invalid to set this option to true when the action field is set.
	It's an invalid option until effective cluster version is 9_16_0.

	*/
	SkipNodesAtTargetVersion *bool

	/* SkipWarnings.

	   Ignore warnings and proceed with the install.
	*/
	SkipWarnings *bool

	/* StabilizeMinutes.

	   Sets a custom value between 1 to 60 minutes for the upgrade, allowing each node a specified amount of time to stabilize after a reboot.
	*/
	StabilizeMinutes *int64

	/* ValidateOnly.

	   Validate the operation and its parameters, without actually performing the operation.
	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the software modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwareModifyParams) WithDefaults() *SoftwareModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the software modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwareModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := SoftwareModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the software modify params
func (o *SoftwareModifyParams) WithTimeout(timeout time.Duration) *SoftwareModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software modify params
func (o *SoftwareModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software modify params
func (o *SoftwareModifyParams) WithContext(ctx context.Context) *SoftwareModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software modify params
func (o *SoftwareModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software modify params
func (o *SoftwareModifyParams) WithHTTPClient(client *http.Client) *SoftwareModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software modify params
func (o *SoftwareModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the software modify params
func (o *SoftwareModifyParams) WithAction(action *string) *SoftwareModifyParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the software modify params
func (o *SoftwareModifyParams) SetAction(action *string) {
	o.Action = action
}

// WithEstimateOnly adds the estimateOnly to the software modify params
func (o *SoftwareModifyParams) WithEstimateOnly(estimateOnly *bool) *SoftwareModifyParams {
	o.SetEstimateOnly(estimateOnly)
	return o
}

// SetEstimateOnly adds the estimateOnly to the software modify params
func (o *SoftwareModifyParams) SetEstimateOnly(estimateOnly *bool) {
	o.EstimateOnly = estimateOnly
}

// WithForceRolling adds the forceRolling to the software modify params
func (o *SoftwareModifyParams) WithForceRolling(forceRolling *bool) *SoftwareModifyParams {
	o.SetForceRolling(forceRolling)
	return o
}

// SetForceRolling adds the forceRolling to the software modify params
func (o *SoftwareModifyParams) SetForceRolling(forceRolling *bool) {
	o.ForceRolling = forceRolling
}

// WithInfo adds the info to the software modify params
func (o *SoftwareModifyParams) WithInfo(info *models.SoftwareReference) *SoftwareModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the software modify params
func (o *SoftwareModifyParams) SetInfo(info *models.SoftwareReference) {
	o.Info = info
}

// WithNodesToUpdate adds the nodesToUpdate to the software modify params
func (o *SoftwareModifyParams) WithNodesToUpdate(nodesToUpdate *string) *SoftwareModifyParams {
	o.SetNodesToUpdate(nodesToUpdate)
	return o
}

// SetNodesToUpdate adds the nodesToUpdate to the software modify params
func (o *SoftwareModifyParams) SetNodesToUpdate(nodesToUpdate *string) {
	o.NodesToUpdate = nodesToUpdate
}

// WithPauseAfter adds the pauseAfter to the software modify params
func (o *SoftwareModifyParams) WithPauseAfter(pauseAfter *string) *SoftwareModifyParams {
	o.SetPauseAfter(pauseAfter)
	return o
}

// SetPauseAfter adds the pauseAfter to the software modify params
func (o *SoftwareModifyParams) SetPauseAfter(pauseAfter *string) {
	o.PauseAfter = pauseAfter
}

// WithReturnTimeout adds the returnTimeout to the software modify params
func (o *SoftwareModifyParams) WithReturnTimeout(returnTimeout *int64) *SoftwareModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the software modify params
func (o *SoftwareModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithShowValidationDetails adds the showValidationDetails to the software modify params
func (o *SoftwareModifyParams) WithShowValidationDetails(showValidationDetails *bool) *SoftwareModifyParams {
	o.SetShowValidationDetails(showValidationDetails)
	return o
}

// SetShowValidationDetails adds the showValidationDetails to the software modify params
func (o *SoftwareModifyParams) SetShowValidationDetails(showValidationDetails *bool) {
	o.ShowValidationDetails = showValidationDetails
}

// WithSkipNodesAtTargetVersion adds the skipNodesAtTargetVersion to the software modify params
func (o *SoftwareModifyParams) WithSkipNodesAtTargetVersion(skipNodesAtTargetVersion *bool) *SoftwareModifyParams {
	o.SetSkipNodesAtTargetVersion(skipNodesAtTargetVersion)
	return o
}

// SetSkipNodesAtTargetVersion adds the skipNodesAtTargetVersion to the software modify params
func (o *SoftwareModifyParams) SetSkipNodesAtTargetVersion(skipNodesAtTargetVersion *bool) {
	o.SkipNodesAtTargetVersion = skipNodesAtTargetVersion
}

// WithSkipWarnings adds the skipWarnings to the software modify params
func (o *SoftwareModifyParams) WithSkipWarnings(skipWarnings *bool) *SoftwareModifyParams {
	o.SetSkipWarnings(skipWarnings)
	return o
}

// SetSkipWarnings adds the skipWarnings to the software modify params
func (o *SoftwareModifyParams) SetSkipWarnings(skipWarnings *bool) {
	o.SkipWarnings = skipWarnings
}

// WithStabilizeMinutes adds the stabilizeMinutes to the software modify params
func (o *SoftwareModifyParams) WithStabilizeMinutes(stabilizeMinutes *int64) *SoftwareModifyParams {
	o.SetStabilizeMinutes(stabilizeMinutes)
	return o
}

// SetStabilizeMinutes adds the stabilizeMinutes to the software modify params
func (o *SoftwareModifyParams) SetStabilizeMinutes(stabilizeMinutes *int64) {
	o.StabilizeMinutes = stabilizeMinutes
}

// WithValidateOnly adds the validateOnly to the software modify params
func (o *SoftwareModifyParams) WithValidateOnly(validateOnly *bool) *SoftwareModifyParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the software modify params
func (o *SoftwareModifyParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwareModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string

		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	if o.EstimateOnly != nil {

		// query param estimate_only
		var qrEstimateOnly bool

		if o.EstimateOnly != nil {
			qrEstimateOnly = *o.EstimateOnly
		}
		qEstimateOnly := swag.FormatBool(qrEstimateOnly)
		if qEstimateOnly != "" {

			if err := r.SetQueryParam("estimate_only", qEstimateOnly); err != nil {
				return err
			}
		}
	}

	if o.ForceRolling != nil {

		// query param force_rolling
		var qrForceRolling bool

		if o.ForceRolling != nil {
			qrForceRolling = *o.ForceRolling
		}
		qForceRolling := swag.FormatBool(qrForceRolling)
		if qForceRolling != "" {

			if err := r.SetQueryParam("force_rolling", qForceRolling); err != nil {
				return err
			}
		}
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.NodesToUpdate != nil {

		// query param nodes_to_update
		var qrNodesToUpdate string

		if o.NodesToUpdate != nil {
			qrNodesToUpdate = *o.NodesToUpdate
		}
		qNodesToUpdate := qrNodesToUpdate
		if qNodesToUpdate != "" {

			if err := r.SetQueryParam("nodes_to_update", qNodesToUpdate); err != nil {
				return err
			}
		}
	}

	if o.PauseAfter != nil {

		// query param pause_after
		var qrPauseAfter string

		if o.PauseAfter != nil {
			qrPauseAfter = *o.PauseAfter
		}
		qPauseAfter := qrPauseAfter
		if qPauseAfter != "" {

			if err := r.SetQueryParam("pause_after", qPauseAfter); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.ShowValidationDetails != nil {

		// query param show_validation_details
		var qrShowValidationDetails bool

		if o.ShowValidationDetails != nil {
			qrShowValidationDetails = *o.ShowValidationDetails
		}
		qShowValidationDetails := swag.FormatBool(qrShowValidationDetails)
		if qShowValidationDetails != "" {

			if err := r.SetQueryParam("show_validation_details", qShowValidationDetails); err != nil {
				return err
			}
		}
	}

	if o.SkipNodesAtTargetVersion != nil {

		// query param skip_nodes_at_target_version
		var qrSkipNodesAtTargetVersion bool

		if o.SkipNodesAtTargetVersion != nil {
			qrSkipNodesAtTargetVersion = *o.SkipNodesAtTargetVersion
		}
		qSkipNodesAtTargetVersion := swag.FormatBool(qrSkipNodesAtTargetVersion)
		if qSkipNodesAtTargetVersion != "" {

			if err := r.SetQueryParam("skip_nodes_at_target_version", qSkipNodesAtTargetVersion); err != nil {
				return err
			}
		}
	}

	if o.SkipWarnings != nil {

		// query param skip_warnings
		var qrSkipWarnings bool

		if o.SkipWarnings != nil {
			qrSkipWarnings = *o.SkipWarnings
		}
		qSkipWarnings := swag.FormatBool(qrSkipWarnings)
		if qSkipWarnings != "" {

			if err := r.SetQueryParam("skip_warnings", qSkipWarnings); err != nil {
				return err
			}
		}
	}

	if o.StabilizeMinutes != nil {

		// query param stabilize_minutes
		var qrStabilizeMinutes int64

		if o.StabilizeMinutes != nil {
			qrStabilizeMinutes = *o.StabilizeMinutes
		}
		qStabilizeMinutes := swag.FormatInt64(qrStabilizeMinutes)
		if qStabilizeMinutes != "" {

			if err := r.SetQueryParam("stabilize_minutes", qStabilizeMinutes); err != nil {
				return err
			}
		}
	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool

		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {

			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
