package transport

import (
	"net/http"
	"newsletter-platform/transport/model/contracts"
	"newsletter-platform/transport/util"
)

// Method for handling requests for generating a new password reset token.
func (handler *Handler) GenerateNewPassResetToken(writer http.ResponseWriter, request *http.Request) {
	var requestData contracts.PassResetTokenRequestData
	err := util.UnmarshallRequest(request, &requestData)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = handler.Services.PassResetService.GenerateNewToken(request.Context(), requestData.Email)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, nil)
}

// Method for handling requests for changing user's password.
func (handler *Handler) ResetUsersPassword(writer http.ResponseWriter, request *http.Request) {
	var requestData contracts.PassResetRequestData
	err := util.UnmarshallRequest(request, &requestData)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = handler.Services.PassResetService.ResetPassword(
		request.Context(),
		requestData.Email,
		requestData.Password,
		requestData.TokenId,
	)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, nil)
}
