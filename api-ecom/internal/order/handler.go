package order

import (
	"api-ecom/internal/json"
	"log/slog"
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

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var tempOrder placeOrderPrams
	if err := json.Read(r, &tempOrder); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.service.PlaceOrder(r.Context(), tempOrder)

	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusAccepted, order)
}
