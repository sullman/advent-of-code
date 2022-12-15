package main

import (
	"fmt"
)

type Space byte

const (
	Air Space = iota
	Rock
	Sand
)

const StartX = 500
const StartY = 0

func main() {
	grid := make(map[string]Space)
	maxY := 0

	for {
		var x, y, nextX, nextY int
		var sep rune
		n, _ := fmt.Scanf("%d,%d%c", &x, &y, &sep)
		if n != 3 { break }
		grid[fmt.Sprintf("%d,%d", x, y)] = Rock
		if y > maxY { maxY = y }

		for sep == ' ' {
			fmt.Scanf("-> %d,%d%c", &nextX, &nextY, &sep)
			if nextY > maxY { maxY = nextY }
			// fmt.Printf("Filling from %d,%d -> %d,%d\n", x, y, nextX, nextY)
			var dx, dy int
			if nextX == x && nextY > y {
				dy = 1
			} else if nextX == x {
				dy = -1
			} else if nextX > x {
				dx = 1
			} else {
				dx = -1
			}

			for ; x != nextX || y != nextY; {
				x, y = x + dx, y + dy
				// fmt.Printf("Filling %d,%d\n", x, y)
				grid[fmt.Sprintf("%d,%d", x, y)] = Rock
			}
		}
	}

	fmt.Printf("Maximum Y = %d\n", maxY)
	fmt.Printf("Rocks = %d\n", len(grid))

	numSand := 0
	var x, y int

	for y < maxY {
		x, y = StartX, StartY
		for y < maxY {
			if grid[fmt.Sprintf("%d,%d", x, y + 1)] == Air {
				y++
			} else if grid[fmt.Sprintf("%d,%d", x - 1, y + 1)] == Air {
				x--
				y++
			} else if grid[fmt.Sprintf("%d,%d", x + 1, y + 1)] == Air {
				x++
				y++
			} else {
				grid[fmt.Sprintf("%d,%d", x, y)] = Sand
				numSand++
				break
			}
		}
	}

	fmt.Println(numSand)

	// The beginning will proceed the same as part 1, so we can simply avoid
	// resetting grid and numSand. It wouldn't be that much harder to combine
	// the loops instead of copy/paste, but I'm lazy.
	startStr := fmt.Sprintf("%d,%d", StartX, StartY)

	for grid[startStr] == Air {
		x, y = StartX, StartY
		for {
			if y == maxY + 1 {
				grid[fmt.Sprintf("%d,%d", x, y)] = Sand
				numSand++
				break
			} else if grid[fmt.Sprintf("%d,%d", x, y + 1)] == Air {
				y++
			} else if grid[fmt.Sprintf("%d,%d", x - 1, y + 1)] == Air {
				x--
				y++
			} else if grid[fmt.Sprintf("%d,%d", x + 1, y + 1)] == Air {
				x++
				y++
			} else {
				grid[fmt.Sprintf("%d,%d", x, y)] = Sand
				numSand++
				break
			}
		}
	}

	fmt.Println(numSand)
}
