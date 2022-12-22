package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	name string
	value int
	known bool
	left *Element
	right *Element
	operator rune
}

func GetValue(elem *Element, part2 bool) (int, bool) {
	if part2 && elem.name == "humn" {
		return 0, false
	} else if elem.known {
		return elem.value, true
	}

	leftVal, leftKnown := GetValue(elem.left, part2)
	rightVal, rightKnown := GetValue(elem.right, part2)

	if !leftKnown || !rightKnown {
		return 0, false
	}

	var value int

	switch elem.operator {
	case '+':
		value = leftVal + rightVal
	case '-':
		value = leftVal - rightVal
	case '*':
		value = leftVal * rightVal
	case '/':
		value = leftVal / rightVal
	default:
		panic("Unknown operator")
	}

	elem.value = value
	elem.known = true

	return value, true
}

func MakeEqual(elem *Element, value int) {
	fmt.Printf("%s needs to equal %d\n", elem.name, value)
	if elem.name == "humn" {
		fmt.Println(value)
		return
	}

	leftVal, leftKnown := GetValue(elem.left, true)
	rightVal, rightKnown := GetValue(elem.right, true)

	if leftKnown {
		var needed int
		switch elem.operator {
		case '+':
			needed = value - leftVal
		case '-':
			needed = leftVal - value
		case '*':
			needed = value / leftVal
		case '/':
			needed = leftVal / value
		default:
			panic("Unknown operator")
		}
		MakeEqual(elem.right, needed)
	} else if rightKnown {
		var needed int
		switch elem.operator {
		case '+':
			needed = value - rightVal
		case '-':
			needed = value + rightVal
		case '*':
			needed = value / rightVal
		case '/':
			needed = value * rightVal
		default:
			panic("Unknown operator")
		}
		MakeEqual(elem.left, needed)
	} else {
		panic("humn is on both sides!?")
	}
}

func main() {
	elements := make(map[string]*Element)

	for {
		var name string
		n, _ := fmt.Scanf("%s", &name)
		if n != 1 { break }
		name = name[0:(len(name) - 1)]

		elem := elements[name]
		if elem == nil {
			elem = &Element{name: name}
			elements[name] = elem
		}

		var left, right string
		var operator rune
		n, _ = fmt.Scanf("%s %c %s\n", &left, &operator, &right)

		if n == 1 {
			elem.value, _ = strconv.Atoi(left)
			elem.known = true
		} else {
			elem.operator = operator

			elem.left = elements[left]
			if elem.left == nil {
				elem.left = &Element{name: left}
				elements[left] = elem.left
			}

			elem.right = elements[right]
			if elem.right == nil {
				elem.right = &Element{name: right}
				elements[right] = elem.right
			}
		}
	}

	root := elements["root"]
	GetValue(root, false)
	fmt.Println(root.value)

	// Reset computed state for part 2
	for _, elem := range elements {
		if elem.left != nil {
			elem.value = 0
			elem.known = false
		}
	}

	leftVal, leftKnown := GetValue(root.left, true)
	rightVal, rightKnown := GetValue(root.right, true)

	if leftKnown {
		MakeEqual(root.right, leftVal)
	} else if rightKnown {
		MakeEqual(root.left, rightVal)
	} else {
		panic("humn is on both sides!?")
	}
}
