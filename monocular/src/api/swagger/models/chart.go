package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Chart chart

swagger:model chart
*/
type Chart struct {

	/* description

	Required: true
	Min Length: 1
	*/
	Description *string `json:"description"`

	/* home

	Required: true
	Min Length: 1
	*/
	Home *string `json:"home"`

	/* keywords
	 */
	Keywords []string `json:"keywords,omitempty"`

	/* maintainers

	Required: true
	*/
	Maintainers []*Maintainer `json:"maintainers"`

	/* name

	Required: true
	Min Length: 1
	*/
	Name *string `json:"name"`

	/* repo

	Required: true
	*/
	Repo *Repo `json:"repo"`

	/* repo URL
	 */
	RepoURL string `json:"repoURL,omitempty"`

	/* sources

	Required: true
	*/
	Sources []string `json:"sources"`
}

// Validate validates this chart
func (m *Chart) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHome(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKeywords(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaintainers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Chart) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *Chart) validateHome(formats strfmt.Registry) error {

	if err := validate.Required("home", "body", m.Home); err != nil {
		return err
	}

	if err := validate.MinLength("home", "body", string(*m.Home), 1); err != nil {
		return err
	}

	return nil
}

func (m *Chart) validateKeywords(formats strfmt.Registry) error {

	if swag.IsZero(m.Keywords) { // not required
		return nil
	}

	return nil
}

func (m *Chart) validateMaintainers(formats strfmt.Registry) error {

	if err := validate.Required("maintainers", "body", m.Maintainers); err != nil {
		return err
	}

	for i := 0; i < len(m.Maintainers); i++ {

		if swag.IsZero(m.Maintainers[i]) { // not required
			continue
		}

		if m.Maintainers[i] != nil {

			if err := m.Maintainers[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Chart) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *Chart) validateRepo(formats strfmt.Registry) error {

	if err := validate.Required("repo", "body", m.Repo); err != nil {
		return err
	}

	if m.Repo != nil {

		if err := m.Repo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Chart) validateSources(formats strfmt.Registry) error {

	if err := validate.Required("sources", "body", m.Sources); err != nil {
		return err
	}

	return nil
}
