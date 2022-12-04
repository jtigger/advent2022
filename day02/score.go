package main

import (
	"bufio"
	"fmt"
	"os"
)

type Throw int

const (
	Rock Throw = iota
	Paper
	Scissor
)

var throwByRank = [...]Throw{Scissor, Rock, Paper, Scissor, Rock}

func (t Throw) String() string {
	return [...]string{"Rock", "Paper", "Scissors"}[t]
}

func (t Throw) Beats() Throw {
	return throwByRank[int(t+1)-1]
}

func (t Throw) BeatenBy() Throw {
	return throwByRank[int(t+1)+1]
}

type Round struct {
	player1 Throw
	player2 Throw
}

// P2Score turns the points player 2 earned in this Round: 0 if lossed; 3 if tired; 6 if won.
func (r Round) P2Score() int {
	var p2Score int
	switch {
	case r.player2.BeatenBy() == r.player1:
		p2Score = 0
	case r.player2 == r.player1:
		p2Score = 3
	case r.player2.Beats() == r.player1:
		p2Score = 6
	}
	return p2Score
}

func player1Throw(roundInput string) (Throw, error) {
	switch roundInput[0] {
	case 'A':
		return Rock, nil
	case 'B':
		return Paper, nil
	case 'C':
		return Scissor, nil
	default:
		return 0, fmt.Errorf("could not find player one's throw in \"%s\" (expected one of \"A\", \"B\", \"C\")", roundInput)
	}
}

func player2Throw(p1Throw Throw, roundInput string) (Throw, error) {
	switch roundInput[2] {
	case 'X':
		return p1Throw.Beats(), nil
	case 'Y':
		return p1Throw, nil
	case 'Z':
		return p1Throw.BeatenBy(), nil
	default:
		return 0, fmt.Errorf("could not find player two's throw in \"%s\" (expected one of \"X\", \"Y\", \"Z\")", roundInput)
	}
}

func bonus(player2 Throw) int {
	return int(player2) + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	line := 0
	totalScore := 0
	for scanner.Scan() {
		line += 1
		roundInput := scanner.Text()
		p1, err := player1Throw(roundInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Problem with round #%d: %s", line, err)
			os.Exit(1)
		}
		p2, err := player2Throw(p1, roundInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Problem with round #%d: %s", line, err)
			os.Exit(1)
		}

		round := Round{p1, p2}
		score := round.P2Score() + bonus(p2)
		totalScore += score
	}
	fmt.Printf("Total score: %d\n", totalScore)
}
