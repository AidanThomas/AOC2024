package day3

import (
	"fmt"
	"iter"
	"regexp"
	"strconv"
	"strings"

	"github.com/AidanThomas/AOC2024/parser"
)

var input, parseErr = parser.ParseByLines("inputs/day3/real.txt")

func Solution() error {
	if parseErr != nil {
		return parseErr
	}
	part1()
	part2()
	return nil
}

func part1() {
	in := fold(input)
	matches := regexp.MustCompile("mul\\((\\d*,\\d*)\\)").FindAllSubmatch([]byte(in), -1)
	sum := sum(multiply(extractMatch(matches)))
	fmt.Printf("Part 1: %v\n", sum)
}
func part2() {
	in := fold(input)
	matches := regexp.MustCompile("/mul\\((\\d*,\\d*)\\)|do(?:n't)?\\(\\)/g").FindAllSubmatch([]byte(in), -1)
	fmt.Printf("Part 2: \n")
}

func extractMatch(in [][][]byte) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for _, m := range in {
			numbers := strings.Split(string(m[1]), ",")
			var out []int
			for _, n := range numbers {
				number, _ := strconv.Atoi(n)
				out = append(out, number)
			}
			if !yield(out) {
				return
			}
		}
	}
}

func multiply(in iter.Seq[[]int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		for n := range in {
			if !yield(n[0] * n[1]) {
				return
			}
		}
	}
}

func sum(in iter.Seq[int]) int {
	var out int
	for n := range in {
		out += n
	}
	return out
}

func fold(in []string) string {
	var out string
	for _, s := range in {
		out += s
	}
	return out
}
