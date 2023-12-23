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

const (
	Up = uint8(0)
	Right = uint8(1)
	Down = uint8(2)
	Left = uint8(3)
	Visited = uint8(0b00000100)
	Blocked = uint8(0b00001100)
	Open    = uint8(0b00010000)
	Unknown = uint8(0b00100000)
	Finish  = uint8(0b01010000)
)

var Movements = []Movement {
	Movement{-1, 0},
	Movement{0, 1},
	Movement{1, 0},
	Movement{0, -1},
}

type State struct {
	grid []uint8
	row int
	col int
	length int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([]uint8, 0)
	numRows, numCols := 0, 0
	var firstCol, lastCol int
	var empty []uint8

	for scanner.Scan() {
		line := scanner.Text()
		if numCols == 0 {
			numCols = len(line)
			empty = make([]uint8, numCols)
		} else if numCols != len(line) {
			break
		}

		grid = append(grid, empty...)

		for col, ch := range line {
			index := numRows * numCols + col
			switch ch {
			case '#':
				grid[index] = Blocked
			case '.':
				grid[index] = Open
				if firstCol == 0 { firstCol = col }
				lastCol = col
			case '>':
				grid[index] = Right
			case '<':
				grid[index] = Left
			case 'v':
				grid[index] = Down
			case '^':
				grid[index] = Up
			}
		}

		numRows++
	}

	grid[(numRows - 1) * numCols + lastCol] |= Finish

	longest := 0
	candidates := make([]*State, 1, 32)
	candidates[0] = &State{grid, 0, firstCol, 0}

	for len(candidates) != 0 {
		state := candidates[0]
		moved := false
		index := state.row * numCols + state.col

		if state.grid[index] == Finish {
			if state.length > longest {
				fmt.Printf("Found new longest path %d\n", state.length)
				longest = state.length
			}
		} else if state.grid[index] < uint8(len(Movements)) {
			move := Movements[state.grid[index]]
			state.grid[index] |= Visited
			state.row += move.row
			state.col += move.col
			state.length++
			moved = state.grid[state.row * numCols + state.col] & Visited == 0
		} else {
			row, col := state.row, state.col
			state.grid[index] |= Visited
			state.length++
			for _, move := range Movements {
				newRow, newCol := row + move.row, col + move.col
				if newRow < 0 || newRow >= numRows || newCol < 0 || newCol >= numCols { continue }
				if state.grid[newRow * numCols + newCol] & Visited == 0 {
					if moved {
						clone := make([]uint8, len(state.grid))
						copy(clone, state.grid)
						candidates = append(candidates, &State{clone, newRow, newCol, state.length})
					} else {
						moved = true
						state.row, state.col = newRow, newCol
					}
				}
			}
		}

		if !moved { candidates = candidates[1:] }
	}

	fmt.Printf("Part 1: %d\n", longest)
}
