package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	defer log.Print(time.Since(startTime))

	input, _ := getInput()
	xmas := inputToXMAS(input)

	answer := solve(xmas)

	log.Printf("anwser: %v", answer)
}

func solve(xmas *xmasEncoding) int {
	for {
		isValid, err := xmas.validNext()
		if err != nil {
			log.Fatal(err)
		}
		if !isValid {
			return xmas.numbers[xmas.currentIndex-1]
		}
	}
}

func inputToXMAS(input []string) *xmasEncoding {
	numbers := []int{}

	for _, in := range input {
		numbers = append(numbers, atoi(in))
	}

	xmas := &xmasEncoding{
		numbers:      numbers,
		currentIndex: 25,
		searchLen:    25,
	}

	return xmas
}

type xmasEncoding struct {
	numbers      []int
	currentIndex int
	searchLen    int
}

func (x *xmasEncoding) validNext() (bool, error) {
	defer func() {
		x.currentIndex++
	}()

	if len(x.numbers) < x.currentIndex {
		return false, fmt.Errorf("out of numbers")
	}

	sumSearchSpace := x.numbers[x.currentIndex-x.searchLen : x.currentIndex]
	if !findSum(sumSearchSpace, x.currentNumber()) {
		return false, nil
	}

	return true, nil
}

func (x xmasEncoding) currentNumber() int {
	return x.numbers[x.currentIndex]
}

func findSum(numbers []int, sum int) bool {
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1 != n2 && n1+n2 == sum {
				return true
			}
		}
	}
	return false
}

var inputFile = "../input.txt"

func getInput() ([]string, error) {
	values := []string{}
	file, err := os.Open(inputFile)

	if err != nil {
		return values, err
	}

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		values = append(values, line)
	}

	return values, nil
}

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
