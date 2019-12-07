package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"vapour/cache"
	"vapour/handlers"
	"vapour/middlewares"
	"vapour/util"

	"github.com/gorilla/mux"
)

func main() {
	cache.InitCache(5 * time.Minute)
	router := mux.NewRouter()
	router.Use(middlewares.JSONMiddleware)
	router.HandleFunc("/get/{key}", handlers.GetKey).Methods("GET")
	router.HandleFunc("/set", handlers.SetKey).Methods("POST")
	router.HandleFunc("/counter/get/{name}", handlers.GetCounter).Methods("GET")
	router.HandleFunc("/counter/increment/{name}", handlers.IncrementCounter).Methods("GET")
	router.HandleFunc("/queue/create", handlers.CreateQueue).Methods("POST")
	router.HandleFunc("/queue/{name}/enqueue", handlers.AddToQueue).Methods("POST")
	router.HandleFunc("/queue/{name}/dequeue", handlers.AddToQueue).Methods("POST")
	router.HandleFunc("/status", handlers.GetStatus).Methods("GET")
	router.HandleFunc("/analytics/main", handlers.GetAllShards).Methods("GET")
	const PORT = 3009
	fmt.Printf("Server started on PORT:%d at %d\n", PORT, util.GetMsSinceEpoch())
	server := new(http.Server)
	server.ReadTimeout = 5 * time.Second
	server.WriteTimeout = 5 * time.Second
	server.Addr = fmt.Sprintf(":%d", PORT)
	server.Handler = router
	log.Fatal(server.ListenAndServe())
}
