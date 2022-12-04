package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	completeOverlaps := 0
	for scanner.Scan() {
		assignPair := scanner.Text()

		assigns := strings.Split(assignPair, ",")

		assignRange := strings.Split(assigns[0], "-")
		startA, err := strconv.ParseInt(assignRange[0], 10, 32)
		if err != nil {
			panic(err)
		}
		endA, err := strconv.ParseInt(assignRange[1], 10, 32)
		if err != nil {
			panic(err)
		}

		assignRange = strings.Split(assigns[1], "-")
		startB, err := strconv.ParseInt(assignRange[0], 10, 32)
		if err != nil {
			panic(err)
		}
		endB, err := strconv.ParseInt(assignRange[1], 10, 32)
		if err != nil {
			panic(err)
		}

		if (startA <= startB && endA >= endB) ||
			(startB <= startA && endB >= endA) {
			completeOverlaps++
		}
	}
	fmt.Printf("Total complete overlaps = %d\n", completeOverlaps)
}
