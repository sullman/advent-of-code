package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Part1(in chan string, out chan int) {
	sum := 0

	for line := range in {
		found := false
		first := 0
		last := 0

		for _, r := range line {
			if r >= '0' && r <= '9' {
				last = int(r - '0')
				if !found {
					found = true
					first = last
				}
			}
		}

		sum += 10 * first + last
	}

	out <- sum
}

func Part2(in chan string, out chan int) {
	sum := 0
	mappings := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"zero": 0,
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	keys := make([]string, len(mappings))
	var i int
	for key := range mappings {
		keys[i] = key
		i++
	}
	pattern := strings.Join(keys, "|")
	forwardRe := regexp.MustCompile(pattern)
	backwardRe := regexp.MustCompile(Reverse(pattern))

	for line := range in {
		first := 0
		last := 0

		match := forwardRe.FindString(line)
		if len(match) > 0 {
			first = mappings[match]
			match = backwardRe.FindString(Reverse(line))
			last = mappings[Reverse(match)]
			sum += 10 * first + last
		}
	}

	out <- sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	in1 := make(chan string)
	out1 := make(chan int)
	go Part1(in1, out1)

	in2 := make(chan string)
	out2 := make(chan int)
	go Part2(in2, out2)

	for scanner.Scan() {
		line := scanner.Text()

		in1 <- line
		in2 <- line
	}

	close(in1)
	fmt.Printf("Part 1: %d\n", <-out1)

	close(in2)
	fmt.Printf("Part 2: %d\n", <-out2)
}
