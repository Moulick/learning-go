package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handlerearth)

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf(string(r.Header))
	fmt.Fprint(w, "Hello World")

}

func handlerearth(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello Earth")

}
