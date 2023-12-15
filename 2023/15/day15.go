package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Lens struct {
	Label uint64
	FocalLength byte
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var label uint64
	sum := 0
	value := 0
	hash := -1

	var boxes [256]*list.List

	for i := 0; i < len(boxes); i++ {
		boxes[i] = list.New()
	}

	for {
		ch, err := reader.ReadByte()
		if err != nil { break }
		if ch == '\n' { continue }

		if ch == ',' {
			sum += value
			label, value, hash = 0, 0, -1
			continue
		}

		if ch == '-' {
			hash = value
			for e := boxes[hash].Front(); e != nil; e = e.Next() {
				if e.Value.(*Lens).Label == label {
					boxes[hash].Remove(e)
					break
				}
			}
		} else if ch == '=' {
			hash = value
		} else if hash != -1 {
			focalLength := ch - '0'
			replaced := false
			for e := boxes[hash].Front(); e != nil; e = e.Next() {
				if e.Value.(*Lens).Label == label {
					e.Value.(*Lens).FocalLength = focalLength
					replaced = true
					break
				}
			}
			if !replaced {
				boxes[hash].PushBack(&Lens{label, focalLength})
			}
		} else {
			label = (label << 8) | uint64(ch)
		}

		value = ((value + int(ch)) * 17) & 0xff
	}

	sum += value

	fmt.Printf("Part 1: %d\n", sum)

	sum = 0

	for i, lst := range boxes {
		for slot, e := 1, lst.Front(); e != nil; slot, e = slot + 1, e.Next() {
			sum += (i + 1) * slot * int(e.Value.(*Lens).FocalLength)
		}
	}

	fmt.Printf("Part 2: %d\n", sum)
}
