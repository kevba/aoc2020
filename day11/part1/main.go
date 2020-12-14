package main

import (
	"aoc2020"
	"fmt"
	"log"
)

func main() {
	defer aoc2020.Time()()

	input := aoc2020.GetInput()
	g := inputToGrid(input)
	answer := solve(g)
	log.Printf("anwser: %v", answer)
}

func solve(g *grid) int {
	for {
		g = round(g)

		if g.totalChanges == 0 {
			return countType(g.seats, occ)
		}
	}
}

var occ = '#'
var empty = 'L'
var floor = '.'

func round(g *grid) *grid {
	newGrid := &grid{
		width:        g.width,
		seats:        make([]rune, len(g.seats)),
		totalChanges: 0,
	}

	for i, curSeat := range g.seats {
		neighbours := getNeighbours(g, i)
		occCount := countType(neighbours, occ)

		if curSeat == empty && occCount == 0 {
			newGrid.totalChanges++
			newGrid.seats[i] = occ
		} else if curSeat == occ && occCount >= 4 {
			newGrid.totalChanges++
			newGrid.seats[i] = empty
		} else {
			newGrid.seats[i] = curSeat
		}

	}

	return newGrid
}

func countType(seats []rune, t rune) int {
	c := 0
	for _, s := range seats {
		if s == t {
			c++
		}
	}

	return c
}

func getNeighbours(g *grid, index int) []rune {
	neighbours := []rune{}
	indexes := []int{}

	mask := []struct {
		row int
		col int
	}{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	relativeIndex := index % g.width

	for _, m := range mask {
		rowMod := m.row * g.width
		neighbourIndex := index - m.col + rowMod
		relativeNeighbourIndex := relativeIndex - m.col

		if neighbourIndex < 0 || neighbourIndex > len(g.seats)-1 {
			continue
		}

		if relativeNeighbourIndex < 0 || relativeNeighbourIndex >= g.width {
			continue
		}

		indexes = append(indexes, neighbourIndex)
	}

	for _, i := range indexes {
		if i == index {
			continue
		}
		neighbours = append(neighbours, g.seats[i])
	}

	return neighbours
}

type grid struct {
	width        int
	seats        []rune
	totalChanges int
}

func (g *grid) Print() {
	for i, s := range g.seats {
		if i%g.width == 0 {
			fmt.Print("\n")
		}
		fmt.Print(string(s))
	}
	fmt.Print("\n")
}

func inputToGrid(input []string) *grid {
	newGrid := &grid{
		width: len(input[0]),
		seats: []rune{},
	}

	for _, line := range input {
		for _, c := range line {
			newGrid.seats = append(newGrid.seats, c)
		}
	}

	return newGrid
}
