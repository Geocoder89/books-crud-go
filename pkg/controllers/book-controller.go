package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Geocoder89/go-books-crud/pkg/models"
	"github.com/Geocoder89/go-books-crud/pkg/utils"
	
	"github.com/gorilla/mux"
)



var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks:= models.GetAllBooks()
	res,_ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err:= strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ :=models.GetBookById(ID)
	
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook (w http.ResponseWriter, r *http.Request) {
	// we set to json initially
	bookCreated := &models.Book{}

	// convert what was sent to data the db can understand
	utils.ParseBody(r,bookCreated)

	//  we create the book
	b:=  bookCreated.CreateBook()

	// we then return it back to a json response
	res, _ := json.Marshal(b)

	// we then write back a suitable response back to the user
	w.Write(res)

}


func DeleteBookById (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBookById(ID)

	res,_ := json.Marshal(book)

	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}

	utils.ParseBody(r,updatedBook)
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails,db := models.GetBookById(ID)

	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}


	if updatedBook.Author !="" {
		bookDetails.Author = updatedBook.Author
	}


	if updatedBook.Publication != "" {
		bookDetails.Publication= updatedBook.Publication
	}

	db.Save(&bookDetails)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

