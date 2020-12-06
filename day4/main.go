package main

import (
	"bufio"
	ppkg "github.com/attiss/advent-of-code-2020/day4/passports"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

func main() {
	passports, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	validCount := 0
	for _, passport := range passports {
		valid, err := passport.IsValid()
		if err != nil {
			log.Printf("invalid passport (%+v): %v", passport, err)
		}
		if valid {
			validCount++
		}
	}

	log.Printf("found %d valid passports", validCount)
}

func readInputFromFile(filePath string) (passports []ppkg.Passport, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	var passport ppkg.Passport
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			passports = append(passports, passport)
			passport = ppkg.Passport{}
			continue
		}

		lineParts := strings.Split(line, " ")
		for _, linePart := range lineParts {
			keyValuePair := strings.Split(linePart, ":")
			if len(keyValuePair) != 2 {
				log.Fatalf("invalid key-value pair found: %+v", keyValuePair)
				break
			}

			switch keyValuePair[0] {
			case "byr":
				byr, err := strconv.Atoi(keyValuePair[1])
				if err != nil {
					log.Fatalf("invalid birtday year value: %v (%v)", keyValuePair[1], err)
					break
				}
				passport.BirthYear = byr
			case "iyr":
				iyr, err := strconv.Atoi(keyValuePair[1])
				if err != nil {
					log.Fatalf("invalid issue year value: %v (%v)", keyValuePair[1], err)
					break
				}
				passport.IssueYear = iyr
			case "eyr":
				eyr, err := strconv.Atoi(keyValuePair[1])
				if err != nil {
					log.Fatalf("invalid expiration year value: %v (%v)", keyValuePair[1], err)
					break
				}
				passport.ExpirationYear = eyr
			case "hgt":
				passport.Height = keyValuePair[1]
			case "hcl":
				passport.HairColor = keyValuePair[1]
			case "ecl":
				passport.EyeColor = keyValuePair[1]
			case "pid":
				passport.PassportID = keyValuePair[1]
			case "cid":
				passport.CountryID = keyValuePair[1]
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
