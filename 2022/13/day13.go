package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	EOL = -3
	OpenList = -2
	CloseList = -1
)

type Tokenizer struct {
	str string
	index int
	bufferedVal int
	bufferedParens int
}

func NewTokenizer(str string) *Tokenizer {
	return &Tokenizer{str: str, bufferedVal: -1}
}

func (t *Tokenizer) Next() int {
	if t.bufferedVal >= 0 {
		val := t.bufferedVal
		t.bufferedVal = -1
		return val
	} else if t.bufferedParens > 0 {
		t.bufferedParens--
		return CloseList
	} else if t.index >= len(t.str) {
		return EOL
	}

	if t.str[t.index] == '[' {
		t.index++
		return OpenList
	} else if t.str[t.index] == ']' {
		t.index++
		return CloseList
	}

	val := 0
	for ; t.str[t.index] >= '0' && t.str[t.index] <= '9'; t.index++ {
		val = 10 * val + int(t.str[t.index] - '0')
	}

	if t.str[t.index] == ',' {
		t.index++
	}

	return val
}

func (t *Tokenizer) SynthesizeList(val int) {
	t.bufferedVal = val
	t.bufferedParens++
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	index := 0
	sum := 0

	for scanner.Scan() {
		index++
		left := NewTokenizer(scanner.Text())
		scanner.Scan()
		right := NewTokenizer(scanner.Text())
		scanner.Scan()

		for l, r := left.Next(), right.Next(); l != EOL && r != EOL; l, r = left.Next(), right.Next() {
			if l == r {
				// Doesn't matter what they are if they're the same, just keep going
				continue
			} else if r == CloseList {
				fmt.Printf("Pair %d is in the wrong order, right ran out of items\n", index)
				break
			} else if l == CloseList {
				fmt.Printf("Pair %d is in the right order, left ran out of items\n", index)
				sum += index
				break
			} else if l == OpenList {
				right.SynthesizeList(r)
			} else if r == OpenList {
				left.SynthesizeList(l)
			} else if l < r {
				fmt.Printf("Pair %d is in the right order, %d < %d\n", index, l, r)
				sum += index
				break
			} else { // l > r
				fmt.Printf("Pair %d is in the wrong order, %d > %d\n", index, l, r)
				break
			}
		}
	}

	fmt.Println(sum)
}
