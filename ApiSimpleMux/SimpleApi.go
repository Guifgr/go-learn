package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var count = 0

func main() {
	println("Api in http://localhost:8080/contato")

	people = append(people, Person{
		ID:        "1",
		Firstname: "Name",
		Lastname:  "Lastname",
		Address: &Address{
			City:  "TheCity",
			State: "TheState",
		},
	})

	people = append(people, Person{
		ID:        "2",
		Firstname: "Name2",
		Lastname:  "Lastname2",
		Address: &Address{
			City:  "TheCity2",
			State: "TheState2",
		},
	})

	router := mux.NewRouter()
	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Println(count, " \n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var person Person
	err := decoder.Decode(&person)

	if err != nil {
		panic(err)
	}
	people = append(people, person)
	log.Println("\nID:", person.ID,
		"\nFirst Name:", person.Firstname,
		"\nLast Name:", person.Lastname,
		"\nAddress:", person.Address.City, person.Address.State)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(people)
	}
}
