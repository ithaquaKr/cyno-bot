package http

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func RunServer() {

	//http handlers
	r := mux.NewRouter()
	// r.HandleFunc("/")

	//http loop
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
