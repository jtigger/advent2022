package main

import (
	"bufio"
	"fmt"
	"os"
)

func isOccludedNorth(grid [][]int, treeRow, treeCol int) (treesVisible int, isOccluded bool) {
	treeHeight := grid[treeRow][treeCol]
	for gridRow := treeRow - 1; gridRow >= 0; gridRow-- {
		if grid[gridRow][treeCol] >= treeHeight {
			// e.g. tree at row 2, occluded by tree at row 0; treesVisible = 2
			return treeRow - gridRow, true
		}
	}
	return treeRow, false
}

func isOccludedSouth(grid [][]int, treeRow, treeCol int) (treesVisible int, isOccluded bool) {
	treeHeight := grid[treeRow][treeCol]
	for gridRow := treeRow + 1; gridRow < len(grid); gridRow++ {
		if grid[gridRow][treeCol] >= treeHeight {
			// e.g. tree at row 8, occluded by tree at row 10; treesVisible = 2
			return gridRow - treeRow, true
		}
	}
	// e.g. tree at row 8 on a grid with 12 rows; treesVisible = 3
	return (len(grid) - 1) - treeRow, false
}

func isOccludedEast(grid [][]int, treeRow, treeCol int) (treesVisible int, isOccluded bool) {
	treeHeight := grid[treeRow][treeCol]
	for gridCol := treeCol + 1; gridCol < len(grid[treeRow]); gridCol++ {
		if grid[treeRow][gridCol] >= treeHeight {
			return gridCol - treeCol, true
		}
	}
	return (len(grid[treeRow]) - 1) - treeCol, false
}

func isOccludedWest(grid [][]int, treeRow, treeCol int) (treesVisible int, isOccluded bool) {
	treeHeight := grid[treeRow][treeCol]
	for gridCol := treeCol - 1; gridCol >= 0; gridCol-- {
		if grid[treeRow][gridCol] >= treeHeight {
			return treeCol - gridCol, true
		}
	}
	return treeCol, false
}

type Visibility struct {
	Visible     bool
	ScenicScore int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var heightGrid [][]int
	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		line := scanner.Text()
		heightGrid = append(heightGrid, make([]int, len(line)))
		for colIdx, char := range line {
			heightGrid[rowIdx][colIdx] = int(char - '0')
		}
	}

	numVisible := 0
	maxScenicScore := 0
	var visGrid [][]Visibility
	for rowIdx, row := range heightGrid {
		visGrid = append(visGrid, make([]Visibility, len(row)))
		for colIdx, _ := range row {
			treesToNorth, occludedNorth := isOccludedNorth(heightGrid, rowIdx, colIdx)
			treesToSouth, occludedSouth := isOccludedSouth(heightGrid, rowIdx, colIdx)
			treesToEast, occludedEast := isOccludedEast(heightGrid, rowIdx, colIdx)
			treesToWest, occludedWest := isOccludedWest(heightGrid, rowIdx, colIdx)

			visGrid[rowIdx][colIdx] = Visibility{
				!occludedNorth || !occludedSouth || !occludedEast || !occludedWest,
				treesToNorth * treesToSouth * treesToEast * treesToWest,
			}
			if visGrid[rowIdx][colIdx].Visible {
				numVisible++
			}
			if visGrid[rowIdx][colIdx].ScenicScore > maxScenicScore {
				maxScenicScore = visGrid[rowIdx][colIdx].ScenicScore
			}
		}
	}
	fmt.Printf("Number of visible trees: %d\n", numVisible)
	fmt.Printf("Max scenic score: %d\n", maxScenicScore)
}
