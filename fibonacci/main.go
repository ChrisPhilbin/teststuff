package main

import "fmt"

func main() {
	var a int = 1
	var b int = 1
	var c int = 0
	var sum int = 0
	for ; sum <= 4000000; {
		c = a + b
		if c % 2 == 0 {
			sum += c
			a = b
			b = c
		} else {
			a = b
			b = c
		}
	}
fmt.Println(sum)
}