package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	HighCard byte = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var CardRanks = map[byte]byte {
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type Hand struct {
	Cards [5]byte
	Type byte
	Bid int
}

func ClassifyHand(cards []byte) byte {
	counts := make([]byte, 15)
	for _, card := range cards {
		counts[card] += 1
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	if counts[0] == 5 {
		return FiveOfAKind
	} else if counts[0] == 4 {
		return FourOfAKind
	} else if counts[0] == 3 && counts[1] == 2 {
		return FullHouse
	} else if counts[0] == 3 {
		return ThreeOfAKind
	} else if counts[1] == 2 {
		return TwoPair
	} else if counts[0] == 2 {
		return OnePair
	}

	return HighCard
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hands := make([]Hand, 0)
	var cards string
	var bid int

	for scanner.Scan() {
		n, _ := fmt.Sscanf(scanner.Text(), "%s %d", &cards, &bid)
		if (n != 2) { continue }

		hand := Hand{}

		for i := 0; i < len(hand.Cards); i++ {
			hand.Cards[i] = CardRanks[cards[i]]
		}
		hand.Bid = bid
		hand.Type = ClassifyHand(hand.Cards[0:])

		hands = append(hands, hand)
	}

	total := 0

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			for k := 0; k < len(hands[i].Cards); k++ {
				if hands[i].Cards[k] != hands[j].Cards[k] {
					return hands[i].Cards[k] < hands[j].Cards[k]
				}
			}
		} else {
			return hands[i].Type < hands[j].Type
		}

		panic("Huh?")
	})

	for i, hand := range hands {
		total += (i + 1) * hand.Bid
	}

	fmt.Printf("Part 1: %d\n", total)
}
