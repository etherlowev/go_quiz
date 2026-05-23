package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	s "strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseFile(path string) (int, []string, []string, error) {
	fileData, err := os.ReadFile(path)
	check(err)

	strFileData := string(fileData)

	splittedData := s.Split(strFileData, "\n")

	var amountQuestions = len(splittedData)

	var questions = make([]string, amountQuestions)
	var answers = make([]string, amountQuestions)

	for i := 0; i < len(splittedData); i++ {
		curr := splittedData[i]
		splittedCurr := s.Split(curr, ",")
		if len(splittedCurr) > 1 {
			question := s.Trim(splittedCurr[0], " ")
			answer := s.Trim(splittedCurr[1], " ")

			questions[i] = question
			answers[i] = answer
		} else {
			fmt.Printf("csv structure is not correct on line %v\n", i)
			return 0, nil, nil, errors.New("csv structure is not correct")
		}
	}
	return amountQuestions, questions, answers, nil
}

func main() {
	// flags
	var pathPtr = flag.String("path", "problems.csv", "relative or absolute path")
	var timeLimitPtr = flag.Duration("time", 30, "time limit for question in seconds")
	flag.Parse()
	// flags end

	// vars
	const questionFormat string = "%v: "
	endChan := make(chan int)
	var score = 0
	// vars end

	questionCount, questions, answers, err := parseFile(*pathPtr)
	check(err)

	fmt.Print("Press enter to start")
	_, err = fmt.Scanln()

	go func(ch chan int) {
		for i := 0; i < questionCount; i++ {

			currQ := questions[i]
			currA := answers[i]
			fmt.Printf(questionFormat, currQ)

			var userAnswer string
			_, err = fmt.Scanln(&userAnswer)

			if err != nil {
				continue
			}

			if userAnswer == currA {
				score++
			}
		}
		ch <- 0
	}(endChan)

	select {
	case <-endChan:
		break
	case <-time.After(time.Second * *timeLimitPtr):
		fmt.Println("\ntimeout")
	}

	fmt.Printf("Result: %v/%v", score, questionCount)
}
