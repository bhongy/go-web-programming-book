package api

import "net/http"

func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/account/create", createAccount)
	mux.HandleFunc("/authenticate", authenticate)
	return mux
}
