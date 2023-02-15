package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/awesomeProject/internal/models"
	"github.com/awesomeProject/internal/service"
)

type handler struct {
}

// func (handler *handler)
func Post(w http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	var st models.Student

	json.Unmarshal(data, &st)
	st, er := service.Post(st)
	if er != "" {
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(st)

}
