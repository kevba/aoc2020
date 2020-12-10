package main

import (
	"aoc2020"
	"log"
	"sort"
	"time"
)

func main() {
	startTime := time.Now()
	defer log.Print(time.Since(startTime))

	input := aoc2020.GetInput()
	adapters := inputToAdapters(input)

	answer := solve(adapters)

	log.Printf("anwser: %v", answer)
}

func solve(adapters []int) int {
	// The outlet itself has a rating of 0, which is not part of the input.
	adapters = append(adapters, 0)
	sort.Sort(sort.IntSlice(adapters))

	oneJoltDIffCounter := 0
	threeJoltDIffCounter := 0

	for prevIndex, a := range adapters[1:] {
		joltDiff := a - adapters[prevIndex]

		switch joltDiff {
		case 1:
			oneJoltDIffCounter++
		case 3:
			threeJoltDIffCounter++
		default:
			log.Printf("joltdiff %v is %v, this adapter does not fit", prevIndex, joltDiff)
		}
	}

	// The built in adapter is always 3 higher then the highest
	threeJoltDIffCounter++

	return oneJoltDIffCounter * threeJoltDIffCounter
}

func inputToAdapters(input []string) []int {
	adapters := []int{}
	for _, i := range input {
		adapters = append(adapters, aoc2020.Atoi(i))
	}

	return adapters
}
