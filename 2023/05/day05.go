package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	Length int
	Moved bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	seeds := strings.Split(scanner.Text()[7:], " ")

	values := make([]int, len(seeds))
	newValues := make([]int, len(seeds))
	ranges := make([]Range, len(seeds) / 2)

	for i, str := range seeds {
		newValues[i], _ = strconv.Atoi(str)
	}

	for i := 0; i < len(ranges); i++ {
		ranges[i] = Range{newValues[2 * i], newValues[2 * i + 1], false}
	}

	var dstStart, srcStart, rangeLen int

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			scanner.Scan()
			for i, val := range newValues {
				values[i] = val
			}
			for i, _ := range ranges {
				ranges[i].Moved = false
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

		for i := 0; i < len(ranges); i++ {
			if ranges[i].Moved { continue }
			start := ranges[i].Start
			end := start + ranges[i].Length
			if start < srcStart && end > srcStart + rangeLen {
				ranges = append(ranges, Range{start, srcStart - start, false})
				ranges = append(ranges, Range{srcStart + rangeLen, end - srcStart - rangeLen, false})
				ranges[i].Start = dstStart
				ranges[i].Length = rangeLen
				ranges[i].Moved = true
			} else if start < srcStart && end > srcStart {
				ranges = append(ranges, Range{start, srcStart - start, false})
				ranges[i].Start = dstStart
				ranges[i].Length = end - srcStart
				ranges[i].Moved = true
			} else if start >= srcStart && end <= srcStart + rangeLen {
				ranges[i].Start = dstStart + start - srcStart
				ranges[i].Moved = true
			} else if start < srcStart + rangeLen && end > srcStart + rangeLen {
				ranges = append(ranges, Range{srcStart + rangeLen, end - srcStart - rangeLen, false})
				ranges[i].Start = dstStart + start - srcStart
				ranges[i].Length = srcStart + rangeLen - start
				ranges[i].Moved = true
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

	min = ranges[0].Start
	for _, r := range ranges {
		if r.Start < min {
			min = r.Start
		}
	}

	fmt.Printf("Part 2: %d\n", min)
}
