package main

import "fmt"

var totalbookscount int = 1

//Declaring 2 dimensional map(map->map->string)
var booksstoremap = map[string]map[string]string{}

//created a temprary map and append that map into new count books
func addbook(name string, author string, pbdate string) {
	// temparray := [3]string{name, email, ph}
	tempmap := map[string]string{
		"BookName":    name,
		"Author":      author,
		"PublishDate": pbdate,
	}

	bookNo := fmt.Sprint("book_", totalbookscount)
	booksstoremap[bookNo] = tempmap
	totalbookscount++
}

func main() {
	var choice int = 1
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

		addbook(name, author, pbdate)

		fmt.Printf("Do you want to enter another book (yes:1,no:0) : ")
		fmt.Scanln(&choice)
	}
	// Looping through 2 Dimensional Maps
	for key, element := range booksstoremap {
		fmt.Println(key, " : ")
		for index, element := range element {
			fmt.Println("\t", index, " : ", element)
		}
	}
}
