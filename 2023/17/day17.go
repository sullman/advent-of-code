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

func ShortestPath(grid [][]uint8, minConsecutive, maxConsecutive uint8) uint {
	visited := make([][][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([][]bool, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			visited[i][j] = make([]bool, len(Movements) << 6)
		}
	}

	var state *State
	candidates := list.New()
	candidates.PushFront(&State{0, 0, 0, 0, 0})
	candidates.PushFront(&State{0, 0, 0, 0, 3})
	// Push a sentinel so we can always InsertBefore
	candidates.PushBack(&State{0, 0, 0, ^uint(0), 1})

	for {
		elem := candidates.Front()
		candidates.Remove(elem)
		state = elem.Value.(*State)

		if state.row == len(grid) - 1 && state.col == len(grid[0]) - 1 { break }

		for i, delta := range Movements {
			var newState *State

			if i == state.dir && state.straightRemaining > 0 {
				newRow := state.row + delta.row
				newCol := state.col + delta.col
				if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) { continue }
				newCost := state.cost + uint(grid[newRow][newCol])
				newState = &State{newRow, newCol, state.straightRemaining - 1, newCost, i}
			} else if i % 2 != state.dir % 2 {
				newState = &State{state.row, state.col, maxConsecutive, state.cost, i}
				newRow := state.row
				newCol := state.col
				for j := uint8(0); j < minConsecutive; j++ {
					newRow += delta.row
					newCol += delta.col
					if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
						newState = nil
						break
					}
					newState.row, newState.col = newRow, newCol
					newState.cost += uint(grid[newRow][newCol])
					newState.straightRemaining -= 1
				}
			}

			if newState == nil { continue }

			key := int(newState.straightRemaining << 2) | i

			if visited[newState.row][newState.col][key] {
				continue
			} else {
				visited[newState.row][newState.col][key] = true
			}

			for e := candidates.Front(); e != nil; e = e.Next() {
				if newState.cost < e.Value.(*State).cost {
					candidates.InsertBefore(newState, e)
					break
				}
			}
		}
	}

	return state.cost
}

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

	fmt.Printf("Part 1: %d\n", ShortestPath(grid, 1, 3))
	fmt.Printf("Part 2: %d\n", ShortestPath(grid, 4, 10))
}
