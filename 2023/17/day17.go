package main

import (
	"bufio"
	"container/list"
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

type State struct {
	row int
	col int
	straightRemaining uint8
	cost uint
	dir int
}

const MaxConsecutive = 3

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]uint8, 0, 128)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { break }
		row := make([]uint8, len(line))
		for i, ch := range line {
			row[i] = uint8(ch - '0')
		}
		grid = append(grid, row)
	}

	visited := make([][][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([][]bool, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			visited[i][j] = make([]bool, len(Movements) << 2)
		}
	}

	var state *State
	candidates := list.New()
	candidates.PushFront(&State{0, 0, MaxConsecutive, 0, 1})
	// Push a sentinel so we can always InsertBefore
	candidates.PushBack(&State{0, 0, 0, ^uint(0), 1})

	for {
		elem := candidates.Front()
		candidates.Remove(elem)
		state = elem.Value.(*State)

		if state.row == len(grid) - 1 && state.col == len(grid[0]) - 1 { break }

		for i, delta := range Movements {
			newRow := state.row + delta.row
			newCol := state.col + delta.col
			if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) { continue }

			newCost := state.cost + uint(grid[newRow][newCol])
			newState := &State{newRow, newCol, MaxConsecutive - 1, newCost, i}

			if i == state.dir {
				if state.straightRemaining > 0 {
					newState.straightRemaining = state.straightRemaining - 1
				} else {
					continue
				}
			} else if i % 2 == state.dir % 2 {
				continue
			}

			key := int(newState.straightRemaining << 2) | i

			if visited[newRow][newCol][key] {
				continue
			} else {
				visited[newRow][newCol][key] = true
			}

			for e := candidates.Front(); e != nil; e = e.Next() {
				if newCost < e.Value.(*State).cost {
					candidates.InsertBefore(newState, e)
					break
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", state.cost)
}
