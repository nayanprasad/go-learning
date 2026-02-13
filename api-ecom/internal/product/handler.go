package product

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	p := []string{"mobile", "tws"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
