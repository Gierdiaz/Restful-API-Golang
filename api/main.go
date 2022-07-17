package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user", getUsers)
	fmt.Println("api is on :6049")

	log.Fatal(http.ListenAndServe(":6049", nil))
}

type User struct {
	ID int `json: "id"`
	Name string `json: "name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}


	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			ID: 1,
			Name: "Argus",
		},
		{
			ID: 2,
			Name: "Gudi",
		},
	})
}