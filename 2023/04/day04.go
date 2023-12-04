package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Preamble = iota
	Winners
	Numbers
)

func main() {
	var winners [100]bool
	reader := bufio.NewReader(os.Stdin)
	sum := 0
	section := Preamble
	num := 0
	nextValue := 1

	for {
		ch, err := reader.ReadByte()
		if err != nil {
			break
		}

		if section == Preamble {
			if ch == ':' {
				section = Winners
			}
		} else if section == Winners {
			switch ch {
			case ' ':
				winners[num] = true
				num = 0
			case '|':
				section = Numbers
				num = 0
			default:
				num = num * 10 + int(ch - '0')
			}
		} else if section == Numbers {
			switch ch {
			case ' ':
				if num != 0 && winners[num] {
					nextValue = nextValue << 1
				}
				num = 0
			case '\n':
				if num != 0 && winners[num] {
					nextValue = nextValue << 1
				}
				sum += nextValue >> 1

				// Reset state for next game
				section = Preamble
				winners = [len(winners)]bool{}
				num = 0
				nextValue = 1
			default:
				num = num * 10 + int(ch - '0')
			}
		}
	}

	fmt.Printf("%d\n", sum)
}
