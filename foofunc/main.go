package main

import "fmt"

func foo (a...int) {
	fmt.Println("bar")
}

func main() {
	foo(1,2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}