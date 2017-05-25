package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Server Entry Point
func main() {
	r := mux.NewRouter()

	//Public Endpoints
	r.Handle("/", notImplemented).Methods("GET")

	log.Println("Now listening...")

	//handle server interrupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Stopping Server ...")
		os.Exit(1)
	}()

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	json.NewEncoder(w).Encode("Hurray")

})
