// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListFilesHandlerFunc turns a function with the right signature into a list files handler
type ListFilesHandlerFunc func(ListFilesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListFilesHandlerFunc) Handle(params ListFilesParams) middleware.Responder {
	return fn(params)
}

// ListFilesHandler interface for that can handle valid list files params
type ListFilesHandler interface {
	Handle(ListFilesParams) middleware.Responder
}

// NewListFiles creates a new http.Handler for the list files operation
func NewListFiles(ctx *middleware.Context, handler ListFilesHandler) *ListFiles {
	return &ListFiles{Context: ctx, Handler: handler}
}

/*
	ListFiles swagger:route GET /files listFiles

Get a list of all files
*/
type ListFiles struct {
	Context *middleware.Context
	Handler ListFilesHandler
}

func (o *ListFiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListFilesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListFilesOKBodyItems0 list files o k body items0
//
// swagger:model ListFilesOKBodyItems0
type ListFilesOKBodyItems0 struct {

	// The ID of the file
	ID string `json:"id,omitempty"`

	// The URL of the file
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this list files o k body items0
func (o *ListFilesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListFilesOKBodyItems0) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(o.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", o.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list files o k body items0 based on context it is used
func (o *ListFilesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListFilesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListFilesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res ListFilesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
