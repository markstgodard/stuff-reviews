package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Review struct {
	ID          int    `json:"id"`
	ProductID   int    `json:"product_id"`
	Reviewer    string `json:"reviewer"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
}

var reviews []Review

func init() {
	raw, err := ioutil.ReadFile("./reviews.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(raw, &reviews)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getReviews(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(reviews)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", health)
	mux.HandleFunc("/api/reviews", getReviews)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mux,
	}

	log.Println("Starting server..")
	log.Fatal(server.ListenAndServe(), nil)
}
