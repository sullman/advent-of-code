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
	orig [2]Point
	id int
	supporting map[int]bool
	supportedBy map[int]bool
}

var nextId int
func NewBlock(p1 Point, p2 Point) *Block {
	block := new(Block)

	if p1.x != p2.x {
		if p1.x > p2.x { p1.x, p2.x = p2.x, p1.x }
	} else if p1.y != p2.y {
		if p1.y > p2.y { p1.y, p2.y = p2.y, p1.y }
	} else {
		if p1.z > p2.z { p1.z, p2.z = p2.z, p1.z }
	}

	block.start, block.end = p1, p2
	block.orig[0], block.orig[1] = p1, p2
	block.id = nextId
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
		if block.end.z > max.z { max.z = block.end.z }

		blocks = append(blocks, block)
	}

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].start.z < blocks[j].start.z
	})

	ground := NewBlock(Point{0, 0, 0}, Point{max.x, max.y, 0})

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
	byId := make(map[int]*Block)
	for _, block := range blocks {
		block.orig[0], block.orig[1] = block.start, block.end
		byId[block.id] = block
		if len(block.supportedBy) == 1 {
			for id, _ := range block.supportedBy {
				if id == ground.id { continue }
				unsafe[id] = true
			}
		}
	}

	fmt.Printf("Part 1: %d\n", len(blocks) - len(unsafe))

	// Part 2
	total := 0
	for _, initial := range blocks {
		space := make([][][]*Block, max.x + 1)

		for x := 0; x < len(space); x++ {
			space[x] = make([][]*Block, max.y + 1)
			for y := 0; y < len(space[0]); y++ {
				space[x][y] = make([]*Block, max.z + 1)
				space[x][y][0] = ground
			}
		}

		for _, block := range blocks {
			block.start, block.end = block.orig[0], block.orig[1]
			if block.id != initial.id {
				for x := block.start.x; x <= block.end.x; x++ {
					for y := block.start.y; y <= block.end.y; y++ {
						for z := block.start.z; z <= block.end.z; z++ {
							space[x][y][z] = block
						}
					}
				}
			}
		}

		moved := make(map[int]bool)
		queue := make([]*Block, 0)
		for id, _ := range initial.supporting {
			queue = append(queue, byId[id])
		}
		for len(queue) != 0 {
			block := queue[0]
			queue = queue[1:]

			canFall, fell := true, false
			for canFall {
				for x := block.start.x; x <= block.end.x; x++ {
					for y := block.start.y; y <= block.end.y; y++ {
						if space[x][y][block.start.z - 1] != nil {
							canFall = false
						}
					}
				}

				if canFall {
					fell = true
					for x := block.start.x; x <= block.end.x; x++ {
						for y := block.start.y; y <= block.end.y; y++ {
							for z := block.start.z; z <= block.end.z; z++ {
								space[x][y][z - 1], space[x][y][z] = block, nil
							}
						}
					}
					block.start.z--
					block.end.z--
				}
			}
			if fell {
				moved[block.id] = true
				for id, _ := range block.supporting {
					queue = append(queue, byId[id])
				}
			}
		}

		total += len(moved)
	}

	fmt.Printf("Part 2: %d\n", total)
}
