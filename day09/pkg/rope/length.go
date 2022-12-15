package rope

import (
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

func (k *Knot) Right() {
	k.pos.X += 1
	if k.follower != nil {
		k.follower.FollowRight(k.pos)
	}
}

func (k *Knot) FollowRight(lead Position) {
	if int(math.Abs(float64(lead.X-k.pos.X))) > 1 {
		k.pos.X += 1
		k.pos.Y = lead.Y
		k.trail[k.pos] = true
	}
	if k.follower != nil {
		k.follower.FollowRight(k.pos)
	}
}

func (k *Knot) Left() {
	k.pos.X -= 1
	if k.follower != nil {
		k.follower.FollowLeft(k.pos)
	}
}

func (k *Knot) FollowLeft(lead Position) {
	if int(math.Abs(float64(lead.X-k.pos.X))) > 1 {
		k.pos.X -= 1
		k.pos.Y = lead.Y
		k.trail[k.pos] = true
	}
	if k.follower != nil {
		k.follower.FollowLeft(k.pos)
	}
}

func (k *Knot) Up() {
	k.pos.Y += 1
	if k.follower != nil {
		k.follower.FollowUp(k.pos)
	}
}

func (k *Knot) FollowUp(lead Position) {
	if int(math.Abs(float64(lead.Y-k.pos.Y))) > 1 {
		k.pos.Y += 1
		k.pos.X = lead.X
		k.trail[k.pos] = true
	}
	if k.follower != nil {
		k.follower.FollowUp(k.pos)
	}
}

func (k *Knot) Down() {
	k.pos.Y -= 1
	if k.follower != nil {
		k.follower.FollowDown(k.pos)
	}
}

func (k *Knot) FollowDown(lead Position) {
	if int(math.Abs(float64(lead.Y-k.pos.Y))) > 1 {
		k.pos.Y -= 1
		k.pos.X = lead.X
		k.trail[k.pos] = true
	}
	if k.follower != nil {
		k.follower.FollowDown(k.pos)
	}
}

type Length struct {
	head *Knot
	tail *Knot
}

func NewLength() *Length {
	numKnots := 2
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
