package main

import (
	"BookProject/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

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
	book := &models.Books{}
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), book)
	fmt.Println(body)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks:= models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/book/create", CreateNewBook)
	router.HandleFunc("/books/", GetBooks)
	log.Fatal(http.ListenAndServe(":8080", router))

}
