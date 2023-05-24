package transport

import (
	"net/http"
	"newsletter-platform/service/model"
	"newsletter-platform/transport/model/contracts"
	"newsletter-platform/transport/util"

	"github.com/google/uuid"
)

// Method for handling requests for creating a new newsletter.
func (handler *Handler) CreateNewsletter(writer http.ResponseWriter, request *http.Request) {
	var requestData contracts.CreateNewsletterRequestData
	err := util.UnmarshallRequest(request, &requestData)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	id, err := handler.Services.NewsletterService.CreateNewsletter(
		request.Context(),
		requestData.Name,
		requestData.Description,
	)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(writer, http.StatusCreated, contracts.CreateNewsletterResponse{
		Id:          id,
		Name:        requestData.Name,
		Description: requestData.Description,
	})
}

// Method for handling requests for obtaining public newsletter data.
func (handler *Handler) GetNewsletter(writer http.ResponseWriter, request *http.Request) {
	newsletterId, err := util.GetUuidFromUrl(request)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	data, err := handler.Services.NewsletterService.GetNewsletter(request.Context(), newsletterId)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, data)
}

// Method for handling requests for updating a specific newsletter.
func (handler *Handler) UpdateNewsletter(writer http.ResponseWriter, request *http.Request) {
	newsletterId, err := util.GetUuidFromUrl(request)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	var requestData contracts.UpdateNewsletterRequestData
	err = util.UnmarshallRequest(request, &requestData)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	oldConcurrencyStamp, err := uuid.Parse(requestData.ConcurrencyStamp)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = handler.Services.NewsletterService.UpdateNewsletter(request.Context(), model.NewsletterUpdateModel{
		NewsletterId:        newsletterId,
		Name:                requestData.Name,
		Description:         requestData.Description,
		OldConcurrencyStamp: oldConcurrencyStamp,
	})
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, nil)
}

// Method for handling requests for deleting a specific newsletter.
func (handler *Handler) DeleteNewsletter(writer http.ResponseWriter, request *http.Request) {
	newsletterId, err := util.GetUuidFromUrl(request)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = handler.Services.NewsletterService.DeleteNewsletter(request.Context(), newsletterId)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, nil)
}

// Method for handling requests for obtaining all posts of a newsletter
func (handler *Handler) GetNewsletterPosts(w http.ResponseWriter, r *http.Request) {
	newsletterId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	data, err := handler.Services.NewsletterService.GetPosts(r.Context(), newsletterId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, data)
}
