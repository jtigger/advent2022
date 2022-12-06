package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Crate struct {
	Content rune
}

type Stack struct {
	contents []Crate
}

// Pull takes in a new crate at the bottom of the stack
func (s *Stack) Pull(crate Crate) {
	s.contents = append([]Crate{crate}, s.contents...)
}

func (s *Stack) Push(crate Crate) {
	s.contents = append(s.contents, crate)
}

func (s *Stack) Pop() (Crate, error) {
	if len(s.contents) == 0 {
		return Crate{}, fmt.Errorf("attempted to pop and empty stack")
	}
	crate := s.contents[len(s.contents)-1]
	s.contents = s.contents[:len(s.contents)-1]
	return crate, nil
}

const SpaceBetweenStacks = 4

// Instruction is an action taken on a list of Stack's
type Instruction interface {
	// ApplyOn executes instruction on the supplied list of Stack's
	ApplyOn([]Stack) error
}

type MoveInstr9000 struct {
	Quantity       int
	SourceStackNum int
	SinkStackNum   int
}

func (i MoveInstr9000) ApplyOn(stacks []Stack) error {
	if i.SourceStackNum < 1 || i.SourceStackNum > len(stacks) {
		return fmt.Errorf("invalid SourceStackNum value, %d (expected between 1 and number of stacks, %d)",
			i.SourceStackNum, len(stacks))
	}
	if i.Quantity < 1 || i.Quantity > len(stacks[i.SourceStackNum-1].contents) {
		return fmt.Errorf("invalid Quantity value, %d (expected between 1 and number of crates in source stack, %d)",
			i.Quantity, len(stacks[i.SourceStackNum-1].contents))
	}
	if i.SinkStackNum < 1 || i.SinkStackNum > len(stacks) {
		return fmt.Errorf("invalid SinkStackNum value, %d (expected between 1 and number of stacks, %d)",
			i.SinkStackNum, len(stacks))
	}

	for idx := 1; idx <= i.Quantity; idx++ {
		crate, err := stacks[i.SourceStackNum-1].Pop()
		if err != nil {
			log.Panicf("'move' failed: %s", err)
		}
		stacks[i.SinkStackNum-1].Push(crate)
	}
	return nil
}

type MoveInstr9001 struct {
	Quantity       int
	SourceStackNum int
	SinkStackNum   int
}

func (i MoveInstr9001) ApplyOn(stacks []Stack) error {
	if i.SourceStackNum < 1 || i.SourceStackNum > len(stacks) {
		return fmt.Errorf("invalid SourceStackNum value, %d (expected between 1 and number of stacks, %d)",
			i.SourceStackNum, len(stacks))
	}
	if i.Quantity < 1 || i.Quantity > len(stacks[i.SourceStackNum-1].contents) {
		return fmt.Errorf("invalid Quantity value, %d (expected between 1 and number of crates in source stack, %d)",
			i.Quantity, len(stacks[i.SourceStackNum-1].contents))
	}
	if i.SinkStackNum < 1 || i.SinkStackNum > len(stacks) {
		return fmt.Errorf("invalid SinkStackNum value, %d (expected between 1 and number of stacks, %d)",
			i.SinkStackNum, len(stacks))
	}
	var craneStorage Stack

	for idx := 1; idx <= i.Quantity; idx++ {
		crate, err := stacks[i.SourceStackNum-1].Pop()
		if err != nil {
			log.Panicf("'move' failed: %s", err)
		}
		craneStorage.Push(crate)
	}
	for idx := 1; idx <= i.Quantity; idx++ {
		crate, err := craneStorage.Pop()
		if err != nil {
			log.Panicf("'move' failed: %s", err)
		}
		stacks[i.SinkStackNum-1].Push(crate)
	}
	return nil
}

func parse(line string) (quantity, sourceStackNum, sinkStackNum int, err error) {
	tokens := strings.Split(line, " ")

	if tokens[0] != "move" {
		return 0, 0, 0,
			fmt.Errorf("expected keyword 'move'; was '%s'", tokens[0])
	}
	qty, err := strconv.ParseInt(tokens[1], 10, 16)
	if err != nil {
		return 0, 0, 0,
			fmt.Errorf("expected quantity (as an integer); was: '%s' ", tokens[1])
	}
	if tokens[2] != "from" {
		return 0, 0, 0,
			fmt.Errorf("expected keyword 'from'; was '%s'", tokens[2])
	}
	source, err := strconv.ParseInt(tokens[3], 10, 16)
	if err != nil {
		return 0, 0, 0,
			fmt.Errorf("expected source stack number (as an integer); was: '%s' ", tokens[3])
	}
	if tokens[4] != "to" {
		return 0, 0, 0,
			fmt.Errorf("expected keyword 'to'; was '%s'", tokens[4])
	}
	sink, err := strconv.ParseInt(tokens[5], 10, 16)
	if err != nil {
		return 0, 0, 0,
			fmt.Errorf("expected sink stack number (as an integer); was: '%s' ", tokens[5])
	}
	return int(qty), int(source), int(sink), nil
}

func main() {
	version := flag.String("instrver", "9000", "Version of instruction set to use")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	// pre-allocate the number of stacks to avoid index bound errors
	// note: index into this list is one less than the "stack number"
	stacks := make([]Stack, 16)
	var prog []Instruction

	// either "parse-pictogram" or "load-program"
	mode := "parse-pictogram"
	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()

		switch mode {
		case "parse-pictogram":
			for stackNum, idx := 1, 1; idx < len(line); stackNum, idx = stackNum+1, idx+SpaceBetweenStacks {
				r := rune(line[idx])
				if unicode.IsLetter(r) {
					stacks[stackNum-1].Pull(Crate{r})
				}
			}
			// blank line indicates end of slack pictogram / start of program
			if len(line) == 0 {
				mode = "load-program"
			}
		case "load-program":
			qty, source, sink, err := parse(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to parse program line:\n  %d | %s\n\n  > %s\n", lineNo, line, err)
				os.Exit(1)
			}
			switch *version {
			case "9000":
				prog = append(prog, MoveInstr9000{qty, source, sink})
			case "9001":
				prog = append(prog, MoveInstr9001{qty, source, sink})
			}
		default:
			log.Panicf("invalid mode: %s. (expected one of [parse-pictogram,load-program]", mode)
		}
		lineNo++
	}
	for idx, instr := range prog {
		err := instr.ApplyOn(stacks)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error executing line %d (%#v): %s\n", idx+1, instr, err)
			os.Exit(1)
		}
	}
	for _, stack := range stacks {
		topCrate, _ := stack.Pop()
		fmt.Printf("%c", topCrate.Content)
	}
	fmt.Println()
}
