package main

import (
    //"encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

func main(){
	fmt.Printf("Starting Test Web Service")
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting people\n")
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting people\n")
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting people\n")
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting people\n")
}