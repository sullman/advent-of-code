package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var Limits = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	re := regexp.MustCompile(`(\d+) (blue|red|green)(,|;|$)`)

	for scanner.Scan() {
		line := scanner.Text()
		before, after, found := strings.Cut(line, ":")
		if !found { break }
		possible := true

		gameNum, _ := strconv.Atoi(before[5:])
		matches := re.FindAllStringSubmatch(after, -1)
		for _, match := range matches {
			count, _ := strconv.Atoi(match[1])
			color := match[2]
			if Limits[color] < count {
				possible = false
				break
			}
		}

		if possible {
			sum += gameNum
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
}
