package day1

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
	"strings"

	"github.com/AidanThomas/AOC2024/parser"
)

var input, parseErr = parser.ParseByLines("inputs/day1/real.txt")

func Solution() error {
	if parseErr != nil {
		return parseErr
	}
	part1()
	part2()
	return nil
}

func part1() {
	left := slices.Sorted(getInt(input, "left"))
	right := slices.Sorted(getInt(input, "right"))

	diff := slices.Collect(diff(left, right))
	var sum int
	for _, v := range diff {
		sum += v
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func part2() {
	left := slices.Collect(getString(input, "left"))
	right := slices.Collect(getString(input, "right"))

	count := slices.Collect(count(left, right))
	var sum int
	for _, v := range count {
		sum += v
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func getInt(in []string, side string) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, v := range in {
			var val int
			switch side {
			case "left":
				val, _ = strconv.Atoi(strings.Split(v, "   ")[0])
			case "right":
				val, _ = strconv.Atoi(strings.Split(v, "   ")[1])
			}
			if !yield(val) {
				return
			}
		}
	}
}

func getString(in []string, side string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, v := range in {
			var val string
			switch side {
			case "left":
				val = strings.Split(v, "   ")[0]
			case "right":
				val = strings.Split(v, "   ")[1]
			}
			if !yield(val) {
				return
			}
		}
	}
}

func diff(left, right []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < len(left); i++ {
			diff := abs(right[i] - left[i])
			if !yield(diff) {
				return
			}
		}
	}
}

func count(left, right []string) iter.Seq[int] {
	return func(yield func(int) bool) {
		r := strings.Join(right, " ")
		for _, l := range left {
			count := strings.Count(r, l)
			li, _ := strconv.Atoi(l)
			if !yield(li * count) {
				return
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
