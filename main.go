package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/janritter/kvb-api/controller"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/departures/stations/{key}", controller.GetDeparturesForStationHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
