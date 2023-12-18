package main

import (
	"fmt"
)

type Pair struct {
	x int
	y int
}

var Movements = map[rune]Pair{
	'U': Pair{0, 1},
	'R': Pair{1, 0},
	'D': Pair{0, -1},
	'L': Pair{-1, 0},
}

func Encode(x, y int) uint {
	return uint(((uint32(x) & 0xffff) << 16) | (uint32(y) & 0xffff))
}

func main() {
	var dir rune
	var n, length, color int
	var minX, minY, maxX, maxY int
	var x, y int

	terrain := make(map[uint]int)

	for {
		n, _ = fmt.Scanf("%c %d (#%x)\n", &dir, &length, &color)
		if n != 3 { break }

		move := Movements[dir]
		for i := 0; i < length; i++ {
			x += move.x
			y += move.y
			if x < minX {
				minX = x
			} else if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			} else if y > maxY {
				maxY = y
			}
			terrain[Encode(x, y)] = color
		}
	}

	if x != 0 || y != 0 {
		panic("Not a closed loop?")
	}

	// Add a moat to ensure we can flood the whole outside
	minX--
	minY--
	maxX++
	maxY++

	empty := make(map[uint]bool)
	queue := make([]Pair, 1, len(terrain))
	queue[0] = Pair{minX, minY}
	empty[Encode(minX, minY)] = true

	for len(queue) != 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, move := range Movements {
			x, y = pos.x + move.x, pos.y + move.y
			if x < minX || x > maxX || y < minY || y > maxY { continue }
			enc := Encode(x, y)
			if terrain[enc] == 0 && !empty[enc] {
				empty[enc] = true
				queue = append(queue, Pair{x, y})
			}
		}
	}

	totalArea := (maxX - minX + 1) * (maxY - minY + 1)
	fmt.Printf("Part 1: %d\n", totalArea - len(empty))
}
