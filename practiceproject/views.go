package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HomePage2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: SecondPage")
	w.Write([]byte(`{"message":"hello second api page!"}`))
}
func CreateNewBook(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body

	fmt.Println("Endpoint Hit: create a new book")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book 
	json.Unmarshal(reqBody, &book)
	DB.Create(&book)
}

// func returnAllBooks(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnAllBooks")
// 	json.NewEncoder(w).Encode(Books)
// }
