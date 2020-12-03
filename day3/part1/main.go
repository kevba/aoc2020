package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	defer log.Print(time.Since(startTime))

	input, _ := getInput()

	c := countTreesForSlope(input, 3, 1)

	log.Printf("found trees: %v", c)
}

func countTreesForSlope(lines []treeline, right, down int) int {
	trees := 0

	for i := 0; i*down < len(lines); i++ {
		currentLine := lines[i*down]
		if currentLine.isTree(i * right) {
			trees++
		}
	}

	return trees
}

var inputFile = "../input.txt"

func getInput() ([]treeline, error) {
	values := []treeline{}
	file, err := os.Open(inputFile)

	if err != nil {
		return values, err
	}

	s := bufio.NewScanner(file)

	for s.Scan() {
		values = append(values, treeline(s.Text()))
	}

	return values, nil
}

var tree = "#"

type treeline string

func (t treeline) isTree(index int) bool {
	actualIndex := index % len(t)

	return string(t[actualIndex]) == tree
}
