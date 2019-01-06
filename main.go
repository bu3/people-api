package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/bu3/people-api/people"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", people.GetPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
