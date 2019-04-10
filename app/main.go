package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Minikube! (v0.1)")
	})

	http.ListenAndServe(":5000", nil)
}
