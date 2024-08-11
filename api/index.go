package api

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("id")

	_, _ = w.Write([]byte(sayHello(name)))
}

func sayHello(name string) string {
	return "Hello " + name + "!"
}
