package main

import (
	"bufio"
	"fmt"
	"os"
)

const TotalCycles int = 1000000000 - 1

func Score(rows [][]rune) int {
	sum := 0

	for r, row := range rows {
		weight := len(rows) - r
		for _, ch := range row {
			if ch == 'O' {
				sum += weight
			}
		}
	}

	return sum
}

func Vertical(rows [][]rune, dir int) {
	var from, to, step int
	if dir > 0 {
		from = -1
		to = len(rows) - 1
		step = 1
	} else {
		from = len(rows)
		to = 0
		step = -1
	}

	for r := to; r != from; r -= step {
		row := rows[r]
		for c, ch := range row {
			if ch == 'O' {
				for r2 := r + step; r2 != to + step; r2 += step {
					if rows[r2][c] == '.' {
						rows[r2 - step][c], rows[r2][c] = '.', 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func Horizontal(rows [][]rune, dir int) {
	var from, to, step int
	if dir > 0 {
		from = -1
		to = len(rows[0]) - 1
		step = 1
	} else {
		from = len(rows[0])
		to = 0
		step = -1
	}

	for _, row := range rows {
		for c := to; c != from; c -= step {
			if row[c] == 'O' {
				for c2 := c + step; c2 != to + step; c2 += step {
					if row[c2] == '.' {
						row[c2 - step], row[c2] = '.', 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rows := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { break }
		rows = append(rows, []rune(line))
	}

	// Do the first round separately in order to print part 1
	Vertical(rows, -1)
	fmt.Printf("Part 1: %d\n", Score(rows))
	Horizontal(rows, -1)
	Vertical(rows, 1)
	Horizontal(rows, 1)

	scores := make([]int, 1, 1000)
	scores[0] = Score(rows)

	// Serializing the rows to act as a key to be confident we're in the exact
	// same position seems like it would require large keys. Try just looking at
	// patterns in the scores instead, roughly based on Brent's algorithm. Avoid
	// any misleading short cycles at the very beginning.
	cycleLength := 0
	pivot, nextPivot := 7, 15

	for i := 1; i <= TotalCycles; i++ {
		Vertical(rows, -1)
		Horizontal(rows, -1)
		Vertical(rows, 1)
		Horizontal(rows, 1)
		scores = append(scores, Score(rows))

		if i <= pivot { continue }
		if i == nextPivot {
			pivot = nextPivot
			nextPivot = (nextPivot << 1) + 1
		}

		if scores[i] != scores[i - cycleLength] {
			cycleLength = 0
		} else if cycleLength != 0 && i == pivot + 2 * cycleLength {
			break
		}

		if cycleLength == 0 && scores[pivot] == scores[i] {
			cycleLength = i - pivot
		}
	}

	prev := TotalCycles - ((1 + (TotalCycles - len(scores) + 1) / cycleLength) * cycleLength)
	fmt.Printf("Part 2: %d\n", scores[prev])
}
