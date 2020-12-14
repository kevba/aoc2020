package aoc2020

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

var inputFile = "../input.txt"

// GetInput returns the puzzle input in rows.
func GetInput() []string {
	values := []string{}
	file, err := os.Open(inputFile)

	if err != nil {
		return []string{}
	}

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		values = append(values, line)
	}

	return values
}

// Atoi converts a string to an int. It fatals when this is not possible.
func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("cannot convert %v to int: %v", s, err)
	}
	return num
}

func Time() func() {
	start := time.Now()
	return func() {
		log.Printf("solved in %v \n", time.Since(start))
	}
}
