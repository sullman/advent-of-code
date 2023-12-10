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

	row, col := 0, 0
	edge := North
	flowNorth := func() {
		row--
		length++
		edge = South
	}
	flowSouth := func() {
		row++
		length++
		edge = North
	}
	flowEast := func() {
		col++
		length++
		edge = West
	}
	flowWest := func() {
		col--
		length++
		edge = East
	}

	done:
	for dir := North; ; dir++ {
		length = 0
		row, col = startRow, startCol

		switch dir {
		case North:
			flowNorth()
		case East:
			flowEast()
		case South:
			flowSouth()
		case West:
			flowWest()
		}

		for {
			if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) { break }

			tile := lines[row][col]
			if tile == 'S' {
				break done
			} else if tile == '|' && edge == South {
				flowNorth()
			} else if tile == '|' && edge == North {
				flowSouth()
			} else if tile == '-' && edge == West {
				flowEast()
			} else if tile == '-' && edge == East {
				flowWest()
			} else if tile == 'L' && edge == North {
				flowEast()
			} else if tile == 'L' && edge == East {
				flowNorth()
			} else if tile == 'J' && edge == North {
				flowWest()
			} else if tile == 'J' && edge == West {
				flowNorth()
			} else if tile == '7' && edge == South {
				flowWest()
			} else if tile == '7' && edge == West {
				flowSouth()
			} else if tile == 'F' && edge == South {
				flowEast()
			} else if tile == 'F' && edge == East {
				flowSouth()
			} else {
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", (length + 1) / 2)
}
