package main

import (
	"io"
	"io/ioutil"
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

	externalURL := os.Getenv("EXTERNAL_URL")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, customMessage)
	}

	canaryHandler := func(w http.ResponseWriter, req *http.Request) {
		res, err := http.Get(externalURL)
		if err != nil {
			panic(err)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(b))
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/canary", canaryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
