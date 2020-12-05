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
	seats := getSeatsFromInput(input)
	highest := getHighestID(seats)

	log.Printf("anwser: %v", highest)
}

func getSeatsFromInput(input []string) []seat {
	seats := []seat{}

	for _, seatString := range input {

		row := getRowNum(seatString)
		col := getColNum(seatString)

		seats = append(seats, seat{row: row, col: col})

	}

	return seats
}

func getRowNum(str string) int {
	rowCount := 127
	_, upper := div(str[0:7], 0, rowCount)
	return upper
}

func getColNum(str string) int {
	colCount := 7
	_, upper := div(str[7:10], 0, colCount)
	return upper
}

func div(str string, lower, upper int) (int, int) {
	if len(str) == 0 {
		return lower, upper
	}

	curChar := str[0]
	remain := str[1:]

	modval := ((upper - lower) + 1) / 2
	if curChar == 'F' || curChar == 'L' {
		upper = upper - modval
	}
	if curChar == 'B' || curChar == 'R' {
		lower = lower + modval
	}

	return div(remain, lower, upper)
}

func getHighestID(seats []seat) int {
	n := 0

	for _, s := range seats {
		if s.id() > n {
			n = s.id()
		}
	}
	return n
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

type seat struct {
	row int
	col int
}

func (s *seat) id() int {
	return s.row*8 + s.col
}
