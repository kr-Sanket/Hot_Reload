package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello version 4 | time: %v", time.Now())
	})

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}