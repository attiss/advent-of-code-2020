package main

import (
	"bufio"
	"log"
	"os"

	"github.com/attiss/advent-of-code-2020/day6/group_answers"
)

const (
	inputFile = "input.txt"
)

func main() {
	aggregatedGroupAnswers, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	var sum int
	for _, groupAnswers := range aggregatedGroupAnswers {
		sum += groupAnswers.CountUniqueItems()
	}

	log.Printf("sum of unique group answers: %d", sum)
}

func readInputFromFile(filePath string) (aggregatedGroupAnswers []group_answers.GroupAnswers, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	var groupAnswers group_answers.GroupAnswers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			aggregatedGroupAnswers = append(aggregatedGroupAnswers, groupAnswers)
			groupAnswers = group_answers.GroupAnswers{}
		} else {
			groupAnswers = append(groupAnswers, line)
		}
	}
	aggregatedGroupAnswers = append(aggregatedGroupAnswers, groupAnswers)

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
