package main

import (
	"fmt"
	"sort"
)

type Point struct {
	x int
	y int
	z int
}

type Block struct {
	start Point
	end Point
	id int
	axis int
	supporting map[int]bool
	supportedBy map[int]bool
}

var nextId int
func NewBlock(p1 Point, p2 Point) *Block {
	var axis int
	block := new(Block)

	if p1.x != p2.x {
		axis = 0
		if p1.x > p2.x { p1.x, p2.x = p2.x, p1.x }
	} else if p1.y != p2.y {
		axis = 1
		if p1.y > p2.y { p1.y, p2.y = p2.y, p1.y }
	} else {
		axis = 2
		if p1.z > p2.z { p1.z, p2.z = p2.z, p1.z }
	}

	block.start, block.end = p1, p2
	block.id = nextId
	block.axis = axis
	block.supporting = make(map[int]bool)
	block.supportedBy = make(map[int]bool)
	nextId++

	return block
}

func main() {
	blocks := make([]*Block, 0)
	var x1, y1, z1, x2, y2, z2, n int
	var max Point

	for {
		n, _ = fmt.Scanf("%d,%d,%d~%d,%d,%d", &x1, &y1, &z1, &x2, &y2, &z2)
		if n != 6 { break }

		block := NewBlock(Point{x1, y1, z1}, Point{x2, y2, z2})

		if block.end.x > max.x { max.x = block.end.x }
		if block.end.y > max.y { max.y = block.end.y }

		blocks = append(blocks, block)
	}

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].start.z < blocks[j].start.z
	})

	ground := NewBlock(Point{0, 0, 0}, max)

	settled := make([][]*Block, max.x + 1)
	for x := 0; x < len(settled); x++ {
		settled[x] = make([]*Block, max.y + 1)
		for y := 0; y < len(settled[x]); y++ {
			settled[x][y] = ground
		}
	}

	for _, block := range blocks {
		canFall := true
		for canFall {
			for x := block.start.x; x <= block.end.x; x++ {
				for y := block.start.y; y <= block.end.y; y++ {
					if settled[x][y].end.z == block.start.z - 1 {
						canFall = false
						settled[x][y].supporting[block.id] = true
						block.supportedBy[settled[x][y].id] = true
					}
				}
			}
			if canFall {
				block.start.z--
				block.end.z--
			}
		}
		for x := block.start.x; x <= block.end.x; x++ {
			for y := block.start.y; y <= block.end.y; y++ {
				settled[x][y] = block
			}
		}
	}

	unsafe := make(map[int]bool)
	for _, block := range blocks {
		if len(block.supportedBy) == 1 {
			for id, _ := range block.supportedBy {
				if id == ground.id { continue }
				unsafe[id] = true
			}
		}
	}

	fmt.Printf("Part 1: %d\n", len(blocks) - len(unsafe))
}
