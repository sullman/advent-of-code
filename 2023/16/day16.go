package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Up byte    = 0b0001
	Right byte = 0b0010
	Down byte  = 0b0100
	Left byte  = 0b1000
)

type Beam struct {
	row int
	col int
	dir byte
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]byte, 0)
	visited := make([][]byte, 0)
	numEnergized := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { break }
		grid = append(grid, []byte(line))
		visited = append(visited, make([]byte, len(line)))
	}

	beams := make([]Beam, 1, 16)
	beams[0] = Beam{0, 0, Right}

	for len(beams) > 0 {
		beam := &beams[0]

		outOfBounds := beam.row < 0 || beam.row >= len(grid) || beam.col < 0 || beam.col >= len(grid[0])

		if outOfBounds || visited[beam.row][beam.col] & beam.dir != 0 {
			beams = beams[1:]
			continue
		}

		if visited[beam.row][beam.col] == 0 {
			numEnergized++
		}

		visited[beam.row][beam.col] |= beam.dir

		switch beam.dir {
		case Up:
			switch grid[beam.row][beam.col] {
			case '.', '|':
				beam.row -= 1
			case '/':
				beam.col, beam.dir = beam.col + 1, Right
			case '\\':
				beam.col, beam.dir = beam.col - 1, Left
			case '-':
				beam.col, beam.dir = beam.col + 1, Right
				beams = append(beams, Beam{beam.row, beam.col - 1, Left})
			}
		case Right:
			switch grid[beam.row][beam.col] {
			case '.', '-':
				beam.col += 1
			case '/':
				beam.row, beam.dir = beam.row - 1, Up
			case '\\':
				beam.row, beam.dir = beam.row + 1, Down
			case '|':
				beam.row, beam.dir = beam.row - 1, Up
				beams = append(beams, Beam{beam.row + 1, beam.col, Down})
			}
		case Down:
			switch grid[beam.row][beam.col] {
			case '.', '|':
				beam.row += 1
			case '/':
				beam.col, beam.dir = beam.col - 1, Left
			case '\\':
				beam.col, beam.dir = beam.col + 1, Right
			case '-':
				beam.col, beam.dir = beam.col + 1, Right
				beams = append(beams, Beam{beam.row, beam.col - 1, Left})
			}
		case Left:
			switch grid[beam.row][beam.col] {
			case '.', '-':
				beam.col -= 1
			case '/':
				beam.row, beam.dir = beam.row + 1, Down
			case '\\':
				beam.row, beam.dir = beam.row - 1, Up
			case '|':
				beam.row, beam.dir = beam.row - 1, Up
				beams = append(beams, Beam{beam.row + 1, beam.col, Down})
			}
		}
	}

	fmt.Printf("Part 1: %d\n", numEnergized)
}
