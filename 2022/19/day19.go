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
	bestPossible int
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
		state.bestPossible = other.bestPossible
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

var OrderedTypes = []string{"ore", "clay", "obsidian", "geode"}

func NextBuilds(state *State, newStates *[]*State) {
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
		}
	}
}

// const MaxMinutes = 24
const MaxMinutes = 32

func main() {
	blueprints := make([]*Blueprint, 0)

	for {
		b := ReadBlueprint()
		if b == nil { break }
		blueprints = append(blueprints, b)
	}

	total := 0
	product := 1

	for _, bp := range blueprints {
		if MaxMinutes == 32 && bp.id > 3 { break }
		fmt.Printf("Running Blueprint %d\n", bp.id)

		initial := NewState(nil)
		initial.blueprint = bp
		initial.robots["ore"] = 1
		initial.bestPossible = MaxMinutes * (MaxMinutes + 1) / 2
		states := make([]*State, 0, 200)
		NextBuilds(initial, &states)
		best := initial
		visited := make(map[string]bool)

		for len(states) > 0 {
			// Pop from the end of the stack to make this depth first. By going depth
			// first, we can hopefully find a plausible solution quickly and then
			// aggressively prune anything that can't top it.
			state := states[len(states) - 1]
			states = states[0:(len(states) - 1)]

			if state.bestPossible < best.supply["geode"] {
				continue
			}

			if state.time == MaxMinutes {
				if state.supply["geode"] > best.supply["geode"] {
					best = state
				}
				continue
			}

			visitedKey := fmt.Sprintf("%d %s %v %v", state.time, state.nextBuild, state.robots, state.supply)
			if visited[visitedKey] {
				continue
			}
			visited[visitedKey] = true

			state.time++

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

				// Calculate the best possible geode count from here. The tighter this
				// is, the more efficient the solution. But it's currently *very*
				// generous.
				timeRemaining := MaxMinutes - s.time
				s.bestPossible = s.supply["geode"] + (timeRemaining * s.robots["geode"]) + (timeRemaining * (timeRemaining + 1) / 2)

				states = append(states, s)
			}
		}

		fmt.Println(best.supply["geode"])

		total += bp.id * best.supply["geode"]
		product *= best.supply["geode"]
	}

	fmt.Printf("Quality score: %d\n", total)
	fmt.Printf("Product: %d\n", product)
}
