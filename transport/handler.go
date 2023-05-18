package transport

import (
	"net/http"
	"newsletter-platform/transport/model/ioc"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type Handler struct {
	Port             int
	Mux              *chi.Mux
	AuthService      ioc.IAuthService
	UserService      ioc.IUserService
	PassResetService ioc.IPassResetService
}

// Initializer function for HTTP handler.
func Initialize(port int) *Handler {
	handler := &Handler{
		Port: port,
		Mux:  chi.NewRouter(),
	}
	handler.Mux.Use(middleware.Logger)

	// Protected routes
	handler.Mux.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(handler.AuthService.GetJwtAuth()))
		r.Use(jwtauth.Authenticator)
	})

	// Public routes
	handler.Mux.Get("/health", health)

	return handler
}

// Controller endpoint function for handling requests on /health.
func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}
