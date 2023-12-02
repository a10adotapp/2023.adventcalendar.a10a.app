package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintf(w, "{\"message\":\"hello from %s\"}\n", r.Host)
	}))
}
