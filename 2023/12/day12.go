package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CountArrangements(conditions []rune, blocks []int, memo map[string]int) int {
	if len(conditions) == 0 && len(blocks) == 0 {
		return 1
	} else if len(conditions) == 0 {
		return 0
	}

	key := fmt.Sprintf("%c,%d,%d", conditions[0], len(conditions), len(blocks))
	if val, found := memo[key]; found {
		return val
	}

	if conditions[0] == '.' {
		var i int
		for i = 1; i < len(conditions); i++ {
			if conditions[i] != '.' { break }
		}
		return CountArrangements(conditions[i:], blocks, memo)
	} else if conditions[0] == '#' && len(blocks) != 0 {
		var i int
		for i = 1; i < blocks[0]; i++ {
			if conditions[i] == '.' { return 0 }
		}

		if conditions[i] == '#' { return 0 }
		return CountArrangements(conditions[(i + 1):], blocks[1:], memo)
	} else if conditions[0] == '?' {
		conditions[0] = '.'
		ifOperational := CountArrangements(conditions, blocks, memo)
		conditions[0] = '#'
		ifDamaged := CountArrangements(conditions, blocks, memo)
		conditions[0] = '?'
		memo[key] = ifOperational + ifDamaged
		return ifOperational + ifDamaged
	}

	return 0
}

func Solve(before, after string) int {
	conditions := []rune(before + ".")
	nums := strings.Split(after, ",")
	blocks := make([]int, len(nums))
	for i, num := range nums {
		blocks[i], _ = strconv.Atoi(num)
	}

	memo := make(map[string]int)

	return CountArrangements(conditions, blocks, memo)
}

func main() {
	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		before, after, found := strings.Cut(scanner.Text(), " ")
		if !found { break }

		sum1 += Solve(before, after)

		before = fmt.Sprintf("%s?%s?%s?%s?%s", before, before, before, before, before)
		after = fmt.Sprintf("%s,%s,%s,%s,%s", after, after, after, after, after)

		sum2 += Solve(before, after)
	}

	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: %d\n", sum2)
}
