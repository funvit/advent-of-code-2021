package main

import (
	"flag"
	"fmt"
	"os"

	"aoc"

	_ "aoc/day1"
	_ "aoc/day10"
	_ "aoc/day11"
	_ "aoc/day12"
	_ "aoc/day13"
	_ "aoc/day14"
	_ "aoc/day15"
	_ "aoc/day16"
	_ "aoc/day2"
	_ "aoc/day3"
	_ "aoc/day4"
	_ "aoc/day5"
	_ "aoc/day6"
	_ "aoc/day7"
	_ "aoc/day8"
	_ "aoc/day9"
)

func main() {
	fmt.Println("Advent of code 2021")

	var (
		day, part uint
	)
	flag.UintVar(&day, "day", 1, "day of advent")
	flag.UintVar(&part, "part", 1, "part of day")
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage of %s:\n",
			os.Args[0],
		)
		flag.PrintDefaults()

		fmt.Fprintf(
			flag.CommandLine.Output(),
			"\nExample: \n\n\t%s -day 1 -part 1 ./day1.input.txt\n\n",
			os.Args[0],
		)
	}

	flag.Parse()

	//
	// process
	//
	inFile := flag.Arg(0)
	if inFile == "" {
		fmt.Println("ERROR: specify input data file as last arg")
		os.Exit(1)
	}

	lines, err := aoc.ReadLinesFromFile(inFile)
	if err != nil {
		fmt.Println("ERROR: read lines:", err)
		os.Exit(1)
	}

	fmt.Println("Day:", day, "Part:", part)
	fmt.Println("Lines:", len(lines))

	//
	// get solver and call it
	//
	// Note: Using registry helps to avoid large code changes in this file.
	//       Just add import like `_ "aoc/dayN"` and viola!
	sol, ok := aoc.DefaultRegistry().GetSolver(day, part)
	if !ok {
		printErrAndExit(fmt.Sprintf(
			"No solver for day %d and part %d. "+
				"Are you forgot to import day solution?",
			day, part))
	}

	r := sol(lines)
	printAnswerAndExit(r)
}

func printErrAndExit(args ...interface{}) {
	fmt.Print("ERROR ")
	fmt.Println(args...)
	os.Exit(1)
}

func printAnswerAndExit(val interface{}) {
	fmt.Println("Answer:", val)
	os.Exit(0)
}
