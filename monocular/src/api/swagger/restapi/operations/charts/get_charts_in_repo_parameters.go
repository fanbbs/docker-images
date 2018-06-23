package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetChartsInRepoParams creates a new GetChartsInRepoParams object
// with the default values initialized.
func NewGetChartsInRepoParams() GetChartsInRepoParams {
	var ()
	return GetChartsInRepoParams{}
}

// GetChartsInRepoParams contains all the bound params for the get charts in repo operation
// typically these are obtained from a http.Request
//
// swagger:parameters getChartsInRepo
type GetChartsInRepoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	Repo string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetChartsInRepoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rRepo, rhkRepo, _ := route.Params.GetOK("repo")
	if err := o.bindRepo(rRepo, rhkRepo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetChartsInRepoParams) bindRepo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Repo = raw

	return nil
}
