package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var allStacks []*list.List
	initialized := false
	numStacks := 0

	for scanner.Scan() {
		line := scanner.Text()

		if initialized {
			var numToMove, from, to int
			n, _ := fmt.Sscanf(line, "move %d from %d to %d", &numToMove, &from, &to)
			if n == 3 {
				for ; numToMove > 0; numToMove-- {
					elem := allStacks[from - 1].Back()
					allStacks[to - 1].PushBack(elem.Value)
					allStacks[from - 1].Remove(elem)
				}
			}
		} else if line[1] == '1' {
			initialized = true
		} else {
			if numStacks == 0 {
				numStacks = (len(line) + 1) / 4
				allStacks = make([]*list.List, numStacks)
				for i := 0; i < numStacks; i++ {
					allStacks[i] = list.New()
				}
			}

			for i := 0; i < numStacks; i++ {
				if line[4 * i + 1] != ' ' {
					allStacks[i].PushFront(line[4 * i + 1])
				}
			}
		}
	}

	result := ""

	for i := 0; i < numStacks; i++ {
		result = result + string(allStacks[i].Back().Value.(byte))
	}

	fmt.Println(result)
}
