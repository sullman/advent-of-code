package main

import (
	"fmt"
)

func FindNextInSequence(seq []int) int {
	sum := 0
	for l := len(seq) - 1; l > 0; l-- {
		sum += seq[l]
		nonzero := false
		for i := 0; i < l; i++ {
			seq[i] = seq[i + 1] - seq[i]
			if seq[i] != 0 { nonzero = true }
		}
		if !nonzero { break }
	}
	return sum
}

func main() {
	sum := 0
	var num int
	var sep rune
	seq := make([]int, 0)

	for {
		n, _ := fmt.Scanf("%d%c", &num, &sep)
		if n != 2 { break }

		seq = append(seq, num)

		if sep != ' ' {
			sum += FindNextInSequence(seq)
			seq = []int{}
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
}
