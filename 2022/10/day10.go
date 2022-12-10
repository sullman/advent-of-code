package main

import (
	"bufio"
	"fmt"
	"os"
)

func tick(register *int, cycle *int) int {
	*cycle++

	if *cycle % 20 == 0 && (*cycle / 20) % 2 == 1 {
		return *cycle * *register
	}

	return 0
}

func pixel(register *int, cycle *int) rune {
	drawPos := (*cycle - 1) % 40
	delta := *register - drawPos
	r := '.'

	if delta <= 1 && delta >= -1 {
		r = '#'
	}

	// fmt.Printf("During cycle %d, register is %d, delta is %d. Drawing %c at %d\n", *cycle, *register, delta, r, drawPos)

	return r
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	register := 1
	cycle := 0
	total := 0
	var pixels [240]rune

	for scanner.Scan() {
		line := scanner.Text()
		if line == "noop" {
			total += tick(&register, &cycle)
			pixels[cycle - 1] = pixel(&register, &cycle)
		} else {
			var addend int
			fmt.Sscanf(line, "addx %d", &addend)
			total += tick(&register, &cycle)
			pixels[cycle - 1] = pixel(&register, &cycle)
			total += tick(&register, &cycle)
			pixels[cycle - 1] = pixel(&register, &cycle)
			register += addend
		}
	}

	fmt.Println(total)
	for i := 0; i < 6; i++ {
		fmt.Println(string(pixels[i * 40:(i + 1) * 40]))
	}
}
