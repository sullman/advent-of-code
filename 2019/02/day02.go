package main

import (
	"fmt"
	"errors"
)

const (
	Add = iota + 1
	Multiply
	Halt = 99
)

func runIntcode(program []int) (int, error) {
	memory := make([]int, len(program))
	copy(memory, program)

	instructionPointer := 0

	for {
		switch memory[instructionPointer] {
		case Halt:
			instructionPointer += 1
			return memory[0], nil
		case Add:
			memory[memory[instructionPointer + 3]] = memory[memory[instructionPointer + 1]] + memory[memory[instructionPointer + 2]]
			instructionPointer += 4
		case Multiply:
			memory[memory[instructionPointer + 3]] = memory[memory[instructionPointer + 1]] * memory[memory[instructionPointer + 2]]
			instructionPointer += 4
		default:
			return 0, errors.New("Unexpected operator")
		}
	}
}

func main() {
	memory := make([]int, 0, 10)

	for {
		var value int
		var sep rune
		fmt.Scanf("%d%c", &value, &sep)
		memory = append(memory, value)
		if sep != ',' {
			break
		}
	}

	// Part 1: Run after restoring 1202 state
	memory[1], memory[2] = 12, 2

	result, err := runIntcode(memory)
	fmt.Printf("Part 1: %d\n", result)

	// Part 2: Brute force to find inputs that result in 19690720
	done:
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			memory[1], memory[2] = noun, verb
			result, err = runIntcode(memory)
			if err == nil && result == 19690720 {
				fmt.Printf("Part 2: %d\n", 100 * noun + verb)
				break done
			}
		}
	}
}
