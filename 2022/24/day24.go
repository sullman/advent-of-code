package main

import (
	"container/list"
	"fmt"
)

type Blizzard struct {
	row int
	col int
	rowMotion int
	colMotion int
}

type State struct {
	row int
	col int
	minutes int
	bestPossible int
	blizzards []Blizzard
}

type Movement struct {
	row int
	col int
}

var Movements = []Movement {
	Movement{0, 0},
	Movement{0, 1},
	Movement{1, 0},
	Movement{0, -1},
	Movement{-1, 0},
}

func Distance(row1 int, col1 int, row2 int, col2 int) int {
	rows := row2 - row1
	cols := col2 - col1
	if rows < 0 { rows *= -1 }
	if cols < 0 { cols *= -1 }
	return rows + cols
}

func InsertSorted(l *list.List, state *State) {
	for e := l.Front(); e != nil; e = e.Next() {
		if state.bestPossible <= e.Value.(*State).bestPossible {
			l.InsertBefore(state, e)
			return
		}
	}

	l.PushBack(state)
}

func main() {
	numRows, numCols := 0, 0
	row, col := 0, 0
	var startCol, endCol int
	var lastOpenCol int
	var char byte

	initial := &State{}
	initial.blizzards = make([]Blizzard, 0)

	for {
		n, _ := fmt.Scanf("%c", &char)
		if n == 0 { break }

		if char == '\n' {
			row++
			numCols = col
			col = 0
			continue
		} else if row == 0 && char == '.' {
			startCol = col
		} else if char == '.' {
			lastOpenCol = col
		} else if char == '>' {
			initial.blizzards = append(initial.blizzards, Blizzard{row, col, 0, 1})
		} else if char == 'v' {
			initial.blizzards = append(initial.blizzards, Blizzard{row, col, 1, 0})
		} else if char == '<' {
			initial.blizzards = append(initial.blizzards, Blizzard{row, col, 0, -1})
		} else if char == '^' {
			initial.blizzards = append(initial.blizzards, Blizzard{row, col, -1, 0})
		}

		col++
	}

	numRows = row
	endCol = lastOpenCol
	maxRow, maxCol := numRows - 1, numCols - 1
	initial.col = startCol
	initial.bestPossible = Distance(initial.row, initial.col, maxRow, endCol)

	fmt.Printf("Read %d rows, %d cols, %d blizzards, starting at 0,%d and ending at %d,%d\n", numRows, numCols, len(initial.blizzards), startCol, maxRow, endCol)

	states := list.New()
	states.PushFront(initial)
	visited := make(map[string]bool)

	for states.Len() > 0 {
		elem := states.Front()
		state := states.Remove(elem).(*State)
		// fmt.Printf("Checking %v\n", state)

		if state.row == maxRow {
			fmt.Printf("Made it out in %d minutes\n", state.minutes)
			break
		}

		newBlizzards := make([]Blizzard, len(state.blizzards))
		blocked := make(map[string]bool)

		for i := 0; i < len(state.blizzards); i++ {
			newRow := state.blizzards[i].row + state.blizzards[i].rowMotion
			newCol := state.blizzards[i].col + state.blizzards[i].colMotion

			if newRow == 0 { newRow = maxRow - 1 }
			if newRow == maxRow { newRow = 1 }
			if newCol == 0 { newCol = maxCol - 1 }
			if newCol == maxCol { newCol = 1 }

			newBlizzards[i] = state.blizzards[i]
			newBlizzards[i].row = newRow
			newBlizzards[i].col = newCol

			blocked[fmt.Sprintf("%d,%d", newRow, newCol)] = true
		}

		for _, move := range Movements {
			newRow := state.row + move.row
			newCol := state.col + move.col

			// fmt.Printf("Checking %d,%d blocked=%t\n", newRow, newCol, blocked[fmt.Sprintf("%d,%d", newRow, newCol)])

			if blocked[fmt.Sprintf("%d,%d", newRow, newCol)] { continue }

			memoKey := fmt.Sprintf("%d,%d,%d", state.minutes + 1, newRow, newCol)
			if visited[memoKey] { continue }

			if (newRow > 0 && newRow < maxRow && newCol > 0 && newCol < maxCol) || (newRow == 0 && newCol == startCol) || (newRow == maxRow && newCol == endCol) {
				visited[memoKey] = true
				s := &State{
					row: newRow,
					col: newCol,
					minutes: state.minutes + 1,
					bestPossible: state.minutes + 1 + Distance(newRow, newCol, maxRow, endCol),
					blizzards: newBlizzards,
				}
				InsertSorted(states, s)
			}
		}
	}
}
