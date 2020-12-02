package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	defer log.Print(time.Since(startTime))

	input, _ := getInput()

	valid := getValid(input)

	log.Printf("Valid password count: %v", len(valid))
}

func getValid(passwordLines []passwordLine) []passwordLine {
	result := []passwordLine{}

	for _, pl := range passwordLines {
		if validatePassword(pl) {
			result = append(result, pl)
		}
	}

	return result
}

func validatePassword(pl passwordLine) bool {
	if pl.max-1 > len(pl.password) {
		return false
	}

	char1 := string(pl.password[pl.min-1])
	char2 := string(pl.password[pl.max-1])

	if char1 == pl.char && char2 != pl.char {
		return true
	}

	if char1 != pl.char && char2 == pl.char {
		return true
	}

	return false

}

var inputFile = "../input.txt"

func getInput() ([]passwordLine, error) {
	values := []passwordLine{}
	file, err := os.Open(inputFile)

	if err != nil {
		return values, err
	}

	s := bufio.NewScanner(file)

	for s.Scan() {
		pass := parsePasswordLine(s.Text())
		values = append(values, pass)
	}

	return values, nil
}

var passwordLineRegex = regexp.MustCompile("([0-9]*)-([0-9]*) ([a-z]): ([a-z]*)")

func parsePasswordLine(line string) passwordLine {
	res := passwordLineRegex.FindStringSubmatch(line)

	min, _ := strconv.Atoi(res[1])
	max, _ := strconv.Atoi(res[2])

	return passwordLine{
		min:      min,
		max:      max,
		char:     res[3],
		password: res[4],
	}
}

type passwordLine struct {
	min      int
	max      int
	char     string
	password string
}
