package main

import (
	"encoding/json"
	"net/http"
)

func renderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
