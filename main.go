package main

import (
	"flag"
	"fmt"
)

var day = flag.Int("day", 0, "Specify which day's solution to run")

func main() {
	flag.Parse()
	if *day == 0 {
		fmt.Println("Please input a valid int to '--day'")
	}

	var call func() error
	switch *day {
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
