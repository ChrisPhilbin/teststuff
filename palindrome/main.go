package main

import (
	"fmt"
	"strconv"
)
func convert(n int) string {
	s := strconv.Itoa(n)
	return s
}
func reverse(s string) string {
	var reverse string
	for i := len(s)-1; i >= 0 ; i-- {
		reverse += string(s[i])
	}
	return reverse
}
func main() {
a := 103
//b := convert(a)
c := convert(a)
d := reverse(c)

for i := 1; 1 < 900; i++ {
	num := 100+i * 100+i
	convert(num)
	var new_num string
	new_num = reverse(new_num)
	if new_num == num {
		fmt.Println(num)
	}
}

if c == d {
	fmt.Println(d, "is a palidrome")
} else {
	fmt.Println(d, "is not a palidrome")
}


}