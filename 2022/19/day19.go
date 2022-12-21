package main

import (
	"fmt"
)

type Counts map[string]int

type Blueprint struct {
	id int
	costs map[string]Counts
	max Counts
}

type State struct {
	time int
	robots Counts
	supply Counts
	building Counts
	nextBuild string
	blueprint *Blueprint
}

func NewState(other *State) *State {
	state := new(State)
	state.robots = make(Counts)
	state.supply = make(Counts)
	state.building = make(Counts)

	if other != nil {
		state.time = other.time
		state.blueprint = other.blueprint
		for key, val := range other.robots {
			state.robots[key] = val
		}
		for key, val := range other.supply {
			state.supply[key] = val
		}
		for key, val := range other.building {
			state.building[key] = val
		}
	}

	return state
}

func ReadBlueprint() *Blueprint {
	b := new(Blueprint)

	var robotType, costType string
	var cost int
	var sep rune
	n, _ := fmt.Scanf("Blueprint %d:%c", &b.id, &sep)

	if n != 2 { return nil }

	b.costs = make(map[string]Counts)
	b.max = make(Counts)

	for sep == ' ' {
		costs := make(Counts)
		fmt.Scanf("Each %s robot costs", &robotType)
		for more := true; more; {
			fmt.Scanf("%d %s%c", &cost, &costType, &sep)

			if costType[len(costType) - 1] == '.' {
				costType = costType[:len(costType) - 1]
				more = false
			} else {
				var and string
				fmt.Scanf("%s", &and)
			}

			costs[costType] = cost
			if cost > b.max[costType] {
				b.max[costType] = cost
			}
		}

		b.costs[robotType] = costs
	}

	return b
}

var OrderedTypes = []string{"geode", "obsidian", "clay", "ore"}

func NextBuilds(state *State, newStates *[]*State) {
	numBuilds := 0
	for _, buildType := range OrderedTypes {
		costs := state.blueprint.costs[buildType]
		shouldBuildNext := buildType == "geode" || state.robots[buildType] < state.blueprint.max[buildType]
		for resource, _ := range costs {
			if state.robots[resource] + state.building[resource] == 0 {
				shouldBuildNext = false
				break
			}
		}
		if shouldBuildNext {
			newState := NewState(state)
			newState.nextBuild = buildType
			*newStates = append(*newStates, newState)
			numBuilds++
			// Bad heuristic, neutering it
			if numBuilds >= 4 { break }
		}
	}
}

const MaxMinutes = 24

func main() {
	blueprints := make([]*Blueprint, 0)

	for {
		b := ReadBlueprint()
		if b == nil { break }
		blueprints = append(blueprints, b)
	}

	total := 0

	for _, bp := range blueprints {
		fmt.Printf("Running Blueprint %d\n", bp.id)

		initial := NewState(nil)
		initial.blueprint = bp
		initial.robots["ore"] = 1
		states := make([]*State, 0, 200)
		NextBuilds(initial, &states)
		best := initial
		visited := make(map[string]bool)
		lastTime := 0

		for len(states) > 0 {
			state := states[0]
			states = states[1:]

			visitedKey := fmt.Sprintf("%d %s %v %v", state.time, state.nextBuild, state.robots, state.supply)
			if visited[visitedKey] {
				continue
			}
			visited[visitedKey] = true

			if state.time == MaxMinutes {
				if state.supply["geode"] > best.supply["geode"] {
					best = state
				}
				continue
			}

			state.time++

			// fmt.Println(state)
			if state.time > lastTime || state.time > 20 {
				lastTime = state.time
				// fmt.Println(state)
			}

			// Build, if we can
			canBuild := true
			for costType, cost := range bp.costs[state.nextBuild] {
				if cost > state.supply[costType] {
					canBuild = false
					break
				}
			}

			newStates := make([]*State, 0, 4)

			if canBuild {
				state.building[state.nextBuild] += 1
				for costType, cost := range bp.costs[state.nextBuild] {
					state.supply[costType] -= cost
				}

				NextBuilds(state, &newStates)
			} else {
				newStates = append(newStates, state)
			}

			for _, s := range newStates {
				// Harvest
				for key, val := range s.robots {
					s.supply[key] += val
				}

				// Finish building
				for key, val := range s.building {
					s.robots[key] += val
					s.building[key] = 0
				}

				states = append(states, s)
			}
		}

		fmt.Println(best.supply["geode"])

		total += bp.id * best.supply["geode"]
	}

	fmt.Println(total)
}
