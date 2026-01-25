package routers
import (
	"net/http"
	"github.com/esmrzv/go_shop/internal/auth"
	"github.com/esmrzv/go_shop/internal/http/handler"
)


func RegisterOrderRouter(
	mux *http.ServeMux,
	h *handler.OrderHandler,
	jwt string,
) {
	mux.Handle(
		"/orders",
		auth.JWTMiddleware(jwt, http.HandlerFunc(h.Create)),
	)
}