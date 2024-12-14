package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open input.txt: %v\n", err)
		os.Exit(1)
	}

	sum, err := runPartTwo(string(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in run(): %v\n", err)
		os.Exit(2)
	}

	fmt.Fprintf(os.Stdout, "sum: %v\n", sum)
}

func runPartTwo(input string) (int, error) {
	instructions, err := parseMulAndIfInstructions(input)
	if err != nil {
		return 0, err
	}

	sum := 0
	isEnabled := true
	for _, instruction := range instructions {
		if instruction == "do()" {
			isEnabled = true
		} else if instruction == "don't()" {
			isEnabled = false
		} else if isEnabled {
			mulArg1, mulArg2, err := parseMulArgs(instruction)
			if err != nil {
				return 0, err
			}
			sum += mulArg1 * mulArg2
		}
	}

	return sum, nil
}

func parseMulAndIfInstructions(input string) ([]string, error) {
	r, err := regexp.Compile("mul\\(\\d+,\\d+\\)|don't\\(\\)|do\\(\\)")
	if err != nil {
		return nil, err
	}

	return r.FindAllString(input, -1), nil
}

func run(input string) (int, error) {
	mulInstructions, err := parseMulInstructions(input)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, mulInstruction := range mulInstructions {
		mulArg1, mulArg2, err := parseMulArgs(mulInstruction)
		if err != nil {
			return 0, err
		}
		sum += mulArg1 * mulArg2
	}

	return sum, nil
}

func parseMulArgs(mulInstruction string) (int, int, error) {
	r, err := regexp.Compile("\\d+")
	if err != nil {
		return 0, 0, err
	}

	rawMulArgs := r.FindAllString(mulInstruction, 2)
	if len(rawMulArgs) != 2 {
		return 0, 0, fmt.Errorf("unexpected length not 2 for rawMulArgs: %v\n", rawMulArgs)
	}

	mulArg1, err := strconv.Atoi(rawMulArgs[0])
	if err != nil {
		return 0, 0, err
	}
	mulArg2, err := strconv.Atoi(rawMulArgs[1])
	if err != nil {
		return 0, 0, err
	}

	return mulArg1, mulArg2, nil
}

func parseMulInstructions(input string) ([]string, error) {
	r, err := regexp.Compile("mul\\(\\d+,\\d+\\)")
	if err != nil {
		return nil, err
	}

	return r.FindAllString(input, -1), nil
}
