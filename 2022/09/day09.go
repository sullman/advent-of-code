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

func reconcile(positions []Position, visited []map[string]bool, index int) {
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
		visited[index][fmt.Sprintf("%d,%d", positions[index].x, positions[index].y)] = true
		if index < len(positions) - 1 {
			reconcile(positions, visited, index + 1)
		}
	}
}

const NumKnots = 10

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	positions := make([]Position, NumKnots)
	visited := make([]map[string]bool, NumKnots)

	for i := 0; i < NumKnots; i++ {
		visited[i] = map[string]bool { "0,0": true }
	}

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
			reconcile(positions, visited, 1)
			visited[0][fmt.Sprintf("%d,%d", positions[0].x, positions[0].y)] = true
		}
	}

	fmt.Println(len(visited[1]))
	fmt.Println(len(visited[NumKnots - 1]))
}
