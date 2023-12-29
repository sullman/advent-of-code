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

func InferStartingPosition(x1, y1, vx1, vy1, x2, y2, vx2, vy2, vx, vy int) (int, int) {
	if vx == vx1 || vx == vx2 || vy == vy1 || vy == vy2 { return 0, 0 }
	a := float64(vy1 - vy)
	b := float64(vx1 - vx)
	c := float64(vx2 - vx)
	d := float64(vy2 - vy)
	e := b * d / a - c
	if e == 0 { return 0, 0 }

	t2 := (b * float64(y1) / a - b * float64(y2) / a + float64(x2) - float64(x1)) / e
	t1 := (float64(x2) + t2 * c - float64(x1)) / b

	if t1 < 0 || t2 < 0 { return 0, 0 }

	x := int(float64(x1) + t1 * float64(vx1) - t1 * float64(vx) + 0.5)
	y := int(float64(y1) + t1 * float64(vy1) - t1 * float64(vy) + 0.5)

	return x, y
}

func Part2(hailstones []Hailstone) int {
	var delta, start, velocity Triple
	var h1, h2 *Hailstone
	delta.x = 1
	spiral := 1

	outer:
	for spiral <= 1000 {
		for 2 * velocity.x * delta.x < spiral && 2 * velocity.y * delta.y < spiral {
			if velocity.x != 0 && velocity.y != 0 {
				h1, h2 = nil, nil
				var i int
				for i < len(hailstones) {
					h2 = nil
					for ; h2 == nil; i++ {
						if hailstones[i].velocity.x == velocity.x || hailstones[i].velocity.y == velocity.y { continue }
						if h1 == nil {
							h1 = &(hailstones[i])
						} else {
							h2 = &(hailstones[i])
						}
					}

					start.x, start.y = InferStartingPosition(h1.point.x, h1.point.y, h1.velocity.x, h1.velocity.y, h2.point.x, h2.point.y, h2.velocity.x, h2.velocity.y, velocity.x, velocity.y)
					if start.x != 0 { break }
				}

				inner:
				for i < len(hailstones) {
					h2 = nil
					for ; h2 == nil && i < len(hailstones); i++ {
						if hailstones[i].velocity.x == velocity.x || hailstones[i].velocity.y == velocity.y { continue }
						h2 = &(hailstones[i])
					}
					if h2 != nil {
						x, y := InferStartingPosition(h1.point.x, h1.point.y, h1.velocity.x, h1.velocity.y, h2.point.x, h2.point.y, h2.velocity.x, h2.velocity.y, velocity.x, velocity.y)
						if x == start.x && y == start.y {
							fmt.Printf("Success?\n")
							break outer
						} else if x != 0 {
							// fmt.Printf("For %d,%d start=%d,%d does not equal check %d,%d\n", velocity.x, velocity.y, start.x, start.y, x, y)
							break inner
						}
					}
				}
			}
			velocity.x += delta.x
			velocity.y += delta.y
		}
		if delta.x == 0 {
			delta.x, delta.y = delta.y * -1, 0
			spiral++
		} else {
			delta.y, delta.x = delta.x, 0
		}
	}

	fmt.Printf("Solved for x,y=%d,%d\n", start.x, start.y)
	t1 := (start.x - h1.point.x) / (h1.velocity.x - velocity.x)
	t2 := (start.x - h2.point.x) / (h2.velocity.x - velocity.x)
	velocity.z = (h1.point.z - h2.point.z + t1 * h1.velocity.z - t2 * h2.velocity.z) / (t1 - t2)
	start.z = h1.point.z + t1 * h1.velocity.z - t1 * velocity.z

	fmt.Printf("Think the answer is %v @ %v\n", start, velocity)

	return start.x + start.y + start.z
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

	fmt.Printf("Part 2: %d\n", Part2(hailstones))
}
