package routers

import (
	"net/http"

	"github.com/esmrzv/go_shop/internal/auth"
	"github.com/esmrzv/go_shop/internal/http/handler"

)

func RegisterCartRouter(
	mux *http.ServeMux,
	h *handler.CartHandler,
	jwtSecret string,
) {
	mux.Handle(
		"/cart/add",
		auth.JWTMiddleware(jwtSecret, http.HandlerFunc(h.Add)),
	)

	mux.Handle(
		"/cart",
		auth.JWTMiddleware(jwtSecret, http.HandlerFunc(h.Get)),
	)
}