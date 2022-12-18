package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
	z int
	str string
}

var Adjacent = [][]int {
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
}

func main() {
	points := make([]Point, 0)
	pointsByKey := make(map[string]bool)
	neighbors := make(map[string]int)
	min := Point{x: 1 << 30}

	for {
		var point Point

		n, _ := fmt.Scanln(&point.str)

		if n == 0 { break }

		fmt.Sscanf(point.str, "%d,%d,%d", &point.x, &point.y, &point.z)

		points = append(points, point)
		pointsByKey[point.str] = true

		if point.x < min.x {
			min = point
		}

		for _, d := range Adjacent {
			key := fmt.Sprintf("%d,%d,%d", point.x + d[0], point.y + d[1], point.z + d[2])
			neighbors[key] = neighbors[key] + 1
		}
	}

	overlap := 0

	for _, point := range points {
		overlap += neighbors[point.str]
		delete(neighbors, point.str)
	}

	part1 := len(points) * 6 - overlap
	fmt.Println(part1)

	// Part 2: Start from a single known neighbor outside the surface and walk
	// all the way around it.
	// (In hindsight, I guess this is a quirky version of flood fill. Instead of
	// defining a bounding box and checking the entire thing by walking the six
	// adjacent directions, this pretends like there's no bounding box and only
	// keeps walking if it's still on the surface. That means moving diagonally
	// has to be allowed, which turns out to be slightly tricky. It _might_ be
	// more efficient on certain shapes, but probably isn't worth the complexity.
	exterior := 0
	visited := make(map[string]bool)
	candidates := make([]Point, 0, part1)

	min.x -= 1
	min.str = fmt.Sprintf("%d,%d,%d", min.x, min.y, min.z)
	candidates = append(candidates, min)

	for len(candidates) > 0 {
		point := candidates[0]
		candidates = candidates[1:]
		visited[point.str] = true

		touching := 0

		for _, d := range Adjacent {
			key := fmt.Sprintf("%d,%d,%d", point.x + d[0], point.y + d[1], point.z + d[2])
			if pointsByKey[key] {
				touching++
			}
		}

		if touching == 0 { continue }

		exterior += touching

		for i, d1 := range Adjacent {
			var p Point
			p.x = point.x + d1[0]
			p.y = point.y + d1[1]
			p.z = point.z + d1[2]
			p.str = fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)

			if !pointsByKey[p.str] && !visited[p.str] {
				visited[p.str] = true
				candidates = append(candidates, p)
			}

			for j, d2 := range Adjacent {
				if i / 2 == j / 2 { continue }
				intermediateKey := fmt.Sprintf("%d,%d,%d", point.x + d2[0], point.y + d2[1], point.z + d2[2])
				if pointsByKey[p.str] != pointsByKey[intermediateKey] {
					var p2 Point
					p2.x = point.x + d1[0] + d2[0]
					p2.y = point.y + d1[1] + d2[1]
					p2.z = point.z + d1[2] + d2[2]
					p2.str = fmt.Sprintf("%d,%d,%d", p2.x, p2.y, p2.z)
					if !pointsByKey[p2.str] && !visited[p2.str] {
						visited[p2.str] = true
						candidates = append(candidates, p2)
					}
				}
			}
		}
	}

	fmt.Println(exterior)
}
