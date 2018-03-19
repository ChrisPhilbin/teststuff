package main

import (
	"fmt"
//	"strconv"
)

func main() {
	fmt.Println("Enter a decimal to convert to binary:")
	var convert int
	fmt.Scan(&convert)
	fmt.Printf("%b \n", convert)
	fmt.Println("Enter a binary number to convert to decimal:")
	fmt.Scan(&convert)
	fmt.Printf("%d", convert)
}