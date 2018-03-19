package main

import (
	"fmt"
)

func main() {

	something := make([]string, 5)

	for i := 0; i < 5; i++ {
		fmt.Println("Enter in a name:")
		var name string
		fmt.Scan(&name)
		something[i] = name
	}

	fmt.Println(something)
}