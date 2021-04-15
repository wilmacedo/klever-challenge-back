package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Wallet struct {
	ID string `json:"id"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/balance/{address}", GetBalance).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetBalance(w http.ResponseWriter, r *http.Request) {

}