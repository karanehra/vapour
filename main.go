package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	"vapour/cache"
	"vapour/handlers"
	"vapour/middlewares"
	"vapour/util"

	"github.com/gorilla/mux"
)

func main() {
	EXPIRY, err := strconv.Atoi(os.Getenv("CACHE_DEFAULT_EXPIRY_MINUTES"))
	if err != nil {
		log.Fatal("Env variable 'CACHE_DEFAULT_EXPIRY_MINUTES' not specified")
	}

	stoppedServer := make(chan bool, 1)
	quitServer := make(chan os.Signal, 1)

	signal.Notify(quitServer, os.Interrupt)

	cache.InitCache(time.Duration(EXPIRY) * time.Minute)

	router := mux.NewRouter()
	router.Use(middlewares.JSONMiddleware)
	router.Use(middlewares.CORSMiddleware)
	router.HandleFunc("/get/{key}", handlers.GetKey).Methods("GET")
	router.HandleFunc("/set", handlers.SetKey).Methods("POST")
	router.HandleFunc("/counter/get/{name}", handlers.GetCounter).Methods("GET")
	router.HandleFunc("/counter/increment/{name}", handlers.IncrementCounter).Methods("GET")
	router.HandleFunc("/queue/create", handlers.CreateQueue).Methods("POST")
	router.HandleFunc("/queue/{name}/enqueue", handlers.AddToQueue).Methods("POST")
	router.HandleFunc("/queue/{name}/dequeue", handlers.AddToQueue).Methods("POST")
	router.HandleFunc("/status", handlers.GetStatus).Methods("GET")
	router.HandleFunc("/analytics/main", handlers.GetAllShards).Methods("GET")

	router.HandleFunc("/bucket", handlers.CreateBucket).Methods("POST")

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("Env variable 'PORT' not specified")
	}
	fmt.Printf("Server started on PORT:%s at %d\n", PORT, util.GetMsSinceEpoch())

	server := new(http.Server)
	server.ReadTimeout = 5 * time.Second
	server.WriteTimeout = 5 * time.Second
	server.Addr = fmt.Sprintf(":%s", PORT)
	server.Handler = router

	go gracefulShutdown(server, quitServer, stoppedServer)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}

	<-stoppedServer
	fmt.Println("Bye")
}

func gracefulShutdown(server *http.Server, quitChan <-chan os.Signal, stopChan chan<- bool) {
	<-quitChan
	fmt.Println("\nShutting down Vapour..")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println(cache.MasterCache.KeyCount)

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Error while closing server")
	}
	close(stopChan)
}
