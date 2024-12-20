package main

import (
	"flag"
	"fmt"

	"github.com/AidanThomas/AOC2024/day1"
	"github.com/AidanThomas/AOC2024/day2"
	"github.com/AidanThomas/AOC2024/day3"
	"github.com/AidanThomas/AOC2024/day4"
	"github.com/AidanThomas/AOC2024/day5"
)

var day = flag.Int("day", 0, "Specify which day's solution to run")

func main() {
	flag.Parse()
	if *day == 0 {
		fmt.Println("Please input a valid int to '--day'")
	}

	var call func() error
	switch *day {
	case 1:
		call = day1.Solution
	case 2:
		call = day2.Solution
	case 3:
		call = day3.Solution
	case 4:
		call = day4.Solution
	case 5:
		call = day5.Solution
	default:
		call = func() error {
			return fmt.Errorf("no day was specified")
		}
	}

	err := call()
	if err != nil {
		panic(err)
	}
}
