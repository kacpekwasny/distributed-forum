package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func GetQueryParamDefault(r *http.Request, name string, def string) string {
	v := r.URL.Query().Get(name)
	if v == "" {
		return def
	}
	return v
}

func GetQueryParamParser[T any](r *http.Request, name string, def T, parse func(v string) (T, error)) T {
	parsed, err := parse(r.URL.Query().Get(name))
	if err != nil {
		return def
	}
	return parsed
}

func GetQueryParamInt(r *http.Request, name string, def int64) int64 {
	return GetQueryParamParser(r, name, def,
		func(v string) (int64, error) {
			return strconv.ParseInt(v, 10, 32)
		})
}
