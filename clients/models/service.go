package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

type ServiceParameters struct {
	Parameters []*Parameter `xml:"Parameter,omitempty"`
}

// Service SLMP is a protocol for managing SL processes. It provides facilities for querying SL services available on managed object, accessing historical information associated with a finished processes and creation of new processes
// swagger:model Service
type Service struct {
	XMLName xml.Name `xml:"http://www.sap.com/lmsl/slp Service"`

	// description
	Description string `xml:"description,omitempty"`

	// display name
	DisplayName string `xml:"displayName,omitempty"`

	// external info
	ExternalInfo strfmt.URI `xml:"externalInfo,omitempty"`

	// files
	Files strfmt.URI `xml:"files,omitempty"`

	// id
	// Required: true
	ID *strfmt.URI `xml:"id"`

	// parameters
	Parameters ServiceParameters `xml:"parameters,omitempty"`

	// processes
	// Required: true
	Processes *strfmt.URI `xml:"processes"`

	// slppversion
	// Required: true
	Slppversion *string `xml:"slppversion"`

	// synchronous
	Synchronous bool `xml:"synchronous,omitempty"`

	// versions
	Versions strfmt.URI `xml:"versions,omitempty"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProcesses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSlppversion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Service) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters.Parameters); i++ {

		if swag.IsZero(m.Parameters.Parameters[i]) { // not required
			continue
		}

		if m.Parameters.Parameters[i] != nil {

			if err := m.Parameters.Parameters[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateProcesses(formats strfmt.Registry) error {

	if err := validate.Required("processes", "body", m.Processes); err != nil {
		return err
	}

	return nil
}

func (m *Service) validateSlppversion(formats strfmt.Registry) error {

	if err := validate.Required("slppversion", "body", m.Slppversion); err != nil {
		return err
	}

	return nil
}
