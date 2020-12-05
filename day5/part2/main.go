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
	seatNum := findMySeat(seats)

	log.Printf("anwser: %v", seatNum)
}

var maxID = (128 * 8) + 8

func findMySeat(seats map[int]seat) int {
	for id := 0; id <= maxID; id++ {
		if _, ok := seats[id]; !ok {
			if checkIfMine(id, seats) {
				return id
			}
		}
	}

	return 0
}

func checkIfMine(id int, seats map[int]seat) bool {
	_, ok1 := seats[id-1]
	_, ok2 := seats[id+1]

	return ok1 && ok2
}

func getSeatsFromInput(input []string) map[int]seat {
	seats := map[int]seat{}

	for _, seatString := range input {

		row := getRowNum(seatString)
		col := getColNum(seatString)
		seat := seat{row: row, col: col}
		seats[seat.id()] = seat

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
