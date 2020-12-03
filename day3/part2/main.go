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

	c1 := countTreesForSlope(input, 1, 1)
	c2 := countTreesForSlope(input, 3, 1)
	c3 := countTreesForSlope(input, 5, 1)
	c4 := countTreesForSlope(input, 7, 1)
	c5 := countTreesForSlope(input, 1, 2)

	answer := c1 * c2 * c3 * c4 * c5

	log.Printf("anwser: %v", answer)
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
