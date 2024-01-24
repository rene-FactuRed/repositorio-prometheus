package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "HOLA MUNDO")
	})

	http.ListenAndServe("localhost:8080", nil)
}
