package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	seeds := strings.Split(scanner.Text()[7:], " ")

	values := make([]int, len(seeds))
	newValues := make([]int, len(seeds))

	for i, str := range seeds {
		newValues[i], _ = strconv.Atoi(str)
	}

	var dstStart, srcStart, rangeLen int

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			scanner.Scan()
			for i, val := range newValues {
				values[i] = val
			}
			continue
		}

		n, _ := fmt.Sscanf(line, "%d %d %d", &dstStart, &srcStart, &rangeLen)
		if n != 3 { panic("WTF") }

		for i, val := range values {
			if val >= srcStart && val < srcStart + rangeLen {
				newValues[i] = dstStart + val - srcStart
			}
		}
	}

	min := newValues[0]
	for _, val := range newValues {
		if val < min {
			min = val
		}
	}

	fmt.Printf("Part 1: %d\n", min)
}
