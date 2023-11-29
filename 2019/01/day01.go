package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			mass, _ := strconv.Atoi(line)
			fuel := (mass / 3) - 2
			sum += fuel
		}
	}

	fmt.Println(sum)
}
