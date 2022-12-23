package main

import (
	"fmt"
	"sort"
)

type Valve struct {
	name string
	flowRate int
}

type State struct {
	time int
	target string
	distance int
	released int
	releasing int
	remaining []*Valve
	bestPossible int
}

func ComputeBestPossible(state *State, valves *map[string]*Valve, maxTime int) int {
	timeRemaining := maxTime - state.time + 1
	possible := state.released + state.releasing * timeRemaining

	target := (*valves)[state.target]
	if target != nil && state.distance < timeRemaining {
		possible += target.flowRate * (timeRemaining - state.distance)
	}

	for _, v := range state.remaining {
		if timeRemaining <= 0 { break }
		possible += v.flowRate * timeRemaining
		timeRemaining--
	}

	return possible
}

func Part1(valves *map[string]*Valve, distances *map[string]map[string]int) int {
	relevantValves := make([]*Valve, 0)
	for _, v := range *valves {
		if v.flowRate > 0 {
			relevantValves = append(relevantValves, v)
		}
	}
	sort.Slice(relevantValves, func(i, j int) bool {
		return relevantValves[i].flowRate > relevantValves[j].flowRate
	})

	best := 0
	stack := make([]*State, 1)
	stack[0] = &State{
		time: 0,
		target: "AA",
		distance: 0,
		released: 0,
		releasing: 0,
		remaining: relevantValves,
		bestPossible: 1000000,
	}

	for len(stack) > 0 {
		// Using a stack makes this depth first, which will hopefully let us find a
		// plausible solution quickly and then aggressively prune anything that can't
		// top it.
		state := stack[len(stack) - 1]
		stack = stack[0:(len(stack) - 1)]

		// fmt.Printf("Evaluating t=%d target=%s(%d) released=%d releasing=%d remaining=%v best=%d\n", state.time, state.target, state.distance, state.released, state.releasing, state.remaining, state.bestPossible)

		if state.time == 31 {
			if state.released > best {
				fmt.Printf("New best: %d\n", state.released)
				best = state.released
			}
			continue
		} else if state.bestPossible < best {
			continue
		}

		// Track visited states?

		state.time++
		state.released += state.releasing

		if state.distance > 0 {
			state.distance--
			state.bestPossible = ComputeBestPossible(state, valves, 30)
			stack = append(stack, state)
		} else {
			state.releasing += (*valves)[state.target].flowRate
			// fmt.Printf("Opening %s, now releasing %d\n", state.target, state.releasing)

			if len(state.remaining) == 0 {
				state.released += state.releasing * (30 - state.time + 1)
				state.time = 31
				stack = append(stack, state)
			} else {
				for _, next := range state.remaining {
					remaining := make([]*Valve, 0, len(state.remaining) - 1)
					for _, v := range state.remaining {
						if v.name != next.name {
							remaining = append(remaining, v)
						}
					}
					s := &State{
						time: state.time,
						target: next.name,
						distance: (*distances)[state.target][next.name],
						released: state.released,
						releasing: state.releasing,
						remaining: remaining,
					}
					s.bestPossible = ComputeBestPossible(s, valves, 30)
					stack = append(stack, s)
				}
			}
		}
	}

	return best
}

func main() {
	valves := make(map[string]*Valve)
	distances := make(map[string]map[string]int)

	for {
		var name string
		var rate int
		var skip1, skip2, skip3 string
		var next string
		sep := ','

		n, _ := fmt.Scanf("Valve %s has flow rate=%d; %s %s to %s", &name, &rate, &skip1, &skip2, &skip3)
		if n != 5 { break }

		valve := &Valve{
			name: name,
			flowRate: rate,
		}
		valves[name] = valve

		distances[name] = make(map[string]int)

		for sep == ',' {
			fmt.Scanf("%2s%c", &next, &sep)
			distances[name][next] = 1
		}
	}

	for through, _ := range valves {
		for from, _ := range valves {
			for to, _ := range valves {
				before, after := distances[from][through], distances[through][to]
				if before != 0 && after != 0 {
					dist := before + after
					if distances[from][to] == 0 || dist < distances[from][to] {
						distances[from][to] = dist
					}
				}
			}
		}
	}

	/*
	for from, _ := range valves {
		fmt.Printf("Distances from %s: %v\n", from, distances[from])
	}
	*/

	fmt.Println(Part1(&valves, &distances))
}
