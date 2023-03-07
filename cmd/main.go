package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt_prac2 "github.com/ekholme/go_jwt_prac2"
	"github.com/gorilla/mux"
)

const listenAddr = ":8080"

//basic hello world
func main() {
	fmt.Printf("Running go jwt practice server on %s\n", listenAddr)

	//create a server
	us := jwt_prac2.NewUserStore()
	r := mux.NewRouter()

	s := jwt_prac2.NewServer(r, us)

	s.Router.HandleFunc("/", handleIndex).Methods("GET")

	s.Srvr.ListenAndServe()

}

//handle the index
func handleIndex(w http.ResponseWriter, r *http.Request) {
	msg := make(map[string]string)

	msg["hello"] = "world"

	writeJSON(w, http.StatusOK, msg)
}

//helper to write JSON
func writeJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(v)
}
