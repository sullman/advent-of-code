package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var reNumbers = regexp.MustCompile(`\d+`)

func ExtractNumbers(s string) []int {
	matches := reNumbers.FindAllString(s, -1)
	nums := make([]int, len(matches))
	for i, num := range matches {
		nums[i], _ = strconv.Atoi(num)
	}
	return nums
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	times := ExtractNumbers(scanner.Text())

	scanner.Scan()
	distances := ExtractNumbers(scanner.Text())

	total := 1

	// Solve each one as a math problem using the quadratic formula
	for i := 0; i < len(times); i++ {
		// We need to beat it, not tie
		dist := float64(distances[i] + 1)
		time := float64(times[i])
		min := int(math.Ceil((time - math.Sqrt(time * time - 4 * dist)) / 2))
		max := int(math.Floor((time + math.Sqrt(time * time - 4 * dist)) / 2))
		total *= (max - min + 1)
	}

	fmt.Printf("Part 1: %d\n", total)
}
