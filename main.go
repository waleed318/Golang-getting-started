package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/subosito/gotenv"
	"github.com/waleed318/Golang-getting-started/driver"
	"github.com/waleed318/Golang-getting-started/helper"
	"github.com/waleed318/Golang-getting-started/models"
	"github.com/waleed318/Golang-getting-started/repository/bookRepository"
)

var db *sql.DB

func init() {
	helper.LogFatal(gotenv.Load())
}

func main() {
	db = driver.ConnectDB()
	// var books []models.Book

	// controller := controllers.Controller{}

	// router := mux.NewRouter()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileserver := template.Must(template.ParseFiles("./static/index.html"))
		var book models.Book
		var books = []models.Book{}

		bookRepo := bookRepository.BookRepository{}
		books = bookRepo.GetBooks(db, book, books)
		fmt.Println(books)
		fileserver.Execute(w, books)
	})

	http.HandleFunc("/add-product", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./static/add-products.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		var book models.Book

		fmt.Println(r.Body)
		// helper.LogFatal(json.NewDecoder(r.Body).Decode(&book))

		fmt.Println("Got here")
		iD, _ := strconv.Atoi(r.FormValue("ID"))

		book.ID = iD
		book.Name = r.FormValue("name")
		book.Author = r.FormValue("author")
		book.Pbdate = r.FormValue("pbdate")

		fmt.Println(book)

		bookRepo := bookRepository.BookRepository{}
		bookRepo.AddBook(db, book)

		tmpl.Execute(w, struct{ Success bool }{true})
	})
	println("Server started on port 8080")
	// helper.LogFatal(http.ListenAndServe(":80", router))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
