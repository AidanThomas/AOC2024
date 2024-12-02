package day2

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
	"strings"

	"github.com/AidanThomas/AOC2024/parser"
)

var input, parseErr = parser.ParseByLines("inputs/day2/real.txt")

func Solution() error {
	if parseErr != nil {
		return parseErr
	}
	part1()
	part2()

	return nil
}

func part1() {
	reports := slices.Collect(filter(parseReports(input), isValid))
	fmt.Printf("Part 1: %d\n", len(reports))
}

func part2() {
	reports := slices.Collect(filter(parseReports(input), isValidSingleTolerance))
	fmt.Printf("Part 2: %d\n", len(reports))
}
func parseReports(in []string) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for _, r := range in {
			report := slices.Collect(toInt(strings.Split(r, " ")))
			if !yield(report) {
				return
			}
		}
	}
}

func toInt(in []string) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, v := range in {
			val, _ := strconv.Atoi(v)
			if !yield(val) {
				return
			}
		}
	}
}

func filter(in iter.Seq[[]int], predicate func([]int) bool) iter.Seq[bool] {
	return func(yield func(bool) bool) {
		for i := range in {
			if !predicate(i) {
				continue
			}

			if !yield(true) {
				return
			}
		}
	}
}

func isValid(in []int) bool {
	diff := in[1] - in[0]
	if diff == 0 {
		return false
	}
	direction := diff > 0
	for i := 0; i < len(in)-1; i++ {
		diff := in[i+1] - in[i]
		if diff == 0 {
			return false
		}
		if (direction && diff < 0) || (!direction && diff > 0) {
			return false
		}
		if diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}

func isValidSingleTolerance(in []int) bool {
	fmt.Println(in)
	diff := in[1] - in[0]
	if diff == 0 {
		return (isValid(in[1:]))
	}
	direction := diff > 0
	nextDiff := in[2] - in[1]
	if !(direction == (nextDiff > 0)) {
		result := isValid(in[1:])
		if !result {
			out := make([]int, len(in))
			copy(out, in)
			out = append(out[:1], out[2:]...)
			return isValid(out)
		}
		return (isValid(in[1:]))
	}

	if diff == 0 {
		return (isValid(in[1:]))
	}
	for i := 0; i < len(in)-1; i++ {
		diff := in[i+1] - in[i]
		if diff == 0 {
			out := make([]int, len(in))
			copy(out, in)
			out = append(out[:i], out[i+1:]...)
			result := isValid(out)
			if !result {
				out := make([]int, len(in))
				copy(out, in)
				out = append(out[:i+1], out[i+2:]...)
				return isValid(out)
			}
			return true
		}
		if (direction && diff < 0) || (!direction && diff > 0) {
			out := make([]int, len(in))
			copy(out, in)
			out = append(out[:i], out[i+1:]...)
			result := isValid(out)
			if !result {
				out := make([]int, len(in))
				copy(out, in)
				out = append(out[:i+1], out[i+2:]...)
				return isValid(out)
			}
			return true
		}
		if diff > 3 || diff < -3 {
			out := make([]int, len(in))
			copy(out, in)
			out = append(out[:i], out[i+1:]...)
			result := isValid(out)
			if !result {
				out := make([]int, len(in))
				copy(out, in)
				out = append(out[:i+1], out[i+2:]...)
				return isValid(out)
			}
			return true
		}
	}
	return true
}
