package main

import "fmt"

func main() {
	var a int64 = 1
	var b int64 = 1
	var c int64 = 0
	var sum int64 = 0
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