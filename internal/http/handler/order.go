package handler

import (
	"net/http"
	"github.com/esmrzv/go_shop/internal/service"

)
type OrderHandler struct {
	service service.OrderService
}


func NewOrderHandler(s service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: s,
	}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	if err := h.service.Create(r.Context(), userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
