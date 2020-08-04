package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoProjects/practiceproject/models"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func main() {
	models.DBConnector()
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	// http.HandleFunc("/sec/", HomePage2)
	// // http.HandleFunc("/books/", ReturnAllBooks)
	router.HandleFunc("/book/", CreateNewBook)
	log.Fatal(http.ListenAndServe(":8080", router))

}
