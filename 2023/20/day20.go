package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ModuleType uint8

const (
	Noop ModuleType = iota
	FlipFlop
	Conjunction
	Broadcaster
)

type Module struct {
	name string
	moduleType ModuleType
	state bool
	destinations []string
	sources map[string]bool
}

type Pulse struct {
	source string
	destination string
	high bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	modules := make(map[string]*Module)
	queue := make([]Pulse, 0, 1024)
	numLow, numHigh := 0, 0

	ensureModule := func(name string) *Module {
		m := modules[name]
		if m == nil {
			m = new(Module)
			m.name = name
			m.destinations = make([]string, 0, 1)
			m.sources = make(map[string]bool)
			modules[name] = m
		}
		return m
	}
	trigger := func(pulse Pulse) {
		mod := modules[pulse.destination]
		switch mod.moduleType {
		case FlipFlop:
			if !pulse.high {
				mod.state = !mod.state
				for _, dest := range mod.destinations {
					queue = append(queue, Pulse{mod.name, dest, mod.state})
				}
			}
		case Conjunction:
			mod.sources[pulse.source] = pulse.high
			allHigh := true
			for _, high := range mod.sources {
				if !high {
					allHigh = false
					break
				}
			}
			for _, dest := range mod.destinations {
				queue = append(queue, Pulse{mod.name, dest, !allHigh})
			}
		case Broadcaster:
			for _, dest := range mod.destinations {
				queue = append(queue, Pulse{mod.name, dest, pulse.high})
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 { break }
		before, after, _ := strings.Cut(line, " -> ")
		destinations := strings.Split(after, ", ")

		name, moduleType := before, Noop
		if name[0] == '%' {
			name, moduleType = name[1:], FlipFlop
		} else if name[0] == '&' {
			name, moduleType = name[1:], Conjunction
		} else if name[0] == 'b' {
			moduleType = Broadcaster
		}

		mod := ensureModule(name)
		mod.moduleType = moduleType
		mod.destinations = destinations

		for _, dest := range destinations {
			ensureModule(dest).sources[name] = false
		}
	}

	for i := 0; i < 1000; i++ {
		queue = append(queue, Pulse{"button", "broadcaster", false})

		for len(queue) != 0 {
			p := queue[0]
			queue = queue[1:]
			if p.high {
				numHigh++
			} else {
				numLow++
			}
			trigger(p)
		}
	}

	fmt.Printf("Part 1: %d\n", numLow * numHigh)
}
