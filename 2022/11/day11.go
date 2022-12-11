package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Operator byte

const (
	Undefined Operator = iota
	Add
	Multiply
)

const NumRounds = 10000
const WorryFactor = 1

type Monkey struct {
	items []int
	op Operator
	operand int
	test int
	next [2]int
	numInspected int
}

func main() {
	monkeys := make([]*Monkey, 0, 3)
	lcm := 1

	for {
		monkey := new(Monkey)
		var index int
		var operator, operand string
		n, err := fmt.Scanf("Monkey %d:\n", &index)

		if n != 1 || err != nil {
			break
		}

		if index != len(monkeys) {
			panic("Unexpected monkey index")
		}

		// Reading the whole line and using strings.Split or a regex would probably be easier :shrug:
		fmt.Scanf("  Starting items:")
		monkey.items = make([]int, 0, 10)
		for {
			var item int
			var sep rune
			fmt.Scanf(" %d%c", &item, &sep)
			monkey.items = append(monkey.items, item)
			if sep != ',' {
				break
			}
		}

		fmt.Scanf("  Operation: new = old %s %s\n", &operator, &operand)
		fmt.Scanf("  Test: divisible by %d\n", &monkey.test)
		fmt.Scanf("    If true: throw to monkey %d\n", &monkey.next[0])
		fmt.Scanf("    If false: throw to monkey %d\n", &monkey.next[1])

		switch operator {
		case "*":
			monkey.op = Multiply
		case "+":
			monkey.op = Add
		default:
			panic("Unexpected operator")
		}

		if operand == "old" {
			monkey.operand = -1
		} else {
			monkey.operand, _ = strconv.Atoi(operand)
		}

		// We're calling this the LCM, but we're not actually doing any work to
		// make it the *least* common multiple. It's just *a* multiple. The input
		// uses primes, so it turns out to be right.
		lcm = lcm * monkey.test
		monkeys = append(monkeys, monkey)
		fmt.Scanln()
	}

	for round := 0; round < NumRounds; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				monkey.numInspected++
				newValue := item
				if monkey.op == Add {
					if monkey.operand == -1 {
						newValue += item
					} else {
						newValue += monkey.operand
					}
				} else {
					if monkey.operand == -1 {
						newValue *= item
					} else {
						newValue *= monkey.operand
					}
				}

				if newValue < item {
					panic("Overflow!")
				}

				newValue = (newValue / WorryFactor) % lcm
				next := monkey.next[1]
				if newValue % monkey.test == 0 {
					next = monkey.next[0]
				}

				monkeys[next].items = append(monkeys[next].items, newValue)
			}

			// We could also set it to nil, but since they're just ints and we don't
			// care about garbage collection, preserving the underlying capacity
			// should be slightly more efficient.
			monkey.items = monkey.items[:0]
		}
	}

	// This sort is destructive, but fine since we're done now
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].numInspected > monkeys[j].numInspected
	})

	fmt.Println(monkeys[0].numInspected * monkeys[1].numInspected)
}
