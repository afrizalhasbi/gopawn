package handler

import (
	"encoding/json"
	"gopawn/internal/data/payload"
	"gopawn/internal/service"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var reg payload.Login

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Service.Register(&reg)
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login payload.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Service.Login(&login)

	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Login endpoint not implemented yet"))
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var reg payload.Reset

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Service.ResetPassword(&reg)
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
