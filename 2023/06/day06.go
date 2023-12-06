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

func CombineNumbers(a, b int) int {
	for tmp := b; tmp > 0; tmp /= 10 {
		a *= 10
	}

	return a + b
}

// Solve each one as a math problem using the quadratic formula
func WaysToWin(time int, distance int) int {
	// We need to beat it, not tie
	d := float64(distance + 1)
	t := float64(time)
	min := int(math.Ceil((t - math.Sqrt(t * t - 4 * d)) / 2))
	max := int(math.Floor((t + math.Sqrt(t * t - 4 * d)) / 2))
	return max - min + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	times := ExtractNumbers(scanner.Text())

	scanner.Scan()
	distances := ExtractNumbers(scanner.Text())

	total := 1
	fullTime := 0
	fullDistance := 0

	for i := 0; i < len(times); i++ {
		total *= WaysToWin(times[i], distances[i])
		fullTime = CombineNumbers(fullTime, times[i])
		fullDistance = CombineNumbers(fullDistance, distances[i])
	}

	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", WaysToWin(fullTime, fullDistance))
}
