package main

import "fmt"

var array = []string{}

func user(name string, email string, ph string) {
	temparray := [3]string{name, email, ph}
	array = temparray[:]
}

func main() {
	var name string
	var email string
	var ph string
	fmt.Printf("Enter your Name: ")
	fmt.Scanln(&name)
	fmt.Printf("Enter your Email: ")
	fmt.Scanln(&email)
	fmt.Printf("Enter your Phone No: ")
	fmt.Scanln(&ph)

	user(name, email, ph)
	for index, element := range array {
		fmt.Println("At index", index, "value is", element)
	}
}
