package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User
type User struct {
	ID        int    `json:"id"`
	username  string `json:"username"`
	followers int    `json:"followers"`
}

// NewUser
func NewUser() []User {
	usr := []User{
		User{
			ID:        1,
			username:  "SammyShark",
			followers: 987,
		},
		User{
			ID:        2,
			username:  "JesseOctopus",
			followers: 432,
		},
		User{
			ID:        3,
			username:  "DrewSquid",
			followers: 321,
		},
		User{
			ID:        4,
			username:  "JamieMantisShrimp",
			followers: 654,
		},
	}
	return usr
}

// GetUser
func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		usr := NewUser()
		datauser, err := json.Marshal(usr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datauser)
		return
	}

	http.Error(w, "response failed !", http.StatusNotFound)
}

// Server
func main() {
	http.HandleFunc("/DMXK", GetUser)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
