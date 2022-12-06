package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bytesRead := 0
	numUnique := 14

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line) - numUnique; i++ {
			dup := false
			for j := i; j < i + numUnique - 1; j++ {
				for k := j + 1; k < i + numUnique; k++ {
					if line[j] == line[k] {
						dup = true
						break
					}
				}
			}

			if !dup {
				bytesRead = i + numUnique
				break
			}
		}

		fmt.Println(bytesRead)
	}
}
