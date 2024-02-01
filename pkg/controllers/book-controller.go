package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohammedfuta2000/go-bookstore/pkg/models"
	"github.com/mohammedfuta2000/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("start Getting book by id")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("error while parsing: %s", err.Error())
	}

	log.Default().Println("Retrieved id")
	bookDetails, _ := models.GetBookById(ID)
	log.Default().Println("Retrieved book details")
	res, _ := json.Marshal(bookDetails)
	log.Default().Println("Retrieved response")
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	log.Default().Println("sent response")
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := models.Book{}
	utils.ParseBody(r, &createBook)
	log.Default().Println(createBook.Name)

	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("start Getting book by id")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("error while parsing: %s", err.Error())
	}
	b := models.DeleteBook(ID)
	res, _ := json.Marshal(b)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Updatebook 1")
	updatebook := models.Book{}
	utils.ParseBody(r, &updatebook)
	log.Default().Printf("Updatebook 2: %s", updatebook.Publication)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	log.Default().Printf("Updatebook 3: %s", bookId)
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("problem parsing ID: %s", err.Error())
	}
	log.Default().Printf("Updatebook 3: %v", ID)

	bookDetails, db := models.GetBookById(ID)
	if updatebook.Name != "" {
		bookDetails.Name = updatebook.Name
	}
	if updatebook.Author != "" {
		bookDetails.Author = updatebook.Author
	}
	if updatebook.Publication != "" {
		bookDetails.Publication = updatebook.Publication
	}
	log.Default().Printf("Updatebook 3: %s", bookDetails.Name)
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
