// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wesniff/models"
)

// CallbackHandleHandlerFunc turns a function with the right signature into a callback handle handler
type CallbackHandleHandlerFunc func(CallbackHandleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CallbackHandleHandlerFunc) Handle(params CallbackHandleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CallbackHandleHandler interface for that can handle valid callback handle params
type CallbackHandleHandler interface {
	Handle(CallbackHandleParams, interface{}) middleware.Responder
}

// NewCallbackHandle creates a new http.Handler for the callback handle operation
func NewCallbackHandle(ctx *middleware.Context, handler CallbackHandleHandler) *CallbackHandle {
	return &CallbackHandle{Context: ctx, Handler: handler}
}

/*CallbackHandle swagger:route POST /events events callbackHandle

accept callbacks

*/
type CallbackHandle struct {
	Context *middleware.Context
	Handler CallbackHandleHandler
}

func (o *CallbackHandle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCallbackHandleParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CallbackHandleOKBody callback handle o k body
//
// swagger:model CallbackHandleOKBody
type CallbackHandleOKBody struct {

	// code
	Code models.SuccessCode `json:"code,omitempty"`

	// data
	Data *CallbackHandleOKBodyData `json:"data,omitempty"`
}

// Validate validates this callback handle o k body
func (o *CallbackHandleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CallbackHandleOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callbackHandleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CallbackHandleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CallbackHandleOKBody) UnmarshalBinary(b []byte) error {
	var res CallbackHandleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CallbackHandleOKBodyData callback handle o k body data
//
// swagger:model CallbackHandleOKBodyData
type CallbackHandleOKBodyData struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this callback handle o k body data
func (o *CallbackHandleOKBodyData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CallbackHandleOKBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CallbackHandleOKBodyData) UnmarshalBinary(b []byte) error {
	var res CallbackHandleOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
