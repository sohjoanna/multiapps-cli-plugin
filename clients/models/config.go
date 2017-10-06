package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Config config
// swagger:model config
type Config struct {
	XMLName    xml.Name     `xml:"http://www.sap.com/lmsl/slp config"`
	Parameters []*Parameter `xml:"Parameter,omitempty"`
}

// Validate validates this config
func (m Config) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m.Parameters); i++ {

		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {

			if err := m.Parameters[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
