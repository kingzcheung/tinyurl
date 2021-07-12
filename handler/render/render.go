package render

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	bytes, _ := json.Marshal(data)
	_, _ = w.Write(bytes)
}

func Error(w http.ResponseWriter, error string, code int) {
	Json(w, map[string]string{
		"error": error,
	}, code)
}
