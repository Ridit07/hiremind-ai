package http

import (
	"net/http"

	authserver "gateway-service/server/authserver"
	interviewserver "gateway-service/server/interviewserver"

	authhttp "gateway-service/transport/authtransport"
	interviewhttp "gateway-service/transport/interviewtransport"

	common "gateway-service/transport/common"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(authSvc authserver.GatewayServiceInterface, interviewSvc interviewserver.InterviewServiceInterface) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(common.CORSMiddleware)
	router.Use(common.LoggingMiddleware)

	authHandlers := authhttp.NewAuthHTTPHandlers(authSvc)
	interviewHandlers := interviewhttp.NewInterviewHTTPHandlers(interviewSvc)

	router.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/signup", authHandlers.Signup)
		r.Post("/login", authHandlers.Login)
		r.Post("/refresh", authHandlers.RefreshToken)
		r.Post("/logout", authHandlers.Logout)
	})

	router.Route("/api/v1/interviews", func(r chi.Router) {
		r.Use(common.AuthMiddleware)
		r.Get("/", interviewHandlers.GetInterviews)
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	return router
}
