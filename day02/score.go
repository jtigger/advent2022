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

func (t Throw) String() string {
	return [...]string{"Rock", "Paper", "Scissors"}[t]
}

type Round struct {
	player1 Throw
	player2 Throw
}

// P2Score turns the points player 2 earned in this Round: 0 if lossed; 3 if tired; 6 if won.
func (r Round) P2Score() int {
	p1Rank, p2Rank := r.rankThrows()

	var p2Score int
	switch {
	case p2Rank < p1Rank:
		p2Score = 0
	case p2Rank == p1Rank:
		p2Score = 3
	case p2Rank > p1Rank:
		p2Score = 6
	}
	return p2Score
}

func (r Round) rankThrows() (int, int) {
	// the three possible throws, ranked; prefixed with what throw loses to Rock and suffixed with what beats Scissor.
	var throwByRank = [...]Throw{Scissor, Rock, Paper, Scissor, Rock}

	p1Rank := int(r.player1) + 1
	var p2Rank int
	for idx := p1Rank - 1; idx <= p1Rank+1; idx++ {
		if throwByRank[idx] == r.player2 {
			p2Rank = idx
		}
	}
	return p1Rank, p2Rank
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

func player2Throw(roundInput string) (Throw, error) {
	switch roundInput[2] {
	case 'X':
		return Rock, nil
	case 'Y':
		return Paper, nil
	case 'Z':
		return Scissor, nil
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
		p2, err := player2Throw(roundInput)
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
