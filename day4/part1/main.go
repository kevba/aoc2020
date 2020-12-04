package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	defer log.Print(time.Since(startTime))

	input, _ := getInput()
	valid := countValid(input)

	log.Printf("anwser: %v", valid)
}

func countValid(ps []passport) int {
	validCounter := 0
	for _, p := range ps {
		if isValidPassword(p) {
			validCounter++
		}
	}

	return validCounter
}

func isValidPassword(p passport) bool {
	if p.byr == "" {
		return false
	}
	if p.iyr == "" {
		return false
	}
	if p.eyr == "" {
		return false
	}
	if p.hgt == "" {
		return false
	}
	if p.hcl == "" {
		return false
	}
	if p.ecl == "" {
		return false
	}
	if p.pid == "" {
		return false
	}
	// if p.cid == "" {
	// 	return false
	// }

	return true
}

var inputFile = "../input.txt"

func getInput() ([]passport, error) {
	values := []passport{}
	file, err := os.Open(inputFile)

	if err != nil {
		return values, err
	}

	s := bufio.NewScanner(file)

	passwordLine := ""
	for s.Scan() {
		line := s.Text()

		if line != "" {
			if passwordLine != "" {
				passwordLine += " "
			}

			passwordLine += line

		} else {
			values = append(values, parseLine(passwordLine))
			passwordLine = ""
		}
	}

	return values, nil
}

func parseLine(line string) passport {
	p := passport{}

	parts := strings.Split(line, " ")

	for _, propval := range parts {
		keyval := strings.Split(propval, ":")

		key := keyval[0]
		val := keyval[1]

		switch key {
		case "byr":
			p.byr = val
		case "iyr":
			p.iyr = val
		case "eyr":
			p.eyr = val
		case "hgt":
			p.hgt = val
		case "hcl":
			p.hcl = val
		case "ecl":
			p.ecl = val
		case "pid":
			p.pid = val
		case "cid":
			p.cid = val
		}

	}

	return p
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}
