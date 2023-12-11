package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	North = byte(0b0001)
	East  = byte(0b0010)
	South = byte(0b0100)
	West  = byte(0b1000)
)

var TileTypes = map[rune]byte {
	'|': North | South,
	'-': East | West,
	'F': South | East,
	'J': North | West,
	'L': North | East,
	'7': South | West,
	'S': North | East | South | West,
}

func main() {
	var Inverse [15]byte
	Inverse[North] = South
	Inverse[East] = West
	Inverse[South] = North
	Inverse[West] = East

	scanner := bufio.NewScanner(os.Stdin)
	tiles := make([][]byte, 0, 10)
	startRow, startCol := -1, -1

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]byte, len(line))

		for i, ch := range line {
			row[i] = TileTypes[ch]
			if ch == 'S' {
				startRow = len(tiles)
				startCol = i
			}
		}
		tiles = append(tiles, row)
	}

	numRows, numCols := len(tiles), len(tiles[0])
	var length int
	pipes := make([][]byte, len(tiles))

	done:
	for firstDirection := North; ; firstDirection <<= 1 {
		length = 0
		row, col := startRow, startCol
		for i := 0; i < len(tiles); i++ {
			pipes[i] = make([]byte, len(tiles[i]))
		}

		for flow := firstDirection; ; {
			inv := Inverse[flow]
			oldRow, oldCol := row, col

			switch flow {
			case North:
				row--
			case East:
				col++
			case South:
				row++
			case West:
				col--
			}

			if row < 0 || row >= numRows || col < 0 || col >= numCols { break }

			if tiles[row][col] & inv == 0 {
				if inv == 0 {
					break done
				} else {
					break
				}
			}

			length++
			pipes[oldRow][oldCol] |= flow
			pipes[row][col] |= inv
			flow = tiles[row][col] & ^inv
		}
	}

	fmt.Printf("Part 1: %d\n", (length + 1) / 2)

	numInside := 0

	for _, row := range pipes {
		inside := false
		for _, pipe := range row {
			if pipe & North != 0 {
				inside = !inside
			} else if inside && pipe == 0 {
				numInside++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", numInside)
}
