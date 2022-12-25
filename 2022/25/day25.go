package main

import (
	"fmt"
)

func ConvertToSnafu(num int) string {
	s := ""
	carry := 0

	for num > 0 {
		switch num % 5 {
		case 0:
			s = "0" + s
			carry = 0
		case 1:
			s = "1" + s
			carry = 0
		case 2:
			s = "2" + s
			carry = 0
		case 3:
			s = "=" + s
			carry = 1
		case 4:
			s = "-" + s
			carry = 1
		}

		num = num / 5 + carry
	}

	return s
}

func main() {
	counts := make([]int, 0, 20)
	var decode = map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'-': -1,
		'=': -2,
	}

	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil { break }

		for i, j := 0, len(line) - 1; i < len(line); i, j = i + 1, j - 1 {
			if len(counts) == i {
				counts = append(counts, 0)
			}
			counts[i] += decode[line[j]]
		}
	}

	decimal := 0
	power := 1

	for _, count := range counts {
		decimal += count * power
		power *= 5
	}

	fmt.Println(counts)
	fmt.Println(decimal)
	fmt.Println(ConvertToSnafu(decimal))
}
