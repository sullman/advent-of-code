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

type CycleState struct {
	numRocks int
	maxHeight int
}

// const MaxRocks = 2022
const MaxRocks = 1000000000000

func MemoKey(grid [][7]Space, maxRow int, shape int, jet int) string {
	key := fmt.Sprintf("%d %d ", shape, jet)
	for row := maxRow; row > 0 && row > maxRow - 20; row-- {
		for i := 0; i < len(grid[0]); i++ {
			key = key + fmt.Sprint(grid[row % len(grid)][i])
		}
	}
	return key
}

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
	var grid [100][7]Space
	fallingTop := 0
	fallingBottom := 0
	var jets string
	jetIndex := 0
	maxSettledHeight := 0
	numRocks := 0
	maxRow := len(grid) - 1
	patterns := make(map[string]CycleState)

	for i := 0; i < len(grid[0]); i++ {
		grid[0][i] = Settled
	}

	fmt.Scanln(&jets)

	for numRocks < MaxRocks {
		shape := numRocks % len(shapeTops)
		memo := MemoKey(grid[0:], maxSettledHeight, shape, jetIndex)
		prev, pattern := patterns[memo]
		if pattern {
			// fmt.Printf("Found a pattern after %d rocks / %d height! Last time %d / %d\n", numRocks, maxSettledHeight, prev.numRocks, prev.maxHeight)
			cycleLength := numRocks - prev.numRocks
			if (MaxRocks - numRocks) % cycleLength == 0 {
				// fmt.Println("Skipping ahead based on cycle!")
				// fmt.Printf("Cycle height is %d\n", maxSettledHeight - prev.maxHeight)
				// fmt.Printf("Number of cycles to repeat is %d\n", (MaxRocks - numRocks) / cycleLength)
				maxSettledHeight += (maxSettledHeight - prev.maxHeight) * ((MaxRocks - numRocks) / cycleLength)
				numRocks = MaxRocks
				break
			} else {
				// fmt.Println("Need to keep going to make the cycle finish on the right number")
			}
		} else {
			patterns[memo] = CycleState{numRocks, maxSettledHeight}
		}

		neededRows := maxSettledHeight + shapeTops[shape][7] + 1
		fallingBottom = maxSettledHeight + 4
		fallingTop = maxSettledHeight + 4
		for row := maxRow + 1; row <= neededRows; row++ {
			for i := 0; i < len(grid[0]); i++ {
				grid[row % len(grid)][i] = Air
			}
		}
		maxRow = neededRows

		for i := 0; i < len(grid[0]); i++ {
			for row := shapeBottoms[shape][i]; row > 0 && row <= shapeTops[shape][i]; row++ {
				fallingTop = Max(fallingTop, maxSettledHeight + row)
				grid[(maxSettledHeight + row) % len(grid)][i] = Falling
			}
		}

		for {
			// Jet left/right
			if jets[jetIndex] == '<' {
				canMove := true
				for row := fallingBottom; row <= fallingTop; row++ {
					if grid[row % len(grid)][0] == Falling {
						canMove = false
						break
					}
					for i := 1; i < len(grid[0]); i++ {
						if grid[row % len(grid)][i] == Falling && grid[row % len(grid)][i - 1] == Settled {
							canMove = false
							break
						}
					}
				}

				if canMove {
					for row := fallingBottom; row <= fallingTop; row++ {
						for i := 1; i < len(grid[0]); i++ {
							if grid[row % len(grid)][i] == Falling {
								grid[row % len(grid)][i - 1] = Falling
								grid[row % len(grid)][i] = Air
							}
						}
					}
				}
			} else {
				canMove := true
				for row := fallingBottom; row <= fallingTop; row++ {
					if grid[row % len(grid)][len(grid[0]) - 1] == Falling {
						canMove = false
						break
					}
					for i := 0; i < len(grid[0]) - 1; i++ {
						if grid[row % len(grid)][i] == Falling && grid[row % len(grid)][i + 1] == Settled {
							canMove = false
							break
						}
					}
				}

				if canMove {
					for row := fallingBottom; row <= fallingTop; row++ {
						for i := len(grid[0]) - 2; i >= 0; i-- {
							if grid[row % len(grid)][i] == Falling {
								grid[row % len(grid)][i + 1] = Falling
								grid[row % len(grid)][i] = Air
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
					if grid[row % len(grid)][i] == Falling && grid[(row - 1) % len(grid)][i] == Settled {
						canFall = false
						break
					}
				}
			}

			if canFall {
				for row := fallingBottom; row <= fallingTop; row++ {
					for i := 0; i < len(grid[0]); i++ {
						if grid[row % len(grid)][i] == Falling {
							grid[(row - 1) % len(grid)][i] = Falling
							grid[row % len(grid)][i] = Air
						}
					}
				}
				fallingBottom--
				fallingTop--
			} else {
				for row := fallingBottom; row <= fallingTop; row++ {
					for i := 0; i < len(grid[0]); i++ {
						if grid[row % len(grid)][i] == Falling {
							grid[row % len(grid)][i] = Settled
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
