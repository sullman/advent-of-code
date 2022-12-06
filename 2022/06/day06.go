package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var rollingCounts [26]int
	windowSize := 4

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			rollingCounts[line[i] - 'a']++
			if i >= windowSize {
				rollingCounts[line[i - windowSize] - 'a']--
				dup := rollingCounts[line[i - windowSize] - 'a'] != 1 || rollingCounts[line[i] - 'a'] != 1
				for j := 0; !dup && j < len(rollingCounts); j++ {
					if rollingCounts[j] > 1 {
						dup = true
					}
				}

				if !dup {
					fmt.Println(i + 1)
					break
				}
			}
		}
	}
}
