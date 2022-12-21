package main

import (
	"fmt"
)

type Element struct {
	value int
	next *Element
	prev *Element
}

func main() {
	originalOrder := make([]*Element, 0)
	var zero *Element
	var prev *Element

	for {
		var num int
		n, _ := fmt.Scanf("%d\n", &num)
		if n != 1 { break }
		elem := new(Element)
		elem.value = num
		if prev == nil {
			elem.next = elem
			elem.prev = elem
		} else {
			elem.next = prev.next
			elem.prev = prev
			elem.prev.next = elem
			elem.next.prev = elem
		}
		prev = elem
		originalOrder = append(originalOrder, elem)
		if num == 0 {
			zero = elem
		}
	}

	for _, elem := range originalOrder {
		/*
		fmt.Println("")
		for i, e := 0, zero; i < len(originalOrder); i, e = i + 1, e.next {
			fmt.Println(e.value)
		}
		*/

		if elem.value > 0 {
			for i := 0; i < elem.value; i++ {
				prev, next, nextNext := elem.prev, elem.next, elem.next.next

				prev.next = next
				next.prev = prev
				next.next = elem
				elem.prev = next
				elem.next = nextNext
				nextNext.prev = elem
			}
		} else if elem.value < 0 {
			for i := 0; i > elem.value; i-- {
				prevPrev, prev, next := elem.prev.prev, elem.prev, elem.next

				prevPrev.next = elem
				elem.prev = prevPrev
				elem.next = prev
				prev.prev = elem
				prev.next = next
				next.prev = prev
			}
		}
	}

	sum := 0

	fmt.Println("")
	for i, e := 1, zero.next; i <= 3000; i, e = i + 1, e.next {
		if i % 1000 == 0 {
			fmt.Println(e.value)
			sum += e.value
		}
	}

	fmt.Println(sum)
}
