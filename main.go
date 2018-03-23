package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":3000", r))
}
