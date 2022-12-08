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

	numVisible := 0
	maxScenicScore := 0
	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[row]); col++ {
			visible := false

			visibleUp := 1
			for r := row - 1; r >= 0; r-- {
				if trees[row][col] > trees[r][col] {
					visibleUp++
				} else {
					break
				}
			}

			if row == visibleUp - 1 {
				visibleUp--
				visible = true
			}

			visibleDown := 1
			for r := row + 1; r < len(trees); r++ {
				if trees[row][col] > trees[r][col] {
					visibleDown++
				} else {
					break
				}
			}

			if row + visibleDown == len(trees) {
				visibleDown--
				visible = true
			}

			visibleLeft := 1
			for c := col - 1; c >= 0; c-- {
				if trees[row][col] > trees[row][c] {
					visibleLeft++
				} else {
					break
				}
			}

			if col == visibleLeft - 1 {
				visibleLeft--
				visible = true
			}

			visibleRight := 1
			for c := col + 1; c < len(trees[row]); c++ {
				if trees[row][col] > trees[row][c] {
					visibleRight++
				} else {
					break
				}
			}

			if col + visibleRight == len(trees[row]) {
				visibleRight--
				visible = true
			}

			if visible {
				numVisible++
			}

			scenicScore := visibleLeft * visibleUp * visibleRight * visibleDown
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Println(numVisible)
	fmt.Println(maxScenicScore)
}
