package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"day09/pkg/rope"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numKnots := flag.Int("knots", 2, "Number of knots in the rope.")
	flag.Parse()

	rope := rope.NewLength(*numKnots)

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
			rope.Up(distance)
		case "D":
			rope.Down(distance)
		case "L":
			rope.Left(distance)
		case "R":
			rope.Right(distance)
		default:
			log.Panicf("Unknown direction %s", tokens[0])
		}
	}
	fmt.Printf("Total positions: %d\n", rope.TailJourneyLength())
}
