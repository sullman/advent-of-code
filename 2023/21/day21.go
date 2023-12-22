package main

import (
	"bufio"
	"fmt"
	"os"
)

const TotalSteps = 26501365

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

type Distances struct {
	reachable []int
	totalReachable [2]int
}

func CalculateDistances(garden map[uint16]bool, startRow int, startCol int) *Distances {
	distances := new(Distances)
	distances.reachable = make([]int, 1, 256)

	positions := make(map[uint16]bool)
	positions[Encode(startRow, startCol)] = true
	reached := make(map[uint16]bool)
	reached[Encode(startRow, startCol)] = true

	distances.reachable[0] = 1
	distances.totalReachable[0] = 1
	prev := 0

	for steps := 1; len(positions) != 0; steps++ {
		next := make(map[uint16]bool)

		for enc, _ := range positions {
			r, c := Decode(enc)
			for _, move := range Movements {
				newR, newC := r + move.row, c + move.col
				enc = Encode(newR, newC)
				if garden[enc] && !reached[enc] {
					reached[enc] = true
					next[enc] = true
				}
			}
		}

		distances.reachable = append(distances.reachable, prev + len(next))
		distances.totalReachable[steps % 2] += len(next)
		prev = distances.reachable[steps - 1]
		positions = next
	}

	return distances
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	garden := make(map[uint16]bool)
	positions := make(map[uint16]bool)
	row := 0
	var startRow, startCol int
	var numRows, numCols int

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 { numRows++ }
		if row == 0 { numCols = len(line) }
		row++
		for col, ch := range line {
			switch ch {
			case '.':
				garden[Encode(row, col + 1)] = true
			case 'S':
				startRow, startCol = row, col + 1
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
	}

	fmt.Printf("Part 1: %d\n", len(positions))

	// This solution relies on the fact that there are no blocks in the middle or
	// along the edges.
	total := 0
	distances := CalculateDistances(garden, startRow, startCol)
	if TotalSteps >= len(distances.reachable) {
		total += distances.totalReachable[TotalSteps % 2]
	} else {
		total += distances.reachable[TotalSteps]
	}

	edges := make([]*Distances, 4)
	edges[0] = CalculateDistances(garden, startRow, 1)
	edges[1] = CalculateDistances(garden, startRow, numCols)
	edges[2] = CalculateDistances(garden, 1, startCol)
	edges[3] = CalculateDistances(garden, numRows, startCol)

	for steps := numRows - startRow + 1; steps <= TotalSteps; steps += numRows {
		for _, distances = range edges {
			if TotalSteps - steps >= len(distances.reachable) {
				total += distances.totalReachable[(TotalSteps - steps) % 2]
			} else {
				total += distances.reachable[TotalSteps - steps]
			}
		}
	}

	corners := make([]*Distances, 4)
	corners[0] = CalculateDistances(garden, 1, 1)
	corners[1] = CalculateDistances(garden, 1, numCols)
	corners[2] = CalculateDistances(garden, numRows, numCols)
	corners[3] = CalculateDistances(garden, numRows, 1)

	for mult, steps := 1, numRows + 1; steps <= TotalSteps; mult, steps = mult + 1, steps + numRows {
		for _, distances = range corners {
			if TotalSteps - steps >= len(distances.reachable) {
				total += mult * distances.totalReachable[(TotalSteps - steps) % 2]
			} else {
				total += mult * distances.reachable[TotalSteps - steps]
			}
		}
	}

	fmt.Printf("Part 2: %d\n", total)
}
