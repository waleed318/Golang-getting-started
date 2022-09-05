package main

import (
	"fmt"

	"example.com/packages/mymodule"
)

var totalbookscount int = 1

func main() {
	var choice int = 1

	for choice == 1 {
		var name_ string
		var auth string
		var date string
		fmt.Printf("Enter Book Name: ")
		fmt.Scanln(&name_)
		fmt.Printf("Enter Book Author name: ")
		fmt.Scanln(&auth)
		fmt.Printf("Enter Publish Date: ")
		fmt.Scanln(&date)
		books := mymodule.BookdataStore{}
		books.SetValues(name_, auth, date)
		fmt.Println("Book Name: ", books.GetName())
		fmt.Println("Publish date: ", books.Getauthor())
		fmt.Println("Author: ", books.Getpbdate())
		totalbookscount++
		fmt.Printf("Do you want to enter another book (yes:1,no:0) : ")
		fmt.Scanln(&choice)
	}
}
