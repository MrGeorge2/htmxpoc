package routerutils

import "net/http"

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bad request"))
	w.WriteHeader(http.StatusBadRequest)
}

func NotAlowed(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not allowed"))
	w.WriteHeader(http.StatusMethodNotAllowed)
}
