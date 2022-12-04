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
	totalCompPri := 0
	totalBadgePri := 0
	var group []string
	for scanner.Scan() {
		rucksack := scanner.Text()
		comp1 := rucksack[:len(rucksack)/2]
		comp2 := rucksack[len(rucksack)/2:]

		common := runeInBoth(comp1, comp2)
		pri := priority(common[0])
		totalCompPri += pri

		group = append(group, rucksack)
		if len(group) == 3 {
			common := runeInBoth(group[0], group[1])
			common = runeInBoth(string(common), group[2])

			pri := priority(common[0])
			totalBadgePri += pri

			group = nil
		}

		// fmt.Printf("Rucksack:\n  left  = %s\n  right = %s\n  common = %c (%d)\n", comp1, comp2, common, pri)
	}
	fmt.Printf("Total comparment priorities: %d\n", totalCompPri)
	fmt.Printf("Total badge priorities: %d\n", totalBadgePri)
}
