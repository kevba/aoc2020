package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	defer log.Print(time.Since(startTime))

	input, _ := getInput()
	a, b, c := findSumFor(input, 2020)

	log.Printf("Result is %v * %v * %v = %v", a, b, c, a*b*c)
}

func findSumFor(input []int, sumValue int) (int, int, int) {
	for _, i := range input {
		if i > sumValue {
			continue
		}

		for _, j := range input {
			if i+j > sumValue {
				continue
			}

			for _, k := range input {
				if i+j+k == sumValue {
					return i, j, k
				}
			}

		}
	}

	return 0, 0, 0
}

var inputFile = "../input.txt"

func getInput() ([]int, error) {
	values := []int{}
	file, err := os.Open(inputFile)

	if err != nil {
		return values, err
	}

	s := bufio.NewScanner(file)

	for s.Scan() {
		val, _ := strconv.Atoi(s.Text())
		values = append(values, val)
	}

	return values, nil
}
