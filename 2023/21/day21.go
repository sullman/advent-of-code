package main

import (
	"bufio"
	"fmt"
	"os"
)

type Movement struct {
	row int
	col int
}

var Movements = []Movement {
	Movement{-1, 0},
	Movement{0, 1},
	Movement{1, 0},
	Movement{0, -1},
}

func Encode(row, col int) uint16 {
	return (uint16(row) << 8) | (uint16(col) & 0xff)
}

func Decode(enc uint16) (int, int) {
	return int(enc >> 8), int(enc & 0xff)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	garden := make(map[uint16]bool)
	positions := make(map[uint16]bool)
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		row++
		for col, ch := range line {
			switch ch {
			case '.':
				garden[Encode(row, col + 1)] = true
			case 'S':
				garden[Encode(row, col + 1)] = true
				positions[Encode(row, col + 1)] = true
			}
		}
	}

	for steps := 0; steps < 64; steps++ {
		next := make(map[uint16]bool)
		for enc, _ := range positions {
			r, c := Decode(enc)
			for _, move := range Movements {
				enc = Encode(r + move.row, c + move.col)
				if garden[enc] {
					next[enc] = true
				}
			}
		}
		positions = next
		// fmt.Printf("After %d steps, %d positions are reachable\n", steps + 1, len(positions))
	}

	fmt.Printf("Part 1: %d\n", len(positions))
}
