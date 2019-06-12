// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchVolumesIDParamsBody VolumePatch
// swagger:model patchVolumesIdParamsBody
type PatchVolumesIDParamsBody struct {

	// The capacity of the volume in gigabytes
	// Maximum: 64000
	// Minimum: 10
	Capacity int64 `json:"capacity,omitempty"`

	// The bandwidth for the volume
	// Enum: [1000 10000 100000]
	Iops int64 `json:"iops,omitempty"`

	// The user-defined name for this volume
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// profile
	Profile *PatchVolumesIDParamsBodyProfile `json:"profile,omitempty"`
}

// Validate validates this patch volumes Id params body
func (m *PatchVolumesIDParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchVolumesIDParamsBody) validateCapacity(formats strfmt.Registry) error {

	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if err := validate.MinimumInt("capacity", "body", int64(m.Capacity), 10, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("capacity", "body", int64(m.Capacity), 64000, false); err != nil {
		return err
	}

	return nil
}

var patchVolumesIdParamsBodyTypeIopsPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1000,10000,100000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchVolumesIdParamsBodyTypeIopsPropEnum = append(patchVolumesIdParamsBodyTypeIopsPropEnum, v)
	}
}

// prop value enum
func (m *PatchVolumesIDParamsBody) validateIopsEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, patchVolumesIdParamsBodyTypeIopsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PatchVolumesIDParamsBody) validateIops(formats strfmt.Registry) error {

	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	// value enum
	if err := m.validateIopsEnum("iops", "body", m.Iops); err != nil {
		return err
	}

	return nil
}

func (m *PatchVolumesIDParamsBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PatchVolumesIDParamsBody) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchVolumesIDParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchVolumesIDParamsBody) UnmarshalBinary(b []byte) error {
	var res PatchVolumesIDParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}