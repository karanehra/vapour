package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	vapour "vapour/cache"
	"vapour/handlers"
	"vapour/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	vapour.InitCache()
	router := mux.NewRouter()
	router.Use(middlewares.JSONMiddleware)
	router.HandleFunc("/get/{key}", handlers.GetKey).Methods("GET")
	router.HandleFunc("/set", handlers.SetKey).Methods("POST")
	const PORT = 3009
	fmt.Printf("Server started on PORT:%d at %d\n", PORT, time.Now().UnixNano()/int64(time.Millisecond))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}
