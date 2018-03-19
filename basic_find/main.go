package main

import (
	"fmt"
)

func main() {

	findme := make([]int, 4)

	for i := 0; i < 4; i++ {
		fmt.Println("Enter a number to add to the slice: ")
		var number int
		fmt.Scan(&number)
		findme[i] = number
	}

	fmt.Println(findme)

	fmt.Println("Enter a number to search for:")
	var searchfor int
	fmt.Scan(&searchfor)
	for i := range findme {
		if findme[i] == searchfor {
			fmt.Println("Found it!")			
		} 
	}
}