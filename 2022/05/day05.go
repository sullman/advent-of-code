package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s []byte) []byte {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var allStacks [][]byte
	initialized := false
	numStacks := 0
	part2 := true

	for scanner.Scan() {
		line := scanner.Text()

		if initialized {
			var numToMove, from, to int
			n, _ := fmt.Sscanf(line, "move %d from %d to %d", &numToMove, &from, &to)
			if n == 3 {
				from--
				to--
				if part2 {
					allStacks[to] = append(allStacks[to], allStacks[from][len(allStacks[from]) - numToMove:]...)
				} else {
					allStacks[to] = append(allStacks[to], reverse(allStacks[from][len(allStacks[from]) - numToMove:])...)
				}
				allStacks[from] = allStacks[from][:len(allStacks[from]) - numToMove]
			}
		} else if line[1] == '1' {
			initialized = true
			for i := 0; i < numStacks; i++ {
				reverse(allStacks[i])
			}
		} else {
			if numStacks == 0 {
				numStacks = (len(line) + 1) / 4
				allStacks = make([][]byte, numStacks)
				for i := 0; i < numStacks; i++ {
					allStacks[i] = make([]byte, 0, 25)
				}
			}

			for i := 0; i < numStacks; i++ {
				if line[4 * i + 1] != ' ' {
					allStacks[i] = append(allStacks[i], line[4 * i + 1])
				}
			}
		}
	}

	result := ""

	for i := 0; i < numStacks; i++ {
		result = result + string(allStacks[i][len(allStacks[i]) - 1])
	}

	fmt.Println(result)
}
