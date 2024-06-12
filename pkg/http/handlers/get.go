package handlers

import (
	"encoding/json"
	"net/http"
	"wbtech-l0/pkg/cache"
)

type OrderGetter interface {
	GetOrderHandler(OrderUID string) (string, error)
}

func (h *Handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("order_id")
	if id == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	order, err := cache.GetOrderFromCache(id, h.CacheStorage)
	if err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "failed to encode order to JSON", http.StatusInternalServerError)
	}
}
