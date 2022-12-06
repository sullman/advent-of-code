package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bytesRead := 0

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line) - 4; i++ {
			dup := false
			for j := i; j < i + 3; j++ {
				for k := j + 1; k < i + 4; k++ {
					if line[j] == line[k] {
						dup = true
						break
					}
				}
			}

			if !dup {
				bytesRead = i + 4
				break
			}
		}

		fmt.Println(bytesRead)
	}
}
