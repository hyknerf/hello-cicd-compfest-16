package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/{id}", helloHandler)

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("id")

	_, _ = w.Write([]byte(sayHello(name)))
}

func sayHello(name string) string {
	return "Hello " + name + "!"
}
