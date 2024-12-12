package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open input.txt: %v\n", err)
		os.Exit(1)
	}

	finalStoneCount, err := run(inputFile, 25)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred in run(): %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("finalStoneCount: %v\n", finalStoneCount)

	inputFile.Close()

	inputFile, err = os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open input.txt: %v\n", err)
		os.Exit(3)
	}

	finalStoneCountFast, err := runFast(inputFile, 75)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred in runFast(): %v\n", err)
		os.Exit(4)
	}

	fmt.Printf("finalStoneCountFast: %v\n", finalStoneCountFast)
}

// runFast is an optimized version of run
// Using the fact that a stone splitting in two is independent of every other stone
// Each stone is processed individually
// A DP cache stores the number of children produced by a stones number
func runFast(input io.Reader, numBlinks int) (int, error) {
	initialStoneState, err := readInitialStoneState(input)
	if err != nil {
		return 0, err
	}

	cache := make([]map[int]int, numBlinks)
	for i := range len(cache) {
		cache[i] = make(map[int]int)
	}

	stoneCount, err := getStoneCountFast(initialStoneState, numBlinks, cache)
	if err != nil {
		return 0, err
	}

	return stoneCount, nil
}

func getStoneCountFast(stones []int, blinksRemaining int, cache []map[int]int) (int, error) {
	if blinksRemaining == 0 {
		return len(stones), nil
	}

	stoneCount := 0
	for _, stone := range stones {
		// Check if this stone/blink combo is already in cache
		if _, ok := cache[blinksRemaining-1][stone]; ok {
			stoneCount += cache[blinksRemaining-1][stone]
			continue
		}

		// Add this stone/blink combo to the cache
		if stone == 0 {
			var err error
			cache[blinksRemaining-1][stone], err = getStoneCountFast([]int{1}, blinksRemaining-1, cache)
			if err != nil {
				return 0, err
			}
		} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
			newLeftStone, err := strconv.Atoi(stoneStr[:len(stoneStr)/2])
			if err != nil {
				return 0, err
			}
			newRightStone, err := strconv.Atoi(stoneStr[len(stoneStr)/2:])
			if err != nil {
				return 0, err
			}
			cache[blinksRemaining-1][stone], err = getStoneCountFast([]int{newLeftStone, newRightStone}, blinksRemaining-1, cache)
			if err != nil {
				return 0, err
			}
		} else {
			var err error
			cache[blinksRemaining-1][stone], err = getStoneCountFast([]int{stone * 2024}, blinksRemaining-1, cache)
			if err != nil {
				return 0, err
			}
		}

		// Update stoneCount with the newly added cache value
		stoneCount += cache[blinksRemaining-1][stone]
	}

	return stoneCount, nil
}

func run(input io.Reader, numBlinks int) (int, error) {
	initialStoneState, err := readInitialStoneState(input)
	if err != nil {
		return 0, err
	}

	finalStoneState, err := calculateStoneState(initialStoneState, numBlinks)
	if err != nil {
		return 0, err
	}

	return len(finalStoneState), nil
}

func calculateStoneState(initialStoneState []int, numBlinks int) ([]int, error) {
	if numBlinks < 1 {
		return nil, fmt.Errorf("out of bounds numBlinks: %v", numBlinks)
	}

	currentStoneState := initialStoneState
	for range numBlinks {
		var err error
		currentStoneState, err = iterateStoneState(currentStoneState)
		if err != nil {
			return nil, err
		}
	}
	return currentStoneState, nil
}

func iterateStoneState(currentStoneState []int) ([]int, error) {
	newStoneState := make([]int, 0, 2*len(currentStoneState))

	for _, stone := range currentStoneState {
		if stone == 0 {
			newStoneState = append(newStoneState, 1)
		} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
			newLeftStone, err := strconv.Atoi(stoneStr[:len(stoneStr)/2])
			if err != nil {
				return nil, err
			}
			newRightStone, err := strconv.Atoi(stoneStr[len(stoneStr)/2:])
			if err != nil {
				return nil, err
			}
			newStoneState = append(newStoneState, newLeftStone)
			newStoneState = append(newStoneState, newRightStone)
		} else {
			newStoneState = append(newStoneState, stone*2024)
		}
	}

	return newStoneState, nil
}

func readInitialStoneState(input io.Reader) ([]int, error) {
	initialStoneState := make([]int, 0)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		rawStone := scanner.Text()

		stone, err := strconv.Atoi(rawStone)
		if err != nil {
			return nil, fmt.Errorf("Atoi() failed on rawStone: %v", rawStone)
		}

		initialStoneState = append(initialStoneState, stone)
	}

	return initialStoneState, nil
}
