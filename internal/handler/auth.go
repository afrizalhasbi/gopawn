package handler

import (
	"encoding/json"
	"fmt"
	"gopawn/internal/data/payload"
	"gopawn/internal/service"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var reg payload.Register

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.Register(&reg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login payload.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwtToken, err := h.Service.Login(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"token": "%s"}`, jwtToken)
	}
}

func (h *AuthHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var reg payload.Delete

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.Delete(&reg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var reg payload.ResetPassword
	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.ResetPassword(&reg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/auth/register":
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.Register(w, r)

	case "/auth/login":
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.Login(w, r)

	default:
		http.NotFound(w, r)
	}
}
