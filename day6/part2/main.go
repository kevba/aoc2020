package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {
	startTime := time.Now()
	defer log.Print(time.Since(startTime))

	input, _ := getInput()

	groups := inputToGroups(input)
	answer := solve(groups)

	log.Printf("anwser: %v", answer)
}

func solve(groups []group) int {
	yesCount := 0

	for _, g := range groups {
		yesCount += len(g.getYes())
	}

	return yesCount
}

type group struct {
	answers []string
}

func (g group) getYes() []string {
	yesList := []string{}

	for _, c := range alphabet {
		allYes := true
		for _, anwers := range g.answers {
			if !strings.Contains(anwers, c) {
				allYes = false
				break
			}
		}
		if allYes {
			yesList = append(yesList, c)
		}
	}

	return yesList
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

func inputToGroups(input []string) []group {
	groups := []group{}

	curGroup := group{}
	for i, line := range input {
		if line != "" {
			curGroup.answers = append(curGroup.answers, line)
		}

		if line == "" || i >= len(input)-1 {
			groups = append(groups, curGroup)
			curGroup = group{}
		}
	}

	return groups
}
