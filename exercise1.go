package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	questions_file_name, timer_value := commandLineParams()
	fmt.Println(timer_value) // remove this line when we implement the timer.
	file, _ := os.Open(string(*questions_file_name))
	reader := csv.NewReader(file)
	correct_points, incorrect_points := 0, 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[0])
		var user_input string
		fmt.Scanln(&user_input)
		if user_input == record[1] {
			correct_points += 1
		} else {
			incorrect_points += 1
		}
	}
	fmt.Println("Your score was as follows - ", "total correct points:", correct_points, "total incorrect points:", incorrect_points)
}

func commandLineParams() (filename *string, timer_value int) {
	questions_file_name := flag.String("file", "problems.csv", "")
	flag.Parse()
	quiz_timer := flag.String("timer", "30", "")
	flag.Parse()
	timer_value, _ = strconv.Atoi(*quiz_timer)
	return questions_file_name, timer_value
}
