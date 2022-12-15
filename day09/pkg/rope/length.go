package rope

import (
	"math"
)

type Position struct {
	X, Y int
}

type Knot struct {
	pos   Position
	tail  *Knot
	trail map[Position]bool
}

func NewKnot(tail *Knot) *Knot {
	k := &Knot{}
	k.tail = tail
	k.trail = make(map[Position]bool)
	k.trail[k.pos] = true
	return k
}

func (k *Knot) Right() {
	k.pos.X += 1
	if int(math.Abs(float64(k.pos.X-k.tail.pos.X))) > 1 {
		k.tail.pos.X += 1
		k.tail.pos.Y = k.pos.Y
		k.tail.trail[k.tail.pos] = true
	}
}

func (k *Knot) Left() {
	k.pos.X -= 1
	if int(math.Abs(float64(k.pos.X-k.tail.pos.X))) > 1 {
		k.tail.pos.X -= 1
		k.tail.pos.Y = k.pos.Y
		k.tail.trail[k.tail.pos] = true
	}
}

func (k *Knot) Up() {
	k.pos.Y += 1
	if int(math.Abs(float64(k.pos.Y-k.tail.pos.Y))) > 1 {
		k.tail.pos.Y += 1
		k.tail.pos.X = k.pos.X
		k.tail.trail[k.tail.pos] = true
	}
}

func (k *Knot) Down() {
	k.pos.Y -= 1
	if int(math.Abs(float64(k.pos.Y-k.tail.pos.Y))) > 1 {
		k.tail.pos.Y -= 1
		k.tail.pos.X = k.pos.X
		k.tail.trail[k.tail.pos] = true
	}
}

type Length struct {
	head *Knot
	tail *Knot
}

func NewLength() *Length {
	l := &Length{}
	l.tail = NewKnot(nil)
	l.head = NewKnot(l.tail)
	return l
}

func (l *Length) TailJourneyLength() int {
	return len(l.tail.trail)
}

func (l *Length) Right(steps int) {
	for step := 1; step <= steps; step++ {
		l.head.Right()
	}
}

func (l *Length) Left(steps int) {
	for step := 1; step <= steps; step++ {
		l.head.Left()
	}
}

func (l *Length) Up(steps int) {
	for step := 1; step <= steps; step++ {
		l.head.Up()
	}
}

func (l *Length) Down(steps int) {
	for step := 1; step <= steps; step++ {
		l.head.Down()
	}
}
