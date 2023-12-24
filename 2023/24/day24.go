package main

import (
	"fmt"
)

type Triple struct {
	x int
	y int
	z int
}

type Hailstone struct {
	point Triple
	velocity Triple
	slope float64
	intercept float64
}

type Range struct {
	min float64
	max float64
}

func (r *Range) Contains(val float64) bool {
	return val >= r.min && val <= r.max
}

func main() {
	hailstones := make([]Hailstone, 0, 128)
	// window := &Range{7, 27}
	window := &Range{200000000000000, 400000000000000}
	epsilon := &Range{-0.00001, 0.00001}

	for {
		var point, velocity Triple
		var slope, intercept float64
		n, _ := fmt.Scanf("%d, %d, %d @ %d, %d, %d", &point.x, &point.y, &point.z, &velocity.x, &velocity.y, &velocity.z)
		if n != 6 { break }
		slope = float64(velocity.y) / float64(velocity.x)
		intercept = float64(point.y) - float64(point.x) * slope
		hailstones = append(hailstones, Hailstone{point, velocity, slope, intercept})
	}

	total := 0

	for i := 1; i < len(hailstones); i++ {
		hail := hailstones[i]
		for j := 0; j < i; j++ {
			other := hailstones[j]
			intercepts := hail.intercept - other.intercept
			slopes := other.slope - hail.slope
			if epsilon.Contains(slopes) {
				continue
			}
			x := intercepts / slopes
			if !window.Contains(x) { continue }
			y := hail.slope * x + hail.intercept
			if !window.Contains(y) { continue }
			if (hail.velocity.x > 0 && x < float64(hail.point.x)) { continue }
			if (hail.velocity.x < 0 && x > float64(hail.point.x)) { continue }
			if (other.velocity.x > 0 && x < float64(other.point.x)) { continue }
			if (other.velocity.x < 0 && x > float64(other.point.x)) { continue }
			total++
		}
	}

	fmt.Printf("Part 1: %d\n", total)
}
