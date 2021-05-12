// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wesniff/models"
)

// GetUsersHandlerFunc turns a function with the right signature into a get users handler
type GetUsersHandlerFunc func(GetUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersHandlerFunc) Handle(params GetUsersParams) middleware.Responder {
	return fn(params)
}

// GetUsersHandler interface for that can handle valid get users params
type GetUsersHandler interface {
	Handle(GetUsersParams) middleware.Responder
}

// NewGetUsers creates a new http.Handler for the get users operation
func NewGetUsers(ctx *middleware.Context, handler GetUsersHandler) *GetUsers {
	return &GetUsers{Context: ctx, Handler: handler}
}

/*GetUsers swagger:route GET /users users getUsers

get Users

*/
type GetUsers struct {
	Context *middleware.Context
	Handler GetUsersHandler
}

func (o *GetUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetUsersOKBody get users o k body
//
// swagger:model GetUsersOKBody
type GetUsersOKBody struct {

	// code
	Code models.SuccessCode `json:"code,omitempty"`

	// data
	Data []*models.UserDto `json:"data"`
}

// Validate validates this get users o k body
func (o *GetUsersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUsersOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersOKBody) UnmarshalBinary(b []byte) error {
	var res GetUsersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
