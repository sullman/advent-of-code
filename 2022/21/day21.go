package main

import (
	"fmt"
	"strconv"
)

type Formula struct {
	left string
	right string
	operator rune
}

func main() {
	values := make(map[string]int)
	formulas := make(map[string]*Formula)

	for {
		var name string
		n, _ := fmt.Scanf("%s", &name)
		if n != 1 { break }
		name = name[0:(len(name) - 1)]

		var left, right string
		var operator rune
		n, _ = fmt.Scanf("%s %c %s\n", &left, &operator, &right)

		if n == 1 {
			values[name], _ = strconv.Atoi(left)
		} else {
			formula := &Formula{left, right, operator}
			formulas[name] = formula
		}
	}

	fmt.Printf("Parsed %d values, %d formulas\n", len(values), len(formulas))
	fmt.Println(formulas["root"])

	stack := make([]string, 1, len(formulas))
	stack[0] = "root"

	for len(stack) > 0 {
		monkey := stack[len(stack) - 1]
		if _, exists := values[monkey]; exists {
			stack = stack[0:(len(stack) - 1)]
			continue
		}

		formula := formulas[monkey]
		leftVal, leftExists := values[formula.left]
		rightVal, rightExists := values[formula.right]

		if leftExists && rightExists {
			switch formula.operator {
			case '+':
				values[monkey] = leftVal + rightVal
			case '-':
				values[monkey] = leftVal - rightVal
			case '*':
				values[monkey] = leftVal * rightVal
			case '/':
				values[monkey] = leftVal / rightVal
			default:
				panic("Unknown operator")
			}
			stack = stack[0:(len(stack) - 1)]
		} else {
			if !leftExists {
				stack = append(stack, formula.left)
			}
			if !rightExists {
				stack = append(stack, formula.right)
			}
		}
	}

	fmt.Println(values["root"])
}
