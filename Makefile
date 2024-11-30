day ?= 0

run:
	@go run main.go --day ${day}
.PHONY: run
	
new:
	@mkdir day${day}
	@echo "package day${day}; func Solution() error { return nil }" \
	> day${day}/day${day}.go
	@mkdir inputs/day${day}
	@touch inputs/day${day}/real.txt
	@touch inputs/day${day}/test.txt
.PHONY: new