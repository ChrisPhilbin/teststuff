package main

import (
	"fmt"
)

func test() {
	for i := 0; i < 40; i++ {
		fmt.Println("Hello there")
	}
}

func anothertest() {
	for a := 0; a < 40; a++ {
		fmt.Println("Another test")
	}
}

func main() {

	go test()
	go anothertest()
	// fmt.Println("another function")

}