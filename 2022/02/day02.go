package main

import (
	"bufio"
	"fmt"
	"os"
)

var scores1 = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

var scores2 = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total1 := 0
	total2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		total1 += scores1[line]
		total2 += scores2[line]
	}

	fmt.Println(total1)
	fmt.Println(total2)
}
