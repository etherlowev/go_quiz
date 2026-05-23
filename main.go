package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	_ "unicode/utf8"
)
import _ "fmt"
import _ "bufio"
import _ "os"
import _ "path/filepath"
import s "strings"
import _ "flag"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var path = *flag.String("path", "problems.csv", "relative or absolute path")
	var timeLimit = *flag.Duration("time", 30, "time limit for question in seconds")

	const questionFormat string = "%v: "
	fileData, err := os.ReadFile(path)
	check(err)

	strFileData := string(fileData)

	splittedData := s.Split(strFileData, "\n")

	var score int = 0
	var questionCount int = 0

	go func(limit time.Duration) {
		time.Sleep(timeLimit * time.Second)
		fmt.Println("time out")
	}(timeLimit)

	for ; questionCount < len(splittedData); questionCount++ {

		curr := splittedData[questionCount]
		splittedCurr := s.Split(curr, ",")
		question := s.Trim(splittedCurr[0], " ")
		answer := s.Trim(splittedCurr[1], " ")
		fmt.Printf(questionFormat, question)
		var userAnswer string

		_, err := fmt.Scanln(&userAnswer)

		check(err)
		if userAnswer == answer {
			score++
		}
	}

	defer fmt.Printf("Result: %v/%v", score, questionCount)
}
