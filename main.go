package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"strings"
	"time"

	"log"
	"os"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Printf("failed to open the CSV file : %s\n", *csvFile)
		return
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Printf("failed to parse the provided CSV file")
		return
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
problemloop:

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer

		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerChannel:
			if answer == p.answer {
				correct++
			}

		}

	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	question string
	answer   string
}
