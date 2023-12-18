package main

import (
	"fmt"
)

type Lagoon struct {
	x int
	y int
	area int
	perimeter int
}

func (l *Lagoon) Dig(dir rune, length int) {
	l.perimeter += length

	switch dir {
	case 'R', '0':
		l.x += length
		l.area -= l.y * length
	case 'D', '1':
		l.y -= length
		l.area -= l.x * length
	case 'L', '2':
		l.x -= length
		l.area += l.y * length
	case 'U', '3':
		l.y += length
		l.area += l.x * length
	}
}

func (l *Lagoon) Size() int {
	total := l.perimeter
	if l.area > 0 {
		total += l.area
	} else {
		total -= l.area
	}

	return (total >> 1) + 1
}

func main() {
	var dir rune
	var n, length int

	lagoon1, lagoon2 := new(Lagoon), new(Lagoon)

	for {
		n, _ = fmt.Scanf("%c %d", &dir, &length)
		if n != 2 { break }

		lagoon1.Dig(dir, length)

		n, _ = fmt.Scanf("(#%5x%c)\n", &length, &dir)
		if n != 2 { panic("WTF") }

		lagoon2.Dig(dir, length)
	}

	if lagoon1.x != 0 || lagoon1.y != 0 {
		panic("Part 1: Not a closed loop?")
	}
	if lagoon2.x != 0 || lagoon2.y != 0 {
		panic("Part 2: Not a closed loop?")
	}

	fmt.Printf("Part 1: %d\n", lagoon1.Size())
	fmt.Printf("Part 2: %d\n", lagoon2.Size())
}
