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
	sort.Sort(sort.IntSlice(adapters))
	// The outlet itself has a rating of 0, which is not part of the input.
	adapters = append(adapters, 0)
	// The builtin adapter is always 3 more then the last adapter.
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	skipable := countSkippable(adapters)
	log.Println("---", skipable)
	answer := 0
	for i := 0; i <= skipable; i++ {
		answer = i
	}

	return answer
}

func countSkippable(adapters []int) int {
	skipable := 0
	for i := 0; i < len(adapters); i++ {
		curAdapter := adapters[i]

		for j, nextAdapter := range adapters[i:] {
			if nextAdapter-curAdapter == 3 {
				log.Println(curAdapter, nextAdapter)
				skipable += j - 1
			}
		}
	}

	return skipable
}

// lookAhead looks ahead and returns a slice of all adapters that can be reordered.
// Effectively is searches for a diff of 3, and returns everything before that.
// func lookAhead(adapters []int) []int {
// 	endIndex := 1
// 	for prevIndex, a := range adapters[1:] {
// 		joltDiff := a - adapters[prevIndex]

// 		if joltDiff == 3 {
// 			endIndex = prevIndex + 1
// 			break
// 		}
// 	}

// 	return adapters[:endIndex]
// }

// func findJoltArrangement(adapters []int, preferredJolt int) []int {
// 	newAdapters := []int{}

// 	for i := 0; i < len(adapters); i++ {
// 		curJolt := adapters[i]

// 		for _, j := range []int{1, 2, 3} {
// 			if j+i < len(adapters) {
// 				nextJolts := adapters[j+i]
// 				if nextJolts-curJolt == preferredJolt {
// 					newAdapters = append(newAdapters, nextJolts)
// 					i = +j
// 					break
// 				}
// 			}

// 			newAdapters = append(newAdapters, curJolt)
// 		}
// 	}

// 	return newAdapters
// }

func inputToAdapters(input []string) []int {
	adapters := []int{}
	for _, i := range input {
		adapters = append(adapters, aoc2020.Atoi(i))
	}

	return adapters
}
