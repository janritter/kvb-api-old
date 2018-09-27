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

	name, err := utils.FindClosestMatchingStation(searchStation)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}

	response := typedef.DeparturesResponse{
		Available: true,
		Station:   name,
		Response:  model.GetDeparturesByStationID(utils.GetStationIDForName(name)),
	}

	payload, err := json.Marshal(response)
	if err != nil {
		utils.LogError(err, map[string]string{"module": "controller/stations"})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)

}
