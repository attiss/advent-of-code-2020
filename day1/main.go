package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	inputFile = "input.txt"
	target    = 2020
)

func main() {
	input, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	sort.Ints(input)

	var foundTarget bool
	var result int
	var steps int
	for i := 0; i < len(input); i++ {
		if foundTarget {
			break
		}

		for j := len(input) - 1; j >= 0; j-- {
			steps++

			if (input[i] + input[j]) == target {
				foundTarget = true
				result = input[i] * input[j]
				break
			}
		}
	}

	if foundTarget {
		log.Printf("successfully found target (in %d steps), calculated result: %d", steps, result)
	} else {
		log.Println("failed to find target")
	}
}

func readInputFromFile(filePath string) (input []int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var parsedInt int
		parsedInt, err = strconv.Atoi(scanner.Text())
		input = append(input, parsedInt)
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
