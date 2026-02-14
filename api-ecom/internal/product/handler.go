package product

import (
	"api-ecom/internal/json"
	"log/slog"
	"net/http"

	repo "api-ecom/internal/adapters/psql/sqlc"
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

	p, err := h.service.ListProducts(r.Context())

	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, p)
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProductParams repo.CreateProductParams
	if err := json.Read(r, &createProductParams); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validate here

	p, err := h.service.CreateProduct(r.Context(), createProductParams)

	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, p)

}
