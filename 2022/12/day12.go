package main

import (
	"container/list"
	"fmt"
)

type Path struct {
	row int
	col int
	steps int
	possible int
}

func distance(x1 int, y1 int, x2 int, y2 int) int {
	dx := x2 - x1
	dy := y2 - y1
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}

	return dx + dy
}

func insertSorted(list *list.List, path *Path) {
	for e := list.Front(); e != nil; e = e.Next() {
		if path.possible <= e.Value.(*Path).possible {
			list.InsertBefore(path, e)
			return
		}
	}

	list.PushBack(path)
}

func main() {
	grid := make([][]rune, 0)
	visited := make([][]bool, 0)
	deltas := [5]int{-1, 0, 1, 0, -1}
	var rows, cols, startRow, startCol, endRow, endCol int

	for {
		var line string
		_, err := fmt.Scanln(&line)

		if err != nil {
			break
		}

		row := len(grid)
		grid = append(grid, make([]rune, len(line)))
		visited = append(visited, make([]bool, len(line)))
		for col, ch := range line {
			switch ch {
			case 'S':
				grid[row][col] = 0
				startRow = row
				startCol = col
			case 'E':
				grid[row][col] = 25
				endRow = row
				endCol = col
			default:
				grid[row][col] = ch - 'a'
			}
		}
	}

	rows = len(grid)
	cols = len(grid[0])

	candidates := list.New()
	initial := &Path{ row: startRow, col: startCol, steps: 0, possible: distance(startRow, startCol, endRow, endCol) }
	candidates.PushFront(initial)

	for {
		elem := candidates.Front()
		path := candidates.Remove(elem).(*Path)

		if path.row == endRow && path.col == endCol {
			fmt.Println(path.steps)
			break
		}

		visited[path.row][path.col] = true

		for i := 0; i < 4; i++ {
			newRow := path.row + deltas[i]
			newCol := path.col + deltas[i+1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && !visited[newRow][newCol] && grid[newRow][newCol] - 1 <= grid[path.row][path.col] {
				newPath := &Path{ row: newRow, col: newCol, steps: path.steps + 1, possible: path.steps + 1 + distance(newRow, newCol, endRow, endCol) }
				insertSorted(candidates, newPath)
			}
		}
	}
}
