package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func runeInBoth(one, other string) []rune {
	var inBoth []rune
	for _, oneChar := range one {
		for _, otherChar := range other {
			if oneChar == otherChar {
				inBoth = append(inBoth, oneChar)
			}
		}
	}
	return inBoth
}

func priority(char rune) int {
	if unicode.IsUpper(char) {
		return int(char-'A') + 27
	}
	if unicode.IsLower(char) {
		return int(char-'a') + 1
	}
	log.Panicf("unrecognized rune: %c (%d)", char, char)
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		comp1 := rucksack[:len(rucksack)/2]
		comp2 := rucksack[len(rucksack)/2:]

		common := runeInBoth(comp1, comp2)
		pri := priority(common[0])
		total += pri

		// fmt.Printf("Rucksack:\n  left  = %s\n  right = %s\n  common = %c (%d)\n", comp1, comp2, common, pri)
	}
	fmt.Printf("Total comparment priorities: %d\n", total)
}
