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

var Primes = []int{2, 3, 5, 7, 11, 13, 17, 23, 29}

func ZeroOut(row1, row2 []int, col int) []int {
	rows := make([][]int, 2)
	rows[0] = make([]int, len(row1))
	copy(rows[0], row1)
	rows[1] = make([]int, len(row2))
	copy(rows[1], row2)

	for _, r := range rows {
		if r[col] < 0 {
			for c := col; c < len(r); c++ {
				r[c] *= -1
			}
		}
	}

	if rows[1][col] > rows[0][col] {
		rows[0], rows[1] = rows[1], rows[0]
	}

	for rows[1][col] != 0 {
		row := make([]int, len(row1))
		for c := col; c < len(row); c++ {
			// TODO: Can't guarantee this won't overflow int64 :/
			row[c] = rows[0][c] - rows[1][c]
		}
		for i := 0; i < len(Primes); i++ {
			p := Primes[i]
			divis := true
			for c := col; c < len(row) && divis; c++ {
				if row[c] % p != 0 { divis = false }
			}
			if divis {
				fmt.Printf("Reducing by /%d\n", p)
				for c := col; c < len(row); c++ {
					row[c] /= p
				}
				i--
			}
		}

		if row[col] > rows[1][col] {
			rows[0] = row
		} else {
			rows[0], rows[1] = rows[1], row
		}
	}

	return rows[1]
}

func SolveLinearEquations(values [][]int) []int {
	solution := make([]int, len(values))
	var row int

	for col := 0; col < len(solution); col++ {
		// First, make sure the first row is nonzero
		for row = col; row < len(values); row++ {
			if values[row][col] != 0 { break }
		}
		if row != col {
			values[col], values[row] = values[row], values[col]
		}

		// Now zero out all but the first
		for row = col + 1; row < len(values); row++ {
			values[row] = ZeroOut(values[col], values[row], col)
		}
	}

	lastIndex := len(values[0]) - 1
	for col := len(solution) - 1; col >= 0; col-- {
		val := values[col][lastIndex]
		for c := col + 1; c < lastIndex; c++ {
			val -= values[col][c] * solution[c]
		}
		solution[col] = val / values[col][col]
	}

	return solution
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

	// The underlying algebra is worked out elsewhere, but the general idea is
	// that for our `rock` and any `hail`:
	//
	// rock.point + t * rock.velocity = hail.point + t * hail.velocity
	// (rearranges to)
	// (rock.point - hail.point) = -t * (rock.velocity - hail.velocity)
	//
	// You can ignore the actual value of t and consider these parallel vectors,
	// which means their cross product must be zero. Those cross products are
	// where we derive the coefficients for our linear equations. 6 equations
	// for 6 unknowns if we use three hailstones.
	equations := make([][]int, 6)
	h1, h2, h3 := hailstones[0], hailstones[1], hailstones[2]
	// [x, y, z, vx, vy, vz, c]
	equations[0] = []int{
		h2.velocity.y - h1.velocity.y,
		h1.velocity.x - h2.velocity.x,
		0,
		h1.point.y - h2.point.y,
		h2.point.x - h1.point.x,
		0,
		h2.point.x * h2.velocity.y - h1.point.x * h1.velocity.y + h1.point.y * h1.velocity.x - h2.point.y * h2.velocity.x,
	}

	equations[1] = []int {
		h3.velocity.y - h1.velocity.y,
		h1.velocity.x - h3.velocity.x,
		0,
		h1.point.y - h3.point.y,
		h3.point.x - h1.point.x,
		0,
		h3.point.x * h3.velocity.y - h1.point.x * h1.velocity.y + h1.point.y * h1.velocity.x - h3.point.y * h3.velocity.x,
	}

	equations[2] = []int {
		h2.velocity.z - h1.velocity.z,
		0,
		h1.velocity.x - h2.velocity.x,
		h1.point.z - h2.point.z,
		0,
		h2.point.x - h1.point.x,
		h2.point.x * h2.velocity.z - h1.point.x * h1.velocity.z + h1.point.z * h1.velocity.x - h2.point.z * h2.velocity.x,
	}

	equations[3] = []int {
		h3.velocity.z - h1.velocity.z,
		0,
		h1.velocity.x - h3.velocity.x,
		h1.point.z - h3.point.z,
		0,
		h3.point.x - h1.point.x,
		h3.point.x * h3.velocity.z - h1.point.x * h1.velocity.z + h1.point.z * h1.velocity.x - h3.point.z * h3.velocity.x,
	}

	equations[4] = []int {
		0,
		h2.velocity.z - h1.velocity.z,
		h1.velocity.y - h2.velocity.y,
		0,
		h1.point.z - h2.point.z,
		h2.point.y - h1.point.y,
		h2.point.y * h2.velocity.z - h1.point.y * h1.velocity.z + h1.point.z * h1.velocity.y - h2.point.z * h2.velocity.y,
	}

	equations[5] = []int {
		0,
		h3.velocity.z - h1.velocity.z,
		h1.velocity.y - h3.velocity.y,
		0,
		h1.point.z - h3.point.z,
		h3.point.y - h1.point.y,
		h3.point.y * h3.velocity.z - h1.point.y * h1.velocity.z + h1.point.z * h1.velocity.y - h3.point.z * h3.velocity.y,
	}

	solution := SolveLinearEquations(equations)
	fmt.Printf("Solution: %d, %d, %d @ %d, %d, %d\n", solution[0], solution[1], solution[2], solution[3], solution[4], solution[5])
	fmt.Printf("Part 2: %d\n", solution[0] + solution[1] + solution[2])
}
