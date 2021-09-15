package main

import "fmt"

func main() {

	var name = "jim"
	var age int32 = 37
	const isCool = true
	var size float32 = 2.4

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", age)
}
