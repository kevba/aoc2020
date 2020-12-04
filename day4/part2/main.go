package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
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
	if !validBirthYear(p.byr) {
		return false
	}
	if !validIssueYear(p.iyr) {
		return false
	}
	if !validExpireYear(p.eyr) {
		return false
	}
	if !validHeigth(p.hgt) {
		return false
	}
	if !validHairColor(p.hcl) {
		return false
	}
	if !validEyeColor(p.ecl) {
		return false
	}
	if !validPID(p.pid) {
		return false
	}
	// if p.cid == "" {
	// 	return false
	// }

	return true
}

func validBirthYear(byr string) bool {
	return byr != "" && atoi(byr) >= 1920 && atoi(byr) <= 2002
}

func validIssueYear(iyr string) bool {
	return iyr != "" && atoi(iyr) >= 2010 && atoi(iyr) <= 2020
}

func validExpireYear(eyr string) bool {
	return eyr != "" && atoi(eyr) >= 2020 && atoi(eyr) <= 2030
}

func validHeigth(h string) bool {
	if h == "" {
		return false
	}

	if !strings.Contains(h, "in") && !strings.Contains(h, "cm") {
		return false
	}

	numPart := atoi(string(h[0 : len(h)-2]))
	if strings.Contains(h, "cm") && (numPart < 150 || numPart > 193) {
		return false
	}

	if strings.Contains(h, "in") && (numPart < 59 || numPart > 76) {
		return false
	}

	return true
}

func validHairColor(hcl string) bool {
	if hcl == "" {
		return false
	}

	return match("#[a-f0-9]{6}", hcl)
}

func validEyeColor(ecl string) bool {
	if ecl == "" {
		return false
	}

	return strings.Contains("amb blu brn gry grn hzl oth", ecl)
}

func validPID(pid string) bool {
	if pid == "" {
		return false
	}

	return match("^[0-9]{9}$", pid)
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

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func match(p, s string) bool {
	matched, _ := regexp.MatchString(p, s)

	return matched
}
