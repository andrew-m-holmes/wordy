package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func handlerHome(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received the following request: %v", r.Body)
	_, err := fmt.Fprintf(w, "Home page")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	log.Printf("HTTP server starting on localhost%v", portNumber)
	http.HandleFunc("/", handlerHome)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal(err)
	}
}
