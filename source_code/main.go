package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Assignment/source_code/Handling"
	"github.com/gorilla/mux"
)

var port string = ":8080"

func main() {
	fmt.Println("Executing main()")
	router := mux.NewRouter().StrictSlash(true) // Create a "mux" package router instance to handle RESTapi's

	// handle Route paths
	router.HandleFunc("/create", Handling.CreateData).Methods("POST")
	router.HandleFunc("/read/{ID}", Handling.ReadData).Methods("GET")
	router.HandleFunc("/update/{ID}", Handling.UpdateData).Methods("PUT")
	router.HandleFunc("/delete/{ID}", Handling.DeleteData).Methods("DELETE")

	// Listen and Serve using the port defined above
	log.Fatal(http.ListenAndServe(port, router))
}
