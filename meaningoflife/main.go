package main

import "fmt"

func main() {
	var submission int
	for submission != 42 {
		fmt.Println("Enter a number:")
		fmt.Scan(&submission)
	}
	fmt.Println("You found the meaning of life!")
}