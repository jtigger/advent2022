package grid

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X, Y int
}

type Grid struct {
	head      Position
	tail      Position
	tailTrail map[Position]bool
}

func NewGrid() *Grid {
	g := &Grid{}
	g.tailTrail = make(map[Position]bool)
	g.tailTrail[g.tail] = true
	return g
}

func (g *Grid) TailJourneyLength() int {
	return len(g.tailTrail)
}

func (g *Grid) Right(steps int) {
	for step := 1; step <= steps; step++ {
		g.head.X += 1
		if int(math.Abs(float64(g.head.X-g.tail.X))) > 1 {
			g.tail.X += 1
			g.tail.Y = g.head.Y
			g.tailTrail[g.tail] = true
		}
	}
}

func (g *Grid) Left(steps int) {
	for step := 1; step <= steps; step++ {
		g.head.X -= 1
		if int(math.Abs(float64(g.head.X-g.tail.X))) > 1 {
			g.tail.X -= 1
			g.tail.Y = g.head.Y
			g.tailTrail[g.tail] = true
		}
	}
}

func (g *Grid) Up(steps int) {
	for step := 1; step <= steps; step++ {
		g.head.Y += 1
		if int(math.Abs(float64(g.head.Y-g.tail.Y))) > 1 {
			g.tail.Y += 1
			g.tail.X = g.head.X
			g.tailTrail[g.tail] = true
		}
	}
}

func (g *Grid) Down(steps int) {
	for step := 1; step <= steps; step++ {
		g.head.Y -= 1
		if int(math.Abs(float64(g.head.Y-g.tail.Y))) > 1 {
			g.tail.Y -= 1
			g.tail.X = g.head.X
			g.tailTrail[g.tail] = true
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := NewGrid()
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		length64, err := strconv.ParseInt(tokens[1], 10, 32)
		if err != nil {
			log.Panicf("Unable to parse input \"%s\"", line)
		}
		distance := int(length64)
		switch tokens[0] {
		case "U":
			grid.Up(distance)
		case "D":
			grid.Down(distance)
		case "L":
			grid.Left(distance)
		case "R":
			grid.Right(distance)
		default:
			log.Panicf("Unknown direction %s", tokens[0])
		}
	}
	fmt.Printf("Total positions: %d\n", grid.TailJourneyLength())
}
