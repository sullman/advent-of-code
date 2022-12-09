package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	visited := make(map[string]bool)

	headX := 0
	headY := 0
	tailX := 0
	tailY := 0

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
			headX += dx
			headY += dy

			tdx := headX - tailX
			tdy := headY - tailY
			move := false

			switch {
			case tdx == 2:
				tdx = 1
				move = true
			case tdx == -2:
				tdx = -1
				move = true
			case tdy == 2:
				tdy = 1
				move = true
			case tdy == -2:
				tdy = -1
				move = true
			}

			if move {
				tailX += tdx
				tailY += tdy
			}

			// fmt.Printf("Head now at %d,%d and tail at %d,%d\n", headX, headY, tailX, tailY)

			visited[fmt.Sprintf("%d,%d", tailX, tailY)] = true
		}
	}

	fmt.Println(len(visited))
}
