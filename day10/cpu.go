package main

import (
	"bufio"
	"fmt"
	"log"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	prog := []Instr{}

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "addx":
			prog = append(prog, Noop{})
			addend, err := strconv.ParseInt(tokens[1], 10, 32)
			if err != nil {
				log.Panicf("Expected integer parameter for %s; was %s; %s", tokens[0], tokens[1], err)
			}
			prog = append(prog, AddX{int(addend)})
		case "noop":
			prog = append(prog, Noop{})
		default:
			log.Panicf("Unknown opcode %s", tokens[0])
		}
	}

	probe := []int{20, 60, 100, 140, 180, 220}
	probeIdx := 0
	samples := make(map[int]int)

	registerX := 1
	for progIdx := 0; progIdx < 220; progIdx++ {
		cycleNum := progIdx + 1
		if probe[probeIdx] == cycleNum {
			samples[cycleNum] = registerX
			probeIdx++
		}
		prog[progIdx].PostCycleAction(&registerX)
	}

	fmt.Printf("Samples were:\n")
	totalSignal := 0
	for cycle, x := range samples {
		fmt.Printf("  @%d = %d\n", cycle, x)
		totalSignal += cycle * x
	}

	fmt.Printf("Total signal strength: %d\n", totalSignal)
}
