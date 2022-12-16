package main

import (
	"container/list"
	"fmt"
	"strings"
)

type Valve struct {
	name string
	flowRate int
	tunnels []string
}

type State struct {
	time int
	valve string
	released int
	releasing int
	open string
	maxPossible int
}

const Minutes = 30
var BestYet = 0

func InsertSorted(l *list.List, state *State, visited map[string]int) {
	key := fmt.Sprintf("%s %s", state.valve, state.open)
	prev := visited[key]
	if prev > 0 && prev <= state.time {
		// fmt.Printf("Not revisiting %s %d < %d\n", key, prev, state.time)
		return
	}
	visited[key] = state.time
	// fmt.Printf("Going to visit %s\n", key)
	if state.released > BestYet {
		fmt.Printf("New best yet %d %s\n", state.released, state.open)
		BestYet = state.released
	}

	for e := l.Front(); e != nil; e = e.Next() {
		other := e.Value.(*State)
		if state.maxPossible >= other.maxPossible {
			l.InsertBefore(state, e)
			return
		}
	}

	l.PushBack(state)
}

func Dump(l *list.List, prefix string) {
	fmt.Printf("%s ", prefix)
	for e := l.Front(); e != nil; e = e.Next() {
		state := e.Value.(*State)
		fmt.Printf("%s %s;", state.valve, state.open)
	}
	fmt.Println()
}

func main() {
	valves := make(map[string]*Valve)
	totalRate := 0
	initialValve := ""

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
			tunnels: make([]string, 0),
		}
		valves[name] = valve
		totalRate += rate

		if initialValve == "" {
			initialValve = name
		}

		for sep == ',' {
			fmt.Scanf("%2s%c", &next, &sep)
			valve.tunnels = append(valve.tunnels, next)
		}
	}

	initialValve = "AA"
	visited := make(map[string]int)
	candidates := list.New()
	initial := &State{
		time: 0,
		valve: initialValve,
		released: 0,
		releasing: 0,
		open: "",
		maxPossible: Minutes * totalRate,
	}
	candidates.PushFront(initial)

	for candidates.Len() > 0 {
		// Dump(candidates, "Before:")
		e := candidates.Front()
		state := candidates.Remove(e).(*State)
		// fmt.Printf("%d %v\n", len(visited), state)

		if state.time == Minutes + 1 {
			fmt.Println(state.released)
			candidates.Init()
			break
		}

		if state.releasing == totalRate {
			total := state.released + (state.releasing * (Minutes - state.time))
			next := &State{
				time: Minutes + 1,
				valve: state.valve + "!",
				released: total,
				releasing: state.releasing,
				open: state.open,
				maxPossible: total,
			}
			InsertSorted(candidates, next, visited)
			continue
		}

		valve := valves[state.valve]

		if valve.flowRate > 0 && !strings.Contains(state.open, state.valve) {
			next := &State{
				time: state.time + 1,
				valve: state.valve,
				released: state.released + state.releasing,
				releasing: state.releasing + valve.flowRate,
				open: state.open + "," + state.valve,
				maxPossible: state.released + state.releasing + (Minutes - state.time - 1) * totalRate,
			}
			InsertSorted(candidates, next, visited)
		}

		for _, nextValve := range valve.tunnels {
			next := &State{
				time: state.time + 1,
				valve: nextValve,
				released: state.released + state.releasing,
				releasing: state.releasing,
				open: state.open,
				maxPossible: state.released + state.releasing + (Minutes - state.time - 1) * totalRate,
			}
			InsertSorted(candidates, next, visited)
		}
		// Dump(candidates, "After:")
	}
}
