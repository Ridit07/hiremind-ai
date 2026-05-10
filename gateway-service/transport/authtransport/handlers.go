package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"gateway-service/errors"
	server "gateway-service/server/authserver"
)

type HTTPHandlers struct {
	service server.GatewayServiceInterface
}

func NewAuthHTTPHandlers(service server.GatewayServiceInterface) *HTTPHandlers {
	return &HTTPHandlers{
		service: service,
	}
}

func (h *HTTPHandlers) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req server.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	resp, err := h.service.Signup(ctx, &req)
	if err != nil {
		if appErr, ok := errors.IsAppError(err); ok {
			respondWithError(w, appErr.Code, appErr.Message)
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}

func (h *HTTPHandlers) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req server.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	resp, err := h.service.Login(ctx, &req)
	if err != nil {
		if appErr, ok := errors.IsAppError(err); ok {
			respondWithError(w, appErr.Code, appErr.Message)
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}

func (h *HTTPHandlers) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	resp, err := h.service.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		if appErr, ok := errors.IsAppError(err); ok {
			respondWithError(w, appErr.Code, appErr.Message)
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}

func (h *HTTPHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	resp, err := h.service.Logout(ctx, req.RefreshToken)
	if err != nil {
		if appErr, ok := errors.IsAppError(err); ok {
			respondWithError(w, appErr.Code, appErr.Message)
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}
