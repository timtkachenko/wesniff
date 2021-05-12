package controllers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"wesniff/infra"
	"wesniff/lib"
	"wesniff/wesniff/operations/events"
	"wesniff/wesniff/operations/users"
)

type UserHandler struct {
	veriff lib.Veriff
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (ah *UserHandler) CreateUser(params users.CreateUserParams) middleware.Responder {
	session, err := ah.veriff.Create(params)
	if err != nil {
		return infra.ErrorResponse(err)
	}
	upload, err := ah.veriff.Upload(session, params.Image)
	if err != nil {
		return infra.ErrorResponse(err)
	}
	if upload != nil {
		fmt.Printf("upload %v/n", upload)
	}
	update, err := ah.veriff.Update(session)
	if err != nil {
		return infra.ErrorResponse(err)
	}
	if upload != nil {
		fmt.Printf("update %v/n", update)
	}
	return users.NewCreateUserOK().WithPayload(&users.CreateUserOKBody{
		Code: 1,
		Data: nil,
	})
}

func (ah *UserHandler) CallbackHandler(params events.CallbackHandleParams, principal interface{}) middleware.Responder {
	fmt.Printf("principal %v/n", principal)
	fmt.Printf("Callback %v/n", params.Body)

	return events.NewCallbackHandleOK().WithPayload(&events.CallbackHandleOKBody{
		Code: 1,
		Data: nil,
	})
}
