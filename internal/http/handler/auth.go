package handler

import (
	"encoding/json"
	"net/http"

	"github.com/esmrzv/go_shop/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler{
	return &AuthHandler{
		service: s,
	}
}
func (h *AuthHandler) Register (w http.ResponseWriter, r *http.Request) {
	var req struct{
		Email  string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	if err := h.service.Register(r.Context(), req.Email, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}


func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
	var req struct{
		Email  string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	token, err := h.service.Login(r.Context(), req.Email, req.Password)
	if err != nil{
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}