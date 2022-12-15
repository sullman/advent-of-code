package main

import (
	"container/list"
	"fmt"
)

type Range struct {
	start int
	end int
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

func main() {
	beacons := make(map[int]int)
	ranges := list.New()
	allRanges := make([]*list.List, MaxSize + 1)

	for i := 0; i <= MaxSize; i++ {
		allRanges[i] = list.New()
	}

	for {
		var sensorX, sensorY, beaconX, beaconY int
		n, _ := fmt.Scanf("Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n", &sensorX, &sensorY, &beaconX, &beaconY)
		if n != 4 { break }

		distance := Abs(sensorX - beaconX) + Abs(sensorY - beaconY)
		minRow := Max(0, sensorY - distance)
		maxRow := Min(MaxSize, sensorY + distance)

		for row := minRow; row <= maxRow; row++ {
			dx := distance - Abs(sensorY - row)
			AddRange(allRanges[row], Max(0, sensorX - dx), Min(MaxSize, sensorX + dx))
			if row == TargetY {
				AddRange(ranges, sensorX - dx, sensorX + dx)
			}
		}

		if beaconY == TargetY {
			beacons[beaconX] = beaconY
		}
	}

	impossible := 0
	for e := ranges.Front(); e != nil; e = e.Next() {
		r := e.Value.(*Range)
		impossible += (r.end - r.start + 1)
	}
	impossible -= len(beacons)
	fmt.Println(impossible)

	for row := 0; row <= MaxSize; row++ {
		if allRanges[row].Len() != 1 {
			// We're not sorting the ranges, so have to look at both and take the min
			col := 1 + Min(allRanges[row].Front().Value.(*Range).end, allRanges[row].Back().Value.(*Range).end)
			fmt.Println(row + 4000000 * col)
		}
	}
}
