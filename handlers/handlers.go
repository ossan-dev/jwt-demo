package handlers

import (
	"encoding/json"
	"net/http"
)

type message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		panic(err)
	}
	r.Body.Close()
	err = json.NewEncoder(w).Encode(msg)
	if err != nil {
		panic(err)
	}
}

func ReservedPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := &message{Status: "private endpoint", Info: "this is a private endpoint, you must be authenticated for seeing it"}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
}
