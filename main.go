package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Endpoint requested: %s\n", r.URL.Path)
	})
	fmt.Println("Starting server on port 8000 ...")
	http.ListenAndServe(":8000", nil)
}
