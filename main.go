package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/health", healthCheck).Methods("GET")
	router.HandleFunc("/java", java).Methods("GET")
	router.HandleFunc("/php", php).Methods("GET")
	router.HandleFunc("/demo", demo).Methods("GET")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}

func java(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://c2c-java.apps.internal:8080/demo")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		json.NewEncoder(w).Encode(string(data))
	}
}

func php(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://c2c-java.apps.internal:8080/php")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		json.NewEncoder(w).Encode(string(data))
	}
}

func demo(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://c2c.apps.internal:8080/demo")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		json.NewEncoder(w).Encode(string(data))
	}
}
