package main

import (
	"bytes"
	"slices"
	"testing"
)

func TestDistances(t *testing.T) {
	t.Run("Calculate total distance from two lists.", func(t *testing.T) {
		input := bytes.NewReader([]byte(`3   4
										 4   3
										 2   5
										 1   3
										 3   9
										 3   3`))

		want := 11
		got, err := run(input)

		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Populate lists from puzzle input", func(t *testing.T) {
		input := bytes.NewReader([]byte(`3   4
										 4   3
										 2   5
										 1   3
										 3   9
										 3   3`))

		expList1 := []int{1, 2, 3, 3, 3, 4}
		expList2 := []int{3, 3, 3, 4, 5, 9}

		gotList1, gotList2, err := readLists(input)

		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if !slices.Equal(expList1, gotList1) {
			t.Errorf("got %v want %v", gotList1, expList1)
		}

		if !slices.Equal(expList2, gotList2) {
			t.Errorf("got %v want %v", gotList2, expList2)
		}
	})
	t.Run("Calculate distances from lists", func(t *testing.T) {
		inpList1 := []int{1, 2, 3, 3, 3, 4}
		inpList2 := []int{3, 3, 3, 4, 5, 9}

		want := []int{2, 1, 0, 1, 2, 5}
		got, err := calculateDistances(inpList1, inpList2)

		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if !slices.Equal(want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Calculate total distance from distances", func(t *testing.T) {
		input := []int{2, 1, 0, 1, 2, 5}

		want := 11
		got := calculateDistance(input)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
