package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Instr is a cycle instruction.
type Instr interface {
	PostCycleAction(x *int)
}

type AddX struct {
	addend int
}

func (i AddX) PostCycleAction(x *int) {
	*x += i.addend
}

type Noop struct {
}

func (i Noop) PostCycleAction(x *int) {
}

type Device interface {
	Signal(cycle, x int)
}

type CPU struct {
	prog      []Instr
	registerX int
	devices   []Device
}

func NewCPU() *CPU {
	cpu := &CPU{}
	cpu.registerX = 1
	return cpu
}

func (c *CPU) Load(program []string) error {
	for idx, line := range program {
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "addx":
			if len(tokens) != 2 {
				return fmt.Errorf("on line %d, \"%s\" expected 2 arguments; was %d", idx+1, tokens[0], len(tokens))
			}
			c.prog = append(c.prog, Noop{})
			addend, err := strconv.ParseInt(tokens[1], 10, 32)
			if err != nil {
				return fmt.Errorf("on line %d, \"%s\" expected integer parameter; was \"%s\"; %s", idx+1, tokens[0], tokens[1], err)
			}
			c.prog = append(c.prog, AddX{int(addend)})
		case "noop":
			c.prog = append(c.prog, Noop{})
		case "", "#":
			// allow (and ignore) blank lines or comments
		default:
			return fmt.Errorf("on line %d, unknown instruction \"%s\"", idx+1, tokens[0])
		}
	}
	return nil
}

func (c *CPU) Connect(device Device) {
	c.devices = append(c.devices, device)
}

func (c *CPU) Run() {
	for progIdx := 0; progIdx < len(c.prog); progIdx++ {
		cycleNum := progIdx + 1
		for _, device := range c.devices {
			device.Signal(cycleNum, c.registerX)
		}
		c.prog[progIdx].PostCycleAction(&c.registerX)
	}
}

type Probe struct {
	probe    []int
	probeIdx int
	samples  map[int]int
}

func NewProbe() *Probe {
	p := &Probe{
		[]int{20, 60, 100, 140, 180, 220},
		0,
		make(map[int]int),
	}
	return p
}

func (p *Probe) TotalSignal() int {
	totalSignal := 0
	for cycle, x := range p.samples {
		totalSignal += cycle * x
	}
	return totalSignal
}

func (p *Probe) Signal(cycle, x int) {
	if p.probeIdx < len(p.probe) && p.probe[p.probeIdx] == cycle {
		p.samples[cycle] = x
		p.probeIdx++
	}
}

type Display struct {
}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) Signal(cycle, x int) {
	progIdx := cycle - 1
	pixelPos := progIdx % 40
	spriteLeftEdge := x - 1
	spriteRightEdge := x + 1

	if pixelPos == 0 {
		fmt.Println()
	}
	if pixelPos >= spriteLeftEdge && pixelPos <= spriteRightEdge {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}

func loadListing(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)

	program := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		program = append(program, line)
	}
	return program
}

func main() {
	program := loadListing(os.Stdin)

	cpu := NewCPU()
	probe := NewProbe()
	display := NewDisplay()

	err := cpu.Load(program)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load program: %s", err)
		os.Exit(1)
	}
	cpu.Connect(probe)
	cpu.Connect(display)
	cpu.Run()

	fmt.Printf("\nTotal signal strength: %d\n", probe.TotalSignal())
}
