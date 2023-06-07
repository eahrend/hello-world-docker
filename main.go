package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Hello world, the web server

	customMessage := os.Getenv("MESSAGE")
	if customMessage == "" {
		customMessage = "Hello, world!"
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, customMessage)
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
