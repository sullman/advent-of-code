package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	num int
	sum int
}

func main() {
	var elves []Elf
	var max Elf

	scanner := bufio.NewScanner(os.Stdin)
	num := 0
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			calories, _ := strconv.Atoi(line)
			num++
			sum += calories
		} else {
			elf := Elf{num, sum}
			elves = append(elves, elf)
			if sum > max.sum {
				max = elf
			}

			sum = 0
			num = 0

			// fmt.Printf("Added elf with %d item(s) totaling %d calories\n", elf.num, elf.sum)
		}
	}

	if sum > 0 {
		elves = append(elves, Elf{num, sum})
	}

	fmt.Println("\nTop Elves:")

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].sum > elves[j].sum
	})

	total := 0
	for i := 0; i < 3; i++ {
		total += elves[i].sum
		fmt.Printf("- %d calories (%d item(s))\n", elves[i].sum, elves[i].num)
	}

	fmt.Println(total)
}
