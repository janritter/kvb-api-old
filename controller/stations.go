package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/janritter/kvb-api/model"
	"github.com/janritter/kvb-api/typedef"
	"github.com/janritter/kvb-api/utils"

	"github.com/gorilla/mux"
)

func GetDeparturesForStationHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	searchStation := vars["key"]

	name := utils.FindClosestMatchingStation(searchStation)

	response := typedef.DeparturesResponse{
		Available: true,
		Station:   name,
		Response:  model.GetDeparturesByStationID(utils.GetStationIDForName(name)),
	}

	payload, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
