package main

import "fmt"

func main() {
	x := 5
	y := 6

	// if else

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	// else if

	color := "purple"

	if color == "purple" {
		fmt.Println("color is purple")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("neither of the colors")
	}

	// switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "purple":
		fmt.Println("color is purple")
	default:
		fmt.Println("color is neither blue nor red")
	}
}
