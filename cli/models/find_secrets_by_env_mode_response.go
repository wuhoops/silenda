// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FindSecretsByEnvModeResponse find secrets by env mode response
//
// swagger:model FindSecretsByEnvModeResponse
type FindSecretsByEnvModeResponse struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this find secrets by env mode response
func (m *FindSecretsByEnvModeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this find secrets by env mode response based on context it is used
func (m *FindSecretsByEnvModeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FindSecretsByEnvModeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FindSecretsByEnvModeResponse) UnmarshalBinary(b []byte) error {
	var res FindSecretsByEnvModeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
