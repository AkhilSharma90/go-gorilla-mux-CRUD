package main

import (
	"fmt"

	//"log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/posts", getItem).Methods("GET")
	http.ListenAndServe(":8080", router)

}

func getItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")

	fmt.Println("Your details")

	json.NewEncoder(w).Encode("Hitting the GET API")

}
