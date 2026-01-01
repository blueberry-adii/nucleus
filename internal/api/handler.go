package api

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Healthy and Working"))
}
