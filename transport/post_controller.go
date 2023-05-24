package transport

import (
	"net/http"
	"newsletter-platform/service/model"
	"newsletter-platform/transport/model/contracts"
	"newsletter-platform/transport/util"

	"github.com/google/uuid"
)

// Method for handling requests for creating new posts.
func (handler *Handler) SendPost(w http.ResponseWriter, r *http.Request) {
	var requestData contracts.CreatePostRequestData
	err := util.UnmarshallRequest(r, &requestData)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	newsletterId, err := uuid.Parse(requestData.NewsletterId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	postId, err := handler.Services.PostService.CreateNewPost(r.Context(), model.PostCreateModel{
		Title:        requestData.Title,
		Content:      requestData.Content,
		NewsletterId: newsletterId,
	})
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusCreated, contracts.CreatePostResponse{
		PostId:    postId,
		PostTitle: requestData.Title,
	})
}

// Method for handling request to obtain information about a post.
func (handler *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	postId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	data, err := handler.Services.PostService.GetPost(r.Context(), postId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, data)
}
