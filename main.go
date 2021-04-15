package main

import (
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type Transaction struct {
	TXID string
	Value string
	Confirmations int
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// transactionsJson := `[{"txid":"74772f6296638990cbb6107792ad370a5bc0a878c649cc763d42506e5178acdf","vout":0,"value":"13394","confirmations":0,"lockTime":672872}]`
	var transactions []Transaction

	response, err := http.Get("https://blockbosok-bitcoin.tronwallet.me/api/v2/utxo/1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2")

	if err == nil {
		data,_ := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(data), &transactions)
		json.NewEncoder(w).Encode(transactions)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/balance/{address}", GetBalance).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}