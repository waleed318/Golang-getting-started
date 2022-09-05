package main

import (
	"fmt"
)

var totalbookscount int = 1

type BookdataStore struct {
	name   string
	author string
	pbdate string
}

func main() {
	var choice int = 1
	books := BookdataStore{}
	for choice == 1 {
		var name string
		var author string
		var pbdate string
		fmt.Printf("Enter Book Name: ")
		fmt.Scanln(&name)
		fmt.Printf("Enter Book Author name: ")
		fmt.Scanln(&author)
		fmt.Printf("Enter Publish Date: ")
		fmt.Scanln(&pbdate)
		books.name = name
		books.author = author
		books.pbdate = pbdate
		fmt.Println("Book Name: ", books.name)
		fmt.Println("Publish date: ", books.pbdate)
		fmt.Println("Author: ", books.author)
		totalbookscount++
		fmt.Printf("Do you want to enter another book (yes:1,no:0) : ")
		fmt.Scanln(&choice)
	}
}
