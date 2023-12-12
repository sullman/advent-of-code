package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CountArrangements(conditions []rune, blocks []int) int {
	if len(conditions) == 0 && len(blocks) == 0 {
		return 1
	} else if len(conditions) == 0 {
		return 0
	}

	if conditions[0] == '.' {
		var i int
		for i = 1; i < len(conditions); i++ {
			if conditions[i] != '.' { break }
		}
		return CountArrangements(conditions[i:], blocks)
	} else if conditions[0] == '#' && len(blocks) != 0 {
		var i int
		for i = 1; i < blocks[0]; i++ {
			if conditions[i] == '.' { return 0 }
		}

		if conditions[i] == '#' { return 0 }
		return CountArrangements(conditions[(i + 1):], blocks[1:])
	} else if conditions[0] == '?' {
		conditions[0] = '.'
		ifOperational := CountArrangements(conditions, blocks)
		conditions[0] = '#'
		ifDamaged := CountArrangements(conditions, blocks)
		conditions[0] = '?'
		return ifOperational + ifDamaged
	}

	return 0
}

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		before, after, found := strings.Cut(scanner.Text(), " ")
		if !found { break }

		conditions := []rune(before + ".")
		nums := strings.Split(after, ",")
		blocks := make([]int, len(nums))
		for i, num := range nums {
			blocks[i], _ = strconv.Atoi(num)
		}

		sum += CountArrangements(conditions, blocks)
	}

	fmt.Printf("Part 1: %d\n", sum)
}
