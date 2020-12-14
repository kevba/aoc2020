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

	// The builtin adapter is always 3 more then the last adapter.
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	skipable := countSkippable(adapters)
	answer := prod(skipable)

	return answer
}

func countSkippable(adapters []int) []int {
	skipSets := []int{}

	for i := 0; i < len(adapters); i++ {
		possibleSkipSet := lookAhead(adapters[i:])
		i += len(possibleSkipSet)
		skipSets = append(skipSets, countPossibleWithSkip(possibleSkipSet))
	}

	log.Println(skipSets)
	return skipSets
}

func countPossibleWithSkip(adapters []int) int {

	count := 0

	if len(adapters) == 2 {
		return 4
	}

	possibleSlices := [][]int{}

	for i := 1; i < len(adapters)-1; i++ {
		newAdapterSlice := append(adapters[:i], adapters[i:]...)
		possibleSlices = append(possibleSlices, newAdapterSlice)
	}

	for _, al := range possibleSlices {
		for i, a := range al[1:] {
			if a-al[i] < 3 {
				count++
			}
		}
	}

	return count
}

// lookAhead looks ahead and returns a slice of all adapters that can be reordered.
// Effectively is searches for a diff of 3, and returns everything before that.
func lookAhead(adapters []int) []int {
	endIndex := 1
	for prevIndex, a := range adapters[1:] {
		joltDiff := a - adapters[prevIndex]

		if joltDiff == 3 {
			endIndex = prevIndex
			break
		}
	}

	return adapters[:endIndex]
}

func fac(n int) int {
	n2 := 0
	for i := 0; i <= n; i++ {
		n2 += n - i
	}

	return n2
}

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

func sum(slice []int) int {
	p := 0
	for _, n := range slice {
		p = p + n
	}

	return p
}

func prod(slice []int) int {
	p := 1
	for _, n := range slice {
		p = p * n
	}

	return p
}
