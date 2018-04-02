package main

import "fmt"

type person struct {
	name string
	email string
	runningtotal int64
}

func main(){

	fmt.Println("Let's add the first user:")
	check := true
	for check == true {
		fmt.Println("Enter name:")
		var name string
		fmt.Scanln(&name)
		fmt.Println("Enter email address:")
		var email string
		fmt.Scanln(&email)
		fmt.Println("Enter outstanding balance:")
		var runningtotal int64
		fmt.Scanln(&runningtotal)
		p := person{name: name, email: email, runningtotal: runningtotal}
		fmt.Println("Do you wish to enter another person? Enter 1 to add another.")
		var answer int
		fmt.Scanln(&answer)
		if answer != 1 {
			check = false
		} 
	fmt.Println(p)
	}
}