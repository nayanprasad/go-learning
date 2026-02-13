package product

import (
	"api-ecom/internal/json"
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

	json.Write(w, http.StatusOK, p)
}
