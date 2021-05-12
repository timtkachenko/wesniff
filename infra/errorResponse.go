package infra

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"wesniff/models"
)

func ErrorResponse(err error) middleware.Responder {
	resp := &ResponderFunc{
		Payload: &models.Error{
			Message: err.Error(),
		},
	}
	if apierr, ok := err.(errors.Error); ok {
		resp.Code = int(apierr.Code())
	}
	return resp
}

// ResponderFunc wraps a func as a Responder interface
type ResponderFunc struct {
	Code    int
	Payload *models.Error `json:"body,omitempty"`
}

// WriteResponse writes to the response
func (fn ResponderFunc) WriteResponse(rw http.ResponseWriter, pr runtime.Producer) {
	if fn.Code != 0 {
		rw.WriteHeader(fn.Code)

	} else {
		rw.WriteHeader(http.StatusBadRequest)
	}
	if fn.Payload != nil {
		payload := fn.Payload
		if err := pr.Produce(rw, payload); err != nil {
			panic(err)
		}
	}
}
