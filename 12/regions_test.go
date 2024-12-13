package main

import (
	"slices"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("large example", func(t *testing.T) {
		input :=
			`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

		expPrice := 1930

		outPrice, err := run(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if outPrice != expPrice {
			t.Errorf("outPrice %v but expPrice %v", outPrice, expPrice)
		}
	})
	t.Run("X/O example", func(t *testing.T) {
		input :=
			`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
`
		expPrice := 772

		outPrice, err := run(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if outPrice != expPrice {
			t.Errorf("outPrice %v but expPrice %v", outPrice, expPrice)
		}
	})
	t.Run("small example", func(t *testing.T) {
		input :=
			`AAAA
BBCD
BBCC
EEEC`

		expPrice := 140

		outPrice, err := run(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if outPrice != expPrice {
			t.Errorf("outPrice %v but expPrice %v", outPrice, expPrice)
		}
	})
}

func TestRunPartTwo(t *testing.T) {
	t.Run("large example", func(t *testing.T) {
		input :=
			`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

		expPrice := 1206

		outPrice, err := runPartTwo(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if outPrice != expPrice {
			t.Errorf("outPrice %v but expPrice %v", outPrice, expPrice)
		}
	})
	t.Run("X/O example", func(t *testing.T) {
		input :=
			`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
`
		expPrice := 436

		outPrice, err := runPartTwo(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if outPrice != expPrice {
			t.Errorf("outPrice %v but expPrice %v", outPrice, expPrice)
		}
	})
	t.Run("small example", func(t *testing.T) {
		input :=
			`AAAA
BBCD
BBCC
EEEC`

		expPrice := 80

		outPrice, err := runPartTwo(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if outPrice != expPrice {
			t.Errorf("outPrice %v but expPrice %v", outPrice, expPrice)
		}
	})
	t.Run("E-shaped example", func(t *testing.T) {
		input :=
			`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

		expPrice := 236

		outPrice, err := runPartTwo(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if outPrice != expPrice {
			t.Errorf("outPrice %v but expPrice %v", outPrice, expPrice)
		}
	})
	t.Run("A-B example", func(t *testing.T) {
		input :=
			`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

		expPrice := 368

		outPrice, err := runPartTwo(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}

		if outPrice != expPrice {
			t.Errorf("outPrice %v but expPrice %v", outPrice, expPrice)
		}
	})
}

func TestCalculateSides(t *testing.T) {
	testCases := []struct {
		name      string
		plan      [][]plot
		regionIds int
		expSides  []int
	}{
		{
			"example 1",
			[][]plot{
				{{'A', 1}, {'A', 1}, {'A', 1}, {'A', 1}},
				{{'B', 2}, {'B', 2}, {'C', 3}, {'D', 4}},
				{{'B', 2}, {'B', 2}, {'C', 3}, {'C', 3}},
				{{'E', 5}, {'E', 5}, {'E', 5}, {'C', 3}},
			},
			5,
			[]int{4, 4, 8, 4, 4},
		},
		{
			"E-shaped example",
			[][]plot{
				{{'E', 1}, {'E', 1}, {'E', 1}, {'E', 1}, {'E', 1}},
				{{'E', 1}, {'X', 2}, {'X', 2}, {'X', 2}, {'X', 2}},
				{{'E', 1}, {'E', 1}, {'E', 1}, {'E', 1}, {'E', 1}},
				{{'E', 1}, {'X', 3}, {'X', 3}, {'X', 3}, {'X', 3}},
				{{'E', 1}, {'E', 1}, {'E', 1}, {'E', 1}, {'E', 1}},
			},
			3,
			[]int{12, 4, 4},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			regionId := 1
			for regionId <= testCase.regionIds {
				outSides := calculateSides(testCase.plan, regionId)

				if outSides != testCase.expSides[regionId-1] {
					t.Errorf("outSides %v expSides %v for regionId %v", outSides, testCase.expSides[regionId-1], regionId)
				}
				regionId++
			}
		})
	}
}

func TestCalculatePerimeter(t *testing.T) {
	testCases := []struct {
		name          string
		plan          [][]plot
		regionIds     int
		expPerimeters []int
	}{
		{
			"example 1",
			[][]plot{
				{{'A', 1}, {'A', 1}, {'A', 1}, {'A', 1}},
				{{'B', 2}, {'B', 2}, {'C', 3}, {'D', 4}},
				{{'B', 2}, {'B', 2}, {'C', 3}, {'C', 3}},
				{{'E', 5}, {'E', 5}, {'E', 5}, {'C', 3}},
			},
			5,
			[]int{10, 8, 10, 4, 8},
		},
		{
			"example 2",
			[][]plot{
				{{'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}},
				{{'O', 1}, {'X', 2}, {'O', 1}, {'X', 3}, {'O', 1}},
				{{'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}},
				{{'O', 1}, {'X', 4}, {'O', 1}, {'X', 5}, {'O', 1}},
				{{'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}},
			},
			5,
			[]int{36, 4, 4, 4, 4},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			regionId := 1
			for regionId <= testCase.regionIds {
				outPerimeter := calculatePerimeter(testCase.plan, regionId)

				if outPerimeter != testCase.expPerimeters[regionId-1] {
					t.Errorf("outPerimeter %v expPerimeter %v for regionId %v and plan %v", outPerimeter, testCase.expPerimeters[regionId-1], regionId, testCase.plan)
				}
				regionId++
			}
		})
	}
}

func TestCalculateArea(t *testing.T) {
	testCases := []struct {
		name      string
		plan      [][]plot
		regionIds int
		expAreas  []int
	}{
		{
			"example 1",
			[][]plot{
				{{'A', 1}, {'A', 1}, {'A', 1}, {'A', 1}},
				{{'B', 2}, {'B', 2}, {'C', 3}, {'D', 4}},
				{{'B', 2}, {'B', 2}, {'C', 3}, {'C', 3}},
				{{'E', 5}, {'E', 5}, {'E', 5}, {'C', 3}},
			},
			5,
			[]int{4, 4, 4, 1, 3},
		},
		{
			"example 2",
			[][]plot{
				{{'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}},
				{{'O', 1}, {'X', 2}, {'O', 1}, {'X', 3}, {'O', 1}},
				{{'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}},
				{{'O', 1}, {'X', 4}, {'O', 1}, {'X', 5}, {'O', 1}},
				{{'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}},
			},
			5,
			[]int{21, 1, 1, 1, 1},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			regionId := 1
			for regionId <= testCase.regionIds {
				outArea := calculateArea(testCase.plan, regionId)

				if outArea != testCase.expAreas[regionId-1] {
					t.Errorf("outArea %v expArea %v for regionId %v and plan %v", outArea, testCase.expAreas[regionId-1], regionId, testCase.plan)
				}
				regionId++
			}
		})
	}
}

func TestCalculateAndSetRegionIds(t *testing.T) {
	testCases := []struct {
		name           string
		inpPlan        [][]plot
		expPlan        [][]plot
		expRegionCount int
	}{
		{
			"example 1",
			[][]plot{
				{{'A', 0}, {'A', 0}, {'A', 0}, {'A', 0}},
				{{'B', 0}, {'B', 0}, {'C', 0}, {'D', 0}},
				{{'B', 0}, {'B', 0}, {'C', 0}, {'C', 0}},
				{{'E', 0}, {'E', 0}, {'E', 0}, {'C', 0}},
			},
			[][]plot{
				{{'A', 1}, {'A', 1}, {'A', 1}, {'A', 1}},
				{{'B', 2}, {'B', 2}, {'C', 3}, {'D', 4}},
				{{'B', 2}, {'B', 2}, {'C', 3}, {'C', 3}},
				{{'E', 5}, {'E', 5}, {'E', 5}, {'C', 3}},
			},
			5,
		},
		{
			"example 2",
			[][]plot{
				{{'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}},
				{{'O', 0}, {'X', 0}, {'O', 0}, {'X', 0}, {'O', 0}},
				{{'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}},
				{{'O', 0}, {'X', 0}, {'O', 0}, {'X', 0}, {'O', 0}},
				{{'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}},
			},
			[][]plot{
				{{'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}},
				{{'O', 1}, {'X', 2}, {'O', 1}, {'X', 3}, {'O', 1}},
				{{'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}},
				{{'O', 1}, {'X', 4}, {'O', 1}, {'X', 5}, {'O', 1}},
				{{'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}, {'O', 1}},
			},
			5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := calculateAndSetRegionIds(testCase.inpPlan)

			if got != testCase.expRegionCount {
				t.Errorf("got %v want %v", got, testCase.expRegionCount)
			}

			expectEqualPlans(t, testCase.expPlan, testCase.inpPlan)
		})
	}
}

func TestInitializePlan(t *testing.T) {
	testCases := []struct {
		name       string
		inpRawPlan [][]byte
		expPlan    [][]plot
	}{
		{
			"example 1",
			[][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
			[][]plot{
				{{'A', 0}, {'A', 0}, {'A', 0}, {'A', 0}},
				{{'B', 0}, {'B', 0}, {'C', 0}, {'D', 0}},
				{{'B', 0}, {'B', 0}, {'C', 0}, {'C', 0}},
				{{'E', 0}, {'E', 0}, {'E', 0}, {'C', 0}},
			},
		},
		{
			"example 2",
			[][]byte{
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
			},
			[][]plot{
				{{'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}},
				{{'O', 0}, {'X', 0}, {'O', 0}, {'X', 0}, {'O', 0}},
				{{'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}},
				{{'O', 0}, {'X', 0}, {'O', 0}, {'X', 0}, {'O', 0}},
				{{'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}, {'O', 0}},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outPlan := initializePlan(testCase.inpRawPlan)

			if len(outPlan) != len(testCase.expPlan) {
				t.Fatalf("len(outPlan) is %v but len(expPlan) is %v", len(outPlan), len(testCase.expPlan))
			}

			expectEqualPlans(t, testCase.expPlan, outPlan)
		})
	}
}

func TestReadRawPlan(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		expRawPlan [][]byte
	}{
		{
			"example 1",
			`AAAA
BBCD
BBCC
EEEC`,
			[][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
		},
		{
			"example 2",
			`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			[][]byte{
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			input := strings.NewReader(testCase.input)

			outRawPlan := readRawPlan(input)

			if len(outRawPlan) != len(testCase.expRawPlan) {
				t.Fatalf("len(outRawPlan) is %v but len(expRawPlan) is %v", len(outRawPlan), len(testCase.expRawPlan))
			}

			for i := range testCase.expRawPlan {
				if !slices.Equal(outRawPlan[i], testCase.expRawPlan[i]) {
					t.Errorf("outRawPlan: %v\n expRawPlan: %v\n", outRawPlan, testCase.expRawPlan)
					break
				}
			}
		})
	}
}

func expectEqualPlans(t *testing.T, expPlan, outPlan [][]plot) {
	t.Helper()

	for i, row := range expPlan {
		if len(outPlan[i]) != len(expPlan[i]) {
			t.Fatalf("outPlan: %v\n expPlan: %v\n", outPlan, expPlan)
		}
		for j := range row {
			if outPlan[i][j].plant != expPlan[i][j].plant || outPlan[i][j].regionId != expPlan[i][j].regionId {
				t.Fatalf("outPlan: %v\n expPlan: %v\n", outPlan, expPlan)
			}
		}
	}
}
