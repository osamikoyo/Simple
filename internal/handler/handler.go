package handler

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PONG")
}

func Main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from main handler")
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", Main)
	mux.HandleFunc("/ping", Ping)
}