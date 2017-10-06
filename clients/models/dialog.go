package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Dialog Dialog element describes the request for parameter specification
// swagger:model Dialog
type Dialog struct {
	XMLName xml.Name `xml:"http://www.sap.com/lmsl/slp Dialog"`

	// description
	// Required: true
	Description *string `xml:"description"`

	// display name
	// Required: true
	DisplayName *string `xml:"displayName"`

	// id
	// Required: true
	ID *string `xml:"id"`

	// metadialog
	Metadialog strfmt.URI `xml:"metadialog,omitempty"`
}

// Validate validates this dialog
func (m *Dialog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dialog) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Dialog) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *Dialog) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}