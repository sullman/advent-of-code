package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	value := 0

	for {
		ch, err := reader.ReadByte()
		if err != nil { break }
		if ch == '\n' { continue }

		if ch == ',' {
			sum += value
			value = 0
		} else {
			value = ((value + int(ch)) * 17) & 0xff
		}
	}

	sum += value

	fmt.Printf("Part 1: %d\n", sum)
}
