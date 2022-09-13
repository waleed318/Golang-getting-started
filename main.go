package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/subosito/gotenv"
	"github.com/waleed318/Golang-getting-started/driver"
	"github.com/waleed318/Golang-getting-started/helper"
	"github.com/waleed318/Golang-getting-started/models"
	_ "github.com/waleed318/Golang-getting-started/repository/bookRepository"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	helper.LogFatal(gotenv.Load())
}

func main() {
	client := driver.ConnectMongo()
	booksCollection := client.Database("book_store").Collection("books")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileserver := template.Must(template.ParseFiles("./static/index.html"))

		// BookRepo := bookRepository.BookRepo{}
		// uc := BookRepo.getAllBooks(usersCollection)
		var books []models.BookModel
		Cursor, err := booksCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		if err = Cursor.All(context.TODO(), &books); err != nil {
			panic(err)
		}
		// fmt.Println(books)

		fileserver.Execute(w, books)
	})

	http.HandleFunc("/add-product", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./static/add-products.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		manyContacts := models.BookModel{
			Name:   r.FormValue("name"),
			Author: r.FormValue("author"),
			Pbdate: r.FormValue("pbdate"),
		}

		result, err := booksCollection.InsertOne(context.TODO(), manyContacts)
		if err != nil {
			panic(err)
		}
		fmt.Println("Object Inserted: ", result.InsertedID)

		tmpl.Execute(w, struct{ Success bool }{true})
	})
	println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
