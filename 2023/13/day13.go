package main

import (
	"bufio"
	"fmt"
	"os"
)

func FindReflection(nums []uint, requiredSmudges int) int {
	outer:
	for i := 1; i < len(nums); i++ {
		smudges := 0

		for j := i; j < len(nums) && i - (j - i) > 0 && smudges <= requiredSmudges; j++ {
			diff := nums[j] ^ nums[i - (j - i) - 1]
			if diff != 0 && diff & (diff - 1) == 0 {
				smudges++
			} else if diff != 0 {
				continue outer
			}
		}

		if smudges == requiredSmudges {
			return i
		}
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum1, sum2 := 0, 0

	rows := make([]uint, 0)
	var cols []uint
	row := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			summary := FindReflection(cols, 0)
			if (summary == 0) {
				summary = FindReflection(rows, 0) * 100
			}
			sum1 += summary

			summary = FindReflection(cols, 1)
			if (summary == 0) {
				summary = FindReflection(rows, 1) * 100
			}
			sum2 += summary

			row = 0
			rows = make([]uint, 0)
			continue
		}

		rows = append(rows, 0)
		if row == 0 {
			cols = make([]uint, len(line))
		}

		for col, ch := range line {
			var val uint
			if ch == '#' { val = 1 }
			rows[row] = (rows[row] << 1) + val
			cols[col] = (cols[col] << 1) + val
		}

		row++
	}

	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: %d\n", sum2)
}
