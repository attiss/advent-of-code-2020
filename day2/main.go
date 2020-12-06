package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/attiss/advent-of-code-2020/day2/records"
)

const (
	inputFile = "input.txt"
)

func main() {
	var policyType int
	switch os.Getenv("POLICY_TYPE") {
	case "count":
		policyType = records.CountPolicyType
	case "position":
		policyType = records.PositionPolicyType
	}

	input, err := readInputFromFile(inputFile, policyType)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	var validRecords int
	for _, record := range input {
		if record.Policy.ValidatePassword(record.Password) {
			validRecords++
		}
	}

	log.Println(validRecords)
}

func readInputFromFile(filePath string, policyType int) (input []records.Record, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var rec records.Record

		switch policyType {
		case records.CountPolicyType:
			var policy records.CountPolicy
			n, parseErr := fmt.Sscanf(line, "%d-%d %c: %s", &policy.MinOccurances, &policy.MaxOccurantes, &policy.MustContain, &rec.Password)
			if n != 4 || parseErr != nil {
				err = fmt.Errorf("failed to parse input: %s (%d; %v)", line, n, parseErr)
				return
			}
			rec.Policy = policy
		case records.PositionPolicyType:
			var policy records.PositionPolicy
			n, parseErr := fmt.Sscanf(line, "%d-%d %c: %s", &policy.PosA, &policy.PosB, &policy.MustContain, &rec.Password)
			if n != 4 || parseErr != nil {
				err = fmt.Errorf("failed to parse input: %s (%d; %v)", line, n, parseErr)
				return
			}
			rec.Policy = policy
		}

		input = append(input, rec)
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
