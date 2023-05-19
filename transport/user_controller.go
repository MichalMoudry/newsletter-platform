package transport

import (
	"net/http"
	"newsletter-platform/transport/model/contracts"
	"newsletter-platform/transport/model/dto"
	"newsletter-platform/transport/util"
)

// Method for handling user registration requests.
func (handler *Handler) RegisterUser(writer http.ResponseWriter, request *http.Request) {
	var registerRequest contracts.RegisterUserRequest
	err := util.UnmarshallRequest(request, &registerRequest)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
	}

	err = handler.Services.UserService.CreateUser(request.Context(), dto.ToNewUserData(registerRequest))
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
	}

	util.WriteResponse(writer, http.StatusCreated, registerRequest)
}
