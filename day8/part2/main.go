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
	attemptCounter := 0

	for {
		acc, err := attemptReplaceOperation("nop", "jmp", attemptCounter, instructions)

		if err == nil {
			return acc
		}

		if err.msg == "cannot find more instructions" {
			break
		}

		attemptCounter++
	}

	attemptCounter = 0
	for {
		acc, err := attemptReplaceOperation("jmp", "nop", attemptCounter, instructions)

		if err == nil {
			return acc
		}

		if err.msg == "cannot find more instructions" {
			break
		}

		attemptCounter++
	}

	return -1
}

func runBootcode(instructions []*instruction) (int, *runError) {
	defer func() {
		// Clear runtime info
		for _, ins := range instructions {
			ins.execCount = 0
		}
	}()

	execIndex := 0
	acc := 0

	lastInstructionIdex := len(instructions) - 1

	for {
		if execIndex > lastInstructionIdex {
			break
		}

		instruction := instructions[execIndex]
		instruction.execCount++

		if instruction.execCount > 1 {
			return acc, newRunError("loop detected", execIndex)
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

	return acc, nil
}

func attemptReplaceOperation(op1, op2 string, count int, instructions []*instruction) (int, *runError) {
	opIndex := findOperationIndex(op1, count, instructions)
	if opIndex < 0 {
		return 0, newRunError("cannot find more instructions", -1)
	}

	instructions[opIndex].operation = op2
	acc, runErr := runBootcode(instructions)
	instructions[opIndex].operation = op1

	return acc, runErr
}

func findOperationIndex(op string, count int, instructions []*instruction) int {
	counter := 0

	for i, ins := range instructions {
		if ins.operation == op {
			if counter == count {
				return i
			}
			counter++
		}
	}

	return -1
}

func attemptFixNOPtoJMP(instructions []*instruction, attempt int) []*instruction {
	nopCount := 0

	for _, ins := range instructions {
		if ins.operation == "nop" {
			nopCount++
			if nopCount > attempt {
				ins.operation = "jmp"
			}
		}
	}

	return instructions
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

func newRunError(msg string, index int) *runError {
	return &runError{
		msg:   msg,
		index: index,
	}
}

type runError struct {
	msg   string
	index int
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
