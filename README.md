# Advent of code 2021

[![Go](https://github.com/funvit/advent-of-code-2021/actions/workflows/go.yml/badge.svg)](https://github.com/funvit/advent-of-code-2021/actions/workflows/go.yml)

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

## Notes

### Day 4

It was slow because I don't know the rules of bingo.

### Day 6

Find a simpler algorithm. Come on!

### Day 8

Fun! Digits, digits are everywhere!

### Day 14

Oh, my memory!

### Day 17

What's the trick? Oh, I see!

### Day 18

The first attempt was to use a graph. But it gets too difficult.

### Day 19

Hit me too hard in the head. 

Yes, I know about 3D. But the coordinate axes are undefined...

I looked a little for tips, but implemented it in my own way.

### Day 20

Very interesting. Infinite image, huh?

### Day 21

What? Why so easy... wait, whaaat?

### Day 22

Was an ice memory eater :) 

Remember a puzzle with 2 water cans in "Die hard" movie?