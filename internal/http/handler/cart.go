package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/esmrzv/go_shop/internal/service"
)

type CartHandler struct {
	service service.CartService
}

func NewCartHandler(s service.CartService) *CartHandler {
	return &CartHandler{
		service: s,
	}
}

func (h *CartHandler) Add(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	productID, _ := strconv.Atoi(r.URL.Query().Get("product_id"))

	h.service.AddItem(userID, productID)
	w.WriteHeader(http.StatusOK)
}


func (h *CartHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	cart := h.service.GetCart(userID)
	if cart == nil {
		http.Error(w, "cart not found", http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(cart)

}