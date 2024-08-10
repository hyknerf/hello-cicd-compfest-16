package main

import (
	"net/http"

	"github.com/hyknerf/hello-cicd-compfest-16/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/{id}", api.Handler)

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
