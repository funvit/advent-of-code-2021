# Advent of code 2021

Single repo for all days of 2021.


## Directory struct

Dirs like `dayN` contains related day solution.

Dir `input` contains sample data and my input data for puzzles (yes, it is
differs by account).

Dir `puzzles` contains descriptions of puzzles for those who have not registered
on AoC site.

Single `./cmd/main.go` file for any day or part solution.

All solutions can be tested via:

```shell
$ make test
```

If you want to get answer for your input data use:

```shell
$ go run ./cmd/... -day X -part Y ${your.file}
```