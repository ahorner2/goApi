package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/ahorner2/goApi/server/db"
	"github.com/ahorner2/goApi/server/users"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// define data and struct
type Tool struct {
	ID 					string `json:"id"`
	Name				string `json:"name"`
	Description string `json:"description"`
}

var tools []Tool = []Tool{
	{
		ID:						"1", 
		Name: 			 	"Carl",
		Description:	"Wanted for conspiracy to commit acts of Muay Thai",
	}, 
	{
		ID:						"2", 
		Name: 			 	"Steve",
		Description:	"Yeah, that's just Steve",
	}, 
	{
		ID:						"3", 
		Name: 			 	"John",
		Description:	"Can throw a football at least 15 yards",
	}, 
	{
		ID:						"4", 
		Name: 			 	"Daniel",
		Description:	"Can vibrate through walls, burrow underground",
	}, 
}

// get all tools 
func getTools(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(tools)
}

func main() {
	// load env vars 
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if os.Getenv("DATABASE_URL") == "" {
			log.Fatal("DATABASE_URL not set in env")
	}
	// connect to db 
	db.GetDBConnection()

	// init router (r) 
	r := mux.NewRouter() 

	// pprof handlers for profiling
	r.HandleFunc("/debug/pprof/", http.HandlerFunc(pprof.Index))
	r.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	r.HandleFunc("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	r.HandleFunc("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	r.HandleFunc("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	// "fetch" mock data & endpoint to get all tools 
	r.HandleFunc("/tools", getTools).Methods("GET")
	// create new user endpoint 
	r.HandleFunc("/user", users.CreateNewUser).Methods("POST")
	// update current user creds
	r.HandleFunc("/user/{id}", users.UpdateUser).Methods("PUT")
	// handle delete user 
	r.HandleFunc("/user/{id}", users.DeleteUser).Methods("DELETE")
	// admin view of all mock users 
	r.HandleFunc("/mockusers", users.FetchAllUsers).Methods("GET")

	// start server 
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}