package routers

import (
	"net/http"

	"github.com/esmrzv/go_shop/internal/auth"
	"github.com/esmrzv/go_shop/internal/http/handler"
)

func RegisterProductRouter(mux *http.ServeMux, h *handler.ProductHandler, secret string){
	mux.Handle("/products", auth.JWTMiddleware(secret, http.HandlerFunc(h.Create)),
	)
	mux.Handle("/products/list", auth.JWTMiddleware(secret, http.HandlerFunc(h.List)),)
}