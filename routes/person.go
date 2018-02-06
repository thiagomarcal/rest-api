package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thiagomarcal/rest-api/models"
)

// RouteError teste
type RouteError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// GetPeople Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Persons)
}

// GetPerson Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person, ok := models.Persons[params["id"]]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&RouteError{404, "Person not found"})
	} else {
		json.NewEncoder(w).Encode(person)
	}
}

// CreatePerson create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	models.Persons[person.ID] = person
	json.NewEncoder(w).Encode(models.Persons)
}

// DeletePerson Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	delete(models.Persons, params["id"])
	json.NewEncoder(w).Encode(models.Persons)

}
