day ?= 0

run:
	@go run main.go --day ${day}
.PHONY: run
	
new:
	@mkdir day${day}
	@echo 'package day${day};\n' \
		'import ("fmt"; "github.com/AidanThomas/AOC2024/parser")\n' \
		'var input, parseErr = parser.ParseByLines("inputs/day3/test.txt")\n' \
		'func Solution() error { if parseErr != nil {return parseErr}; part1(); part2(); return nil }\n' \
		'func part1() { fmt.Printf("Part 1: \\n") }\n' \
		'func part2() { fmt.Printf("Part 2: \\n") }\n' \
	> day${day}/day${day}.go
	@mkdir inputs/day${day}
	@touch inputs/day${day}/real.txt
	@touch inputs/day${day}/test.txt
.PHONY: new