package main

import (
	"log"
	"net/http"
	"os"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handleRequest)
	log.Println(http.ListenAndServe(":"+port, nil))
}
