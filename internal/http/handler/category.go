package handler

import (
	"encoding/json"
	"net/http"

	"github.com/esmrzv/go_shop/internal/service"
)

type CategoryHandler struct {
	service service.CategoryService
}


func NewCategoryHandler( s service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request){
	type request struct {
		Name string `json:"name"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(r.Context(), req.Name)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}


func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request){
	categories, err := h.service.List(r.Context())
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)

}