package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Person representa a una persona
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirtsName string   `json:"firtsname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

//Address representa la direccion
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

//GetPeopleEndpoint traer todas las personas
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//GetPersonEndpoint traer persona especifica
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//CreatePersonEndpoint crear una nueva persona
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var personaNueva Person
	_ = json.NewDecoder(req.Body).Decode(&personaNueva)
	personaNueva.ID = params["id"]
	people = append(people, personaNueva)
	json.NewEncoder(w).Encode(people)
}

//DeletePersonEndpoint borrar persona
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirtsName: "Juan", LastName: "Topo",
		Address: &Address{City: "Bogota", State: "Cundinamarca"}})

	people = append(people, Person{ID: "2", FirtsName: "Pedro", LastName: "Melindes",
		Address: &Address{City: "Madrid", State: "España"}})

	people = append(people, Person{ID: "3", FirtsName: "Camila", LastName: "Duran",
		Address: &Address{City: "Buenos Aires", State: "Argentina"}})

	//endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	//se ejecuta el lebantaniemto dels servidor utilizando log por siu existe un error
	log.Fatal(http.ListenAndServe(":3000", router))

}
