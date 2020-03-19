package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	body := `{"error": "` + err.Error() + `"}`
	io.WriteString(w, body)
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, v interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(v)
}
