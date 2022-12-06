package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total1 := 0
	total2 := 0

	var start1, end1, start2, end2 int

	for scanner.Scan() {
		n, _ := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &start1, &end1, &start2, &end2)

		if n != 4 { continue }

		if (start1 >= start2 && end1 <= end2) || (start2 >= start1 && end2 <= end1) {
			total1++
			total2++
		} else if (end1 >= start2 && end1 <= end2) || (start1 >= start2 && start1 <= end2) {
			total2++
		}
	}

	fmt.Println(total1)
	fmt.Println(total2)
}
