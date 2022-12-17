package main

import (
	"fmt"
)

func Max(x int, y int) int {
	if x > y { return x }
	return y
}

const MaxRocks = 2022

func main() {
	shapeTops := [5][7]int {
		{0, 0, 4, 4, 4, 4, 0},
		{0, 0, 5, 6, 5, 0, 0},
		{0, 0, 4, 4, 6, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 5, 5, 0, 0, 0},
	}
	shapeBottoms := [5][7]int {
		{0, 0, 4, 4, 4, 4, 0},
		{0, 0, 5, 4, 5, 0, 0},
		{0, 0, 4, 4, 4, 0, 0},
		{0, 0, 4, 0, 0, 0, 0},
		{0, 0, 4, 4, 0, 0, 0},
	}
	var settled [7]int
	var fallingBottom [7]int
	var fallingTop [7]int
	var jets string
	jetIndex := 0
	maxSettledHeight := 0
	numRocks := 0

	fmt.Scanln(&jets)

	for numRocks < MaxRocks {
		// fmt.Printf("After %d settled rocks, max height is %d\n", numRocks, maxSettledHeight)
		shape := numRocks % len(shapeTops)
		for i := 0; i < len(fallingTop); i++ {
			if shapeBottoms[shape][i] != 0 {
				fallingBottom[i] = maxSettledHeight + shapeBottoms[shape][i]
				fallingTop[i] = maxSettledHeight + shapeTops[shape][i]
			} else {
				fallingBottom[i] = 0
				fallingTop[i] = 0
			}
		}

		for {
			// fmt.Printf("Starting %c tick with settled=%v and falling=%v\n", jets[jetIndex], settled, fallingBottom)
			// Jet left/right
			if jets[jetIndex] == '<' {
				canMove := fallingBottom[0] == 0
				for i := 0; i < len(settled) - 1; i++ {
					if fallingBottom[i + 1] > 0 && settled[i] >= fallingBottom[i + 1] {
						canMove = false
					}
				}

				if canMove {
					for i := 0; i < len(settled) - 1; i++ {
						fallingBottom[i] = fallingBottom[i + 1]
						fallingTop[i] = fallingTop[i + 1]
					}
					fallingBottom[len(settled) - 1] = 0
					fallingTop[len(settled) - 1] = 0
				}
			} else {
				canMove := fallingBottom[len(settled) - 1] == 0
				for i := 1; i < len(settled); i++ {
					if fallingBottom[i - 1] > 0 && settled[i] >= fallingBottom[i - 1] {
						canMove = false
					}
				}

				if canMove {
					for i := len(settled) - 1; i > 0; i-- {
						fallingBottom[i] = fallingBottom[i - 1]
						fallingTop[i] = fallingTop[i - 1]
					}
					fallingBottom[0] = 0
					fallingTop[0] = 0
				}
			}

			jetIndex = (jetIndex + 1) % len(jets)

			// Fall down
			canFall := true
			for i := 0; i < len(settled); i++ {
				if fallingBottom[i] > 0 && settled[i] >= fallingBottom[i] - 1 {
					canFall = false
				}
			}

			if canFall {
				for i := 0; i < len(settled); i++ {
					fallingBottom[i] = Max(0, fallingBottom[i] - 1)
					fallingTop[i] = Max(0, fallingTop[i] - 1)
				}
			} else {
				for i := 0; i < len(settled); i++ {
					settled[i] = Max(settled[i], fallingTop[i])
					maxSettledHeight = Max(maxSettledHeight, settled[i])
				}
				numRocks++
				break
			}
		}
	}

	fmt.Println(maxSettledHeight)
}
