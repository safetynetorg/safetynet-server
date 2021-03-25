package handlers

import "net/http"

func Cors(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
