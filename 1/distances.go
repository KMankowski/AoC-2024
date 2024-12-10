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

	puzzleOutputPart1, puzzleOutputPart2, err := run(inputFile)
	outputError(err)

	fmt.Printf("part 1: %v\npart 2: %v\n", puzzleOutputPart1, puzzleOutputPart2)
}

func run(puzzleInput io.Reader) (int, int, error) {
	list1, list2, err := readLists(puzzleInput)
	if err != nil {
		return 0, 0, err
	}

	// Part 1
	distances, err := calculateDistances(list1, list2)
	if err != nil {
		return 0, 0, err
	}

	totalDistance := calculateDistance(distances)

	// Part 2
	similarity, err := calculateSimilarity(list1, list2)
	if err != nil {
		return 0, 0, err
	}

	return totalDistance, similarity, nil
}

func calculateSimilarity(list1 []int, list2 []int) (int, error) {
	if len(list1) != len(list2) {
		return 0, errors.New("input lists are not the same length to similarity")
	}

	similarityScore := 0
	pointer1 := 0
	pointer2 := 0
	for pointer1 != len(list1) && pointer2 != len(list2) {
		list1Val := list1[pointer1]
		list2Val := list2[pointer2]

		if list1Val < list2Val {
			pointer1++
			continue
		}

		if list1Val > list2Val {
			pointer2++
			continue
		}

		tempPointer2 := pointer2
		for tempPointer2 != len(list2) {
			if list1Val == list2[tempPointer2] {
				similarityScore += list1Val
				tempPointer2++
			} else {
				break
			}
		}
		pointer1++
	}

	return similarityScore, nil
}

func calculateDistances(list1 []int, list2 []int) ([]int, error) {
	if len(list1) != len(list2) {
		return nil, errors.New("input lists are not the same length to distances")
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
