package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type WordHandler struct {
	http.Handler
}

const port int = 8080

func (h *WordHandler) serveHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var word Word

		if err := json.NewDecoder(r.Body).Decode(&word); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
		} else {
			word.DateAdded = time.Now()
			word.DateReviewed = time.Now()
			word.Understanding = 1

			// TODO: add to DB
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(word)
		}

	} else if r.Method == http.MethodGet {
		w.Write([]byte("Word retrieval not yet implemented"))

	} else {
		http.Error(w, "Method not implemented", http.StatusMethodNotAllowed)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Wordy server running"))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Wordy homepage"))
}

func main() {

	wordHandler := &WordHandler{}
	http.HandleFunc("/home", homepage)
	http.HandleFunc("/health", healthCheck)
	http.Handle("/api/words", wordHandler)

	log.Printf("Starting Wordy server on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Server error:", err)
	}

}
