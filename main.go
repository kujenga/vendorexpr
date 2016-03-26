package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handling request")
		w.Write([]byte("you got it!"))
	})

	log.Println("starting server...")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
