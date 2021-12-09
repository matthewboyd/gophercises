package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//initalising scoring system
	correct_points, incorrect_points := 0, 0

	//opening problem files
	fileName, timeLimit, shuffle := getFile()

	if bool(*shuffle) == true {

	}
	//getting the problems into an array
	data := getData(fileName)

	problems := parseData(data)
	if bool(*shuffle) == true {
		problems = randomNumberGenerator(problems)
	}
	fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, question := range problems {
		fmt.Println("Problem number", i+1, "Question: ", question.Q) // asking the user the question
		answerCh := make(chan string)
		go func() {
			var input string
			fmt.Scanln(&input)
			answerCh <- input
		}()
		select {
		case <-timer.C:
			fmt.Println("Correct answers:", correct_points, "Incorrect answers:", incorrect_points, "Out of a total:", len(problems))
			return
		case answer := <-answerCh:
			if answer == question.A {
				correct_points += 1
			} else {
				incorrect_points += 1
			}
		}

	}
	fmt.Println("Correct answers:", correct_points, "Incorrect answers:", incorrect_points, "Out of a total:", len(problems))
}

type problem struct {
	Q string
	A string
}

func parseData(data [][]string) []problem {
	ret := make([]problem, len(data))
	for i, line := range data {
		ret[i] = problem{line[0], line[1]}
	}
	return ret
}

func getFile() (*string, *int, *bool) {
	file := flag.String("file", "problems.csv", "a file that will contain questions and answers")
	timeLimit := flag.Int("time", 30, "A time limit for how long the user has to complete the quiz")
	shuffle := flag.Bool("shuffle", false, "A flag to shuffle the problems so that they're not shown in the same order each time")
	flag.Parse()
	return file, timeLimit, shuffle
}

func getData(fileName *string) [][]string {
	file, _ := os.Open(*fileName)
	defer file.Close()

	//importing file into array
	csvReader := csv.NewReader(file)
	data, _ := csvReader.ReadAll()
	return data
}

func randomNumberGenerator(problems []problem) []problem {

	for i := len(problems) / 2; i < len(problems); i++ {
		randNum1 := rand.Intn(len(problems) - 1)
		randNum2 := rand.Intn(len(problems) - 1)
		temp := problems[randNum1]
		problems[randNum1] = problems[randNum2]
		problems[randNum2] = temp
	}
	return problems
}
