day ?= 0

run:
	@go run main.go --day ${day}
.PHONY: run
	
new:
	@mkdir day${day}
	@echo 'package day${day};' \
		'import ("fmt"; "github.com/AidanThomas/AOC2024/parser");' \
		'var input, parseErr = parser.ParseByLines("inputs/day${day}/test.txt");' \
		'func Solution() error { if parseErr != nil {return parseErr}; part1(); part2(); return nil };' \
		'func part1() { fmt.Printf("Part 1: \n") };' \
		'func part2() { fmt.Printf("Part 2: \n") };' \
	> day${day}/day${day}.go
	@mkdir inputs/day${day}
	@touch inputs/day${day}/real.txt
	@touch inputs/day${day}/test.txt
.PHONY: new