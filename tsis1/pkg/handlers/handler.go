package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tsis1/api"	
	"github.com/gorilla/mux"
)
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to F1 Api made by Darhan")
}

func GetTeams(w http.ResponseWriter, r *http.Request) {
	teams := api.Teams

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(teams)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
    	return
	}
	w.Write(jsonResponse)
}
func GetTeamByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["team"]

	team := api.GetTeamByName(name)
	w.Header().Set("Content-Type", "application/json")

	if team == nil {
		http.Error(w, "Team not found", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(team)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}