package main

import (
	"slices"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name                    string
		inpRawInitialStoneState string
		inpNumBlinks            int
		expFinalStoneCount      int
	}{
		{
			"puzzle example 1 input yields 7 stones",
			"0 1 10 99 999",
			1,
			7,
		},
		{
			"puzzle example 2 input with 6 blinks yields 22 stones",
			"125 17",
			6,
			22,
		},
		{
			"puzzle example 2 input with 25 blinks yields 55312 stones",
			"125 17",
			25,
			55312,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outFinalStoneCount, err := run(strings.NewReader(testCase.inpRawInitialStoneState), testCase.inpNumBlinks)

			if err != nil {
				t.Fatalf("unexpected error in run(): %v\n", err)
			}

			if outFinalStoneCount != testCase.expFinalStoneCount {
				t.Errorf("outFinalStoneCount %v expFinalStoneCount %v", outFinalStoneCount, testCase.expFinalStoneCount)
			}
		})
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outFinalStoneCount, err := runFast(strings.NewReader(testCase.inpRawInitialStoneState), testCase.inpNumBlinks)

			if err != nil {
				t.Fatalf("unexpected error in runFast(): %v\n", err)
			}

			if outFinalStoneCount != testCase.expFinalStoneCount {
				t.Errorf("outFinalStoneCount %v expFinalStoneCount %v", outFinalStoneCount, testCase.expFinalStoneCount)
			}
		})
	}
}

func TestReadInitialStoneState(t *testing.T) {
	testCases := []struct {
		name                    string
		inpRawInitialStoneState string
		expInitialStoneState    []int
	}{
		{
			"puzzle example 1 input yields correct initialStoneState",
			"0 1 10 99 999",
			[]int{0, 1, 10, 99, 999},
		},
		{
			"puzzle example 2 input yields correct initialStoneState",
			"125 17",
			[]int{125, 17},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outInitialStoneState, err := readInitialStoneState(strings.NewReader(testCase.inpRawInitialStoneState))

			if err != nil {
				t.Fatalf("unexpected error in readInitialStoneState(): %v\n", err)
			}

			if !slices.Equal(outInitialStoneState, testCase.expInitialStoneState) {
				t.Errorf("outInitialStoneState %v expInitialStoneState %v", outInitialStoneState, testCase.expInitialStoneState)
			}
		})
	}
}

func TestCalculateStoneState(t *testing.T) {
	testCases := []struct {
		name                 string
		inpInitialStoneState []int
		inpNumBlinks         int
		expFinalStoneState   []int
	}{
		{
			"puzzle example 1 input yields correct finalStoneState",
			[]int{0, 1, 10, 99, 999},
			1,
			[]int{1, 2024, 1, 0, 9, 9, 2021976},
		},
		{
			"puzzle example 2 input yields correct finalStoneState",
			[]int{125, 17},
			6,
			[]int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outFinalStoneState, err := calculateStoneState(testCase.inpInitialStoneState, testCase.inpNumBlinks)

			if err != nil {
				t.Fatalf("unexpected error in calculateStoneState(): %v\n", err)
			}

			if !slices.Equal(outFinalStoneState, testCase.expFinalStoneState) {
				t.Errorf("outFinalStoneState %v expFinalStoneState %v", outFinalStoneState, testCase.expFinalStoneState)
			}
		})
	}
}
