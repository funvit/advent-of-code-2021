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

### Day 23

Damn. It is like puzzle 15, but how you gonna prefill graph?

I am lost 1 day due to wrong path calculation (hallway you know)...

Fnx to community for tips (pod only moves twice) and for cool ascii visualisation style:

```
    0 ........... DABC BADC  
    5 A.......... D.BC BADC  
   10 AA......... D.BC B.DC  
   60 AA......... D..C BBDC  
  460 AA...C..... D... BBDC  
 3460 AA...C.D... D... BB.C  
 3760 AA.....D... D... BBCC  
 4060 AA.....D.C. D... BBC.  
 8060 AA...D.D.C. .... BBC.  
 8110 AA...D.D.C. .B.. .BC.  
 8113 A....D.D.C. .B.. ABC.  
11113 A....D...C. .B.. ABCD  
11513 A....D..... .BC. ABCD  
15513 A.......... .BCD ABCD  
15516 ........... ABCD ABCD  
```

### Day 24

First I implemented ALU ... But this is too slow. 

Then I spied a solution from the community.
