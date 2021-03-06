// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicCurrenciesResponse public currencies response
// swagger:model public_currencies_response
type PublicCurrenciesResponse struct {

	// result
	// Required: true
	Result []Currency `json:"result"`
}

// Validate validates this public currencies response
func (m *PublicCurrenciesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicCurrenciesResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	for i := 0; i < len(m.Result); i++ {

		if err := m.Result[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicCurrenciesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicCurrenciesResponse) UnmarshalBinary(b []byte) error {
	var res PublicCurrenciesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
