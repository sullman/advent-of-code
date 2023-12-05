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
	points := 0
	section := Preamble
	num := 0
	nextValue := 1
	numWinners := 0
	cardCounts := make([]int, 0, 100)
	numPlayed := 0

	for {
		ch, err := reader.ReadByte()
		if err != nil {
			break
		}

		if section == Preamble {
			if ch == ':' {
				section = Winners
				if len(cardCounts) == 0 {
					cardCounts = append(cardCounts, 0)
				}
				cardCounts[0] += 1
				numPlayed += cardCounts[0]
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
					numWinners++
				}
				num = 0
			case '\n':
				if num != 0 && winners[num] {
					nextValue = nextValue << 1
					numWinners++
				}
				points += nextValue >> 1

				for len(cardCounts) <= numWinners {
					cardCounts = append(cardCounts, 0)
				}
				for i := 1; i <= numWinners; i++ {
					cardCounts[i] += cardCounts[0]
				}

				// Reset state for next game
				section = Preamble
				winners = [len(winners)]bool{}
				cardCounts = cardCounts[1:]
				num = 0
				nextValue = 1
				numWinners = 0
			default:
				num = num * 10 + int(ch - '0')
			}
		}
	}

	fmt.Printf("Part 1: %d\n", points)
	fmt.Printf("Part 2: %d\n", numPlayed)
}
