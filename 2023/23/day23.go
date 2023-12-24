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

type Path struct {
	dest *Node
	length int
}

type Node struct {
	row int
	col int
	visited bool
	finish bool
	paths [4]*Path
}

func Encode(row, col int) uint16 {
	return (uint16(row) << 8) | uint16(col)
}

func FindLongestPath(node *Node) int {
	if node.finish { return 0 }

	node.visited = true
	longest := -10000

	for _, path := range node.paths {
		if path != nil && !path.dest.visited {
			length := path.length + FindLongestPath(path.dest)
			if length > longest { longest = length }
		}
	}

	node.visited = false

	return longest
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
	grid = make([]uint8, len(grid))
	copy(grid, candidates[0].grid)

	for len(candidates) != 0 {
		state := candidates[0]
		moved := false
		index := state.row * numCols + state.col

		if state.grid[index] == Finish {
			if state.length > longest {
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

	// Part 2: Collapse the grid down to a proper graph
	nodes := make(map[uint16]*Node)
	start := new(Node)
	start.col = firstCol
	nodes[Encode(0, firstCol)] = start
	queue := make([]*Node, 1, 128)
	queue[0] = start

	travel := func(node *Node, initialDir int) *Node {
		move := Movements[initialDir]
		row, col := node.row + move.row, node.col + move.col

		if row < 0 || row >= numRows || col < 0 || col >= numCols { return nil }
		if grid[row * numCols + col] & Blocked != 0 { return nil }
		ignoreDir := (initialDir + 2) % 4
		length := 0

		for {
			var newDir int
			numPaths := 0
			length++

			for dir, move := range Movements {
				if dir == ignoreDir { continue }
				newRow, newCol := row + move.row, col + move.col
				if newRow < 0 || newRow >= numRows || newCol < 0 || newCol >= numCols { continue }
				if grid[newRow * numCols + newCol] & Blocked == 0 {
					numPaths++
					newDir = dir
				}
			}

			if numPaths == 1 {
				ignoreDir = (newDir + 2) % 4
				row += Movements[newDir].row
				col += Movements[newDir].col
			} else if numPaths == 0 {
				if grid[row * numCols + col] & Finish == Finish {
					break
				} else {
					return nil
				}
			} else {
				break
			}
		}

		index := Encode(row, col)
		dest := nodes[index]
		if dest == nil {
			dest = new(Node)
			dest.row, dest.col = row, col
			nodes[index] = dest
		}

		// fmt.Printf("Creating path from (%d,%d) -> (%d,%d) length=%d\n", node.row, node.col, dest.row, dest.col, length)

		node.paths[initialDir] = &Path{dest, length}
		dest.paths[ignoreDir] = &Path{node, length}

		return dest
	}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		for i := 0; i < len(Movements); i++ {
			if node.paths[i] != nil { continue }
			dest := travel(node, i)
			if dest != nil {
				queue = append(queue, dest)
			}
		}
	}

	fmt.Printf("Collapsed grid into %d nodes\n", len(nodes))

	nodes[Encode(numRows - 1, lastCol)].finish = true

	longest = FindLongestPath(start)

	fmt.Printf("Part 2: %d\n", longest)
}
