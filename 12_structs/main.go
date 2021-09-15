package main

import "fmt"

// Define person struct
type Teacher struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	teacherOne := Teacher{firstname: "jim", lastname: "polo", age: 23}
	fmt.Println(teacherOne)
	fmt.Println(teacherOne.firstname)
	fmt.Println(teacherOne.lastname)
}
