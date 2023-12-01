package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		found := false
		first := 0
		last := 0

		for _, r := range line {
			if r >= '0' && r <= '9' {
				last = int(r - '0')
				if !found {
					found = true
					first = last
				}
			}
		}

		if found {
			sum += 10 * first + last
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
}
