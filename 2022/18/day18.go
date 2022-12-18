package main

import (
	"fmt"
)

var Deltas = [][]int {
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
}

func main() {
	points := make([]string, 0)
	neighbors := make(map[string]int)

	for {
		var point string
		var x, y, z int

		n, _ := fmt.Scanln(&point)

		if n == 0 { break }

		points = append(points, point)
		fmt.Sscanf(point, "%d,%d,%d", &x, &y, &z)
		for _, d := range Deltas {
			key := fmt.Sprintf("%d,%d,%d", x + d[0], y + d[1], z + d[2])
			neighbors[key] = neighbors[key] + 1
		}
	}

	overlap := 0

	for _, point := range points {
		overlap += neighbors[point]
		delete(neighbors, point)
	}

	fmt.Println(len(points) * 6 - overlap)
}
