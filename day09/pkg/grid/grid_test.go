package grid_test

import (
	"testing"

	"day09/pkg/grid"
)

func TestGrid_TailJourneyLength(t *testing.T) {
	type TestCase struct {
		name   string
		prog   func(grid *grid.Grid)
		length int
	}
	testCases := []TestCase{
		{
			"Moving right, tail follows head",
			func(grid *grid.Grid) {
				grid.Right(4)
			},
			4,
		},
		{
			"Moving left, tail follows head",
			func(grid *grid.Grid) {
				grid.Left(3)
			},
			3,
		},
		{
			"Moving up, tail follows head",
			func(grid *grid.Grid) {
				grid.Up(5)
			},
			5,
		},
		{
			"Moving down, tail follows head",
			func(grid *grid.Grid) {
				grid.Up(6)
			},
			6,
		},
		{
			"Moving diagonal, tail falls in line behind head",
			func(grid *grid.Grid) {
				grid.Right(1)
				grid.Up(1)
				grid.Right(1)
				grid.Up(1)
			},
			2,
		},
		{
			"When the tail revisits a position, it is *not* counted again",
			func(grid *grid.Grid) {
				grid.Up(6)
				grid.Down(6)
			},
			6,
		},
		{
			"Example from Advent of Code works",
			func(grid *grid.Grid) {
				grid.Right(4)
				grid.Up(4)
				grid.Left(3)
				grid.Down(1)
				grid.Right(4)
				grid.Down(1)
				grid.Left(5)
				grid.Right(2)
			},
			13,
		},
	}
	for idx, tc := range testCases {
		rope := grid.NewGrid()
		tc.prog(rope)
		if rope.TailJourneyLength() != tc.length {
			t.Fatalf("Test %d: \"%s\"\n  expected tail trail of length %d; was %d",
				idx, tc.name, tc.length, rope.TailJourneyLength())
		}
	}
}
