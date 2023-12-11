package main

import (
	"bufio"
	"fmt"
	"os"
)

type Galaxy struct {
	X int
	Y int
}

func Dist(x1, y1, x2, y2 int) int {
	dx := x2 - x1
	if dx < 0 { dx *= -1 }
	dy := y2 - y1
	if dy < 0 { dy *= -1 }
	return dx + dy
}

func Expand(lists [][]*Galaxy, factor int, x bool) {
	numToAdd := 0

	for _, list := range lists {
		if len(list) == 0 {
			numToAdd += factor
		} else if numToAdd != 0 {
			for _, galaxy := range list {
				if x {
					galaxy.X += numToAdd
				} else {
					galaxy.Y += numToAdd
				}
			}
		}
	}
}

func TotalDistance(galaxies []*Galaxy) int {
	sum := 0
	for i := 0; i < len(galaxies) - 1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += Dist(galaxies[i].X, galaxies[i].Y, galaxies[j].X, galaxies[j].Y)
		}
	}
	return sum
}

func main() {
	var galaxy *Galaxy
	scanner := bufio.NewScanner(os.Stdin)
	galaxies := make([]*Galaxy, 0)
	rows := make([][]*Galaxy, 0)
	var cols [][]*Galaxy = nil

	for scanner.Scan() {
		line := scanner.Text()
		row := len(rows)
		rows = append(rows, make([]*Galaxy, 0))
		if cols == nil {
			cols = make([][]*Galaxy, len(line))
		}

		for col, ch := range line {
			if ch == '#' {
				galaxy = &Galaxy{col, row}
				galaxies = append(galaxies, galaxy)
				rows[row] = append(rows[row], galaxy)
				cols[col] = append(cols[col], galaxy)
			}
		}
	}

	Expand(rows, 1, false)
	Expand(cols, 1, true)
	fmt.Printf("Part 1: %d\n", TotalDistance(galaxies))

	Expand(rows, 999998, false)
	Expand(cols, 999998, true)
	fmt.Printf("Part 2: %d\n", TotalDistance(galaxies))
}
