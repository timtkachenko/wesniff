// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserUpdateDto user update dto
//
// swagger:model UserUpdateDto
type UserUpdateDto struct {
	roleField *string
}

// Role gets the role of this subtype
func (m *UserUpdateDto) Role() *string {
	return m.roleField
}

// SetRole sets the role of this subtype
func (m *UserUpdateDto) SetRole(val *string) {
	m.roleField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *UserUpdateDto) UnmarshalJSON(raw []byte) error {
	var data struct {
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Role *string `json:"role,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result UserUpdateDto

	result.roleField = base.Role

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m UserUpdateDto) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Role *string `json:"role,omitempty"`
	}{

		Role: m.Role(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this user update dto
func (m *UserUpdateDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userUpdateDtoTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["admin","member"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userUpdateDtoTypeRolePropEnum = append(userUpdateDtoTypeRolePropEnum, v)
	}
}

// property enum
func (m *UserUpdateDto) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userUpdateDtoTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserUpdateDto) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role()) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserUpdateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUpdateDto) UnmarshalBinary(b []byte) error {
	var res UserUpdateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}