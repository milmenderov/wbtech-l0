package handlers

import (
	"html/template"
	"net/http"
	wilb_l0 "wbtech-l0"
	"wbtech-l0/pkg/http/service"
)

type Handler struct {
	services     *service.Service
	CacheStorage map[string]wilb_l0.Order
}

func NewHandler(services *service.Service, cacheStorage map[string]wilb_l0.Order) *Handler {
	return &Handler{
		services:     services,
		CacheStorage: cacheStorage,
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/get_order", h.GetOrderByID)
	return mux
}
