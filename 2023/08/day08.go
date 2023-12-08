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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, left, right string
	numMoves := 0
	nodes := make(map[string]Node)

	scanner.Scan()
	moves := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		n, _ := fmt.Sscanf(scanner.Text(), "%3s = (%3s, %3s)", &name, &left, &right)
		if n != 3 { continue }
		nodes[name] = Node{left, right}
	}

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
