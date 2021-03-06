// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SentryConfig Sentry.io configuration
// swagger:model sentry_config
type SentryConfig struct {

	// sample rate
	SampleRate float32 `json:"sample_rate,omitempty"`

	// upload mme log
	UploadMmeLog bool `json:"upload_mme_log,omitempty"`

	// url native
	// Min Length: 0
	URLNative string `json:"url_native,omitempty"`

	// url python
	// Min Length: 0
	URLPython string `json:"url_python,omitempty"`
}

// Validate validates this sentry config
func (m *SentryConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURLNative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURLPython(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SentryConfig) validateURLNative(formats strfmt.Registry) error {

	if swag.IsZero(m.URLNative) { // not required
		return nil
	}

	if err := validate.MinLength("url_native", "body", string(m.URLNative), 0); err != nil {
		return err
	}

	return nil
}

func (m *SentryConfig) validateURLPython(formats strfmt.Registry) error {

	if swag.IsZero(m.URLPython) { // not required
		return nil
	}

	if err := validate.MinLength("url_python", "body", string(m.URLPython), 0); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SentryConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SentryConfig) UnmarshalBinary(b []byte) error {
	var res SentryConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
