package bookRepository

import (
	"context"
	"fmt"
	"log"

	"github.com/waleed318/Golang-getting-started/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepo struct{}

func (BookRepo) getAllBooks(booksCollection *mongo.Collection) []models.BookModel {
	var uc []models.BookModel
	Cursor, err := booksCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = Cursor.All(context.TODO(), &uc); err != nil {
		panic(err)
	}
	fmt.Println(uc)
	return uc
}
