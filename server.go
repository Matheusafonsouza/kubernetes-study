package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello</h1>"))
}
