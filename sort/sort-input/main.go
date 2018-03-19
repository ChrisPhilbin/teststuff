package main

import (
	"fmt"
//	"sort"
)

type userdef []string

// func (u userdef) Len() int			{ return len(u) }
// func (u userdef) Swap(i, j int)		{ u[i], u[j] = u[j], u[i] }
// func (u userdef) Less(i, j int) bool { return u[i] < u[j] }

func main() {

	list := make([]userdef, 5)

	for i := 0; i < 5; i++ {
		fmt.Println("Enter in a name:")
		var input userdef
		fmt.Scan(&input)
	}
	fmt.Println(list)
}