package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/enceeobee/crib/hand"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h := &hand.Hand{}
	err = json.Unmarshal(body, h)
	if err != nil {
		log.Printf("Error unmarshaling hand: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := &hand.CountResponse{
		Points: h.Count(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Cribbage!")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/count", countHandler)

	fmt.Println("\nCribbage server started")

	log.Fatal(http.ListenAndServe(":8456", nil))
}
