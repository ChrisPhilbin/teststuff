package main

import "fmt"

func main() {

	take :=	func(number int) (int, bool) {
		number = number%2
		var even bool
		if number == 0 {
			even = true
		} else {
			even = false
		}
		return number, even
	}

	number, even := take(3)
	fmt.Println(number)
	fmt.Println(even)
}