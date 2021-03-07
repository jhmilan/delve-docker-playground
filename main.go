package main

import (
	"log"
	"net/http"
	"strings"

	quoteV3 "rsc.io/quote"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Path
		message = strings.TrimPrefix(message, "/")
		message = "Hello, " + message + "!"

		w.Write([]byte(message))
	})

	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		message := "Hello, " + quoteV3.Go()

		w.Write([]byte(message))
	})

	http.HandleFunc("/glass", func(w http.ResponseWriter, r *http.Request) {
		message := "Hello, " + quoteV3.Glass()

		w.Write([]byte(message))
	})

	log.Print("starting web server")
	if err := http.ListenAndServe(":38081", nil); err != nil {
		log.Fatal(err)
	}
}
