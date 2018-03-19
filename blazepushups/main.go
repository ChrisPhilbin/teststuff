// At UAB football games, Blaze does push ups after each Blazer score. After the first Blazer touchdown (and point after), Blaze does 7 push ups. After the second touchdown and point after, the score is now 14 and Blaze does 14 push ups.

// Write a program that calculates how many total push ups Blaze does during the whole game. Assume that only 7 point touchdowns (including the point after) occur. Prompt for the final score and print out how many push ups Blaze has done.

package main

import "fmt"

func main() {
	fmt.Println("What was the final score of the game?")
	var score int
	var pushups int
	fmt.Scan(&score)
	for i := 7; i <= score; i = i + 7 {
		pushups = pushups + i
	}
	fmt.Println("Blaze did", pushups, "pushups during the game")
}