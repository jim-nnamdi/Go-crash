package main

import "fmt"

func main() {

	// Define map

	emails := make(map[string]string)

	emails["bob"] = "bob@gmail.com"
	emails["jim"] = "jim@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"])

	delete(emails, "bob")
	fmt.Println(emails)
}
