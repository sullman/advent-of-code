package main

import (
	"container/list"
	"fmt"
)

type Range struct {
	start int
	end int
}

type Rectangle struct {
	N, E, S, W int
}

func (r *Rectangle) Valid() bool {
	return r.N > r.S && r.E > r.W
}

func Rotate(x int, y int) (int, int) {
	return x + y, y - x
}

func RotateBack(x int, y int) (int, int) {
	return (x - y) / 2, (x + y) / 2
}

// const TargetY = 10
const TargetY = 2000000
const MaxSize = TargetY * 2

func Abs(num int) int {
	if num < 0 {
		return -1 * num
	}

	return num
}

func Min(x int, y int) int {
	if x < y { return x }
	return y
}

func Max(x int, y int) int {
	if x > y { return x }
	return y
}

func AddRange(l *list.List, start int, end int) {
	r := &Range{start: start, end: end}
	for e := l.Front(); e != nil; {
		other := e.Value.(*Range)
		var remove *list.Element

		if r.start <= other.start && r.end >= other.start - 1 {
			remove = e
			r.end = Max(r.end, other.end)
		} else if other.start <= r.start && other.end >= r.start - 1 {
			remove = e
			r.start = Min(r.start, other.start)
			r.end = Max(r.end, other.end)
		}

		e = e.Next()

		if remove != nil {
			l.Remove(remove)
		}
	}

	l.PushBack(r)
}

func SplitRectangles(l *list.List, off *Rectangle) {
	// fmt.Printf("Checking %v against current set of rectangles\n", *off)
	for e := l.Front(); e != nil; {
		rect := e.Value.(*Rectangle)
		overlap := &Rectangle{}
		// fmt.Printf("Comparing to %v\n", *rect)

		overlap.N = Min(rect.N, off.N)
		overlap.E = Min(rect.E, off.E)
		overlap.S = Max(rect.S, off.S)
		overlap.W = Max(rect.W, off.W)
		// fmt.Printf("Overlap is %v\n", *overlap)

		if !overlap.Valid() {
			e = e.Next()
			continue
		}

		// We can create four _potential_ rectangles that result from removing the
		// overlap. For the sake of visualization, imagine a 3x3 square where the
		// overlap is in the middle:
		//
		// LTR
		// LOR
		// LBR
		//
		// So there's a rectangle on the left, the right, and smaller ones on the
		// top and bottom. If the overlap isn't in the middle, some of the
		// rectangles will be invalid, but they're computed the same way.

		left := &Rectangle{N: rect.N, E: overlap.W, S: rect.S, W: rect.W}
		if (left.Valid()) { l.PushFront(left) }
		top := &Rectangle{N: rect.N, E: overlap.E, S: overlap.N, W: overlap.W}
		if (top.Valid()) { l.PushFront(top) }
		right := &Rectangle{N: rect.N, E: rect.E, S: rect.S, W: overlap.E}
		if (right.Valid()) { l.PushFront(right) }
		bottom := &Rectangle{N: overlap.S, E: overlap.E, S: rect.S, W: overlap.W}
		if (bottom.Valid()) { l.PushFront(bottom) }

		remove := e
		e = e.Next()
		l.Remove(remove)
	}

	// fmt.Printf("Now tracking %d rectangles\n", l.Len())
}

func main() {
	beacons := make(map[int]int)
	ranges := list.New()
	candidates := list.New()
	square := Rectangle{}

	// Rotating the rest of the squares is simple, but rotating the initial grid is awkward
	possible := &Rectangle{}
	_, possible.N = Rotate(0, MaxSize)
	possible.E, _ = Rotate(MaxSize, MaxSize)
	_, possible.S = Rotate(MaxSize, 0)
	possible.W, _ = Rotate(0, 0)
	candidates.PushFront(possible)

	for {
		var sensorX, sensorY, beaconX, beaconY int
		n, _ := fmt.Scanf("Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n", &sensorX, &sensorY, &beaconX, &beaconY)
		if n != 4 { break }

		distance := Abs(sensorX - beaconX) + Abs(sensorY - beaconY)
		dx := distance - Abs(sensorY - TargetY)

		if dx > 0 {
			AddRange(ranges, sensorX - dx, sensorX + dx)
		}

		if beaconY == TargetY {
			beacons[beaconX] = beaconY
		}

		// top -> E,N
		// bottom -> W,S
		square.E, square.N = Rotate(sensorX, sensorY + distance)
		square.W, square.S = Rotate(sensorX, sensorY - distance)
		SplitRectangles(candidates, &square)
	}

	impossible := 0
	for e := ranges.Front(); e != nil; e = e.Next() {
		r := e.Value.(*Range)
		impossible += (r.end - r.start + 1)
	}
	impossible -= len(beacons)
	fmt.Println(impossible)

	for e := candidates.Front(); e != nil; e = e.Next() {
		rect := e.Value.(*Rectangle)
		// Checking for a delta of 2 here feels a bit like an off-by-one error
		// we're getting away with?
		if rect.N - 2 == rect.S && rect.E - 2 == rect.W {
			x, y := RotateBack(rect.W + 1, rect.S + 1)
			fmt.Println(x, y)
			fmt.Println(x * 4000000 + y)
		}
		// fmt.Println(e.Value.(*Rectangle))
	}
}
