package main

import (
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Transaction struct {
	TXID string
	Value string
	Confirmations int
}

type Result struct {
	Confirmed int
	Unconfirmed int
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var transactions []Transaction
	var result Result

	response, err := http.Get(os.Getenv("API_URL") + mux.Vars(r)["address"])

	if err == nil {
		data,_ := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(data), &transactions)

		var length = len(transactions)

		if (length == 0) {
			handleError(w)
			return
		}
		
		for i := 0; i < length; i++ {
			res,_ := strconv.Atoi(transactions[i].Value)
			if transactions[i].Confirmations > 2 {
				result.Confirmed += res
			} else {
				result.Unconfirmed += res
			}
		}

		json.NewEncoder(w).Encode(result)
	} else {
		handleError(w)
	}
}

func handleError(w http.ResponseWriter) {
	jsonData := map[string]string{"error": "Wallet not found."}
	jsonValue,_ := json.Marshal(jsonData)
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonValue)
}

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

	router := mux.NewRouter()

	router.HandleFunc("/balance/{address}", GetBalance).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}