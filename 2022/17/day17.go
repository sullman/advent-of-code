package main

import (
	"fmt"
)

type Space byte

const (
	Air Space = iota
	Settled
	Falling
)

func Max(x int, y int) int {
	if x > y { return x }
	return y
}

const MaxRocks = 2022

func main() {
	shapeTops := [5][8]int {
		{0, 0, 4, 4, 4, 4, 0, 4},
		{0, 0, 5, 6, 5, 0, 0, 6},
		{0, 0, 4, 4, 6, 0, 0, 6},
		{0, 0, 7, 0, 0, 0, 0, 7},
		{0, 0, 5, 5, 0, 0, 0, 5},
	}
	shapeBottoms := [5][7]int {
		{0, 0, 4, 4, 4, 4, 0},
		{0, 0, 5, 4, 5, 0, 0},
		{0, 0, 4, 4, 4, 0, 0},
		{0, 0, 4, 0, 0, 0, 0},
		{0, 0, 4, 4, 0, 0, 0},
	}
	grid := make([][]Space, 1)
	fallingTop := 0
	fallingBottom := 0
	var jets string
	jetIndex := 0
	maxSettledHeight := 0
	numRocks := 0

	grid[0] = make([]Space, 7)
	for i := 0; i < len(grid[0]); i++ {
		grid[0][i] = Settled
	}

	fmt.Scanln(&jets)

	for numRocks < MaxRocks {
		shape := numRocks % len(shapeTops)
		neededRows := maxSettledHeight + shapeTops[shape][7] + 1
		fallingBottom = maxSettledHeight + 4
		fallingTop = maxSettledHeight + 4
		for len(grid) < neededRows {
			grid = append(grid, make([]Space, 7))
		}

		for i := 0; i < len(grid[0]); i++ {
			for row := shapeBottoms[shape][i]; row > 0 && row <= shapeTops[shape][i]; row++ {
				fallingTop = Max(fallingTop, maxSettledHeight + row)
				grid[maxSettledHeight + row][i] = Falling
			}
		}

		for {
			// Jet left/right
			if jets[jetIndex] == '<' {
				canMove := true
				for row := fallingBottom; row <= fallingTop; row++ {
					if grid[row][0] == Falling {
						canMove = false
						break
					}
					for i := 1; i < len(grid[0]); i++ {
						if grid[row][i] == Falling && grid[row][i - 1] == Settled {
							canMove = false
							break
						}
					}
				}

				if canMove {
					for row := fallingBottom; row <= fallingTop; row++ {
						for i := 1; i < len(grid[0]); i++ {
							if grid[row][i] == Falling {
								grid[row][i - 1] = Falling
								grid[row][i] = Air
							}
						}
					}
				}
			} else {
				canMove := true
				for row := fallingBottom; row <= fallingTop; row++ {
					if grid[row][len(grid[0]) - 1] == Falling {
						canMove = false
						break
					}
					for i := 0; i < len(grid[0]) - 1; i++ {
						if grid[row][i] == Falling && grid[row][i + 1] == Settled {
							canMove = false
							break
						}
					}
				}

				if canMove {
					for row := fallingBottom; row <= fallingTop; row++ {
						for i := len(grid[0]) - 2; i >= 0; i-- {
							if grid[row][i] == Falling {
								grid[row][i + 1] = Falling
								grid[row][i] = Air
							}
						}
					}
				}
			}

			jetIndex = (jetIndex + 1) % len(jets)

			// Fall down
			canFall := true
			for row := fallingBottom; canFall && row <= fallingTop; row++ {
				for i := 0; i < len(grid[0]); i++ {
					if grid[row][i] == Falling && grid[row - 1][i] == Settled {
						canFall = false
						break
					}
				}
			}

			if canFall {
				for row := fallingBottom; row <= fallingTop; row++ {
					for i := 0; i < len(grid[0]); i++ {
						if grid[row][i] == Falling {
							grid[row - 1][i] = Falling
							grid[row][i] = Air
						}
					}
				}
				fallingBottom--
				fallingTop--
			} else {
				for row := fallingBottom; row <= fallingTop; row++ {
					for i := 0; i < len(grid[0]); i++ {
						if grid[row][i] == Falling {
							grid[row][i] = Settled
							maxSettledHeight = Max(maxSettledHeight, row)
						}
					}
				}
				numRocks++
				break
			}
		}
	}

	fmt.Println(maxSettledHeight)
}
