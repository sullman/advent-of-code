package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operator byte

const (
	None Operator = iota
	GT
	LT
)

type Step struct {
	field string
	operator Operator
	operand int
	destination string
}

type Part map[string]int
type Workflow []Step

func IsAccepted(workflows map[string]Workflow, part Part) bool {
	name := "in"

	for {
		w, found := workflows[name]
		if !found { return name == "A" }

		inner:
		for _, step := range w {
			switch step.operator {
			case None:
				name = step.destination
				break inner
			case GT:
				if part[step.field] > step.operand {
					name = step.destination
					break inner
				}
			case LT:
				if part[step.field] < step.operand {
					name = step.destination
					break inner
				}
			}
		}
	}
}

func Score(part Part) int {
	sum := 0
	for _, val := range part {
		sum += val
	}
	return sum
}

type Range struct { min, max int }
type Block struct {
	name string
	ranges map[string]Range
}

func FindAllRanges(workflows map[string]Workflow) int {
	sum := 0
	blocks := make([]Block, 1, len(workflows))
	blocks[0] = Block{"in", map[string]Range{
		"x": Range{1, 4000},
		"m": Range{1, 4000},
		"a": Range{1, 4000},
		"s": Range{1, 4000},
	}}

	for len(blocks) != 0 {
		block := blocks[0]
		blocks = blocks[1:]

		if block.name == "R" {
			continue
		} else if block.name == "A" {
			product := 1
			for _, r := range block.ranges {
				product *= (r.max - r.min + 1)
			}
			sum += product
			continue
		}

		w := workflows[block.name]

		for _, step := range w {
			if step.operator == None {
				block.name = step.destination
				blocks = append(blocks, block)
				break
			} else {
				var branch Block
				branch.name = step.destination
				branch.ranges = make(map[string]Range)
				for key, val := range block.ranges {
					branch.ranges[key] = val
				}

				rng := block.ranges[step.field]
				if step.operator == GT {
					if step.operand >= rng.min && step.operand < rng.max {
						block.ranges[step.field] = Range{rng.min, step.operand}
						branch.ranges[step.field] = Range{step.operand + 1, rng.max}
						blocks = append(blocks, branch)
					} else if step.operand < rng.min {
						blocks = append(blocks, branch)
						break
					}
				} else if step.operator == LT {
					if step.operand > rng.min && step.operand <= rng.max {
						block.ranges[step.field] = Range{step.operand, rng.max}
						branch.ranges[step.field] = Range{rng.min, step.operand - 1}
						blocks = append(blocks, branch)
					} else if step.operand > rng.max {
						blocks = append(blocks, branch)
						break
					}
				}
			}
		}
	}

	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	finishedSteps := false
	stepPattern := regexp.MustCompile(`(?:([xmas])([<>])(\d+):)?([a-zAR]+)`)
	fieldPattern := regexp.MustCompile(`([xmas])=(\d+)`)

	workflows := make(map[string]Workflow)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			finishedSteps = true
		} else if finishedSteps {
			part := make(Part)
			matches := fieldPattern.FindAllStringSubmatch(line, -1)

			for _, match := range matches {
				part[match[1]], _ = strconv.Atoi(match[2])
			}

			if IsAccepted(workflows, part) {
				sum += Score(part)
			}
		} else {
			name, stepsAll, _ := strings.Cut(line, "{")
			matches := stepPattern.FindAllStringSubmatch(stepsAll, -1)
			steps := make([]Step, len(matches))

			for i, match := range matches {
				field, dest := match[1], match[4]
				val := 0
				op := None
				if match[2] == "<" {
					op = LT
					val, _ = strconv.Atoi(match[3])
				} else if match[2] == ">" {
					op = GT
					val, _ = strconv.Atoi(match[3])
				}
				steps[i] = Step{field, op, val, dest}
			}

			workflows[name] = steps
		}
	}

	fmt.Printf("Part 1: %d\n", sum)

	fmt.Printf("Part 2: %d\n", FindAllRanges(workflows))
}
