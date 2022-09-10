package bookRepository

import (
	"database/sql"
	"fmt"

	"github.com/waleed318/Golang-getting-started/helper"
	"github.com/waleed318/Golang-getting-started/models"
)

type BookRepository struct{}

const (
	bookStore = "book_store"
)

func (BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("SELECT * FROM " + bookStore + ";")
	helper.LogFatal(err)

	defer helper.LogFatal(rows.Close())

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Author, &book.Name, &book.Pbdate)
		helper.LogFatal(err)

		books = append(books, book)
		fmt.Println(books)
	}

	return books
}

func (BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
	row := db.QueryRow("select * from "+bookStore+" where ID=$1", id)

	err := row.Scan(&book.ID, &book.Author, &book.Name, &book.Pbdate)
	helper.LogFatal(err)

	return book
}

func (BookRepository) AddBook(db *sql.DB, book models.Book) int {
	err := db.QueryRow("insert into "+bookStore+" (ID,name, author, pbdate) values($1, $2, $3,$4) RETURNING ID;", book.ID, book.Name, book.Author, book.Pbdate).Scan(&book.ID)
	helper.LogFatal(err)

	return book.ID
}

func (BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update "+bookStore+" set name=$1, author=$2, pbdate=$3 where id=$4 RETURNING ID", &book.Name, &book.Author, &book.Pbdate, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	helper.LogFatal(err)

	return rowsUpdated
}

func (BookRepository) DeleteBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from "+bookStore+" where ID=$1", id)
	helper.LogFatal(err)

	rowDeleted, err := result.RowsAffected()
	helper.LogFatal(err)

	return rowDeleted
}
