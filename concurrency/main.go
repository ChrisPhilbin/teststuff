package main

import (
	"fmt"
	"sync"

)
var wg sync.WaitGroup

func test() {
	for i := 0; i < 40; i++ {
		fmt.Printf("test func -- %d\n", i)
	}
	wg.Done()
}

func anothertest() {
	for a := 0; a < 40; a++ {
		fmt.Printf("another test func -- %d\n", a)
	}
	wg.Done()
}

func main() {

	wg.Add(2)
	go test()
	go anothertest()
	wg.Wait()
	// fmt.Println("another function")

}