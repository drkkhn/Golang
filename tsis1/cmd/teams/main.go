package main 
import (
	"log"
	"net/http"
	"tsis1/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("tsis1 starts...")
	router := mux.NewRouter()
	log.Println("Creating routes...")

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/teams", handlers.GetTeams).Methods("GET")
	router.HandleFunc("/teams/{team}", handlers.GetTeamByNameHandler).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)

}