package main

import (
	"fmt"
	"flag"
	"encoding/csv"
	"os"
	"strings"
)

csvFile := flag.String("csv", "problems.csv", "csv file for the game in question,answer format")

func main() {

	flag.Parse()

	file, error := os.Open(*csvFile)
		if error != nil {
			fmt.Println("There was an error opening file %s\n", csvFile)
			return
		}
	defer file.Close()

	csvFile := csv.NewReader(file)
	csvData, error := csvReader.ReadAll()
		if error != nil {
			fmt.Println("There was an issue reading from the CSV file")
			return
		}

	orderedpair := make(map[string]string, len(csvData))
	for _, data := range csvData {
		orderedpair[data[0]] = data[1]
	}

	var userAnswer string
	var questionNumber, correct int

	for question, answer := range orderedpair {
		questionNumber++

		fmt.Println("Question number %d: %s = ", questionNumber, question)

		scanner.Scan()
		userAnswer = scanner.Text()
		userAnswer = strings.ToLower(userAnswer)
		userAnswer = strings.TrimSpace(userAnswer)
		answer = strings.ToLower(answer)

		if answer == userAnswer {
			correct++
		}

	}

}