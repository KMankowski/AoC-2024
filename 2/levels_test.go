package main

import (
	"slices"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("given puzzle example, the correct number of safe reports is found", func(t *testing.T) {
		inputReports :=
			`7 6 4 2 1
		 1 2 7 8 9
		 9 7 6 2 1
		 1 3 2 4 5
		 8 6 4 4 1
		 1 3 6 7 9`

		want := 2
		got, err := run(strings.NewReader(inputReports))

		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestReadReports(t *testing.T) {
	t.Run("given puzzle example, each report is correctly populated into a []int", func(t *testing.T) {
		inputReports :=
			`7 6 4 2 1
			 1 2 7 8 9
			 9 7 6 2 1
			 1 3 2 4 5
			 8 6 4 4 1
			 1 3 6 7 9`

		want := [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		}

		got, err := readReports(strings.NewReader(inputReports))

		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if len(got) != len(want) {
			t.Errorf("expected %v reports, got %v reports", len(want), len(got))
		}

		isCorrect := true
		for i, report := range want {
			if !slices.Equal(report, want[i]) {
				isCorrect = false
				break
			}
		}

		if !isCorrect {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestIsSafe(t *testing.T) {
	t.Run("given ascending safe report, isSafe() returns true", func(t *testing.T) {
		inpReport := []int{1, 3, 6, 7, 9}

		want := true
		got := isSafe(inpReport)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("given ascending unsafe report, isSafe() returns false", func(t *testing.T) {
		inpReport := []int{1, 2, 6, 7, 9}

		want := false
		got := isSafe(inpReport)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("given descending safe report, isSafe() returns true", func(t *testing.T) {
		inpReport := []int{7, 6, 4, 2, 1}

		want := true
		got := isSafe(inpReport)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("given descending unsafe report, isSafe() returns false", func(t *testing.T) {
		inpReport := []int{10, 6, 4, 2, 1}

		want := false
		got := isSafe(inpReport)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("given report that is neither ascending nor descending, isSafe() returns false", func(t *testing.T) {
		inpReport := []int{7, 6, 4, 5, 2}

		want := false
		got := isSafe(inpReport)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
