package routers

import (
	"net/http"

	"github.com/esmrzv/go_shop/internal/auth"
	"github.com/esmrzv/go_shop/internal/http/handler"
)


func RegisterCategoryRouter(mux *http.ServeMux, h * handler.CategoryHandler, secret string){
	mux.Handle(
		"/categories",
		auth.JWTMiddleware(secret, 
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			switch r.Method{
			case http.MethodPost:
				h.CreateCategory(w, r)
			case http.MethodGet:
				h.ListCategories(w, r)
			default:
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		})),
	)
}