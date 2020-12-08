package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	defer log.Print(time.Since(startTime))

	input, _ := getInput()
	instructions := inputToInstructions(input)

	answer := solve(instructions)

	log.Printf("anwser: %v", answer)
}

func solve(instructions []*instruction) int {
	execIndex := 0
	acc := 0

	for {
		instruction := instructions[execIndex]
		instruction.execCount++

		if instruction.execCount > 1 {
			break
		}

		switch instruction.operation {
		case "nop":
			execIndex++
			break
		case "jmp":
			execIndex += instruction.argument
			break
		case "acc":
			execIndex++
			acc += instruction.argument
			break
		}

	}

	return acc
}

func inputToInstructions(input []string) []*instruction {
	instructions := []*instruction{}

	for _, in := range input {
		op := in[:3]
		arg := atoi(in[4:])

		instructions = append(instructions, &instruction{
			operation: op,
			argument:  arg,
		})
	}

	return instructions
}

type instruction struct {
	operation string
	argument  int
	execCount int
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
