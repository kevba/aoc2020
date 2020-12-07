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
	bagMap := inputToBagMap(input)

	answer := solve(bagMap)

	log.Printf("anwser: %v", answer)
}

func solve(bags map[string]bag) int {
	ends := findAllPossibleContaining("shiny gold bag", bags)

	// The golden bag does not count of course
	return len(ends) - 1
}

func findAllPossibleContaining(bagName string, bags map[string]bag) map[string]bool {
	ends := map[string]bool{bagName: true}

	if len(bags[bagName].containedBy) == 0 {
		return map[string]bool{bagName: true}
	}

	for _, containedBy := range bags[bagName].containedBy {
		e := findAllPossibleContaining(containedBy, bags)
		for name := range e {
			ends[name] = true
		}
	}

	return ends
}

type bag struct {
	name        string
	containedBy []string
	contains    map[string]int
}

func inputToBagMap(input []string) map[string]bag {
	bagMap := map[string]bag{}

	for _, inLine := range input {
		newBag := lineToBag(inLine)

		if existingBag, ok := bagMap[newBag.name]; ok {
			newBag.containedBy = existingBag.containedBy
		}

		bagMap[newBag.name] = newBag

		for containedName := range newBag.contains {
			containedBag, ok := bagMap[containedName]

			if !ok {
				containedBag = bag{
					name:        containedName,
					containedBy: []string{},
					contains:    map[string]int{},
				}
			}

			containedBag.containedBy = append(containedBag.containedBy, newBag.name)
			bagMap[containedName] = containedBag
		}
	}

	return bagMap
}

var lineRegex = regexp.MustCompile("([a-z]* [a-z]* [a-z]*) contain(.*)")

func lineToBag(in string) bag {
	newBag := bag{
		name:        "",
		containedBy: []string{},
		contains:    map[string]int{},
	}

	match := lineRegex.FindStringSubmatch(in)

	name := match[1]
	name = name[:len(name)-1]
	newBag.name = name

	for _, containBag := range strings.Split(match[2], ",") {
		if containBag == " no other bags." {
			continue
		}

		containBagAmount := atoi(containBag[1:2])
		containBagName := containBag[3:]

		if containBagName[len(containBagName)-1] == '.' {
			containBagName = containBagName[:len(containBagName)-1]
		}

		if containBagAmount > 1 {
			containBagName = containBagName[:len(containBagName)-1]
		}

		newBag.contains[containBagName] = containBagAmount
	}

	return newBag
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
