package main

import "fmt"

func  main () {
	var large, small int
	fmt.Println("Enter a number:")
	fmt.Scan(&large)
	fmt.Println("Enter a number smaller than the one you just entered:")
	fmt.Scan(&small)
	ans := large/small
	fmt.Printf("The answer is: %d", ans)

}