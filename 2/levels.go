package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open input.txt: %v\n", err)
		os.Exit(1)
	}

	safeReportCount, safeReportWithDampenerCount, err := run(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred in run(): %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("safeReportCount: %v\n", safeReportCount)
	fmt.Printf("safeReportWithDampenerCount: %v\n", safeReportWithDampenerCount)
}

func run(input io.Reader) (int, int, error) {
	reports, err := readReports(input)
	if err != nil {
		return 0, 0, err
	}

	safeReportCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReportCount++
		}
	}

	safeReportWithDampenerCount := 0
	for _, report := range reports {
		if isSafeWithDampener(report) {
			safeReportWithDampenerCount++
		}
	}

	return safeReportCount, safeReportWithDampenerCount, nil
}

func readReports(input io.Reader) ([][]int, error) {
	reportScanner := bufio.NewScanner(input)
	reportScanner.Split(bufio.ScanLines)

	reports := make([][]int, 0)

	for reportScanner.Scan() {
		rawReport := reportScanner.Text()

		levelScanner := bufio.NewScanner(strings.NewReader(rawReport))
		levelScanner.Split(bufio.ScanWords)

		report := make([]int, 0)

		for levelScanner.Scan() {
			rawLevel := levelScanner.Text()

			level, err := strconv.Atoi(rawLevel)
			if err != nil {
				return nil, err
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func isSafe(report []int) bool {
	isAscending := report[1]-report[0] >= 0
	for i := range report[:len(report)-1] {
		difference := report[i+1] - report[i]
		if isAscending {
			if difference <= 0 || difference > 3 {
				return false
			}
		} else {
			if difference >= 0 || difference < -3 {
				return false
			}
		}
	}
	return true
}

func isSafeWithDampener(report []int) bool {
	if isSafe(report[1:]) {
		return true
	}

	isAscending := report[1]-report[0] >= 0
	for i := range report[:len(report)-1] {
		difference := report[i+1] - report[i]
		if isAscending {
			if difference <= 0 || difference > 3 {
				return isSafe(slices.Concat(report[:i+1], report[i+2:])) || isSafe(slices.Concat(report[:i], report[i+1:]))
			}
		} else {
			if difference >= 0 || difference < -3 {
				return isSafe(slices.Concat(report[:i+1], report[i+2:])) || isSafe(slices.Concat(report[:i], report[i+1:]))
			}
		}
	}
	return true
}
