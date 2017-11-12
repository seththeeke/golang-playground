package main

import (
    //"encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

type Person struct {
    firstName string
    lastName string
    age  int
}

func main(){
	fmt.Printf("Starting Test Web Service\n")
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
    larry := new(Person)
    larry.firstName = "Larry"
    larry.lastName = "Theeke"
    larry.age = 240
	fmt.Printf(larry.firstName + "\n")
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting person\n")
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("creating people\n")
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("deleting people\n")
}