package main

//import
import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"example.com/packages/mymodule"
)

type TempBook struct {
	name   string
	author string
	pbdate string
}
type BookList struct {
	Items []TempBook
}

func (box *BookList) AddItem(item TempBook) []TempBook {
	box.Items = append(box.Items, item)
	return box.Items
}

var items = []TempBook{}
var bookslist = BookList{items}

// main func
func main() {
	fmt.Println("Starting Server at port 8080")
	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileserver := template.Must(template.ParseFiles("./static/index.html"))
		fileserver.Execute(w, bookslist)
	})

	http.HandleFunc("/add-product", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./static/add-products.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		books := mymodule.BookdataStore{}
		books.SetValues(r.FormValue("name"), r.FormValue("author"), r.FormValue("pbdate"))

		fmt.Println("Book Name: ", books.GetName())
		fmt.Println("Author Name: ", books.Getauthor())
		fmt.Println("Publish Date: ", books.Getpbdate())
		fmt.Print("\n")
		tempBook := TempBook{
			name:   books.GetName(),
			author: books.Getauthor(),
			pbdate: books.Getpbdate(),
		}
		items = append(items, tempBook)
		bookslist.AddItem(tempBook)

		fmt.Printf("%#v", items[len(items)-1])

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
