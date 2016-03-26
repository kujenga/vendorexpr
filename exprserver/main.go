package exprserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handling request")
		w.Write([]byte("you got it!"))
	})

	http.Handle("/", r)
}
