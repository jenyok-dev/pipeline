package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func getMessage() string {
	msg := os.Getenv("APP_MESSAGE")
	if msg == "" {
		return "Hello, World!"
	}
	return msg
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: getMessage(),
		Status:  "running",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthHandler)

	port := ":8080"
	log.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
