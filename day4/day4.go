package day4

import (
	"fmt"
	"iter"
	"maps"

	"github.com/AidanThomas/AOC2024/parser"
)

var input, parseErr = parser.ParseByLines("inputs/day4/real.txt")

type coordinate struct {
	x int
	y int
}

func Solution() error {
	if parseErr != nil {
		return parseErr
	}
	part1()
	part2()
	return nil
}
func part1() {
	grid := maps.Collect(toGrid(input))
	fmt.Printf("Part 1: %v\n", findXmas(grid))
}
func part2() {
	grid := maps.Collect(toGrid(input))
	fmt.Printf("Part 2: %v\n", findMas(grid))
}

func toGrid(in []string) iter.Seq2[coordinate, rune] {
	return func(yield func(coordinate, rune) bool) {
		for y, l := range in {
			chars := []rune(l)
			for x, c := range chars {
				if !yield(coordinate{x, y}, c) {
					return
				}
			}
		}
	}
}

// part 1
func findXmas(grid map[coordinate]rune) int {
	var count int
	for c, v := range grid {
		if v != 'X' {
			continue
		}

		directions := []coordinate{
			{0, -1},
			{1, -1},
			{1, 0},
			{1, 1},
			{0, 1},
			{-1, 1},
			{-1, 0},
			{-1, -1},
		}
		ms, dirs := findChar(grid, c, directions, 'M')
		if len(ms) == 0 {
			continue
		}

		for i := 0; i < len(ms); i++ {
			coord := coordinate{ms[i].x + dirs[i].x, ms[i].y + dirs[i].y}
			if find(grid, coord, 'A') {
				coord.x += dirs[i].x
				coord.y += dirs[i].y
				if find(grid, coord, 'S') {
					count++
				}
			}
		}
	}
	return count
}

// part 2
func findMas(grid map[coordinate]rune) int {
	var count int
	for c, v := range grid {
		if v != 'A' {
			continue
		}
		directions := []coordinate{
			{1, -1},
			{1, 1},
			{-1, 1},
			{-1, -1},
		}
		ms, _ := findChar(grid, c, directions, 'M')
		if len(ms) != 2 {
			continue
		}
		if !(ms[0].x == ms[1].x || ms[0].y == ms[1].y) {
			continue
		}
		ss, _ := findChar(grid, c, directions, 'S')
		if len(ss) != 2 {
			continue
		}
		count++
	}
	return count
}

func findChar(grid map[coordinate]rune, tcoord coordinate, directions []coordinate, target rune) ([]coordinate, []coordinate) {
	var (
		coords []coordinate
		dirs   []coordinate
	)

	for _, d := range directions {
		newCoord := coordinate{x: tcoord.x + d.x, y: tcoord.y + d.y}
		if val, ok := grid[newCoord]; ok && val == target {
			coords = append(coords, newCoord)
			dirs = append(dirs, d)
		}
	}

	return coords, dirs
}

func find(grid map[coordinate]rune, tcoord coordinate, target rune) bool {
	if val, ok := grid[tcoord]; ok && val == target {
		return true
	}
	return false
}
