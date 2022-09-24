package main

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FullNme  string
	UserName string
	Email    string
}

type Post struct {
	Title  string
	Body   string
	Author User
}

var data []Post = []Post{}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/posts", getItem).Methods("GET")
	router.HandleFunc("/posts", addItems).Methods("POST")
	http.ListenAndServe(":8080", router)

}

func addItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var newPost Post

	json.NewDecoder(r.Body).Decode(&newPost)

	data = append(data, newPost)

	json.NewEncoder(w).Encode(data)

}

func getItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(data)

}
