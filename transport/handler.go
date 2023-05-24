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

		r.Route("/newsletters", func(r chi.Router) {
			r.Post("/", handler.CreateNewsletter)
			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/posts", handler.GetNewsletterPosts)
				r.Put("/", handler.UpdateNewsletter)
				r.Delete("/", handler.DeleteNewsletter)
			})
		})

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", handler.SendPost)
			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/", handler.GetPost)
				r.Delete("/", handler.DeletePost)
			})
		})
	})

	// Public routes
	handler.Mux.Get("/health", health)
	handler.Mux.Post("/register", handler.RegisterUser)
	handler.Mux.Post("/login", handler.Login)
	handler.Mux.Post("/passwordresetrequests", handler.GenerateNewPassResetToken)
	handler.Mux.Patch("/.well-known/change-password", handler.ResetUsersPassword)

	handler.Mux.Route("/subscriptions", func(r chi.Router) {
		r.Post("/", handler.RegisterSubscription)
		r.Delete("/{uuid}", handler.NewsletterUnsubscribe)
	})

	handler.Mux.Get("/newsletters/{uuid}", handler.GetNewsletter)

	return handler
}

// Controller endpoint function for handling requests on /health.
func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}
