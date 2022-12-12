package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	packetMarker := FindMarker(4, string(data))
	messageMarker := FindMarker(14, string(data))

	fmt.Printf("packet starts at character %d\n", packetMarker)
	fmt.Printf("message starts at character %d\n", messageMarker)
}

// FindMarker locates the first position (1-based index) in `data` where
//   none of the previous `length` number of characters are the same.
func FindMarker(length int, data string) int {
	found := false
	for idx := length; idx < len(data); idx++ {
		chunk := data[idx-length : idx]
		found = true
		for _, letter := range chunk {
			if strings.Count(chunk, string(letter)) > 1 {
				found = false
			}
		}
		if found == true {
			return idx
		}
	}
	return -1
}
