package main

import (
	"fmt"
)

type CycleState struct {
	numRocks int
	maxHeight int
}

// const MaxRocks = 2022
const MaxRocks = 1000000000000

var Shapes = [][]byte{
	{0b00011110},
	{0b00001000, 0b00011100, 0b00001000},
	{0b00011100, 0b00000100, 0b00000100},
	{0b00010000, 0b00010000, 0b00010000, 0b00010000},
	{0b00011000, 0b00011000},
}

func main() {
	// We'll add to it if necessary, but since each row is just a byte we might
	// as well preallocate a decent chunk of memory instead of shifting it every
	// time we append.
	grid := make([]byte, 4096)
	grid[0] = 0b11111111

	fallingBottom := 0
	var jets string
	jetIndex := 0
	maxSettledHeight := 0
	numRocks := 0
	patterns := make(map[uint64]CycleState)

	fmt.Scanln(&jets)

	for numRocks < MaxRocks {
		shapeIndex := numRocks % len(Shapes)

		// Construct a memo key that factors in relevant state. We definitely need
		// the shape index (3 bits) and jet index (14 bits), which leaves enough
		// room to include the top 6 rows from the grid and use a uint64.
		var memo uint64
		for row := maxSettledHeight; row > 0 && row > maxSettledHeight - 6; row-- {
			memo = (memo << 8) | uint64(grid[row])
		}
		memo = memo | (uint64(jetIndex) << 47) | (uint64(shapeIndex) << 61)

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

		fallingBottom = maxSettledHeight + 4
		falling := make([]byte, len(Shapes[shapeIndex]))
		for i := 0; i < len(falling); i++ {
			falling[i] = Shapes[shapeIndex][i]
		}

		neededRows := fallingBottom + len(falling) + 1
		for len(grid) < neededRows {
			grid = append(grid, 0)
		}


		for {
			// Jet left/right
			canMove := true
			for i, mask := range falling {
				check := grid[fallingBottom + i]
				if jets[jetIndex] == '<' {
					if mask & 0b01000000 != 0 || (mask << 1) & check != 0 {
						canMove = false
						break
					}
				} else {
					if mask & 0b00000001 != 0 || (mask >> 1) & check != 0 {
						canMove = false
						break
					}
				}
			}

			if canMove {
				for i := 0; i < len(falling); i++ {
					if jets[jetIndex] == '<' {
						falling[i] = falling[i] << 1
					} else {
						falling[i] = falling[i] >> 1
					}
				}
			}

			jetIndex = (jetIndex + 1) % len(jets)

			// Fall down
			canFall := true
			for i, mask := range falling {
				check := grid[fallingBottom + i - 1]
				if mask & check != 0 {
					canFall = false
					break
				}
			}

			if canFall {
				fallingBottom--
			} else {
				for i, mask := range falling {
					if fallingBottom + i > maxSettledHeight {
						maxSettledHeight = fallingBottom + i
					}
					grid[fallingBottom + i] = grid[fallingBottom + i] | mask
				}
				numRocks++
				break
			}
		}
	}

	fmt.Println(maxSettledHeight)
}
