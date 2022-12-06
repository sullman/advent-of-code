package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	} else if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total1 := 0
	total2 := 0

	var badgeBits [53]uint8
	badge := 0
	row := 0

	for scanner.Scan() {
		var seen [53]bool
		line := scanner.Text()
		divider := len(line) / 2
		dupe := 0
		badgeBit := uint8(1 << (row % 3))

		for i, r := range line {
			p := priority(r)

			if i < divider {
				seen[p] = true
			} else if seen[p] {
				dupe = p
			}

			badgeBits[p] |= badgeBit
			if badgeBits[p] == 0b111 {
				badge = p
			}
		}

		total1 += dupe
		if badge != 0 {
			total2 += badge
			badge = 0
			for i := 0; i < len(badgeBits); i++ {
				badgeBits[i] = 0
			}
		}

		row++
	}

	fmt.Println(total1)
	fmt.Println(total2)
}
