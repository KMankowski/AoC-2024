package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
)

func main() {
	outputError := outputErrorFunc()

	inputFile, err := os.Open("input.txt")
	outputError(err)

	puzzleOutput, err := run(inputFile)
	outputError(err)

	fmt.Printf("%v", puzzleOutput)
}

func run(puzzleInput io.Reader) (int, error) {
	list1, list2, err := readLists(puzzleInput)
	if err != nil {
		return 0, err
	}

	distances, err := calculateDistances(list1, list2)
	if err != nil {
		return 0, err
	}

	totalDistance := calculateDistance(distances)

	return totalDistance, nil
}

func readLists(puzzleInput io.Reader) ([]int, []int, error) {
	scanner := bufio.NewScanner(puzzleInput)
	scanner.Split(bufio.ScanWords)

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	isAppendToFirstList := true
	for scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			return nil, nil, err
		}

		if isAppendToFirstList {
			list1 = append(list1, num)
		} else {
			list2 = append(list2, num)
		}

		isAppendToFirstList = !isAppendToFirstList
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	slices.Sort(list1)
	slices.Sort(list2)

	return list1, list2, nil
}

func calculateDistances(list1 []int, list2 []int) ([]int, error) {
	if len(list1) != len(list2) {
		return nil, errors.New("input lists are not the same length")
	}

	distances := make([]int, len(list1))

	for i := range distances {
		distance := list1[i] - list2[i]
		distances[i] = int(math.Abs(float64(distance)))
	}

	return distances, nil
}

func calculateDistance(distances []int) int {
	sum := 0
	for _, distance := range distances {
		sum += distance
	}
	return sum
}

func outputErrorFunc() func(error) {
	errNum := 0
	return func(err error) {
		errNum++
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(errNum)
		}
	}
}
