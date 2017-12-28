package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Person struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	people = append(people, Person{ID: "1", FirstName: "User 1 First Name", LastName: "User 1 Last Name", Address: &Address{City: "City 1", State: "State 1"}})
	people = append(people, Person{ID: "2", FirstName: "User 2 First Name", LastName: "User 2 Last Name", Address: &Address{City: "City 2", State: "State 2"}})
	people = append(people, Person{ID: "3", FirstName: "User 3 First Name", LastName: "User 3 Last Name"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	for _, person := range people {
		if person.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(person)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, person := range people {
		if person.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(people)
}
