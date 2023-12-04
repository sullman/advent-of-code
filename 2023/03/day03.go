package main

import (
	"bufio"
	"fmt"
	"os"
)

type Part struct {
	Id int
	Used bool
	Used2 int
}

type Symbol struct {
	Row int
	Col int
	Char rune
}

var Adjacent = [][]int {
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var currentPart *Part
	var symbols []Symbol
	parts := make(map[string]*Part)
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		currentPart = nil

		for col, ch := range line {
			if ch >= '0' && ch <= '9' {
				if currentPart == nil {
					currentPart = new(Part)
					currentPart.Used2 = -1
				}

				currentPart.Id = currentPart.Id * 10 + int(ch - '0')
				parts[fmt.Sprintf("%d,%d", row, col)] = currentPart
			} else if ch != '.' {
				currentPart = nil
				symbol := Symbol{row, col, ch}
				symbols = append(symbols, symbol)
			} else if currentPart != nil {
				currentPart = nil
			}
		}

		row++
	}

	sum := 0
	ratioSum := 0
	for index, symbol := range symbols {
		numNeighbors := 0
		ratio := 0
		if symbol.Char == '*' {
			ratio = 1
		}

		for _, delta := range Adjacent {
			key := fmt.Sprintf("%d,%d", symbol.Row + delta[0], symbol.Col + delta[1])
			part := parts[key]
			if part != nil {
				if !part.Used {
					part.Used = true
					sum += part.Id
				}
				if part.Used2 != index {
					ratio *= part.Id
					part.Used2 = index
					numNeighbors++
				}
			}
		}

		if numNeighbors == 2 {
			ratioSum += ratio
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", ratioSum)
}
