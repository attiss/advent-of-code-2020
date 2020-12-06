package main

import (
	"bufio"
	"log"
	"os"

	"github.com/attiss/advent-of-code-2020/day3/treemap"
)

const (
	inputFile = "input.txt"
)

func main() {
	treeMap, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	checkPaths := []struct {
		right int
		down  int
	}{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	productOfTrees := 1
	for _, path := range checkPaths {
		trees := treeMap.CountPathTrees(path.right, path.down)
		log.Printf("path (right %d; down %d) encounters %d trees", path.right, path.down, trees)
		productOfTrees *= trees
	}
	log.Printf("product of trees: %d", productOfTrees)
}

func readInputFromFile(filePath string) (treeMap treemap.TreeMap, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var rowTrees []bool
		for _, e := range line {
			if string(e) == "#" {
				rowTrees = append(rowTrees, true)
			} else {
				rowTrees = append(rowTrees, false)
			}
		}
		treeMap = append(treeMap, rowTrees)
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
