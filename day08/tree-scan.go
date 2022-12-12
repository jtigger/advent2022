package main

import (
	"bufio"
	"fmt"
	"os"
)

func isOccludedNorth(grid [][]int, rowIdx, colIdx int) bool {
	treeHeight := grid[rowIdx][colIdx]
	for gridRow := rowIdx - 1; gridRow >= 0; gridRow-- {
		if grid[gridRow][colIdx] >= treeHeight {
			return true
		}
	}
	return false
}

func isOccludedSouth(grid [][]int, rowIdx, colIdx int) bool {
	treeHeight := grid[rowIdx][colIdx]
	for gridRow := rowIdx + 1; gridRow < len(grid); gridRow++ {
		if grid[gridRow][colIdx] >= treeHeight {
			return true
		}
	}
	return false
}

func isOccludedEast(grid [][]int, rowIdx, colIdx int) bool {
	treeHeight := grid[rowIdx][colIdx]
	for gridCol := colIdx + 1; gridCol < len(grid[rowIdx]); gridCol++ {
		if grid[rowIdx][gridCol] >= treeHeight {
			return true
		}
	}
	return false
}

func isOccludedWest(grid [][]int, rowIdx, colIdx int) bool {
	treeHeight := grid[rowIdx][colIdx]
	for gridCol := colIdx - 1; gridCol >= 0; gridCol-- {
		if grid[rowIdx][gridCol] >= treeHeight {
			return true
		}
	}
	return false
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
	var visGrid [][]bool
	for rowIdx, row := range heightGrid {
		visGrid = append(visGrid, make([]bool, len(row)))
		for colIdx, _ := range row {
			visNorth := !isOccludedNorth(heightGrid, rowIdx, colIdx)
			visSouth := !isOccludedSouth(heightGrid, rowIdx, colIdx)
			visEast := !isOccludedEast(heightGrid, rowIdx, colIdx)
			visWest := !isOccludedWest(heightGrid, rowIdx, colIdx)
			visGrid[rowIdx][colIdx] = visNorth || visSouth || visEast || visWest
			if visGrid[rowIdx][colIdx] {
				numVisible++
			}
		}
	}
	fmt.Printf("Number of visible trees: %d\n", numVisible)
}
