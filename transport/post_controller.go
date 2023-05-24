package transport

import (
	"net/http"
	"newsletter-platform/service/model"
)

// Method for handling requests for creating new posts.
func (handler *Handler) SendPost(w http.ResponseWriter, r *http.Request) {
	handler.Services.PostService.CreateNewPost(r.Context(), model.PostCreateModel{})
}
