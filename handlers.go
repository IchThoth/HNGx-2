package main

import (
	"context"
	"encoding/json"
	"net/http"
)

var usercollection = openCollection(Client, "user")

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user *User

	_, err := usercollection.InsertOne(context.Background(), user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
