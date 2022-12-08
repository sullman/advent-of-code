package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var trees [][]byte
	var visible [][]bool

	for scanner.Scan() {
		line := scanner.Text()
		trees = append(trees, []byte(line))
		visible = append(visible, make([]bool, len(line)))
	}

	// Check horizontal visibility
	for row := 0; row < len(trees); row++ {
		prev := byte(0)
		for col := 0; col < len(trees[row]); col++ {
			if trees[row][col] > prev {
				visible[row][col] = true
				prev = trees[row][col]
			}
		}

		prev = 0
		for col := len(trees[row]) - 1; col >= 0; col-- {
			if trees[row][col] > prev {
				visible[row][col] = true
				prev = trees[row][col]
			}
		}
	}

	// Check vertical visibility
	for col := 0; col < len(trees[0]); col++ {
		prev := byte(0)
		for row := 0; row < len(trees); row++ {
			if trees[row][col] > prev {
				visible[row][col] = true
				prev = trees[row][col]
			}
		}

		prev = 0
		for row := len(trees) - 1; row >= 0; row-- {
			if trees[row][col] > prev {
				visible[row][col] = true
				prev = trees[row][col]
			}
		}
	}

	numVisible := 0
	for row := 0; row < len(visible); row++ {
		for col := 0; col < len(visible[row]); col++ {
			if visible[row][col] {
				numVisible++
			}
		}
	}

	fmt.Println(numVisible)
}
