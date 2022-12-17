package rope

import (
	"fmt"
	"io"
	"math"
)

type Position struct {
	X, Y int
}

type Knot struct {
	pos      Position
	follower *Knot
	trail    map[Position]bool
}

func NewKnot(follower *Knot) *Knot {
	k := &Knot{}
	k.follower = follower
	k.trail = make(map[Position]bool)
	k.trail[k.pos] = true
	return k
}

func (k *Knot) Follow(lead Position) {
	if int(math.Abs(float64(lead.X-k.pos.X))) > 1 ||
		int(math.Abs(float64(lead.Y-k.pos.Y))) > 1 {
		if k.pos.X < lead.X {
			k.pos.X += 1
		}
		if k.pos.X > lead.X {
			k.pos.X -= 1
		}
		if k.pos.Y < lead.Y {
			k.pos.Y += 1
		}
		if k.pos.Y > lead.Y {
			k.pos.Y -= 1
		}
		k.trail[k.pos] = true
		if k.follower != nil {
			k.follower.Follow(k.pos)
		}
	}
}

func (k *Knot) Right() {
	k.Follow(Position{k.pos.X + 2, k.pos.Y})
}

func (k *Knot) Left() {
	k.Follow(Position{k.pos.X - 2, k.pos.Y})
}

func (k *Knot) Up() {
	k.Follow(Position{k.pos.X, k.pos.Y + 2})
}

func (k *Knot) Down() {
	k.Follow(Position{k.pos.X, k.pos.Y - 2})
}

type Length struct {
	head   *Knot
	tail   *Knot
	Output io.Writer
}

func (l *Length) Printf(format string, a ...any) {
	if l.Output != nil {
		_, err := fmt.Fprintf(l.Output, format, a...)
		if err != nil {
			panic(err)
		}
	}
}

func NewLength(numKnots int) *Length {
	g := &Length{}
	g.tail = NewKnot(nil)
	prev := g.tail
	for c := 1; c < numKnots; c++ {
		prev = NewKnot(prev)
	}
	g.head = prev
	return g
}

func (l *Length) TailJourneyLength() int {
	return len(l.tail.trail)
}

func (l *Length) AsGrid() []string {
	minX, maxX := -5, 5
	minY, maxY := -5, 5
	for pos, _ := range l.head.trail {
		if pos.X < minX {
			minX = pos.X
		}
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y < minY {
			minY = pos.Y
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}
	var grid []string
	for rowIdx := maxY; rowIdx >= minY; rowIdx-- {
		row := ""
		for colIdx := minX; colIdx <= maxX; colIdx++ {
			knot := l.head
			knotIdx := 0
			knotFound := false
			for knot != nil && !knotFound {
				if knot.pos.X == colIdx && knot.pos.Y == rowIdx {
					row += fmt.Sprintf("%d", knotIdx)
					knotFound = true
				}
				knot = knot.follower
				knotIdx++
			}
			if !knotFound {
				switch {
				case l.tail.trail[Position{colIdx, rowIdx}]:
					row += "*"
				case colIdx == 0 && rowIdx == 0:
					row += "+"
				case rowIdx == 0:
					row += "-"
				case colIdx == 0:
					row += "|"
				default:
					row += "."
				}
			}
		}
		grid = append(grid, row)
	}
	return grid
}

func (l *Length) PrintGrid() {
	if l.Output != nil {
		for _, row := range l.AsGrid() {
			l.Printf("%s\n", row)
		}
		l.Printf("\n")
	}
}

func (l *Length) Right(steps int) {
	l.Printf("=======\nR %d\n", steps)
	for step := 1; step <= steps; step++ {
		l.head.Right()
		l.PrintGrid()
	}
}

func (l *Length) Left(steps int) {
	l.Printf("=======\nL %d\n", steps)
	for step := 1; step <= steps; step++ {
		l.head.Left()
		l.PrintGrid()
	}
}

func (l *Length) Up(steps int) {
	l.Printf("=======\nU %d\n", steps)
	for step := 1; step <= steps; step++ {
		l.head.Up()
		l.PrintGrid()
	}
}

func (l *Length) Down(steps int) {
	l.Printf("=======\nD %d\n", steps)
	for step := 1; step <= steps; step++ {
		l.head.Down()
		l.PrintGrid()
	}
}
