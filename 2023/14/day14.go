package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rows := make([][]rune, 0)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { break }
		row := len(rows)
		rows = append(rows, []rune(line))

		for c, ch := range rows[row] {
			if ch == 'O' {
				for r := row - 1; r >= 0; r-- {
					if rows[r][c] == '.' {
						rows[r + 1][c], rows[r][c] = '.', 'O'
					} else {
						break
					}
				}
			}
		}
	}

	for r, row := range rows {
		weight := len(rows) - r
		for _, ch := range row {
			if ch == 'O' {
				sum += weight
			}
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
}
