package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

type craneGame struct {
	buttonA button
	buttonB button
	prize   coords
}

type button struct {
	xMove int64
	yMove int64
}

type coords struct {
	x int64
	y int64
}

// 33427
func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in Open(): %v", err)
		os.Exit(1)
	}

	// minTokens, err := run(inputFile)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error in run(): %v", err)
	// 	os.Exit(2)
	//}

	minTokens, err := runPartTwo(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in runPartTwo(): %v", err)
		os.Exit(3)
	}

	fmt.Printf("minTokens: %v\n", minTokens)
}

func runPartTwo(r io.Reader) (int64, error) {
	craneGames, err := parseCraneGames(r)
	if err != nil {
		return 0, err
	}

	minTokens := int64(0)
	for _, craneGame := range craneGames {
		craneGame.prize.x += 10000000000000
		craneGame.prize.y += 10000000000000
		minTokens += calculateMinTokens(craneGame)
	}

	return minTokens, nil
}

func run(r io.Reader) (int64, error) {
	craneGames, err := parseCraneGames(r)
	if err != nil {
		return 0, err
	}

	minTokens := int64(0)
	for _, craneGame := range craneGames {
		minTokens += calculateMinTokens(craneGame)
	}

	return minTokens, nil
}

func calculateMinTokens(game craneGame) int64 {
	determinant := float64((game.buttonA.xMove * game.buttonB.yMove) - (game.buttonA.yMove * game.buttonB.xMove))
	dA := float64((game.prize.x * game.buttonB.yMove) - (game.prize.y * game.buttonB.xMove))
	dB := float64((game.buttonA.xMove * game.prize.y) - (game.buttonA.yMove * game.prize.x))

	aPushes := dA / determinant
	bPushes := dB / determinant

	if !isInteger(aPushes) || !isInteger(bPushes) {
		return 0
	}

	minTokens := (3 * int64(aPushes)) + (1 * int64(bPushes))

	return minTokens
}

func isInteger(f float64) bool {
	return f == float64(int64(math.Floor(f)))
}

func parseCraneGames(r io.Reader) ([]craneGame, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	pattern, err := regexp.Compile("\\d+")
	if err != nil {
		return nil, err
	}

	craneGames := make([]craneGame, 0, 320)
	for scanner.Scan() {
		nextCraneGame := craneGame{}

		rawButtonALine := scanner.Text()
		rawButtonA := pattern.FindAllString(rawButtonALine, -1)
		if len(rawButtonA) != 2 {
			return nil, fmt.Errorf("parseCraneGames() regex pattern matched %v strings in %q", len(rawButtonA), rawButtonALine)
		}

		aX, err := strconv.Atoi(rawButtonA[0])
		if err != nil {
			return nil, err
		}
		nextCraneGame.buttonA.xMove = int64(aX)

		aY, err := strconv.Atoi(rawButtonA[1])
		if err != nil {
			return nil, err
		}
		nextCraneGame.buttonA.yMove = int64(aY)

		scanner.Scan()
		rawButtonBLine := scanner.Text()
		rawButtonB := pattern.FindAllString(rawButtonBLine, -1)
		if len(rawButtonB) != 2 {
			return nil, fmt.Errorf("parseCraneGames() regex pattern matched %v strings in %q", len(rawButtonB), rawButtonBLine)
		}

		bX, err := strconv.Atoi(rawButtonB[0])
		if err != nil {
			return nil, err
		}
		nextCraneGame.buttonB.xMove = int64(bX)

		bY, err := strconv.Atoi(rawButtonB[1])
		if err != nil {
			return nil, err
		}
		nextCraneGame.buttonB.yMove = int64(bY)

		scanner.Scan()
		rawPrizeLine := scanner.Text()
		rawPrize := pattern.FindAllString(rawPrizeLine, -1)
		if len(rawPrize) != 2 {
			return nil, fmt.Errorf("parseCraneGames() regex pattern matched %v strings in %q", len(rawPrize), rawPrizeLine)
		}

		x, err := strconv.Atoi(rawPrize[0])
		if err != nil {
			return nil, err
		}
		nextCraneGame.prize.x = int64(x)

		y, err := strconv.Atoi(rawPrize[1])
		if err != nil {
			return nil, err
		}
		nextCraneGame.prize.y = int64(y)

		// Empty line
		scanner.Scan()

		craneGames = append(craneGames, nextCraneGame)
	}

	return craneGames, nil
}

func (g1 craneGame) equals(g2 craneGame) bool {
	if g1.buttonA.xMove != g2.buttonA.xMove {
		return false
	}
	if g1.buttonA.yMove != g2.buttonA.yMove {
		return false
	}
	if g1.buttonB.xMove != g2.buttonB.xMove {
		return false
	}
	if g1.buttonB.yMove != g2.buttonB.yMove {
		return false
	}
	if g1.prize.x != g2.prize.x {
		return false
	}
	if g1.prize.y != g2.prize.y {
		return false
	}
	return true
}
