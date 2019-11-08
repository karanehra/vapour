package main

import (
	"fmt"
	"log"
	"net/http"
	"vapour/handlers"
	"vapour/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.JSONMiddleware)
	router.HandleFunc("/get/{key}", handlers.GetKey).Methods("GET")
	router.HandleFunc("/set", handlers.SetKey).Methods("POST")
	const PORT = 3009
	fmt.Printf("Server started on PORT:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}
