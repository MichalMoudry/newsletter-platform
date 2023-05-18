package transport

import (
	"net/http"
	"newsletter-platform/transport/model/contracts"
	"newsletter-platform/transport/util"
)

// Method for handling user registration requests.
func (handler *Handler) RegisterUser(writer http.ResponseWriter, request *http.Request) {
	var userDto contracts.RegisterUserRequest
	err := util.UnmarshallRequest(request, &userDto)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
	}

	err = handler.UserService.CreateUser(request.Context())
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
	}

	util.WriteResponse(writer, http.StatusCreated, userDto)
}
