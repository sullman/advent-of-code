package main

import (
	"bufio"
	"fmt"
	"os"
)

func FindReflection(nums []uint) int {
	outer:
	for i := 1; i < len(nums); i++ {
		for j := i; j < len(nums) && i - (j - i) > 0; j++ {
			if nums[j] != nums[i - (j - i) - 1] { continue outer }
		}

		return i
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	rows := make([]uint, 0)
	var cols []uint
	row := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			summary := FindReflection(cols)
			if (summary == 0) {
				summary = FindReflection(rows) * 100
			}

			sum += summary

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

	fmt.Printf("Part 1: %d\n", sum)
}
