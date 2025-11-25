package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, ".")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	addr := ":8080"
	log.Printf("Server starting on %s\n", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}