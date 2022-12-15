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

func main() {
	beacons := make(map[int]int)
	ranges := list.New()

	for {
		var sensorX, sensorY, beaconX, beaconY int
		n, _ := fmt.Scanf("Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n", &sensorX, &sensorY, &beaconX, &beaconY)
		if n != 4 { break }

		distance := Abs(sensorX - beaconX) + Abs(sensorY - beaconY)
		dx := distance - Abs(sensorY - TargetY)

		if dx > 0 {
			r := &Range{start: sensorX - dx, end: sensorX + dx}
			fmt.Printf("Can't be in %d..%d\n", r.start, r.end)
			for e := ranges.Front(); e != nil; {
				other := e.Value.(*Range)
				var remove *list.Element

				if r.start <= other.start && r.end >= other.start {
					remove = e
					r.end = Max(r.end, other.end)
				} else if other.start <= r.start && other.end >= r.start {
					remove = e
					r.start = Min(r.start, other.start)
					r.end = Max(r.end, other.end)
				}

				e = e.Next()

				if remove != nil {
					// fmt.Printf("Combined %d..%d + %d..%d -> %d..%d\n", other.start, other.end, sensorX - dx, sensorX + dx, r.start, r.end)
					ranges.Remove(remove)
				}
			}

			ranges.PushBack(r)
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
}
