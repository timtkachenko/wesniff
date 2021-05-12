// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wesniff/models"
)

// CallbackHandleOKCode is the HTTP code returned for type CallbackHandleOK
const CallbackHandleOKCode int = 200

/*CallbackHandleOK Success

swagger:response callbackHandleOK
*/
type CallbackHandleOK struct {

	/*
	  In: Body
	*/
	Payload *CallbackHandleOKBody `json:"body,omitempty"`
}

// NewCallbackHandleOK creates CallbackHandleOK with default headers values
func NewCallbackHandleOK() *CallbackHandleOK {

	return &CallbackHandleOK{}
}

// WithPayload adds the payload to the callback handle o k response
func (o *CallbackHandleOK) WithPayload(payload *CallbackHandleOKBody) *CallbackHandleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the callback handle o k response
func (o *CallbackHandleOK) SetPayload(payload *CallbackHandleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CallbackHandleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CallbackHandleDefault Error

swagger:response callbackHandleDefault
*/
type CallbackHandleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCallbackHandleDefault creates CallbackHandleDefault with default headers values
func NewCallbackHandleDefault(code int) *CallbackHandleDefault {
	if code <= 0 {
		code = 500
	}

	return &CallbackHandleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the callback handle default response
func (o *CallbackHandleDefault) WithStatusCode(code int) *CallbackHandleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the callback handle default response
func (o *CallbackHandleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the callback handle default response
func (o *CallbackHandleDefault) WithPayload(payload *models.Error) *CallbackHandleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the callback handle default response
func (o *CallbackHandleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CallbackHandleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}