package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

type Word struct {
	Term         string
	PartOfSpeech string
	Definition   string
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	data := Word{
		Term:         "dog",
		PartOfSpeech: "noun",
		Definition:   "Four legged animal descending from wolves.",
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, data)
}

func main() {
	fmt.Printf("HTTP server listening on localhost%v\n", portNumber)
	http.HandleFunc("/", handlerHome)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal(err)
	}
}
