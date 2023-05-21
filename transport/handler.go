package transport

import (
	"net/http"
	"newsletter-platform/transport/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type Handler struct {
	Port     int
	Mux      *chi.Mux
	Services model.ServiceCollection
}

// Initializer function for HTTP handler.
func Initialize(port int, tokenAuth *jwtauth.JWTAuth, services model.ServiceCollection) *Handler {
	handler := &Handler{
		Port:     port,
		Mux:      chi.NewRouter(),
		Services: services,
	}
	handler.Mux.Use(middleware.Logger)

	// Protected routes
	handler.Mux.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/users", func(r chi.Router) {
			r.Route("/{email}", func(r chi.Router) {
				r.Get("/", handler.GetUserData)
				r.Delete("/", handler.DeleteUser)
				r.Put("/", handler.UpdateUserInfo)
			})
		})
	})

	// Public routes
	handler.Mux.Get("/health", health)
	handler.Mux.Post("/register", handler.RegisterUser)
	handler.Mux.Post("/login", handler.Login)
	handler.Mux.Post("/passwordresetrequests", handler.GenerateNewPassResetToken)
	handler.Mux.Patch("/.well-known/change-password", handler.ResetUsersPassword)

	return handler
}

// Controller endpoint function for handling requests on /health.
func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}
