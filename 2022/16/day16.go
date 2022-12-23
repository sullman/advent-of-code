package main

import (
	"fmt"
	"sort"
)

type Valve struct {
	name string
	flowRate int
}

type Agent struct {
	target string
	distance int
}

type State struct {
	time int
	agents []Agent
	released int
	releasing int
	remaining []*Valve
	bestPossible int
}

func ComputeBestPossible(state *State, valves *map[string]*Valve, maxTime int) int {
	timeRemaining := maxTime - state.time + 1
	possible := state.released + state.releasing * timeRemaining

	agentCount := 0
	for _, agent := range state.agents {
		target := (*valves)[agent.target]
		if target != nil && agent.distance < timeRemaining {
			possible += target.flowRate * (timeRemaining - agent.distance)
		}
		if agent.distance > 0 {
			agentCount++
		}
	}

	for _, v := range state.remaining {
		if timeRemaining <= 0 { break }
		agentCount++
		possible += v.flowRate * timeRemaining
		if agentCount % len(state.agents) == 0 {
			timeRemaining--
		}
	}

	return possible
}

func GenerateCombinations(all *[]*Valve, num int, permutations bool) <-chan []string {
	c := make(chan []string)

	go func(c chan []string) {
		defer close(c)

		if permutations {
			AddTarget(c, make([]string, 0), all, -1, num)
		} else {
			AddTarget(c, make([]string, 0), all, 0, num)
		}
	}(c)

	return c
}

func AddTarget(c chan []string, slice []string, all *[]*Valve, start int, numToAdd int) {
	if numToAdd == 0 {
		c <- slice
		return
	}

	// Permutations during the main process, but combinations at the start
	if start < 0 {
		visited := make(map[string]bool)
		for _, target := range slice {
			visited[target] = true
		}
		for _, v := range *all {
			if !visited[v.name] {
				newSlice := make([]string, len(slice) + 1)
				for j, str := range slice {
					newSlice[j] = str
				}
				newSlice[len(slice)] = v.name
				AddTarget(c, newSlice, all, start, numToAdd - 1)
			}
		}
	} else {
		for i := start; i <= len(*all) - numToAdd; i++ {
			newSlice := make([]string, len(slice) + 1)
			for j, str := range slice {
				newSlice[j] = str
			}
			newSlice[len(slice)] = (*all)[i].name
			AddTarget(c, newSlice, all, i + 1, numToAdd - 1)
		}
	}
}

func GenericSolve(valves *map[string]*Valve, distances *map[string]map[string]int, numAgents int, maxTime int) int {
	relevantValves := make([]*Valve, 0)
	for _, v := range *valves {
		if v.flowRate > 0 {
			relevantValves = append(relevantValves, v)
		}
	}
	sort.Slice(relevantValves, func(i, j int) bool {
		return relevantValves[i].flowRate < relevantValves[j].flowRate
	})

	best := 0
	stack := make([]*State, 1)
	initial := &State{
		time: 0,
		released: 0,
		releasing: 0,
		remaining: relevantValves,
		bestPossible: 1000000,
	}
	initial.agents = make([]Agent, numAgents)
	for i := 0; i < numAgents; i++ {
		initial.agents[i] = Agent{"AA", 0}
	}
	stack[0] = initial

	for len(stack) > 0 {
		// Using a stack makes this depth first, which will hopefully let us find a
		// plausible solution quickly and then aggressively prune anything that can't
		// top it.
		state := stack[len(stack) - 1]
		stack = stack[0:(len(stack) - 1)]

		// fmt.Printf("Evaluating t=%d agents=%v released=%d releasing=%d remaining=%v best=%d\n", state.time, state.agents, state.released, state.releasing, state.remaining, state.bestPossible)

		if state.time == maxTime + 1 {
			if state.released > best {
				// fmt.Printf("New best: %d\n", state.released)
				best = state.released
			}
			continue
		} else if state.bestPossible < best {
			continue
		}

		// Track visited states?

		state.time++
		state.released += state.releasing

		freeAgents := make([]int, 0, len(state.agents))

		for index, agent := range state.agents {
			if agent.distance > 0 {
				state.agents[index].distance--
			} else {
				state.releasing += (*valves)[agent.target].flowRate
				if len(state.remaining) == 0 {
					state.agents[index].target = "AA"
				}
				// fmt.Printf("Opening %s, now releasing %d\n", agent.target, state.releasing)
				freeAgents = append(freeAgents, index)
			}
		}

		if len(freeAgents) > 0 {
			if len(state.remaining) == 0 {
				if len(freeAgents) == numAgents {
					state.released += state.releasing * (maxTime - state.time + 1)
					state.time = maxTime + 1
				} else {
					state.bestPossible = ComputeBestPossible(state, valves, maxTime)

				}
				// state.released += state.releasing * (maxTime - state.time + 1)
				// state.time = maxTime + 1
				stack = append(stack, state)
			} else {
				for targets := range GenerateCombinations(&state.remaining, len(freeAgents), state.time != 1) {
					// fmt.Printf("Generated combo: %v\n", targets)
					targetIndexes := make(map[string]int)
					for idx, target := range targets {
						targetIndexes[target] = idx
					}
					s := &State{
						time: state.time,
						agents: make([]Agent, len(state.agents)),
						released: state.released,
						releasing: state.releasing,
						remaining: make([]*Valve, 0, len(state.remaining) - len(targets)),
					}
					for i := 0; i < len(state.agents); i++ {
						s.agents[i] = state.agents[i]
					}
					for _, v := range state.remaining {
						if agentIndex, exists := targetIndexes[v.name]; exists {
							idx := freeAgents[agentIndex]
							s.agents[idx].distance = (*distances)[s.agents[idx].target][v.name]
							s.agents[idx].target = v.name
						} else {
							s.remaining = append(s.remaining, v)
						}
					}
					s.bestPossible = ComputeBestPossible(s, valves, maxTime)
					// fmt.Printf("Appending t=%d agents=%v released=%d releasing=%d remaining=%v best=%d\n", s.time, s.agents, s.released, s.releasing, s.remaining, s.bestPossible)
					stack = append(stack, s)
				}
			}
		} else {
			state.bestPossible = ComputeBestPossible(state, valves, maxTime)
			stack = append(stack, state)
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

	fmt.Println(GenericSolve(&valves, &distances, 1, 30))
	fmt.Println(GenericSolve(&valves, &distances, 2, 26))
}
