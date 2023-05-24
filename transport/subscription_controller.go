package transport

import (
	"net/http"
	"newsletter-platform/transport/model/contracts"
	"newsletter-platform/transport/util"

	"github.com/google/uuid"
)

// Method for handling requests for registering new subscriptions.
func (handler *Handler) RegisterSubscription(writer http.ResponseWriter, request *http.Request) {
	var requestData contracts.RegisterSubscriptionRequestData
	if err := util.UnmarshallRequest(request, &requestData); err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	newsletterId, err := uuid.Parse(requestData.NewsletterId)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = handler.Services.SubscriptionService.CreateSubscription(request.Context(), requestData.Email, newsletterId)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusCreated, nil)
}

// Method for handling request for unsubscribing from a newsletter.
func (handler *Handler) NewsletterUnsubscribe(w http.ResponseWriter, r *http.Request) {
	var requestData contracts.CancelSubscriptionRequestData
	if err := util.UnmarshallRequest(r, &requestData); err != nil {
		return
	}

	newsletterId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	err = handler.Services.SubscriptionService.CancelSubscription(r.Context(), requestData.Email, newsletterId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, nil)
}
