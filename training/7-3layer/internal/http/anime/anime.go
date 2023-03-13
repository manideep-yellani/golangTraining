package anime

import (
	"encoding/json"
	"net/http"
	"training/7-3layer/internal/services"
)

var table = "MOVIES"

type handler struct {
	service services.Animeservicer
}

func New(service services.Animeservicer) *handler {
	return &handler{service}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("fdf")
}
