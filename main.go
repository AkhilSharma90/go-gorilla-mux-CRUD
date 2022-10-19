package main

import (
	"encoding/json"
	"strconv"

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
	router.HandleFunc("/posts/{id}", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", updateItem).Methods("PUT")
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

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could  not be  converted  to  integer"))
		return
	}

	if id >= len(data) {
		w.WriteHeader(404)
		w.Write([]byte("No data found with  specified ID"))
		return
	}

	post := data[id]

	json.NewEncoder(w).Encode(post)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	var updatedItem Post

	//updateItem := Post

	json.NewDecoder(r.Body).Decode(&updatedItem)

	data[id] = updatedItem

	json.NewEncoder(w).Encode(updatedItem)

}
