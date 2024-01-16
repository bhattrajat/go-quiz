package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func getQueAnsFromCsvFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read csv file", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse content from csv file", err)
	}
	return rows
}

func startQuiz(queAns [][]string, timeLimit int) {
	fmt.Println("Welcome to the math quiz 👨‍🔬")
	fmt.Println("Try to enter correct answers for each questions")
	totalQuestions := len(queAns)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	score := 0
	for ind, qA := range queAns {
		ansChan := make(chan string)
		go func() {
			question := qA[0]
			fmt.Printf("Question %d: %s = ", ind+1, question)
			var userAns string
			fmt.Scan(&userAns)
			ansChan <- userAns
		}()
		select {
		case <-timer.C:
			fmt.Println()
			fmt.Println("Time Up")
			fmt.Println("Your score: ", score, "out of", totalQuestions)
			return
		case userAns := <-ansChan:
			answer := qA[1]
			if userAns == answer {
				score++
			}
		}
	}
	fmt.Println("Your score: ", score, "out of", totalQuestions)
}

func main() {
	fileName := flag.String("file", "problems.csv", "A path to csv file to read and parse questions from")
	timeLimit := flag.Int("timer", 30, "A quiz timer in seconds")
	flag.Parse()
	queAns := getQueAnsFromCsvFile(*fileName)

	startQuiz(queAns, *timeLimit)
}
