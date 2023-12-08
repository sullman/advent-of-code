package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Left string
	Right string
}

func LCM(nums []int) int {
	vals := make([]int, len(nums))
	vals[0] = nums[0]

	for i := 0; i < len(nums) - 1; {
		if vals[i] < vals[i + 1] {
			vals[i] += nums[i]
			i = 0
		} else if vals[i] > vals[i + 1] {
			vals[i + 1] += nums[i + 1]
		} else {
			i++
		}
	}

	return vals[0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, left, right string
	numMoves := 0
	nodes := make(map[string]Node)
	ghosts := make([]string, 0)

	scanner.Scan()
	moves := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		n, _ := fmt.Sscanf(scanner.Text(), "%3s = (%3s, %3s)", &name, &left, &right)
		if n != 3 { continue }
		nodes[name] = Node{left, right}
		if name[2] == 'A' {
			ghosts = append(ghosts, name)
		}
	}

	// Sort of annoying, the second demo input doesn't have ZZZ
	if _, hasZZZ := nodes["ZZZ"]; hasZZZ {
		current := "AAA"
		for current != "ZZZ" {
			if moves[numMoves % len(moves)] == 'L' {
				current = nodes[current].Left
			} else {
				current = nodes[current].Right
			}
			numMoves++
		}

		fmt.Printf("Part 1: %d\n", numMoves)
	}

	cycles := make([]int, len(ghosts))
	numFound := 0

	numMoves = 0
	for numFound < len(cycles) {
		goLeft := moves[numMoves % len(moves)] == 'L'
		numMoves++
		for i, name := range ghosts {
			if goLeft {
				ghosts[i] = nodes[name].Left
			} else {
				ghosts[i] = nodes[name].Right
			}
			if ghosts[i][2] == 'Z' && cycles[i] == 0 {
				cycles[i] = numMoves
				numFound++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", LCM(cycles))
}
