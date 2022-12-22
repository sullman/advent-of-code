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
	Right = iota
	Down
	Left
	Up
)

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

func WrapPart1(grid *[]string, row *int, col *int, facing *int) bool {
	wrappedRow, wrappedCol := *row, *col
	reverse := Directions[(*facing + 2) % 4]
	for ; GetSpaceType(grid, wrappedRow, wrappedCol) != Void; wrappedRow, wrappedCol = wrappedRow + reverse.row, wrappedCol + reverse.col {}
	wrappedRow += Directions[*facing].row
	wrappedCol += Directions[*facing].col
	if GetSpaceType(grid, wrappedRow, wrappedCol) == Open {
		*row, *col = wrappedRow, wrappedCol
	} else {
		return false
	}

	return true
}

func WrapPart2(grid *[]string, row *int, col *int, facing *int) bool {
	// Hardcode the transitions :[
	newRow, newCol := *row, *col
	newFacing := *facing
	if *facing == Left && *row < 50 {
		newFacing = Right
		newRow, newCol = 149 - *row, 0
	} else if *facing == Left && *row < 100 {
		newFacing = Down
		newRow, newCol = 100, *row - 50
	} else if *facing == Left && *row < 150 {
		newFacing = Right
		newRow, newCol = 149 - *row, 50
	} else if *facing == Left && *row < 200 {
		newFacing = Down
		newRow, newCol = 0, *row - 100
	} else if *facing == Up && *col < 50 {
		newFacing = Right
		newRow, newCol = *col + 50, 50
	} else if *facing == Up && *col < 100 {
		newFacing = Right
		newRow, newCol = *col + 100, 0
	} else if *facing == Up && *col < 150 {
		newFacing = Up
		newRow, newCol = 199, *col - 100
	} else if *facing == Right && *row < 50 {
		newFacing = Left
		newRow, newCol = 149 - *row, 99
	} else if *facing == Right && *row < 100 {
		newFacing = Up
		newRow, newCol = 49, 50 + *row
	} else if *facing == Right && *row < 150 {
		newFacing = Left
		newRow, newCol = 149 - *row, 149
	} else if *facing == Right && *row < 200 {
		newFacing = Up
		newRow, newCol = 149, *row - 100
	} else if *facing == Down && *col < 50 {
		newFacing = Down
		newRow, newCol = 0, 100 + *col
	} else if *facing == Down && *col < 100 {
		newFacing = Left
		newRow, newCol = *col + 100, 49
	} else if *facing == Down && *col < 150 {
		newFacing = Left
		newRow, newCol = *col - 50, 99
	} else {
		panic("Huh?")
	}

	if (*grid)[newRow][newCol] == Wall {
		return false
	}

	*facing = newFacing
	*row, *col = newRow, newCol
	return true
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

	wrapFunctions := []func(*[]string, *int, *int, *int) bool{WrapPart1, WrapPart2}

	for _, wrap := range wrapFunctions {
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
					if !wrap(&grid, &row, &col, &facing) {
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
}
