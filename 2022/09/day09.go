package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x int
	y int
}

func reconcile(positions []Position, index int) {
	dx := positions[index - 1].x - positions[index].x
	dy := positions[index - 1].y - positions[index].y
	move := false

	if dx == 2 {
		dx = 1
		move = true
	} else if dx == -2 {
		dx = -1
		move = true
	}

	if dy == 2 {
		dy = 1
		move = true
	} else if dy == -2 {
		dy = -1
		move = true
	}

	if move {
		positions[index].x += dx
		positions[index].y += dy
		if index < len(positions) - 1 {
			reconcile(positions, index + 1)
		}
	}
}

const NumKnots = 10

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	positions := make([]Position, NumKnots)
	visited := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		var direction string
		var count int
		fmt.Sscanf(line, "%s %d", &direction, &count)

		var dx, dy int

		switch direction {
		case "U":
			dy = 1
		case "R":
			dx = 1
		case "D":
			dy = -1
		case "L":
			dx = -1
		}

		for i := 0; i < count; i++ {
			positions[0].x += dx
			positions[0].y += dy
			reconcile(positions, 1)
			visited[fmt.Sprintf("%d,%d", positions[NumKnots - 1].x, positions[NumKnots - 1].y)] = true
		}
	}

	fmt.Println(len(visited))
}
