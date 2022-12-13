package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	EOL = -3
	OpenList = -2
	CloseList = -1
)

type Packet struct {
	str string
	index int
	bufferedVal int
	bufferedParens int
	special bool
}

func NewPacket(str string) *Packet {
	p := &Packet{str: str}
	p.Reset()
	return p
}

func (p *Packet) Reset() {
	p.index = 0
	p.bufferedVal = -1
	p.bufferedParens = 0
}

func (p *Packet) Next() int {
	if p.bufferedVal >= 0 {
		val := p.bufferedVal
		p.bufferedVal = -1
		return val
	} else if p.bufferedParens > 0 {
		p.bufferedParens--
		return CloseList
	} else if p.index >= len(p.str) {
		return EOL
	}

	if p.str[p.index] == '[' {
		p.index++
		return OpenList
	} else if p.str[p.index] == ']' {
		p.index++
		return CloseList
	}

	val := 0
	for ; p.str[p.index] >= '0' && p.str[p.index] <= '9'; p.index++ {
		val = 10 * val + int(p.str[p.index] - '0')
	}

	if p.str[p.index] == ',' {
		p.index++
	}

	return val
}

func (p *Packet) SynthesizeList(val int) {
	p.bufferedVal = val
	p.bufferedParens++
}

func Compare(left *Packet, right *Packet) bool {
	left.Reset()
	right.Reset()

	for l, r := left.Next(), right.Next(); l != EOL && r != EOL; l, r = left.Next(), right.Next() {
		if l == r {
			// Doesn't matter what they are if they're the same, just keep going
			continue
		} else if r == CloseList {
			return false
		} else if l == CloseList {
			return true
		} else if l == OpenList {
			right.SynthesizeList(r)
		} else if r == OpenList {
			left.SynthesizeList(l)
		} else if l < r {
			return true
		} else { // l > r
			return false
		}
	}

	panic("Packets were the same?")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	index := 0
	sum := 0
	packets := make([]*Packet, 2)
	packets[0] = NewPacket("[[2]]")
	packets[0].special = true
	packets[1] = NewPacket("[[6]]")
	packets[1].special = true

	for scanner.Scan() {
		index++
		left := NewPacket(scanner.Text())
		packets = append(packets, left)
		scanner.Scan()
		right := NewPacket(scanner.Text())
		packets = append(packets, right)
		scanner.Scan()

		if Compare(left, right) {
			sum += index
		}
	}

	fmt.Println(sum)

	sort.Slice(packets, func(i, j int) bool {
		return Compare(packets[i], packets[j])
	})

	product := 1
	for i, p := range packets {
		if p.special {
			product *= i + 1
		}
	}

	fmt.Println(product)
}
