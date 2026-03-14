// Package httpres contains HTTP response helpers.
package httpres

import (
	"encoding/json"
	"net/http"
)

func ConvertToJSON(w http.ResponseWriter, response any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
