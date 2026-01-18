package handler

import (
	"encoding/json"
	"net/http"

	"github.com/esmrzv/go_shop/internal/auth"
	"github.com/esmrzv/go_shop/internal/service"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler{
	return &ProductHandler{
		service: s,
	}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request){
	userID := r.Context().Value(auth.UserContextKey).(int)
	_ = userID

	var req struct{
		Name string `json:"name"`
		Price float64 `json:"price"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	if err := h.service.Create(r.Context(), req.Name, req.Price); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request){
	products, err := h.service.List(r.Context())
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}