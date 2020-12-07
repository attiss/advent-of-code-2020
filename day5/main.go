package main

import (
	"bufio"
	"github.com/attiss/advent-of-code-2020/day5/plane_seat"
	"log"
	"os"
	"sort"
)

const (
	inputFile = "input.txt"
)

func main() {
	seats, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	rows := 128
	columns := 8

	var seatIDs []int
	for _, seat := range seats {
		seatID := seat.GetSeatID(rows, columns)
		log.Printf("seat ID for %s is %d", seat, seatID)
		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)
	log.Printf("highest SeatID: %d", seatIDs[len(seatIDs)-1])
}

func readInputFromFile(filePath string) (seats []plane_seat.PlaneSeat, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		seats = append(seats, plane_seat.PlaneSeat(line))
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
