package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thiagomarcal/rest-api/routes"
)

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", routes.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", routes.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", routes.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", routes.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
