package day5

import (
	"fmt"
	"iter"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/AidanThomas/AOC2024/parser"
)

var (
	input, parseErr = parser.ParseByLines("inputs/day5/test.txt")
	index           = getSeparatorIndex(input)
	instructions    = parseInstructions(input[:index])
	updates         = parseUpdates(input[index+1:])
)

func Solution() error {
	if parseErr != nil {
		return parseErr
	}
	part1()
	part2() // note - order is transitive - wip
	return nil
}

func part1() {
	var total int
	for u := range updates {
		if executeInstructions(u) {
			mid := int(math.Floor(float64(len(u) / 2)))
			total += u[mid]
		}
	}

	fmt.Printf("Part 1: %d\n", total)
}

func part2() {
	var total int
	for u := range updates {
		if !executeInstructions(u) {
			sorted := sortUpdate(u)
			mid := int(math.Floor(float64(len(sorted) / 2)))
			total += sorted[mid]
		}
	}
	fmt.Printf("Part 2: %d\n", total)
}

func getSeparatorIndex(in []string) int {
	for i, s := range in {
		if s == "" {
			return i
		}
	}
	return 0
}

func parseInstructions(in []string) map[int][]int {
	out := make(map[int][]int)

	for _, s := range in {
		numbers := toInts(strings.Split(s, "|"))

		if slice, ok := out[numbers[0]]; ok {
			out[numbers[0]] = append(slice, numbers[1])
			continue
		}

		out[numbers[0]] = []int{numbers[1]}
	}

	return out
}

func parseUpdates(in []string) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for _, s := range in {
			numbers := toInts(strings.Split(s, ","))
			if !yield(numbers) {
				return
			}
		}
	}
}

// part 1
func executeInstructions(update []int) bool {
	for i, u := range update {
		if rules, ok := instructions[u]; ok {
			for _, r := range rules {
				if slices.Contains(update[:i], r) {
					return false
				}
			}
		}
	}

	return true
}

// part 2
func sortUpdate(update []int) []int {
	out := makeCopy(update)
	for !executeInstructions(out) {
		out = sortOnce(out)
	}

	fmt.Println(out)
	return out

	// 97 13 75 29 47
	// append beginning
	// append i
	// append j
	// append end
	// 97 75 13 29 47
	// 97 75 29 13 47
	// 97 75 29 47 13
	// potentially need to then recheck it recursively
}

func sortOnce(update []int) []int {
	out := makeCopy(update)
	for i, u := range update {
		if i == 0 {
			continue
		}
		if rules, ok := instructions[u]; ok {
			for _, r := range rules {
				if j := slices.Index(update[:i], r); j != -1 {
					beginning := makeCopy(update[:j])
					bad := update[j]
					end := makeCopy(update[i+1:])
					out = make([]int, 0)
					out = append(out, beginning...)
					out = append(out, u)
					out = append(out, bad)
					out = append(out, end...)
					return out
				}
			}
		}
	}
	return out
}

func toInts(in []string) []int {
	var out []int
	for _, s := range in {
		n, _ := strconv.Atoi(s)
		out = append(out, n)
	}
	return out
}

func makeCopy(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	return out
}
