package aoc

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func ReadLines(reader io.Reader) ([]string, error) {

	var lines []string

	r := bufio.NewReader(reader)
	var idx int
	for {
		idx++

		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}

		lines = append(lines, string(l))
	}

	return lines, nil
}

func ReadLinesFromFile(name string) ([]string, error) {

	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer func() { _ = f.Close() }()

	return ReadLines(f)
}

func MustReadLinesFromFile(name string) []string {
	s, err := ReadLinesFromFile(name)
	if err != nil {
		panic(err)
	}

	return s
}
