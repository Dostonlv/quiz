package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"strings"

	"log"
	"os"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
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

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
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
