package main

import (
	"slices"
	"testing"
)

func TestRunPartTwo(t *testing.T) {
	t.Run("Part 2 example", func(t *testing.T) {
		input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

		want := 48
		got, err := runPartTwo(input)

		if err != nil {
			t.Errorf("unexpected error: %v\n", err)
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestRun(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

		want := 161
		got, err := run(input)

		if err != nil {
			t.Errorf("unexpected error: %v\n", err)
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestParseMulInstructions(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

		want := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
		got, err := parseMulInstructions(input)

		if err != nil {
			t.Errorf("unexpected error: %v\n", err)
		}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestParseMulArgs(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := "mul(2,4)"

		expMulArg1 := 2
		expMulArg2 := 4
		outMulArg1, outMulArg2, err := parseMulArgs(input)

		if err != nil {
			t.Errorf("unexpected error: %v\n", err)
		}

		if outMulArg1 != expMulArg1 {
			t.Errorf("outMulArg1 %v expMulArg1 %v", outMulArg1, expMulArg1)
		}

		if outMulArg2 != expMulArg2 {
			t.Errorf("outMulArg2 %v expMulArg2 %v", outMulArg2, expMulArg2)
		}
	})
}
