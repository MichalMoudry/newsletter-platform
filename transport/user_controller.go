package transport

import (
	"net/http"
	"newsletter-platform/transport/model/contracts"
	"newsletter-platform/transport/model/dto"
	"newsletter-platform/transport/util"

	"github.com/go-chi/chi/v5"
)

// Function for obtaining an email from the request URL.
func getEmailFromURL(r *http.Request) (string, error) {
	email := chi.URLParam(r, "email")
	//TODO: Add email validation
	return email, nil
}

// Method for handling user registration requests.
func (handler *Handler) RegisterUser(writer http.ResponseWriter, request *http.Request) {
	var registerRequest contracts.RegisterUserRequestData
	err := util.UnmarshallRequest(request, &registerRequest)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = handler.Services.UserService.CreateUser(request.Context(), dto.ToNewUserData(registerRequest))
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusCreated, registerRequest)
}

// Method for handling login requests.
func (handler *Handler) Login(writer http.ResponseWriter, request *http.Request) {
	var requestData contracts.LoginRequestData
	err := util.UnmarshallRequest(request, &requestData)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	token, err := handler.Services.UserService.ValidateLogin(
		request.Context(),
		requestData.Email,
		requestData.Password,
	)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, token)
}

// Method for handling user's requests for their data.
func (handler *Handler) GetUserData(writer http.ResponseWriter, request *http.Request) {
	email, err := getEmailFromURL(request)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	data, err := handler.Services.UserService.GetUser(request.Context(), email)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, data)
}

// Method for handling user's requests to delete their accounts.
func (handler *Handler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	var requestData contracts.DeleteUserRequestData
	err := util.UnmarshallRequest(request, &requestData)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	email, err := getEmailFromURL(request)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	err = handler.Services.UserService.DeleteUser(request.Context(), email, requestData.ConcurrencyStamp)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, nil)
}

// Method for handling request for updataing user's general information.
func (handler *Handler) UpdateUserInfo(writer http.ResponseWriter, request *http.Request) {
	email, err := getEmailFromURL(request)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	var requestData contracts.UpdateUserRequestData
	if err = util.UnmarshallRequest(request, &requestData); err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = handler.Services.UserService.UpdateUsersInfo(
		request.Context(),
		email,
		requestData.UserName,
		requestData.ConcurrencyStamp,
	)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(writer, http.StatusOK, nil)
}
