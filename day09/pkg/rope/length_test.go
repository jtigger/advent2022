package rope_test

import (
	"testing"

	"day09/pkg/rope"
)

func TestGrid_TailJourneyLength(t *testing.T) {
	type TestCase struct {
		name   string
		prog   func(length *rope.Length)
		length int
	}
	testCases := []TestCase{
		{
			"Moving right, tail follows head",
			func(length *rope.Length) {
				length.Right(4)
			},
			4,
		},
		{
			"Moving left, tail follows head",
			func(length *rope.Length) {
				length.Left(3)
			},
			3,
		},
		{
			"Moving up, tail follows head",
			func(length *rope.Length) {
				length.Up(5)
			},
			5,
		},
		{
			"Moving down, tail follows head",
			func(length *rope.Length) {
				length.Up(6)
			},
			6,
		},
		{
			"Moving diagonal, tail falls in line behind head",
			func(length *rope.Length) {
				length.Right(1)
				length.Up(1)
				length.Right(1)
				length.Up(1)
			},
			2,
		},
		{
			"When the tail revisits a position, it is *not* counted again",
			func(length *rope.Length) {
				length.Up(6)
				length.Down(6)
			},
			6,
		},
		{
			"Example from Advent of Code works",
			func(length *rope.Length) {
				length.Right(4)
				length.Up(4)
				length.Left(3)
				length.Down(1)
				length.Right(4)
				length.Down(1)
				length.Left(5)
				length.Right(2)
			},
			13,
		},
	}
	for idx, tc := range testCases {
		rope := rope.NewLength()
		tc.prog(rope)
		if rope.TailJourneyLength() != tc.length {
			t.Fatalf("Test %d: \"%s\"\n  expected tail trail of length %d; was %d",
				idx, tc.name, tc.length, rope.TailJourneyLength())
		}
	}
}
