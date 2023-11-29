package main

import (
	"fmt"
)

const (
	Add = iota + 1
	Multiply
	Halt = 99
)

func runIntcode(positions *[]int) {
	position := 0

	for {
		switch (*positions)[position] {
		case Halt:
			return
		case Add:
			(*positions)[(*positions)[position + 3]] = (*positions)[(*positions)[position + 1]] + (*positions)[(*positions)[position + 2]]
		case Multiply:
			(*positions)[(*positions)[position + 3]] = (*positions)[(*positions)[position + 1]] * (*positions)[(*positions)[position + 2]]
		default:
			panic("Unexpected operator")
		}

		position += 4
	}
}

func main() {
	positions := make([]int, 0, 10)

	for {
		var value int
		var sep rune
		fmt.Scanf("%d%c", &value, &sep)
		positions = append(positions, value)
		if sep != ',' {
			break
		}
	}

	// Restore 1202 state
	positions[1], positions[2] = 12, 2

	runIntcode(&positions)

	fmt.Println(positions[0])
}
