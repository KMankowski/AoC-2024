package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		rawCraneGames :=
			`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

		expMinTokens := int64(480)

		outMinTokens, err := run(strings.NewReader(rawCraneGames))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if outMinTokens != expMinTokens {
			t.Errorf("outMinTokens is %v but expMinTokens is %v", outMinTokens, expMinTokens)
		}
	})
}

func TestCalculateMinTokens(t *testing.T) {
	testCases := []struct {
		name         string
		game         craneGame
		expMinTokens int64
	}{
		{
			"example 1",
			craneGame{
				button{94, 34},
				button{22, 67},
				coords{8400, 5400},
			},
			280,
		},
		{
			"example 2",
			craneGame{
				button{26, 66},
				button{67, 21},
				coords{12748, 12176},
			},
			0,
		},
		{
			"example 3",
			craneGame{
				button{17, 86},
				button{84, 37},
				coords{7870, 6450},
			},
			200,
		},
		{
			"example 4",
			craneGame{
				button{69, 23},
				button{27, 71},
				coords{18641, 10279},
			},
			0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outMinTokens := calculateMinTokens(testCase.game)

			if outMinTokens != testCase.expMinTokens {
				t.Errorf("outMinTokens is %v but expMinTokens is %v", outMinTokens, testCase.expMinTokens)
			}
		})
	}
}

func TestParseCraneGames(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		rawCraneGames :=
			`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

		expCraneGames := []craneGame{
			{
				button{
					94, 34,
				},
				button{
					22, 67,
				},
				coords{
					8400, 5400,
				},
			},
			{
				button{
					26, 66,
				},
				button{
					67, 21,
				},
				coords{
					12748, 12176,
				},
			},
			{
				button{
					17, 86,
				},
				button{
					84, 37,
				},
				coords{
					7870, 6450,
				},
			},
			{
				button{
					69, 23,
				},
				button{
					27, 71,
				},
				coords{
					18641, 10279,
				},
			},
		}

		outCraneGames, err := parseCraneGames(strings.NewReader(rawCraneGames))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(outCraneGames) != len(expCraneGames) {
			t.Fatalf("len(outCraneGames) is %v but len(expCraneGames) is %v", len(outCraneGames), len(expCraneGames))
		}

		for i := range expCraneGames {
			if !expCraneGames[i].equals(outCraneGames[i]) {
				t.Fatalf("expCraneGames: %v\noutCraneGames: %v", expCraneGames, outCraneGames)
			}
		}
	})
}
