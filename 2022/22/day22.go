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

var Directions = []Movement {
	Movement{0, 1},
	Movement{1, 0},
	Movement{0, -1},
	Movement{-1, 0},
}

const (
	Open = '.'
	Void = ' '
	Wall = '#'
)

func GetSpaceType(grid *[]string, row int, col int) byte {
	if row < 0 || row >= len(*grid) {
		return Void
	}
	if col < 0 || col >= len((*grid)[row]) {
		return Void
	}

	return (*grid)[row][col]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		grid = append(grid, line)
	}

	scanner.Scan()
	instructions := scanner.Text()

	row := 0
	col := 0
	facing := 0

	for ; grid[row][col] != Open; col++ {}

	for i := 0; i < len(instructions); i++ {
		numSpaces := 0

		for ; i < len(instructions) && instructions[i] <= '9'; i++ {
			numSpaces = numSpaces * 10 + int(instructions[i] - '0')
		}

		for n := 0; n < numSpaces; n++ {
			newRow := row + Directions[facing].row
			newCol := col + Directions[facing].col
			switch GetSpaceType(&grid, newRow, newCol) {
			case Open:
				row, col = newRow, newCol
			case Wall:
				break
			case Void:
				wrappedRow, wrappedCol := row, col
				reverse := Directions[(facing + 2) % 4]
				for ; GetSpaceType(&grid, wrappedRow, wrappedCol) != Void; wrappedRow, wrappedCol = wrappedRow + reverse.row, wrappedCol + reverse.col {}
				wrappedRow += Directions[facing].row
				wrappedCol += Directions[facing].col
				if GetSpaceType(&grid, wrappedRow, wrappedCol) == Open {
					row, col = wrappedRow, wrappedCol
				} else {
					break
				}
			default:
				panic("Unexpected tile!")
			}
		}

		if i < len(instructions) {
			switch instructions[i] {
			case 'L':
				facing = (facing + 3) % 4
			case 'R':
				facing = (facing + 1) % 4
			default:
				panic("Unexpected instruction!")
			}
		}
	}

	row++
	col++
	fmt.Printf("Finished at %d, %d facing %d = %d\n", row, col, facing, 1000 * row + 4 * col + facing)
}
