package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonWithStatus(w http.ResponseWriter, v any, code int) error {
	w.WriteHeader(code)
	return WriteJson(w, v)
}

// Set Content-Type: application/json.
// Marshal the `v` and write it.
//
// Remember to set `w.WriteHeader(200, 201, etc.)` before calling this function.
func WriteJson(w http.ResponseWriter, marshallable any) error {
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(marshallable)

	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(jsonResp)
	return nil
}
