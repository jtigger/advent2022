package rope_test

import (
	"testing"

	"day09/pkg/rope"
)

type TestCase struct {
	name   string
	prog   func(length *rope.Length)
	length int
}

func TestLength_2Knots_TailJourneyLength(t *testing.T) {
	testCases := []TestCase{
		{
			"Moving right, follower follows head",
			func(length *rope.Length) {
				length.Right(4)
			},
			4,
		},
		{
			"Moving left, follower follows head",
			func(length *rope.Length) {
				length.Left(3)
			},
			3,
		},
		{
			"Moving up, follower follows head",
			func(length *rope.Length) {
				length.Up(5)
			},
			5,
		},
		{
			"Moving down, follower follows head",
			func(length *rope.Length) {
				length.Up(6)
			},
			6,
		},
		{
			"Moving diagonal, follower falls in line behind head",
			func(length *rope.Length) {
				length.Right(1)
				length.Up(1)
				length.Right(1)
				length.Up(1)
			},
			2,
		},
		{
			"When the follower revisits a position, it is *not* counted again",
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
		rope := rope.NewLength(2)
		tc.prog(rope)
		if rope.TailJourneyLength() != tc.length {
			t.Fatalf("Test %d: \"%s\"\n  expected tail trail of length %d; was %d",
				idx, tc.name, tc.length, rope.TailJourneyLength())
		}
	}
}

func TestLength_10Knots_TailJourneyLength(t *testing.T) {
	testCases := []TestCase{
		{
			"Example with tail movement (from Advent of Code)",
			func(length *rope.Length) {
				length.Right(5)
				length.Up(8)
				length.Left(8)
				length.Down(3)
				length.Right(17)
				length.Down(10)
				length.Left(25)
				length.Up(20)
			},
			36,
		},
	}
	for idx, tc := range testCases {
		rope := rope.NewLength(10)
		tc.prog(rope)
		if rope.TailJourneyLength() != tc.length {
			t.Fatalf("Test %d: \"%s\"\n  expected tail trail of length %d; was %d",
				idx, tc.name, tc.length, rope.TailJourneyLength())
		}
	}
}
