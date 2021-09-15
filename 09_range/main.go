package main

import "fmt"

func main() {
	ids := []int{2, 4, 5, 6, 7}

	// for _, id := range ids {
	// 	fmt.Println(id)
	// }

	// sum

	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println(sum)

	// range with maps

	emails := make(map[string]string)
	emails["bob"] = "bob@gmail.com"
	emails["kate"] = "kate@me.com"

	for key, value := range emails {
		fmt.Printf("%s : %s\n", key, value)
	}
}
