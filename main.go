package main

import (
	"fmt"
	"log"
	"net/http"
	"vapour/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get/{key}", handlers.GetKey).Methods("GET")
	router.HandleFunc("/set", handlers.GetKey).Methods("POST")
	const PORT = 3000
	fmt.Printf("Server started on PORT:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}
