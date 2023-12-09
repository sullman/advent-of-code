package main

import (
	"fmt"
)

func ExtrapolateSequence(seq []int) (int, int) {
	next := 0
	prev := 0
	sign := 1
	for l := len(seq) - 1; l > 0; l-- {
		next += seq[l]
		prev += sign * seq[0]
		sign *= -1
		nonzero := false
		for i := 0; i < l; i++ {
			seq[i] = seq[i + 1] - seq[i]
			if seq[i] != 0 { nonzero = true }
		}
		if !nonzero { break }
	}
	return prev, next
}

func main() {
	sum1, sum2 := 0, 0
	var num int
	var sep rune
	seq := make([]int, 0)

	for {
		n, _ := fmt.Scanf("%d%c", &num, &sep)
		if n != 2 { break }

		seq = append(seq, num)

		if sep != ' ' {
			prev, next := ExtrapolateSequence(seq)
			sum1 += next
			sum2 += prev
			seq = []int{}
		}
	}

	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: %d\n", sum2)
}
