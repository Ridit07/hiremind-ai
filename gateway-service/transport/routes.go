package http

import (
	"net/http"

	server "gateway-service/server/authserver"

	authhttp "gateway-service/transport/authtransport"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(service server.GatewayServiceInterface) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(LoggingMiddleware)

	handlers := authhttp.NewAuthHTTPHandlers(service)

	router.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/signup", handlers.Signup)
		r.Post("/login", handlers.Login)
		r.Post("/refresh", handlers.RefreshToken)
		r.Post("/logout", handlers.Logout)
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	return router
}
