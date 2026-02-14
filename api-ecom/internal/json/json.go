package json

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Read(r *http.Request, data any) error {
	if r.Body == nil || r.ContentLength == 0 {
		return errors.New("request body is empty")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}
