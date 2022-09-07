package main

//import
import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"example.com/packages/mymodule"
)

var items = []mymodule.BookdataStore{}
var bookslist = mymodule.BookList{items}

type Context struct {
	Success   bool
	bookslist mymodule.BookList
}

// main func
func main() {
	fmt.Println("Starting Server at port 8080")
	//Initializing index Html file
	fileserver := template.Must(template.ParseFiles("./static/index.html"))
	//Request handle
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fileserver.Execute(w, nil)
			return
		}

		fileserver.Execute(w, bookslist)
	})

	//Initializing Form Html File
	tmpl := template.Must(template.ParseFiles("./static/add-products.html"))
	//Request Parsing
	http.HandleFunc("/add-product", func(w http.ResponseWriter, r *http.Request) {
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
		bookslist.AddItem(books)
		fmt.Println(bookslist)
		//  counter++
		// context := Context{
		// 	Success:   true,
		// 	bookslist: bookslist,
		// }
		tmpl.Execute(w, struct{ Success bool }{true})
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
