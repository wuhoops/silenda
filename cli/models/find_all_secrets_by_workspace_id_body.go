// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FindAllSecretsByWorkspaceIDBody find all secrets by workspace Id body
//
// swagger:model FindAllSecretsByWorkspaceIdBody
type FindAllSecretsByWorkspaceIDBody struct {

	// workspace id
	WorkspaceID string `json:"workspace_id,omitempty"`
}

// Validate validates this find all secrets by workspace Id body
func (m *FindAllSecretsByWorkspaceIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this find all secrets by workspace Id body based on context it is used
func (m *FindAllSecretsByWorkspaceIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FindAllSecretsByWorkspaceIDBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FindAllSecretsByWorkspaceIDBody) UnmarshalBinary(b []byte) error {
	var res FindAllSecretsByWorkspaceIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
