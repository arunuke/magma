// Code generated by go-swagger; DO NOT EDIT.

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetChannelsChannelIDParams creates a new GetChannelsChannelIDParams object
// with the default values initialized.
func NewGetChannelsChannelIDParams() *GetChannelsChannelIDParams {
	var ()
	return &GetChannelsChannelIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChannelsChannelIDParamsWithTimeout creates a new GetChannelsChannelIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChannelsChannelIDParamsWithTimeout(timeout time.Duration) *GetChannelsChannelIDParams {
	var ()
	return &GetChannelsChannelIDParams{

		timeout: timeout,
	}
}

// NewGetChannelsChannelIDParamsWithContext creates a new GetChannelsChannelIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChannelsChannelIDParamsWithContext(ctx context.Context) *GetChannelsChannelIDParams {
	var ()
	return &GetChannelsChannelIDParams{

		Context: ctx,
	}
}

// NewGetChannelsChannelIDParamsWithHTTPClient creates a new GetChannelsChannelIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChannelsChannelIDParamsWithHTTPClient(client *http.Client) *GetChannelsChannelIDParams {
	var ()
	return &GetChannelsChannelIDParams{
		HTTPClient: client,
	}
}

/*GetChannelsChannelIDParams contains all the parameters to send to the API endpoint
for the get channels channel ID operation typically these are written to a http.Request
*/
type GetChannelsChannelIDParams struct {

	/*ChannelID
	  Release Channel ID

	*/
	ChannelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get channels channel ID params
func (o *GetChannelsChannelIDParams) WithTimeout(timeout time.Duration) *GetChannelsChannelIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get channels channel ID params
func (o *GetChannelsChannelIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get channels channel ID params
func (o *GetChannelsChannelIDParams) WithContext(ctx context.Context) *GetChannelsChannelIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get channels channel ID params
func (o *GetChannelsChannelIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get channels channel ID params
func (o *GetChannelsChannelIDParams) WithHTTPClient(client *http.Client) *GetChannelsChannelIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get channels channel ID params
func (o *GetChannelsChannelIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the get channels channel ID params
func (o *GetChannelsChannelIDParams) WithChannelID(channelID string) *GetChannelsChannelIDParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the get channels channel ID params
func (o *GetChannelsChannelIDParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WriteToRequest writes these params to a swagger request
func (o *GetChannelsChannelIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel_id
	if err := r.SetPathParam("channel_id", o.ChannelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
