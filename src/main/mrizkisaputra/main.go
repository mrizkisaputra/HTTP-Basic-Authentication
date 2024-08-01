package main

import (
	. "basic_authentication/src/main/mrizkisaputra/controllers"
	"basic_authentication/src/main/mrizkisaputra/utils"
	"fmt"
	"net/http"
)

var mux = http.DefaultServeMux

func main() {
	mux.HandleFunc("/contacts", GetContactHandler)

	var handler http.Handler = mux
	handler = utils.MiddlewareAuth(handler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("cannot start server: %w", err))
	}
}
