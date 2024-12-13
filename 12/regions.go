package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type plot struct {
	plant    byte
	regionId int
}

type sides struct {
	isLeftProcessed   bool
	isRightProcessed  bool
	isTopProcessed    bool
	isBottomProcessed bool
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
		fmt.Fprintf(os.Stderr, "error reading input.txt: %v\n", err)
	}

	// totalPrice, err := run(inputFile)
	// if err != nil {
	// 	os.Exit(2)
	// 	fmt.Fprintf(os.Stderr, "error in run(): %v\n", err)
	// }

	// fmt.Printf("totalPrice: %v\n", totalPrice)

	totalPrice, err := runPartTwo(inputFile)
	if err != nil {
		os.Exit(3)
		fmt.Fprintf(os.Stderr, "error in runPartTwo(): %v\n", err)
	}

	fmt.Printf("totalPrice: %v\n", totalPrice)
}

func run(input io.Reader) (int, error) {
	rawPlan := readRawPlan(input)
	plan := initializePlan(rawPlan)
	regionIdsSet := calculateAndSetRegionIds(plan)

	totalCost := 0
	regionId := 1
	for regionId <= regionIdsSet {
		area := calculateArea(plan, regionId)
		perimeter := calculatePerimeter(plan, regionId)

		cost := area * perimeter
		totalCost += cost

		regionId++
	}

	return totalCost, nil
}

func runPartTwo(input io.Reader) (int, error) {
	rawPlan := readRawPlan(input)
	plan := initializePlan(rawPlan)
	regionIdsSet := calculateAndSetRegionIds(plan)

	totalCost := 0
	regionId := 1
	for regionId <= regionIdsSet {
		area := calculateArea(plan, regionId)
		sides := calculateSides(plan, regionId)

		cost := area * sides
		totalCost += cost

		regionId++
	}

	return totalCost, nil
}

func calculateSides(plan [][]plot, regionId int) int {
	sidesProcessed := make([][]sides, len(plan))
	for rowIndex := range sidesProcessed {
		sidesProcessed[rowIndex] = make([]sides, len(plan[rowIndex]))
	}

	sides := 0
	for rowIndex, row := range plan {
		for plotIndex, plot := range row {
			if plot.regionId == regionId {
				sides += processLeft(plan, sidesProcessed, regionId, rowIndex, plotIndex)
				sides += processRight(plan, sidesProcessed, regionId, rowIndex, plotIndex)
				sides += processTop(plan, sidesProcessed, regionId, rowIndex, plotIndex)
				sides += processBottom(plan, sidesProcessed, regionId, rowIndex, plotIndex)
			}
		}
	}

	return sides
}

func processRight(plan [][]plot, sidesProcessed [][]sides, regionId, rowIndex, plotIndex int) int {
	if sidesProcessed[rowIndex][plotIndex].isRightProcessed {
		return 0
	}
	sidesProcessed[rowIndex][plotIndex].isRightProcessed = true

	if plotIndex+1 >= len(plan[rowIndex]) || plan[rowIndex][plotIndex+1].regionId != regionId {
		row := rowIndex - 1
		for row >= 0 {
			// If the next plot up is not in this region, stop going up
			if plan[row][plotIndex].regionId != regionId {
				break
			}
			// If the right neighbor of the next plot up does not share this same side, stop going up
			if plotIndex+1 < len(plan[row]) && plan[row][plotIndex+1].regionId == regionId {
				break
			}
			sidesProcessed[row][plotIndex].isRightProcessed = true
			row--
		}
		row = rowIndex + 1
		for row < len(plan) {
			if plan[row][plotIndex].regionId != regionId {
				break
			}
			if plotIndex+1 < len(plan[row]) && plan[row][plotIndex+1].regionId == regionId {
				break
			}
			sidesProcessed[row][plotIndex].isRightProcessed = true
			row++
		}
		return 1
	}

	return 0
}

func processTop(plan [][]plot, sidesProcessed [][]sides, regionId, rowIndex, plotIndex int) int {
	if sidesProcessed[rowIndex][plotIndex].isTopProcessed {
		return 0
	}
	sidesProcessed[rowIndex][plotIndex].isTopProcessed = true

	if rowIndex-1 < 0 || plan[rowIndex-1][plotIndex].regionId != regionId {
		column := plotIndex - 1
		for column >= 0 {
			// If the next plot up is not in this region, stop going up
			if plan[rowIndex][column].regionId != regionId {
				break
			}
			// If the right neighbor of the next plot up does not share this same side, stop going up
			if rowIndex-1 >= 0 && plan[rowIndex-1][column].regionId == regionId {
				break
			}
			sidesProcessed[rowIndex][column].isTopProcessed = true
			column--
		}
		column = plotIndex + 1
		for column < len(plan) {
			if plan[rowIndex][column].regionId != regionId {
				break
			}
			if rowIndex-1 >= 0 && plan[rowIndex-1][column].regionId == regionId {
				break
			}
			sidesProcessed[rowIndex][column].isTopProcessed = true
			column++
		}
		return 1
	}

	return 0
}

func processBottom(plan [][]plot, sidesProcessed [][]sides, regionId, rowIndex, plotIndex int) int {
	if sidesProcessed[rowIndex][plotIndex].isBottomProcessed {
		return 0
	}
	sidesProcessed[rowIndex][plotIndex].isBottomProcessed = true

	if rowIndex+1 >= len(plan) || plan[rowIndex+1][plotIndex].regionId != regionId {
		column := plotIndex - 1
		for column >= 0 {
			// If the next plot up is not in this region, stop going up
			if plan[rowIndex][column].regionId != regionId {
				break
			}
			// If the right neighbor of the next plot up does not share this same side, stop going up
			if rowIndex+1 < len(plan) && plan[rowIndex+1][column].regionId == regionId {
				break
			}
			sidesProcessed[rowIndex][column].isBottomProcessed = true
			column--
		}
		column = plotIndex + 1
		for column < len(plan) {
			if plan[rowIndex][column].regionId != regionId {
				break
			}
			if rowIndex+1 < len(plan) && plan[rowIndex+1][column].regionId == regionId {
				break
			}
			sidesProcessed[rowIndex][column].isBottomProcessed = true
			column++
		}
		return 1
	}

	return 0
}

func processLeft(plan [][]plot, sidesProcessed [][]sides, regionId, rowIndex, plotIndex int) int {
	if sidesProcessed[rowIndex][plotIndex].isLeftProcessed {
		return 0
	}
	sidesProcessed[rowIndex][plotIndex].isLeftProcessed = true

	if plotIndex-1 < 0 || plan[rowIndex][plotIndex-1].regionId != regionId {
		row := rowIndex - 1
		for row >= 0 {
			// If the next plot up is not in this region, stop going up
			if plan[row][plotIndex].regionId != regionId {
				break
			}
			// If the left neighbor of the next plot up does not share this same side, stop going up
			if plotIndex-1 >= 0 && plan[row][plotIndex-1].regionId == regionId {
				break
			}
			sidesProcessed[row][plotIndex].isLeftProcessed = true
			row--
		}
		row = rowIndex + 1
		for row < len(plan) {
			if plan[row][plotIndex].regionId != regionId {
				break
			}
			if plotIndex-1 >= 0 && plan[row][plotIndex-1].regionId == regionId {
				break
			}
			sidesProcessed[row][plotIndex].isLeftProcessed = true
			row++
		}
		return 1
	}

	return 0
}

func calculatePerimeter(plan [][]plot, regionId int) int {
	perimeter := 0
	for rowIndex, row := range plan {
		for plotIndex, plot := range row {
			if plot.regionId == regionId {
				if isFenceNeeded(plan, regionId, rowIndex+1, plotIndex) {
					perimeter++
				}
				if isFenceNeeded(plan, regionId, rowIndex-1, plotIndex) {
					perimeter++
				}
				if isFenceNeeded(plan, regionId, rowIndex, plotIndex+1) {
					perimeter++
				}
				if isFenceNeeded(plan, regionId, rowIndex, plotIndex-1) {
					perimeter++
				}
			}
		}
	}
	return perimeter
}

func isFenceNeeded(plan [][]plot, regionId, rowIndex, plotIndex int) bool {
	if rowIndex < 0 || rowIndex > len(plan)-1 {
		return true
	}
	if plotIndex < 0 || plotIndex > len(plan[rowIndex])-1 {
		return true
	}
	if plan[rowIndex][plotIndex].regionId != regionId {
		return true
	}
	return false
}

func calculateArea(plan [][]plot, regionId int) int {
	area := 0
	for _, row := range plan {
		for _, plot := range row {
			if plot.regionId == regionId {
				area++
			}
		}
	}
	return area
}

func calculateAndSetRegionIds(plan [][]plot) int {
	nextRegionId := 1
	for rowIndex, row := range plan {
		for plotIndex, plot := range row {
			if plot.regionId != 0 {
				continue
			}
			setRegion(plan, plot.plant, nextRegionId, rowIndex, plotIndex)
			nextRegionId++
		}
	}
	return nextRegionId - 1
}

func setRegion(plan [][]plot, plant byte, regionId, rowIndex, plotIndex int) {
	if rowIndex < 0 || rowIndex > len(plan)-1 {
		return
	}
	if plotIndex < 0 || plotIndex > len(plan[rowIndex])-1 {
		return
	}
	if plan[rowIndex][plotIndex].regionId != 0 || plan[rowIndex][plotIndex].plant != plant {
		return
	}

	plan[rowIndex][plotIndex].regionId = regionId

	setRegion(plan, plant, regionId, rowIndex-1, plotIndex)
	setRegion(plan, plant, regionId, rowIndex+1, plotIndex)
	setRegion(plan, plant, regionId, rowIndex, plotIndex-1)
	setRegion(plan, plant, regionId, rowIndex, plotIndex+1)
}

func initializePlan(rawPlan [][]byte) [][]plot {
	plan := make([][]plot, len(rawPlan))
	for rowIndex := range rawPlan {
		plan[rowIndex] = make([]plot, len(rawPlan[rowIndex]))
	}

	for rowIndex, row := range plan {
		for plantIndex := range row {
			plant := rawPlan[rowIndex][plantIndex]
			plan[rowIndex][plantIndex] = plot{plant, 0}
		}
	}

	return plan
}

func readRawPlan(input io.Reader) [][]byte {
	scanner := bufio.NewScanner(input)

	// TODO: replace hard-coded # of lines in file
	rawPlan := make([][]byte, 0, 140)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// .Bytes() did not work here ONLY for very large input... why?
		nextLine := scanner.Text()
		rawPlan = append(rawPlan, []byte(nextLine))
	}

	return rawPlan
}
