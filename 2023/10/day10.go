package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	North = iota
	East
	South
	West
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 10)
	startRow, startCol := -1, -1
	var length int

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if startRow == -1 {
			startCol = strings.IndexByte(line, 'S')
			if startCol != -1 {
				startRow = len(lines) - 1
			}
		}
	}

	fmt.Printf("Starting at %d,%d\n", startRow, startCol)

	done:
	for dir := North; ; dir++ {
		row, col := startRow, startCol

		switch dir {
		case North:
			row += 1
		case East:
			col -= 1
		case South:
			row -= 1
		case West:
			col += 1
		}

		for length = 1; ; length++ {
			if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) { break }

			tile := lines[row][col]
			if tile == 'S' {
				break done
			} else if tile == '|' && dir == South {
				row -= 1
			} else if tile == '|' && dir == North {
				row += 1
			} else if tile == '-' && dir == West {
				col += 1
			} else if tile == '-' && dir == East {
				col -= 1
			} else if tile == 'L' && dir == North {
				dir = West
				col += 1
			} else if tile == 'L' && dir == East {
				dir = South
				row -= 1
			} else if tile == 'J' && dir == North {
				dir = East
				col -= 1
			} else if tile == 'J' && dir == West {
				dir = South
				row -= 1
			} else if tile == '7' && dir == South {
				dir = East
				col -= 1
			} else if tile == '7' && dir == West {
				dir = North
				row += 1
			} else if tile == 'F' && dir == South {
				dir = West
				col += 1
			} else if tile == 'F' && dir == East {
				dir = North
				row += 1
			} else {
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", (length + 1) / 2)
}
