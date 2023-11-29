package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	naiveFuel := 0
	fullFuel := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			mass, _ := strconv.Atoi(line)
			fuel := (mass / 3) - 2
			naiveFuel += fuel
			fullFuel += fuel

			for {
				fuel = (fuel / 3) - 2
				if fuel > 0 {
					fullFuel += fuel
				} else {
					break
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", naiveFuel)
	fmt.Printf("Part 2: %d\n", fullFuel)
}
