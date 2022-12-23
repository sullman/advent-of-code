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

type Elf struct {
	row int
	col int
	pos string
	proposed Movement
}

func NewElf(row int, col int) *Elf {
	return &Elf{
		row: row,
		col: col,
		pos: fmt.Sprintf("%d,%d", row, col),
	}
}

var Checks = [][]Movement {
	{Movement{-1, -1}, Movement{-1, 0}, Movement{-1, 1}},
	{Movement{1, -1}, Movement{1, 0}, Movement{1, 1}},
	{Movement{-1, -1}, Movement{0, -1}, Movement{1, -1}},
	{Movement{-1, 1}, Movement{0, 1}, Movement{1, 1}},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	elfMap := make(map[string]*Elf)
	elfList := make([]*Elf, 0)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		for col, char := range line {
			if char == '#' {
				elf := NewElf(row, col)
				elfMap[elf.pos] = elf
				elfList = append(elfList, elf)
			}
		}
	}

	for round := 0; ; round++ {
		proposed := make(map[string]int)

		minRow, maxRow := 1000000, -1000000
		minCol, maxCol := 1000000, -1000000

		for _, elf := range elfList {
			if elf.row < minRow { minRow = elf.row }
			if elf.row > maxRow { maxRow = elf.row }
			if elf.col < minCol { minCol = elf.col }
			if elf.col > maxCol { maxCol = elf.col }

			var proposal *Movement
			allClear := true
			for i := round; i < round + len(Checks); i++ {
				checks := Checks[i % len(Checks)]
				clear := true
				for _, move := range checks {
					row, col := elf.row + move.row, elf.col + move.col
					pos := fmt.Sprintf("%d,%d", row, col)
					if _, exists := elfMap[pos]; exists {
						clear = false
						allClear = false
						break
					}
				}
				if clear && proposal == nil {
					proposal = &checks[1]
				}
			}
			if !allClear && proposal != nil {
				row, col := elf.row + proposal.row, elf.col + proposal.col
				pos := fmt.Sprintf("%d,%d", row, col)
				proposed[pos] = proposed[pos] + 1
				elf.proposed = *proposal
			} else {
				elf.proposed = Movement{0,0}
			}
		}

		if round == 10 {
			fmt.Printf("In round %d, bounding rectangle is %dx%d\n", round, maxRow - minRow + 1, maxCol - minCol + 1)
			fmt.Println((maxRow - minRow + 1) * (maxCol - minCol + 1) - len(elfMap))
		}

		if len(proposed) == 0 {
			fmt.Printf("In round %d, everything is stable!\n", round + 1)
			break
		}

		for _, elf := range elfList {
			row, col := elf.row + elf.proposed.row, elf.col + elf.proposed.col
			pos := fmt.Sprintf("%d,%d", row, col)
			if proposed[pos] == 1 {
				delete(elfMap, elf.pos)
				elf.row, elf.col = row, col
				elf.pos = pos
				elfMap[pos] = elf
			}
		}
	}
}
