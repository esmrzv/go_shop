package routers

import (
	"github.com/esmrzv/go_shop/internal/http/handler"
	"net/http"
)

func RegisterAuthRouter(mux *http.ServeMux, h *handler.AuthHandler) {
	mux.HandleFunc("/register", h.Register)
	mux.HandleFunc("/login", h.Login)
}