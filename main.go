package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name"`
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/users", AllUsers).Methods("GET")
	r.HandleFunc("/", HomePage).Methods("GET")
	http.ListenAndServe(":3000", r)
}

func GetAllUsers() []User {
	return []User{
		{Name: "Ahmad"},
		{Name: "Sharif"},
	}
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	users := GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

// HomePage is a handler function that returns a json encoded string
func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to the HomePage!")
}